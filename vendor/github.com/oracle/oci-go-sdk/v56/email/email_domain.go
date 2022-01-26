// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// API for the Email Delivery service. Use this API to send high-volume, application-generated
// emails. For more information, see Overview of the Email Delivery Service (https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//
// **Note:** Write actions (POST, UPDATE, DELETE) may take several minutes to propagate and be reflected by the API.
// If a subsequent read request fails to reflect your changes, wait a few minutes and try again.
//

package email

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// EmailDomain The properties that define a email domain.
// A Email Domain contains configuration used to assert responsibility for emails sent from that domain.
type EmailDomain struct {

	// The name of the email domain in the Internet Domain Name System (DNS).
	// Example: `example.net`
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the email domain.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains this email domain.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The current state of the email domain.
	LifecycleState EmailDomainLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DKIM key
	// that will be used to sign mail sent from this email domain.
	ActiveDkimId *string `mandatory:"false" json:"activeDkimId"`

	// Value of the SPF field. For more information about SPF, please see
	// SPF Authentication (https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm#components).
	IsSpf *bool `mandatory:"false" json:"isSpf"`

	// The description of a email domain.
	Description *string `mandatory:"false" json:"description"`

	// The time the email domain was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format, "YYYY-MM-ddThh:mmZ".
	// Example: `2021-02-12T22:47:12.613Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m EmailDomain) String() string {
	return common.PointerString(m)
}

// EmailDomainLifecycleStateEnum Enum with underlying type: string
type EmailDomainLifecycleStateEnum string

// Set of constants representing the allowable values for EmailDomainLifecycleStateEnum
const (
	EmailDomainLifecycleStateActive   EmailDomainLifecycleStateEnum = "ACTIVE"
	EmailDomainLifecycleStateCreating EmailDomainLifecycleStateEnum = "CREATING"
	EmailDomainLifecycleStateDeleting EmailDomainLifecycleStateEnum = "DELETING"
	EmailDomainLifecycleStateDeleted  EmailDomainLifecycleStateEnum = "DELETED"
	EmailDomainLifecycleStateFailed   EmailDomainLifecycleStateEnum = "FAILED"
	EmailDomainLifecycleStateUpdating EmailDomainLifecycleStateEnum = "UPDATING"
)

var mappingEmailDomainLifecycleState = map[string]EmailDomainLifecycleStateEnum{
	"ACTIVE":   EmailDomainLifecycleStateActive,
	"CREATING": EmailDomainLifecycleStateCreating,
	"DELETING": EmailDomainLifecycleStateDeleting,
	"DELETED":  EmailDomainLifecycleStateDeleted,
	"FAILED":   EmailDomainLifecycleStateFailed,
	"UPDATING": EmailDomainLifecycleStateUpdating,
}

// GetEmailDomainLifecycleStateEnumValues Enumerates the set of values for EmailDomainLifecycleStateEnum
func GetEmailDomainLifecycleStateEnumValues() []EmailDomainLifecycleStateEnum {
	values := make([]EmailDomainLifecycleStateEnum, 0)
	for _, v := range mappingEmailDomainLifecycleState {
		values = append(values, v)
	}
	return values
}
