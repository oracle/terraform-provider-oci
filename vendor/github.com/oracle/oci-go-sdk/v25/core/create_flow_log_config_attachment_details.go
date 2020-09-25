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

// CreateFlowLogConfigAttachmentDetails The representation of CreateFlowLogConfigAttachmentDetails
type CreateFlowLogConfigAttachmentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the resource to attach the
	// flow log configuration to. Attaching the configuration enables flow logs for the resource.
	TargetEntityId *string `mandatory:"true" json:"targetEntityId"`

	// The type of resource to attach the flow log configuration to.
	TargetEntityType CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum `mandatory:"true" json:"targetEntityType"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the flow log configuration to attach.
	FlowLogConfigId *string `mandatory:"true" json:"flowLogConfigId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid
	// entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m CreateFlowLogConfigAttachmentDetails) String() string {
	return common.PointerString(m)
}

// CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum Enum with underlying type: string
type CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum string

// Set of constants representing the allowable values for CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum
const (
	CreateFlowLogConfigAttachmentDetailsTargetEntityTypeSubnet CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum = "SUBNET"
)

var mappingCreateFlowLogConfigAttachmentDetailsTargetEntityType = map[string]CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum{
	"SUBNET": CreateFlowLogConfigAttachmentDetailsTargetEntityTypeSubnet,
}

// GetCreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnumValues Enumerates the set of values for CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum
func GetCreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnumValues() []CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum {
	values := make([]CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum, 0)
	for _, v := range mappingCreateFlowLogConfigAttachmentDetailsTargetEntityType {
		values = append(values, v)
	}
	return values
}
