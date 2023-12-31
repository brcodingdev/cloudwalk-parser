package main

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Parser to read log app was generated by a Quake 3 Arena server",
	Run: func(c *cobra.Command, args []string) {
	},
	PersistentPreRun: func(c *cobra.Command, args []string) {
		cmdCalled := c.CalledAs()
		if cmdCalled == "version" || cmdCalled == "help" {
			return
		}
		var err error

		if err != nil {
			os.Exit(1)
		}
	},
}

func loadRepository() *redis.Client {
	host := os.Getenv("HOST_REDIS")
	if len(host) == 0 {
		host = "localhost:6379"
	}

	fmt.Println("Redis host ", host)
	return redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func init() {
	rootCmd.AddCommand(loadCmd)
	rootCmd.AddCommand(reportCmd)
}
