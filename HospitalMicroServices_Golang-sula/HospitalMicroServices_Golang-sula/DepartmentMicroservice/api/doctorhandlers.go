package api

import (
	"fmt"
	"github.com/Fring02/HospitalMicroservices/DepartmentMicroservice/core"
	"github.com/Fring02/HospitalMicroservices/DepartmentMicroservice/pkg/repositories"
	"github.com/gin-gonic/gin"
	"strconv"
)

var DoctorRepository repositories.DoctorRepository

func RouteDoctors(router *gin.Engine) {

	router.GET("/doctors", GetDoctors)
	router.GET("/doctors/:id", GetDoctorByID)
	router.POST("/doctors", CreateDoctor)
	router.DELETE("/doctors/:id", DeleteDoctor)
	router.PUT("/doctors/:id", UpdateDoctor)

}

func GetDoctors(c *gin.Context) {
	doctors := DoctorRepository.GetDoctors()
	c.JSON(200, doctors)
}

func GetDoctorByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect format"))
	}
	doctor := DoctorRepository.GetDoctorByID(id)
	c.JSON(200, doctor)
}

func CreateDoctor(c *gin.Context) {

	doctor := &core.Doctor{}
	err := c.BindJSON(doctor)
	if err != nil {
		c.Data(400, jsonContentType, []byte("Fill all fields"))
	}
	id, err := DoctorRepository.CreateDoctor(*doctor)
	if err != nil {
		c.Data(500, jsonContentType, []byte("Failed to create a doctor"))
	} else {
		c.Data(200, jsonContentType, []byte(fmt.Sprintf("The doctor with id %v has been successfully created", id)))
	}

}

func DeleteDoctor(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect format"))
	}
	doctor := DoctorRepository.GetDoctorByID(id)
	if doctor == nil {
		c.Data(400, jsonContentType, []byte("No such doctor in database"))
		return
	}
	_, err = DoctorRepository.DeleteDoctor(doctor.ID)
	if err != nil {
		c.Data(500, jsonContentType, []byte("Failed to delete the doctor"))
	} else {
		c.Data(200, jsonContentType, []byte(fmt.Sprintf("The doctor with id %v has been successfully deleted", id)))
	}

}

func UpdateDoctor(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect format"))
	}
	doctor := DoctorRepository.GetDoctorByID(id)

	if doctor == nil {
		c.Data(400, jsonContentType, []byte("There is no such doctor"))
		return
	}

	updatedDoctor := &core.Doctor{}
	err = c.BindJSON(updatedDoctor)
	if err != nil {
		c.Data(400, jsonContentType, []byte("Fill all fields"))
		return
	}
	updatedDoctor.ID = id
	_, err = DoctorRepository.UpdateDoctor(*updatedDoctor)
	if err != nil {
		c.Data(500, jsonContentType, []byte("Failed to update the doctor"))
		return
	} else {
		c.Data(200, jsonContentType, []byte(fmt.Sprintf("The doctor with id %v has been successfully updated", id)))
	}

}
