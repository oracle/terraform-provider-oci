// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TestMetricExtensionData The Test result details
type TestMetricExtensionData struct {

	// Test Run Id
	TestRunId *string `mandatory:"true" json:"testRunId"`

	// Test Run Metric Suffix
	TestRunMetricSuffix *string `mandatory:"true" json:"testRunMetricSuffix"`

	// Test Run Namespace name
	TestRunNamespaceName *string `mandatory:"true" json:"testRunNamespaceName"`

	// Test Run Resource Group name
	TestRunResourceGroupName *string `mandatory:"false" json:"testRunResourceGroupName"`
}

func (m TestMetricExtensionData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TestMetricExtensionData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
