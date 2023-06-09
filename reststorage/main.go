package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化 Gin Web框架
	r := gin.Default()
	// 注册 API 路由
	r.Any("resource/:type", restHandler)
	//启动 Web 服务器
	_ = r.Run(":8080")
}

// Task 结构体
type Task struct {
	Name        string
	Description string
}

// Service 服务
type Service struct {
	Name        string
	Description string
}

type TaskRegistry interface {
	ListTasks() ([]Task, error)
	CreateTask(task Task) error
}

type MockTaskRegistry struct {
}

func (t *MockTaskRegistry) ListTasks() ([]Task, error) {
	return []Task{
		{
			Name:        "testMockTaskRegistry",
			Description: "testMockTaskRegistry",
		},
	}, nil
}

func (t *MockTaskRegistry) CreateTask(task Task) error {
	return nil
}

type MysqlTaskRegistry struct {
}

func (t *MysqlTaskRegistry) ListTasks() ([]Task, error) {
	return []Task{
		{
			Name:        "testMysqlTaskRegistry",
			Description: "testMysqlTaskRegistry",
		},
	}, nil
}

func (t *MysqlTaskRegistry) CreateTask(task Task) error {
	return nil
}

type ServiceRegistry struct {
}

func (s *ServiceRegistry) ListServices() ([]Service, error) {
	return []Service{
		{
			Name:        "testService",
			Description: "testService",
		},
	}, nil
}

func (s *ServiceRegistry) CreateService(service Service) error {
	return nil
}

type handlerStorage interface {
	List(c *gin.Context)
	Create(c *gin.Context)
}

type TaskStorage struct {
}

func (t *TaskStorage) List(c *gin.Context) {
	// 这里需要使用哪个 Registry 直接修改即可
	taskRegistry := MysqlTaskRegistry{}
	tasks, err := taskRegistry.ListTasks()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to list tasks"})
		return
	}
	c.JSON(200, tasks)
}

func (t *TaskStorage) Create(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}
	// 这里需要使用哪个 Registry 直接修改即可
	taskRegistry := MysqlTaskRegistry{}
	err := taskRegistry.CreateTask(task)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to createTask"})
		return
	}
	c.JSON(200, gin.H{"message": "Task created successfully"})
}

type ServiceStorage struct {
}

func (s *ServiceStorage) List(c *gin.Context) {
	serviceRegistry := ServiceRegistry{}
	tasks, err := serviceRegistry.ListServices()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to list tasks"})
		return
	}
	c.JSON(200, tasks)
}

func (s *ServiceStorage) Create(c *gin.Context) {
	var service Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}
	serviceRegistry := ServiceRegistry{}
	err := serviceRegistry.CreateService(service)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to createTask"})
		return
	}
	c.JSON(200, gin.H{"message": "Task created successfully"})
}

func restHandler(c *gin.Context) {
	m := map[string]handlerStorage{
		"task":    &TaskStorage{},
		"service": &ServiceStorage{},
	}
	resourceType := c.Param("type")
	switch c.Request.Method {
	case "GET":
		m[resourceType].List(c)
	case "POST":
		m[resourceType].Create(c)
	}
}
