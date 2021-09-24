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

var json_flag_i bool

func init() {	
	jsonCmd.Flags().BoolVarP(&json_flag_i, "indent", "i", false, "Indent the output (prettyprint)")
	rootCmd.AddCommand(jsonCmd)
}

var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "Format your text as JSON",
	Aliases: []string {},
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

		p := processors.FormatJSON{}
		flags := make([]processors.Flag, 0)
		flags = append(flags, processors.Flag{Short: "i", Value: json_flag_i})

		out, err = p.Transform(in, flags...)
		if err != nil {
			return err
		}

		_, err = fmt.Fprint(os.Stdout, out)
		return err
	},
}
