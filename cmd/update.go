package cmd

import (
	"net/http"
	"path/filepath"

	"github.com/luoruofeng/naval-cli/util"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "naval更新任务",
	Long:  `根据给定的navel文件，向naval平台更新任务。`,
	Run: func(cmd *cobra.Command, args []string) {
		input, _ := cmd.Flags().GetString("input")
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetInt("port")
		if input == "" {
			cmd.Help()
			return
		}
		util.AddHttpRequest(host, port, input, http.MethodPut, "task")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	lp, _ := filepath.Abs("./")
	updateCmd.Flags().StringP("input", "i", lp, "输入的文件路径")
	updateCmd.Flags().IntP("port", "p", 8080, "naval端口号")
	updateCmd.Flags().StringP("host", "l", "localhost", "naval主机地址")
}
