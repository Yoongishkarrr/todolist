/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"todolist/task"

	"github.com/spf13/cobra"
)

// addCmd представляет команду add
var addCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "Add a new task to the to-do list",
	Long:  `This command allows you to add a new task to your to-do list.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := args[0]

		newTask := task.NewTask(description) 
		task.AddTask(newTask) 

		fmt.Printf("Task added: ID: %d, Description: %s\n", newTask.ID, newTask.Description)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
