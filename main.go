package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	vocals     = []string{"a", "e", "i", "o", "u"}
	consonants = []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z"}
	titles     = []string{"Lord", "Dark", "Mistress", "King", "Queen", "Lady", "Mister", "Little", "Big"}
	worlds     = []string{"Adra", "Alumbra", "Antica", "Ardera", "Astera", "Bastia", "Batabra", "Belobra", "Bona", "Cadebra", "Calmera", "Celebra", "Celesta", "Collabra", "Damora", "Descubra", "Dibra", "Epoca", "Famosa", "Fera", "Ferobra", "Firmera", "Gentebra", "Gladera", "Harmonia", "Havera", "Honbra", "Illusera", "Impulsa", "Inabra", "Kalibra", "Karna", "Libertabra", "Lobera", "Luminera", "Lutabra", "Marbera", "Marcia", "Menera", "Monza", "Mudabra", "Mykera", "Nefera", "Nossobra", "Ocebra", "Olima", "Ombra", "Optera", "Pacera", "Peloria", "Premia", "Quelibra", "Quintera", "Refugia", "Reinobra", "Seanera", "Secura", "Serdebra", "Solidera", "Suna", "Talera", "Tembra", "Thyria", "Trona", "Utobra", "Venebra", "Versa", "Visabra", "Vunira", "Wintera", "Wizera", "Xandebra", "Yonabra", "Zenobra", "Zuna", "Zunera"}
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	length := flag.Int("length", getRandomInt(5, 12), "length of name")
	includeWorld := flag.Bool("include-world", false, "include world in name")
	includeTitle := flag.Bool("include-title", false, "include title in name")
	count := flag.Int("count", 1, "number of result")
	flag.Parse()

	for i := 0; i < *count; i += 1 {
		fmt.Println(createRandomName(*length, *includeTitle, *includeWorld))
	}
}

func getRandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func createRandomName(length int, includeTitle, includeWorld bool) string {
	// Create a nickname consiting of consonants+vocals+consonants+vocals etc..
	var chars []string
	for i := 0; i < length; i += 1 {
		if i%2 != 0 {
			chars = append(chars, vocals[getRandomInt(0, len(vocals)-1)])
		} else {
			chars = append(chars, consonants[getRandomInt(0, len(consonants)-1)])
		}
	}

	name := strings.Title(strings.Join(chars, ""))

	if includeTitle {
		name = fmt.Sprintf("%s %s", titles[getRandomInt(0, len(titles)-1)], name)
	}

	if includeWorld {
		name = fmt.Sprintf("%s of %s", name, worlds[getRandomInt(0, len(worlds)-1)])
	}

	return name
}
