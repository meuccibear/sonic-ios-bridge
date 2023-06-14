/*
 *   sonic-ios-bridge  Connect to your iOS Devices.
 *   Copyright (C) 2022 SonicCloudOrg
 *
 *   This program is free software: you can redistribute it and/or modify
 *   it under the terms of the GNU Affero General Public License as published
 *   by the Free Software Foundation, either version 3 of the License, or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU Affero General Public License for more details.
 *
 *   You should have received a copy of the GNU Affero General Public License
 *   along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */
package remote

import "github.com/spf13/cobra"

var (
	remoteCmd *cobra.Command
	udid      string
	port      int
	host      string
)

const (
	OnLine  = "onLine"
	OffLine = "offLine"
	Unknown = "unknown"
)

func InitRemote(remote *cobra.Command) {
	remoteCmd = remote
	shareInit()
	connectInit()
	disConnectInit()
}
