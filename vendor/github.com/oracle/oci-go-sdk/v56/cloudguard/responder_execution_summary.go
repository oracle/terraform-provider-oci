// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ResponderExecutionSummary Summary of the Responder Execution.
type ResponderExecutionSummary struct {

	// The unique identifier of the responder execution
	Id *string `mandatory:"true" json:"id"`

	// Responder Rule id for the responder execution
	ResponderRuleId *string `mandatory:"true" json:"responderRuleId"`

	// Rule Type for the responder execution
	ResponderRuleType ResponderTypeEnum `mandatory:"true" json:"responderRuleType"`

	// Rule name for the responder execution
	ResponderRuleName *string `mandatory:"true" json:"responderRuleName"`

	// Problem id associated with the responder execution
	ProblemId *string `mandatory:"true" json:"problemId"`

	// Problem name associated with the responder execution
	ProblemName *string `mandatory:"true" json:"problemName"`

	// Region where the problem is found
	Region *string `mandatory:"true" json:"region"`

	// Target Id of the problem for the responder execution
	TargetId *string `mandatory:"true" json:"targetId"`

	// compartment id of the problem for the responder execution
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// resource type of the problem for the responder execution
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// resource name of the problem for the responder execution. TODO-DOC link to resource definition doc
	ResourceName *string `mandatory:"true" json:"resourceName"`

	// The date and time the responder execution was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// current execution status of the responder
	ResponderExecutionStatus ResponderExecutionStatesEnum `mandatory:"true" json:"responderExecutionStatus"`

	// possible type of responder execution modes
	ResponderExecutionMode ResponderExecutionModesEnum `mandatory:"true" json:"responderExecutionMode"`

	// The date and time the responder execution was updated. Format defined by RFC3339.
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`

	// Message about the responder execution.
	Message *string `mandatory:"false" json:"message"`

	ResponderRuleExecutionDetails *ResponderRuleExecutionDetails `mandatory:"false" json:"responderRuleExecutionDetails"`
}

func (m ResponderExecutionSummary) String() string {
	return common.PointerString(m)
}
