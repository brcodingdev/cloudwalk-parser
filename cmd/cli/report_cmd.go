package main

import (
	"context"
	"github.com/brcodingdev/cloudwalk-parser/internal/pkg/app"
	"github.com/brcodingdev/cloudwalk-parser/internal/ports/repository"
	"github.com/spf13/cobra"
	"log"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Report the last load log parse",
	Run: func(c *cobra.Command, args []string) {
		client := loadRepository()
		ctx := context.Background()
		repository := repository.NewRedisMatch(ctx, client)
		app := app.NewMatchReport(repository)

		_, err := app.PrintMatch()
		if err != nil {
			log.Fatalf("%v", err)
		}
	},
}
