package view

import (
	"fmt"
	"baju/controller"
)

func InsertBaju() {
	var id int
	var merk, kondisi string
	var ukuran float32
	fmt.Print("Masukkan id baju : ")
	fmt.Scan(&id)
	fmt.Print("Masukkan merk baju : ")
	fmt.Scan(&merk)
	fmt.Print("Masukkan kondisi baju : ")
	fmt.Scan(&kondisi)
	fmt.Print("Masukkan ukuran baju : ")
	fmt.Scan(&ukuran)
	cek := controller.InsertBaju(id, merk, kondisi, ukuran)
	if cek == true {
		fmt.Println("data berhasil diinput")
	} else {
		fmt.Println("data gagal diinput")
	}
}

func UpdateBaju() {
	var id int
	var merk, kondisi string
	var ukuran float32
	fmt.Print("Masukkan id baju yang ingin diupdate: ")
	fmt.Scan(&id)
	fmt.Print("Masukkan merk baru baju: ")
	fmt.Scan(&merk)
	fmt.Print("Masukkan kondisi baru baju: ")
	fmt.Scan(&kondisi)
	fmt.Print("Masukkan ukuran baru baju: ")
	fmt.Scan(&ukuran)
	cek := controller.UpdateBaju(id, merk, kondisi, ukuran)
	if cek == true {
		fmt.Println("Data berhasil diupdate")
	} else {
		fmt.Println("Data gagal diupdate atau baju tidak ditemukan")
	}
}

func ReadAllBaju() {
	bajus := controller.ReadAllBaju()
	if bajus != nil {
		for bajus.Next != nil {
			fmt.Println("Id Baju : ", bajus.Next.Data.Id)
			fmt.Println("Merk Baju : ", bajus.Next.Data.Merk)
			fmt.Println("Ukuran Baju : ", bajus.Next.Data.Ukuran)
			fmt.Println("Kondisi Baju : ", bajus.Next.Data.Kondisi)
			fmt.Println("---------------------")
			bajus = bajus.Next
		}
	}
}

func SearchBaju() {
	var id int
	var merk, kondisi string
	var ukuran float32
	fmt.Print("Masukkan id baju yang ingin dicari: ")
	fmt.Scan(&id)
	baju := controller.SearchBaju(id, merk, kondisi, ukuran)
	if baju != nil {
		fmt.Println("Id Baju : ", baju.Id)
		fmt.Println("Merk Baju : ", baju.Merk)
		fmt.Println("Ukuran Baju : ", baju.Ukuran)
		fmt.Println("Kondisi Baju : ", baju.Kondisi)
	} else {
		fmt.Println("Baju tidak ditemukan")
	}
}

func DeleteBaju() {
	var id int
	fmt.Println("Masukkan id baju yang ingin dihapus: ")
	fmt.Scan(&id)
	cek := controller.DeleteBaju(id)
	if cek == true {
		fmt.Println("Data berhasil dihapus")
	} else {
		fmt.Println("Data gagal dihapus atau baju tidak ditemukan")
	}
}
