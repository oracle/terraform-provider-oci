// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AdbServerlesTablespaceDetails Migration tablespace settings valid for ADB-D target type using remap feature
type AdbServerlesTablespaceDetails struct {

	// Name of tablespace at target to which the source database tablespace need to be remapped.
	RemapTarget AdbServerlesTablespaceDetailsRemapTargetEnum `mandatory:"false" json:"remapTarget,omitempty"`
}

func (m AdbServerlesTablespaceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AdbServerlesTablespaceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAdbServerlesTablespaceDetailsRemapTargetEnum(string(m.RemapTarget)); !ok && m.RemapTarget != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RemapTarget: %s. Supported values are: %s.", m.RemapTarget, strings.Join(GetAdbServerlesTablespaceDetailsRemapTargetEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AdbServerlesTablespaceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAdbServerlesTablespaceDetails AdbServerlesTablespaceDetails
	s := struct {
		DiscriminatorParam string `json:"targetType"`
		MarshalTypeAdbServerlesTablespaceDetails
	}{
		"ADB_S_REMAP",
		(MarshalTypeAdbServerlesTablespaceDetails)(m),
	}

	return json.Marshal(&s)
}

// AdbServerlesTablespaceDetailsRemapTargetEnum Enum with underlying type: string
type AdbServerlesTablespaceDetailsRemapTargetEnum string

// Set of constants representing the allowable values for AdbServerlesTablespaceDetailsRemapTargetEnum
const (
	AdbServerlesTablespaceDetailsRemapTargetData AdbServerlesTablespaceDetailsRemapTargetEnum = "DATA"
)

var mappingAdbServerlesTablespaceDetailsRemapTargetEnum = map[string]AdbServerlesTablespaceDetailsRemapTargetEnum{
	"DATA": AdbServerlesTablespaceDetailsRemapTargetData,
}

var mappingAdbServerlesTablespaceDetailsRemapTargetEnumLowerCase = map[string]AdbServerlesTablespaceDetailsRemapTargetEnum{
	"data": AdbServerlesTablespaceDetailsRemapTargetData,
}

// GetAdbServerlesTablespaceDetailsRemapTargetEnumValues Enumerates the set of values for AdbServerlesTablespaceDetailsRemapTargetEnum
func GetAdbServerlesTablespaceDetailsRemapTargetEnumValues() []AdbServerlesTablespaceDetailsRemapTargetEnum {
	values := make([]AdbServerlesTablespaceDetailsRemapTargetEnum, 0)
	for _, v := range mappingAdbServerlesTablespaceDetailsRemapTargetEnum {
		values = append(values, v)
	}
	return values
}

// GetAdbServerlesTablespaceDetailsRemapTargetEnumStringValues Enumerates the set of values in String for AdbServerlesTablespaceDetailsRemapTargetEnum
func GetAdbServerlesTablespaceDetailsRemapTargetEnumStringValues() []string {
	return []string{
		"DATA",
	}
}

// GetMappingAdbServerlesTablespaceDetailsRemapTargetEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAdbServerlesTablespaceDetailsRemapTargetEnum(val string) (AdbServerlesTablespaceDetailsRemapTargetEnum, bool) {
	enum, ok := mappingAdbServerlesTablespaceDetailsRemapTargetEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
