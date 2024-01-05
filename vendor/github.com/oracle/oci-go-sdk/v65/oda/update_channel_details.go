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

// UpdateChannelDetails Properties to update a Channel.
type UpdateChannelDetails interface {

	// The Channel's name. The name can contain only letters, numbers, periods, and underscores. The name must begin with a letter.
	GetName() *string

	// A short description of the Channel.
	GetDescription() *string

	// The number of milliseconds before a session expires.
	GetSessionExpiryDurationInMilliseconds() *int64

	// Simple key-value pair that is applied without any predefined name, type, or scope.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type updatechanneldetails struct {
	JsonData                            []byte
	Name                                *string                           `mandatory:"false" json:"name"`
	Description                         *string                           `mandatory:"false" json:"description"`
	SessionExpiryDurationInMilliseconds *int64                            `mandatory:"false" json:"sessionExpiryDurationInMilliseconds"`
	FreeformTags                        map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags                         map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	Type                                string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *updatechanneldetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatechanneldetails updatechanneldetails
	s := struct {
		Model Unmarshalerupdatechanneldetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.Description = s.Model.Description
	m.SessionExpiryDurationInMilliseconds = s.Model.SessionExpiryDurationInMilliseconds
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatechanneldetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "OSVC":
		mm := UpdateOsvcChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OSS":
		mm := UpdateOssChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ANDROID":
		mm := UpdateAndroidChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MSTEAMS":
		mm := UpdateMsTeamsChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "APPEVENT":
		mm := UpdateAppEventChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "WEB":
		mm := UpdateWebChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IOS":
		mm := UpdateIosChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SLACK":
		mm := UpdateSlackChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SERVICECLOUD":
		mm := UpdateServiceCloudChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TWILIO":
		mm := UpdateTwilioChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "WEBHOOK":
		mm := UpdateWebhookChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "APPLICATION":
		mm := UpdateApplicationChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FACEBOOK":
		mm := UpdateFacebookChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CORTANA":
		mm := UpdateCortanaChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UpdateChannelDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetName returns Name
func (m updatechanneldetails) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m updatechanneldetails) GetDescription() *string {
	return m.Description
}

// GetSessionExpiryDurationInMilliseconds returns SessionExpiryDurationInMilliseconds
func (m updatechanneldetails) GetSessionExpiryDurationInMilliseconds() *int64 {
	return m.SessionExpiryDurationInMilliseconds
}

// GetFreeformTags returns FreeformTags
func (m updatechanneldetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m updatechanneldetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m updatechanneldetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatechanneldetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
