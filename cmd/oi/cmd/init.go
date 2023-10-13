/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init short des",
	Long:  `init long des`,
	Run: func(cmd *cobra.Command, args []string) {

		fileName := "infra.o"

		if err := Init(fileName); err != nil {
			fmt.Println("init failed")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func Init(fileName string) error {
	fd, err := os.Create(fileName)
	if err != nil {
		return err
	}

	_, err = fd.Write([]byte(InitContent()))
	if err != nil {
		return err
	}
	return nil
}

func InitContent() string {
	return `network: bridge
service: [redis, rabbitmq, nginx]
`
}
