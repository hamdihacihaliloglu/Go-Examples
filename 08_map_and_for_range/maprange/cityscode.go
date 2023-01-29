package maprange

import "fmt"

func CitysCode() {
	citys := make(map[string]string)
	citys["Ankara"] = "06"
	citys["Samsun"] = "55"
	citys["İstanbul"] = "34"
	citys["Adana"] = "01"

	for k, v := range citys {
		fmt.Print("Şehir: ", k)
		fmt.Println(" ", v)
	}

}

func CitysPhoneCode() {
	citys2 := map[string]string{"Amasya": "358", "Ankara": "312"}
	for k, v := range citys2 {
		fmt.Print("Şehir ve Telefon Kodu: ", k)
		fmt.Println(" ", v)
	}
}
