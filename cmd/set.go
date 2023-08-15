/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"go_pass/model"
	"go_pass/util"
	"os"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set new password.",
	Long:  ``,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("set called")
	// },
	RunE: runInteractive,
}

func runInteractive(cmd *cobra.Command, args []string) error {
	fmt.Println("set password info.")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("entity name: ")
	entityName, _ := reader.ReadString('\n')
	fmt.Println("entityName: ", entityName)
	fmt.Print("password: ")
	password, _ := reader.ReadString('\n')
	fmt.Printf("the password is %s", password)

	key := []byte("example-key-1234")
	passEntity, err := model.New(entityName, password, key)
	if err != nil {
		fmt.Println("Creating password entity is failed.")
		return err
	}

	passEntityStr, err := passEntity.GetPassEntryString()
	if err != nil {
		return err
	}

	err = util.FileWrite("./password.txt", ([]byte)(passEntityStr))
	if err != nil {
		fmt.Println("Cannot save password.")
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
