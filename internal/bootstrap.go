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
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/yungkei/yhttpcmd/config"
)

func InitializeConfig(cfgFile string) *viper.Viper {
	if cfgFile != "" {
		config.ConfigYamlPath = cfgFile
	}
	cfg := config.ConfigYamlPath

	v := viper.New()
	v.SetConfigFile(cfg)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("[yhttpcmd] [WARNNING] Read config failed: %s.\n", err)
		fmt.Println("[yhttpcmd] Use default config.")
	}

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		if err := v.Unmarshal(&config.Cfg); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&config.Cfg); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("[yhttpcmd] Loaded config(%s).\n", cfg)
	return v
}
