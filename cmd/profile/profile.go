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

package profile

import (
	"github.com/edgexfoundry/edgex-cli/cmd/profile/add"
	"github.com/edgexfoundry/edgex-cli/cmd/profile/list"
	"github.com/edgexfoundry/edgex-cli/cmd/profile/rm"
	"github.com/edgexfoundry/edgex-cli/cmd/profile/update"

	"github.com/spf13/cobra"
)

// NewCommand returns the profile command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "profile [option]",
		Short: "Device profile command.",
		Long: `Device profile
		
Device profiles describe the attributes of a common class of devices 
managed by a specific DS. A device profile is made up of one or more 
device resources, optional resources, and optional commands. 
Commands are added to Core Metadata, and are used to validate REST API 
requests made to Core Command's command endpoint. 
`,
	}

	cmd.AddCommand(rm.NewCommand())
	cmd.AddCommand(list.NewCommand())
	cmd.AddCommand(add.NewCommand())
	cmd.AddCommand(update.NewCommand())
	return cmd
}
