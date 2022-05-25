package main

import (
	"banktest_account/src/cmd"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	defer stop()

	var rootCmd = &cobra.Command{Use: "app"}

	signApi := &cobra.Command{
		Use:   "start",
		Short: "Initial http server",
		Run: func(cli *cobra.Command, args []string) {
			cmd.StartHttp(ctx)
		},
	}

	rootCmd.AddCommand(signApi)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(context.Background(), err.Error())
	}
}
