// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateJobDetails Job properties that can be updated.
type UpdateJobDetails struct {

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detailed description of the job.
	Description *string `mandatory:"false" json:"description"`

	// Schedule specified in the cron expression format that has seven fields for second, minute, hour, day-of-month, month, day-of-week, year.
	// It can also include special characters like * for all and ? for any. There are also pre-defined schedules that can be specified using
	// special strings. For example, @hourly will run the job every hour.
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
