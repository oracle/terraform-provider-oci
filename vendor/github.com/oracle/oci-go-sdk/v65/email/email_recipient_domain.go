// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// EmailRecipientDomain The properties that define a Recipient Domain
type EmailRecipientDomain struct {

	// The recipient domain in the Internet Domain Name System (DNS). The recipient domain name must be globally unique for this tenancy.
	// Example: `example.com`
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the recipient domain.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains this recipient domain.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The recipient domain's current lifecycle state.
	LifecycleState EmailRecipientDomainLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the recipient domain was created.
	// Times are expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format, "YYYY-MM-ddThh:mmZ".
	// Example: `2021-02-12T22:47:12.613Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time of the last change to the Recipient Domain, due to
	// an update operation.
	// Times are expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format, "YYYY-MM-ddThh:mmZ".
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The description of a recipient domain.
	Description *string `mandatory:"false" json:"description"`

	// The Email Delivery Config OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) used to submit an email by Email Delivery when sent from this recipient domain.
	EmailDeliveryConfigId *string `mandatory:"false" json:"emailDeliveryConfigId"`

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

func (m EmailRecipientDomain) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EmailRecipientDomain) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEmailRecipientDomainLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEmailRecipientDomainLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EmailRecipientDomainLifecycleStateEnum Enum with underlying type: string
type EmailRecipientDomainLifecycleStateEnum string

// Set of constants representing the allowable values for EmailRecipientDomainLifecycleStateEnum
const (
	EmailRecipientDomainLifecycleStateCreating EmailRecipientDomainLifecycleStateEnum = "CREATING"
	EmailRecipientDomainLifecycleStateUpdating EmailRecipientDomainLifecycleStateEnum = "UPDATING"
	EmailRecipientDomainLifecycleStateActive   EmailRecipientDomainLifecycleStateEnum = "ACTIVE"
	EmailRecipientDomainLifecycleStateDeleting EmailRecipientDomainLifecycleStateEnum = "DELETING"
	EmailRecipientDomainLifecycleStateDeleted  EmailRecipientDomainLifecycleStateEnum = "DELETED"
	EmailRecipientDomainLifecycleStateFailed   EmailRecipientDomainLifecycleStateEnum = "FAILED"
)

var mappingEmailRecipientDomainLifecycleStateEnum = map[string]EmailRecipientDomainLifecycleStateEnum{
	"CREATING": EmailRecipientDomainLifecycleStateCreating,
	"UPDATING": EmailRecipientDomainLifecycleStateUpdating,
	"ACTIVE":   EmailRecipientDomainLifecycleStateActive,
	"DELETING": EmailRecipientDomainLifecycleStateDeleting,
	"DELETED":  EmailRecipientDomainLifecycleStateDeleted,
	"FAILED":   EmailRecipientDomainLifecycleStateFailed,
}

var mappingEmailRecipientDomainLifecycleStateEnumLowerCase = map[string]EmailRecipientDomainLifecycleStateEnum{
	"creating": EmailRecipientDomainLifecycleStateCreating,
	"updating": EmailRecipientDomainLifecycleStateUpdating,
	"active":   EmailRecipientDomainLifecycleStateActive,
	"deleting": EmailRecipientDomainLifecycleStateDeleting,
	"deleted":  EmailRecipientDomainLifecycleStateDeleted,
	"failed":   EmailRecipientDomainLifecycleStateFailed,
}

// GetEmailRecipientDomainLifecycleStateEnumValues Enumerates the set of values for EmailRecipientDomainLifecycleStateEnum
func GetEmailRecipientDomainLifecycleStateEnumValues() []EmailRecipientDomainLifecycleStateEnum {
	values := make([]EmailRecipientDomainLifecycleStateEnum, 0)
	for _, v := range mappingEmailRecipientDomainLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetEmailRecipientDomainLifecycleStateEnumStringValues Enumerates the set of values in String for EmailRecipientDomainLifecycleStateEnum
func GetEmailRecipientDomainLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingEmailRecipientDomainLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmailRecipientDomainLifecycleStateEnum(val string) (EmailRecipientDomainLifecycleStateEnum, bool) {
	enum, ok := mappingEmailRecipientDomainLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
