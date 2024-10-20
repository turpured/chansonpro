// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package list

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	models "github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/spf13/cobra"
)

type deviceServiceList struct {
	rd []models.DeviceService
}

// NewCommand returns the list device services command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists existing devices services",
		Long:  `Return the list fo current device services.`,
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := http.Get("http://localhost:48081/api/v1/deviceservice")
			if err != nil {
				fmt.Println("An error occured. Is EdgeX running?")
				fmt.Println(err)
			}
			defer resp.Body.Close()

			data, _ := ioutil.ReadAll(resp.Body)

			deviceServiceList1 := deviceServiceList{}

			errjson := json.Unmarshal(data, &deviceServiceList1.rd)
			if errjson != nil {
				fmt.Println(errjson)
			}
			for i, device := range deviceServiceList1.rd {
				fmt.Printf("%v\t%s\t%v\n", i, device.Name, device.Created)
			}
		},
	}
	return cmd
}
