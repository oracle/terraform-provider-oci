// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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
