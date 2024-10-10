package main
import "fmt"

func main() {
	var angka1, angka2, angka3, angka4, angka5 int
	var char1, char2, char3 byte

	fmt.Print("Masukkan 5 angka integer (dipisah spasi): ")
	fmt.Scanf("%d %d %d %d %d\n", &angka1, &angka2, &angka3, &angka4, &angka5)

	fmt.Print("Masukkan 3 karakter (tanpa spasi): ")
	fmt.Scanf("%c%c%c", &char1, &char2, &char3)

	fmt.Printf("Output pertama: %c %c %c %c %c\n", angka1, angka2, angka3, angka4, angka5)
	fmt.Printf("Output kedua: %c%c%c\n", char1+1, char2+1, char3+1)
}