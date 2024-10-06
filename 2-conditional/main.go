//Dalam bahasa pemrograman Golang, terdapat 2 tipe percabangan, yaitu dengan menggunakan if else dan switch.

package main

import ("fmt")

//Penggunaan if else dasar

// func main()  {
// 	numChild := 5
// 	if numChild > 2{
// 		fmt.Println("Banyak anak banyak rezeki")
// 	} else {
// 		fmt.Println("2 anak sudah cukup")
// 	}
// }



// Percabangan if else yang ada di Golang dapat menggunakan temporary variable yang mungkin di bahasa
// pemrograman lain hal itu belum tersedia. Temporary Variable yang dimaksud adalah sebuah variabel yang 
//hanya digunakan pada satu kode blok tersebut saja. Agar lebih mudah dalam memahami, silakan perhatikan 
//blok kode di bawah ini.
//Penggunaan temporary variable

// func main()  {
// 	numChild := 5; if numChild > 2{
// 		fmt.Println("Banyak anak banyak rezeki")
// 	} else{
// 		fmt.Println("Cukup aja")
// 	}
// 	fmt.Println(numChild)
// }



//Penggunaan switch case

//Penggunaan switch case Dasar
// func main(){
// 	numChild := 1
// 	switch numChild{
// 	case 5:
// 		fmt.Println("Mantap")
// 	case 3:
// 		fmt.Println("Bagus")
// 	default:
// 		fmt.Println("Semangat")
// 	}
// }


//switch case bisa juga digunakan layaknya seperti memakai percabangan if else, caranya adalah 
//sebagai berikut.

func main(){
	numChild:= 5
	switch {
	case numChild > 5:
		fmt.Println("lebih dari lima")
	case numChild > 3:
		fmt.Println("lebih dari tiga")
	case numChild < 5:
		fmt.Println("kurang dari lima")
	}
}


