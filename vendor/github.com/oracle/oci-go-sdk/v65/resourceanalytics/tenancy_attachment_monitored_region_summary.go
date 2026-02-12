// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Analytics API
//
// Use the Resource Analytics API to manage Resource Analytics Instances.
//

package resourceanalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TenancyAttachmentMonitoredRegionSummary Information about a monitored region in a tenancy.
type TenancyAttachmentMonitoredRegionSummary struct {

	// The Region Identifier (https://docs.oracle.com/iaas/Content/General/Concepts/regions.htm) of the monitored region. E.g. us-ashburn-1
	RegionId *string `mandatory:"true" json:"regionId"`

	DataPopulation *TenancyAttachmentDataPopulation `mandatory:"true" json:"dataPopulation"`
}

func (m TenancyAttachmentMonitoredRegionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TenancyAttachmentMonitoredRegionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
