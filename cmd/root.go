/*
Copyright © 2022 Muhammad Yasser <mdmmdyssraamr@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/mdyssr/pray/pkg/utils"
	"github.com/mdyssr/prayers"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pray",
	Short: "Prayer Times CLI",
	Long: `Prayer Times CLI that shows the Islamic prayer times in your area.

For any bugs or suggestions, don't hesitate to
email me at mhmmdyssraamr@gmail.com`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		showPrayerTimes()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func showPrayerTimes() {
	prayersData, err := prayers.GetPrayersData()
	if err != nil {
		fmt.Println("Error getting prayers data...")
		os.Exit(1)
	}

	utils.PrintFormattedPrayersData(prayersData)
}
