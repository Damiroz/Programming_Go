package service

type Task struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    DueDate     string `json:"due_date"`
    Done        bool   `json:"done"`
}

type TaskService struct {
    tasks map[string]Task
}

func NewTaskService() *TaskService {
    return &TaskService{
        tasks: make(map[string]Task),
    }
}

func (s *TaskService) Create(t Task) Task {
    s.tasks[t.ID] = t
    return t
}

func (s *TaskService) GetAll() []Task {
    result := []Task{}
    for _, t := range s.tasks {
        result = append(result, t)
    }
    return result
}
