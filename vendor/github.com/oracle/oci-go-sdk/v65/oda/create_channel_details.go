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

// CreateChannelDetails Properties that are required to create a Channel.
type CreateChannelDetails interface {

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

type createchanneldetails struct {
	JsonData                            []byte
	Description                         *string                           `mandatory:"false" json:"description"`
	SessionExpiryDurationInMilliseconds *int64                            `mandatory:"false" json:"sessionExpiryDurationInMilliseconds"`
	FreeformTags                        map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags                         map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	Name                                *string                           `mandatory:"true" json:"name"`
	Type                                string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createchanneldetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatechanneldetails createchanneldetails
	s := struct {
		Model Unmarshalercreatechanneldetails
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
func (m *createchanneldetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "MSTEAMS":
		mm := CreateMsTeamsChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "WEB":
		mm := CreateWebChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FACEBOOK":
		mm := CreateFacebookChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "APPLICATION":
		mm := CreateApplicationChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SERVICECLOUD":
		mm := CreateServiceCloudChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SLACK":
		mm := CreateSlackChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OSVC":
		mm := CreateOsvcChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "APPEVENT":
		mm := CreateAppEventChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OSS":
		mm := CreateOssChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CORTANA":
		mm := CreateCortanaChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ANDROID":
		mm := CreateAndroidChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TWILIO":
		mm := CreateTwilioChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "WEBHOOK":
		mm := CreateWebhookChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IOS":
		mm := CreateIosChannelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateChannelDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetDescription returns Description
func (m createchanneldetails) GetDescription() *string {
	return m.Description
}

// GetSessionExpiryDurationInMilliseconds returns SessionExpiryDurationInMilliseconds
func (m createchanneldetails) GetSessionExpiryDurationInMilliseconds() *int64 {
	return m.SessionExpiryDurationInMilliseconds
}

// GetFreeformTags returns FreeformTags
func (m createchanneldetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createchanneldetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetName returns Name
func (m createchanneldetails) GetName() *string {
	return m.Name
}

func (m createchanneldetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createchanneldetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
