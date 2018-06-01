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
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/viper"

	"github.com/axamon/cripta"

	"github.com/spf13/cobra"
)

// onCmd represents the on command
var onCmd = &cobra.Command{
	Use:   "on",
	Short: "Attiva il proxy",
	Long:  `Attiva il proxy selezionato in una nuova shell`,
	Run: func(cmd *cobra.Command, args []string) {

		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		err := viper.ReadInConfig() // Find and read the config file
		if err != nil {             // Handle errors reading the config file
			fmt.Println("Devi prima lanciare 'timproxy set' per creare il file di configurazione.")
			os.Exit(1)
		}

		passwordaziendale := cripta.Decifra(viper.GetString("passcriptata"), "")
		//fmt.Println(result) //Debug
		matricola := viper.GetString("matricola")
		//fmt.Println(matricola) //Debug
		host := viper.GetString("proxy")

		httpproxy := "http://" + matricola + ":" + passwordaziendale + host
		httpsproxy := "https://" + matricola + ":" + passwordaziendale + host

		os.Setenv("HTTP_PROXY", httpproxy)
		os.Setenv("HTTPS_PROXY", httpsproxy)

		cmdpath := viper.GetString("cmdpath")
		cmdshell := exec.Command(cmdpath, "/C", "start", cmdpath)
		errshell := cmdshell.Start()
		if errshell != nil {
			fmt.Fprintln(os.Stderr, "Error", errshell.Error())
		}

		//fmt.Println(httpproxy, httpsproxy) //Debug
		//fmt.Println("on called")
	},
}

func init() {
	rootCmd.AddCommand(onCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// onCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// onCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
