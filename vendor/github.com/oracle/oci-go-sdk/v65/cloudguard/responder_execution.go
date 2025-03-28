// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResponderExecution Attributes for a responder execution (ResponderExecution resource).
type ResponderExecution struct {

	// The unique identifier of the responder execution
	Id *string `mandatory:"true" json:"id"`

	// Responder rule ID for the responder execution
	ResponderRuleId *string `mandatory:"true" json:"responderRuleId"`

	// Responder rule type for the responder execution
	ResponderRuleType ResponderTypeEnum `mandatory:"true" json:"responderRuleType"`

	// Responder rule name for the responder execution
	ResponderRuleName *string `mandatory:"true" json:"responderRuleName"`

	// Problem ID associated with the responder execution
	ProblemId *string `mandatory:"true" json:"problemId"`

	// Region where the problem is found
	Region *string `mandatory:"true" json:"region"`

	// Target ID of the problem for the responder execution
	TargetId *string `mandatory:"true" json:"targetId"`

	// Compartment OCID of the responder execution for the problem
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Resource type of the problem for the responder execution
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// Resource name of the problem for the responder execution.
	ResourceName *string `mandatory:"true" json:"resourceName"`

	// The date and time the responder execution was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Current execution status of the responder
	ResponderExecutionStatus ResponderExecutionStatesEnum `mandatory:"true" json:"responderExecutionStatus"`

	// Execution mode of the responder
	ResponderExecutionMode ResponderExecutionModesEnum `mandatory:"true" json:"responderExecutionMode"`

	// The date and time the responder execution was updated. Format defined by RFC3339.
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`

	// Message about the responder execution
	Message *string `mandatory:"false" json:"message"`

	ResponderRuleExecutionDetails *ResponderRuleExecutionDetails `mandatory:"false" json:"responderRuleExecutionDetails"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`
}

func (m ResponderExecution) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResponderExecution) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResponderTypeEnum(string(m.ResponderRuleType)); !ok && m.ResponderRuleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResponderRuleType: %s. Supported values are: %s.", m.ResponderRuleType, strings.Join(GetResponderTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingResponderExecutionStatesEnum(string(m.ResponderExecutionStatus)); !ok && m.ResponderExecutionStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResponderExecutionStatus: %s. Supported values are: %s.", m.ResponderExecutionStatus, strings.Join(GetResponderExecutionStatesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingResponderExecutionModesEnum(string(m.ResponderExecutionMode)); !ok && m.ResponderExecutionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResponderExecutionMode: %s. Supported values are: %s.", m.ResponderExecutionMode, strings.Join(GetResponderExecutionModesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
