package slices

import "fmt"

func Name() {
	names := make([]string, 0)

	var capacity int
	var namesadd string
	fmt.Println("Özel ders kaydı için kişi sayısı girin(maks:8):")
	fmt.Scanln(&capacity)
	if capacity <= 8 {
		for i := 1; i <= capacity; i++ {
			fmt.Printf("Lütfen %v . ismi giriniz:", i)
			fmt.Scanln(&namesadd)
			names = append(names, namesadd)
		}
		fmt.Println("Kaydı oluşturulan isimler:", names)
	} else {
		fmt.Println("Mevcut çok fazla, lütfen 8 kişi veya daha azını girerek tekrar deneyiniz.")
	}

}
