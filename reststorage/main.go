package main

import (
	"github.com/gin-gonic/gin"
)

// 创建任务的 API 处理函数
func createTaskHandler(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}
	taskRegistry := TaskRegistry{}
	err := taskRegistry.CreateTask(task)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create task"})
		return
	}
	c.JSON(201, gin.H{"message": "Task created successfully"})
}

// 获取任务列表的API处理函数
func listTasksHandler(c *gin.Context) {
	taskRegistry := TaskRegistry{}
	tasks, err := taskRegistry.ListTasks()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to listtasks"})
		return
	}
	c.JSON(200, tasks)
}

// 展示服务的 API 处理函数
func listServicesHandler(c *gin.Context) {
	c.JSON(200, nil)
}

// 创建服务的 API 处理函数
func createServicesHandler(c *gin.Context) {
	c.JSON(200, nil)
}

func main() {
	// 初始化 Gin Web框架
	r := gin.Default()
	// 注册 API 路由
	r.POST("/tasks", createTaskHandler)
	r.POST("/services", createServicesHandler)
	r.GET("/tasks", listTasksHandler)
	r.GET("/services", listServicesHandler)
	//启动 Web 服务器
	_ = r.Run(":8080")
}

// Task 结构体
type Task struct {
	Name        string
	Description string
}

type TaskRegistry struct {
}

func (t *TaskRegistry) ListTasks() ([]Task, error) {
	return []Task{
		{
			Name:        "test",
			Description: "test",
		},
	}, nil
}

func (t *TaskRegistry) CreateTask(task Task) error {
	return nil
}
