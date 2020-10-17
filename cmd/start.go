/*
Copyright © 2020 Scott Xiong <mbp98k@h=gmail.com>
This file is part of CLI application foo.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var daemon bool

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your start command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if daemon {
			command := exec.Command("test", "start")
			command.Start()
			fmt.Printf("test start, [PID] %d running...\n", command.Process.Pid)
			ioutil.WriteFile("/Users/scottxiong/go/src/github.com/scott-x/deamon-demo-go/cmd.lock", []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)
			daemon = false
			os.Exit(0)
		} else {
			fmt.Println("gonne start")
		}

		fmt.Println("server is running at: http://127.0.0.1:9006")
		startHttp()
	},
}

func init() {
	startCmd.Flags().BoolVarP(&daemon, "deamon", "d", false, "is daemon?")
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func startHttp() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello cmd!")
	})
	if err := http.ListenAndServe(":9006", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
