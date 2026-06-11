// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database Autonomous Recovery Service API
//
// Use Oracle Database Autonomous Recovery Service API to manage Protected Databases.
//

package recovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RequestSummarizedProtectedDatabaseAnalyticsDetails Request body with equality filters to scope analytics prior to aggregation. The filters applied after compartment scoping.
// Filtering can be done on any of the `groupBy` properties like `health`, `isRedoLogsEnabled`, `protectionPolicyId`, `lifecycleState`.
// For example (any 1 or more filters can be used for filtering):
// ```
//
//	{
//	  "filters": {
//	    "lifecycleState": ["ACTIVE", "UPDATING"],
//	    "health": ["PROTECTED"],
//	    "isRedoLogsEnabled": ["true"],
//	    "protectionPolicyId": ["ocid1.policy.silver.."],
//	  }
//	}
//
// ```
type RequestSummarizedProtectedDatabaseAnalyticsDetails struct {

	// Equality filters applied before aggregation.
	// Rules
	//   - Filters on different fields are treated as AND
	//   - Multiple filters on the same field are treated as OR
	// For example if filters is :
	// ```
	// "filters": {
	//   "lifecycleState": ["ACTIVE", "UPDATING"],
	//   "health": ["PROTECTED"]
	// }
	// ```
	// This Matches Protected Databases where `lifecycleState` is `ACTIVE` OR `UPDATING`, AND `health` is `PROTECTED`.
	Filters map[string][]string `mandatory:"false" json:"filters"`
}

func (m RequestSummarizedProtectedDatabaseAnalyticsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RequestSummarizedProtectedDatabaseAnalyticsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
