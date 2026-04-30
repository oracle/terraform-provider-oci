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

// DataSafeRegistrationRequestDetails The requested Data Safe registration action details for the database.
type DataSafeRegistrationRequestDetails struct {

	// The requested Data Safe registration action for the database.
	Action DataSafeRegistrationRequestDetailsActionEnum `mandatory:"true" json:"action"`

	ConnectionOption *ConnectionOption `mandatory:"false" json:"connectionOption"`
}

func (m DataSafeRegistrationRequestDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataSafeRegistrationRequestDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataSafeRegistrationRequestDetailsActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetDataSafeRegistrationRequestDetailsActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DataSafeRegistrationRequestDetailsActionEnum Enum with underlying type: string
type DataSafeRegistrationRequestDetailsActionEnum string

// Set of constants representing the allowable values for DataSafeRegistrationRequestDetailsActionEnum
const (
	DataSafeRegistrationRequestDetailsActionRegister   DataSafeRegistrationRequestDetailsActionEnum = "REGISTER"
	DataSafeRegistrationRequestDetailsActionUnregister DataSafeRegistrationRequestDetailsActionEnum = "UNREGISTER"
)

var mappingDataSafeRegistrationRequestDetailsActionEnum = map[string]DataSafeRegistrationRequestDetailsActionEnum{
	"REGISTER":   DataSafeRegistrationRequestDetailsActionRegister,
	"UNREGISTER": DataSafeRegistrationRequestDetailsActionUnregister,
}

var mappingDataSafeRegistrationRequestDetailsActionEnumLowerCase = map[string]DataSafeRegistrationRequestDetailsActionEnum{
	"register":   DataSafeRegistrationRequestDetailsActionRegister,
	"unregister": DataSafeRegistrationRequestDetailsActionUnregister,
}

// GetDataSafeRegistrationRequestDetailsActionEnumValues Enumerates the set of values for DataSafeRegistrationRequestDetailsActionEnum
func GetDataSafeRegistrationRequestDetailsActionEnumValues() []DataSafeRegistrationRequestDetailsActionEnum {
	values := make([]DataSafeRegistrationRequestDetailsActionEnum, 0)
	for _, v := range mappingDataSafeRegistrationRequestDetailsActionEnum {
		values = append(values, v)
	}
	return values
}

// GetDataSafeRegistrationRequestDetailsActionEnumStringValues Enumerates the set of values in String for DataSafeRegistrationRequestDetailsActionEnum
func GetDataSafeRegistrationRequestDetailsActionEnumStringValues() []string {
	return []string{
		"REGISTER",
		"UNREGISTER",
	}
}

// GetMappingDataSafeRegistrationRequestDetailsActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataSafeRegistrationRequestDetailsActionEnum(val string) (DataSafeRegistrationRequestDetailsActionEnum, bool) {
	enum, ok := mappingDataSafeRegistrationRequestDetailsActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
