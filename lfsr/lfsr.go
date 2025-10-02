package lfsr

import (
	"fmt"
	"strings"
)

type LFSR struct {
	Size     uint
	Taps     []uint
	Formulas []string
	Name     string
}

func NewLFSR(size uint, taps []uint, name string) *LFSR {
	formulas := make([]string, size)
	for i := range formulas {
		formulas[i] = "0"
	}
	return &LFSR{Size: size, Taps: taps, Formulas: formulas, Name: name}
}

func (l *LFSR) InjectBit(label string) {
	feedback := l.Formulas[l.Taps[0]]
	for _, tap := range l.Taps[1:] {
		feedback = fmt.Sprintf("%s ⊕ %s", feedback, l.Formulas[tap])
	}
	feedback = fmt.Sprintf("%s ⊕ %s", feedback, l.Formulas[l.Size-1])

	for i := l.Size - 1; i > 0; i-- {
		l.Formulas[i] = l.Formulas[i-1]
	}
	l.Formulas[0] = fmt.Sprintf("%s ⊕ %s", label, feedback)
}

func (l *LFSR) StateFormulas() string {
	var indexed []string
	for i := 0; i < int(l.Size); i++ {
		indexed = append(indexed,
			fmt.Sprintf("[%d]=%s", i, SimplifyXOR(l.Formulas[i])),
		)
	}
	return strings.Join(indexed, "; ")
}
