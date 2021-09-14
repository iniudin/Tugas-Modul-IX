package main

import (
	"errors"
	"fmt"
	"strings"
)

var menuPilihan = []string{"tahu", "tempe", "sambal", "lele", "ayam", "nasi"}

func checkList(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// fungsi ini digunakan untuk mengecek input user kosong / tidak
// dan mengecek apakah pesanan terdapat pada menu / tidak
func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("input tidak boleh kosong")
	}

	check := checkList(menuPilihan, input)
	if !check {
		return false, errors.New("menu tidak terdapat pada daftar menu")
	}

	return true, nil
}

func MenuMakanan() {
	fmt.Println("Menu Toko Makanan")
	fmt.Println("=================")
	for _, v := range menuPilihan {
		fmt.Printf("- %s\n", strings.Title(v))
	}
	fmt.Println("=================")
}

func main() {
	var pesanan = []string{}
	var inputPesanan string
	var prompt string

	MenuMakanan()

	for {
		fmt.Printf("Masukkan menu pilihan anda: ")
		fmt.Scan(&inputPesanan)

		if valid, err := validate(inputPesanan); valid {
			pesanan = append(pesanan, inputPesanan)
		} else {
			fmt.Println(err.Error())
		}

		fmt.Printf("Lanjut memesan? (y/t): ")
		fmt.Scan(&prompt)

		if strings.ToLower(prompt) == "y" {
			continue
		} else if strings.ToLower(prompt) == "t" {
			break
		} else {
			fmt.Println("silahkan ketik y/t")
			continue
		}
	}

	fmt.Println("Anda Memesan:")
	for _, v := range pesanan {
		fmt.Printf("- %s\n", strings.Title(v))
	}

	fmt.Println("Terimakasih atas pesanan anda!")
}
