package conditional

import "fmt"

func BankDemo() {
	hesapbakiyesi := 1000.00
	cekilmekistenen := 400.00
	if cekilmekistenen > hesapbakiyesi {
		fmt.Println("Bakiye yetersiz")
	} else if cekilmekistenen <= hesapbakiyesi {
		kalanpara := hesapbakiyesi - cekilmekistenen
		fmt.Println("Kalan Bakiyeniz : " + fmt.Sprintf("%f", kalanpara))
	} else {
		fmt.Println("Geçersiz İşlem")
	}

}
