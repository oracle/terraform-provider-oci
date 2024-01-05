// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// Use the Cloud Advisor API to find potential inefficiencies in your tenancy and address them.
// Cloud Advisor can help you save money, improve performance, strengthen system resilience, and improve security.
// For more information, see Cloud Advisor (https://docs.cloud.oracle.com/Content/CloudAdvisor/Concepts/cloudadvisoroverview.htm).
//

package optimizer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ProfileLevelSummary The metadata associated with the profile level summary.
type ProfileLevelSummary struct {

	// A unique name for the profile level.
	Name *string `mandatory:"true" json:"name"`

	// The name of the recommendation this profile level applies to.
	RecommendationName *string `mandatory:"true" json:"recommendationName"`

	// The metrics that will be evaluated by profiles using this profile level.
	Metrics []EvaluatedMetric `mandatory:"true" json:"metrics"`

	// The default aggregation interval (in days) for profiles using this profile level.
	DefaultInterval *int `mandatory:"true" json:"defaultInterval"`

	// An array of aggregation intervals (in days) allowed for profiles using this profile level.
	ValidIntervals []int `mandatory:"true" json:"validIntervals"`

	// The date and time the category details were created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the category details were last updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`
}

func (m ProfileLevelSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProfileLevelSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
