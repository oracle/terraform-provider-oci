// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InitialDataLoad Options required for the pipeline Initial Data Load. If enabled, copies existing data from source to target before replication.
type InitialDataLoad struct {

	// If ENABLED, then existing source data is also synchronized to the target when creating or updating the pipeline.
	IsInitialLoad InitialDataLoadIsInitialLoadEnum `mandatory:"true" json:"isInitialLoad"`

	// Action upon existing tables in target when initial Data Load is set i.e., isInitialLoad=true.
	ActionOnExistingTable InitialLoadActionEnum `mandatory:"false" json:"actionOnExistingTable,omitempty"`
}

func (m InitialDataLoad) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InitialDataLoad) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInitialDataLoadIsInitialLoadEnum(string(m.IsInitialLoad)); !ok && m.IsInitialLoad != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IsInitialLoad: %s. Supported values are: %s.", m.IsInitialLoad, strings.Join(GetInitialDataLoadIsInitialLoadEnumStringValues(), ",")))
	}

	if _, ok := GetMappingInitialLoadActionEnum(string(m.ActionOnExistingTable)); !ok && m.ActionOnExistingTable != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionOnExistingTable: %s. Supported values are: %s.", m.ActionOnExistingTable, strings.Join(GetInitialLoadActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InitialDataLoadIsInitialLoadEnum Enum with underlying type: string
type InitialDataLoadIsInitialLoadEnum string

// Set of constants representing the allowable values for InitialDataLoadIsInitialLoadEnum
const (
	InitialDataLoadIsInitialLoadEnabled  InitialDataLoadIsInitialLoadEnum = "ENABLED"
	InitialDataLoadIsInitialLoadDisabled InitialDataLoadIsInitialLoadEnum = "DISABLED"
)

var mappingInitialDataLoadIsInitialLoadEnum = map[string]InitialDataLoadIsInitialLoadEnum{
	"ENABLED":  InitialDataLoadIsInitialLoadEnabled,
	"DISABLED": InitialDataLoadIsInitialLoadDisabled,
}

var mappingInitialDataLoadIsInitialLoadEnumLowerCase = map[string]InitialDataLoadIsInitialLoadEnum{
	"enabled":  InitialDataLoadIsInitialLoadEnabled,
	"disabled": InitialDataLoadIsInitialLoadDisabled,
}

// GetInitialDataLoadIsInitialLoadEnumValues Enumerates the set of values for InitialDataLoadIsInitialLoadEnum
func GetInitialDataLoadIsInitialLoadEnumValues() []InitialDataLoadIsInitialLoadEnum {
	values := make([]InitialDataLoadIsInitialLoadEnum, 0)
	for _, v := range mappingInitialDataLoadIsInitialLoadEnum {
		values = append(values, v)
	}
	return values
}

// GetInitialDataLoadIsInitialLoadEnumStringValues Enumerates the set of values in String for InitialDataLoadIsInitialLoadEnum
func GetInitialDataLoadIsInitialLoadEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingInitialDataLoadIsInitialLoadEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInitialDataLoadIsInitialLoadEnum(val string) (InitialDataLoadIsInitialLoadEnum, bool) {
	enum, ok := mappingInitialDataLoadIsInitialLoadEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
