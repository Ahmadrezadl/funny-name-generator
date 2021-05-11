package NameGenerator

import (
	"math/rand"
	"time"
)

var (
	Words1 = [...]string{"اصغر", "صغرا", "علی", "حسن", "ممد", "کمال", "جک", "کامبیز", "عباس"}
	Words2 = [...]string{"سیبیل", "ممه", "پا", "مغز", "دست", "چشم", "شیکم", "مو", "انگشت", "کون"}
	Words3 = [...]string{"قشنگ", "گنده", "کوچولو", "دست", "تیز", "طلایی", "سیاه", "کبابی"}
)

var (
	FinglishWords1 = [...]string{"Asghar", "Soghra", "Ali", "Hasan", "Mamad", "Kamal", "Jack", "Kambiz", "Abbas"}
	FinglishWords2 = [...]string{"Sibil", "Mame", "Pa", "Maghz", "Dast", "Cheshm", "Shikam", "Moo", "Angosht", "Koon"}
	FinglishWords3 = [...]string{"Ghashang", "Gonde", "Kocholo", "Dast", "Tiz", "Talaei", "Siah", "Kababi"}
)

func Randomize() {
	rand.Seed(time.Now().UnixNano())
}

func PersianName() string {
	return Words1[rand.Intn(len(Words1))] + " " + Words2[rand.Intn(len(Words2))] + " " + Words3[rand.Intn(len(Words3))]
}

func FinglishName() string {
	return FinglishWords1[rand.Intn(len(FinglishWords1))] + " " + FinglishWords2[rand.Intn(len(FinglishWords2))] + " " + FinglishWords3[rand.Intn(len(FinglishWords3))]
}
