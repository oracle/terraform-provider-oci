// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateJobDetails Job properties that can be updated.
type UpdateJobDetails struct {

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detailed description of the job.
	Description *string `mandatory:"false" json:"description"`

	// Interval on which the job will be run. Value is specified as a cron-supported time specification "nickname".
	// The following subset of those is supported: @monthly, @weekly, @daily, @hourly.
	// For metastore sync, an additional option @default is supported, which will schedule jobs at a more granular frequency.
	ScheduleCronExpression *string `mandatory:"false" json:"scheduleCronExpression"`

	// Date that the schedule should be operational. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeScheduleBegin *common.SDKTime `mandatory:"false" json:"timeScheduleBegin"`

	// Date that the schedule should end from being operational. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeScheduleEnd *common.SDKTime `mandatory:"false" json:"timeScheduleEnd"`

	// The key of the connection resource that is used for the harvest by this job.
	ConnectionKey *string `mandatory:"false" json:"connectionKey"`
}

func (m UpdateJobDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
