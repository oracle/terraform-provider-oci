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

// OperatorAction Details of the operator action. Operator actions are a pre-defined set of commands available to the operator on different layers of the infrastructure. Although the groupings may differ depending on the infrastructure layers,
// the groups are designed to enable the operator access to commands to resolve a specific set of issues. The infrastructure layers controlled by the Operator Control include Dom0, CellServer, and Control Plane Server (CPS).
// There are five groups available to the operator. x-obmcs-top-level-enum: '#/definitions/OperatorActionCategories' enum: *OPERATORACTIONCATEGORIES
// The following infrastructure layers are controlled by the operator actions x-obmcs-top-level-enum: '#/definitions/InfrastructureLayers' enum: *INFRASTRUCTURELAYERS
type OperatorAction struct {

	// Unique Oracle assigned identifier for the operator action.
	Id *string `mandatory:"true" json:"id"`

	// Unique name of the operator action.
	Name *string `mandatory:"true" json:"name"`

	// Display Name of the operator action.
	CustomerDisplayName *string `mandatory:"false" json:"customerDisplayName"`

	// Name of the infrastructure layer associated with the operator action.
	Component *string `mandatory:"false" json:"component"`

	// resourceType for which the OperatorAction is applicable
	ResourceType ResourceTypesEnum `mandatory:"false" json:"resourceType,omitempty"`

	// Description of the operator action in terms of associated risk profile, and characteristics of the operating system commands made
	// available to the operator under this operator action.
	Description *string `mandatory:"false" json:"description"`

	// Fine grained properties associated with the operator control.
	Properties []OperatorActionProperties `mandatory:"false" json:"properties"`
}

func (m OperatorAction) String() string {
	return common.PointerString(m)
}
