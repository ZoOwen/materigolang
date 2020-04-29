package main


import(
	"fmt"
) 
// untuk menyatukan string dan text di print harus menggunakan koma (,)
//untuk menggabungkan string menggunakan + +
//perbedaan int dan float , int untuk bilangan bulat, float bisa berisi koma (,)
//penggunaan := (sorthand var ) hanya bisa didalam function
//, jika hanya di deklarasikan tetapi tidak di pakai maka akan terjadi error
var namaDepan = "Raif"
var namaBelakang = "Fahmi"

var namaLengkap = namaDepan + " " + namaBelakang

var umur = 18

func main() {
	sekolah := "di smakzie"
	fmt.Println("hello nama saya " + namaLengkap + " saya berumur", umur ,
		"tahun, saya bersekolah di " + sekolah)
}