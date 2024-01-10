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

// WorkersSummary Details of the workers in a specific On-premise vantage point.
type WorkersSummary struct {

	// Total number of workers in a specific On-premise vantage point.
	Total *int `mandatory:"true" json:"total"`

	// Number of available workers in a specific On-premise vantage point.
	Available *int `mandatory:"true" json:"available"`

	// Number of occupied workers in a specific On-premise vantage point.
	Used *int `mandatory:"true" json:"used"`

	// Number of disabled workers in a specific On-premise vantage point.
	Disabled *int `mandatory:"true" json:"disabled"`

	// Minimum version among the workers in a specific On-premise vantage point.
	MinVersion *string `mandatory:"true" json:"minVersion"`

	// List of available capabilities in a specific On-premise vantage point.
	AvailableCapabilities []AvailableCapability `mandatory:"true" json:"availableCapabilities"`
}

func (m WorkersSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkersSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
