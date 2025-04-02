// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// Use the Email Delivery API to do the necessary set up to send high-volume and application-generated emails through the OCI Email Delivery service.
// For more information, see Overview of the Email Delivery Service (https://docs.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//  **Note:** Write actions (POST, UPDATE, DELETE) may take several minutes to propagate and be reflected by the API.
//  If a subsequent read request fails to reflect your changes, wait a few minutes and try again.
//

package email

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateEmailIpPoolDetails The attributes to update an IpPool which will be used to route emails.
type UpdateEmailIpPoolDetails struct {

	// The description of the IpPool. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The ADD/REMOVE operations are used to assign or un-assign public IPs from an IpPool, respectively.
	OutboundIpOperationType UpdateEmailIpPoolDetailsOutboundIpOperationTypeEnum `mandatory:"false" json:"outboundIpOperationType,omitempty"`

	// List of public IPs to ADD or REMOVE from the IpPool.
	// Public IPs must be in the AVAILABLE state to be assigned to the IpPool.
	// After a public IP is unassigned, it will be marked as AVAILABLE and can be assigned to another IpPool.
	// The last IP removed from the Pool will be deleted from the IP Pool after 24 hours.
	OutboundIps []string `mandatory:"false" json:"outboundIps"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateEmailIpPoolDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateEmailIpPoolDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateEmailIpPoolDetailsOutboundIpOperationTypeEnum(string(m.OutboundIpOperationType)); !ok && m.OutboundIpOperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OutboundIpOperationType: %s. Supported values are: %s.", m.OutboundIpOperationType, strings.Join(GetUpdateEmailIpPoolDetailsOutboundIpOperationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateEmailIpPoolDetailsOutboundIpOperationTypeEnum Enum with underlying type: string
type UpdateEmailIpPoolDetailsOutboundIpOperationTypeEnum string

// Set of constants representing the allowable values for UpdateEmailIpPoolDetailsOutboundIpOperationTypeEnum
const (
	UpdateEmailIpPoolDetailsOutboundIpOperationTypeAdd    UpdateEmailIpPoolDetailsOutboundIpOperationTypeEnum = "ADD"
	UpdateEmailIpPoolDetailsOutboundIpOperationTypeRemove UpdateEmailIpPoolDetailsOutboundIpOperationTypeEnum = "REMOVE"
)

var mappingUpdateEmailIpPoolDetailsOutboundIpOperationTypeEnum = map[string]UpdateEmailIpPoolDetailsOutboundIpOperationTypeEnum{
	"ADD":    UpdateEmailIpPoolDetailsOutboundIpOperationTypeAdd,
	"REMOVE": UpdateEmailIpPoolDetailsOutboundIpOperationTypeRemove,
}

var mappingUpdateEmailIpPoolDetailsOutboundIpOperationTypeEnumLowerCase = map[string]UpdateEmailIpPoolDetailsOutboundIpOperationTypeEnum{
	"add":    UpdateEmailIpPoolDetailsOutboundIpOperationTypeAdd,
	"remove": UpdateEmailIpPoolDetailsOutboundIpOperationTypeRemove,
}

// GetUpdateEmailIpPoolDetailsOutboundIpOperationTypeEnumValues Enumerates the set of values for UpdateEmailIpPoolDetailsOutboundIpOperationTypeEnum
func GetUpdateEmailIpPoolDetailsOutboundIpOperationTypeEnumValues() []UpdateEmailIpPoolDetailsOutboundIpOperationTypeEnum {
	values := make([]UpdateEmailIpPoolDetailsOutboundIpOperationTypeEnum, 0)
	for _, v := range mappingUpdateEmailIpPoolDetailsOutboundIpOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateEmailIpPoolDetailsOutboundIpOperationTypeEnumStringValues Enumerates the set of values in String for UpdateEmailIpPoolDetailsOutboundIpOperationTypeEnum
func GetUpdateEmailIpPoolDetailsOutboundIpOperationTypeEnumStringValues() []string {
	return []string{
		"ADD",
		"REMOVE",
	}
}

// GetMappingUpdateEmailIpPoolDetailsOutboundIpOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateEmailIpPoolDetailsOutboundIpOperationTypeEnum(val string) (UpdateEmailIpPoolDetailsOutboundIpOperationTypeEnum, bool) {
	enum, ok := mappingUpdateEmailIpPoolDetailsOutboundIpOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
