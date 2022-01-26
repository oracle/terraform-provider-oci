// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ResourceActionSummary The metadata associated with the resource action summary.
type ResourceActionSummary struct {

	// The unique OCID associated with the resource action.
	Id *string `mandatory:"true" json:"id"`

	// The unique OCID associated with the category.
	CategoryId *string `mandatory:"true" json:"categoryId"`

	// The unique OCID associated with the recommendation.
	RecommendationId *string `mandatory:"true" json:"recommendationId"`

	// The unique OCID associated with the resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The name assigned to the resource.
	Name *string `mandatory:"true" json:"name"`

	// The kind of resource.
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name associated with the compartment.
	CompartmentName *string `mandatory:"true" json:"compartmentName"`

	Action *Action `mandatory:"true" json:"action"`

	// The resource action's current state.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The estimated cost savings, in dollars, for the resource action.
	EstimatedCostSaving *float64 `mandatory:"true" json:"estimatedCostSaving"`

	// The current status of the resource action.
	Status StatusEnum `mandatory:"true" json:"status"`

	// The date and time that the resource action entered its current status. The format is defined by RFC3339.
	// For example, "The status of the resource action changed from `pending` to `current(ignored)` on this date and time."
	TimeStatusBegin *common.SDKTime `mandatory:"true" json:"timeStatusBegin"`

	// The date and time the current status will change. The format is defined by RFC3339.
	// For example, "The current `postponed` status of the resource action will end and change to `pending` on this
	// date and time."
	TimeStatusEnd *common.SDKTime `mandatory:"false" json:"timeStatusEnd"`

	// Custom metadata key/value pairs for the resource action.
	//  **Metadata Example**
	//       "metadata" : {
	//          "cpuRecommendedShape": "VM.Standard1.1",
	//          "computeMemoryUtilization": "26.05734124418388",
	//          "currentShape": "VM.Standard1.2",
	//          "instanceRecommendedShape": "VM.Standard1.1",
	//          "computeCpuUtilization": "7.930035319720132",
	//          "memoryRecommendedShape": "None"
	//       }
	Metadata map[string]string `mandatory:"false" json:"metadata"`

	// Additional metadata key/value pairs that you provide.
	// They serve the same purpose and functionality as fields in the `metadata` object.
	// They are distinguished from `metadata` fields in that these can be nested JSON objects (whereas `metadata` fields are string/string maps only).
	// For example:
	// `{"CurrentShape": {"name":"VM.Standard2.16"}, "RecommendedShape": {"name":"VM.Standard2.8"}}`
	ExtendedMetadata map[string]interface{} `mandatory:"false" json:"extendedMetadata"`

	// The date and time the resource action details were created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the resource action details were last updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m ResourceActionSummary) String() string {
	return common.PointerString(m)
}
