// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AggregatedNetworkData Details of the aggregated network data.
type AggregatedNetworkData struct {

	// Status of the aggregated network data result.
	ResultState AggregatedNetworkDataResultStateEnum `mandatory:"true" json:"resultState"`

	// List of vantage point nodes.
	VantagePointNodes []VantagePointNode `mandatory:"false" json:"vantagePointNodes"`

	// An array of node arrays where each internal array corresponds to nodes at one level.
	NodesByLevel [][]Node `mandatory:"false" json:"nodesByLevel"`

	// Map of link objects.
	Links map[string]Link `mandatory:"false" json:"links"`

	// String containing error details.
	ErrorDetails *string `mandatory:"false" json:"errorDetails"`
}

func (m AggregatedNetworkData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AggregatedNetworkData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAggregatedNetworkDataResultStateEnum(string(m.ResultState)); !ok && m.ResultState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResultState: %s. Supported values are: %s.", m.ResultState, strings.Join(GetAggregatedNetworkDataResultStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AggregatedNetworkDataResultStateEnum Enum with underlying type: string
type AggregatedNetworkDataResultStateEnum string

// Set of constants representing the allowable values for AggregatedNetworkDataResultStateEnum
const (
	AggregatedNetworkDataResultStateSuccess AggregatedNetworkDataResultStateEnum = "SUCCESS"
	AggregatedNetworkDataResultStateFailure AggregatedNetworkDataResultStateEnum = "FAILURE"
	AggregatedNetworkDataResultStatePartial AggregatedNetworkDataResultStateEnum = "PARTIAL"
)

var mappingAggregatedNetworkDataResultStateEnum = map[string]AggregatedNetworkDataResultStateEnum{
	"SUCCESS": AggregatedNetworkDataResultStateSuccess,
	"FAILURE": AggregatedNetworkDataResultStateFailure,
	"PARTIAL": AggregatedNetworkDataResultStatePartial,
}

var mappingAggregatedNetworkDataResultStateEnumLowerCase = map[string]AggregatedNetworkDataResultStateEnum{
	"success": AggregatedNetworkDataResultStateSuccess,
	"failure": AggregatedNetworkDataResultStateFailure,
	"partial": AggregatedNetworkDataResultStatePartial,
}

// GetAggregatedNetworkDataResultStateEnumValues Enumerates the set of values for AggregatedNetworkDataResultStateEnum
func GetAggregatedNetworkDataResultStateEnumValues() []AggregatedNetworkDataResultStateEnum {
	values := make([]AggregatedNetworkDataResultStateEnum, 0)
	for _, v := range mappingAggregatedNetworkDataResultStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAggregatedNetworkDataResultStateEnumStringValues Enumerates the set of values in String for AggregatedNetworkDataResultStateEnum
func GetAggregatedNetworkDataResultStateEnumStringValues() []string {
	return []string{
		"SUCCESS",
		"FAILURE",
		"PARTIAL",
	}
}

// GetMappingAggregatedNetworkDataResultStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAggregatedNetworkDataResultStateEnum(val string) (AggregatedNetworkDataResultStateEnum, bool) {
	enum, ok := mappingAggregatedNetworkDataResultStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
