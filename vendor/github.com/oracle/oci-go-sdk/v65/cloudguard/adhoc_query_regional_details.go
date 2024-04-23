// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AdhocQueryRegionalDetails Instance level status for each region.
type AdhocQueryRegionalDetails struct {

	// Region name
	Region *string `mandatory:"true" json:"region"`

	// adhoc query status of the region
	RegionalStatus AdhocQueryStatusEnum `mandatory:"false" json:"regionalStatus,omitempty"`

	// error message to show if adhoc query fails in a region
	RegionalError *string `mandatory:"false" json:"regionalError"`

	// Expected number of instances on which query should run
	ExpectedCount *string `mandatory:"false" json:"expectedCount"`

	// Number of instances on which query failed
	FailedCount *string `mandatory:"false" json:"failedCount"`

	// Number of instances on which query succeeded
	SucceededCount *string `mandatory:"false" json:"succeededCount"`

	// Number of instances on which query expired
	ExpiredCount *string `mandatory:"false" json:"expiredCount"`
}

func (m AdhocQueryRegionalDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AdhocQueryRegionalDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAdhocQueryStatusEnum(string(m.RegionalStatus)); !ok && m.RegionalStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RegionalStatus: %s. Supported values are: %s.", m.RegionalStatus, strings.Join(GetAdhocQueryStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
