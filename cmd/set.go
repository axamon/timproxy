// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"

	"github.com/axamon/cripta"
	"github.com/howeyc/gopass"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inserisci la matricola:\n")

		text := bufio.NewReader(os.Stdin)
		matricola, _ := text.ReadString('\n')
		//fmt.Println(matricola) //debug
		matricola = strings.TrimSuffix(matricola, "\n")

		viper.Set("matricola", matricola)

		fmt.Printf("Inserisci password aziendale (non verrà mostrata a video): ")

		pass, err := gopass.GetPasswd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
		//La password viene maltrattata criptata compressa e basata a 64
		passcriptata := cripta.Cifra(string(pass), "")
		//fmt.Println(passcriptata) //Debug

		fmt.Println("Informazioni salvate nel file config.yaml")
		fmt.Println("Ora puoi lanciare 'timproxy on' per avviare un terminale con proxy attivo.")
		viper.Set("passcriptata", passcriptata)
		viper.Set("proxy", "@lelapomi.telecomitlaia.it:8080")
		viper.Set("cmdpath", "C:\\Windows\\System32\\cmd.exe")
		viper.WriteConfigAs("config.yaml")

	},
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
