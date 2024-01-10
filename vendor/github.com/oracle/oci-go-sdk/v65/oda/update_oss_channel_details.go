// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateOssChannelDetails Properties to update an Oracle Streaming Service (OSS) channel.
type UpdateOssChannelDetails struct {

	// The Channel's name. The name can contain only letters, numbers, periods, and underscores. The name must begin with a letter.
	Name *string `mandatory:"false" json:"name"`

	// A short description of the Channel.
	Description *string `mandatory:"false" json:"description"`

	// The number of milliseconds before a session expires.
	SessionExpiryDurationInMilliseconds *int64 `mandatory:"false" json:"sessionExpiryDurationInMilliseconds"`

	// Simple key-value pair that is applied without any predefined name, type, or scope.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The IDs of the Skills and Digital Assistants that the Channel is routed to.
	EventSinkBotIds []string `mandatory:"false" json:"eventSinkBotIds"`

	// The topic inbound messages are received on.
	InboundMessageTopic *string `mandatory:"false" json:"inboundMessageTopic"`

	// The topic outbound messages are sent on.
	OutboundMessageTopic *string `mandatory:"false" json:"outboundMessageTopic"`

	// The Oracle Streaming Service bootstrap servers.
	BootstrapServers *string `mandatory:"false" json:"bootstrapServers"`

	// The security protocol to use when conecting to the Oracle Streaming Service. See Oracle Streaming Service documentation for a list of valid values.
	SecurityProtocol *string `mandatory:"false" json:"securityProtocol"`

	// The SASL mechanmism to use when conecting to the Oracle Streaming Service. See Oracle Streaming Service documentation for a list of valid values.
	SaslMechanism *string `mandatory:"false" json:"saslMechanism"`

	// The tenancy to use when connecting to the Oracle Streaming Service.
	TenancyName *string `mandatory:"false" json:"tenancyName"`

	// The user name to use when connecting to the Oracle Streaming Service.
	UserName *string `mandatory:"false" json:"userName"`

	// The stream pool OCI to use when connecting to the Oracle Streaming Service.
	StreamPoolId *string `mandatory:"false" json:"streamPoolId"`

	// The authentication token to use when connecting to the Oracle Streaming Service.
	AuthToken *string `mandatory:"false" json:"authToken"`
}

// GetName returns Name
func (m UpdateOssChannelDetails) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m UpdateOssChannelDetails) GetDescription() *string {
	return m.Description
}

// GetSessionExpiryDurationInMilliseconds returns SessionExpiryDurationInMilliseconds
func (m UpdateOssChannelDetails) GetSessionExpiryDurationInMilliseconds() *int64 {
	return m.SessionExpiryDurationInMilliseconds
}

// GetFreeformTags returns FreeformTags
func (m UpdateOssChannelDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateOssChannelDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateOssChannelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOssChannelDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateOssChannelDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateOssChannelDetails UpdateOssChannelDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdateOssChannelDetails
	}{
		"OSS",
		(MarshalTypeUpdateOssChannelDetails)(m),
	}

	return json.Marshal(&s)
}
