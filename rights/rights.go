// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package rights

const (
	// AppSettings is the right to read and write access to the settings, devices and access keys of the application
	AppSettings = "settings"

	// AppCollaborators is the right to edit and modify collaborators of the application
	AppCollaborators = "collaborators"

	// is the right to delete the application
	AppDelete = "delete"

	// ReadUplink is the right to view messages sent by devices of the application
	ReadUplink = "messages:up:r"

	// WriteUplink is the right to send messages to the application
	WriteUplink = "messages:up:w"

	// WriteDownlink is the right to send messages to devices of the application
	WriteDownlink = "messages:down:w"

	// evices is the right to list, edit and remove devices for the application on a handler
	Devices = "devices"

	// GatewaySettings is the right to read and write access to the gateway settings
	GatewaySettings = "gateway:settings"

	// GatewayCollaborators is the right to edit the gateway collaborators
	GatewayCollaborators = "gateway:collaborators"

	// GatewayDelete is the right to delete a gatweay
	GatewayDelete = "gateway:delete"

	// GatewayLocation is the right to view the exact location of the gateway, otherwise only approximate location will be shown
	GatewayLocation = "gateway:location"

	// GatewayStatus is the right to view the gateway status and metrics about the gateway
	GatewayStatus = "gateway:status"

	// GatewayOwner is the right to view the owner of the gateway
	GatewayOwner = "gateway:owner"

	// ComponentSettings is the right to read and write access to the settings and access key of a network component
	ComponentSettings = "component:settings"

	// ComponentDelete is the right to delete the network component
	ComponentDelete = "component:delete"
)
