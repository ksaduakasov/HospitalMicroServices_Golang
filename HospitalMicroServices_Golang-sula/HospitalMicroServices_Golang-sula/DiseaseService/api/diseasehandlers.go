package api

import (
	"fmt"
	"github.com/Fring02/HospitalMicroservices/DiseaseService/core"
	"github.com/Fring02/HospitalMicroservices/DiseaseService/pkg/repositories"
	"github.com/gin-gonic/gin"
	"strconv"
)

var DiseaseRepository repositories.DiseaseRepository

func RouteDiseases(router *gin.Engine) {
	router.GET("/diseases", GetDiseases)
	router.GET("/diseases/:id", GetDiseaseByID)
	router.POST("/diseases", CreateDisease)
	router.DELETE("/diseases/:id", DeleteDisease)
	router.PUT("/diseases/:id", UpdateDisease)
}

func GetDiseases(c *gin.Context) {
	diseases := DiseaseRepository.GetDiseases()
	c.JSON(200, diseases)
}

func GetDiseaseByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect format"))
	}
	disease := DiseaseRepository.GetDiseaseByID(id)
	c.JSON(200, disease)
}

func CreateDisease(c *gin.Context) {

	disease := &core.Disease{}
	err := c.BindJSON(disease)

	if err != nil {
		c.Data(400, jsonContentType, []byte("Fill all fields"))
	}

	exist := DiseaseRepository.CheckForDisease(*disease)

	if exist == false {
		id, err := DiseaseRepository.CreateDisease(*disease)
		if err != nil {
			c.Data(500, jsonContentType, []byte("Failed to create a disease"))
			return
		} else {
			c.Data(200, jsonContentType, []byte(fmt.Sprintf("The disease with id %v has been successfully created", id)))
		}
	} else {
		c.Data(400, jsonContentType, []byte("The disease already exits"))
		return
	}
}

func DeleteDisease(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect format"))
	}
	disease := DiseaseRepository.GetDiseaseByID(id)
	if disease == nil {
		c.Data(400, jsonContentType, []byte("No such disease in database"))
		return
	}
	_, err = DiseaseRepository.DeleteDisease(disease.ID)
	if err != nil {
		c.Data(500, jsonContentType, []byte("Failed to delete the disease"))
	} else {
		c.Data(200, jsonContentType, []byte(fmt.Sprintf("The department with id %v has been successfully deleted", id)))
	}

}

func UpdateDisease(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect format"))
	}
	disease := DiseaseRepository.GetDiseaseByID(id)

	if disease == nil {
		c.Data(400, jsonContentType, []byte("There is no such disease"))
		return
	}

	updatedDisease := &core.Disease{}
	err = c.BindJSON(updatedDisease)
	if err != nil {
		c.Data(400, jsonContentType, []byte("Fill all fields"))
		return
	}
	updatedDisease.ID = id
	_, err = DiseaseRepository.UpdateDisease(*updatedDisease)
	if err != nil {
		c.Data(500, jsonContentType, []byte("Failed to update the disease"))
		return
	} else {
		c.Data(200, jsonContentType, []byte(fmt.Sprintf("The disease with id %v has been successfully updated", id)))
	}

}


