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
	"github.com/edgexfoundry/edgex-cli/config"
	"github.com/edgexfoundry/edgex-cli/pkg/formatters"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/metadata"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/urlclient/local"
	"github.com/edgexfoundry/go-mod-core-contracts/models"

	"github.com/spf13/cobra"
)

const deviceTempl = "Device ID\tDevice Name\tDescription\tOperating State\tAdmin State\tDevice Service\tDevice Profile\n" +
	"{{range .}}" +
	"{{.Id}}\t{{.Name}}\t{{.Description}}\t{{.OperatingState}}\t{{.AdminState}}\t{{.Service.Name}}\t{{.Profile.Name}}\n" +
	"{{end}}"

var name string

// NewCommand returns the list device command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "A list of all devices",
		Long:  `Return all devices sorted by id.`,
		RunE:  listHandler,
	}
	cmd.Flags().StringVarP(&name, "name", "n", "", "Returns Device matching the given name")
	return cmd
}

func listHandler(cmd *cobra.Command, args []string) (err error) {
	url := config.Conf.Clients["Metadata"].Url() + clients.ApiDeviceRoute
	mdc := metadata.NewDeviceClient(
		local.New(url),
	)
	var devices []models.Device
	var device models.Device

	if name != "" {
		device, err = mdc.DeviceForName(cmd.Context(), name)
		if err != nil {
			return
		}
		devices = append(devices, device)
	} else {
		devices, err = mdc.Devices(cmd.Context())
		if err != nil {
			return
		}
	}

	formatter := formatters.NewFormatter(deviceTempl, nil)
	err = formatter.Write(devices)
	return
}
