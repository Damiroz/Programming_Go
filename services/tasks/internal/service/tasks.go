package service

import (
    "errors"
    "fmt"
    "sync"
    "time"
)

type Task struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description,omitempty"`
    DueDate     string `json:"due_date,omitempty"`
    Done        bool   `json:"done"`
}

type CreateTaskRequest struct {
    Title       string `json:"title"`
    Description string `json:"description"`
    DueDate     string `json:"due_date"`
}

type UpdateTaskRequest struct {
    Title *string `json:"title,omitempty"`
    Done  *bool   `json:"done,omitempty"`
}

type TasksService struct {
    mu    sync.RWMutex
    tasks map[string]Task
    last  int
}

func NewTasksService() *TasksService {
    return &TasksService{
        tasks: make(map[string]Task),
    }
}

func (s *TasksService) Create(req CreateTaskRequest) (Task, error) {
    if req.Title == "" {
        return Task{}, errors.New("title is required")
    }

    s.mu.Lock()
    defer s.mu.Unlock()

    s.last++
    id := generateID(s.last)

    task := Task{
        ID:          id,
        Title:       req.Title,
        Description: req.Description,
        DueDate:     req.DueDate,
        Done:        false,
    }

    s.tasks[id] = task
    return task, nil
}

func (s *TasksService) List() []Task {
    s.mu.RLock()
    defer s.mu.RUnlock()

    result := make([]Task, 0, len(s.tasks))
    for _, t := range s.tasks {
        result = append(result, t)
    }
    return result
}

func (s *TasksService) Get(id string) (Task, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()

    t, ok := s.tasks[id]
    return t, ok
}

func (s *TasksService) Update(id string, req UpdateTaskRequest) (Task, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    t, ok := s.tasks[id]
    if !ok {
        return Task{}, errors.New("not found")
    }

    if req.Title != nil {
        t.Title = *req.Title
    }
    if req.Done != nil {
        t.Done = *req.Done
    }

    s.tasks[id] = t
    return t, nil
}

func (s *TasksService) Delete(id string) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    if _, ok := s.tasks[id]; !ok {
        return errors.New("not found")
    }
    delete(s.tasks, id)
    return nil
}

func generateID(n int) string {
    return "t_" + time.Now().Format("20060102150405") + "_" + fmtInt(n)
}

func fmtInt(n int) string {
    return fmt.Sprintf("%03d", n)
}
