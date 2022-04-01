package main

import (
	controllers "TugasFramework/controllers"
	model "TugasFramework/model"
	"context"
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	fmt.Print("Kamu mau email 1 atau banyak orang ? ")
	var input int
	fmt.Scanln(&input)

	fmt.Print("Masukan kode jurusan mahasiswa yang akan diemail : ")
	var kode string
	fmt.Scanln(&kode)

	fmt.Print("Masukan angkatan mahasiswa yang akan diemail : ")
	var angkatan int
	fmt.Scanln(&angkatan)

	var absen int
	var jumlahAnak int
	if input == 1 {
		fmt.Print("Masukan absen anak yang akan diemail : ")
		fmt.Scanln(&absen)
		jumlahAnak = 1
	} else {
		fmt.Print("Masukan absen pertama anak yang akan diemail : ")
		fmt.Scanln(&absen)

		fmt.Print("Masukan absen akhir anak yang akan diemail : ")
		var absenAkhir int
		fmt.Scanln(&absenAkhir)
		jumlahAnak = absenAkhir - absen
	}
	var informasi = model.Informasi{KodeJurusan: kode, Angkatan: angkatan, Absen: absen, JumlahAnak: jumlahAnak}

	fmt.Print("Isi subjek : ")
	var subjek string
	fmt.Scanln(&subjek)

	fmt.Print("Isi text : ")
	var text string
	fmt.Scanln(&text)

	cc := []string{"tralalala@gmail.com"}
	to := []string{}
	go controllers.EvenEmail(to, informasi)
	go controllers.OddEmail(to, informasi)
	time.Sleep(3000 * time.Millisecond)

	var email string
	var password string
	email, err := client.Get(ctx, "email").Result()
	password, err1 := client.Get(ctx, "password").Result()
	if err1 != nil || err != nil {
		fmt.Print("email: ")
		fmt.Scanln(&email)
		fmt.Print("password: ")
		fmt.Scanln(&password)
		err := client.Set(ctx, "key", "value", 0).Err()
		if err != nil {
			fmt.Println(err)
		}
		err = client.Set(ctx, "password", password, 0).Err()
		if err != nil {
			fmt.Println(err)
		}

	}

	var nama string
	fmt.Scanln(&nama)
	sender := controllers.GenerateSender(nama, email)
	controllers.SendMail(to, cc, subjek, text, email, password, sender)
	s := gocron.NewScheduler(time.UTC)
	// s.Every(1).Minute().Do(func() { controllers.SendMail(to, cc, subjek, text) })
	// s.Every(7).Day().At("10:30").Do(func() { controllers.SendMail(to, cc, subjek, text) })
	s.StartBlocking()
}
