// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the Data Connectivity Management Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConnectivityUsage Contains details of ConnectivityUsage.
type ConnectivityUsage struct {

	// The status of the usage report/update.
	Status ConnectivityUsageStatusEnum `mandatory:"true" json:"status"`

	// Error message when usage report/update.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`
}

func (m ConnectivityUsage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConnectivityUsage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConnectivityUsageStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetConnectivityUsageStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConnectivityUsageStatusEnum Enum with underlying type: string
type ConnectivityUsageStatusEnum string

// Set of constants representing the allowable values for ConnectivityUsageStatusEnum
const (
	ConnectivityUsageStatusFailed  ConnectivityUsageStatusEnum = "FAILED"
	ConnectivityUsageStatusSuccess ConnectivityUsageStatusEnum = "SUCCESS"
)

var mappingConnectivityUsageStatusEnum = map[string]ConnectivityUsageStatusEnum{
	"FAILED":  ConnectivityUsageStatusFailed,
	"SUCCESS": ConnectivityUsageStatusSuccess,
}

var mappingConnectivityUsageStatusEnumLowerCase = map[string]ConnectivityUsageStatusEnum{
	"failed":  ConnectivityUsageStatusFailed,
	"success": ConnectivityUsageStatusSuccess,
}

// GetConnectivityUsageStatusEnumValues Enumerates the set of values for ConnectivityUsageStatusEnum
func GetConnectivityUsageStatusEnumValues() []ConnectivityUsageStatusEnum {
	values := make([]ConnectivityUsageStatusEnum, 0)
	for _, v := range mappingConnectivityUsageStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetConnectivityUsageStatusEnumStringValues Enumerates the set of values in String for ConnectivityUsageStatusEnum
func GetConnectivityUsageStatusEnumStringValues() []string {
	return []string{
		"FAILED",
		"SUCCESS",
	}
}

// GetMappingConnectivityUsageStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectivityUsageStatusEnum(val string) (ConnectivityUsageStatusEnum, bool) {
	enum, ok := mappingConnectivityUsageStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
