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

// ProcessRecommendationDetails Details of recommendation to be processed.
type ProcessRecommendationDetails struct {

	// Unique identifier of the recommendation.
	RecommendationKey *string `mandatory:"true" json:"recommendationKey"`

	// The status of a recommendation.
	RecommendationStatus RecommendationStatusEnum `mandatory:"true" json:"recommendationStatus"`

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

func (m ProcessRecommendationDetails) String() string {
	return common.PointerString(m)
}
