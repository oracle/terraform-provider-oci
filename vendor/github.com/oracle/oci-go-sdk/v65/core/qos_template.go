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

// QosTemplate Quality of Service template to map DSCP values to respective Class of Services; PREMIUM (P1), DEFAULT (P2), BULK (P3), SCAVENGER (P4).
type QosTemplate struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Quality of Service template.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The Oracle ID of Quality of Service template (OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
	Id *string `mandatory:"true" json:"id"`

	// The Tenancy's Oracle ID (OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) for which Quality of Service template is applicable.
	TenancyId *string `mandatory:"true" json:"tenancyId"`

	// The current state of Quality of Service template.
	LifecycleState QosTemplateLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// List of QosMappings consisting of DSCP values with their respective ClassOfService. Eg {43 - PREMIUM}
	QosMappings []QosMappings `mandatory:"true" json:"qosMappings"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The date and time the Quality of Service template was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m QosTemplate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QosTemplate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingQosTemplateLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetQosTemplateLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// QosTemplateLifecycleStateEnum Enum with underlying type: string
type QosTemplateLifecycleStateEnum string

// Set of constants representing the allowable values for QosTemplateLifecycleStateEnum
const (
	QosTemplateLifecycleStateProvisioning QosTemplateLifecycleStateEnum = "PROVISIONING"
	QosTemplateLifecycleStateAvailable    QosTemplateLifecycleStateEnum = "AVAILABLE"
	QosTemplateLifecycleStateTerminating  QosTemplateLifecycleStateEnum = "TERMINATING"
	QosTemplateLifecycleStateTerminated   QosTemplateLifecycleStateEnum = "TERMINATED"
	QosTemplateLifecycleStateUpdating     QosTemplateLifecycleStateEnum = "UPDATING"
)

var mappingQosTemplateLifecycleStateEnum = map[string]QosTemplateLifecycleStateEnum{
	"PROVISIONING": QosTemplateLifecycleStateProvisioning,
	"AVAILABLE":    QosTemplateLifecycleStateAvailable,
	"TERMINATING":  QosTemplateLifecycleStateTerminating,
	"TERMINATED":   QosTemplateLifecycleStateTerminated,
	"UPDATING":     QosTemplateLifecycleStateUpdating,
}

var mappingQosTemplateLifecycleStateEnumLowerCase = map[string]QosTemplateLifecycleStateEnum{
	"provisioning": QosTemplateLifecycleStateProvisioning,
	"available":    QosTemplateLifecycleStateAvailable,
	"terminating":  QosTemplateLifecycleStateTerminating,
	"terminated":   QosTemplateLifecycleStateTerminated,
	"updating":     QosTemplateLifecycleStateUpdating,
}

// GetQosTemplateLifecycleStateEnumValues Enumerates the set of values for QosTemplateLifecycleStateEnum
func GetQosTemplateLifecycleStateEnumValues() []QosTemplateLifecycleStateEnum {
	values := make([]QosTemplateLifecycleStateEnum, 0)
	for _, v := range mappingQosTemplateLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetQosTemplateLifecycleStateEnumStringValues Enumerates the set of values in String for QosTemplateLifecycleStateEnum
func GetQosTemplateLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
		"UPDATING",
	}
}

// GetMappingQosTemplateLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQosTemplateLifecycleStateEnum(val string) (QosTemplateLifecycleStateEnum, bool) {
	enum, ok := mappingQosTemplateLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
