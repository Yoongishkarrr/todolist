/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"todolist/task"

	"github.com/spf13/cobra"
)

var (
	id   int
	item string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "just add smth to do smth.. easy right? :)",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		taskID := strconv.Itoa(id)
		newTask := task.NewTask(taskID + ": " + item)
		fmt.Println("Task added:", newTask)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&id, "id", "i", 0, "Task ID")
	addCmd.Flags().StringVarP(&item, "item", "I", "", "Task description")
}

// Here you will define your flags and configuration settings.

// Cobra supports Persistent Flags which will work for this command
// and all subcommands, e.g.:
// addCmd.PersistentFlags().String("foo", "", "A help for foo")

// Cobra supports local flags which will only run when this command
// is called directly, e.g.:
// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
