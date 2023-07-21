package middleware

import (
	"encoding/json"
	"log"
	"strconv"
	"net/http"

	"github.com/mikeleo03/Course-Scheduler_Backend/models"
	"github.com/mikeleo03/Course-Scheduler_Backend/repository"
	"github.com/mikeleo03/Course-Scheduler_Backend/algorithm"
)

// response format
type response struct {
	Row     int64  `json:"row,omitempty"`
	Message string `json:"message,omitempty"`
}

type ScheduleResponse struct {
	Status     bool  				`json:"status,omitempty"`
	Message    string 				`json:"message,omitempty"`
	Value      []models.MataKuliah  `json:"value,omitempty"`
	Total      float64              `json:"total,omitempty"`
	SKS        int              	`json:"sks,omitempty"`
}

// Create a global instance of Repo
var repo *repository.Repo

// SetRepo sets the instance of Repo
func SetRepo(r *repository.Repo) {
	repo = r
}

// AddMataKuliah create a mata kuliah in the postgres db
func AddMataKuliah(w http.ResponseWriter, r *http.Request) {
	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	// w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty matkul of type models.MataKuliahAdd
	var matkul models.MataKuliahAdd

	// decode the json request to matkul
	if (r.Method == "POST") {
		err := json.NewDecoder(r.Body).Decode(&matkul)
		if err != nil {
			log.Fatalf("Unable to decode the request body.  %v", err)
		}

		// Gather the fakultas data
		fakultasData, err := repo.GetAllFakul()
		if err != nil {
			log.Fatalf("Unable to get all fakul. %v", err)
		}

		// call insert matkul function and pass the matkul
		rowCount := repo.InsertMataKuliah(matkul, fakultasData)

		// format a response object
		res := response {
			Row:   rowCount,
			Message: "Matkul created successfully",
		}

		// send the response
		json.NewEncoder(w).Encode(res)
	}
}

// AddFakultas create a new fakultas in the postgres db
func AddFakultas(w http.ResponseWriter, r *http.Request) {
	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	// w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty fakul of type models.Fakultas
	var fakul models.Fakultas

	// decode the json request to fakul
	if (r.Method == "POST") {
		err := json.NewDecoder(r.Body).Decode(&fakul)
		if err != nil {
			log.Fatalf("Unable to decode the request body.  %v", err)
		}

		// call insert fakul function and pass the fakul
		rowCount := repo.InsertFakultas(fakul)

		// format a response object
		res := response {
			Row:   rowCount,
			Message: "Fakultas created successfully",
		}

		// send the response
		json.NewEncoder(w).Encode(res)
	}
}

// AddJurusan create a new jurusan in the postgres db
func AddJurusan(w http.ResponseWriter, r *http.Request) {
	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	// w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty jurusan of type models.Jurusan
	var jurusan models.Jurusan

	// decode the json request to jurusan
	if (r.Method == "POST") {
		err := json.NewDecoder(r.Body).Decode(&jurusan)
		if err != nil {
			log.Fatalf("Unable to decode the request body.  %v", err)
		}

		// call insert jurusan function and pass the jurusan
		rowCount := repo.InsertJurusan(jurusan)

		// format a response object
		res := response {
			Row:   rowCount,
			Message: "Jurusan created successfully",
		}

		// send the response
		json.NewEncoder(w).Encode(res)
	}
}

// GetAllMataKuliah will return all the matkul
func GetAllMataKuliah(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get all the matkul in the db
	matkuls, err := repo.GetAllMatkul()
	if err != nil {
		log.Fatalf("Unable to get all matkul. %v", err)
	}

	// send all the users as response
	json.NewEncoder(w).Encode(matkuls)
}

// GetAllFakultas will return all the matkul
func GetAllFakultas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get all the matkul in the db
	fakuls, err := repo.GetAllFakul()
	if err != nil {
		log.Fatalf("Unable to get all fakul. %v", err)
	}

	// send all the users as response
	json.NewEncoder(w).Encode(fakuls)
}

// GetAllJurusan will return all the jurusan
func GetAllJurusan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get all the matkul in the db
	jurus, err := repo.GetAllJurus()
	if err != nil {
		log.Fatalf("Unable to get all jurusan. %v", err)
	}

	// send all the users as response
	json.NewEncoder(w).Encode(jurus)
}

// GenerateSchedule generates a schedule based on the provided data
func GenerateSchedule(w http.ResponseWriter, r *http.Request) {
	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	// w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	/* // create an empty request struct to receive the data
	var request models.Request */

	// decode the json request to the request struct
	if r.Method == "GET" {
		// Parse the URL parameters
		jurusan := r.URL.Query().Get("jurusan")
		semester := r.URL.Query().Get("semester")
		sksmin := r.URL.Query().Get("sksmin")
		sksmax := r.URL.Query().Get("sksmax")

		// get all the matkul data based on spesification
		sem, err := strconv.Atoi(semester)
		if err != nil {
			// Handle error if the string cannot be converted to int
			log.Fatalf("Error convertting. %v", err)
			return
		}
		mataKuliahList, err := repo.GetMatkulData(jurusan, sem)
		if err != nil {
			log.Fatalf("Unable to retrieve matkul data. %v", err)
		}

		// Check if there is no data after filtering process
		if len(mataKuliahList) == 0 {
			res := ScheduleResponse {
				Status:  false,
				Message: "Tidak terdapat mata kuliah yang memenuhi kondisi. Silahkan ubah data.",
			}
			// send the response
			json.NewEncoder(w).Encode(res)
			return
		}

		// Apply the algorithm to the filtered mata kuliah data
		sksminval, err := strconv.Atoi(sksmin)
		if err != nil {
			// Handle error if the string cannot be converted to int
			log.Fatalf("Error converting. %v", err)
			return
		}

		sksmaxval, err := strconv.Atoi(sksmax)
		if err != nil {
			// Handle error if the string cannot be converted to int
			log.Fatalf("Error converting. %v", err)
			return
		}

		// Initialize variables to keep track of the maximum scheduledMataKuliah and its total
		var maxScheduledMataKuliah []models.MataKuliah
		var maxTotal float64
		var maxSKS int

		// Iterate from sksmin to sksmax (inclusive)
		for sks := sksminval; sks <= sksmaxval; sks++ {
			// Apply the algorithm to the mataKuliahList with current sks value
			total, scheduledMataKuliah := algorithm.ScheduleCourses(mataKuliahList, sks)

			// Check if the current scheduledMataKuliah has a higher total than the previous maximum
			if total > maxTotal {
				maxTotal = total
				maxScheduledMataKuliah = scheduledMataKuliah
				maxSKS = sks
			}
		}
		
		// Return the result to the frontend
		res := ScheduleResponse{
			Status:  true,
			Value:   maxScheduledMataKuliah,
			Total:   maxTotal,
			SKS: 	 maxSKS,
			Message: "",
		}
		// send the response
		json.NewEncoder(w).Encode(res)
	}
}