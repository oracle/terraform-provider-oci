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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// Dkim The properties that define a DKIM.
type Dkim struct {

	// The DKIM selector.
	// If the same domain is managed in more than one region, each region must use different selectors.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DKIM.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the email domain
	// that this DKIM belongs to.
	EmailDomainId *string `mandatory:"true" json:"emailDomainId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains this DKIM.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The current state of the DKIM.
	LifecycleState DkimLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a resource.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The description of the DKIM. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The time the DKIM was created.
	// Times are expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format, "YYYY-MM-ddThh:mmZ".
	// Example: `2021-02-12T22:47:12.613Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time of the last change to the DKIM configuration, due to a state change or
	// an update operation.
	// Times are expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format, "YYYY-MM-ddThh:mmZ".
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The name of the DNS subdomain that must be provisioned to enable email recipients to verify DKIM signatures.
	// It is usually created with a CNAME record set to the cnameRecordValue
	DnsSubdomainName *string `mandatory:"false" json:"dnsSubdomainName"`

	// The DNS CNAME record value to provision to the DKIM DNS subdomain, when using the CNAME method for DKIM setup (preferred).
	CnameRecordValue *string `mandatory:"false" json:"cnameRecordValue"`

	// The DNS TXT record value to provision to the DKIM DNS subdomain in place of using a CNAME record.
	// This is used in cases where a CNAME can not be used, such as when the cnameRecordValue would exceed the maximum length for a DNS entry.
	// This can also be used by customers who have an existing procedure to directly provision TXT records for DKIM.
	// Be aware that many DNS APIs will require you to break this string into segments of less than 255 characters.
	TxtRecordValue *string `mandatory:"false" json:"txtRecordValue"`

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

func (m Dkim) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Dkim) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDkimLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDkimLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DkimLifecycleStateEnum Enum with underlying type: string
type DkimLifecycleStateEnum string

// Set of constants representing the allowable values for DkimLifecycleStateEnum
const (
	DkimLifecycleStateActive         DkimLifecycleStateEnum = "ACTIVE"
	DkimLifecycleStateCreating       DkimLifecycleStateEnum = "CREATING"
	DkimLifecycleStateDeleting       DkimLifecycleStateEnum = "DELETING"
	DkimLifecycleStateDeleted        DkimLifecycleStateEnum = "DELETED"
	DkimLifecycleStateFailed         DkimLifecycleStateEnum = "FAILED"
	DkimLifecycleStateInactive       DkimLifecycleStateEnum = "INACTIVE"
	DkimLifecycleStateNeedsAttention DkimLifecycleStateEnum = "NEEDS_ATTENTION"
	DkimLifecycleStateUpdating       DkimLifecycleStateEnum = "UPDATING"
)

var mappingDkimLifecycleStateEnum = map[string]DkimLifecycleStateEnum{
	"ACTIVE":          DkimLifecycleStateActive,
	"CREATING":        DkimLifecycleStateCreating,
	"DELETING":        DkimLifecycleStateDeleting,
	"DELETED":         DkimLifecycleStateDeleted,
	"FAILED":          DkimLifecycleStateFailed,
	"INACTIVE":        DkimLifecycleStateInactive,
	"NEEDS_ATTENTION": DkimLifecycleStateNeedsAttention,
	"UPDATING":        DkimLifecycleStateUpdating,
}

// GetDkimLifecycleStateEnumValues Enumerates the set of values for DkimLifecycleStateEnum
func GetDkimLifecycleStateEnumValues() []DkimLifecycleStateEnum {
	values := make([]DkimLifecycleStateEnum, 0)
	for _, v := range mappingDkimLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDkimLifecycleStateEnumStringValues Enumerates the set of values in String for DkimLifecycleStateEnum
func GetDkimLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"INACTIVE",
		"NEEDS_ATTENTION",
		"UPDATING",
	}
}

// GetMappingDkimLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDkimLifecycleStateEnum(val string) (DkimLifecycleStateEnum, bool) {
	mappingDkimLifecycleStateEnumIgnoreCase := make(map[string]DkimLifecycleStateEnum)
	for k, v := range mappingDkimLifecycleStateEnum {
		mappingDkimLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDkimLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
