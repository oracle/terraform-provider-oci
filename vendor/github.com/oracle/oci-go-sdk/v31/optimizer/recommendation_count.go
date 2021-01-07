// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// APIs for managing Cloud Advisor. Cloud Advisor provides recommendations that help you maximize cost savings and improve the security posture of your tenancy.
//

package optimizer

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// RecommendationCount The count of recommendations in a category, grouped by importance.
type RecommendationCount struct {

	// The level of importance assigned to the recommendation.
	Importance ImportanceEnum `mandatory:"true" json:"importance"`

	// The count of recommendations.
	Count *int `mandatory:"true" json:"count"`
}

func (m RecommendationCount) String() string {
	return common.PointerString(m)
}
