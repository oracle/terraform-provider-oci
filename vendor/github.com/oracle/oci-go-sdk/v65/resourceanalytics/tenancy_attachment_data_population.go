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

// TenancyAttachmentDataPopulation Data population status for a monitored region in the tenancy.
type TenancyAttachmentDataPopulation struct {

	// The overall status of the data population from the monitored region of the tenancy.
	Status TenancyAttachmentDataPopulationStatusEnum `mandatory:"true" json:"status"`

	// The date and time the data population task was started, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The number of data population tasks currently in progress.
	InProgressCount *int `mandatory:"true" json:"inProgressCount"`

	// The number of data population tasks that have succeeded.
	SucceededCount *int `mandatory:"true" json:"succeededCount"`

	// The total number of data population tasks.
	TotalCount *int `mandatory:"true" json:"totalCount"`

	// The date and time the data population task completed, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`
}

func (m TenancyAttachmentDataPopulation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TenancyAttachmentDataPopulation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTenancyAttachmentDataPopulationStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetTenancyAttachmentDataPopulationStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
