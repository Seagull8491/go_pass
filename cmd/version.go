/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"go_pass/util"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func getProjectPath() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	projectRootDir := filepath.Join(currentDir, "..")
	return projectRootDir, nil
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// projectRootDir, _ := getProjectPath()
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println("version is unknown")
		}

		versionFilePath := filepath.Join(currentDir, "version.txt")
		version, err := util.FileRead(versionFilePath)
		if err == nil {
			fmt.Println("version:", version)
		} else {
			fmt.Println("version is unknown")
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
