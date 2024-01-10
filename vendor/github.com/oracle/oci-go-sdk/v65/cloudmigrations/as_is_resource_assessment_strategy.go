// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AsIsResourceAssessmentStrategy The 'As-Is' based strategy.
type AsIsResourceAssessmentStrategy struct {

	// The real resource usage is multiplied to this number before making any recommendation.
	AdjustmentMultiplier *float32 `mandatory:"false" json:"adjustmentMultiplier"`

	// The type of resource.
	ResourceType ResourceAssessmentStrategyResourceTypeEnum `mandatory:"true" json:"resourceType"`
}

// GetResourceType returns ResourceType
func (m AsIsResourceAssessmentStrategy) GetResourceType() ResourceAssessmentStrategyResourceTypeEnum {
	return m.ResourceType
}

func (m AsIsResourceAssessmentStrategy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AsIsResourceAssessmentStrategy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingResourceAssessmentStrategyResourceTypeEnum(string(m.ResourceType)); !ok && m.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", m.ResourceType, strings.Join(GetResourceAssessmentStrategyResourceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AsIsResourceAssessmentStrategy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAsIsResourceAssessmentStrategy AsIsResourceAssessmentStrategy
	s := struct {
		DiscriminatorParam string `json:"strategyType"`
		MarshalTypeAsIsResourceAssessmentStrategy
	}{
		"AS_IS",
		(MarshalTypeAsIsResourceAssessmentStrategy)(m),
	}

	return json.Marshal(&s)
}
