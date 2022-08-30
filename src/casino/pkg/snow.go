package pkg

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

// function generates id-num and returns it

func Generate() snowflake.ID {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	id := node.Generate()
	return id
}
