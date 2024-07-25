package node

type Baju struct {
	Id      int
	Merk    string
	Kondisi string
	Ukuran  float32
}

type TableBajustruct {
	Data Baju
	Next *TableBaju
}
