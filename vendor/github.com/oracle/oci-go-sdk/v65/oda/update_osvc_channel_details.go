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

// UpdateOsvcChannelDetails Properties required to update an OSVC channel.
type UpdateOsvcChannelDetails struct {

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

	// The host.
	//
	// For OSVC, you can derive these values from the URL that you use to launch the Agent Browser User Interface
	// or the chat launch page. For example, if the URL is https://sitename.exampledomain.com/app/chat/chat_launch,
	// then the host is sitename.exampledomain.com.
	//
	// For FUSION, this is the host portion of your Oracle Applications Cloud (Fusion) instance's URL.
	// For example: sitename.exampledomain.com.
	Host *string `mandatory:"false" json:"host"`

	// The port.
	Port *string `mandatory:"false" json:"port"`

	// The user name for the digital-assistant agent.
	UserName *string `mandatory:"false" json:"userName"`

	// The password for the digital-assistant agent.
	Password *string `mandatory:"false" json:"password"`

	// The total session count.
	TotalSessionCount *int `mandatory:"false" json:"totalSessionCount"`

	// The name of the Authentication Provider to use to authenticate the user.
	AuthenticationProviderName *string `mandatory:"false" json:"authenticationProviderName"`

	// The ID of the Skill or Digital Assistant that the Channel is routed to.
	BotId *string `mandatory:"false" json:"botId"`

	// The type of OSVC service.
	ChannelService OsvcServiceTypeEnum `mandatory:"false" json:"channelService,omitempty"`
}

// GetName returns Name
func (m UpdateOsvcChannelDetails) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m UpdateOsvcChannelDetails) GetDescription() *string {
	return m.Description
}

// GetSessionExpiryDurationInMilliseconds returns SessionExpiryDurationInMilliseconds
func (m UpdateOsvcChannelDetails) GetSessionExpiryDurationInMilliseconds() *int64 {
	return m.SessionExpiryDurationInMilliseconds
}

// GetFreeformTags returns FreeformTags
func (m UpdateOsvcChannelDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m UpdateOsvcChannelDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateOsvcChannelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOsvcChannelDetails) ValidateEnumValue() (bool, error) {
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
func (m UpdateOsvcChannelDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateOsvcChannelDetails UpdateOsvcChannelDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdateOsvcChannelDetails
	}{
		"OSVC",
		(MarshalTypeUpdateOsvcChannelDetails)(m),
	}

	return json.Marshal(&s)
}
