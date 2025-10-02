package main

import (
	"A5/lfsr"
	"fmt"
)

func main() {
	R1 := lfsr.NewLFSR(19, []uint{17, 16, 13}, "R1")
	R2 := lfsr.NewLFSR(22, []uint{20}, "R2")
	R3 := lfsr.NewLFSR(23, []uint{21, 20, 7}, "R3")

	var historyForR1, historyForR2, historyForR3 []string

	lfsr.LoadKeyAndFrame(R1, &historyForR1)
	lfsr.LoadKeyAndFrame(R2, &historyForR2)
	lfsr.LoadKeyAndFrame(R3, &historyForR3)

	if err := lfsr.WriteToFile("R1.txt", historyForR1); err != nil {
		fmt.Println(err)
	}
	if err := lfsr.WriteToFile("R2.txt", historyForR2); err != nil {
		fmt.Println(err)
	}
	if err := lfsr.WriteToFile("R3.txt", historyForR3); err != nil {
		fmt.Println(err)
	}
}
