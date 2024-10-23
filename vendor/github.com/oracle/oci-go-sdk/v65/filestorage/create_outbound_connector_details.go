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

// CreateOutboundConnectorDetails Details for creating the outbound connector.
type CreateOutboundConnectorDetails interface {

	// The availability domain the outbound connector is in. May be unset
	// as a blank or NULL value.
	// Example: `Uocm:PHX-AD-1`
	GetAvailabilityDomain() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the outbound connector.
	GetCompartmentId() *string

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My outbound connector`
	GetDisplayName() *string

	// Free-form tags for this resource. Each tag is a simple key-value pair
	//  with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Locks associated with this resource.
	GetLocks() []ResourceLock
}

type createoutboundconnectordetails struct {
	JsonData           []byte
	DisplayName        *string                           `mandatory:"false" json:"displayName"`
	FreeformTags       map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags        map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	Locks              []ResourceLock                    `mandatory:"false" json:"locks"`
	AvailabilityDomain *string                           `mandatory:"true" json:"availabilityDomain"`
	CompartmentId      *string                           `mandatory:"true" json:"compartmentId"`
	ConnectorType      string                            `json:"connectorType"`
}

// UnmarshalJSON unmarshals json
func (m *createoutboundconnectordetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateoutboundconnectordetails createoutboundconnectordetails
	s := struct {
		Model Unmarshalercreateoutboundconnectordetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.AvailabilityDomain = s.Model.AvailabilityDomain
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.Locks = s.Model.Locks
	m.ConnectorType = s.Model.ConnectorType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createoutboundconnectordetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectorType {
	case "LDAPBIND":
		mm := CreateLdapBindAccountDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateOutboundConnectorDetails: %s.", m.ConnectorType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m createoutboundconnectordetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetFreeformTags returns FreeformTags
func (m createoutboundconnectordetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createoutboundconnectordetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetLocks returns Locks
func (m createoutboundconnectordetails) GetLocks() []ResourceLock {
	return m.Locks
}

// GetAvailabilityDomain returns AvailabilityDomain
func (m createoutboundconnectordetails) GetAvailabilityDomain() *string {
	return m.AvailabilityDomain
}

// GetCompartmentId returns CompartmentId
func (m createoutboundconnectordetails) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m createoutboundconnectordetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createoutboundconnectordetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
