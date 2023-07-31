package cmd

import (
	"github.com/luoruofeng/naval-cli/util"
	"github.com/spf13/cobra"
)

// composeCmd represents the compose command
var composeCmd = &cobra.Command{
	Use:   "compose",
	Short: "创建convert类型的naval可执行文件",
	Long: `通过docker-compose.yaml文件创建convert类型的naval可执行文件，例如通过如下地址的docker-compose资源文件来生成：

	awesome-compose:
	https://github.com/docker/awesome-compose

其他参考资料:

	Docker官方的Compose ：
	https://github.com/docker/compose

	Docker官方的实验室：
	https://github.com/docker/labs`,
	PreRun: func(cmd *cobra.Command, args []string) {
		var typeFlag int
		cmd.Parent().PersistentFlags().IntVar(&typeFlag, "type", int(util.Convert), "任务类型")
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Parent().Run(cmd.Parent(), args)
	},
}

func init() {
	createCmd.AddCommand(composeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// composeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// composeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
