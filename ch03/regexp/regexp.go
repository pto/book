package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	names := []string{"André Miró", "Èduard Manet", "Jean Garçon Rhône", "J. K. F. Äsbourg"}
	nameRx := regexp.MustCompile(`(?P<forenames>\pL+\.?(?:\s+\pL+\.?)*)\s+(?P<surname>\pL+)`)
	fmt.Printf("TRACE: nameRx.SubexpNames() is %q\n", nameRx.SubexpNames())
	for i := range names {
		names[i] = nameRx.ReplaceAllString(names[i], "${surname}, ${forenames}")
	}
	for _, name := range names {
		fmt.Println(name)
	}
	fmt.Println()

	text := "Now is the the time for good good men to come to the aid of their country."
	wordRx := regexp.MustCompile(`\w+`)
	if matches := wordRx.FindAllString(text, -1); matches != nil {
		previous := ""
		for _, match := range matches {
			if match == previous {
				fmt.Println("Duplicate word:", match)
			}
			previous = match
		}
	}
	fmt.Println()

	lines := `
		first: 1
		second   : two
		Third : trés
		Ülm : Not a tree
		`
	valueForKey := make(map[string]string)
	keyValueRx := regexp.MustCompile(`\s*(\pL[\pL\p{Nd}_]*)\s*:\s*(.+)`)
	if matches := keyValueRx.FindAllStringSubmatch(lines, -1); matches != nil {
		for _, match := range matches {
			valueForKey[match[1]] = strings.TrimRight(match[2], "\t ")
		}
	}
	for value, key := range valueForKey {
		fmt.Printf("%s: %s\n", value, key)
	}
	fmt.Println()

	attrName := "thing"
	attribs := "this='that' hi='there' thing='hello' other='none' thing=\"world\" more='less'"
	attrValueRx := regexp.MustCompile(regexp.QuoteMeta(attrName) +
		`=(?:"([^"]+)"|'([^']+)')`)
	if indexes := attrValueRx.FindAllStringSubmatchIndex(attribs, -1); indexes != nil {
		fmt.Printf("TRACE: %v\n\n", indexes)
		for _, positions := range indexes {
			start, end := positions[2], positions[3]
			if start == -1 {
				start, end = positions[4], positions[5]
			}
			fmt.Printf("'%s'\n", attribs[start:end])
		}
	}
	fmt.Println()

	latin1 := "Éclairs are passé, but ömlauts are dérécherché"
	unaccentedLatin1Rx := regexp.MustCompile(
		`[ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖØÙÚÛÜÝàáâãäåæçèéêëìíîïñðòóôõöøùúûüýÿ]+`)
	fmt.Println(unaccentedLatin1Rx.ReplaceAllStringFunc(latin1, UnaccentedLatin1))
}

func UnaccentedLatin1(s string) string {
	chars := make([]rune, 0, len(s))
	for _, char := range s {
		switch char {
		case 'À', 'Á', 'Â', 'Ã', 'Ä', 'Å':
			char = 'A'
		case 'Æ':
			chars = append(chars, 'A')
			char = 'E'
		case 'Ç':
			char = 'C'
		case 'È', 'É', 'Ê', 'Ë':
			char = 'E'
		case 'Ì', 'Í', 'Î', 'Ï':
			char = 'I'
		case 'Ð':
			char = 'D'
		case 'Ñ':
			char = 'N'
		case 'Ò', 'Ó', 'Ô', 'Ö', 'Õ', 'Ø':
			char = 'O'
		case 'Ù', 'Ú', 'Û', 'Ü':
			char = 'U'
		case 'Ý':
			char = 'Y'
		case 'à', 'á', 'â', 'ã', 'ä', 'å':
			char = 'a'
		case 'æ':
			chars = append(chars, 'a')
			char = 'e'
		case 'ç':
			char = 'c'
		case 'è', 'é', 'ê', 'ë':
			char = 'e'
		case 'ì', 'í', 'î', 'ï':
			char = 'i'
		case 'ð':
			char = 'd'
		case 'ñ':
			char = 'n'
		case 'ò', 'ó', 'ô', 'ö', 'õ', 'ø':
			char = 'o'
		case 'ù', 'ú', 'û', 'ü':
			char = 'u'
		case 'ý', 'ÿ':
			char = 'y'
		}
		chars = append(chars, char)
	}
	return string(chars)
}
