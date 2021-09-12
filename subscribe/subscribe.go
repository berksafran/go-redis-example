package subscribe

import (
	"fmt"
	"go-redis-example/config"
)

// Sub subscribes a channel
func Sub() {
	var ctx = config.Config.Ctx
	var rdb = config.Config.Client
	var channelName = config.Config.ChannelName

	sub := rdb.Subscribe(ctx, channelName)
	result, err := sub.Receive(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Subcribed to", result, "channel.")

	fmt.Println("Listening on", channelName, "channel ...")

	message, err := sub.ReceiveMessage(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("New incoming message from Publisher -->", message.Payload)
}
