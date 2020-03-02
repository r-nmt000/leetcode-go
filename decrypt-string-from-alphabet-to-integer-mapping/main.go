package main

import (
	"fmt"
	"strings"
)

func main() {
	//ex1 := "123410#11#12#20483726#987"
	//ex2 := "10#11#12"
	ex3 := "1326#"
	fmt.Println(freqAlphabets(ex3))
	//fmt.Println(splitStringWithSharp(ex2, false))
	//str := "12345"[2]
	//fmt.Println(str)
	//fmt.Println(string(str))
}

func freqAlphabets(s string) string {
	// シャープが見つかる単位で区切る
	// シャープから前の2つは確実にひとかたまり
	// それより前は1つずつの数字とみなせる
	// 123410#11#12#20483726#987
	// 123410# 11# 12# 20483726# 987
	// 1 2 3 4 10# 11# 12# 2 0 4 8 3 7 26# 9 8 7
	ret := ""
	splitted := splitStringWithSharp(s, hasLastSharp(s))
	for i := 0; i < len(splitted); i++ {
		v := splitted[i]
		if string(v[len(v)-1]) == "#" {
			for j := 0; j < len(v)-3; j++ {
				encrypted, _ := mapping[string(v[j])]
				ret += encrypted
			}
			encrypted, _ := mapping[string(v[len(v)-3:])]
			ret += encrypted
		} else {
			for j := 0; j < len(v); j++ {
				encrypted, _ := mapping[string(v[j])]
				ret += encrypted
			}
		}
	}
	return ret
}

func splitStringWithSharp(s string, hasLastSharp bool) []string {
	splitted := strings.Split(s, "#")
	if splitted[len(splitted)-1] == "" {
		splitted = splitted[:len(splitted)-1]
	}
	for i := 0; i < len(splitted); i++ {
		if i == len(splitted)-1 && !hasLastSharp {
			continue
		}
		splitted[i] = splitted[i] + "#"
	}
	return splitted
}

func hasLastSharp(s string) bool {
	lastCharacter := string(s[len(s)-1])
	if lastCharacter == "#" {
		return true
	}
	return false
}

var mapping = map[string]string{
	"1":   "a",
	"2":   "b",
	"3":   "c",
	"4":   "d",
	"5":   "e",
	"6":   "f",
	"7":   "g",
	"8":   "h",
	"9":   "i",
	"10#": "j",
	"11#": "k",
	"12#": "l",
	"13#": "m",
	"14#": "n",
	"15#": "o",
	"16#": "p",
	"17#": "q",
	"18#": "r",
	"19#": "s",
	"20#": "t",
	"21#": "u",
	"22#": "v",
	"23#": "w",
	"24#": "x",
	"25#": "y",
	"26#": "z",
}
