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

// RecommendationSummary The metadata associated with the recommendation summary.
type RecommendationSummary struct {

	// The unique OCID associated with the recommendation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy. The tenancy is the root compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The unique OCID associated with the category.
	CategoryId *string `mandatory:"true" json:"categoryId"`

	// The name assigned to the recommendation.
	Name *string `mandatory:"true" json:"name"`

	// Text describing the recommendation.
	Description *string `mandatory:"true" json:"description"`

	// The level of importance assigned to the recommendation.
	Importance ImportanceEnum `mandatory:"true" json:"importance"`

	// An array of `ResourceCount` objects grouped by the status of the resource actions.
	ResourceCounts []ResourceCount `mandatory:"true" json:"resourceCounts"`

	// The recommendation's current state.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The estimated cost savings, in dollars, for the recommendation.
	EstimatedCostSaving *float64 `mandatory:"true" json:"estimatedCostSaving"`

	// The current status of the recommendation.
	Status StatusEnum `mandatory:"true" json:"status"`

	// The date and time that the recommendation entered its current status. The format is defined by RFC3339.
	// For example, "The status of the recommendation changed from `pending` to `current(ignored)` on this date and time."
	TimeStatusBegin *common.SDKTime `mandatory:"true" json:"timeStatusBegin"`

	// The date and time the current status will change. The format is defined by RFC3339.
	// For example, "The current `postponed` status of the recommendation will end and change to `pending` on this
	// date and time."
	TimeStatusEnd *common.SDKTime `mandatory:"false" json:"timeStatusEnd"`

	// The date and time the recommendation details were created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the recommendation details were last updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	SupportedLevels *SupportedLevels `mandatory:"false" json:"supportedLevels"`

	// Additional metadata key/value pairs for the recommendation summary.
	// For example:
	// `{"EstimatedSaving": "200"}`
	ExtendedMetadata map[string]string `mandatory:"false" json:"extendedMetadata"`
}

func (m RecommendationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RecommendationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingImportanceEnum(string(m.Importance)); !ok && m.Importance != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Importance: %s. Supported values are: %s.", m.Importance, strings.Join(GetImportanceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
