// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UserAggregation The user aggregation provides information about the overall security state of database users.
// For example, it states how many users have the DBA role and how many users are in the critical category.
type UserAggregation struct {

	// The array of user aggregation data.
	Items []map[string]interface{} `mandatory:"true" json:"items"`
}

func (m UserAggregation) String() string {
	return common.PointerString(m)
}
