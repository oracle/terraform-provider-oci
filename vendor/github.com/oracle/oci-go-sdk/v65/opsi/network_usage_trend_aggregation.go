// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NetworkUsageTrendAggregation Usage data per network interface.
type NetworkUsageTrendAggregation struct {

	// Name of interface.
	InterfaceName *string `mandatory:"true" json:"interfaceName"`

	// Address that is connected to a computer network that uses the Internet Protocol for communication.
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// Unique identifier assigned to a network interface.
	MacAddress *string `mandatory:"true" json:"macAddress"`

	// List of usage data samples for a network interface.
	UsageData []NetworkUsageTrend `mandatory:"true" json:"usageData"`
}

func (m NetworkUsageTrendAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NetworkUsageTrendAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
