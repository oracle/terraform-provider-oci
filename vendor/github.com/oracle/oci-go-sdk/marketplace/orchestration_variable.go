// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
)

// OrchestrationVariable The model of a variable for an orchestration resource.
type OrchestrationVariable struct {

	// The name of the variable.
	Name *string `mandatory:"false" json:"name"`

	// The variable's default value.
	DefaultValue *string `mandatory:"false" json:"defaultValue"`

	// A description of the variable.
	Description *string `mandatory:"false" json:"description"`

	// The data type of the variable.
	DataType OrchestrationVariableTypeEnumEnum `mandatory:"false" json:"dataType,omitempty"`

	// Whether the variable is mandatory.
	IsMandatory *bool `mandatory:"false" json:"isMandatory"`

	// A brief textual description that helps to explain the variable.
	HintMessage *string `mandatory:"false" json:"hintMessage"`
}

func (m OrchestrationVariable) String() string {
	return common.PointerString(m)
}
