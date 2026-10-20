// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// SubsetScope The scope of the subset rule
type SubsetScope interface {
}

type subsetscope struct {
	JsonData  []byte
	ScopeType string `json:"scopeType"`
}

// UnmarshalJSON unmarshals json
func (m *subsetscope) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersubsetscope subsetscope
	s := struct {
		Model Unmarshalersubsetscope
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ScopeType = s.Model.ScopeType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *subsetscope) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ScopeType {
	case "ALL":
		mm := SubsetScopeForAllObjects{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SPECIFIC":
		mm := SubsetScopeForSpecificObjects{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for SubsetScope: %s.", m.ScopeType)
		return *m, nil
	}
}

func (m subsetscope) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m subsetscope) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SubsetScopeScopeTypeEnum Enum with underlying type: string
type SubsetScopeScopeTypeEnum string

// Set of constants representing the allowable values for SubsetScopeScopeTypeEnum
const (
	SubsetScopeScopeTypeAll      SubsetScopeScopeTypeEnum = "ALL"
	SubsetScopeScopeTypeSpecific SubsetScopeScopeTypeEnum = "SPECIFIC"
)

var mappingSubsetScopeScopeTypeEnum = map[string]SubsetScopeScopeTypeEnum{
	"ALL":      SubsetScopeScopeTypeAll,
	"SPECIFIC": SubsetScopeScopeTypeSpecific,
}

var mappingSubsetScopeScopeTypeEnumLowerCase = map[string]SubsetScopeScopeTypeEnum{
	"all":      SubsetScopeScopeTypeAll,
	"specific": SubsetScopeScopeTypeSpecific,
}

// GetSubsetScopeScopeTypeEnumValues Enumerates the set of values for SubsetScopeScopeTypeEnum
func GetSubsetScopeScopeTypeEnumValues() []SubsetScopeScopeTypeEnum {
	values := make([]SubsetScopeScopeTypeEnum, 0)
	for _, v := range mappingSubsetScopeScopeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSubsetScopeScopeTypeEnumStringValues Enumerates the set of values in String for SubsetScopeScopeTypeEnum
func GetSubsetScopeScopeTypeEnumStringValues() []string {
	return []string{
		"ALL",
		"SPECIFIC",
	}
}

// GetMappingSubsetScopeScopeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubsetScopeScopeTypeEnum(val string) (SubsetScopeScopeTypeEnum, bool) {
	enum, ok := mappingSubsetScopeScopeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
