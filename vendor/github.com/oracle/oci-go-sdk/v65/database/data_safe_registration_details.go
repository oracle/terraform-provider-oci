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

// DataSafeRegistrationDetails The Data Safe registration details of the database.
type DataSafeRegistrationDetails struct {

	// The Data Safe registration status of the database.
	RegistrationState DataSafeRegistrationDetailsRegistrationStateEnum `mandatory:"false" json:"registrationState,omitempty"`
}

func (m DataSafeRegistrationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataSafeRegistrationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDataSafeRegistrationDetailsRegistrationStateEnum(string(m.RegistrationState)); !ok && m.RegistrationState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RegistrationState: %s. Supported values are: %s.", m.RegistrationState, strings.Join(GetDataSafeRegistrationDetailsRegistrationStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DataSafeRegistrationDetailsRegistrationStateEnum Enum with underlying type: string
type DataSafeRegistrationDetailsRegistrationStateEnum string

// Set of constants representing the allowable values for DataSafeRegistrationDetailsRegistrationStateEnum
const (
	DataSafeRegistrationDetailsRegistrationStateRegistered   DataSafeRegistrationDetailsRegistrationStateEnum = "REGISTERED"
	DataSafeRegistrationDetailsRegistrationStateUnregistered DataSafeRegistrationDetailsRegistrationStateEnum = "UNREGISTERED"
	DataSafeRegistrationDetailsRegistrationStateFailed       DataSafeRegistrationDetailsRegistrationStateEnum = "FAILED"
)

var mappingDataSafeRegistrationDetailsRegistrationStateEnum = map[string]DataSafeRegistrationDetailsRegistrationStateEnum{
	"REGISTERED":   DataSafeRegistrationDetailsRegistrationStateRegistered,
	"UNREGISTERED": DataSafeRegistrationDetailsRegistrationStateUnregistered,
	"FAILED":       DataSafeRegistrationDetailsRegistrationStateFailed,
}

var mappingDataSafeRegistrationDetailsRegistrationStateEnumLowerCase = map[string]DataSafeRegistrationDetailsRegistrationStateEnum{
	"registered":   DataSafeRegistrationDetailsRegistrationStateRegistered,
	"unregistered": DataSafeRegistrationDetailsRegistrationStateUnregistered,
	"failed":       DataSafeRegistrationDetailsRegistrationStateFailed,
}

// GetDataSafeRegistrationDetailsRegistrationStateEnumValues Enumerates the set of values for DataSafeRegistrationDetailsRegistrationStateEnum
func GetDataSafeRegistrationDetailsRegistrationStateEnumValues() []DataSafeRegistrationDetailsRegistrationStateEnum {
	values := make([]DataSafeRegistrationDetailsRegistrationStateEnum, 0)
	for _, v := range mappingDataSafeRegistrationDetailsRegistrationStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDataSafeRegistrationDetailsRegistrationStateEnumStringValues Enumerates the set of values in String for DataSafeRegistrationDetailsRegistrationStateEnum
func GetDataSafeRegistrationDetailsRegistrationStateEnumStringValues() []string {
	return []string{
		"REGISTERED",
		"UNREGISTERED",
		"FAILED",
	}
}

// GetMappingDataSafeRegistrationDetailsRegistrationStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataSafeRegistrationDetailsRegistrationStateEnum(val string) (DataSafeRegistrationDetailsRegistrationStateEnum, bool) {
	enum, ok := mappingDataSafeRegistrationDetailsRegistrationStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
