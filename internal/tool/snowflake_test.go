package tool

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSnowflake(t *testing.T) {
	snowflake, err := NewSnowflake(0, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	idMap := make(map[int64]struct{}, 0)
	for i := 0; i < 5; i++ {
		id := snowflake.NextVal()
		_, found := idMap[id]
		assert.False(t, found)
		idMap[id] = struct{}{}
		fmt.Printf("%08b\n", id)
	}
}
