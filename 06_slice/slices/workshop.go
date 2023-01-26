package slices

import "fmt"

func Workshop() {
	lessons := []string{"Matematik", "Fizik", "Edebiyat"}
	lessonsCopy := make([]string, len(lessons))
	copy(lessonsCopy, lessons)
	var optionalLessons int
	var optinonalLessonsName string
	fmt.Println("Bir döenmde en fazla 5 seçmeli ders seçeblirsiniz.")
	fmt.Println("Lütfen seçmeli ders adedini giriniz:")
	fmt.Scanln(&optionalLessons)
	if optionalLessons <= 5 {
		for i := 1; i <= optionalLessons; i++ {
			fmt.Printf("%v . ders adını giriniz:", i)
			fmt.Scanln(&optinonalLessonsName)
			lessonsCopy = append(lessonsCopy, optinonalLessonsName)

		}
		fmt.Println("Dönem içindeki zorunlu dersleriniz:", lessons)
		fmt.Println("Dönem içindeki seçmeli dersleriniz:", lessonsCopy[3:])
	} else {
		fmt.Println("Lütfen uyarıyı dikkate alınız.")
	}

}
