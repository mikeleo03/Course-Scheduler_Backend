package algorithm

import (
	"github.com/mikeleo03/Course-Scheduler_Backend/models"
)

func ScheduleCourses(matakuliah []models.MataKuliah, creditLimit int) (float64, []models.MataKuliah) {
	dp := make([][]float64, len(matakuliah)+1)
	selectedCourses := make([][][]models.MataKuliah, len(matakuliah)+1)

	for i := range dp {
		dp[i] = make([]float64, creditLimit+1)
		selectedCourses[i] = make([][]models.MataKuliah, creditLimit+1)
	}

	for i := 1; i <= len(matakuliah); i++ {
		for j := 1; j <= creditLimit; j++ {
			if matakuliah[i-1].SKS <= j {
				score := getScore(matakuliah[i-1].Prediksi, matakuliah[i-1].SKS)
				if dp[i-1][j-matakuliah[i-1].SKS]+score > dp[i-1][j] {
					dp[i][j] = dp[i-1][j-matakuliah[i-1].SKS] + score
					selectedCourses[i][j] = append(selectedCourses[i-1][j-matakuliah[i-1].SKS][:0:0], selectedCourses[i-1][j-matakuliah[i-1].SKS]...)
					selectedCourses[i][j] = append(selectedCourses[i][j], matakuliah[i-1])
				} else {
					dp[i][j] = dp[i-1][j]
					selectedCourses[i][j] = append(selectedCourses[i-1][j][:0:0], selectedCourses[i-1][j]...)
				}
			} else {
				dp[i][j] = dp[i-1][j]
				selectedCourses[i][j] = append(selectedCourses[i-1][j][:0:0], selectedCourses[i-1][j]...)
			}
		}
	}

	maxScore := dp[len(matakuliah)][creditLimit]
	selected := selectedCourses[len(matakuliah)][creditLimit]

	// Calculating total sks
	totalSKS := CalculateTotalSKS(selected)
	ip := maxScore / float64(totalSKS)

	return ip, selected
}

func getScore(prediction string, SKS int) float64 {
	scores := map[string]float64{
		"A":  4,
		"AB": 3.5,
		"B":  3,
		"BC": 2.5,
		"C":  2,
		"D":  1,
		"E":  0,
	}

	return scores[prediction] * float64(SKS)
}

// CalculateTotalSKS calculates the total SKS value from a list of selected MataKuliah
func CalculateTotalSKS(selectedMataKuliah []models.MataKuliah) int {
	totalSKS := 0
	for _, matkul := range selectedMataKuliah {
		totalSKS += matkul.SKS
	}
	return totalSKS
}