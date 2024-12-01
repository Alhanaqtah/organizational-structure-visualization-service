package employees

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"

	http_lib "organizational-structure-visualization-service/internal/lib/http"
	"organizational-structure-visualization-service/pkg/logger/sl"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrNotFound = errors.New("not found")
)

type Employee struct {
	ID           int        `json:"id"`
	FirstName    string     `json:"first_name"`
	MiddleName   string     `json:"middle_name"`
	LastName     string     `json:"last_name"`
	Position     string     `json:"position,omitempty"`
	Department   string     `json:"department,omitempty"`
	Subdivision  string     `json:"subdivision,omitempty"`
	Role         string     `json:"role,omitempty"`
	Project      string     `json:"project,omitempty"`
	City         string     `json:"city,omitempty"`
	HireDate     time.Time  `json:"hire_date,omitempty"`
	HeadID       int        `json:"-"`
	Subordinates []Employee `json:"subordinates,omitempty"`
	Colleagues   []Employee `json:"colleagues,omitempty"`
	Managers     []Employee `json:"managers,omitempty"`
}

type Filters map[string][]string

type Model struct {
	pool *pgxpool.Pool
}

func NewModel(pool *pgxpool.Pool) *Model {
	return &Model{
		pool: pool,
	}
}

func (m *Model) GetAll(ctx context.Context, offset, limit int) ([]Employee, error) {
	const op = "model.employees.GetAll"

	log := http_lib.GetCtxLogger(ctx)
	log = slog.With(slog.String("op", op))

	baseQuery := `
		SELECT e.id, e.first_name, e.middle_name, e.last_name, 
		       p.title AS position, d.title AS department, r.title AS role, pr.title AS project, o.city
		FROM public.employees e
		JOIN public.positions p ON e.position_id = p.id
		JOIN public.divisions d ON e.division_id = d.id
		JOIN public.departments dept ON d.department_id = dept.id
		JOIN public.roles r ON e.role_id = r.id
		JOIN public.projects pr ON e.project_id = pr.id
		JOIN public.offices o ON dept.office_id = o.id
		ORDER BY e.id LIMIT $1 OFFSET $2`

	rows, err := m.pool.Query(ctx, baseQuery, limit, offset)
	if err != nil {
		log.Error("failed to get list of employees", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var e Employee
		if err := rows.Scan(&e.ID, &e.FirstName, &e.MiddleName, &e.LastName, &e.Position, &e.Department, &e.Role, &e.Project, &e.City); err != nil {
			log.Error("failed to scan employee", sl.Err(err))
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		employees = append(employees, e)
	}

	if err := rows.Err(); err != nil {
		log.Error("failed to read rows", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return employees, nil
}

// GetAllPositions возвращает уникальные должности
func (m *Model) GetAllPositions(ctx context.Context) ([]string, error) {
	return m.getDistinctValues(ctx, "title", "positions")
}

// GetAllDepartments возвращает уникальные департаменты
func (m *Model) GetAllDepartments(ctx context.Context) ([]string, error) {
	return m.getDistinctValues(ctx, "title", "departments")
}

// GetAllSubdivisions возвращает уникальные подразделения
func (m *Model) GetAllSubdivisions(ctx context.Context) ([]string, error) {
	return m.getDistinctValues(ctx, "title", "divisions")
}

// GetAllRoles возвращает уникальные роли
func (m *Model) GetAllRoles(ctx context.Context) ([]string, error) {
	return m.getDistinctValues(ctx, "title", "roles")
}

// GetAllProjects возвращает уникальные проекты
func (m *Model) GetAllProjects(ctx context.Context) ([]string, error) {
	return m.getDistinctValues(ctx, "title", "projects")
}

// GetAllCities возвращает уникальные города
func (m *Model) GetAllCities(ctx context.Context) ([]string, error) {
	return m.getDistinctValues(ctx, "city", "offices")
}

// Общий метод для получения уникальных значений из указанной таблицы
func (m *Model) getDistinctValues(ctx context.Context, column, table string) ([]string, error) {
	query := fmt.Sprintf("SELECT DISTINCT %s FROM %s WHERE %s IS NOT NULL ORDER BY %s", column, table, column, column)
	rows, err := m.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var values []string
	for rows.Next() {
		var value string
		if err := rows.Scan(&value); err != nil {
			return nil, err
		}
		values = append(values, value)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return values, nil
}

func (m *Model) GetByID(ctx context.Context, id string) (*Employee, error) {
	const op = "model.employees.GetByID"

	// Проверка валидности ID
	if id == "" {
		return nil, fmt.Errorf("%s: id is required", op)
	}

	query := `
		SELECT e.id, e.first_name, e.middle_name, e.last_name, 
		       p.title AS position, d.title AS department, r.title AS role, pr.title AS project, o.city
		FROM public.employees e
		JOIN public.positions p ON e.position_id = p.id
		JOIN public.divisions d ON e.division_id = d.id
		JOIN public.departments dept ON d.department_id = dept.id
		JOIN public.roles r ON e.role_id = r.id
		JOIN public.projects pr ON e.project_id = pr.id
		JOIN public.offices o ON dept.office_id = o.id
		WHERE 
			e.id = $1
	`

	var employee Employee
	err := m.pool.QueryRow(ctx, query, id).Scan(
		&employee.ID,
		&employee.FirstName,
		&employee.MiddleName,
		&employee.LastName,
		&employee.Position,
		&employee.Department,
		&employee.Role,
		&employee.Project,
		&employee.City,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("%s: employee not found: %w", op, ErrNotFound)
		}
		return nil, fmt.Errorf("%s: failed to query employee: %w", op, err)
	}

	return &employee, nil
}

func (m *Model) GetTree(ctx context.Context, id string) (*Employee, error) {
	const op = "model.employees.GetTree"

	log := http_lib.GetCtxLogger(ctx)
	log = slog.With(slog.String("op", op))

	var headID sql.NullInt64
	// Получаем информацию о сотруднике
	queryEmployee := `SELECT 
        e.id,
        e.first_name,
        e.middle_name,
        e.last_name,
        p.title AS position,
        d.title AS department,
        s.title AS subdivision,
        r.title AS role,
        pr.title AS project,
		e.head_id,
        o.city
    FROM 
        public.employees e
    LEFT JOIN public.positions p ON e.position_id = p.id
    LEFT JOIN public.divisions s ON e.division_id = s.id
    LEFT JOIN public.departments d ON s.department_id = d.id
    LEFT JOIN public.roles r ON e.role_id = r.id
    LEFT JOIN public.projects pr ON e.project_id = pr.id
    LEFT JOIN public.offices o ON o.id = d.office_id
    WHERE 
        e.id = $1`

	row := m.pool.QueryRow(ctx, queryEmployee, id)

	var employee Employee
	if err := row.Scan(
		&employee.ID,
		&employee.FirstName,
		&employee.MiddleName,
		&employee.LastName,
		&employee.Position,
		&employee.Department,
		&employee.Subdivision,
		&employee.Role,
		&employee.Project,
		&headID,
		&employee.City,
	); err != nil {
		log.Error("failed to get employee", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if headID.Valid {
		employee.HeadID = int(headID.Int64)
	}

	var head Employee
	var subordinates []Employee
	var colleagues []Employee

	// Получаем информацию о руководителе
	if headID.Valid {
		queryHead := `SELECT
		e.ID,
        e.first_name,
        e.middle_name,
        e.last_name,
        p.title AS position,
        d.title AS department,
        s.title AS subdivision,
        r.title AS role,
        pr.title AS project,
        o.city
    FROM 
        public.employees e
    LEFT JOIN public.positions p ON e.position_id = p.id
    LEFT JOIN public.divisions s ON e.division_id = s.id
    LEFT JOIN public.departments d ON s.department_id = d.id
    LEFT JOIN public.roles r ON e.role_id = r.id
    LEFT JOIN public.projects pr ON e.project_id = pr.id
    LEFT JOIN public.offices o ON o.id = d.office_id
    WHERE e.id = $1`

		rowHead := m.pool.QueryRow(ctx, queryHead, employee.HeadID)

		if err := rowHead.Scan(
			&head.ID,
			&head.FirstName,
			&head.MiddleName,
			&head.LastName,
			&head.Position,
			&head.Department,
			&head.Subdivision,
			&head.Role,
			&head.Project,
			&head.City,
		); err != nil {
			log.Error("failed to get head", sl.Err(err))
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		// Получаем коллег
		queryColleagues := `SELECT 
        e.id,
        e.first_name,
        e.middle_name,
        e.last_name,
        p.title AS position,
        d.title AS department,
        s.title AS subdivision,
        r.title AS role,
        pr.title AS project,
        o.city
    FROM 
        public.employees e
    LEFT JOIN public.positions p ON e.position_id = p.id
    LEFT JOIN public.divisions s ON e.division_id = s.id
    LEFT JOIN public.departments d ON s.department_id = d.id
    LEFT JOIN public.roles r ON e.role_id = r.id
    LEFT JOIN public.projects pr ON e.project_id = pr.id
    LEFT JOIN public.offices o ON o.id = d.office_id
    WHERE e.head_id = $1 AND e.id <> $2`

		rowsColleagues, err := m.pool.Query(ctx, queryColleagues, employee.HeadID, employee.ID)
		if err != nil {
			log.Error("failed to get colleagues", sl.Err(err))
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		defer rowsColleagues.Close()

		for rowsColleagues.Next() {
			var colleague Employee
			if err := rowsColleagues.Scan(
				&colleague.ID,
				&colleague.FirstName,
				&colleague.MiddleName,
				&colleague.LastName,
				&colleague.Position,
				&colleague.Department,
				&colleague.Subdivision,
				&colleague.Role,
				&colleague.Project,
				&colleague.City,
			); err != nil {
				log.Error("failed to scan colleague", sl.Err(err))
				return nil, fmt.Errorf("%s: %w", op, err)
			}
			colleagues = append(colleagues, colleague)
		}
		if err := rowsColleagues.Err(); err != nil {
			log.Error("failed to iterate over colleagues", sl.Err(err))
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	// Получаем подчинённых
	querySubordinates := `SELECT 
        e.id,
        e.first_name,
        e.middle_name,
        e.last_name,
        p.title AS position,
        d.title AS department,
        s.title AS subdivision,
        r.title AS role,
        pr.title AS project,
        o.city
    FROM 
        public.employees e
    LEFT JOIN public.positions p ON e.position_id = p.id
    LEFT JOIN public.divisions s ON e.division_id = s.id
    LEFT JOIN public.departments d ON s.department_id = d.id
    LEFT JOIN public.roles r ON e.role_id = r.id
    LEFT JOIN public.projects pr ON e.project_id = pr.id
    LEFT JOIN public.offices o ON o.id = d.office_id
    WHERE e.head_id = $1`

	rowsSubordinates, err := m.pool.Query(ctx, querySubordinates, employee.ID)
	if err != nil {
		log.Error("failed to get subordinates", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rowsSubordinates.Close()

	for rowsSubordinates.Next() {
		var subordinate Employee
		if err := rowsSubordinates.Scan(
			&subordinate.ID,
			&subordinate.FirstName,
			&subordinate.MiddleName,
			&subordinate.LastName,
			&subordinate.Position,
			&subordinate.Department,
			&subordinate.Subdivision,
			&subordinate.Role,
			&subordinate.Project,
			&subordinate.City,
		); err != nil {
			log.Error("failed to scan subordinate", sl.Err(err))
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		subordinates = append(subordinates, subordinate)
	}
	if err := rowsSubordinates.Err(); err != nil {
		log.Error("failed to iterate over subordinates", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	employee.Managers = append(employee.Managers, head)
	employee.Colleagues = colleagues
	employee.Subordinates = subordinates

	return &employee, nil
}
