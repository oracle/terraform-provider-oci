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

// SqlBucket Sql bucket type object.
type SqlBucket struct {

	// Collection timestamp
	// Example: `"2020-03-31T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// Unique SQL_ID for a SQL Statement.
	SqlIdentifier *string `mandatory:"true" json:"sqlIdentifier"`

	// Plan hash value for the SQL Execution Plan
	PlanHash *int64 `mandatory:"true" json:"planHash"`

	// SQL Bucket ID, examples <= 3 secs, 3-10 secs, 10-60 secs, 1-5 min, > 5 min
	// Example: `"<= 3 secs"`
	BucketId *string `mandatory:"true" json:"bucketId"`

	// Version
	// Example: `1`
	Version *float32 `mandatory:"false" json:"version"`

	// Ops Insights internal representation of the database type.
	DatabaseType *string `mandatory:"false" json:"databaseType"`

	// Total number of executions
	// Example: `60`
	ExecutionsCount *int `mandatory:"false" json:"executionsCount"`

	// Total CPU time
	// Example: `1046`
	CpuTimeInSec *float32 `mandatory:"false" json:"cpuTimeInSec"`

	// Total IO time
	// Example: `5810`
	IoTimeInSec *float32 `mandatory:"false" json:"ioTimeInSec"`

	// Total other wait time
	// Example: `24061`
	OtherWaitTimeInSec *float32 `mandatory:"false" json:"otherWaitTimeInSec"`

	// Total time
	// Example: `30917`
	TotalTimeInSec *float32 `mandatory:"false" json:"totalTimeInSec"`
}

func (m SqlBucket) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlBucket) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
