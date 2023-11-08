// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SensitiveType A sensitive type defines a particular type or class of sensitive data. It can be a basic sensitive type with regular
// expressions or a sensitive category. While sensitive types are used for data discovery, sensitive categories are used
// for logically grouping the related or similar sensitive types. Learn more (https://docs.oracle.com/en/cloud/paas/data-safe/udscs/sensitive-types.html#GUID-45A5A3CB-5B67-4C75-9ACC-DD511D14E7C4).
type SensitiveType interface {

	// The OCID of the sensitive type.
	GetId() *string

	// The display name of the sensitive type.
	GetDisplayName() *string

	// The OCID of the compartment that contains the sensitive type.
	GetCompartmentId() *string

	// The current state of the sensitive type.
	GetLifecycleState() DiscoveryLifecycleStateEnum

	// Specifies whether the sensitive type is user-defined or predefined.
	GetSource() SensitiveTypeSourceEnum

	// The date and time the sensitive type was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	GetTimeCreated() *common.SDKTime

	// The date and time the sensitive type was last updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	GetTimeUpdated() *common.SDKTime

	// The short name of the sensitive type.
	GetShortName() *string

	// The description of the sensitive type.
	GetDescription() *string

	// The OCID of the parent sensitive category.
	GetParentCategoryId() *string

	// Specifies whether the sensitive type is common. Common sensitive types belong to
	// library sensitive types which are frequently used to perform sensitive data discovery.
	GetIsCommon() *bool

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type sensitivetype struct {
	JsonData         []byte
	ShortName        *string                           `mandatory:"false" json:"shortName"`
	Description      *string                           `mandatory:"false" json:"description"`
	ParentCategoryId *string                           `mandatory:"false" json:"parentCategoryId"`
	IsCommon         *bool                             `mandatory:"false" json:"isCommon"`
	FreeformTags     map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags       map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Id               *string                           `mandatory:"true" json:"id"`
	DisplayName      *string                           `mandatory:"true" json:"displayName"`
	CompartmentId    *string                           `mandatory:"true" json:"compartmentId"`
	LifecycleState   DiscoveryLifecycleStateEnum       `mandatory:"true" json:"lifecycleState"`
	Source           SensitiveTypeSourceEnum           `mandatory:"true" json:"source"`
	TimeCreated      *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	TimeUpdated      *common.SDKTime                   `mandatory:"true" json:"timeUpdated"`
	EntityType       string                            `json:"entityType"`
}

// UnmarshalJSON unmarshals json
func (m *sensitivetype) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersensitivetype sensitivetype
	s := struct {
		Model Unmarshalersensitivetype
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.LifecycleState = s.Model.LifecycleState
	m.Source = s.Model.Source
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.ShortName = s.Model.ShortName
	m.Description = s.Model.Description
	m.ParentCategoryId = s.Model.ParentCategoryId
	m.IsCommon = s.Model.IsCommon
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.EntityType = s.Model.EntityType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *sensitivetype) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntityType {
	case "SENSITIVE_TYPE":
		mm := SensitiveTypePattern{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SENSITIVE_CATEGORY":
		mm := SensitiveCategory{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for SensitiveType: %s.", m.EntityType)
		return *m, nil
	}
}

// GetShortName returns ShortName
func (m sensitivetype) GetShortName() *string {
	return m.ShortName
}

// GetDescription returns Description
func (m sensitivetype) GetDescription() *string {
	return m.Description
}

// GetParentCategoryId returns ParentCategoryId
func (m sensitivetype) GetParentCategoryId() *string {
	return m.ParentCategoryId
}

// GetIsCommon returns IsCommon
func (m sensitivetype) GetIsCommon() *bool {
	return m.IsCommon
}

// GetFreeformTags returns FreeformTags
func (m sensitivetype) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m sensitivetype) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m sensitivetype) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m sensitivetype) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m sensitivetype) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m sensitivetype) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetLifecycleState returns LifecycleState
func (m sensitivetype) GetLifecycleState() DiscoveryLifecycleStateEnum {
	return m.LifecycleState
}

// GetSource returns Source
func (m sensitivetype) GetSource() SensitiveTypeSourceEnum {
	return m.Source
}

// GetTimeCreated returns TimeCreated
func (m sensitivetype) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m sensitivetype) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m sensitivetype) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m sensitivetype) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDiscoveryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDiscoveryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSensitiveTypeSourceEnum(string(m.Source)); !ok && m.Source != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Source: %s. Supported values are: %s.", m.Source, strings.Join(GetSensitiveTypeSourceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
