// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// Use the Organizations API to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and organization resources. For more information, see Organization Management Overview (https://docs.oracle.com/iaas/Content/General/Concepts/organization_management_overview.htm).
//

package tenantmanagercontrolplane

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AssignedSubscription Assigned subscription type, which carries shared properties for any assigned subscription version.
type AssignedSubscription interface {

	// The Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the subscription.
	GetId() *string

	// The Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the owning compartment. Always a tenancy OCID.
	GetCompartmentId() *string

	// The type of subscription, such as 'UCM', 'SAAS', 'ERP', 'CRM'.
	GetServiceName() *string

	// The date and time of creation, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	GetTimeCreated() *common.SDKTime

	// The date and time of update, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	GetTimeUpdated() *common.SDKTime

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type assignedsubscription struct {
	JsonData      []byte
	Id            *string                           `mandatory:"true" json:"id"`
	CompartmentId *string                           `mandatory:"true" json:"compartmentId"`
	ServiceName   *string                           `mandatory:"true" json:"serviceName"`
	TimeCreated   *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	TimeUpdated   *common.SDKTime                   `mandatory:"true" json:"timeUpdated"`
	FreeformTags  map[string]string                 `mandatory:"true" json:"freeformTags"`
	DefinedTags   map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`
	EntityVersion string                            `json:"entityVersion"`
}

// UnmarshalJSON unmarshals json
func (m *assignedsubscription) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerassignedsubscription assignedsubscription
	s := struct {
		Model Unmarshalerassignedsubscription
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.ServiceName = s.Model.ServiceName
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.EntityVersion = s.Model.EntityVersion

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *assignedsubscription) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntityVersion {
	case "V1":
		mm := ClassicAssignedSubscription{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "V2":
		mm := CloudAssignedSubscription{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for AssignedSubscription: %s.", m.EntityVersion)
		return *m, nil
	}
}

// GetId returns Id
func (m assignedsubscription) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m assignedsubscription) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetServiceName returns ServiceName
func (m assignedsubscription) GetServiceName() *string {
	return m.ServiceName
}

// GetTimeCreated returns TimeCreated
func (m assignedsubscription) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m assignedsubscription) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetFreeformTags returns FreeformTags
func (m assignedsubscription) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m assignedsubscription) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m assignedsubscription) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m assignedsubscription) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AssignedSubscriptionEntityVersionEnum Enum with underlying type: string
type AssignedSubscriptionEntityVersionEnum string

// Set of constants representing the allowable values for AssignedSubscriptionEntityVersionEnum
const (
	AssignedSubscriptionEntityVersionV1 AssignedSubscriptionEntityVersionEnum = "V1"
	AssignedSubscriptionEntityVersionV2 AssignedSubscriptionEntityVersionEnum = "V2"
)

var mappingAssignedSubscriptionEntityVersionEnum = map[string]AssignedSubscriptionEntityVersionEnum{
	"V1": AssignedSubscriptionEntityVersionV1,
	"V2": AssignedSubscriptionEntityVersionV2,
}

var mappingAssignedSubscriptionEntityVersionEnumLowerCase = map[string]AssignedSubscriptionEntityVersionEnum{
	"v1": AssignedSubscriptionEntityVersionV1,
	"v2": AssignedSubscriptionEntityVersionV2,
}

// GetAssignedSubscriptionEntityVersionEnumValues Enumerates the set of values for AssignedSubscriptionEntityVersionEnum
func GetAssignedSubscriptionEntityVersionEnumValues() []AssignedSubscriptionEntityVersionEnum {
	values := make([]AssignedSubscriptionEntityVersionEnum, 0)
	for _, v := range mappingAssignedSubscriptionEntityVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetAssignedSubscriptionEntityVersionEnumStringValues Enumerates the set of values in String for AssignedSubscriptionEntityVersionEnum
func GetAssignedSubscriptionEntityVersionEnumStringValues() []string {
	return []string{
		"V1",
		"V2",
	}
}

// GetMappingAssignedSubscriptionEntityVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssignedSubscriptionEntityVersionEnum(val string) (AssignedSubscriptionEntityVersionEnum, bool) {
	enum, ok := mappingAssignedSubscriptionEntityVersionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
