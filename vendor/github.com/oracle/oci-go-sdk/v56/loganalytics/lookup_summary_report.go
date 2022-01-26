// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// LookupSummaryReport Summary report of lookups in the tenancy.
type LookupSummaryReport struct {

	// The number of user created lookups.
	UserCreatedCount *int `mandatory:"false" json:"userCreatedCount"`

	// The number of oracle defined lookups.
	OracleDefinedCount *int `mandatory:"false" json:"oracleDefinedCount"`

	// The total number of lookups.
	TotalCount *int `mandatory:"false" json:"totalCount"`
}

func (m LookupSummaryReport) String() string {
	return common.PointerString(m)
}
