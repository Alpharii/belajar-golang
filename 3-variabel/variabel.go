package main

import("fmt")

// func main(){
// 	var name string = "bintang"
// 	fmt.Println(name)

// 	var namaOpsional string;
// 	namaOpsional = "brighton"
// 	fmt.Println(namaOpsional)

// 	name = "alphari"
// 	fmt.Println(name)

// 	var angka1, angka2, angka3 int
// 	var c, x byte
// 	var isReal bool
// 	var uang, desimal float32
// 	fmt.Println(angka1, angka2, angka3, c, x, isReal, uang, desimal)
// }

func main(){
	var angka float32
	angka = 42.5
	fmt.Println(angka)
	fmt.Printf("tipe data angka adalah %T\n", angka)
}

//perbedaan fmt printf dan println

// fmt.Println: Lebih sederhana dan digunakan untuk mencetak nilai dengan spasi dan newline otomatis. Tidak mendukung format string.
// fmt.Printf: Digunakan untuk mencetak dengan kontrol lebih spesifik terhadap format output, tetapi tidak menambahkan newline secara otomatis.

//Ciri-ciri fmt.Println:
// Menambahkan spasi antara argumen yang dicetak.
// Secara otomatis menambahkan newline (\n) di akhir output.
// Tidak mendukung format string seperti %d, %s, atau %T.

//Ciri-ciri fmt.Printf:
// Menggunakan format string untuk menampilkan output, seperti:
// %d untuk bilangan bulat (integer).
// %f untuk bilangan desimal (float).
// %s untuk string.
// %T untuk tipe data dari variabel.
// Tidak secara otomatis menambahkan newline, jadi jika ingin menambahkannya, gunakan \n dalam string format.
// Lebih fleksibel karena mendukung format yang bisa disesuaikan dengan tipe data.