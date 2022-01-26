// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// RecommendationDetails Details of a recommendation.
type RecommendationDetails struct {

	// Unique identifier of the recommendation.
	RecommendationKey *string `mandatory:"true" json:"recommendationKey"`

	// Type of recommendation.
	RecommendationType RecommendationTypeEnum `mandatory:"true" json:"recommendationType"`

	// Status of a recommendation.
	RecommendationStatus RecommendationStatusEnum `mandatory:"true" json:"recommendationStatus"`

	// Level of confidence, on a scale between 0 and 1, that the recommendation is applicable.
	ConfidenceScore *float32 `mandatory:"false" json:"confidenceScore"`

	// Unique identifier of the source object; the one for which a recommendation is made.
	SourceObjectKey *string `mandatory:"false" json:"sourceObjectKey"`

	// Name of the source object; the one for which a recommendation is made.
	SourceObjectName *string `mandatory:"false" json:"sourceObjectName"`

	// Type of the source object; the one for which a recommendation is made.
	SourceObjectType RecommendationResourceTypeEnum `mandatory:"false" json:"sourceObjectType,omitempty"`

	// Unique identifier of the target object; the one which has been recommended.
	TargetObjectKey *string `mandatory:"false" json:"targetObjectKey"`

	// Name of the target object; the one which has been recommended.
	TargetObjectName *string `mandatory:"false" json:"targetObjectName"`

	// Type of the target object; the one which has been recommended.
	TargetObjectType RecommendationResourceTypeEnum `mandatory:"false" json:"targetObjectType,omitempty"`

	// A map of maps that contains additional properties which are specific to the associated objects.
	// Each associated object defines it's set of required and optional properties.
	// Example: `{
	//             "DataEntity": {
	//               "parentId": "entityId"
	//             },
	//             "Term": {
	//               "parentId": "glossaryId"
	//             }
	//           }`
	Properties map[string]map[string]string `mandatory:"false" json:"properties"`
}

func (m RecommendationDetails) String() string {
	return common.PointerString(m)
}
