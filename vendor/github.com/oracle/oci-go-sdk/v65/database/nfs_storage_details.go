// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NfsStorageDetails Details for attaching/detaching a NFS Storage to an ACD.
type NfsStorageDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the NFS Storage.
	NfsStorageId *string `mandatory:"false" json:"nfsStorageId"`

	// NFS attach/detach operation.
	ActionType NfsStorageDetailsActionTypeEnum `mandatory:"false" json:"actionType,omitempty"`
}

func (m NfsStorageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NfsStorageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingNfsStorageDetailsActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetNfsStorageDetailsActionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NfsStorageDetailsActionTypeEnum Enum with underlying type: string
type NfsStorageDetailsActionTypeEnum string

// Set of constants representing the allowable values for NfsStorageDetailsActionTypeEnum
const (
	NfsStorageDetailsActionTypeAttach NfsStorageDetailsActionTypeEnum = "ATTACH"
	NfsStorageDetailsActionTypeDetach NfsStorageDetailsActionTypeEnum = "DETACH"
)

var mappingNfsStorageDetailsActionTypeEnum = map[string]NfsStorageDetailsActionTypeEnum{
	"ATTACH": NfsStorageDetailsActionTypeAttach,
	"DETACH": NfsStorageDetailsActionTypeDetach,
}

var mappingNfsStorageDetailsActionTypeEnumLowerCase = map[string]NfsStorageDetailsActionTypeEnum{
	"attach": NfsStorageDetailsActionTypeAttach,
	"detach": NfsStorageDetailsActionTypeDetach,
}

// GetNfsStorageDetailsActionTypeEnumValues Enumerates the set of values for NfsStorageDetailsActionTypeEnum
func GetNfsStorageDetailsActionTypeEnumValues() []NfsStorageDetailsActionTypeEnum {
	values := make([]NfsStorageDetailsActionTypeEnum, 0)
	for _, v := range mappingNfsStorageDetailsActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetNfsStorageDetailsActionTypeEnumStringValues Enumerates the set of values in String for NfsStorageDetailsActionTypeEnum
func GetNfsStorageDetailsActionTypeEnumStringValues() []string {
	return []string{
		"ATTACH",
		"DETACH",
	}
}

// GetMappingNfsStorageDetailsActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNfsStorageDetailsActionTypeEnum(val string) (NfsStorageDetailsActionTypeEnum, bool) {
	enum, ok := mappingNfsStorageDetailsActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
