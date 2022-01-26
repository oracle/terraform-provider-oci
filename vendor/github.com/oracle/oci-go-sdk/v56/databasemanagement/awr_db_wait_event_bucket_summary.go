// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// AwrDbWaitEventBucketSummary A summary of the AWR wait event bucket and waits percentage.
type AwrDbWaitEventBucketSummary struct {

	// The name of the wait event frequency category. Normally, it is the upper range of the waits within the AWR wait event bucket.
	Category *string `mandatory:"true" json:"category"`

	// The percentage of waits in a wait event bucket over the total waits of the database.
	Percentage *float64 `mandatory:"true" json:"percentage"`
}

func (m AwrDbWaitEventBucketSummary) String() string {
	return common.PointerString(m)
}
