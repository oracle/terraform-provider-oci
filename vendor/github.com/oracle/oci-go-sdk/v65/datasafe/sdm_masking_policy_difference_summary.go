// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SdmMaskingPolicyDifferenceSummary Summary of a SDM masking policy difference.
type SdmMaskingPolicyDifferenceSummary struct {

	// The OCID of the SDM masking policy difference.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment to contain the SDM masking policy difference.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the SDM masking policy difference.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time the SDM masking policy difference was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the SDM masking policy difference creation started, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreationStarted *common.SDKTime `mandatory:"true" json:"timeCreationStarted"`

	// The OCID of the sensitive data model associated with the SDM masking policy difference.
	SensitiveDataModelId *string `mandatory:"true" json:"sensitiveDataModelId"`

	// The OCID of the masking policy associated with the SDM masking policy difference.
	MaskingPolicyId *string `mandatory:"true" json:"maskingPolicyId"`

	// The current state of the SDM masking policy difference.
	LifecycleState SdmMaskingPolicyDifferenceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The type of difference.
	DifferenceType SdmMaskingPolicyDifferenceDifferenceTypeEnum `mandatory:"true" json:"differenceType"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m SdmMaskingPolicyDifferenceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SdmMaskingPolicyDifferenceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSdmMaskingPolicyDifferenceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSdmMaskingPolicyDifferenceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSdmMaskingPolicyDifferenceDifferenceTypeEnum(string(m.DifferenceType)); !ok && m.DifferenceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DifferenceType: %s. Supported values are: %s.", m.DifferenceType, strings.Join(GetSdmMaskingPolicyDifferenceDifferenceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
