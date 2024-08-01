package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

type Alamat struct {
	Nama, Alamat, Telp string
}

func main() {
	datacsv := "Nama,Alamat,Telp"
	datacsv = datacsv + "Didi,Jl.Mangga,123456\n"
	datacsv = datacsv + "Agus,Jl.Apel,6666\n"
	datacsv = datacsv + "Tono,Jl.Pegangsaan,4444\n"
	fmt.Println(datacsv)

	reader := csv.NewReader(strings.NewReader(datacsv))
	for {
		records, err := reader.Read()
	}

	// Membaca semua data dari file CSV
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Menampilkan data yang telah dibaca
	for _, record := range records {
		fmt.Println(record)
	}
}
