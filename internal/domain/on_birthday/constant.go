package on_birthday

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	DOC_TYPE             = "on_birthday"
	AGE_CALLOUT          = 35
	MALE_YOUNG_CALLOUT   = "Mas "
	MALE_OLD_CALLOUT     = "Pak "
	FEMALE_YOUNG_CALLOUT = "Mbak "
	FEMALE_OLD_CALLOUT   = "Bu "
)

var MessagesIndex []string = []string{
	"Cieee ada yang ulang tahun niich!",
	"Lhooo kok udah ultah aja ?!",
	"Hi guys, aku ultah niih! Ucapin ngapa ?",
}

var FooterIndex []string = []string{
	"Semoga hari ulang tahunmu penuh dengan kebahagiaan dan cinta!",
	"Selamat ulang tahun! Semoga tahun ini membawa banyak keberuntungan dan kesuksesan bagimu.",
	"Ucapan selamat ulang tahun yang hangat untukmu! Semoga semua impianmu menjadi kenyataan.",
	"Selamat ulang tahun yang indah! Semoga hari ini penuh dengan kebahagiaan dan keceriaan.",
	"Hari ini adalah hari istimewa karena kamu lahir. Selamat ulang tahun yang penuh berkat!",
	"Selamat ulang tahun! Semoga setiap langkahmu di tahun ini dipenuhi dengan keberuntungan dan kebahagiaan.",
	"Ucapan selamat ulang tahun yang penuh kasih sayang untukmu! Semoga hidupmu selalu diberkati dengan kebahagiaan dan kesuksesan.",
	"Selamat ulang tahun! Semoga hari ini menjadi awal dari tahun yang fantastis bagimu.",
	"Ucapan selamat ulang tahun yang tulus untukmu! Semoga setiap momen spesial di hari ini menjadi kenangan yang abadi.",
	"Selamat ulang tahun yang menyenangkan! Semoga hari ini menjadi salah satu dari banyak hari indah dalam hidupmu.",
}

func BirthdayContent(index int) string {
	return MessagesIndex[index]
}

func BirthdayWishes(index int) string {
	return FooterIndex[index]
}

func BirthdayDescription(empName, deptName, posName, gender string, age int) string {
	var msg string

	if gender == "M" {
		msg = MALE_YOUNG_CALLOUT

		if age > AGE_CALLOUT {
			msg = MALE_OLD_CALLOUT
		}
	} else {
		msg = FEMALE_YOUNG_CALLOUT

		if age > AGE_CALLOUT {
			msg = FEMALE_OLD_CALLOUT
		}
	}

	for _, c := range []cases.Caser{
		cases.Title(language.Indonesian),
	} {
		msg += "**" + c.String(strings.ToLower(empName)) + "**"
	}

	if len(posName) > 0 {
		msg += fmt.Sprintf(" si **%s** nya", posName)
	}

	msg += fmt.Sprintf(" dept **%s**", deptName)

	msg += " ulang tahun lhoo! 🎂"

	return msg
}
