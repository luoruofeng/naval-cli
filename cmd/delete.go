package cmd

import (
	"net/http"

	"github.com/luoruofeng/naval-cli/util"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "根据任务id进行删除",
	Long:  `根据任务id进行删除`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetInt("port")
		if id == "" {
			cmd.Help()
			return
		}
		util.DeleteHttpRequest(host, port, id, http.MethodDelete, "task")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	deleteCmd.Flags().IntP("port", "p", 8080, "naval端口号")
	deleteCmd.Flags().StringP("host", "l", "localhost", "naval主机地址")
	deleteCmd.Flags().StringP("id", "i", "", "需要删除的task id")
}
