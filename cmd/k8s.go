package cmd

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/luoruofeng/naval-cli/util"
	"github.com/spf13/cobra"
)

func exec(input, output string, id string, waitSecond int) {
	if r, err := util.IsPathExists(output); err != nil {
		fmt.Println(err)
		return
	} else if !r {
		fmt.Println("输出的文件路径不存在")
		return
	}

	if r, err := util.IsPathExists(input); err != nil {
		fmt.Println(err)
		return
	} else if !r {
		fmt.Println("输入的文件路径不存在")
		return
	} else {
		if r := util.IsDirectory(input); !r {
			fmt.Printf("输入的文件:%s\n", input)
			if util.IsDirectory(output) {
				output = filepath.Join(output, util.CreateOutputFileName(util.GetFileNameWithoutExtension(input)))
			}
			util.CreateNavalYAMLFile(output, id, util.Create, waitSecond, util.ReadFile(input))
		} else if r {
			fmt.Printf("输入的文件夹:%s\n", input)
			err := util.ProcessYAMLFiles(input, func(inputPath string, data []byte) error {
				var outputPath string = output
				if util.IsDirectory(output) {
					outputPath = filepath.Join(output, util.CreateOutputFileName(util.GetFileNameWithoutExtension(inputPath)))
				}
				return util.CreateNavalYAMLFile(outputPath, id, util.Create, waitSecond, string(data))
			})
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

// k8sCmd represents the k8s command
var k8sCmd = &cobra.Command{
	Use:   "k8s",
	Short: "创建create类型的naval执行文件",
	Long: `通过该命令可以通过k8s资源yaml文件创建create类型的naval执行文件，例如通过如下地址的k8s资源文件来生成：
	
	k8s官方example：
	https://github.com/kubernetes/examples

	google的microservices-demo：
	https://github.com/GoogleCloudPlatform/microservices-demo/blob/main/release/kubernetes-manifests.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		oflag := cmd.Flags().ShorthandLookup("o")
		iflag := cmd.Flags().ShorthandLookup("i")
		idflag := cmd.Flags().Lookup("id")
		waitflag := cmd.Flags().Lookup("wait")

		input, err := filepath.Abs(iflag.Value.String())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("输入的文件路径:%s\n", input)

		output, err := filepath.Abs(oflag.Value.String())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("输出的文件路径:%s\n", output)

		id := idflag.Value.String()
		var wait int

		if id == "" {
			fmt.Println("id不能为空")
			return
		}

		w, err := strconv.Atoi(waitflag.Value.String())
		if err != nil {
			fmt.Println(err)
			return
		}
		wait = w

		exec(input, output, id, wait)
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
	lp, _ := filepath.Abs("./")
	k8sCmd.Flags().StringP("output", "o", lp, "输出的文件路径")
	k8sCmd.Flags().StringP("input", "i", lp, "输入的文件路径")
	k8sCmd.Flags().String("id", "", "资源的id")
	k8sCmd.Flags().IntP("wait", "w", 0, "等待的时间")
}
