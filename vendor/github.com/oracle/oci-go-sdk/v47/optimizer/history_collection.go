// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// APIs for managing Cloud Advisor. Cloud Advisor provides recommendations that help you maximize cost savings and improve the security posture of your tenancy.
//

package optimizer

import (
	"github.com/oracle/oci-go-sdk/v47/common"
)

// HistoryCollection A list containing the recommendation history items that match filter criteria, if any. Results contain `HistorySummary` objects.
type HistoryCollection struct {

	// A collection of history summaries.
	Items []HistorySummary `mandatory:"true" json:"items"`
}

func (m HistoryCollection) String() string {
	return common.PointerString(m)
}
