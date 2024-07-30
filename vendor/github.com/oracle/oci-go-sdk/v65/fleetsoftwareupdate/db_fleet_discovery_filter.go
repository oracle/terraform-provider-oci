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

// DbFleetDiscoveryFilter Possible Discovery filters for Database targets.
type DbFleetDiscoveryFilter interface {

	// INCLUDE or EXCLUDE the filter results in the discovery for DB targets.
	// Supported for 'FSUCOLLECTION' RESOURCE_ID filter only.
	GetMode() DbFleetDiscoveryFilterModeEnum
}

type dbfleetdiscoveryfilter struct {
	JsonData []byte
	Mode     DbFleetDiscoveryFilterModeEnum `mandatory:"false" json:"mode,omitempty"`
	Type     string                         `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *dbfleetdiscoveryfilter) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdbfleetdiscoveryfilter dbfleetdiscoveryfilter
	s := struct {
		Model Unmarshalerdbfleetdiscoveryfilter
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Mode = s.Model.Mode
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *dbfleetdiscoveryfilter) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DEFINED_TAG":
		mm := DbDefinedTagsFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DB_UNIQUE_NAME":
		mm := DbUniqueNameFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VERSION":
		mm := DbVersionFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RESOURCE_ID":
		mm := DbResourceIdFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DB_HOME_NAME":
		mm := DbHomeNameFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPARTMENT_ID":
		mm := DbCompartmentIdFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DB_NAME":
		mm := DbNameFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FREEFORM_TAG":
		mm := DbFreeformTagsFilter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DbFleetDiscoveryFilter: %s.", m.Type)
		return *m, nil
	}
}

// GetMode returns Mode
func (m dbfleetdiscoveryfilter) GetMode() DbFleetDiscoveryFilterModeEnum {
	return m.Mode
}

func (m dbfleetdiscoveryfilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m dbfleetdiscoveryfilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDbFleetDiscoveryFilterModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetDbFleetDiscoveryFilterModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbFleetDiscoveryFilterModeEnum Enum with underlying type: string
type DbFleetDiscoveryFilterModeEnum string

// Set of constants representing the allowable values for DbFleetDiscoveryFilterModeEnum
const (
	DbFleetDiscoveryFilterModeInclude DbFleetDiscoveryFilterModeEnum = "INCLUDE"
	DbFleetDiscoveryFilterModeExclude DbFleetDiscoveryFilterModeEnum = "EXCLUDE"
)

var mappingDbFleetDiscoveryFilterModeEnum = map[string]DbFleetDiscoveryFilterModeEnum{
	"INCLUDE": DbFleetDiscoveryFilterModeInclude,
	"EXCLUDE": DbFleetDiscoveryFilterModeExclude,
}

var mappingDbFleetDiscoveryFilterModeEnumLowerCase = map[string]DbFleetDiscoveryFilterModeEnum{
	"include": DbFleetDiscoveryFilterModeInclude,
	"exclude": DbFleetDiscoveryFilterModeExclude,
}

// GetDbFleetDiscoveryFilterModeEnumValues Enumerates the set of values for DbFleetDiscoveryFilterModeEnum
func GetDbFleetDiscoveryFilterModeEnumValues() []DbFleetDiscoveryFilterModeEnum {
	values := make([]DbFleetDiscoveryFilterModeEnum, 0)
	for _, v := range mappingDbFleetDiscoveryFilterModeEnum {
		values = append(values, v)
	}
	return values
}

// GetDbFleetDiscoveryFilterModeEnumStringValues Enumerates the set of values in String for DbFleetDiscoveryFilterModeEnum
func GetDbFleetDiscoveryFilterModeEnumStringValues() []string {
	return []string{
		"INCLUDE",
		"EXCLUDE",
	}
}

// GetMappingDbFleetDiscoveryFilterModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbFleetDiscoveryFilterModeEnum(val string) (DbFleetDiscoveryFilterModeEnum, bool) {
	enum, ok := mappingDbFleetDiscoveryFilterModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DbFleetDiscoveryFilterTypeEnum Enum with underlying type: string
type DbFleetDiscoveryFilterTypeEnum string

// Set of constants representing the allowable values for DbFleetDiscoveryFilterTypeEnum
const (
	DbFleetDiscoveryFilterTypeCompartmentId DbFleetDiscoveryFilterTypeEnum = "COMPARTMENT_ID"
	DbFleetDiscoveryFilterTypeVersion       DbFleetDiscoveryFilterTypeEnum = "VERSION"
	DbFleetDiscoveryFilterTypeDbName        DbFleetDiscoveryFilterTypeEnum = "DB_NAME"
	DbFleetDiscoveryFilterTypeDbUniqueName  DbFleetDiscoveryFilterTypeEnum = "DB_UNIQUE_NAME"
	DbFleetDiscoveryFilterTypeDbHomeName    DbFleetDiscoveryFilterTypeEnum = "DB_HOME_NAME"
	DbFleetDiscoveryFilterTypeFreeformTag   DbFleetDiscoveryFilterTypeEnum = "FREEFORM_TAG"
	DbFleetDiscoveryFilterTypeDefinedTag    DbFleetDiscoveryFilterTypeEnum = "DEFINED_TAG"
	DbFleetDiscoveryFilterTypeResourceId    DbFleetDiscoveryFilterTypeEnum = "RESOURCE_ID"
)

var mappingDbFleetDiscoveryFilterTypeEnum = map[string]DbFleetDiscoveryFilterTypeEnum{
	"COMPARTMENT_ID": DbFleetDiscoveryFilterTypeCompartmentId,
	"VERSION":        DbFleetDiscoveryFilterTypeVersion,
	"DB_NAME":        DbFleetDiscoveryFilterTypeDbName,
	"DB_UNIQUE_NAME": DbFleetDiscoveryFilterTypeDbUniqueName,
	"DB_HOME_NAME":   DbFleetDiscoveryFilterTypeDbHomeName,
	"FREEFORM_TAG":   DbFleetDiscoveryFilterTypeFreeformTag,
	"DEFINED_TAG":    DbFleetDiscoveryFilterTypeDefinedTag,
	"RESOURCE_ID":    DbFleetDiscoveryFilterTypeResourceId,
}

var mappingDbFleetDiscoveryFilterTypeEnumLowerCase = map[string]DbFleetDiscoveryFilterTypeEnum{
	"compartment_id": DbFleetDiscoveryFilterTypeCompartmentId,
	"version":        DbFleetDiscoveryFilterTypeVersion,
	"db_name":        DbFleetDiscoveryFilterTypeDbName,
	"db_unique_name": DbFleetDiscoveryFilterTypeDbUniqueName,
	"db_home_name":   DbFleetDiscoveryFilterTypeDbHomeName,
	"freeform_tag":   DbFleetDiscoveryFilterTypeFreeformTag,
	"defined_tag":    DbFleetDiscoveryFilterTypeDefinedTag,
	"resource_id":    DbFleetDiscoveryFilterTypeResourceId,
}

// GetDbFleetDiscoveryFilterTypeEnumValues Enumerates the set of values for DbFleetDiscoveryFilterTypeEnum
func GetDbFleetDiscoveryFilterTypeEnumValues() []DbFleetDiscoveryFilterTypeEnum {
	values := make([]DbFleetDiscoveryFilterTypeEnum, 0)
	for _, v := range mappingDbFleetDiscoveryFilterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDbFleetDiscoveryFilterTypeEnumStringValues Enumerates the set of values in String for DbFleetDiscoveryFilterTypeEnum
func GetDbFleetDiscoveryFilterTypeEnumStringValues() []string {
	return []string{
		"COMPARTMENT_ID",
		"VERSION",
		"DB_NAME",
		"DB_UNIQUE_NAME",
		"DB_HOME_NAME",
		"FREEFORM_TAG",
		"DEFINED_TAG",
		"RESOURCE_ID",
	}
}

// GetMappingDbFleetDiscoveryFilterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbFleetDiscoveryFilterTypeEnum(val string) (DbFleetDiscoveryFilterTypeEnum, bool) {
	enum, ok := mappingDbFleetDiscoveryFilterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
