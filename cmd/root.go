package cmd

import (
	"fmt"
	"github.com/XORbit01/webpalm/core"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   usage(),
	Short: "A web scraping tool",
	Long:  long(),
	Run: func(cmd *cobra.Command, args []string) {
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			//help message
			fmt.Println("Error: ", err)
			return
		}
		level, err := cmd.Flags().GetInt("level")
		if err != nil {
			//help message
			fmt.Println("Error: ", err)
			return
		}

		liveMode, err := cmd.Flags().GetBool("live")
		if err != nil {
			//help message
			fmt.Println("Error: ", err)
			return
		}

		if level < 0 {
			fmt.Println("Error: Level should be greater equal than 0")
			return
		}
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		exportFile, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		regexMap, err := cmd.Flags().GetStringToString("regexes")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		excludedStatus, err := cmd.Flags().GetIntSlice("exclude-code")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		fmt.Println(options(url, level, liveMode, exportFile, regexMap, excludedStatus))
		cr := core.NewCrawler(url, level, liveMode, exportFile, regexMap, excludedStatus)
		cr.Crawl()
	},
	Example: example() + regexestable(),
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "target url / ex: -u https://google.com")
	err := rootCmd.MarkFlagRequired("url")
	if err != nil {
		return
	}
	rootCmd.Flags().IntP("level", "l", 0, "level of palming / ex: -l 2")

	rootCmd.Flags().Bool("live", false, "live output mode (slow but live streaming) / ex: --live")

	rootCmd.Flags().StringP("output", "o", "", "file to export the result (f.json, f.xml, f.txt) / ex: -o result.json")

	rootCmd.Flags().StringToString("regexes", map[string]string{}, "regexes to match in each page / ex: --regexes comments=\"<--.*?-->\"")

	rootCmd.Flags().IntSliceP("exclude-code", "x", []int{}, "status codes to exclude / ex : -x 404,500")
}