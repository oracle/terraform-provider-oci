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

// CreateServiceCloudChannelDetails Properties required to create an Service Cloud channel.
type CreateServiceCloudChannelDetails struct {

	// The Channel's name. The name can contain only letters, numbers, periods, and underscores. The name must begin with a letter.
	Name *string `mandatory:"true" json:"name"`

	// The domain name.
	//
	// If you have access to Oracle B2C Service, you can derive this value from the URL that you use to launch the
	// Agent Browser User Interface. For example, if the URL is sitename.exampledomain.com, then the host name prefix
	// is sitename and the domain name is exampledomain.com.
	//
	// If the channel is connecting to Oracle B2C Service version 19A or later, and you have multiple interfaces,
	// then you must include the interface ID in the host (site) name . For example, for the interface that has an ID of 2, you would use something like sitename-2.exampledomain.com.
	DomainName *string `mandatory:"true" json:"domainName"`

	// The host prefix.
	//
	// If you have access to Oracle B2C Service, you can derive this value from the URL that you use to launch the
	// Agent Browser User Interface. For example, if the URL is sitename.exampledomain.com, then the host name prefix
	// is sitename and the domain name is exampledomain.com.
	//
	// If the channel is connecting to Oracle B2C Service version 19A or later, and you have multiple interfaces,
	// then you must include the interface ID in the host (site) name . For example, for the interface that has an ID of 2, you would use something like sitename-2.exampledomain.com.
	HostNamePrefix *string `mandatory:"true" json:"hostNamePrefix"`

	// The user name for an Oracle B2C Service staff member who has the necessary profile permissions.
	UserName *string `mandatory:"true" json:"userName"`

	// The password for the Oracle B2C Service staff member who has the necessary profile permissions.
	Password *string `mandatory:"true" json:"password"`

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

	// The type of Service Cloud client.
	ClientType ServiceCloudClientTypeEnum `mandatory:"true" json:"clientType"`
}

// GetName returns Name
func (m CreateServiceCloudChannelDetails) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m CreateServiceCloudChannelDetails) GetDescription() *string {
	return m.Description
}

// GetSessionExpiryDurationInMilliseconds returns SessionExpiryDurationInMilliseconds
func (m CreateServiceCloudChannelDetails) GetSessionExpiryDurationInMilliseconds() *int64 {
	return m.SessionExpiryDurationInMilliseconds
}

// GetFreeformTags returns FreeformTags
func (m CreateServiceCloudChannelDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateServiceCloudChannelDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateServiceCloudChannelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateServiceCloudChannelDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingServiceCloudClientTypeEnum(string(m.ClientType)); !ok && m.ClientType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClientType: %s. Supported values are: %s.", m.ClientType, strings.Join(GetServiceCloudClientTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateServiceCloudChannelDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateServiceCloudChannelDetails CreateServiceCloudChannelDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateServiceCloudChannelDetails
	}{
		"SERVICECLOUD",
		(MarshalTypeCreateServiceCloudChannelDetails)(m),
	}

	return json.Marshal(&s)
}
