package main

import (
	"fmt"
	"github.com/Alemkhan/HospitalMicroservices/DepartmentMicroservice/core"
	"github.com/Alemkhan/HospitalMicroservices/DepartmentMicroservice/pkg/repositories"
	"github.com/gin-gonic/gin"
	"strconv"
)

var departmentRepository repositories.DepartmentsRepository

func RouteDepartments(router *gin.Engine)  {

	router.GET("/departments", GetDepartments)
	router.GET("/departments/:id", GetDepartmentByID)
	router.POST("/departments", CreateDepartment)
	router.DELETE("/departments/:id", DeleteDepartment)
	router.PUT("/departments/:id", UpdateDepartment)

}

func GetDepartments(c *gin.Context)  {
	departments := departmentRepository.GetDepartments()
	c.JSON(200, departments)
}

func GetDepartmentByID(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect format"))
	}
	department := departmentRepository.GetDepartmentByID(id)
	c.JSON(200, department)
}

func CreateDepartment(c *gin.Context)  {

	department := &core.Department{}
	err := c.BindJSON(department)
	if err != nil {
		c.Data(400, jsonContentType, []byte("Fill all fields"))
	}
	id, err := departmentRepository.CreateDepartment(*department)
	if err != nil {
		c.Data(500, jsonContentType, []byte("Failed to create a department"))
	} else {
		c.Data(200, jsonContentType, []byte(fmt.Sprintf("The department with id %v has been successfully created", id)))
	}

}

func DeleteDepartment(c *gin.Context)  {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect format"))
	}
	department := departmentRepository.GetDepartmentByID(id)
	if department == nil {
		c.Data(400, jsonContentType, []byte("No such department in database"))
		return
	}
	_, err = departmentRepository.DeleteDepartment(department.ID)
	if err != nil {
		c.Data(500, jsonContentType, []byte("Failed to delete the department"))
	} else {
		c.Data(200, jsonContentType, []byte(fmt.Sprintf("The department with id %v has been successfully deleted", id)))
	}

}

func UpdateDepartment(c *gin.Context)  {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect format"))
	}
	department := departmentRepository.GetDepartmentByID(id)

	if department == nil {
		c.Data(400, jsonContentType, []byte("There is no such department"))
		return
	}

	updatedDepartment := &core.Department{}
	err = c.BindJSON(updatedDepartment)
	if err != nil {
		c.Data(400, jsonContentType, []byte("Fill all fields"))
		return
	}
	updatedDepartment.ID = id
	_, err = departmentRepository.UpdateDepartment(*updatedDepartment)
	if err != nil {
		c.Data(500, jsonContentType, []byte("Failed to update the department"))
		return
	} else {
		c.Data(200, jsonContentType, []byte(fmt.Sprintf("The department with id %v has been successfully updated", id)))
	}

}