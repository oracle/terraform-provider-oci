// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// InternalDrgAttachment A link between a DRG and VCN. For more information, see
// Overview of the Networking Service (https://docs.cloud.oracle.com/Content/Network/Concepts/overview.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you
// supply string values using the API.
type InternalDrgAttachment struct {

	// The OCID of the compartment containing the DRG attachment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the DRG.
	DrgId *string `mandatory:"true" json:"drgId"`

	// The DRG attachment's Oracle ID (OCID).
	Id *string `mandatory:"true" json:"id"`

	// The DRG attachment's current state.
	LifecycleState InternalDrgAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the VCN.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// NextHop target's MPLS label.
	MplsLabel *string `mandatory:"true" json:"mplsLabel"`

	// The string in the form ASN:mplsLabel.
	RouteTarget *string `mandatory:"true" json:"routeTarget"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the route table the DRG attachment is using.
	RouteTableId *string `mandatory:"false" json:"routeTableId"`

	// The date and time the DRG attachment was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m InternalDrgAttachment) String() string {
	return common.PointerString(m)
}

// InternalDrgAttachmentLifecycleStateEnum Enum with underlying type: string
type InternalDrgAttachmentLifecycleStateEnum string

// Set of constants representing the allowable values for InternalDrgAttachmentLifecycleStateEnum
const (
	InternalDrgAttachmentLifecycleStateAttaching InternalDrgAttachmentLifecycleStateEnum = "ATTACHING"
	InternalDrgAttachmentLifecycleStateAttached  InternalDrgAttachmentLifecycleStateEnum = "ATTACHED"
	InternalDrgAttachmentLifecycleStateDetaching InternalDrgAttachmentLifecycleStateEnum = "DETACHING"
	InternalDrgAttachmentLifecycleStateDetached  InternalDrgAttachmentLifecycleStateEnum = "DETACHED"
)

var mappingInternalDrgAttachmentLifecycleState = map[string]InternalDrgAttachmentLifecycleStateEnum{
	"ATTACHING": InternalDrgAttachmentLifecycleStateAttaching,
	"ATTACHED":  InternalDrgAttachmentLifecycleStateAttached,
	"DETACHING": InternalDrgAttachmentLifecycleStateDetaching,
	"DETACHED":  InternalDrgAttachmentLifecycleStateDetached,
}

// GetInternalDrgAttachmentLifecycleStateEnumValues Enumerates the set of values for InternalDrgAttachmentLifecycleStateEnum
func GetInternalDrgAttachmentLifecycleStateEnumValues() []InternalDrgAttachmentLifecycleStateEnum {
	values := make([]InternalDrgAttachmentLifecycleStateEnum, 0)
	for _, v := range mappingInternalDrgAttachmentLifecycleState {
		values = append(values, v)
	}
	return values
}
