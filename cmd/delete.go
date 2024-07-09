/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"
	"todolist/task"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete as delete",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		success := task.DeleteTask(id)
		if success {
			fmt.Println("Task deleted successfully")
		} else {
			fmt.Println("Task not found")
		}
	},
}

var (
	deleteID int
)

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().IntVarP(&deleteID, "id", "i", 0, "")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
