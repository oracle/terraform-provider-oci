// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateAssetSourceDetails Asset source update request.
type UpdateAssetSourceDetails interface {

	// A user-friendly name for the asset source. Does not have to be unique, and it's mutable.
	// Avoid entering confidential information.
	GetDisplayName() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that is going to be used to create assets.
	GetAssetsCompartmentId() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the discovery schedule that is going to be assigned to an asset source.
	GetDiscoveryScheduleId() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. It exists only for cross-compatibility.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}

	// Specifies if this is the Source or Destination point for migration - different assets may be discovered depending on setting.
	GetEnvironmentType() EnvironmentTypeEnum
}

type updateassetsourcedetails struct {
	JsonData            []byte
	DisplayName         *string                           `mandatory:"false" json:"displayName"`
	AssetsCompartmentId *string                           `mandatory:"false" json:"assetsCompartmentId"`
	DiscoveryScheduleId *string                           `mandatory:"false" json:"discoveryScheduleId"`
	FreeformTags        map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags         map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags          map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	EnvironmentType     EnvironmentTypeEnum               `mandatory:"false" json:"environmentType,omitempty"`
	Type                string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *updateassetsourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateassetsourcedetails updateassetsourcedetails
	s := struct {
		Model Unmarshalerupdateassetsourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.AssetsCompartmentId = s.Model.AssetsCompartmentId
	m.DiscoveryScheduleId = s.Model.DiscoveryScheduleId
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.EnvironmentType = s.Model.EnvironmentType
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateassetsourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "VMWARE":
		mm := UpdateVmWareAssetSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OLVM":
		mm := UpdateOlvmAssetSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AWS":
		mm := UpdateAwsAssetSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateAssetSourceDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m updateassetsourcedetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetAssetsCompartmentId returns AssetsCompartmentId
func (m updateassetsourcedetails) GetAssetsCompartmentId() *string {
	return m.AssetsCompartmentId
}

// GetDiscoveryScheduleId returns DiscoveryScheduleId
func (m updateassetsourcedetails) GetDiscoveryScheduleId() *string {
	return m.DiscoveryScheduleId
}

// GetFreeformTags returns FreeformTags
func (m updateassetsourcedetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m updateassetsourcedetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m updateassetsourcedetails) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetEnvironmentType returns EnvironmentType
func (m updateassetsourcedetails) GetEnvironmentType() EnvironmentTypeEnum {
	return m.EnvironmentType
}

func (m updateassetsourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateassetsourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingEnvironmentTypeEnum(string(m.EnvironmentType)); !ok && m.EnvironmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EnvironmentType: %s. Supported values are: %s.", m.EnvironmentType, strings.Join(GetEnvironmentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
