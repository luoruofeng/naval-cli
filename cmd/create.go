package cmd

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/luoruofeng/naval-cli/util"
	"github.com/spf13/cobra"
)

func exec(input, output string, id string, waitSecond int, typeId util.TT) {
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
			util.CreateNavalYAMLFile(output, id, util.TT(typeId), waitSecond, util.ReadFile(input))
		} else if r {
			fmt.Printf("输入的文件夹:%s\n", input)
			err := util.ProcessYAMLFiles(input, func(inputPath string, data []byte) error {
				var outputPath string = output
				if util.IsDirectory(output) {
					outputPath = filepath.Join(output, util.CreateOutputFileName(util.GetFileNameWithoutExtension(inputPath)))
				}
				return util.CreateNavalYAMLFile(outputPath, id, util.TT(typeId), waitSecond, string(data))
			})
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		oflag := cmd.PersistentFlags().ShorthandLookup("o")
		iflag := cmd.PersistentFlags().ShorthandLookup("i")
		idflag := cmd.PersistentFlags().Lookup("id")
		waitflag := cmd.PersistentFlags().Lookup("wait")
		typeId, err := cmd.PersistentFlags().GetInt("type")
		if err != nil {
			fmt.Println(err)
			return
		}

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

		exec(input, output, id, wait, util.TT(typeId))
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	lp, _ := filepath.Abs("./")
	createCmd.PersistentFlags().StringP("output", "o", lp, "输出的文件路径")
	createCmd.PersistentFlags().StringP("input", "i", lp, "输入的文件路径")
	createCmd.PersistentFlags().String("id", "", "资源的id")
	createCmd.PersistentFlags().IntP("wait", "w", 0, "等待的时间")
}
