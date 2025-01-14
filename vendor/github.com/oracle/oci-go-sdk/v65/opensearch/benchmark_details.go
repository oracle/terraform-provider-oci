// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BenchmarkDetails Benchmark test configuration detail.
type BenchmarkDetails struct {

	// Execution Length of Perf Test (default set to STANDARD)
	PerfTestLength BenchmarkDetailsPerfTestLengthEnum `mandatory:"true" json:"perfTestLength"`

	// Control Cluster Id Information
	ControlClusterId *string `mandatory:"true" json:"controlClusterId"`

	// Required Test Cluster Id needed to run perf test
	Test1ClusterId *string `mandatory:"false" json:"test1ClusterId"`

	// Optional Test Cluster Id needed to run perf test
	Test2ClusterId *string `mandatory:"false" json:"test2ClusterId"`
}

func (m BenchmarkDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BenchmarkDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBenchmarkDetailsPerfTestLengthEnum(string(m.PerfTestLength)); !ok && m.PerfTestLength != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PerfTestLength: %s. Supported values are: %s.", m.PerfTestLength, strings.Join(GetBenchmarkDetailsPerfTestLengthEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BenchmarkDetailsPerfTestLengthEnum Enum with underlying type: string
type BenchmarkDetailsPerfTestLengthEnum string

// Set of constants representing the allowable values for BenchmarkDetailsPerfTestLengthEnum
const (
	BenchmarkDetailsPerfTestLengthStandard BenchmarkDetailsPerfTestLengthEnum = "STANDARD"
	BenchmarkDetailsPerfTestLengthLong     BenchmarkDetailsPerfTestLengthEnum = "LONG"
)

var mappingBenchmarkDetailsPerfTestLengthEnum = map[string]BenchmarkDetailsPerfTestLengthEnum{
	"STANDARD": BenchmarkDetailsPerfTestLengthStandard,
	"LONG":     BenchmarkDetailsPerfTestLengthLong,
}

var mappingBenchmarkDetailsPerfTestLengthEnumLowerCase = map[string]BenchmarkDetailsPerfTestLengthEnum{
	"standard": BenchmarkDetailsPerfTestLengthStandard,
	"long":     BenchmarkDetailsPerfTestLengthLong,
}

// GetBenchmarkDetailsPerfTestLengthEnumValues Enumerates the set of values for BenchmarkDetailsPerfTestLengthEnum
func GetBenchmarkDetailsPerfTestLengthEnumValues() []BenchmarkDetailsPerfTestLengthEnum {
	values := make([]BenchmarkDetailsPerfTestLengthEnum, 0)
	for _, v := range mappingBenchmarkDetailsPerfTestLengthEnum {
		values = append(values, v)
	}
	return values
}

// GetBenchmarkDetailsPerfTestLengthEnumStringValues Enumerates the set of values in String for BenchmarkDetailsPerfTestLengthEnum
func GetBenchmarkDetailsPerfTestLengthEnumStringValues() []string {
	return []string{
		"STANDARD",
		"LONG",
	}
}

// GetMappingBenchmarkDetailsPerfTestLengthEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBenchmarkDetailsPerfTestLengthEnum(val string) (BenchmarkDetailsPerfTestLengthEnum, bool) {
	enum, ok := mappingBenchmarkDetailsPerfTestLengthEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
