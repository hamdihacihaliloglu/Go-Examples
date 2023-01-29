package maprange

import "fmt"

func DictionaryRange() {
	dictionary := make(map[string]string)
	dictionary["Apple"] = "Elma"
	dictionary["Phone"] = "Telefon"
	dictionary["Book"] = "Kitap"
	dictionary["Table"] = "Masa"
	dictionary["Computer"] = "Bilgisayar"

	for k, v := range dictionary {
		fmt.Print("En: ", k)
		fmt.Println(" Tr: ", v)
	}

	delete(dictionary, "Phone")
	value1, bool1 := dictionary["Phone"]

	fmt.Println(value1)
	fmt.Println("Yeni listede olma durumu:", bool1)

	for k, v := range dictionary {
		fmt.Print("En: ", k)
		fmt.Println(" Tr: ", v)
	}
}
