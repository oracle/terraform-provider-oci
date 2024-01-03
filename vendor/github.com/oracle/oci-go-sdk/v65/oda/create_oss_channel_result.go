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

// CreateOssChannelResult The configuration for an Oracle Streaming Service (OSS) channel.
type CreateOssChannelResult struct {

	// Unique immutable identifier that was assigned when the Channel was created.
	Id *string `mandatory:"true" json:"id"`

	// The Channel's name. The name can contain only letters, numbers, periods, and underscores. The name must begin with a letter.
	Name *string `mandatory:"true" json:"name"`

	// When the resource was created. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// When the resource was last updated. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The topic inbound messages are received on.
	InboundMessageTopic *string `mandatory:"true" json:"inboundMessageTopic"`

	// The topic outbound messages are sent on.
	OutboundMessageTopic *string `mandatory:"true" json:"outboundMessageTopic"`

	// The Oracle Streaming Service bootstrap servers.
	BootstrapServers *string `mandatory:"true" json:"bootstrapServers"`

	// The security protocol to use when conecting to the Oracle Streaming Service. See Oracle Streaming Service documentation for a list of valid values.
	SecurityProtocol *string `mandatory:"true" json:"securityProtocol"`

	// The SASL mechanmism to use when conecting to the Oracle Streaming Service. See Oracle Streaming Service documentation for a list of valid values.
	SaslMechanism *string `mandatory:"true" json:"saslMechanism"`

	// The tenancy to use when connecting to the Oracle Streaming Service.
	TenancyName *string `mandatory:"true" json:"tenancyName"`

	// The user name to use when connecting to the Oracle Streaming Service.
	UserName *string `mandatory:"true" json:"userName"`

	// The stream pool OCI to use when connecting to the Oracle Streaming Service.
	StreamPoolId *string `mandatory:"true" json:"streamPoolId"`

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

	// The category of the Channel.
	Category ChannelCategoryEnum `mandatory:"true" json:"category"`

	// The Channel's current state.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m CreateOssChannelResult) GetId() *string {
	return m.Id
}

// GetName returns Name
func (m CreateOssChannelResult) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m CreateOssChannelResult) GetDescription() *string {
	return m.Description
}

// GetCategory returns Category
func (m CreateOssChannelResult) GetCategory() ChannelCategoryEnum {
	return m.Category
}

// GetSessionExpiryDurationInMilliseconds returns SessionExpiryDurationInMilliseconds
func (m CreateOssChannelResult) GetSessionExpiryDurationInMilliseconds() *int64 {
	return m.SessionExpiryDurationInMilliseconds
}

// GetLifecycleState returns LifecycleState
func (m CreateOssChannelResult) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m CreateOssChannelResult) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m CreateOssChannelResult) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetFreeformTags returns FreeformTags
func (m CreateOssChannelResult) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateOssChannelResult) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateOssChannelResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOssChannelResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingChannelCategoryEnum(string(m.Category)); !ok && m.Category != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Category: %s. Supported values are: %s.", m.Category, strings.Join(GetChannelCategoryEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateOssChannelResult) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateOssChannelResult CreateOssChannelResult
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateOssChannelResult
	}{
		"OSS",
		(MarshalTypeCreateOssChannelResult)(m),
	}

	return json.Marshal(&s)
}
