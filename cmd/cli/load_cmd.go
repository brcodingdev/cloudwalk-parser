package main

import (
	"context"
	"fmt"
	"github.com/brcodingdev/cloudwalk-parser/internal/pkg/app"
	"github.com/brcodingdev/cloudwalk-parser/internal/ports/repository"
	"github.com/spf13/cobra"
	"log"
	"os"
	"runtime/pprof"
)

var absLogFile string

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load log file",
	Run: func(c *cobra.Command, args []string) {
		cpu, err := os.Create(".cpu.prof")
		if err != nil {
			log.Fatal(err)
		}

		pprof.StartCPUProfile(cpu)
		defer cpu.Close()

		client := loadRepository()
		ctx := context.Background()
		repository := repository.NewRedisMatch(ctx, client)
		repository.Clean()
		app := app.NewMatchFileProcess(absLogFile, repository)

		err = app.Load()
		if err != nil {
			log.Fatalf("%v", err)
			return
		}

		fmt.Print("file processed ", absLogFile)

		mem, err := os.Create(".mem.prof")
		if err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
		pprof.WriteHeapProfile(mem)
		defer mem.Close()
	},
}

func init() {
	loadCmd.PersistentFlags().StringVarP(
		&absLogFile,
		"file",
		"f",
		"",
		"Log app to be processed",
	)
}
