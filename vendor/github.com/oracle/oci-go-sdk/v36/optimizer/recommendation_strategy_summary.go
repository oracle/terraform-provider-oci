// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// APIs for managing Cloud Advisor. Cloud Advisor provides recommendations that help you maximize cost savings and improve the security posture of your tenancy.
//

package optimizer

import (
	"github.com/oracle/oci-go-sdk/v36/common"
)

// RecommendationStrategySummary The metadata associated with the recommendation strategy.
type RecommendationStrategySummary struct {

	// The display name of the recommendation.
	Name *string `mandatory:"true" json:"name"`

	// The list of strategies used.
	Strategies []Strategy `mandatory:"true" json:"strategies"`
}

func (m RecommendationStrategySummary) String() string {
	return common.PointerString(m)
}
