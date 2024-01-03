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

// CreateOsvcChannelDetails Properties required to create an OSVC channel.
type CreateOsvcChannelDetails struct {

	// The Channel's name. The name can contain only letters, numbers, periods, and underscores. The name must begin with a letter.
	Name *string `mandatory:"true" json:"name"`

	// The host.
	//
	// For OSVC, you can derive these values from the URL that you use to launch the Agent Browser User Interface
	// or the chat launch page. For example, if the URL is https://sitename.exampledomain.com/app/chat/chat_launch,
	// then the host is sitename.exampledomain.com.
	//
	// For FUSION, this is the host portion of your Oracle Applications Cloud (Fusion) instance's URL.
	// For example: sitename.exampledomain.com.
	Host *string `mandatory:"true" json:"host"`

	// The port.
	Port *string `mandatory:"true" json:"port"`

	// The user name for the digital-assistant agent.
	UserName *string `mandatory:"true" json:"userName"`

	// The password for the digital-assistant agent.
	Password *string `mandatory:"true" json:"password"`

	// The total session count.
	TotalSessionCount *int `mandatory:"true" json:"totalSessionCount"`

	// The name of the Authentication Provider to use to authenticate the user.
	AuthenticationProviderName *string `mandatory:"true" json:"authenticationProviderName"`

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

	// The ID of the Skill or Digital Assistant that the Channel is routed to.
	BotId *string `mandatory:"false" json:"botId"`

	// The type of OSVC service.
	ChannelService OsvcServiceTypeEnum `mandatory:"false" json:"channelService,omitempty"`
}

// GetName returns Name
func (m CreateOsvcChannelDetails) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m CreateOsvcChannelDetails) GetDescription() *string {
	return m.Description
}

// GetSessionExpiryDurationInMilliseconds returns SessionExpiryDurationInMilliseconds
func (m CreateOsvcChannelDetails) GetSessionExpiryDurationInMilliseconds() *int64 {
	return m.SessionExpiryDurationInMilliseconds
}

// GetFreeformTags returns FreeformTags
func (m CreateOsvcChannelDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateOsvcChannelDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateOsvcChannelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOsvcChannelDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOsvcServiceTypeEnum(string(m.ChannelService)); !ok && m.ChannelService != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ChannelService: %s. Supported values are: %s.", m.ChannelService, strings.Join(GetOsvcServiceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateOsvcChannelDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateOsvcChannelDetails CreateOsvcChannelDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateOsvcChannelDetails
	}{
		"OSVC",
		(MarshalTypeCreateOsvcChannelDetails)(m),
	}

	return json.Marshal(&s)
}
