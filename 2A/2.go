package main
import "fmt"

func main() {
	var tahun int
	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&tahun)

	if tahun%400 == 0 {
		fmt.Println("Tahun", tahun, "adalah tahun kabisat")
	} else if tahun%100 == 0 {
		fmt.Println("Tahun", tahun, "bukan tahun kabisat")
	} else if tahun%4 == 0 {
		fmt.Println("Tahun", tahun, "adalah tahun kabisat")
	} else {
		fmt.Println("Tahun", tahun, "bukan tahun kabisat")
	}
}