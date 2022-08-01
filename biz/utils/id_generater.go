package utils

import "github.com/bwmarrin/snowflake"

var (
	userIDGenerator    *snowflake.Node
	messageIDGenerator *snowflake.Node
	groupIDGenerator   *snowflake.Node
)

func init() {
	var err error
	userIDGenerator, err = snowflake.NewNode(100)
	messageIDGenerator, err = snowflake.NewNode(200)
	groupIDGenerator, err = snowflake.NewNode(300)
	if err != nil {
		panic(err)
	}
}

func GenerateUserID() uint64 {
	return uint64(userIDGenerator.Generate().Int64())
}

func GenerateGroupID() uint64 {
	return uint64(groupIDGenerator.Generate().Int64())
}

func GenerateMessageID() uint64 {
	return uint64(messageIDGenerator.Generate().Int64())
}
