package handlers

import (
	"errors"
	"net/http"
	"task-manager/internal/models"
	"task-manager/internal/storage"
	"task-manager/pkg/utils"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	storage storage.TaskStorage
}

func NewTaskHandler(storage storage.TaskStorage) *TaskHandler {
	return &TaskHandler{
		storage: storage,
	}
}

// CreateTask handles task creation
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := h.storage.Create(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) GetTask(c *gin.Context) {
	id := c.Param("id")

	task, err := h.storage.GetByID(id)
	if err != nil {
		if errors.Is(err, storage.ErrTaskNotFound) {
			utils.NotFound(c, "Task")
			return
		}
		utils.InternalError(c, "Failed to retrieve task")
		return
	}

	utils.Success(c, task, "Task retrieved successfully")
}

func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.storage.GetAll()
	if err != nil {
		utils.InternalError(c, "Failed to retrieve tasks")
		return
	}

	utils.Success(c, tasks, "Tasks retrieved successfully")
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var update models.TaskUpdate
	if err := c.ShouldBindJSON(&update); err != nil {
		utils.BadRequest(c, "Invalid request payload")
		return
	}

	task, err := h.storage.Update(id, &update)
	if err != nil {
		if errors.Is(err, storage.ErrTaskNotFound) {
			utils.NotFound(c, "Task")
			return
		}
		utils.InternalError(c, "Failed to update task")
		return
	}

	utils.Success(c, task, "Task updated successfully")
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	if err := h.storage.Delete(id); err != nil {
		if errors.Is(err, storage.ErrTaskNotFound) {
			utils.NotFound(c, "Task")
			return
		}
		utils.InternalError(c, "Failed to delete task")
		return
	}

	utils.Success(c, nil, "Task deleted successfully")
}

func (h *TaskHandler) GetTasksByStatus(c *gin.Context) {
	status := c.Query("status")
	if status == "" {
		utils.BadRequest(c, "Status query parameter is required")
		return
	}

	tasks, err := h.storage.GetByStatus(status)
	if err != nil {
		utils.InternalError(c, "Failed to retrieve tasks")
		return
	}

	utils.Success(c, tasks, "Tasks retrieved successfully")
}
