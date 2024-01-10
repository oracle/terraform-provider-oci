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

// UpdateTargetAssetDetails Details of the updated target asset.
type UpdateTargetAssetDetails interface {

	// A boolean indicating whether the asset should be migrated.
	GetIsExcludedFromExecution() *bool
}

type updatetargetassetdetails struct {
	JsonData                []byte
	IsExcludedFromExecution *bool  `mandatory:"false" json:"isExcludedFromExecution"`
	Type                    string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *updatetargetassetdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatetargetassetdetails updatetargetassetdetails
	s := struct {
		Model Unmarshalerupdatetargetassetdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.IsExcludedFromExecution = s.Model.IsExcludedFromExecution
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatetargetassetdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "INSTANCE":
		mm := UpdateVmTargetAssetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UpdateTargetAssetDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetIsExcludedFromExecution returns IsExcludedFromExecution
func (m updatetargetassetdetails) GetIsExcludedFromExecution() *bool {
	return m.IsExcludedFromExecution
}

func (m updatetargetassetdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatetargetassetdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateTargetAssetDetailsTypeEnum Enum with underlying type: string
type UpdateTargetAssetDetailsTypeEnum string

// Set of constants representing the allowable values for UpdateTargetAssetDetailsTypeEnum
const (
	UpdateTargetAssetDetailsTypeInstance UpdateTargetAssetDetailsTypeEnum = "INSTANCE"
)

var mappingUpdateTargetAssetDetailsTypeEnum = map[string]UpdateTargetAssetDetailsTypeEnum{
	"INSTANCE": UpdateTargetAssetDetailsTypeInstance,
}

var mappingUpdateTargetAssetDetailsTypeEnumLowerCase = map[string]UpdateTargetAssetDetailsTypeEnum{
	"instance": UpdateTargetAssetDetailsTypeInstance,
}

// GetUpdateTargetAssetDetailsTypeEnumValues Enumerates the set of values for UpdateTargetAssetDetailsTypeEnum
func GetUpdateTargetAssetDetailsTypeEnumValues() []UpdateTargetAssetDetailsTypeEnum {
	values := make([]UpdateTargetAssetDetailsTypeEnum, 0)
	for _, v := range mappingUpdateTargetAssetDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateTargetAssetDetailsTypeEnumStringValues Enumerates the set of values in String for UpdateTargetAssetDetailsTypeEnum
func GetUpdateTargetAssetDetailsTypeEnumStringValues() []string {
	return []string{
		"INSTANCE",
	}
}

// GetMappingUpdateTargetAssetDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateTargetAssetDetailsTypeEnum(val string) (UpdateTargetAssetDetailsTypeEnum, bool) {
	enum, ok := mappingUpdateTargetAssetDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
