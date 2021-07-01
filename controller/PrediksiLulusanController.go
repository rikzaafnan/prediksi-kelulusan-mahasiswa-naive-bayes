package controller

import (
	"math"
	"net/http"
	"prediksi-kelulusan-mahasiswa-naive-bayes/config"
	"prediksi-kelulusan-mahasiswa-naive-bayes/models"

	"github.com/labstack/echo/v4"
)

// func totalDataStatusKelulusan(keteranganStatusKelulusan string) ([]models.ProfileLulusan, error) {

// 	var profileLulusan []models.ProfileLulusan

// 	err := config.DB.Where("status_kelulusan = ?", keteranganStatusKelulusan).Find(&profileLulusan).Error
// 	if err != nil {
// 		return profileLulusan, err
// 	}

// 	return profileLulusan, nil
// }

func validTotalData(profileLulusan []models.ProfileLulusan, totalDataLulusan int) float64 {

	// validTotalData := math.Round((float64(len(profileLulusan))/float64(41))*100) / 100
	validTotalData := math.Round((float64(len(profileLulusan))/float64(totalDataLulusan))*100) / 100

	return validTotalData
}

func peluangTotalDataJenisKelamin(lulusanTepat []models.ProfileLulusan, lulusanTerlambat []models.ProfileLulusan, jenisKelamin string) (peluangTepatGender int, peluangTerlambatGender int) {

	peluangTepatGender = 0
	peluangTerlambatGender = 0

	dataPeluangTepat := []models.ProfileLulusan{}
	dataPeluangTerlambat := []models.ProfileLulusan{}

	// tepat
	if len(lulusanTepat) > 0 {

		for _, detailLulusanTepat := range lulusanTepat {

			if detailLulusanTepat.JenisKelamin == jenisKelamin {

				dataPeluangTepat = append(dataPeluangTepat, detailLulusanTepat)

			}

		}

		peluangTepatGender = len(dataPeluangTepat)

	}

	if len(lulusanTerlambat) > 0 {
		for _, detailLulusanTerlambat := range lulusanTerlambat {

			if detailLulusanTerlambat.JenisKelamin == jenisKelamin {

				dataPeluangTerlambat = append(dataPeluangTepat, detailLulusanTerlambat)

			}

		}

		peluangTerlambatGender = len(dataPeluangTerlambat)
	}

	return peluangTepatGender, peluangTerlambatGender

}

func peluangTotalDataStatusMahasiswa(lulusanTepat []models.ProfileLulusan, lulusanTerlambat []models.ProfileLulusan, statusMahasiswa string) (peluangTepatStatusMahasiswa int, peluangTerlambatStatusMahasiswa int) {
	peluangTepatStatusMahasiswa = 0
	peluangTerlambatStatusMahasiswa = 0

	dataPeluangTepat := []models.ProfileLulusan{}
	dataPeluangTerlambat := []models.ProfileLulusan{}

	// tepat
	if len(lulusanTepat) > 0 {

		for _, detailLulusanTepat := range lulusanTepat {

			if detailLulusanTepat.StatusMahasiswa == statusMahasiswa {

				dataPeluangTepat = append(dataPeluangTepat, detailLulusanTepat)

			}

		}

		peluangTepatStatusMahasiswa = len(dataPeluangTepat)

	}

	if len(lulusanTerlambat) > 0 {
		for _, detailLulusanTerlambat := range lulusanTerlambat {

			if detailLulusanTerlambat.StatusMahasiswa == statusMahasiswa {

				dataPeluangTerlambat = append(dataPeluangTepat, detailLulusanTerlambat)

			}

		}

		peluangTerlambatStatusMahasiswa = len(dataPeluangTerlambat)
	}

	return peluangTepatStatusMahasiswa, peluangTerlambatStatusMahasiswa
}

func peluangTotalDataMasaStudiTepat(lulusanMasaStudi7, lulusanMasaStudi8, lulusanMasaStudi9, lulusanMasaStudi10 []models.ProfileLulusan, jenisKelamin string) (peluangGenderTepatMasaStudi7 int, peluangGenderTepatMasaStudi8 int, peluangGenderTepatMasaStudi9 int, peluangGenderTepatMasaStudi10 int) {

	peluangGenderTepatMasaStudi7 = 0
	peluangGenderTepatMasaStudi8 = 0
	peluangGenderTepatMasaStudi9 = 0
	peluangGenderTepatMasaStudi10 = 0

	dataPeluangTepatMasaStudi7 := []models.ProfileLulusan{}
	dataPeluangTepatMasaStudi8 := []models.ProfileLulusan{}
	dataPeluangTepatMasaStudi9 := []models.ProfileLulusan{}
	dataPeluangTepatMasaStudi10 := []models.ProfileLulusan{}

	// 7 smester
	if len(lulusanMasaStudi7) > 0 {

		for _, detailLulusanTepatMasaStudi7 := range lulusanMasaStudi7 {

			if detailLulusanTepatMasaStudi7.JenisKelamin == jenisKelamin {

				dataPeluangTepatMasaStudi7 = append(dataPeluangTepatMasaStudi7, detailLulusanTepatMasaStudi7)

			}

		}
		peluangGenderTepatMasaStudi7 = len(dataPeluangTepatMasaStudi7)
	}

	// 8 smester
	if len(lulusanMasaStudi8) > 0 {

		for _, detailLulusanTepatMasaStudi8 := range lulusanMasaStudi8 {

			if detailLulusanTepatMasaStudi8.JenisKelamin == jenisKelamin {

				dataPeluangTepatMasaStudi8 = append(dataPeluangTepatMasaStudi8, detailLulusanTepatMasaStudi8)

			}

		}
		peluangGenderTepatMasaStudi8 = len(dataPeluangTepatMasaStudi8)
	}

	// 9 smester
	if len(lulusanMasaStudi9) > 0 {

		for _, detailLulusanTepatMasaStudi9 := range lulusanMasaStudi9 {

			if detailLulusanTepatMasaStudi9.JenisKelamin == jenisKelamin {

				dataPeluangTepatMasaStudi9 = append(dataPeluangTepatMasaStudi9, detailLulusanTepatMasaStudi9)

			}

		}
		peluangGenderTepatMasaStudi9 = len(dataPeluangTepatMasaStudi9)
	}

	// 10 smester
	if len(lulusanMasaStudi10) > 0 {

		for _, detailLulusanTepatMasaStudi10 := range lulusanMasaStudi10 {

			if detailLulusanTepatMasaStudi10.JenisKelamin == jenisKelamin {

				dataPeluangTepatMasaStudi10 = append(dataPeluangTepatMasaStudi10, detailLulusanTepatMasaStudi10)

			}

		}
		peluangGenderTepatMasaStudi10 = len(dataPeluangTepatMasaStudi10)
	}

	return peluangGenderTepatMasaStudi7, peluangGenderTepatMasaStudi8, peluangGenderTepatMasaStudi9, peluangGenderTepatMasaStudi10
}

// status mahasiswa 'pekerja'
func peluangTotalDataMasaStudiStatusMahasiswaTepat(lulusanMasaStudi7, lulusanMasaStudi8, lulusanMasaStudi9, lulusanMasaStudi10 []models.ProfileLulusan, statusMahasiswa string) (peluangStatusMahasiswaTepatMasaStudi7 int, peluangStatusMahasiswaTepatMasaStudi8 int, peluangStatusMahasiswaTepatMasaStudi9 int, peluangStatusMahasiswaTepatMasaStudi10 int) {
	peluangStatusMahasiswaTepatMasaStudi7 = 0
	peluangStatusMahasiswaTepatMasaStudi8 = 0
	peluangStatusMahasiswaTepatMasaStudi9 = 0
	peluangStatusMahasiswaTepatMasaStudi10 = 0

	dataPeluangStatusMahasiswaTepatMasaStudi7 := []models.ProfileLulusan{}
	dataPeluangStatusMahasiswaTepatMasaStudi8 := []models.ProfileLulusan{}
	dataPeluangStatusMahasiswaTepatMasaStudi9 := []models.ProfileLulusan{}
	dataPeluangStatusMahasiswaTepatMasaStudi10 := []models.ProfileLulusan{}

	// 7 smester
	if len(lulusanMasaStudi7) > 0 {

		for _, detailLulusanTepatMasaStudi7 := range lulusanMasaStudi7 {

			if detailLulusanTepatMasaStudi7.StatusMahasiswa == statusMahasiswa {

				dataPeluangStatusMahasiswaTepatMasaStudi7 = append(dataPeluangStatusMahasiswaTepatMasaStudi7, detailLulusanTepatMasaStudi7)

			}

		}
		peluangStatusMahasiswaTepatMasaStudi7 = len(dataPeluangStatusMahasiswaTepatMasaStudi7)
	}
	// 8 smester
	if len(lulusanMasaStudi8) > 0 {

		for _, detailLulusanTepatMasaStudi8 := range lulusanMasaStudi8 {

			if detailLulusanTepatMasaStudi8.StatusMahasiswa == statusMahasiswa {

				dataPeluangStatusMahasiswaTepatMasaStudi8 = append(dataPeluangStatusMahasiswaTepatMasaStudi8, detailLulusanTepatMasaStudi8)

			}

		}
		peluangStatusMahasiswaTepatMasaStudi8 = len(dataPeluangStatusMahasiswaTepatMasaStudi8)
	}
	// 9 smester
	if len(lulusanMasaStudi9) > 0 {

		for _, detailLulusanTepatMasaStudi9 := range lulusanMasaStudi9 {

			if detailLulusanTepatMasaStudi9.StatusMahasiswa == statusMahasiswa {

				dataPeluangStatusMahasiswaTepatMasaStudi9 = append(dataPeluangStatusMahasiswaTepatMasaStudi9, detailLulusanTepatMasaStudi9)

			}

		}
		peluangStatusMahasiswaTepatMasaStudi9 = len(dataPeluangStatusMahasiswaTepatMasaStudi9)
	}
	// 7 smester
	if len(lulusanMasaStudi10) > 0 {

		for _, detailLulusanTepatMasaStudi10 := range lulusanMasaStudi10 {

			if detailLulusanTepatMasaStudi10.StatusMahasiswa == statusMahasiswa {

				dataPeluangStatusMahasiswaTepatMasaStudi10 = append(dataPeluangStatusMahasiswaTepatMasaStudi10, detailLulusanTepatMasaStudi10)

			}

		}
		peluangStatusMahasiswaTepatMasaStudi10 = len(dataPeluangStatusMahasiswaTepatMasaStudi10)
	}

	return peluangStatusMahasiswaTepatMasaStudi7, peluangStatusMahasiswaTepatMasaStudi8, peluangStatusMahasiswaTepatMasaStudi9, peluangStatusMahasiswaTepatMasaStudi10
}

// status mahasiswa 'mahasiswa'

func PrediksiLulus(c echo.Context) error {

	nim := c.FormValue("nim")
	nama := c.FormValue("nama")
	jenisKelamin := c.FormValue("jenis_kelamin")
	statusMahasiswa := c.FormValue("status_mahasiswa")
	// ipk1 := c.FormValue("ipk_1")
	// ipk2 := c.FormValue("ipk_2")
	// ipk3 := c.FormValue("ipk_3")
	// ipk4 := c.FormValue("ipk_4")
	// ips1 := c.FormValue("ips_1")
	// ips2 := c.FormValue("ips_2")
	// ips3 := c.FormValue("ips_3")
	// ips4 := c.FormValue("ips_4")

	// peluang prediksi Status Mahasiswa ( TEPAT atau tidak TEPAT )
	var profileLulusan []models.ProfileLulusan

	// mencari total data training
	err := config.DB.Find(&profileLulusan).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	totalDataLulusan := len(profileLulusan)

	lulusanTepat := []models.ProfileLulusan{}
	lulusanTerlambat := []models.ProfileLulusan{}
	lulusanMasaStudi7 := []models.ProfileLulusan{}
	lulusanMasaStudi8 := []models.ProfileLulusan{}
	lulusanMasaStudi9 := []models.ProfileLulusan{}
	lulusanMasaStudi10 := []models.ProfileLulusan{}
	// ipk lulus cumlaude ipk_lulus >=3.51 and ipk_lulus <=4.00
	ipkLulusCumlaude := []models.ProfileLulusan{}
	// ipk lulus sangat memuaskan ipk_lulus >=2.76 && ipk_lulus <=3.50
	ipkLulusSangatMemuaskan := []models.ProfileLulusan{}
	// ipk lulus memuaskan >=2.00 && ipk_lulus <=2.75
	ipkLulusMemuaskan := []models.ProfileLulusan{}

	for _, detailProfileLulusan := range profileLulusan {

		// get data 'TEPAT' yang ada pada data training
		if detailProfileLulusan.StatusKelulusan == "TEPAT" {
			lulusanTepat = append(lulusanTepat, detailProfileLulusan)

		}

		// get data 'TERLAMBAT' yang ada pada data training
		if detailProfileLulusan.StatusKelulusan == "TERLAMBAT" {
			lulusanTerlambat = append(lulusanTerlambat, detailProfileLulusan)
		}

		// get data 'masa studi = 7' yang ada pada data training
		if detailProfileLulusan.MasaStudi == 7 {
			lulusanMasaStudi7 = append(lulusanMasaStudi7, detailProfileLulusan)
		}
		// get data 'masa studi = 8' yang ada pada data training
		if detailProfileLulusan.MasaStudi == 8 {
			lulusanMasaStudi8 = append(lulusanMasaStudi8, detailProfileLulusan)
		}
		// get data 'masa studi = 9' yang ada pada data training
		if detailProfileLulusan.MasaStudi == 9 {
			lulusanMasaStudi9 = append(lulusanMasaStudi9, detailProfileLulusan)
		}
		// get data 'masa studi = 10' yang ada pada data training
		if detailProfileLulusan.MasaStudi == 10 {
			lulusanMasaStudi10 = append(lulusanMasaStudi10, detailProfileLulusan)
		}

		// total data 'IPK antara 3.51 dan 4.00' yang ada pada data training
		if (detailProfileLulusan.IpkLulus >= 3.51) && (detailProfileLulusan.IpkLulus <= 4.00) {
			ipkLulusCumlaude = append(ipkLulusCumlaude, detailProfileLulusan)
		}

		// total data 'IPK antara 2.76 dan 3.50' yang ada pada data training
		if (detailProfileLulusan.IpkLulus >= 2.76) && (detailProfileLulusan.IpkLulus <= 3.50) {
			ipkLulusSangatMemuaskan = append(ipkLulusSangatMemuaskan, detailProfileLulusan)
		}

		// total data 'IPK antara 2.00 dan 2.75' yang ada pada data training
		if (detailProfileLulusan.IpkLulus >= 2.00) && (detailProfileLulusan.IpkLulus <= 2.75) {
			ipkLulusMemuaskan = append(ipkLulusMemuaskan, detailProfileLulusan)
		}
	}

	// mencari peluang total data 'TEPAT' dan 'TERLAMBAT' dengan jenis kelamin yang dicari ("LAKI-LAKI","PEREMPUAN")
	peluangTepatGender, peluangTerlambatGender := peluangTotalDataJenisKelamin(lulusanTepat, lulusanTerlambat, jenisKelamin)

	// mencari peluang total data 'TEPAT' dan 'TERLAMBAT' dengan jenis status mahasiswa yang dicari ("PEKERJA","MAHASISWA")
	peluangTepatStatusMahasiswa, peluangTerlambatStatusMahasiswa := peluangTotalDataStatusMahasiswa(lulusanTepat, lulusanTerlambat, statusMahasiswa)

	peluangGenderTepatMasaStudi7, peluangGenderTepatMasaStudi8, peluangGenderTepatMasaStudi9, peluangGenderTepatMasaStudi10 := peluangTotalDataMasaStudiTepat(lulusanMasaStudi7, lulusanMasaStudi8, lulusanMasaStudi9, lulusanMasaStudi10, jenisKelamin)

	peluangStatusMahasiswaTepatMasaStudi7, peluangStatusMahasiswaTepatMasaStudi8, peluangStatusMahasiswaTepatMasaStudi9, peluangStatusMahasiswaTepatMasaStudi10 := peluangTotalDataMasaStudiStatusMahasiswaTepat(lulusanMasaStudi7, lulusanMasaStudi8, lulusanMasaStudi9, lulusanMasaStudi10, statusMahasiswa)

	_ = map[string]interface{}{
		"total-data-lulusan":                       float64(totalDataLulusan),
		"total-data-tepat":                         float64(len(lulusanTepat)),
		"total-data-terlambat":                     float64(len(lulusanTerlambat)),
		"total-data-masa-studi-7":                  float64(len(lulusanMasaStudi7)),
		"total-data-masa-studi-8":                  float64(len(lulusanMasaStudi8)),
		"total-data-masa-studi-9":                  float64(len(lulusanMasaStudi9)),
		"total-data-masa-studi-10":                 float64(len(lulusanMasaStudi10)),
		"total-data-ipk-cumlaude":                  float64(len(ipkLulusCumlaude)),
		"total-data-ipk-sangat-memuaskan":          float64(len(ipkLulusSangatMemuaskan)),
		"total-data-ipk-memuaskan":                 float64(len(ipkLulusMemuaskan)),
		"valid-lulusan-tepat":                      validTotalData(lulusanTepat, totalDataLulusan),
		"valid-lulusan-terlambat":                  validTotalData(lulusanTerlambat, totalDataLulusan),
		"valid-masa-studi-7":                       validTotalData(lulusanMasaStudi7, totalDataLulusan),
		"valid-masa-studi-8":                       validTotalData(lulusanMasaStudi8, totalDataLulusan),
		"valid-masa-studi-9":                       validTotalData(lulusanMasaStudi9, totalDataLulusan),
		"valid-masa-studi-10":                      validTotalData(lulusanMasaStudi10, totalDataLulusan),
		"valid-ipk-cumlaude":                       validTotalData(ipkLulusCumlaude, totalDataLulusan),
		"valid-ipk-sangat-memuaskan":               validTotalData(ipkLulusSangatMemuaskan, totalDataLulusan),
		"valid-ipk-memuaskan":                      validTotalData(ipkLulusMemuaskan, totalDataLulusan),
		"peluang-jk-tepat":                         float64(peluangTepatGender),
		"peluang-jk-terlambat":                     float64(peluangTerlambatGender),
		"peluang-status-mahasiswa-tepat":           float64(peluangTepatStatusMahasiswa),
		"peluang-status-mahasiswa-terlambat":       float64(peluangTerlambatStatusMahasiswa),
		"valid-peluang-jk-tepat":                   math.Round((float64(peluangTepatGender)/float64(len(lulusanTepat)))*100) / 100,
		"valid-peluang-jk-terlambat":               math.Round((float64(peluangTerlambatGender)/float64(len(lulusanTerlambat)))*100) / 100,
		"valid-peluang-status-mahasiswa-tepat":     math.Round((float64(peluangTepatStatusMahasiswa)/float64(len(lulusanTepat)))*100) / 100,
		"valid-peluang-status-mahasiswa-terlambat": math.Round((float64(peluangTerlambatStatusMahasiswa)/float64(len(lulusanTerlambat)))*100) / 100,
	}

	// peluang tepat
	peluangTepat := validTotalData(lulusanTepat, totalDataLulusan) * math.Round((float64(peluangTepatGender)/float64(len(lulusanTepat)))*100) / 100 * math.Round((float64(peluangTepatStatusMahasiswa)/float64(len(lulusanTepat)))*100) / 100

	// peluang terlambat
	peluangTerlambat := validTotalData(lulusanTerlambat, totalDataLulusan) * math.Round((float64(peluangTerlambatGender)/float64(len(lulusanTerlambat)))*100) / 100 * math.Round((float64(peluangTerlambatStatusMahasiswa)/float64(len(lulusanTerlambat)))*100) / 100

	// peluang 7 semester
	peluang7semester := validTotalData(lulusanMasaStudi7, totalDataLulusan) * math.Round((float64(peluangGenderTepatMasaStudi7)/float64(len(lulusanMasaStudi7)))*100) * math.Round((float64(peluangStatusMahasiswaTepatMasaStudi7)/float64(len(lulusanMasaStudi7)))*100)

	// peluang 8 semester
	peluang8semester := validTotalData(lulusanMasaStudi8, totalDataLulusan) * math.Round((float64(peluangGenderTepatMasaStudi8)/float64(len(lulusanMasaStudi8)))*100) * math.Round((float64(peluangStatusMahasiswaTepatMasaStudi8)/float64(len(lulusanMasaStudi8)))*100)

	// peluang 9 semester
	peluang9semester := validTotalData(lulusanMasaStudi9, totalDataLulusan) * math.Round((float64(peluangGenderTepatMasaStudi9)/float64(len(lulusanMasaStudi9)))*100) * math.Round((float64(peluangStatusMahasiswaTepatMasaStudi9)/float64(len(lulusanMasaStudi9)))*100)

	// peluang 10 semester
	peluang10semester := validTotalData(lulusanMasaStudi10, totalDataLulusan) * math.Round((float64(peluangGenderTepatMasaStudi10)/float64(len(lulusanMasaStudi10)))*100) * math.Round((float64(peluangStatusMahasiswaTepatMasaStudi10)/float64(len(lulusanMasaStudi10)))*100)

	// prediksi lulus 'TEPAT' atau 'TERLAMBAT'
	prediksiStatusKelulusan := "TERLAMBAT"
	if peluangTepat > peluangTerlambat {
		prediksiStatusKelulusan = "TEPAT"

	}

	// prediksi lulus masa studi 7, 8, 9 atau 10 semester
	prediksiMasaStudi := "10 semester"
	if peluang7semester > peluang8semester {
		prediksiMasaStudi = "7 semester"

	} else if peluang8semester > peluang9semester {
		prediksiMasaStudi = "8 semester"
	} else if peluang9semester > peluang10semester {
		prediksiMasaStudi = "9 semester"
	} else {
		prediksiMasaStudi = "10 semester"
	}

	response := map[string]interface{}{
		// "data-training": dataTraining,
		// "profile-lulusan": profileLulusan,
		"prediksi-status-kelulusan": prediksiStatusKelulusan,
		"prediksi_masa_studi":       prediksiMasaStudi,
		"jumlah_peluang_kelulusan": map[string]interface{}{
			"peluang_tepat":      peluangTepat,
			"peluang_terlambat":  peluangTerlambat,
			"peluang_7_smester":  peluang7semester,
			"peluang_8_smester":  peluang8semester,
			"peluang_9_smester":  peluang9semester,
			"peluang_10_smester": peluang10semester,
		},
		"biodata_pengujian": map[string]interface{}{
			"nama":                       nama,
			"nim":                        nim,
			"jenis_kelamin":              jenisKelamin,
			"status_mahasiswa":           statusMahasiswa,
			"status_prediksi_kelulusan":  prediksiStatusKelulusan,
			"status_prediksi_masa_studi": prediksiStatusKelulusan,
		},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success",
		"response": response,
	})

}
