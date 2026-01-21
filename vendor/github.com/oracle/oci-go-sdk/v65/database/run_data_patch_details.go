// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RunDataPatchDetails Details for running datapatch operation on a database and its pluggable databases
type RunDataPatchDetails struct {

	// List of Pluggable Database OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to run datapatch on. The datapatch would run on the database first and then the given pluggable databases.
	PluggableDatabases []string `mandatory:"false" json:"pluggableDatabases"`

	DataPatchOptions *DataPatchOptions `mandatory:"false" json:"dataPatchOptions"`

	// The action to perform on run database dataPatch operation
	Action RunDataPatchDetailsActionEnum `mandatory:"false" json:"action,omitempty"`
}

func (m RunDataPatchDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RunDataPatchDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRunDataPatchDetailsActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetRunDataPatchDetailsActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RunDataPatchDetailsActionEnum Enum with underlying type: string
type RunDataPatchDetailsActionEnum string

// Set of constants representing the allowable values for RunDataPatchDetailsActionEnum
const (
	RunDataPatchDetailsActionApply    RunDataPatchDetailsActionEnum = "APPLY"
	RunDataPatchDetailsActionPrecheck RunDataPatchDetailsActionEnum = "PRECHECK"
)

var mappingRunDataPatchDetailsActionEnum = map[string]RunDataPatchDetailsActionEnum{
	"APPLY":    RunDataPatchDetailsActionApply,
	"PRECHECK": RunDataPatchDetailsActionPrecheck,
}

var mappingRunDataPatchDetailsActionEnumLowerCase = map[string]RunDataPatchDetailsActionEnum{
	"apply":    RunDataPatchDetailsActionApply,
	"precheck": RunDataPatchDetailsActionPrecheck,
}

// GetRunDataPatchDetailsActionEnumValues Enumerates the set of values for RunDataPatchDetailsActionEnum
func GetRunDataPatchDetailsActionEnumValues() []RunDataPatchDetailsActionEnum {
	values := make([]RunDataPatchDetailsActionEnum, 0)
	for _, v := range mappingRunDataPatchDetailsActionEnum {
		values = append(values, v)
	}
	return values
}

// GetRunDataPatchDetailsActionEnumStringValues Enumerates the set of values in String for RunDataPatchDetailsActionEnum
func GetRunDataPatchDetailsActionEnumStringValues() []string {
	return []string{
		"APPLY",
		"PRECHECK",
	}
}

// GetMappingRunDataPatchDetailsActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRunDataPatchDetailsActionEnum(val string) (RunDataPatchDetailsActionEnum, bool) {
	enum, ok := mappingRunDataPatchDetailsActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
