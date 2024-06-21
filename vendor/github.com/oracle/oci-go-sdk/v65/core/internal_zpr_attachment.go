// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InternalZprAttachment An Internal ZPR Attachment
type InternalZprAttachment struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the ZPR attachment lives.
	// Used to have the ability to list all ZprAttachments in a compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resouce this is attached to.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// Needs to remain same for all zpr attachments owned by the service.
	// If this value does not correspond to the policy that VCN-CP will add/enforce then it will cause authorization issues.
	OwnerId *string `mandatory:"true" json:"ownerId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of this ZprAttachment.
	Id *string `mandatory:"false" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The attachment's Origin Id used to enforce ZPR policies.
	OriginId *int64 `mandatory:"false" json:"originId"`

	// The date and time the InternalZprAttachment was created, in the format defined
	// by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The ZprAttachment's current state.
	LifecycleState InternalZprAttachmentLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Security Attributes for this resource. This is unique to ZPR, and helps identify which resources are allowed to be accessed by what permission controls.
	// Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m InternalZprAttachment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternalZprAttachment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingInternalZprAttachmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetInternalZprAttachmentLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InternalZprAttachmentLifecycleStateEnum Enum with underlying type: string
type InternalZprAttachmentLifecycleStateEnum string

// Set of constants representing the allowable values for InternalZprAttachmentLifecycleStateEnum
const (
	InternalZprAttachmentLifecycleStateProvisioning InternalZprAttachmentLifecycleStateEnum = "PROVISIONING"
	InternalZprAttachmentLifecycleStateAvailable    InternalZprAttachmentLifecycleStateEnum = "AVAILABLE"
	InternalZprAttachmentLifecycleStateTerminating  InternalZprAttachmentLifecycleStateEnum = "TERMINATING"
	InternalZprAttachmentLifecycleStateTerminated   InternalZprAttachmentLifecycleStateEnum = "TERMINATED"
)

var mappingInternalZprAttachmentLifecycleStateEnum = map[string]InternalZprAttachmentLifecycleStateEnum{
	"PROVISIONING": InternalZprAttachmentLifecycleStateProvisioning,
	"AVAILABLE":    InternalZprAttachmentLifecycleStateAvailable,
	"TERMINATING":  InternalZprAttachmentLifecycleStateTerminating,
	"TERMINATED":   InternalZprAttachmentLifecycleStateTerminated,
}

var mappingInternalZprAttachmentLifecycleStateEnumLowerCase = map[string]InternalZprAttachmentLifecycleStateEnum{
	"provisioning": InternalZprAttachmentLifecycleStateProvisioning,
	"available":    InternalZprAttachmentLifecycleStateAvailable,
	"terminating":  InternalZprAttachmentLifecycleStateTerminating,
	"terminated":   InternalZprAttachmentLifecycleStateTerminated,
}

// GetInternalZprAttachmentLifecycleStateEnumValues Enumerates the set of values for InternalZprAttachmentLifecycleStateEnum
func GetInternalZprAttachmentLifecycleStateEnumValues() []InternalZprAttachmentLifecycleStateEnum {
	values := make([]InternalZprAttachmentLifecycleStateEnum, 0)
	for _, v := range mappingInternalZprAttachmentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetInternalZprAttachmentLifecycleStateEnumStringValues Enumerates the set of values in String for InternalZprAttachmentLifecycleStateEnum
func GetInternalZprAttachmentLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
	}
}

// GetMappingInternalZprAttachmentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInternalZprAttachmentLifecycleStateEnum(val string) (InternalZprAttachmentLifecycleStateEnum, bool) {
	enum, ok := mappingInternalZprAttachmentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
