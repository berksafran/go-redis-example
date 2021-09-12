package publisher

import (
	"fmt"
	"go-redis-example/config"
)

// Publish a message on selected channel
func Publish(message string) {
	var ctx = config.Config.Ctx
	var rdb = config.Config.Client
	var channelName = config.Config.ChannelName

	result, err := rdb.Publish(ctx, channelName, message).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Publish Result --> ", result) // Output: 1 (OK)
}
