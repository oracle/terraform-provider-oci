// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// Use the Email Delivery API to do the necessary set up to send high-volume and application-generated emails through the OCI Email Delivery service.
// For more information, see Overview of the Email Delivery Service (https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//  **Note:** Write actions (POST, UPDATE, DELETE) may take several minutes to propagate and be reflected by the API.
//  If a subsequent read request fails to reflect your changes, wait a few minutes and try again.
//

package email

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EmailReturnPath The properties that define a Email Return Path
type EmailReturnPath struct {

	// The email return path domain in the Internet Domain Name System (DNS).
	// Example: `iad1.rp.example.com`
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the email return path.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the EmailDomain
	// that this email return path belongs to.
	ParentResourceId *string `mandatory:"true" json:"parentResourceId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains this email return path.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The current state of the email return path.
	LifecycleState EmailReturnPathLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a
	// resource in 'Failed' state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The description of the email return path. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The name of the DNS subdomain that must be provisioned to enable email recipients to verify Email Return Path.
	// It is usually created with a CNAME record set to the cnameRecordValue.
	DnsSubdomainName *string `mandatory:"false" json:"dnsSubdomainName"`

	// The DNS CNAME record value to provision to the Return Patn DNS subdomain, when using the CNAME method for Email Return Path setup (preferred).
	CnameRecordValue *string `mandatory:"false" json:"cnameRecordValue"`

	// The time the email return path was created.
	// Times are expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format, "YYYY-MM-ddThh:mmZ".
	// Example: `2021-02-12T22:47:12.613Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time of the last change to the Email Return Path configuration, due to a state change or
	// an update operation.
	// Times are expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format, "YYYY-MM-ddThh:mmZ".
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

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

func (m EmailReturnPath) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EmailReturnPath) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingEmailReturnPathLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEmailReturnPathLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EmailReturnPathLifecycleStateEnum Enum with underlying type: string
type EmailReturnPathLifecycleStateEnum string

// Set of constants representing the allowable values for EmailReturnPathLifecycleStateEnum
const (
	EmailReturnPathLifecycleStateActive         EmailReturnPathLifecycleStateEnum = "ACTIVE"
	EmailReturnPathLifecycleStateCreating       EmailReturnPathLifecycleStateEnum = "CREATING"
	EmailReturnPathLifecycleStateDeleting       EmailReturnPathLifecycleStateEnum = "DELETING"
	EmailReturnPathLifecycleStateDeleted        EmailReturnPathLifecycleStateEnum = "DELETED"
	EmailReturnPathLifecycleStateFailed         EmailReturnPathLifecycleStateEnum = "FAILED"
	EmailReturnPathLifecycleStateNeedsAttention EmailReturnPathLifecycleStateEnum = "NEEDS_ATTENTION"
	EmailReturnPathLifecycleStateUpdating       EmailReturnPathLifecycleStateEnum = "UPDATING"
)

var mappingEmailReturnPathLifecycleStateEnum = map[string]EmailReturnPathLifecycleStateEnum{
	"ACTIVE":          EmailReturnPathLifecycleStateActive,
	"CREATING":        EmailReturnPathLifecycleStateCreating,
	"DELETING":        EmailReturnPathLifecycleStateDeleting,
	"DELETED":         EmailReturnPathLifecycleStateDeleted,
	"FAILED":          EmailReturnPathLifecycleStateFailed,
	"NEEDS_ATTENTION": EmailReturnPathLifecycleStateNeedsAttention,
	"UPDATING":        EmailReturnPathLifecycleStateUpdating,
}

var mappingEmailReturnPathLifecycleStateEnumLowerCase = map[string]EmailReturnPathLifecycleStateEnum{
	"active":          EmailReturnPathLifecycleStateActive,
	"creating":        EmailReturnPathLifecycleStateCreating,
	"deleting":        EmailReturnPathLifecycleStateDeleting,
	"deleted":         EmailReturnPathLifecycleStateDeleted,
	"failed":          EmailReturnPathLifecycleStateFailed,
	"needs_attention": EmailReturnPathLifecycleStateNeedsAttention,
	"updating":        EmailReturnPathLifecycleStateUpdating,
}

// GetEmailReturnPathLifecycleStateEnumValues Enumerates the set of values for EmailReturnPathLifecycleStateEnum
func GetEmailReturnPathLifecycleStateEnumValues() []EmailReturnPathLifecycleStateEnum {
	values := make([]EmailReturnPathLifecycleStateEnum, 0)
	for _, v := range mappingEmailReturnPathLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetEmailReturnPathLifecycleStateEnumStringValues Enumerates the set of values in String for EmailReturnPathLifecycleStateEnum
func GetEmailReturnPathLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
		"UPDATING",
	}
}

// GetMappingEmailReturnPathLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmailReturnPathLifecycleStateEnum(val string) (EmailReturnPathLifecycleStateEnum, bool) {
	enum, ok := mappingEmailReturnPathLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
