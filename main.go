package main

import "fmt"

// Задана строка S из малых латинских букв, требуется узнать
// длину наибольшей подстроки, в которой все буквы оьдинаковы.

// "aabaa"
// a - 4
// b - 1

// "aabaaacaaaaa"

func max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

// MaxSequenceLength determine max length of char sequence in string
func MaxSequenceLength(s string) int64 {
    var final int64 = 0
    if s == "" {
        return final
    }
    var counter int64 = 0
    var lastChar byte

	for i := 0; i < len(s); i++ {
		if lastChar == s[i] {
            counter++
            final = max(counter, final)
		} else {
            lastChar = s[i]
            counter = 1
        }
	}

	return final
}

func main() {
	fmt.Println(MaxSequenceLength("aabaa"))
	fmt.Println(MaxSequenceLength("aabaaa"))
}

// "aabaaa"
