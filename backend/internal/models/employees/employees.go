package employees

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	http_lib "organizational-structure-visualization-service/internal/lib/http"
	"organizational-structure-visualization-service/pkg/logger/sl"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Employee struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	MiddleName  string    `json:"middle_name"`
	LastName    string    `json:"last_name"`
	Position    string    `json:"position,omitempty"`
	Department  string    `json:"department,omitempty"`
	Subdivision string    `json:"subdivision,omitempty"`
	Role        string    `json:"role,omitempty"`
	Project     string    `json:"project,omitempty"`
	City        string    `json:"city,omitempty"`
	HireDate    time.Time `json:"hire_date,omitempty"`
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

	// Запрос без фильтрации
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
		// Сканируем результат запроса
		if err := rows.Scan(&e.ID, &e.FirstName, &e.MiddleName, &e.LastName, &e.Position, &e.Department, &e.Role, &e.Project, &e.City); err != nil {
			log.Error("failed to scan employee", sl.Err(err))
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		employees = append(employees, e)
	}

	// Проверка на ошибки при чтении
	if err := rows.Err(); err != nil {
		log.Error("failed to read rows", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return employees, nil
}
