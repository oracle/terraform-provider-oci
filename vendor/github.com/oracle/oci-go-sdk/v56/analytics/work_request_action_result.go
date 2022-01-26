// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

// WorkRequestActionResultEnum Enum with underlying type: string
type WorkRequestActionResultEnum string

// Set of constants representing the allowable values for WorkRequestActionResultEnum
const (
	WorkRequestActionResultCompartmentChanged          WorkRequestActionResultEnum = "COMPARTMENT_CHANGED"
	WorkRequestActionResultCreated                     WorkRequestActionResultEnum = "CREATED"
	WorkRequestActionResultDeleted                     WorkRequestActionResultEnum = "DELETED"
	WorkRequestActionResultStarted                     WorkRequestActionResultEnum = "STARTED"
	WorkRequestActionResultStopped                     WorkRequestActionResultEnum = "STOPPED"
	WorkRequestActionResultScaled                      WorkRequestActionResultEnum = "SCALED"
	WorkRequestActionResultNetworkEndpointChanged      WorkRequestActionResultEnum = "NETWORK_ENDPOINT_CHANGED"
	WorkRequestActionResultVanityUrlCreated            WorkRequestActionResultEnum = "VANITY_URL_CREATED"
	WorkRequestActionResultVanityUrlUpdated            WorkRequestActionResultEnum = "VANITY_URL_UPDATED"
	WorkRequestActionResultVanityUrlDeleted            WorkRequestActionResultEnum = "VANITY_URL_DELETED"
	WorkRequestActionResultPrivateAccessChannelCreated WorkRequestActionResultEnum = "PRIVATE_ACCESS_CHANNEL_CREATED"
	WorkRequestActionResultPrivateAccessChannelUpdated WorkRequestActionResultEnum = "PRIVATE_ACCESS_CHANNEL_UPDATED"
	WorkRequestActionResultPrivateAccessChannelDeleted WorkRequestActionResultEnum = "PRIVATE_ACCESS_CHANNEL_DELETED"
	WorkRequestActionResultNone                        WorkRequestActionResultEnum = "NONE"
)

var mappingWorkRequestActionResult = map[string]WorkRequestActionResultEnum{
	"COMPARTMENT_CHANGED":            WorkRequestActionResultCompartmentChanged,
	"CREATED":                        WorkRequestActionResultCreated,
	"DELETED":                        WorkRequestActionResultDeleted,
	"STARTED":                        WorkRequestActionResultStarted,
	"STOPPED":                        WorkRequestActionResultStopped,
	"SCALED":                         WorkRequestActionResultScaled,
	"NETWORK_ENDPOINT_CHANGED":       WorkRequestActionResultNetworkEndpointChanged,
	"VANITY_URL_CREATED":             WorkRequestActionResultVanityUrlCreated,
	"VANITY_URL_UPDATED":             WorkRequestActionResultVanityUrlUpdated,
	"VANITY_URL_DELETED":             WorkRequestActionResultVanityUrlDeleted,
	"PRIVATE_ACCESS_CHANNEL_CREATED": WorkRequestActionResultPrivateAccessChannelCreated,
	"PRIVATE_ACCESS_CHANNEL_UPDATED": WorkRequestActionResultPrivateAccessChannelUpdated,
	"PRIVATE_ACCESS_CHANNEL_DELETED": WorkRequestActionResultPrivateAccessChannelDeleted,
	"NONE":                           WorkRequestActionResultNone,
}

// GetWorkRequestActionResultEnumValues Enumerates the set of values for WorkRequestActionResultEnum
func GetWorkRequestActionResultEnumValues() []WorkRequestActionResultEnum {
	values := make([]WorkRequestActionResultEnum, 0)
	for _, v := range mappingWorkRequestActionResult {
		values = append(values, v)
	}
	return values
}
