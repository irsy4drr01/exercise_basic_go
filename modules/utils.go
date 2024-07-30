package modules

// exported func dengan uppercase di awal nama func
func Sum(n1, n2 int) (int, int) { // jika semua parameter memiliki tipe data sama bisa ditulis hanya sekali saja di akhir
	return n1 + n2, times(5, 4)
}

// unexported func dengan huruf lowercase di awal nama func
func times(n1, n2 int) int {
	return n1 * n2
}
