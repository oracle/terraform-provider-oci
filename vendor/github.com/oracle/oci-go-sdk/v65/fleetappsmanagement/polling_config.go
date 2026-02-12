// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PollingConfig Configuration for polling a status endpoint until stop condition is met.
type PollingConfig struct {

	// Whether this is polling api task.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// When set to true, the expected response is used to determine the stop condition during polling. If false, a separate stopCondition must be defined.
	IsStopConditionSameAsExpectedResponse *bool `mandatory:"false" json:"isStopConditionSameAsExpectedResponse"`

	StopCondition *ExpectedResponseCondition `mandatory:"false" json:"stopCondition"`
}

func (m PollingConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PollingConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
