package mutants

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_IsMutant(t *testing.T) {
	tests := []struct {
		Name          string
		DNA           []string
		Expected      bool
		ExpectedError error
	}{
		{
			Name:          "Given_Is_Mutant_With_Error_Empty_Dna",
			DNA:           []string{},
			Expected:      false,
			ExpectedError: errors.New("empty dna"),
		},
		{
			Name:          "Given_Is_Mutant_With_Error_Dna_Length_Not_Valid",
			DNA:           []string{"ATGCG", "CTGC", "TTGT", "AGAAGG", "CCCCTA", "TCACTG"},
			Expected:      false,
			ExpectedError: errors.New("dna length not valid"),
		},
		{
			Name:          "Given_Is_Mutant_With_Error_Dna_Chars_Not_Valid",
			DNA:           []string{"ATGCGA", "CXGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
			Expected:      false,
			ExpectedError: errors.New("dna chars not valid"),
		},
		{
			Name:     "Given_Is_Mutant",
			DNA:      []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
			Expected: true,
		},
		{
			Name:     "Given_Is_Human",
			DNA:      []string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"},
			Expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			isMutant, err := IsMutant(test.DNA)

			require.Equal(t, test.Expected, isMutant)
			require.Equal(t, test.ExpectedError, err)
		})
	}
}
