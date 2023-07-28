package cmd

import (
	"github.com/luoruofeng/naval-cli/util"
	"github.com/spf13/cobra"
)

// k8sCmd represents the k8s command
var k8sCmd = &cobra.Command{
	Use:   "k8s",
	Short: "创建create类型的naval可执行文件",
	Long: `通过k8s资源yaml文件创建create类型的naval可执行文件，例如通过如下地址的k8s资源文件来生成：
	
	k8s官方example：
	https://github.com/kubernetes/examples

	google的microservices-demo：
	https://github.com/GoogleCloudPlatform/microservices-demo/blob/main/release/kubernetes-manifests.yaml`,
	PreRun: func(cmd *cobra.Command, args []string) {
		var typeFlag int
		cmd.Parent().PersistentFlags().IntVar(&typeFlag, "type", int(util.Create), "任务类型")
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Parent().Run(cmd.Parent(), args)
	},
}

func init() {
	createCmd.AddCommand(k8sCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// k8sCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// k8sCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
