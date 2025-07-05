// Copyright © 2019 VMware, INC
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
	"os"
	"text/tabwriter"
	"time"

	"github.com/edgexfoundry/edgex-cli/pkg/utils"
	models "github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/spf13/cobra"
)

// var rd []models.Device
type deviceList struct {
	rd []models.Device
}

// NewCommand returns the list device command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "A list of all device services",
		Long:  `Return all device services sorted by id.`,
		Run: func(cmd *cobra.Command, args []string) {

			resp, err := http.Get("http://localhost:48081/api/v1/device")
			if err != nil {
				// handle error
				fmt.Println("An error occurred. Is EdgeX running?")
				fmt.Println(err)
			}
			defer resp.Body.Close()

			data, _ := ioutil.ReadAll(resp.Body)

			deviceList1 := deviceList{}

			errjson := json.Unmarshal(data, &deviceList1.rd)
			if errjson != nil {
				fmt.Println(errjson)
			}

			w := new(tabwriter.Writer)
			w.Init(os.Stdout, 0, 8, 1, '\t', 0)
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t\n", "Device ID", "Device Name", "Created", "Operating State", "Device Service", "Device Profile")
			for _, device := range deviceList1.rd {
				tCreated := time.Unix(device.Created/1000, 0)
				fmt.Fprintf(w, "%s\t%s\t%v\t%v\t%s\t%s\t\n",
					device.Id,
					device.Name,
					utils.HumanDuration(time.Since(tCreated)),
					device.OperatingState,
					device.Service.Name,
					device.Profile.Name,
				)
			}
			w.Flush()
		},
	}
	return cmd
}
