// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalExadataStorageConnectorStatus The status of an Exadata storage server connector.
type ExternalExadataStorageConnectorStatus struct {

	// The connection status of the connector.
	Status ExternalExadataStorageConnectorStatusStatusEnum `mandatory:"true" json:"status"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata storage server connector.
	Id *string `mandatory:"false" json:"id"`

	// The error message indicating the reason for failure or `null` if the connection was successful.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`
}

func (m ExternalExadataStorageConnectorStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalExadataStorageConnectorStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalExadataStorageConnectorStatusStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetExternalExadataStorageConnectorStatusStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExternalExadataStorageConnectorStatusStatusEnum Enum with underlying type: string
type ExternalExadataStorageConnectorStatusStatusEnum string

// Set of constants representing the allowable values for ExternalExadataStorageConnectorStatusStatusEnum
const (
	ExternalExadataStorageConnectorStatusStatusSucceeded ExternalExadataStorageConnectorStatusStatusEnum = "SUCCEEDED"
	ExternalExadataStorageConnectorStatusStatusFailed    ExternalExadataStorageConnectorStatusStatusEnum = "FAILED"
)

var mappingExternalExadataStorageConnectorStatusStatusEnum = map[string]ExternalExadataStorageConnectorStatusStatusEnum{
	"SUCCEEDED": ExternalExadataStorageConnectorStatusStatusSucceeded,
	"FAILED":    ExternalExadataStorageConnectorStatusStatusFailed,
}

var mappingExternalExadataStorageConnectorStatusStatusEnumLowerCase = map[string]ExternalExadataStorageConnectorStatusStatusEnum{
	"succeeded": ExternalExadataStorageConnectorStatusStatusSucceeded,
	"failed":    ExternalExadataStorageConnectorStatusStatusFailed,
}

// GetExternalExadataStorageConnectorStatusStatusEnumValues Enumerates the set of values for ExternalExadataStorageConnectorStatusStatusEnum
func GetExternalExadataStorageConnectorStatusStatusEnumValues() []ExternalExadataStorageConnectorStatusStatusEnum {
	values := make([]ExternalExadataStorageConnectorStatusStatusEnum, 0)
	for _, v := range mappingExternalExadataStorageConnectorStatusStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalExadataStorageConnectorStatusStatusEnumStringValues Enumerates the set of values in String for ExternalExadataStorageConnectorStatusStatusEnum
func GetExternalExadataStorageConnectorStatusStatusEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingExternalExadataStorageConnectorStatusStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalExadataStorageConnectorStatusStatusEnum(val string) (ExternalExadataStorageConnectorStatusStatusEnum, bool) {
	enum, ok := mappingExternalExadataStorageConnectorStatusStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
