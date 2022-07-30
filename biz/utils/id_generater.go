package utils

import "github.com/bwmarrin/snowflake"

var idGenerator *snowflake.Node

func init() {
	var err error
	idGenerator, err = snowflake.NewNode(999)
	if err != nil {
		panic(err)
	}
}

func GetID() uint64 {
	return uint64(idGenerator.Generate().Int64())
}
