package normal

import (
	"fmt"
	"go-redis-example/config"

	"github.com/go-redis/redis/v8"
)

// RunNormalMode ...
func RunNormalMode() {
	var ctx = config.Config.Ctx
	var rdb = config.Config.Client

	/*
		SET
	*/
	// Set a new key-value
	// Using with Result() method that returns message and error about operation
	message, err := rdb.Set(ctx, "user_name", "safran", 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("[SET] Message --> ", message)

	// Using with Err() method that returns only error if needs.
	err = rdb.Set(ctx, "name", "juno", 0).Err()
	if err != nil {
		panic(err)
	}

	err = rdb.Set(ctx, "age", 33, 0).Err()
	if err != nil {
		panic(err)
	}

	/*
		GET
	*/
	// Get "user_name" key's value
	// Result() returns value of key and error
	val, err := rdb.Get(ctx, "user_name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("[GET] user_name value ->", val)

	/*
		DEL
	*/
	// DEL "user_name" key

	msg, err := rdb.Del(ctx, "age").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("[DEL] message: ", msg) // Output: 1 (means OK)

	/*
		EXIST
	*/
	// Check whether "age" key exists or not
	valAge, err := rdb.Get(ctx, "age").Result()
	if err == redis.Nil {
		fmt.Println("[GET-EXIST] age does not exist", err)
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("age value --> ", valAge)
	}

	// Another way for checking
	Exists := rdb.Exists(ctx, "age")
	fmt.Println("[Exist] message -> ", Exists) // Output: exists age: 0 (0 means false)

	/*
		APPEND
	*/
	// Add string to name
	valueAppend, err := rdb.Append(ctx, "name", " The Cat").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("[APPEND] message -> ", valueAppend) // Output: 12 (length of new string)

	// Get new value of key
	val, err = rdb.Get(ctx, "name").Result()
	fmt.Println("New value of key --> ", val) // juno The Cat

}
