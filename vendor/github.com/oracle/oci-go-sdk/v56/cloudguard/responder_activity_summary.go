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

// ResponderActivitySummary Responder Activity summary Definition.
type ResponderActivitySummary struct {

	// Unique id for Responder activity.
	Id *string `mandatory:"true" json:"id"`

	// problemId for which Responder activity is associated to.
	ProblemId *string `mandatory:"true" json:"problemId"`

	// Id of the responder rule for the problem
	ResponderRuleId *string `mandatory:"true" json:"responderRuleId"`

	// responder rule type for performing the operation
	ResponderType ResponderTypeEnum `mandatory:"true" json:"responderType"`

	// responder rule name
	ResponderRuleName *string `mandatory:"true" json:"responderRuleName"`

	// Responder activity types
	ResponderActivityType ResponderActivityTypeEnum `mandatory:"true" json:"responderActivityType"`

	// the responder execution status
	ResponderExecutionStatus ResponderExecutionStatesEnum `mandatory:"true" json:"responderExecutionStatus"`

	// responder activity starting time
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// additional message related to this operation
	Message *string `mandatory:"true" json:"message"`
}

func (m ResponderActivitySummary) String() string {
	return common.PointerString(m)
}
