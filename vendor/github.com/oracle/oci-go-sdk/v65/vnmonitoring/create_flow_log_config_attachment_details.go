// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateFlowLogConfigAttachmentDetails The representation of CreateFlowLogConfigAttachmentDetails
type CreateFlowLogConfigAttachmentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource to attach the
	// flow log configuration to. Attaching the configuration enables flow logs for the resource.
	TargetEntityId *string `mandatory:"true" json:"targetEntityId"`

	// The type of resource to attach the flow log configuration to.
	TargetEntityType CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum `mandatory:"true" json:"targetEntityType"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the flow log configuration to attach.
	FlowLogConfigId *string `mandatory:"true" json:"flowLogConfigId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The capture filter's Oracle ID (OCID).
	CaptureFilterId *string `mandatory:"false" json:"captureFilterId"`
}

func (m CreateFlowLogConfigAttachmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateFlowLogConfigAttachmentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum(string(m.TargetEntityType)); !ok && m.TargetEntityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetEntityType: %s. Supported values are: %s.", m.TargetEntityType, strings.Join(GetCreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum Enum with underlying type: string
type CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum string

// Set of constants representing the allowable values for CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum
const (
	CreateFlowLogConfigAttachmentDetailsTargetEntityTypeSubnet CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum = "SUBNET"
	CreateFlowLogConfigAttachmentDetailsTargetEntityTypeVcn    CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum = "VCN"
	CreateFlowLogConfigAttachmentDetailsTargetEntityTypeVnic   CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum = "VNIC"
)

var mappingCreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum = map[string]CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum{
	"SUBNET": CreateFlowLogConfigAttachmentDetailsTargetEntityTypeSubnet,
	"VCN":    CreateFlowLogConfigAttachmentDetailsTargetEntityTypeVcn,
	"VNIC":   CreateFlowLogConfigAttachmentDetailsTargetEntityTypeVnic,
}

var mappingCreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnumLowerCase = map[string]CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum{
	"subnet": CreateFlowLogConfigAttachmentDetailsTargetEntityTypeSubnet,
	"vcn":    CreateFlowLogConfigAttachmentDetailsTargetEntityTypeVcn,
	"vnic":   CreateFlowLogConfigAttachmentDetailsTargetEntityTypeVnic,
}

// GetCreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnumValues Enumerates the set of values for CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum
func GetCreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnumValues() []CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum {
	values := make([]CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum, 0)
	for _, v := range mappingCreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnumStringValues Enumerates the set of values in String for CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum
func GetCreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnumStringValues() []string {
	return []string{
		"SUBNET",
		"VCN",
		"VNIC",
	}
}

// GetMappingCreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum(val string) (CreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnum, bool) {
	enum, ok := mappingCreateFlowLogConfigAttachmentDetailsTargetEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
