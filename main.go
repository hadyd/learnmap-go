package main

import "fmt"

type produk struct {
	namaProduk  string
	hargaProduk int
	stokProduk  int
}

func main() {
	minimarket := make(map[int]produk)
	minimarket[1] = produk{namaProduk: "Pasta Gigi", hargaProduk: 8000, stokProduk: 11}
	minimarket[2] = produk{namaProduk: "Shampo", hargaProduk: 6000, stokProduk: 11}
	minimarket[3] = produk{namaProduk: "Sabun", hargaProduk: 3000, stokProduk: 12}
	minimarket[4] = produk{namaProduk: "Masker", hargaProduk: 7000, stokProduk: 5}
	for _, item := range minimarket {
		if item.stokProduk <= 10 {
			fmt.Println("Daftar Stock Barang di Bawah 10:", "\n...............................")
			fmt.Println("Nama Barang: ", item.namaProduk, "\nHarga Barang: Rp.", item.hargaProduk, "\nStok Barang:", item.stokProduk, "\n...............................")
		}
	}
}
