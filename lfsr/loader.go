package lfsr

import "fmt"

func LoadKeyAndFrame(l *LFSR, history *[]string) {
	for i := 0; i < 64; i++ {
		l.InjectBit(fmt.Sprintf("K%d", i))
		*history = append(*history, l.StateFormulas())
	}
	for i := 0; i < 22; i++ {
		l.InjectBit(fmt.Sprintf("F%d", i))
		*history = append(*history, l.StateFormulas())
	}
}
