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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RecommendationCount) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingImportanceEnum(string(m.Importance)); !ok && m.Importance != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Importance: %s. Supported values are: %s.", m.Importance, strings.Join(GetImportanceEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
