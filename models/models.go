package models

// Data mata kuliah yang ditambahkan
type MataKuliahAdd struct {
	ID    		int  		`json:"id"`
	Nama		string 		`json:"nama"`
	SKS			int 		`json:"sks"`
	Jurusan     string 		`json:"jurusan"`
	SemMin		int			`json:"semmin"`
	Prediksi 	string 		`json:"prediksi"`
}

// Data mata kuliah
type MataKuliah struct {
	ID    		int  		`json:"id"`
	Nama		string 		`json:"nama"`
	SKS			int 		`json:"sks"`
	Jurusan     string 		`json:"jurusan"`
	Fakultas	string 		`json:"fakultas"`
	SemMin		int			`json:"semmin"`
	Prediksi 	string 		`json:"prediksi"`
}

// Data fakultas
type Fakultas struct {
	ID		    int       	`json:"id"`
	Nama		string		`json:"nama"`
}

// Data jurusan
type Jurusan struct {
	ID		  	int       	`json:"id"`
	Nama		string		`json:"nama"`
	Fakultas	string 		`json:"fakultas"`
}

// Data request
type Request struct {
	Jurusan    		string 		`json:"jurusan"`
	Semester		int			`json:"semester"`
	SKS				int 		`json:"sks"`
}