/*
Copyright © 2020 Scott Xiong <mbp98k@h=gmail.com>
This file is part of CLI application foo.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "A brief description of your stop command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		strb, _ := ioutil.ReadFile("/Users/scottxiong/go/src/github.com/scott-x/deamon-demo-go/cmd.lock")
		command := exec.Command("kill", string(strb))
		command.Start()
		fmt.Println("test stoped")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
