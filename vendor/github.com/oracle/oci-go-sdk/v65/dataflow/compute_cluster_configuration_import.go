// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComputeClusterConfigurationImport Response for an accepted Compute Cluster configuration import request.
type ComputeClusterConfigurationImport struct {

	// The OCID of the Compute Cluster.
	ComputeClusterId *string `mandatory:"false" json:"computeClusterId"`

	// The library IDs created by the configuration import operation.
	LibraryIds []string `mandatory:"false" json:"libraryIds"`

	// The lifecycle state of the configuration import request.
	State ComputeClusterConfigurationImportStateEnum `mandatory:"false" json:"state,omitempty"`
}

func (m ComputeClusterConfigurationImport) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeClusterConfigurationImport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingComputeClusterConfigurationImportStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetComputeClusterConfigurationImportStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ComputeClusterConfigurationImportStateEnum Enum with underlying type: string
type ComputeClusterConfigurationImportStateEnum string

// Set of constants representing the allowable values for ComputeClusterConfigurationImportStateEnum
const (
	ComputeClusterConfigurationImportStatePending ComputeClusterConfigurationImportStateEnum = "PENDING"
)

var mappingComputeClusterConfigurationImportStateEnum = map[string]ComputeClusterConfigurationImportStateEnum{
	"PENDING": ComputeClusterConfigurationImportStatePending,
}

var mappingComputeClusterConfigurationImportStateEnumLowerCase = map[string]ComputeClusterConfigurationImportStateEnum{
	"pending": ComputeClusterConfigurationImportStatePending,
}

// GetComputeClusterConfigurationImportStateEnumValues Enumerates the set of values for ComputeClusterConfigurationImportStateEnum
func GetComputeClusterConfigurationImportStateEnumValues() []ComputeClusterConfigurationImportStateEnum {
	values := make([]ComputeClusterConfigurationImportStateEnum, 0)
	for _, v := range mappingComputeClusterConfigurationImportStateEnum {
		values = append(values, v)
	}
	return values
}

// GetComputeClusterConfigurationImportStateEnumStringValues Enumerates the set of values in String for ComputeClusterConfigurationImportStateEnum
func GetComputeClusterConfigurationImportStateEnumStringValues() []string {
	return []string{
		"PENDING",
	}
}

// GetMappingComputeClusterConfigurationImportStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputeClusterConfigurationImportStateEnum(val string) (ComputeClusterConfigurationImportStateEnum, bool) {
	enum, ok := mappingComputeClusterConfigurationImportStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
