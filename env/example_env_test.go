package env

import (
	"fmt"
	"os"
)

func ExampleNewConfigFromEnv() {
	os.Setenv("TWELVE_EXAMPLE_DB", "mongodb://localhost/myapp")
	os.Setenv("TWELVE_EXAMPLE_DEBUG", "true")
	os.Setenv("TWELVE_EXAMPLE_REDISDB", "1")
	config, _ := NewConfigFromEnv(map[string]VarSpec{
		"db":    {Name: "TWELVE_EXAMPLE_DB"},
		"debug": {Name: "TWELVE_EXAMPLE_DEBUG"},
		"env":   {Name: "TWELVE_EXAMPLE_ENV", DefaultValue: "development"},
		"cache": {
			Name: "TWELVE_EXAMPLE_REDISDB",
			Parser: func(val string) string {
				return fmt.Sprintf("redis://localhost:6379/%s", val)
			},
		},
	})
	fmt.Println("db is:", config["db"])
	fmt.Println("debug is:", config.AsFlag("debug"))
	fmt.Println("env is:", config["env"])
	fmt.Println("cache is:", config["cache"])

	// Output:
	// db is: mongodb://localhost/myapp
	// debug is: true
	// env is: development
	// cache is: redis://localhost:6379/1
}
