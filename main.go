package main

import (
	"fmt"
	"baju/view"
)

func menu() {
	for {
		fmt.Println("Menu Program")
		fmt.Println("1. Insert Baju")
		fmt.Println("2. Update Baju")
		fmt.Println("3. Delete Baju")
		fmt.Println("4. Read All Baju")
		fmt.Println("5. Search Baju")
		fmt.Println("6. Exit")
		fmt.Println("---------------------")
		fmt.Print("masukkan menu pilihan anda: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			view.InsertBaju()
		case 2:
			view.UpdateBaju()
		case 3:
			view.DeleteBaju()
		case 4:
			view.ReadAllBaju()
		case 5:
			view.SearchBaju()
		case 6:
			fmt.Println("exit program")
			return
		}
	}
}

func main() {
	menu()

	
}
