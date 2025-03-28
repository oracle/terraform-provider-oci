// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// CreateCortanaChannelDetails Properties required to create a Cortana channel.
type CreateCortanaChannelDetails struct {

	// The Channel's name. The name can contain only letters, numbers, periods, and underscores. The name must begin with a letter.
	Name *string `mandatory:"true" json:"name"`

	// The Microsoft App ID that you obtained when you created your bot registration in Azure.
	MsaAppId *string `mandatory:"true" json:"msaAppId"`

	// The client secret that you obtained from your bot registration.
	MsaAppPassword *string `mandatory:"true" json:"msaAppPassword"`

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
}

// GetName returns Name
func (m CreateCortanaChannelDetails) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m CreateCortanaChannelDetails) GetDescription() *string {
	return m.Description
}

// GetSessionExpiryDurationInMilliseconds returns SessionExpiryDurationInMilliseconds
func (m CreateCortanaChannelDetails) GetSessionExpiryDurationInMilliseconds() *int64 {
	return m.SessionExpiryDurationInMilliseconds
}

// GetFreeformTags returns FreeformTags
func (m CreateCortanaChannelDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateCortanaChannelDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateCortanaChannelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCortanaChannelDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateCortanaChannelDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateCortanaChannelDetails CreateCortanaChannelDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateCortanaChannelDetails
	}{
		"CORTANA",
		(MarshalTypeCreateCortanaChannelDetails)(m),
	}

	return json.Marshal(&s)
}
