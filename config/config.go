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
package config

var (
	DefaultConfig = Config{
		Server:         DefaultServerConfig,
		CommandConfigs: DefaultCommandConfigs,
	}

	DefaultServerConfig = ServerConfig{
		Port: "8080",
	}

	DefaultCommandConfigs = []*CommandConfig{
		{},
	}
)

var (
	Cfg = DefaultConfig
)

type Config struct {
	Server         ServerConfig     `mapstructure:"server,omitempty" json:"server,omitempty" yaml:"server,omitempty"`
	CommandConfigs []*CommandConfig `mapstructure:"command_configs,omitempty" json:"command_configs,omitempty" yaml:"command_configs,omitempty"`
}

type ServerConfig struct {
	Port string `mapstructure:"port,omitempty" json:"port,omitempty" yaml:"port,omitempty"`
}

type CommandConfig struct {
	Command string `mapstructure:"command,omitempty" json:"command,omitempty" yaml:"command,omitempty"`
	Route   string `mapstructure:"route,omitempty" json:"route,omitempty" yaml:"route,omitempty"`
}

var ConfigYamlPath = "./yhttpcmd.yaml"
