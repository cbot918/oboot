/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/cbot918/oboot/src/pkg/infra"
	"github.com/spf13/cobra"
)

// gendcCmd represents the gendc command
var gendcCmd = &cobra.Command{
	Use:   "gendc",
	Short: "gendc short des",
	Long:  `gendc long des`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := gendc(); err != nil {
			fmt.Println("gen docker-compose failed")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(gendcCmd)
}

func gendc() error {
	infraName := "test"
	inputFile := "infra.o"
	outputFile := "docker-compose.yaml"

	infra, err := infra.NewInfra(infraName, inputFile)
	if err != nil {
		return err
	}

	err = infra.DockerCompose.WriteDockercompose(outputFile)
	if err != nil {
		return err
	}

	return nil
}
