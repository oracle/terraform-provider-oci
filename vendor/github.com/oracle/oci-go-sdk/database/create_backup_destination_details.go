// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateBackupDestinationDetails Details for creating a backup destination.
type CreateBackupDestinationDetails interface {

	// The user-provided name of the backup destination.
	GetDisplayName() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	GetDefinedTags() map[string]map[string]interface{}
}

type createbackupdestinationdetails struct {
	JsonData      []byte
	DisplayName   *string                           `mandatory:"true" json:"displayName"`
	CompartmentId *string                           `mandatory:"true" json:"compartmentId"`
	FreeformTags  map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags   map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	Type          string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createbackupdestinationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatebackupdestinationdetails createbackupdestinationdetails
	s := struct {
		Model Unmarshalercreatebackupdestinationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createbackupdestinationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "NFS":
		mm := CreateNfsBackupDestinationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RECOVERY_APPLIANCE":
		mm := CreateRecoveryApplianceBackupDestinationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDisplayName returns DisplayName
func (m createbackupdestinationdetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetCompartmentId returns CompartmentId
func (m createbackupdestinationdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetFreeformTags returns FreeformTags
func (m createbackupdestinationdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m createbackupdestinationdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m createbackupdestinationdetails) String() string {
	return common.PointerString(m)
}

// CreateBackupDestinationDetailsTypeEnum Enum with underlying type: string
type CreateBackupDestinationDetailsTypeEnum string

// Set of constants representing the allowable values for CreateBackupDestinationDetailsTypeEnum
const (
	CreateBackupDestinationDetailsTypeNfs               CreateBackupDestinationDetailsTypeEnum = "NFS"
	CreateBackupDestinationDetailsTypeRecoveryAppliance CreateBackupDestinationDetailsTypeEnum = "RECOVERY_APPLIANCE"
)

var mappingCreateBackupDestinationDetailsType = map[string]CreateBackupDestinationDetailsTypeEnum{
	"NFS":                CreateBackupDestinationDetailsTypeNfs,
	"RECOVERY_APPLIANCE": CreateBackupDestinationDetailsTypeRecoveryAppliance,
}

// GetCreateBackupDestinationDetailsTypeEnumValues Enumerates the set of values for CreateBackupDestinationDetailsTypeEnum
func GetCreateBackupDestinationDetailsTypeEnumValues() []CreateBackupDestinationDetailsTypeEnum {
	values := make([]CreateBackupDestinationDetailsTypeEnum, 0)
	for _, v := range mappingCreateBackupDestinationDetailsType {
		values = append(values, v)
	}
	return values
}
