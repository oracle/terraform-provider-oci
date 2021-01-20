// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// APIs for managing Cloud Advisor. Cloud Advisor provides recommendations that help you maximize cost savings and improve the security posture of your tenancy.
//

package optimizer

import (
	"github.com/oracle/oci-go-sdk/v33/common"
)

// BulkApplyRecommendationsDetails Details about bulk recommendation actions.
type BulkApplyRecommendationsDetails struct {

	// The unique OCIDs of the resource actions that recommendations are applied to.
	ResourceActionIds []string `mandatory:"true" json:"resourceActionIds"`

	// The current status of the recommendation.
	Status StatusEnum `mandatory:"true" json:"status"`

	// The date and time the current status will change. The format is defined by RFC3339.
	// For example, "The current `postponed` status of the resource action will end and change to `pending` on this
	// date and time."
	TimeStatusEnd *common.SDKTime `mandatory:"false" json:"timeStatusEnd"`
}

func (m BulkApplyRecommendationsDetails) String() string {
	return common.PointerString(m)
}
