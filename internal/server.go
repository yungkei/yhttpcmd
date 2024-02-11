/*
Copyright © 2024 yungkei

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/yungkei/yhttpcmd/config"
)

func setupRouter(v *viper.Viper) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("static/*")
	r.GET("/yhttpcmd", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"version": Version,
		})
	})

	bindCmdRouter(r, v)
	return r
}

func bindCmdRouter(r *gin.Engine, v *viper.Viper) {
	commands := config.Cfg.CommandConfigs
	for _, command := range commands {
		commandName := command.Command
		if commandName == "" {
			return
		}
		route := command.Route
		if route == "" {
			route = "/" + commandName
		}
		r.POST(route, func(c *gin.Context) {
			var cmdParam CmdParam
			err := c.BindJSON(&cmdParam)
			if err == nil {
				httpcmd(c, commandName, cmdParam.Args)
			}
		})
	}
}

func Start(cfgFile string) {
	viper := InitializeConfig(cfgFile)
	r := setupRouter(viper)
	r.Run(":" + config.Cfg.Server.Port)
}
