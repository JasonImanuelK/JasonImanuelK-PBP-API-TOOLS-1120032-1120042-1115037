package model

type Informasi struct {
	KodeJurusan string `json:"kodeJurusan"`
	Angkatan    int    `json:"angkatan"`
	Absen       int    `json:"absen"`
	JumlahAnak  int    `json:"jumlahAnak"`
}
