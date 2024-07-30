// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TargetDetails Details of target member of a Exadata Fleet Update Collection.
type TargetDetails interface {

	// OCID of the target resource in the Exadata Fleet Update Collection.
	GetId() *string

	// Compartment identifier of the target.
	GetCompartmentId() *string
}

type targetdetails struct {
	JsonData      []byte
	Id            *string `mandatory:"false" json:"id"`
	CompartmentId *string `mandatory:"false" json:"compartmentId"`
	EntityType    string  `json:"entityType"`
}

// UnmarshalJSON unmarshals json
func (m *targetdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertargetdetails targetdetails
	s := struct {
		Model Unmarshalertargetdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.EntityType = s.Model.EntityType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *targetdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntityType {
	case "VMCLUSTER":
		mm := VmClusterTargetSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CLOUDVMCLUSTER":
		mm := CloudVmClusterTargetSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATABASE":
		mm := DatabaseTargetSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for TargetDetails: %s.", m.EntityType)
		return *m, nil
	}
}

// GetId returns Id
func (m targetdetails) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m targetdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m targetdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m targetdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TargetDetailsEntityTypeEnum Enum with underlying type: string
type TargetDetailsEntityTypeEnum string

// Set of constants representing the allowable values for TargetDetailsEntityTypeEnum
const (
	TargetDetailsEntityTypeDatabase       TargetDetailsEntityTypeEnum = "DATABASE"
	TargetDetailsEntityTypeVmcluster      TargetDetailsEntityTypeEnum = "VMCLUSTER"
	TargetDetailsEntityTypeCloudvmcluster TargetDetailsEntityTypeEnum = "CLOUDVMCLUSTER"
)

var mappingTargetDetailsEntityTypeEnum = map[string]TargetDetailsEntityTypeEnum{
	"DATABASE":       TargetDetailsEntityTypeDatabase,
	"VMCLUSTER":      TargetDetailsEntityTypeVmcluster,
	"CLOUDVMCLUSTER": TargetDetailsEntityTypeCloudvmcluster,
}

var mappingTargetDetailsEntityTypeEnumLowerCase = map[string]TargetDetailsEntityTypeEnum{
	"database":       TargetDetailsEntityTypeDatabase,
	"vmcluster":      TargetDetailsEntityTypeVmcluster,
	"cloudvmcluster": TargetDetailsEntityTypeCloudvmcluster,
}

// GetTargetDetailsEntityTypeEnumValues Enumerates the set of values for TargetDetailsEntityTypeEnum
func GetTargetDetailsEntityTypeEnumValues() []TargetDetailsEntityTypeEnum {
	values := make([]TargetDetailsEntityTypeEnum, 0)
	for _, v := range mappingTargetDetailsEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetDetailsEntityTypeEnumStringValues Enumerates the set of values in String for TargetDetailsEntityTypeEnum
func GetTargetDetailsEntityTypeEnumStringValues() []string {
	return []string{
		"DATABASE",
		"VMCLUSTER",
		"CLOUDVMCLUSTER",
	}
}

// GetMappingTargetDetailsEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTargetDetailsEntityTypeEnum(val string) (TargetDetailsEntityTypeEnum, bool) {
	enum, ok := mappingTargetDetailsEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
