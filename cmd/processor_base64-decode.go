// Code generated by github.com/abhimanyu003/sttr/cmd/generate.go. DO NOT EDIT
// generated: Fri, 24 Sep 2021 22:31:53 +0000

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/abhimanyu003/sttr/processors"
	"github.com/spf13/cobra"
)


func init() {
	rootCmd.AddCommand(base64DecodeCmd)
}

var base64DecodeCmd = &cobra.Command{
	Use:   "base64-decode",
	Short: "Decode your base64 text",
	Aliases: []string {"b64-dec", "b64-decode"},
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		in, out := "", ""

		if len(args) == 0 {
			all, err := ioutil.ReadAll(cmd.InOrStdin())
			if err != nil {
				return err
			}
			in = string(all)
		} else {
			in = args[0]
		}

		p := processors.Base64Decode{}
		flags := make([]processors.Flag, 0)

		out, err = p.Transform(in, flags...)
		if err != nil {
			return err
		}

		_, err = fmt.Fprint(os.Stdout, out)
		return err
	},
}
