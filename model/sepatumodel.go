package model

import (
	"baju/database"
	"baju/node"
)

func InsertBaju(id int, merk string, kondisi string, ukuran float32) {
	newDataBaju := node.TableBaju{
		Data: node.Baju{id, merk, kondisi, ukuran},
		Next: nil,
	}
	var tempLL *node.TableBaju
	tempLL = &database.DBbaju
	if database.DBbaju.Next == nil {
		database.DBbaju.Next = &newDataBaju
	} else {
		for tempLL.Next != nil {
			tempLL = tempLL.Next
		}
		tempLL.Next = &newDataBaju
	}
}

func ReadAllBaju() *node.TableBaju {
	var tempLL *node.TableBaju
	tempLL = &database.DBbaju
	if database.DBbaju.Next == nil {
		return nil
	} else {
		return tempLL
	}
}

func DeleteBaju(id int) bool {
	var tempLL *node.TableBaju
	tempLL = &database.DBbaju
	if tempLL.Next == nil {
		return false
	} else {
		for tempLL.Next != nil {
			if tempLL.Next.Data.Id == id {
				tempLL.Next = tempLL.Next.Next
				return true
			}
			tempLL = tempLL.Next
		}
		return false
	}

}

func SearchBaju(id int) *node.TableBaju {
	var tempLL *node.TableBaju
	tempLL = database.DBbaju.Next
	cek := false
	if database.DBbaju.Next == nil {
		return nil
	} else {
		for tempLL != nil {
			if id == tempLL.Data.Id {
				cek = true
				return tempLL
			}
			tempLL = tempLL.Next
		}
		if cek != true {
			return nil
		}
	}
	return nil
}

func UpdateBaju(baju node.Baju) bool {
	addr := SearchBaju(baju.Id)
	if addr == nil {
		return false
	} else {
		addr.Data.Merk = baju.Merk
		addr.Data.Kondisi = baju.Kondisi
		addr.Data.Ukuran = baju.Ukuran
		return true
	}
}
