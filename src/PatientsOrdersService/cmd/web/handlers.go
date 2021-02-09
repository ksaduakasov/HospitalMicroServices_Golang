package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Fring02/HospitalMicroservices/src/PatientsOrdersService/pkg/models"
	"github.com/Fring02/HospitalMicroservices/src/PatientsOrdersService/pkg/repositories"
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
	"strconv"
)

func(a *Application) Routes() http.Handler {
	standardMiddleware := alice.New()
	dynamicMiddleware := alice.New()
	mux := pat.New()
	mux.Get("/patients", dynamicMiddleware.ThenFunc(getAllPatients))
	mux.Get("/patients/:id", dynamicMiddleware.ThenFunc(getPatient))
	mux.Post("/patients/create", dynamicMiddleware.ThenFunc(createPatient))
	mux.Del("/patients/delete/:id", dynamicMiddleware.ThenFunc(deletePatient))
	mux.Put("/patients/update/:id", dynamicMiddleware.ThenFunc(updatePatient))
	return standardMiddleware.Then(mux)
}

func updatePatient(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT"{
		App.MethodNotAllowedError(w, nil)
	}
	//trying to parse id
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil{
		App.ServerError(w, err)
		return
	}
	//getting patient from db by id
	patient, err := repositories.GetPatientById(conn, id)
	if err == sql.ErrNoRows{
		App.NotFoundError(w, err)
		return
	} else if err != nil {
		App.ServerError(w, err)
		return
	}
	//deserialize patient from request json body
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&patient)
	if err != nil {
		App.BadRequestError(w, err)
		return
	}
	//trying to update patient...
	updated, err := repositories.UpdatePatient(conn, *patient, id)
	if err != nil {
		App.BadRequestError(w, err)
		return
	} else if !updated {
		App.ServerError(w, err)
		return
	}
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Updated patient with id %v", id)))
}

func deletePatient(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE"{
		App.MethodNotAllowedError(w, nil)
	}
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil{
		App.ServerError(w, err)
		return
	}
	deleted, err := repositories.DeletePatient(conn, id)
	if err == sql.ErrNoRows{
		App.BadRequestError(w, err)
		return
	} else if err != nil || !deleted {
		App.ServerError(w, err)
		return
	}
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Deleted patient with id %v", id)))
}

func createPatient(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST"{
		App.MethodNotAllowedError(w, nil)
	}
	decoder := json.NewDecoder(r.Body)
	var patient models.Patient
	err := decoder.Decode(&patient)
	if err != nil {
		App.BadRequestError(w, err)
		return
	}
	id, err := repositories.CreatePatient(conn, patient)
	if err != nil {
		App.BadRequestError(w, err)
		return
	}
	w.WriteHeader(201)
	w.Write([]byte(fmt.Sprintf("Created patient with id %v", id)))
}

func getPatient(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET"{
		App.MethodNotAllowedError(w, nil)
	}
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil{
		App.ServerError(w, err)
		return
	}
	patient, err := repositories.GetPatientById(conn, id)
	if err == sql.ErrNoRows{
		App.NotFoundError(w, err)
		return
	} else if err != nil {
		App.ServerError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(patient)
	if err != nil {
		App.ServerError(w, err)
		return
	}
}

func getAllPatients(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET"{
		App.MethodNotAllowedError(w, nil)
	}
	patients, err := repositories.GetAllPatients(conn)
	if err == sql.ErrNoRows{
		App.NotFoundError(w, err)
		return
	} else if err != nil {
		App.ServerError(w, err)
		return
	}
	if patients == nil{
		patients = []*models.Patient{}
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(patients)
	if err != nil {
		App.ServerError(w, err)
		return
	}
}
