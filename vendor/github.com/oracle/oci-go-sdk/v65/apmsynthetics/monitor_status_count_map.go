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

// MonitorStatusCountMap Details of the monitor count per state.
// Example: `{ "total" : 5, "enabled" : 3 , "disabled" : 2, "invalid" : 0 }`
type MonitorStatusCountMap struct {

	// Total number of monitors using the script.
	Total *int `mandatory:"true" json:"total"`

	// Number of enabled monitors using the script.
	Enabled *int `mandatory:"true" json:"enabled"`

	// Number of disabled monitors using the script.
	Disabled *int `mandatory:"true" json:"disabled"`

	// Number of invalid monitors using the script.
	Invalid *int `mandatory:"true" json:"invalid"`
}

func (m MonitorStatusCountMap) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MonitorStatusCountMap) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
