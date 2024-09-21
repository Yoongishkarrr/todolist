package task

import (
	"fmt"
	"sync"
)

type Task struct {
	ID          int
	Description string
}

var (
	taskIDCounter int
	taskList      []*Task
	mutex         sync.Mutex
)

func GetAllTasks() []*Task {
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Printf("Current tasks count: %d\n", len(taskList))
	return taskList
}

func AddTask(newTask *Task) {
	mutex.Lock()
	defer mutex.Unlock()
	taskList = append(taskList, newTask) // Используйте указатель на новую задачу
}

func NewTask(description string) *Task {
	mutex.Lock()
	defer mutex.Unlock()
	taskIDCounter++
	task := &Task{
		ID:          taskIDCounter,
		Description: description,
	}
	return task // Возвращаем новую задачу без добавления в список
}

func UpdateTask(id int, newDescription string) *Task {
	mutex.Lock()
	defer mutex.Unlock()
	for _, task := range taskList {
		if task.ID == id {
			task.Description = newDescription
			return task
		}
	}
	return nil
}

func DeleteTask(id int) bool {
	mutex.Lock()
	defer mutex.Unlock()
	for i, task := range taskList {
		if task.ID == id {
			taskList = append(taskList[:i], taskList[i+1:]...)
			return true
		}
	}
	return false
}

func (t *Task) String() string {
	return fmt.Sprintf("ID: %d, Description: %s", t.ID, t.Description)
}
