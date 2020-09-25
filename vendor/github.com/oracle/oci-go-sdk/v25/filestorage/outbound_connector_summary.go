// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// API for the File Storage service. Use this API to manage file systems, mount targets, and snapshots. For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v25/common"
)

// OutboundConnectorSummary Summary information for a outbound connector.
type OutboundConnectorSummary interface {

	// The availability domain the outbound connector is in. May be unset
	// as a blank or NULL value.
	// Example: `Uocm:PHX-AD-1`
	GetAvailabilityDomain() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the outbound connector.
	GetCompartmentId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the outbound connector.
	GetId() *string

	// The current state of this outbound connector.
	GetLifecycleState() OutboundConnectorSummaryLifecycleStateEnum

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My outbound connector`
	GetDisplayName() *string

	// The date and time the outbound connector was created
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	GetTimeCreated() *common.SDKTime

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

type outboundconnectorsummary struct {
	JsonData           []byte
	AvailabilityDomain *string                                    `mandatory:"false" json:"availabilityDomain"`
	CompartmentId      *string                                    `mandatory:"false" json:"compartmentId"`
	Id                 *string                                    `mandatory:"false" json:"id"`
	LifecycleState     OutboundConnectorSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
	DisplayName        *string                                    `mandatory:"false" json:"displayName"`
	TimeCreated        *common.SDKTime                            `mandatory:"false" json:"timeCreated"`
	FreeformTags       map[string]string                          `mandatory:"false" json:"freeformTags"`
	DefinedTags        map[string]map[string]interface{}          `mandatory:"false" json:"definedTags"`
	ConnectorType      string                                     `json:"connectorType"`
}

// UnmarshalJSON unmarshals json
func (m *outboundconnectorsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleroutboundconnectorsummary outboundconnectorsummary
	s := struct {
		Model Unmarshaleroutboundconnectorsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.AvailabilityDomain = s.Model.AvailabilityDomain
	m.CompartmentId = s.Model.CompartmentId
	m.Id = s.Model.Id
	m.LifecycleState = s.Model.LifecycleState
	m.DisplayName = s.Model.DisplayName
	m.TimeCreated = s.Model.TimeCreated
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.ConnectorType = s.Model.ConnectorType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *outboundconnectorsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectorType {
	case "LDAPBIND":
		mm := LdapBindAccountSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetAvailabilityDomain returns AvailabilityDomain
func (m outboundconnectorsummary) GetAvailabilityDomain() *string {
	return m.AvailabilityDomain
}

//GetCompartmentId returns CompartmentId
func (m outboundconnectorsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetId returns Id
func (m outboundconnectorsummary) GetId() *string {
	return m.Id
}

//GetLifecycleState returns LifecycleState
func (m outboundconnectorsummary) GetLifecycleState() OutboundConnectorSummaryLifecycleStateEnum {
	return m.LifecycleState
}

//GetDisplayName returns DisplayName
func (m outboundconnectorsummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetTimeCreated returns TimeCreated
func (m outboundconnectorsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetFreeformTags returns FreeformTags
func (m outboundconnectorsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m outboundconnectorsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m outboundconnectorsummary) String() string {
	return common.PointerString(m)
}

// OutboundConnectorSummaryLifecycleStateEnum Enum with underlying type: string
type OutboundConnectorSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for OutboundConnectorSummaryLifecycleStateEnum
const (
	OutboundConnectorSummaryLifecycleStateCreating OutboundConnectorSummaryLifecycleStateEnum = "CREATING"
	OutboundConnectorSummaryLifecycleStateActive   OutboundConnectorSummaryLifecycleStateEnum = "ACTIVE"
	OutboundConnectorSummaryLifecycleStateDeleting OutboundConnectorSummaryLifecycleStateEnum = "DELETING"
	OutboundConnectorSummaryLifecycleStateDeleted  OutboundConnectorSummaryLifecycleStateEnum = "DELETED"
)

var mappingOutboundConnectorSummaryLifecycleState = map[string]OutboundConnectorSummaryLifecycleStateEnum{
	"CREATING": OutboundConnectorSummaryLifecycleStateCreating,
	"ACTIVE":   OutboundConnectorSummaryLifecycleStateActive,
	"DELETING": OutboundConnectorSummaryLifecycleStateDeleting,
	"DELETED":  OutboundConnectorSummaryLifecycleStateDeleted,
}

// GetOutboundConnectorSummaryLifecycleStateEnumValues Enumerates the set of values for OutboundConnectorSummaryLifecycleStateEnum
func GetOutboundConnectorSummaryLifecycleStateEnumValues() []OutboundConnectorSummaryLifecycleStateEnum {
	values := make([]OutboundConnectorSummaryLifecycleStateEnum, 0)
	for _, v := range mappingOutboundConnectorSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
