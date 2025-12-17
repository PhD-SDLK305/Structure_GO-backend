package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOne(t *testing.T) {
	// var (
	// 	input  = 1
	// 	output = 2
	// )

	// actual := AddOne(1)
	// if actual != output {
	// 	t.Errorf("AddOne(%d), output %d, actual = %d", input, output, actual)
	// }
	assert.Equal(t, AddOne(2), 4, "AddOne(2) should be 3")
}

// func TestRequire(t *testing.T) {
// 	require.Equal(t, 2, 3)
// 	println("Not executing")
// }

// func TestAssert(t *testing.T) {
// 	assert.Equal(t, 2, 3)
// 	println("executing")
// }
