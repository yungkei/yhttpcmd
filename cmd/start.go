/*
Copyright © 2024 yungkei

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yungkei/yhttpcmd/internal"
)

// startCmd represents the start command
var (
	cfgFile  string
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start yhttpcmd server",
		Long: `Start yhttpcmd server to convert command as http service, 
		It will read the configuration file and convert the commands set in the configuration into HTTP services,
		
		Ensure you run this within the root directory of your site.`,
		Run: func(cmd *cobra.Command, args []string) {
			initBanner()
			internal.Start(cfgFile)
		},
	}
)

func initBanner() {
	banner := `
       _     _   _                            _ 
 _   _| |__ | |_| |_ _ __   ___ _ __ ___   __| |
| | | | '_ \| __| __| '_ \ / __| '_ ' _ \ / _' |
| |_| | | | | |_| |_| |_) | (__| | | | | | (_| |
 \__, |_| |_|\__|\__| .__/ \___|_| |_| |_|\__,_|
 |___/              |_|
	
 Copyright © 2024 yungkei

 `
	fmt.Printf("%s\n", banner)
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.yhttpcmd.yaml)")
}
