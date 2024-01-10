// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Process Automation
//
// Process Automation helps you to rapidly design, automate, and manage business processes in the cloud. With the Process Automation design-time (Designer) and the runtime (Workspace) environments, you can easily create, develop, manage, test, and monitor process applications and their components.
//

package opa

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AttachmentDetails Description of an attachment for an instance
type AttachmentDetails struct {

	// The role of the target attachment.
	//    * `PARENT` - The target instance is the parent of this attachment.
	//    * `CHILD` - The target instance is the child of this attachment.
	TargetRole AttachmentDetailsTargetRoleEnum `mandatory:"true" json:"targetRole"`

	// * If role == `PARENT`, the attached instance was created by this service instance
	// * If role == `CHILD`, this instance was created from attached instance on behalf of a user
	IsImplicit *bool `mandatory:"true" json:"isImplicit"`

	// The OCID of the target instance (which could be any other OCI PaaS/SaaS resource), to which this instance is attached.
	TargetId *string `mandatory:"true" json:"targetId"`

	// The dataplane instance URL of the attached instance
	TargetInstanceUrl *string `mandatory:"true" json:"targetInstanceUrl"`

	// The type of the target instance, such as "FUSION".
	TargetServiceType *string `mandatory:"true" json:"targetServiceType"`
}

func (m AttachmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AttachmentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAttachmentDetailsTargetRoleEnum(string(m.TargetRole)); !ok && m.TargetRole != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetRole: %s. Supported values are: %s.", m.TargetRole, strings.Join(GetAttachmentDetailsTargetRoleEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AttachmentDetailsTargetRoleEnum Enum with underlying type: string
type AttachmentDetailsTargetRoleEnum string

// Set of constants representing the allowable values for AttachmentDetailsTargetRoleEnum
const (
	AttachmentDetailsTargetRoleParent AttachmentDetailsTargetRoleEnum = "PARENT"
	AttachmentDetailsTargetRoleChild  AttachmentDetailsTargetRoleEnum = "CHILD"
)

var mappingAttachmentDetailsTargetRoleEnum = map[string]AttachmentDetailsTargetRoleEnum{
	"PARENT": AttachmentDetailsTargetRoleParent,
	"CHILD":  AttachmentDetailsTargetRoleChild,
}

var mappingAttachmentDetailsTargetRoleEnumLowerCase = map[string]AttachmentDetailsTargetRoleEnum{
	"parent": AttachmentDetailsTargetRoleParent,
	"child":  AttachmentDetailsTargetRoleChild,
}

// GetAttachmentDetailsTargetRoleEnumValues Enumerates the set of values for AttachmentDetailsTargetRoleEnum
func GetAttachmentDetailsTargetRoleEnumValues() []AttachmentDetailsTargetRoleEnum {
	values := make([]AttachmentDetailsTargetRoleEnum, 0)
	for _, v := range mappingAttachmentDetailsTargetRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetAttachmentDetailsTargetRoleEnumStringValues Enumerates the set of values in String for AttachmentDetailsTargetRoleEnum
func GetAttachmentDetailsTargetRoleEnumStringValues() []string {
	return []string{
		"PARENT",
		"CHILD",
	}
}

// GetMappingAttachmentDetailsTargetRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAttachmentDetailsTargetRoleEnum(val string) (AttachmentDetailsTargetRoleEnum, bool) {
	enum, ok := mappingAttachmentDetailsTargetRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
