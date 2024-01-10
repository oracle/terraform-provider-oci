// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LookupSummaryReport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
