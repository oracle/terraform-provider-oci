// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResponderActivitySummary Responder activity summary definition.
type ResponderActivitySummary struct {

	// Unique ID for responder activity
	Id *string `mandatory:"true" json:"id"`

	// Unique ID of problem associated with responder activity
	ProblemId *string `mandatory:"true" json:"problemId"`

	// Unique ID of the responder rule associated with the problem
	ResponderRuleId *string `mandatory:"true" json:"responderRuleId"`

	// Responder rule type for performing the operation
	ResponderType ResponderTypeEnum `mandatory:"true" json:"responderType"`

	// Responder rule name
	ResponderRuleName *string `mandatory:"true" json:"responderRuleName"`

	// Responder activity type
	ResponderActivityType ResponderActivityTypeEnum `mandatory:"true" json:"responderActivityType"`

	// Responder execution status
	ResponderExecutionStatus ResponderExecutionStatesEnum `mandatory:"true" json:"responderExecutionStatus"`

	// Responder activity starting time
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Additional message related to this operation
	Message *string `mandatory:"true" json:"message"`
}

func (m ResponderActivitySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResponderActivitySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResponderTypeEnum(string(m.ResponderType)); !ok && m.ResponderType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResponderType: %s. Supported values are: %s.", m.ResponderType, strings.Join(GetResponderTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingResponderActivityTypeEnum(string(m.ResponderActivityType)); !ok && m.ResponderActivityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResponderActivityType: %s. Supported values are: %s.", m.ResponderActivityType, strings.Join(GetResponderActivityTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingResponderExecutionStatesEnum(string(m.ResponderExecutionStatus)); !ok && m.ResponderExecutionStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResponderExecutionStatus: %s. Supported values are: %s.", m.ResponderExecutionStatus, strings.Join(GetResponderExecutionStatesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
