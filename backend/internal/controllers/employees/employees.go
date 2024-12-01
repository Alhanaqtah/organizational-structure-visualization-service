package employees

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	http_lib "organizational-structure-visualization-service/internal/lib/http"
	model "organizational-structure-visualization-service/internal/models/employees"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Model interface {
	GetAll(ctx context.Context, offset, limit int) ([]model.Employee, error)
	GetAllPositions(ctx context.Context) ([]string, error)
	GetAllDepartments(ctx context.Context) ([]string, error)
	GetAllSubdivisions(ctx context.Context) ([]string, error)
	GetAllRoles(ctx context.Context) ([]string, error)
	GetAllProjects(ctx context.Context) ([]string, error)
	GetAllCities(ctx context.Context) ([]string, error)
	GetByID(ctx context.Context, id string) (*model.Employee, error)
}

type Controller struct {
	model Model
}

type ControllerConfig struct {
	Model Model
}

type employeesResponse struct {
	Total     int              `json:"total"`
	Page      int              `json:"page"`
	Limit     int              `json:"limit"`
	Employees []model.Employee `json:"employees"`
}

func New(cfg ControllerConfig) *Controller {
	return &Controller{
		model: cfg.Model,
	}
}

func (c *Controller) Register() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", c.GetAll)
	r.Get("/{id}", c.GetByID)
	r.Get("/filters", c.GetFilters)

	return r
}

func (c *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	const op = "controller.employees.GetAll"

	log := http_lib.GetCtxLogger(r.Context())
	log = log.With(slog.String("op", op))

	filters := make(model.Filters)
	c.parseGetAllRequest(r, filters)

	page, limit := c.parsePagination(r)

	empls, err := c.model.GetAll(r.Context(), (page-1)*limit, limit)
	if err != nil {
		http_lib.ErrInternal(w, r)
		return
	}

	filteredEmployees := filterEmployees(empls, filters)

	filteredTotal := len(filteredEmployees)

	response := employeesResponse{
		Total:     filteredTotal,
		Page:      page,
		Limit:     limit,
		Employees: filteredEmployees,
	}

	render.JSON(w, r, response)
}

func (c *Controller) GetFilters(w http.ResponseWriter, r *http.Request) {
	const op = "controller.employees.GetFilters"

	log := http_lib.GetCtxLogger(r.Context())
	log = log.With(slog.String("op", op))

	positions, err := c.model.GetAllPositions(r.Context())
	if err != nil {
		log.Error("failed to fetch positions", slog.Any("error", err))
		http_lib.ErrInternal(w, r)
		return
	}

	departments, err := c.model.GetAllDepartments(r.Context())
	if err != nil {
		log.Error("failed to fetch departments", slog.Any("error", err))
		http_lib.ErrInternal(w, r)
		return
	}

	subdivisions, err := c.model.GetAllSubdivisions(r.Context())
	if err != nil {
		log.Error("failed to fetch subdivisions", slog.Any("error", err))
		http_lib.ErrInternal(w, r)
		return
	}

	roles, err := c.model.GetAllRoles(r.Context())
	if err != nil {
		log.Error("failed to fetch roles", slog.Any("error", err))
		http_lib.ErrInternal(w, r)
		return
	}

	projects, err := c.model.GetAllProjects(r.Context())
	if err != nil {
		log.Error("failed to fetch projects", slog.Any("error", err))
		http_lib.ErrInternal(w, r)
		return
	}

	cities, err := c.model.GetAllCities(r.Context())
	if err != nil {
		log.Error("failed to fetch cities", slog.Any("error", err))
		http_lib.ErrInternal(w, r)
		return
	}

	filters := []map[string]interface{}{
		{
			"filter": "Должности",
			"values": positions,
		},
		{
			"filter": "Департаменты",
			"values": departments,
		},
		{
			"filter": "Подразделения",
			"values": subdivisions,
		},
		{
			"filter": "Роли",
			"values": roles,
		},
		{
			"filter": "Проекты",
			"values": projects,
		},
		{
			"filter": "Города",
			"values": cities,
		},
	}

	render.JSON(w, r, filters)
}

func (c *Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	const op = "controller.employees.GetByID"

	log := http_lib.GetCtxLogger(r.Context())
	log = log.With(slog.String("op", op))

	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http_lib.ErrBadRequest(w, r)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http_lib.ErrBadRequest(w, r)
		return
	}

	// Получение сотрудника из модели
	employee, err := c.model.GetByID(r.Context(), idStr)
	if err != nil {
		if err == model.ErrNotFound {
			http_lib.ErrNotFound(w, r)
			return
		}

		log.Error("failed to fetch employee", slog.Any("error", err))
		http_lib.ErrInternal(w, r)
		return
	}

	render.JSON(w, r, employee)
}

func (c *Controller) parseGetAllRequest(r *http.Request, filters model.Filters) {
	params := r.URL.Query()

	for key, values := range params {
		switch key {
		case "position":
			filters["position"] = values
		case "department":
			filters["department"] = values
		case "subdivision":
			filters["subdivision"] = values
		case "role":
			filters["role"] = values
		case "project":
			filters["project"] = values
		case "city":
			filters["city"] = values
		case "first_name_search":
			filters["first_name_search"] = values
		case "middle_name_search":
			filters["middle_name_search"] = values
		case "last_last_name":
			filters["last_name_search"] = values
		}
	}
}

func (c *Controller) parsePagination(r *http.Request) (int, int) {
	page := 1
	limit := 20

	if p := r.URL.Query().Get("page"); p != "" {
		parsedPage, err := strconv.Atoi(p)
		if err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	if l := r.URL.Query().Get("limit"); l != "" {
		parsedLimit, err := strconv.Atoi(l)
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	return page, limit
}

// Фильтрация сотрудников в коде
func filterEmployees(employees []model.Employee, filters model.Filters) []model.Employee {
	var filtered []model.Employee

	for _, employee := range employees {
		include := true

		if position, ok := filters["position"]; ok && len(position) > 0 && !contains(position, employee.Position) {
			include = false
		}
		if department, ok := filters["department"]; ok && len(department) > 0 && !contains(department, employee.Department) {
			include = false
		}
		if role, ok := filters["role"]; ok && len(role) > 0 && !contains(role, employee.Role) {
			include = false
		}
		if project, ok := filters["project"]; ok && len(project) > 0 && !contains(project, employee.Project) {
			include = false
		}
		if city, ok := filters["city"]; ok && len(city) > 0 && !contains(city, employee.City) {
			include = false
		}
		if firstNameSearch, ok := filters["first_name_search"]; ok && len(firstNameSearch) > 0 && !containsLike(firstNameSearch, employee.FirstName) {
			include = false
		}
		if middleNameSearch, ok := filters["middle_name_search"]; ok && len(middleNameSearch) > 0 && !containsLike(middleNameSearch, employee.MiddleName) {
			include = false
		}
		if lastNameSearch, ok := filters["last_name_search"]; ok && len(lastNameSearch) > 0 && !containsLike(lastNameSearch, employee.LastName) {
			include = false
		}

		if include {
			filtered = append(filtered, employee)
		}
	}

	return filtered
}

func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func containsLike(slice []string, value string) bool {
	for _, item := range slice {
		if strings.HasPrefix(strings.ToLower(value), strings.ToLower(item)) {
			return true
		}
	}
	return false
}
