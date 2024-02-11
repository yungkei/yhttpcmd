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
	"os/exec"

	"github.com/gin-gonic/gin"
)

func httpcmd(c *gin.Context, command string, args ...string) error {
	cmd := exec.Command(command, args...)

	stdout, err := cmd.Output()
	cmdStr := cmd.String()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Command": cmdStr,
			"Message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Command": cmdStr,
			"Message": string(stdout),
		})
	}

	return err
}

type CmdParam struct {
	Args string `json:"args" form:"args"`
}
