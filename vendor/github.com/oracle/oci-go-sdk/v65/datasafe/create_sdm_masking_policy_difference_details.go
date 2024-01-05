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

// CreateSdmMaskingPolicyDifferenceDetails Details to create a new SDM masking policy difference.
type CreateSdmMaskingPolicyDifferenceDetails struct {

	// The OCID of the masking policy. Note that if the masking policy is not associated with an SDM, CreateSdmMaskingPolicyDifference
	// operation won't be allowed.
	MaskingPolicyId *string `mandatory:"true" json:"maskingPolicyId"`

	// The OCID of the compartment where the SDM masking policy difference resource should be created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The type of the SDM masking policy difference. It defines the difference scope.
	// NEW identifies new sensitive columns in the sensitive data model that are not in the masking policy.
	// DELETED identifies columns that are present in the masking policy but have been deleted from the sensitive data model.
	// MODIFIED identifies columns that are present in the sensitive data model as well as the masking policy but some of their attributes have been modified.
	// ALL covers all the above three scenarios and reports new, deleted and modified columns.
	DifferenceType SdmMaskingPolicyDifferenceDifferenceTypeEnum `mandatory:"false" json:"differenceType,omitempty"`

	// A user-friendly name for the SDM masking policy difference. Does not have to be unique, and it is changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateSdmMaskingPolicyDifferenceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSdmMaskingPolicyDifferenceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSdmMaskingPolicyDifferenceDifferenceTypeEnum(string(m.DifferenceType)); !ok && m.DifferenceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DifferenceType: %s. Supported values are: %s.", m.DifferenceType, strings.Join(GetSdmMaskingPolicyDifferenceDifferenceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
