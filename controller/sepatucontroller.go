package controller

import (
	"baju/model"
	"baju/node"
)

func InsertBaju(id int, merk string, kondisi string, ukuran float32) bool {
	if id > 0 && merk != "" {
		model.InsertBaju(id, merk, kondisi, ukuran)
		return true
	} else {
		return false
	}
}

func UpdateBaju(id int, merk string, kondisi string, ukuran float32) bool {
	if id > 0 && merk != "" {
		sepatu := node.Sepatu{
			Id:      id,
			Merk:    merk,
			Kondisi: kondisi,
			Ukuran:  ukuran,
		}
		return model.UpdateBaju(sepatu)
	}
	return false
}

func SearchBaju(id int, merk string, kondisi string, ukuran float32) *node.Sepatu {
	if id > 0 {
		addr := model.SearchBaju(id)
		if addr != nil {
			return &addr.Data
		}
	}
	return nil
}

func DeleteBaju(id int) bool {
	if id > 0 {
		return model.DeleteBaju(id)
	}
	return false
}

func ReadAllBaju() *node.TableSepatu {
	return model.ReadAllBaju()
}
