package main

import (
	"fmt"
	"golang/beginners/internal/mutants"
	"strings"
)

func main() {

	mutantDNA := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	humanDNA := []string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"}

	isMutant, err := mutants.IsMutant(mutantDNA)

	if err != nil {
		fmt.Printf("%s\n", err)
	}

	if isMutant {
		fmt.Printf("DNA [%s] is a mutant\n", strings.Join(mutantDNA, ","))
	} else {
		fmt.Printf("DNA [%s] is a human\n", strings.Join(mutantDNA, ","))
	}

	isMutant, err = mutants.IsMutant(humanDNA)

	if err != nil {
		fmt.Printf("%s", err)
	}

	if isMutant {
		fmt.Printf("DNA [%s] is a mutant\n", strings.Join(humanDNA, ","))
	} else {
		fmt.Printf("DNA [%s] is a human\n", strings.Join(humanDNA, ","))
	}
}
