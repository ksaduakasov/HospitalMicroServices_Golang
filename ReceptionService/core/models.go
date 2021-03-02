package core

type Order struct {
 Id int
 Title string
 Description string
 DiseaseId int
 PatientId int
}

type Department struct {
 Id int
 Name string
 Description string
 DiseaseId string
}