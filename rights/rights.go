// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package rights

import "github.com/TheThingsNetwork/ttn/core/types"

const (
	// AppSettings is the types.Right to read and write access to the settings, devices and access keys of the application
	AppSettings types.Right = "settings"

	// AppCollaborators is the types.Right to edit and modify collaborators of the application
	AppCollaborators types.Right = "collaborators"

	// AppDelete is the types.Right to delete the application
	AppDelete types.Right = "delete"

	// ReadUplink is the types.Right to view messages sent by devices of the application
	ReadUplink types.Right = "messages:up:r"

	// WriteUplink is the types.Right to send messages to the application
	WriteUplink types.Right = "messages:up:w"

	// WriteDownlink is the types.Right to send messages to devices of the application
	WriteDownlink types.Right = "messages:down:w"

	// Devices is the types.Right to list, edit and remove devices for the application on a handler
	Devices types.Right = "devices"

	// GatewaySettings is the types.Right to read and write access to the gateway settings
	GatewaySettings types.Right = "gateway:settings"

	// GatewayCollaborators is the types.Right to edit the gateway collaborators
	GatewayCollaborators types.Right = "gateway:collaborators"

	// GatewayDelete is the types.Right to delete a gatweay
	GatewayDelete types.Right = "gateway:delete"

	// GatewayLocation is the types.Right to view the exact location of the gateway, otherwise only approximate location will be shown
	GatewayLocation types.Right = "gateway:location"

	// GatewayStatus is the types.Right to view the gateway status and metrics about the gateway
	GatewayStatus types.Right = "gateway:status"

	// GatewayOwner is the types.Right that states that a collaborator is an owner
	GatewayOwner types.Right = "gateway:owner"

	// GatewayMessages is the types.Right to view the messages of a gateway
	GatewayMessages types.Right = "gateway:messages"

	// ComponentSettings is the types.Right to read and write access to the settings and access key of a network component
	ComponentSettings types.Right = "component:settings"

	// ComponentDelete is the types.Right to delete the network component
	ComponentDelete types.Right = "component:delete"

	// ComponentCollaborators is the types.Right to view and edit component collaborators
	ComponentCollaborators = "component:collaborators"
)
