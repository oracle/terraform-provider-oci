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

// InternalVnicAttachment Details of a service VNIC attachment.
type InternalVnicAttachment struct {

	// The OCID of the compartment containing the VNIC attachment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the VNIC.
	Id *string `mandatory:"true" json:"id"`

	// The current state of a VNIC attachment.
	LifecycleState InternalVnicAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The substrate or anycast IP address of the VNICaaS fleet that the VNIC is attached to.
	SubstrateIp *string `mandatory:"true" json:"substrateIp"`

	// The date and time the VNIC attachment was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A user-friendly name. Does not have to be unique.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The slot number of the VNIC.
	SlotId *int `mandatory:"false" json:"slotId"`

	// Shape of VNIC that is used to allocate resource in the data plane.
	VnicShape InternalVnicAttachmentVnicShapeEnum `mandatory:"false" json:"vnicShape,omitempty"`
}

func (m InternalVnicAttachment) String() string {
	return common.PointerString(m)
}

// InternalVnicAttachmentLifecycleStateEnum Enum with underlying type: string
type InternalVnicAttachmentLifecycleStateEnum string

// Set of constants representing the allowable values for InternalVnicAttachmentLifecycleStateEnum
const (
	InternalVnicAttachmentLifecycleStateAttaching InternalVnicAttachmentLifecycleStateEnum = "ATTACHING"
	InternalVnicAttachmentLifecycleStateAttached  InternalVnicAttachmentLifecycleStateEnum = "ATTACHED"
	InternalVnicAttachmentLifecycleStateDetaching InternalVnicAttachmentLifecycleStateEnum = "DETACHING"
	InternalVnicAttachmentLifecycleStateDetached  InternalVnicAttachmentLifecycleStateEnum = "DETACHED"
)

var mappingInternalVnicAttachmentLifecycleState = map[string]InternalVnicAttachmentLifecycleStateEnum{
	"ATTACHING": InternalVnicAttachmentLifecycleStateAttaching,
	"ATTACHED":  InternalVnicAttachmentLifecycleStateAttached,
	"DETACHING": InternalVnicAttachmentLifecycleStateDetaching,
	"DETACHED":  InternalVnicAttachmentLifecycleStateDetached,
}

// GetInternalVnicAttachmentLifecycleStateEnumValues Enumerates the set of values for InternalVnicAttachmentLifecycleStateEnum
func GetInternalVnicAttachmentLifecycleStateEnumValues() []InternalVnicAttachmentLifecycleStateEnum {
	values := make([]InternalVnicAttachmentLifecycleStateEnum, 0)
	for _, v := range mappingInternalVnicAttachmentLifecycleState {
		values = append(values, v)
	}
	return values
}

// InternalVnicAttachmentVnicShapeEnum Enum with underlying type: string
type InternalVnicAttachmentVnicShapeEnum string

// Set of constants representing the allowable values for InternalVnicAttachmentVnicShapeEnum
const (
	InternalVnicAttachmentVnicShapeFixed0200 InternalVnicAttachmentVnicShapeEnum = "VNICAAS_FIXED0200"
	InternalVnicAttachmentVnicShapeFixed0400 InternalVnicAttachmentVnicShapeEnum = "VNICAAS_FIXED0400"
)

var mappingInternalVnicAttachmentVnicShape = map[string]InternalVnicAttachmentVnicShapeEnum{
	"VNICAAS_FIXED0200": InternalVnicAttachmentVnicShapeFixed0200,
	"VNICAAS_FIXED0400": InternalVnicAttachmentVnicShapeFixed0400,
}

// GetInternalVnicAttachmentVnicShapeEnumValues Enumerates the set of values for InternalVnicAttachmentVnicShapeEnum
func GetInternalVnicAttachmentVnicShapeEnumValues() []InternalVnicAttachmentVnicShapeEnum {
	values := make([]InternalVnicAttachmentVnicShapeEnum, 0)
	for _, v := range mappingInternalVnicAttachmentVnicShape {
		values = append(values, v)
	}
	return values
}
