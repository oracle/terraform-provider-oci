// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// Use the File Storage service API to manage file systems, mount targets, and snapshots.
// For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OutboundConnector Outbound connectors are used to help File Storage communicate with an external server, such as an LDAP server.
// An outbound connector contains all the information needed to connect, authenticate, and gain authorization to perform the account's required functions.
type OutboundConnector interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the outbound connector.
	GetCompartmentId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the outbound connector.
	GetId() *string

	// The current state of this outbound connector.
	GetLifecycleState() OutboundConnectorLifecycleStateEnum

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My outbound connector`
	GetDisplayName() *string

	// The date and time the outbound connector was created
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	GetTimeCreated() *common.SDKTime

	// The availability domain the outbound connector is in. May be unset
	// as a blank or NULL value.
	// Example: `Uocm:PHX-AD-1`
	GetAvailabilityDomain() *string

	// Free-form tags for this resource. Each tag is a simple key-value pair
	//  with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type outboundconnector struct {
	JsonData           []byte
	AvailabilityDomain *string                             `mandatory:"false" json:"availabilityDomain"`
	FreeformTags       map[string]string                   `mandatory:"false" json:"freeformTags"`
	DefinedTags        map[string]map[string]interface{}   `mandatory:"false" json:"definedTags"`
	CompartmentId      *string                             `mandatory:"true" json:"compartmentId"`
	Id                 *string                             `mandatory:"true" json:"id"`
	LifecycleState     OutboundConnectorLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	DisplayName        *string                             `mandatory:"true" json:"displayName"`
	TimeCreated        *common.SDKTime                     `mandatory:"true" json:"timeCreated"`
	ConnectorType      string                              `json:"connectorType"`
}

// UnmarshalJSON unmarshals json
func (m *outboundconnector) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleroutboundconnector outboundconnector
	s := struct {
		Model Unmarshaleroutboundconnector
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.Id = s.Model.Id
	m.LifecycleState = s.Model.LifecycleState
	m.DisplayName = s.Model.DisplayName
	m.TimeCreated = s.Model.TimeCreated
	m.AvailabilityDomain = s.Model.AvailabilityDomain
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.ConnectorType = s.Model.ConnectorType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *outboundconnector) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectorType {
	case "LDAPBIND":
		mm := LdapBindAccount{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for OutboundConnector: %s.", m.ConnectorType)
		return *m, nil
	}
}

// GetAvailabilityDomain returns AvailabilityDomain
func (m outboundconnector) GetAvailabilityDomain() *string {
	return m.AvailabilityDomain
}

// GetFreeformTags returns FreeformTags
func (m outboundconnector) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m outboundconnector) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetCompartmentId returns CompartmentId
func (m outboundconnector) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetId returns Id
func (m outboundconnector) GetId() *string {
	return m.Id
}

// GetLifecycleState returns LifecycleState
func (m outboundconnector) GetLifecycleState() OutboundConnectorLifecycleStateEnum {
	return m.LifecycleState
}

// GetDisplayName returns DisplayName
func (m outboundconnector) GetDisplayName() *string {
	return m.DisplayName
}

// GetTimeCreated returns TimeCreated
func (m outboundconnector) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

func (m outboundconnector) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m outboundconnector) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOutboundConnectorLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOutboundConnectorLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OutboundConnectorLifecycleStateEnum Enum with underlying type: string
type OutboundConnectorLifecycleStateEnum string

// Set of constants representing the allowable values for OutboundConnectorLifecycleStateEnum
const (
	OutboundConnectorLifecycleStateCreating OutboundConnectorLifecycleStateEnum = "CREATING"
	OutboundConnectorLifecycleStateActive   OutboundConnectorLifecycleStateEnum = "ACTIVE"
	OutboundConnectorLifecycleStateDeleting OutboundConnectorLifecycleStateEnum = "DELETING"
	OutboundConnectorLifecycleStateDeleted  OutboundConnectorLifecycleStateEnum = "DELETED"
)

var mappingOutboundConnectorLifecycleStateEnum = map[string]OutboundConnectorLifecycleStateEnum{
	"CREATING": OutboundConnectorLifecycleStateCreating,
	"ACTIVE":   OutboundConnectorLifecycleStateActive,
	"DELETING": OutboundConnectorLifecycleStateDeleting,
	"DELETED":  OutboundConnectorLifecycleStateDeleted,
}

var mappingOutboundConnectorLifecycleStateEnumLowerCase = map[string]OutboundConnectorLifecycleStateEnum{
	"creating": OutboundConnectorLifecycleStateCreating,
	"active":   OutboundConnectorLifecycleStateActive,
	"deleting": OutboundConnectorLifecycleStateDeleting,
	"deleted":  OutboundConnectorLifecycleStateDeleted,
}

// GetOutboundConnectorLifecycleStateEnumValues Enumerates the set of values for OutboundConnectorLifecycleStateEnum
func GetOutboundConnectorLifecycleStateEnumValues() []OutboundConnectorLifecycleStateEnum {
	values := make([]OutboundConnectorLifecycleStateEnum, 0)
	for _, v := range mappingOutboundConnectorLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOutboundConnectorLifecycleStateEnumStringValues Enumerates the set of values in String for OutboundConnectorLifecycleStateEnum
func GetOutboundConnectorLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingOutboundConnectorLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOutboundConnectorLifecycleStateEnum(val string) (OutboundConnectorLifecycleStateEnum, bool) {
	enum, ok := mappingOutboundConnectorLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OutboundConnectorConnectorTypeEnum Enum with underlying type: string
type OutboundConnectorConnectorTypeEnum string

// Set of constants representing the allowable values for OutboundConnectorConnectorTypeEnum
const (
	OutboundConnectorConnectorTypeLdapbind OutboundConnectorConnectorTypeEnum = "LDAPBIND"
)

var mappingOutboundConnectorConnectorTypeEnum = map[string]OutboundConnectorConnectorTypeEnum{
	"LDAPBIND": OutboundConnectorConnectorTypeLdapbind,
}

var mappingOutboundConnectorConnectorTypeEnumLowerCase = map[string]OutboundConnectorConnectorTypeEnum{
	"ldapbind": OutboundConnectorConnectorTypeLdapbind,
}

// GetOutboundConnectorConnectorTypeEnumValues Enumerates the set of values for OutboundConnectorConnectorTypeEnum
func GetOutboundConnectorConnectorTypeEnumValues() []OutboundConnectorConnectorTypeEnum {
	values := make([]OutboundConnectorConnectorTypeEnum, 0)
	for _, v := range mappingOutboundConnectorConnectorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOutboundConnectorConnectorTypeEnumStringValues Enumerates the set of values in String for OutboundConnectorConnectorTypeEnum
func GetOutboundConnectorConnectorTypeEnumStringValues() []string {
	return []string{
		"LDAPBIND",
	}
}

// GetMappingOutboundConnectorConnectorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOutboundConnectorConnectorTypeEnum(val string) (OutboundConnectorConnectorTypeEnum, bool) {
	enum, ok := mappingOutboundConnectorConnectorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
