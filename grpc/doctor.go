package hospitalpb

type Doctor struct {
	ID               int
	Firstname        string
	Lastname         string
	Patronymic       string
	Phone            string
	Email            string
	Description      string
	WorkExp          int
	DepartmentID     int
	IsAvailable      bool
	DoctorMultiplier float64
}

type Department struct {
	ID          int
	Name        string
	Description string
	DiseaseID   int
}
