// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstanceConfiguration An instance configuration is a template that defines the settings to use when creating Compute instances.
// An instance configuration is a template that defines the settings to use when creating Compute instances
// or GPU Memory Clusters.
// For more information about instance configurations, see
// Managing Compute Instances (https://docs.oracle.com/iaas/Content/Compute/Concepts/instancemanagement.htm).
type InstanceConfiguration struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment
	// containing the instance configuration.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance configuration.
	Id *string `mandatory:"true" json:"id"`

	// The date and time the instance configuration was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	InstanceDetails InstanceConfigurationInstanceDetails `mandatory:"false" json:"instanceDetails"`

	// The GPU Memory Cluster configuration entries for.
	GmcConfigs []InstanceConfigurationGmcConfigDetail `mandatory:"false" json:"gmcConfigs"`

	// Differentiator for instance configuration.
	// Following values are supported:
	// * INSTANCE : All details related to instance will be passed within instanceDetails.
	// * GMC : All details related to gpu memory cluster will be passed within gmcConfigs.
	Source InstanceConfigurationSourceEnum `mandatory:"false" json:"source,omitempty"`

	// Parameters that were not specified when the instance configuration was created, but that
	// are required to launch an instance from the instance configuration. See the
	// LaunchInstanceConfiguration operation.
	DeferredFields []string `mandatory:"false" json:"deferredFields"`
}

func (m InstanceConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingInstanceConfigurationSourceEnum(string(m.Source)); !ok && m.Source != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Source: %s. Supported values are: %s.", m.Source, strings.Join(GetInstanceConfigurationSourceEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *InstanceConfiguration) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DefinedTags     map[string]map[string]interface{}      `json:"definedTags"`
		DisplayName     *string                                `json:"displayName"`
		FreeformTags    map[string]string                      `json:"freeformTags"`
		InstanceDetails instanceconfigurationinstancedetails   `json:"instanceDetails"`
		GmcConfigs      []InstanceConfigurationGmcConfigDetail `json:"gmcConfigs"`
		Source          InstanceConfigurationSourceEnum        `json:"source"`
		DeferredFields  []string                               `json:"deferredFields"`
		CompartmentId   *string                                `json:"compartmentId"`
		Id              *string                                `json:"id"`
		TimeCreated     *common.SDKTime                        `json:"timeCreated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	nn, e = model.InstanceDetails.UnmarshalPolymorphicJSON(model.InstanceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.InstanceDetails = nn.(InstanceConfigurationInstanceDetails)
	} else {
		m.InstanceDetails = nil
	}

	m.GmcConfigs = make([]InstanceConfigurationGmcConfigDetail, len(model.GmcConfigs))
	copy(m.GmcConfigs, model.GmcConfigs)
	m.Source = model.Source

	m.DeferredFields = make([]string, len(model.DeferredFields))
	copy(m.DeferredFields, model.DeferredFields)
	m.CompartmentId = model.CompartmentId

	m.Id = model.Id

	m.TimeCreated = model.TimeCreated

	return
}

// InstanceConfigurationSourceEnum Enum with underlying type: string
type InstanceConfigurationSourceEnum string

// Set of constants representing the allowable values for InstanceConfigurationSourceEnum
const (
	InstanceConfigurationSourceInstance InstanceConfigurationSourceEnum = "INSTANCE"
	InstanceConfigurationSourceGmc      InstanceConfigurationSourceEnum = "GMC"
)

var mappingInstanceConfigurationSourceEnum = map[string]InstanceConfigurationSourceEnum{
	"INSTANCE": InstanceConfigurationSourceInstance,
	"GMC":      InstanceConfigurationSourceGmc,
}

var mappingInstanceConfigurationSourceEnumLowerCase = map[string]InstanceConfigurationSourceEnum{
	"instance": InstanceConfigurationSourceInstance,
	"gmc":      InstanceConfigurationSourceGmc,
}

// GetInstanceConfigurationSourceEnumValues Enumerates the set of values for InstanceConfigurationSourceEnum
func GetInstanceConfigurationSourceEnumValues() []InstanceConfigurationSourceEnum {
	values := make([]InstanceConfigurationSourceEnum, 0)
	for _, v := range mappingInstanceConfigurationSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetInstanceConfigurationSourceEnumStringValues Enumerates the set of values in String for InstanceConfigurationSourceEnum
func GetInstanceConfigurationSourceEnumStringValues() []string {
	return []string{
		"INSTANCE",
		"GMC",
	}
}

// GetMappingInstanceConfigurationSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstanceConfigurationSourceEnum(val string) (InstanceConfigurationSourceEnum, bool) {
	enum, ok := mappingInstanceConfigurationSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
