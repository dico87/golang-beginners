package mutants

import (
	"errors"
	"strings"
	"sync"
)

const validChars = "ATCG"
const validDnaA = "AAAA"
const validDnaT = "TTTT"
const validDnaC = "CCCC"
const validDnaG = "GGGG"

func IsMutant(dna []string) (bool, error) {

	newDna := upperDNA(dna)
	err := checkDNASize(newDna)
	if err != nil {
		return false, err
	}

	var wg sync.WaitGroup
	var lock sync.RWMutex
	counter := 0
	wg.Add(3)
	go func() {
		defer lock.RUnlock()
		lock.RLock()
		counter = counter + checkHorizontals(newDna)
		wg.Done()
	}()
	go func() {
		defer lock.RUnlock()
		lock.RLock()
		counter = counter + checkVerticals(newDna)
		wg.Done()
	}()
	go func() {
		defer lock.RUnlock()
		lock.RLock()
		counter = counter + checkDiagonal(newDna)
		wg.Done()
	}()

	wg.Wait()

	if counter > 0 {
		return true, nil
	}

	return false, nil
}

func checkDiagonal(dna []string) int {
	matrix := convertToMatrix(dna)
	diagonals := getDiagonals(matrix)
	counter := 0
	for _, row := range diagonals {
		if strings.Contains(row, validDnaA) ||
			strings.Contains(row, validDnaT) ||
			strings.Contains(row, validDnaC) ||
			strings.Contains(row, validDnaG) {
			counter++
		}
	}

	return counter
}

func convertToMatrix(dna []string) [][]string {
	rows := len(dna)
	columns := len(dna[0])
	matrix := make([][]string, rows)

	// fill de matrix because golang needs this
	for i := range matrix {
		matrix[i] = make([]string, columns)
	}

	for i := 0; i < rows; i++ {
		chars := []rune(dna[i])
		for j := 0; j < len(chars); j++ {
			matrix[i][j] = string(chars[j])
		}
	}
	return matrix
}

func getDiagonals(matrix [][]string) []string {
	rows := len(matrix)
	columns := len(matrix[0])
	var diagonals []string

	// check half diagonals
	for k := 0; k <= (rows - 1); k++ {
		i := k
		j := 0
		diagonal := ""
		for i >= 0 {
			diagonal = diagonal + matrix[i][j]
			i--
			j++
		}
		diagonals = append(diagonals, diagonal)
	}

	for k := 1; k <= (columns - 1); k++ {
		i := rows - 1
		j := k
		diagonal := ""
		for j <= (columns - 1) {
			diagonal = diagonal + matrix[i][j]
			i--
			j++
		}
		diagonals = append(diagonals, diagonal)
	}

	return diagonals
}

func checkVerticals(dna []string) int {
	pivotDna := pivot(dna)
	return checkHorizontals(pivotDna)
}

func pivot(dna []string) []string {
	dnaLength := len(dna[0])
	var pivotDna []string
	reverseDna := ""
	for i := 0; i < dnaLength; i++ {
		reverseDna = ""
		for j := 0; j < len(dna); j++ {
			chars := []rune(dna[j])
			reverseDna += string(chars[i])
		}
		pivotDna = append(pivotDna, reverseDna)
	}

	return pivotDna
}

func checkHorizontals(dna []string) int {
	count := 0
	for _, row := range dna {
		if strings.Contains(row, validDnaA) ||
			strings.Contains(row, validDnaT) ||
			strings.Contains(row, validDnaC) ||
			strings.Contains(row, validDnaG) {
			count++
		}
	}

	return count
}

func upperDNA(dna []string) []string {
	var upper []string
	for _, row := range dna {
		upper = append(upper, strings.ToUpper(row))
	}

	return upper
}

func checkDNASize(dna []string) error {
	if len(dna) < 1 {
		return errors.New("empty dna")
	}

	length := len(dna[0])
	for i := 1; i < len(dna); i++ {
		if len(dna[i]) != length {
			return errors.New("dna length not valid")
		}
	}

	for _, row := range dna {
		chars := []rune(row)
		for _, char := range chars {
			if !strings.ContainsAny(string(char), validChars) {
				return errors.New("dna chars not valid")
			}
		}
	}

	return nil
}
