// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperatorAccessControl API
//
// Operator Access Control enables you to control the time duration and the actions an Oracle operator can perform on your Exadata Cloud@Customer infrastructure.
// Using logging service, you can view a near real-time audit report of all actions performed by an Oracle operator.
// Use the table of contents and search tool to explore the OperatorAccessControl API.
//

package operatoraccesscontrol

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// OperatorActionSummary Details of the operator action. Operator actions are pre-defined set of commands available to the operator on different layers of the infrastructure.
type OperatorActionSummary struct {

	// Unique identifier assigned by Oracle to an operator action.
	Id *string `mandatory:"true" json:"id"`

	// Name of the operator action.
	Name *string `mandatory:"true" json:"name"`

	// Name of the component for which the operator action is applicable.
	Component *string `mandatory:"false" json:"component"`

	// compartmentId for which the OperatorAction is applicable
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// resourceType for which the OperatorAction is applicable
	ResourceType ResourceTypesEnum `mandatory:"false" json:"resourceType,omitempty"`

	// The current lifecycle state of the operator action.
	LifecycleState OperatorActionLifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Description of the operator action in terms of associated risk profile, and characteristics of the operating system commands made
	// available to the operator under this operator action.
	Description *string `mandatory:"false" json:"description"`
}

func (m OperatorActionSummary) String() string {
	return common.PointerString(m)
}
