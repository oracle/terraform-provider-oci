// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CreateAlertPolicyDetails The details used to create a new alert policy.
type CreateAlertPolicyDetails struct {

	// Indicates the Data Safe feature the alert policy belongs to
	AlertPolicyType AlertPolicyTypeEnum `mandatory:"true" json:"alertPolicyType"`

	// Severity level of the alert raised by this policy.
	Severity AlertSeverityEnum `mandatory:"true" json:"severity"`

	// The OCID of the compartment where you want to create the alert policy.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the alert policy. The name does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the alert policy.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateAlertPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAlertPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAlertPolicyTypeEnum(string(m.AlertPolicyType)); !ok && m.AlertPolicyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AlertPolicyType: %s. Supported values are: %s.", m.AlertPolicyType, strings.Join(GetAlertPolicyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAlertSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetAlertSeverityEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
