package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAssert(t *testing.T) {
	var (
		input  = 1
		output = 3
	)

	actual := AddOne(1)
	if actual != output {
		t.Errorf("AddOne(%d), input %d, actual = %d", input, output, actual)
	}
	// assert.Equal(t, 2, 3, "AddOne(2) should be 3")
	// fmt.Println("Not executing")
}

func TestRequire(t *testing.T) {
	require.Equal(t, 2, 3)
	fmt.Println("Not executing")
}

// go tool cover -html=coverage -o coverage.html
// go tool cover -html=coverage -o coverage.html
// Mac open .\coverage.html  - Wins start .\coverage.html
