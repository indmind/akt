// Package akt contains utility functions for akt
package akt

import (
	"bytes"
	"strconv"
	"strings"
)

func unit(u string) string {
	switch u {
	case "0":
		return "nol"
	case "1":
		return "satu"
	case "2":
		return "dua"
	case "3":
		return "tiga"
	case "4":
		return "empat"
	case "5":
		return "lima"
	case "6":
		return "enam"
	case "7":
		return "tujuh"
	case "8":
		return "delapan"
	case "9":
		return "sembilan"
	}
	return ""
}

func tens(t string) string {
	temp := strings.Split(t, "")
	head, tail := temp[0], temp[1]

	if t == "10" {
		return "sepuluh"
	} else if t == "11" {
		return "sebelas"
	} else if head == "1" {
		return unit(tail) + " belas"
	} else if head == "0" {
		return unit(tail)
	} else if tail == "0" {
		return unit(head) + " puluh"
	}

	return unit(head) + " puluh " + unit(tail)
}

func hundreds(h string) string {
	temp := strings.Split(h, "")
	head, mid, last := temp[0], temp[1], temp[2]

	tail := mid + last

	if h == "100" {
		return "seratus"
	} else if tail == "00" {
		return unit(head) + " ratus"
	} else if head == "0" || head == "1" {
		return "seratus " + tens(tail)
	}

	return unit(head) + " ratus " + tens(tail)
}

func reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func reverseSlice(ss []string) {
	last := len(ss) - 1

	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

func splitSubN(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))

	l := len(runes)

	for i, r := range runes {
		sub = sub + string(r)

		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}

	return subs
}

func splitNumber(num string) []string {
	reversed := reverse(num)
	splited := splitSubN(reversed, 3)

	reverseSlice(splited)

	for idx, val := range splited {
		splited[idx] = reverse(val)
	}

	return splited
}

func parseNumber(n string) string {
	num, err := strconv.Atoi(n)

	if err != nil {
		panic(err)
	}

	// 010 -> 10 from first conversion
	str := strconv.Itoa(num)

	if n == "000" {
		return ""
	} else if num < 10 {
		return unit(str)
	} else if num < 100 {
		return tens(str)
	} else if num < 1000 {
		return hundreds(str)
	}

	return ""
}

func threeZero(t []string) []string {
	reversed := make([]string, len(t))

	copy(reversed, t)

	reverseSlice(reversed)

	for idx, val := range reversed {
		j := idx + 1

		if len(val) < 1 {
			reversed[idx] = ""
			continue
		}

		switch j {
		case 1:
			reversed[idx] = val
		case 2:
			if val == "satu" {
				reversed[idx] = "seribu"
			} else {
				reversed[idx] = val + " ribu"
			}
		case 3:
			reversed[idx] = val + " juta"
		case 4:
			reversed[idx] = val + " milyar"
		case 5:
			reversed[idx] = val + " triliun"
		case 6:
			reversed[idx] = val + " kuadriliun"
		case 7:
			reversed[idx] = val + " kuantiliun"
		case 8:
			reversed[idx] = val + " sekstiliun"
		case 9:
			reversed[idx] = val + " septiliun"
		case 10:
			reversed[idx] = val + " oktiliun"
		case 11:
			reversed[idx] = val + " noniliun"
		case 12:
			reversed[idx] = val + " desiliun"
		}
	}

	reverseSlice(reversed)

	return reversed
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// Convert string of number to indonesian word (spelled out)
func Convert(number string) string {
	filtered := strings.Replace(number, ".", "", -1)
	splited := splitNumber(filtered)

	for idx, val := range splited {
		splited[idx] = parseNumber(val)
	}

	zerThree := threeZero(splited)

	sentence := strings.Join(zerThree, " ")

	return standardizeSpaces(sentence)
}
