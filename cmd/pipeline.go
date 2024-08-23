// Package cmd /*
/*
Copyright © 2024 IllTamer <mail@illtamer.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"github.com/IUnlimit/sample-sls/internal/conf"
	"github.com/IUnlimit/sample-sls/internal/model"
	"github.com/spf13/cobra"
)

// pipelineCmd represents the pipeline command
var pipelineCmd = &cobra.Command{
	Use:   "pipeline",
	Short: "使用管道模式",
	Long: `使用带管道模式启动服务.
	直接将输入和输出读取到当前环境，以便对兼容数据进行管道处理,
	在开始使用 'pipeline' 命令之前，您应该先配置生成的 'config.yml' 文件.`,
	Run: func(cmd *cobra.Command, args []string) {
		conf.Init()

		pipeApps := make([]*model.App, 0)
		for _, app := range conf.Config.App {
			if !app.Pipeline {
				continue
			}
			pipeApps = append(pipeApps, app)
		}

		//scanner := bufio.NewScanner(os.Stdin)
		//for scanner.Scan() {
		//	// 逐行读取标准输入
		//	fmt.Println(scanner.Text())
		//}
		//if err := scanner.Err(); err != nil {
		//	fmt.Println("Error reading stdin:", err)
		//}
	},
}

func init() {
	rootCmd.AddCommand(pipelineCmd)
}
