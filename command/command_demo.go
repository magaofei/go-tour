package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
				love by spf13 and friends in Go.
				Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("cobra")
		inferCmd := exec.Command("infer")
		inferCmd.Stdin = strings.NewReader("some input")
		var out bytes.Buffer
		err := inferCmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("in all caps %q\n", out.String())

	},
}

func init() {

}

func main() {
	fmt.Println("start")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
