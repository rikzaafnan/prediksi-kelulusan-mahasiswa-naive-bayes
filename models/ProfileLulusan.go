package models

type ProfileLulusan struct {
	// ID              int     `json:"id" form:"id"`
	NIM             string  `json:"nim" form:"nim"`
	Nama            string  `json:"nama" form:"nama"`
	JenisKelamin    string  `json:"jenis_kelamin" form:"jenis_kelamin"`
	StatusMahasiswa string  `json:"status_mahasiswa" form:"status_mahasiswa"`
	Ipk1            float64 `json:"ipk_1" form:"ipk_1"`
	Ipk2            float32 `json:"ipk_2" form:"ipk_2"`
	Ipk3            float32 `json:"ipk_3" form:"ipk_3"`
	Ipk4            float32 `json:"ipk_4" form:"ipk_4"`
	RataIpk         float32 `json:"rata_ipk" form:"rata_ipk"`
	Ips1            float32 `json:"ips_1" form:"ips_1"`
	Ips2            float32 `json:"ips_2" form:"ips_2"`
	Ips3            float32 `json:"ips_3" form:"ips_3"`
	Ips4            float32 `json:"ips_4" form:"ips_4"`
	MasaStudi       int     `json:"masa_studi" form:"masa_studi"`
	StatusKelulusan string  `json:"status_kelusan" form:"status_kelusan"`
	IpkLulus        float32 `json:"ipk_lulus" form:"ipk_lulus"`
}
