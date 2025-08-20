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

// EmailDeliveryConfigIpAssociationSummary The properties that define a a delivery configuration ip association.
type EmailDeliveryConfigIpAssociationSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the delivery configuration ip association.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains this delivery config ip association.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the email delivery configuation.
	EmailDeliveryConfigId *string `mandatory:"true" json:"emailDeliveryConfigId"`

	// Outbound IP to be assigned with delivery configuration. The value can be either an IPv4 address or a an IPv6 address.
	// Example IPv4- 192.168.0.1 IPv6- 2001:0db8:85a3:0000:0000:8a2e:0370:7334
	OutboundIp *string `mandatory:"true" json:"outboundIp"`

	// The delivery configuration's current lifecycle state.
	LifecycleState EmailDeliveryConfigIpAssociationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the delivery configuration was created.
	// Times are expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format, "YYYY-MM-ddThh:mmZ".
	// Example: `2021-02-12T22:47:12.613Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time of the last change to the delivery configuration, due to a state change or
	// an update operation.
	// Times are expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format, "YYYY-MM-ddThh:mmZ".
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a
	// resource in 'INACTIVE' state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m EmailDeliveryConfigIpAssociationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EmailDeliveryConfigIpAssociationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEmailDeliveryConfigIpAssociationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEmailDeliveryConfigIpAssociationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
