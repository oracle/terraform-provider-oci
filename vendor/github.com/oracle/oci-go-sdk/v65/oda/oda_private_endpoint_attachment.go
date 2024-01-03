// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OdaPrivateEndpointAttachment ODA Private Endpoint Attachment is used to attach ODA Private Endpoint to ODA (Digital Assistant) Instance.
type OdaPrivateEndpointAttachment struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ODA Private Endpoint Attachment.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the attached ODA Instance.
	OdaInstanceId *string `mandatory:"true" json:"odaInstanceId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ODA Private Endpoint.
	OdaPrivateEndpointId *string `mandatory:"true" json:"odaPrivateEndpointId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that the ODA private endpoint attachment belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the ODA Private Endpoint attachment.
	LifecycleState OdaPrivateEndpointAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// When the resource was created. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// When the resource was last updated. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m OdaPrivateEndpointAttachment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OdaPrivateEndpointAttachment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOdaPrivateEndpointAttachmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOdaPrivateEndpointAttachmentLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OdaPrivateEndpointAttachmentLifecycleStateEnum Enum with underlying type: string
type OdaPrivateEndpointAttachmentLifecycleStateEnum string

// Set of constants representing the allowable values for OdaPrivateEndpointAttachmentLifecycleStateEnum
const (
	OdaPrivateEndpointAttachmentLifecycleStateCreating OdaPrivateEndpointAttachmentLifecycleStateEnum = "CREATING"
	OdaPrivateEndpointAttachmentLifecycleStateUpdating OdaPrivateEndpointAttachmentLifecycleStateEnum = "UPDATING"
	OdaPrivateEndpointAttachmentLifecycleStateActive   OdaPrivateEndpointAttachmentLifecycleStateEnum = "ACTIVE"
	OdaPrivateEndpointAttachmentLifecycleStateDeleting OdaPrivateEndpointAttachmentLifecycleStateEnum = "DELETING"
	OdaPrivateEndpointAttachmentLifecycleStateDeleted  OdaPrivateEndpointAttachmentLifecycleStateEnum = "DELETED"
	OdaPrivateEndpointAttachmentLifecycleStateFailed   OdaPrivateEndpointAttachmentLifecycleStateEnum = "FAILED"
)

var mappingOdaPrivateEndpointAttachmentLifecycleStateEnum = map[string]OdaPrivateEndpointAttachmentLifecycleStateEnum{
	"CREATING": OdaPrivateEndpointAttachmentLifecycleStateCreating,
	"UPDATING": OdaPrivateEndpointAttachmentLifecycleStateUpdating,
	"ACTIVE":   OdaPrivateEndpointAttachmentLifecycleStateActive,
	"DELETING": OdaPrivateEndpointAttachmentLifecycleStateDeleting,
	"DELETED":  OdaPrivateEndpointAttachmentLifecycleStateDeleted,
	"FAILED":   OdaPrivateEndpointAttachmentLifecycleStateFailed,
}

var mappingOdaPrivateEndpointAttachmentLifecycleStateEnumLowerCase = map[string]OdaPrivateEndpointAttachmentLifecycleStateEnum{
	"creating": OdaPrivateEndpointAttachmentLifecycleStateCreating,
	"updating": OdaPrivateEndpointAttachmentLifecycleStateUpdating,
	"active":   OdaPrivateEndpointAttachmentLifecycleStateActive,
	"deleting": OdaPrivateEndpointAttachmentLifecycleStateDeleting,
	"deleted":  OdaPrivateEndpointAttachmentLifecycleStateDeleted,
	"failed":   OdaPrivateEndpointAttachmentLifecycleStateFailed,
}

// GetOdaPrivateEndpointAttachmentLifecycleStateEnumValues Enumerates the set of values for OdaPrivateEndpointAttachmentLifecycleStateEnum
func GetOdaPrivateEndpointAttachmentLifecycleStateEnumValues() []OdaPrivateEndpointAttachmentLifecycleStateEnum {
	values := make([]OdaPrivateEndpointAttachmentLifecycleStateEnum, 0)
	for _, v := range mappingOdaPrivateEndpointAttachmentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOdaPrivateEndpointAttachmentLifecycleStateEnumStringValues Enumerates the set of values in String for OdaPrivateEndpointAttachmentLifecycleStateEnum
func GetOdaPrivateEndpointAttachmentLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOdaPrivateEndpointAttachmentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOdaPrivateEndpointAttachmentLifecycleStateEnum(val string) (OdaPrivateEndpointAttachmentLifecycleStateEnum, bool) {
	enum, ok := mappingOdaPrivateEndpointAttachmentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
