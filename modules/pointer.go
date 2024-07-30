package modules

import "fmt"

// & digunakan untuk mengubah isi value menjadi alamat memori
// * digunakan untuk mengubah alamat memori menjadi isi value
func Pointer() {
	// mencetak value seperti biasa
	fmt.Println("Tanpa pointer")
	nilai1 := 10
	nilai2 := 10
	fmt.Println(nilai1)
	fmt.Println(nilai2)

	// membuat 2 var berbeda dengan alamat yang sama menggunakan referensi
	// diawali dengan "&" yang menandakan value berisi alamat memori
	fmt.Println("\nMenggunakan pointer")
	nilai3 := &nilai1
	fmt.Println(&nilai1)
	fmt.Println(nilai3)

	// mencetak value dari 2 var berbeda dengan pointer
	// var baru yang direferensikan ke value lain akan menghasilkan value yang sama
	fmt.Println("\nMencetak valuew dari referensi dengan pointer")
	fmt.Println(nilai1)
	fmt.Println(*nilai3)

	// value dari var baru dibuah dengan menggunakan pointer
	// maka value referensi akan ikut berubah sesuai value baru yang dialamatkan sesuai referensinya
	fmt.Println("\nMengubah value dengan pointer")
	fmt.Println("Sebelum diubah - dengan pointer")
	fmt.Println(nilai1)
	fmt.Println(*nilai3)
	fmt.Println("Sebelum diubah - dengan pointer")
	*nilai3 = 20
	fmt.Println(nilai1)
	fmt.Println(*nilai3)
}

func NoPointer() {
	// mengisi value dari var baru tanpa pointer
	// value dari var baru akan sama dengan value dari var
	// jika value baru dirubah maka value referensinya tidak akan ikut berubah
	fmt.Println("\nMengubah value tanpa pointer")
	fmt.Println("Sebelum diubah - tanpa pointer")
	nilai1 := 30
	nilai4 := 30
	fmt.Println(nilai1)
	fmt.Println(nilai4)
	fmt.Println("Setelah diubah - tanpa pointer")
	nilai1 = nilai4
	nilai4 = 40
	fmt.Println(nilai1)
	fmt.Println(nilai4)
}
