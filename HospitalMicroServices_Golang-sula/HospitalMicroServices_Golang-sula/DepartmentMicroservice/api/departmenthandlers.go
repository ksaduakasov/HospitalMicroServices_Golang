package api

import (
	"fmt"
	"github.com/Fring02/HospitalMicroservices/DepartmentMicroservice/core"
	"github.com/Fring02/HospitalMicroservices/DepartmentMicroservice/pkg/repositories"
	"github.com/gin-gonic/gin"
	"strconv"
)

var DepartmentRepository repositories.DepartmentsRepository

func RouteDepartments(router *gin.Engine) {
	router.GET("/departments", GetDepartments)
	router.GET("/departments/:id", GetDepartmentByID)
	router.GET("/departments/:id/disease", GetDepartmentsByDiseaseId)
	router.POST("/departments", CreateDepartment)
	router.DELETE("/departments/:id", DeleteDepartment)
	router.PUT("/departments/:id", UpdateDepartment)
}

func GetDepartments(c *gin.Context) {
	departments := DepartmentRepository.GetDepartments()
	c.JSON(200, departments)
}
func GetDepartmentsByDiseaseId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect format"))
	}
	departments := DepartmentRepository.GetDepartmentsByDiseaseId(id)
	c.JSON(200, departments)
}
func GetDepartmentByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect format"))
	}
	department := DepartmentRepository.GetDepartmentByID(id)
	c.JSON(200, department)
}

func CreateDepartment(c *gin.Context) {

	department := &core.Department{}
	err := c.BindJSON(department)
	if err != nil {
		c.Data(400, jsonContentType, []byte("Fill all fields"))
	}
	id, err := DepartmentRepository.CreateDepartment(*department)
	if err != nil {
		c.Data(500, jsonContentType, []byte("Failed to create a department"))
	} else {
		c.Data(200, jsonContentType, []byte(fmt.Sprintf("The department with id %v has been successfully created", id)))
	}

}

func DeleteDepartment(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect format"))
	}
	department := DepartmentRepository.GetDepartmentByID(id)
	if department == nil {
		c.Data(400, jsonContentType, []byte("No such department in database"))
		return
	}
	_, err = DepartmentRepository.DeleteDepartment(department.ID)
	if err != nil {
		c.Data(500, jsonContentType, []byte("Failed to delete the department"))
	} else {
		c.Data(200, jsonContentType, []byte(fmt.Sprintf("The department with id %v has been successfully deleted", id)))
	}

}

func UpdateDepartment(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect format"))
	}
	department := DepartmentRepository.GetDepartmentByID(id)

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
	_, err = DepartmentRepository.UpdateDepartment(*updatedDepartment)
	if err != nil {
		c.Data(500, jsonContentType, []byte("Failed to update the department"))
		return
	} else {
		c.Data(200, jsonContentType, []byte(fmt.Sprintf("The department with id %v has been successfully updated", id)))
	}

}
