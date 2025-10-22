package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const rootHelpMenu string = `
        ######## ####
     ###        ### ##
   ###  ######+   ## ##         ####             ###           ##    ###
  ##  ###  ######  ## ##      ########                         ##    ###
  ++ ++  ++-    ++ ++ ++      ####     ########  ### ########  ##    ### ########
  +  ++ ++      +. ++ ++       -###### ###   ### ### ###  ###  ##    ### ###  ###
  +. ++ ++    +++ ++  ++     .##   ### ###   ### ### ###  ###  ###  #### ###  ###
  -- --  ------  --  --       ######## ########  ### ###  ###   #######  #######+
  --  --   ------.  --                 ###                               ###
   --. ---        ---                  ###                               ###
     ---- --------

    MIT License - 2025 Elouann Hosta ehosta@student.42lyon.fr


SpinUp is a CLI tool that build fully configurable web apps.

It uses the templates you want for each stack of your project:
 - A front-end service (Vue.js, React, etc.)
 - A back-end service (Go, Python Flask, Python Django, etc.)

⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯
`

var rootCmd = &cobra.Command{
	Use:   "spin",
	Short: "SpinUp — Build customizable and full web apps in an instant.",
	Long:  rootHelpMenu,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
