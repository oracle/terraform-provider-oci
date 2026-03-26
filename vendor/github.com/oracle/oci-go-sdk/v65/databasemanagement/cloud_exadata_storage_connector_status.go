// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CloudExadataStorageConnectorStatus The status of an Exadata storage server connector.
type CloudExadataStorageConnectorStatus struct {

	// The connection status of the connector.
	Status CloudExadataStorageConnectorStatusStatusEnum `mandatory:"true" json:"status"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata storage server connector.
	Id *string `mandatory:"false" json:"id"`

	// The error message indicating the reason for failure or `null` if the connection was successful.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`
}

func (m CloudExadataStorageConnectorStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudExadataStorageConnectorStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCloudExadataStorageConnectorStatusStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetCloudExadataStorageConnectorStatusStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CloudExadataStorageConnectorStatusStatusEnum Enum with underlying type: string
type CloudExadataStorageConnectorStatusStatusEnum string

// Set of constants representing the allowable values for CloudExadataStorageConnectorStatusStatusEnum
const (
	CloudExadataStorageConnectorStatusStatusSucceeded CloudExadataStorageConnectorStatusStatusEnum = "SUCCEEDED"
	CloudExadataStorageConnectorStatusStatusFailed    CloudExadataStorageConnectorStatusStatusEnum = "FAILED"
)

var mappingCloudExadataStorageConnectorStatusStatusEnum = map[string]CloudExadataStorageConnectorStatusStatusEnum{
	"SUCCEEDED": CloudExadataStorageConnectorStatusStatusSucceeded,
	"FAILED":    CloudExadataStorageConnectorStatusStatusFailed,
}

var mappingCloudExadataStorageConnectorStatusStatusEnumLowerCase = map[string]CloudExadataStorageConnectorStatusStatusEnum{
	"succeeded": CloudExadataStorageConnectorStatusStatusSucceeded,
	"failed":    CloudExadataStorageConnectorStatusStatusFailed,
}

// GetCloudExadataStorageConnectorStatusStatusEnumValues Enumerates the set of values for CloudExadataStorageConnectorStatusStatusEnum
func GetCloudExadataStorageConnectorStatusStatusEnumValues() []CloudExadataStorageConnectorStatusStatusEnum {
	values := make([]CloudExadataStorageConnectorStatusStatusEnum, 0)
	for _, v := range mappingCloudExadataStorageConnectorStatusStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudExadataStorageConnectorStatusStatusEnumStringValues Enumerates the set of values in String for CloudExadataStorageConnectorStatusStatusEnum
func GetCloudExadataStorageConnectorStatusStatusEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingCloudExadataStorageConnectorStatusStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudExadataStorageConnectorStatusStatusEnum(val string) (CloudExadataStorageConnectorStatusStatusEnum, bool) {
	enum, ok := mappingCloudExadataStorageConnectorStatusStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
