/*******************************************************************************
 * Copyright 2020 VMWare.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/
package watcher

import (
	"github.com/edgexfoundry/edgex-cli/cmd/watcher/add"
	"github.com/edgexfoundry/edgex-cli/cmd/watcher/list"
	"github.com/edgexfoundry/edgex-cli/cmd/watcher/rm"

	"github.com/spf13/cobra"
)

// NewCommand returns the watcher command of type cobra.Command
func NewCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "watcher",
		Short: "Watcher command",
		Long:  `Actions related to watchers.`,
	}
	cmd.AddCommand(list.NewCommand())
	cmd.AddCommand(rm.NewCommand())
	cmd.AddCommand(add.NewCommand())
	return cmd
}
