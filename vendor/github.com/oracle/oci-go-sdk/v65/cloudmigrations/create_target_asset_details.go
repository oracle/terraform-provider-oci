// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// CreateTargetAssetDetails Details of the new target asset.
type CreateTargetAssetDetails interface {

	// OCID of the associated migration plan.
	GetMigrationPlanId() *string

	// A boolean indicating whether the asset should be migrated.
	GetIsExcludedFromExecution() *bool
}

type createtargetassetdetails struct {
	JsonData                []byte
	MigrationPlanId         *string `mandatory:"true" json:"migrationPlanId"`
	IsExcludedFromExecution *bool   `mandatory:"true" json:"isExcludedFromExecution"`
	Type                    string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createtargetassetdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatetargetassetdetails createtargetassetdetails
	s := struct {
		Model Unmarshalercreatetargetassetdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.MigrationPlanId = s.Model.MigrationPlanId
	m.IsExcludedFromExecution = s.Model.IsExcludedFromExecution
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createtargetassetdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "INSTANCE":
		mm := CreateVmTargetAssetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateTargetAssetDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetMigrationPlanId returns MigrationPlanId
func (m createtargetassetdetails) GetMigrationPlanId() *string {
	return m.MigrationPlanId
}

// GetIsExcludedFromExecution returns IsExcludedFromExecution
func (m createtargetassetdetails) GetIsExcludedFromExecution() *bool {
	return m.IsExcludedFromExecution
}

func (m createtargetassetdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createtargetassetdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateTargetAssetDetailsTypeEnum Enum with underlying type: string
type CreateTargetAssetDetailsTypeEnum string

// Set of constants representing the allowable values for CreateTargetAssetDetailsTypeEnum
const (
	CreateTargetAssetDetailsTypeInstance CreateTargetAssetDetailsTypeEnum = "INSTANCE"
)

var mappingCreateTargetAssetDetailsTypeEnum = map[string]CreateTargetAssetDetailsTypeEnum{
	"INSTANCE": CreateTargetAssetDetailsTypeInstance,
}

var mappingCreateTargetAssetDetailsTypeEnumLowerCase = map[string]CreateTargetAssetDetailsTypeEnum{
	"instance": CreateTargetAssetDetailsTypeInstance,
}

// GetCreateTargetAssetDetailsTypeEnumValues Enumerates the set of values for CreateTargetAssetDetailsTypeEnum
func GetCreateTargetAssetDetailsTypeEnumValues() []CreateTargetAssetDetailsTypeEnum {
	values := make([]CreateTargetAssetDetailsTypeEnum, 0)
	for _, v := range mappingCreateTargetAssetDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateTargetAssetDetailsTypeEnumStringValues Enumerates the set of values in String for CreateTargetAssetDetailsTypeEnum
func GetCreateTargetAssetDetailsTypeEnumStringValues() []string {
	return []string{
		"INSTANCE",
	}
}

// GetMappingCreateTargetAssetDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateTargetAssetDetailsTypeEnum(val string) (CreateTargetAssetDetailsTypeEnum, bool) {
	enum, ok := mappingCreateTargetAssetDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
