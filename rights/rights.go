// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package rights

// Right is the type of rights on The Things Network
type Right string

const (
	// AppSettings is the right to read and write access to the settings, devices and access keys of the application
	AppSettings Right = "settings"

	// AppCollaborators is the right to edit and modify collaborators of the application
	AppCollaborators Right = "collaborators"

	// AppDelete is the right to delete the application
	AppDelete Right = "delete"

	// ReadUplink is the right to view messages sent by devices of the application
	ReadUplink Right = "messages:up:r"

	// WriteUplink is the right to send messages to the application
	WriteUplink Right = "messages:up:w"

	// WriteDownlink is the right to send messages to devices of the application
	WriteDownlink Right = "messages:down:w"

	// Devices is the right to list, edit and remove devices for the application on a handler
	Devices Right = "devices"

	// GatewaySettings is the right to read and write access to the gateway settings
	GatewaySettings Right = "gateway:settings"

	// GatewayCollaborators is the right to edit the gateway collaborators
	GatewayCollaborators Right = "gateway:collaborators"

	// GatewayDelete is the right to delete a gatweay
	GatewayDelete Right = "gateway:delete"

	// GatewayLocation is the right to view the exact location of the gateway, otherwise only approximate location will be shown
	GatewayLocation Right = "gateway:location"

	// GatewayStatus is the right to view the gateway status and metrics about the gateway
	GatewayStatus Right = "gateway:status"

	// GatewayOwner is the right to view the owner of the gateway
	GatewayOwner Right = "gateway:owner"

	// ComponentSettings is the right to read and write access to the settings and access key of a network component
	ComponentSettings Right = "component:settings"

	// ComponentDelete is the right to delete the network component
	ComponentDelete Right = "component:delete"
)
