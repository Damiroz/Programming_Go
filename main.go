package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:admin@l127.0.0.1:8000/todo?sslmode=disable"
	}

	db, err := openDB(dsn)
	if err != nil {
		log.Fatalf("openDB error: %v", err)
	}
	defer db.Close()

	repo := NewRepo(db)

	// Вставим пару задач
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	titles := []string{"Сделать ПЗ №5", "Купить кофе", "Проверить отчёты"}
	for _, title := range titles {
		id, err := repo.CreateTask(ctx, title)
		if err != nil {
			log.Fatalf("CreateTask error: %v", err)
		}
		log.Printf("Inserted task id=%d (%s)", id, title)
	}

	// Прочитаем список задач
	ctxList, cancelList := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelList()

	tasks, err := repo.ListTasks(ctxList)
	if err != nil {
		log.Fatalf("ListTasks error: %v", err)
	}

	// Напечатаем
	fmt.Println("=== Tasks ===")
	for _, t := range tasks {
		fmt.Printf("#%d | %-24s | done=%-5v | %s\n",
			t.ID, t.Title, t.Done, t.CreatedAt.Format(time.RFC3339))
	}
	// Тестирование FindByID (Проверочное задание 2)
fmt.Println("\n--- Find Task by ID (ID=1) ---")
ctxFind, cancelFind := context.WithTimeout(context.Background(), 2*time.Second)
defer cancelFind()

// Попробуем найти задачу с ID=1 
foundTask, err := repo.FindByID(ctxFind, 1)
if err != nil {
	log.Fatalf("FindByID error: %v", err)
}

if foundTask != nil {
	fmt.Printf("Found Task #%d: Title='%s', Done=%v\n", 
		foundTask.ID, foundTask.Title, foundTask.Done)
} else {
	fmt.Println("Task with ID=1 not found.")
}

// Тестирование CreateMany (Проверочное задание 3)
fmt.Println("\n--- Bulk Insert via Transaction ---")
ctxMany, cancelMany := context.WithTimeout(context.Background(), 5*time.Second)
defer cancelMany()

bulkTitles := []string{"Задача из транзакции 1", "Задача из транзакции 2"}
err = repo.CreateMany(ctxMany, bulkTitles)
if err != nil {
	log.Fatalf("CreateMany error: %v", err)
}
log.Println("Successfully inserted bulk tasks via transaction.")

// естирование ListDone (Проверочное задание 1)
fmt.Println("\n--- List Done/Not Done Tasks ---")
ctxListDone, cancelListDone := context.WithTimeout(context.Background(), 3*time.Second)
defer cancelListDone()

// Вставляем одну выполненную задачу для теста
repo.CreateTask(ctxListDone, "Завершенная задача (тест)")
repo.DB.ExecContext(ctxListDone, `UPDATE tasks SET done = TRUE WHERE title = $1`, "Завершенная задача (тест)")

// Задачи, которые НЕ выполнены
notDone, err := repo.ListDone(ctxListDone, false)
if err != nil {
	log.Fatalf("ListDone(false) error: %v", err)
}
fmt.Printf("Total NOT DONE tasks: %d\n", len(notDone))

// Задачи, которые ВЫПОЛНЕНЫ
doneTasks, err := repo.ListDone(ctxListDone, true)
if err != nil {
	log.Fatalf("ListDone(true) error: %v", err)
}
fmt.Printf("Total DONE tasks: %d (Example: %s)\n", len(doneTasks), doneTasks[0].Title)

// Настройки пула (Проверочное задание 4)
log.Println("\n--- Connection Pool Settings ---")
log.Println("SetMaxOpenConns(10), SetMaxIdleConns(5), SetConnMaxLifetime(30 min) chosen for local development/testing.")
log.Println("Эти настройки обеспечивают быстрый доступ к соединениям без перегрузки локального сервера БД.")
}
