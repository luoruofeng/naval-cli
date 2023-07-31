package cmd

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/luoruofeng/naval-cli/util"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "naval添加新的任务",
	Long: `根据给定的navel文件，向naval平台添加新的任务。
	
naval文件可以是单个文件，也可以是文件夹。如果是文件夹，naval会遍历文件夹下的所有yml文件，将所有的文件都添加到naval平台上。
关于如何编写naval文件，请参考：
	
	https://github.com/luoruofeng/Naval`,
	Run: func(cmd *cobra.Command, args []string) {
		input, _ := cmd.Flags().GetString("input")
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetInt("port")

		util.HttpRequest(host, port, input, http.MethodPost, "task")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	lp, _ := filepath.Abs("./")
	input := addCmd.Flags().StringP("input", "i", lp, "输入的文件路径")
	host := addCmd.Flags().StringP("host", "l", "localhost", "输入的文件路径")
	port := addCmd.Flags().IntP("port", "p", 8080, "输入的文件路径")
	fmt.Println(input, host, port)
}
