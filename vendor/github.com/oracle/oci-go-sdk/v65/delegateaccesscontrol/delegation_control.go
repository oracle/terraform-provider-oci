// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Delegate Access Control API
//
// Oracle Delegate Access Control allows ExaCC and ExaCS customers to delegate management of their Exadata resources operators outside their tenancies.
// With Delegate Access Control, Support Providers can deliver managed services using comprehensive and robust tooling built on the OCI platform.
// Customers maintain control over who has access to the delegated resources in their tenancy and what actions can be taken.
// Enterprises managing resources across multiple tenants can use Delegate Access Control to streamline management tasks.
// Using logging service, customers can view a near real-time audit report of all actions performed by a Service Provider operator.
//

package delegateaccesscontrol

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DelegationControl Delegation Control enables you to grant, audit, or revoke the access Oracle has to your Exadata Cloud infrastructure, and obtain audit reports of all actions taken by a human operator, in a near real-time manner.
type DelegationControl struct {

	// The OCID of the Delegation Control.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the Delegation Control.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Name of the Delegation Control. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Resource type for which the Delegation Control is applicable to.
	ResourceType DelegationControlResourceTypeEnum `mandatory:"true" json:"resourceType"`

	// Description of the Delegation Control.
	Description *string `mandatory:"false" json:"description"`

	// number of approvals required.
	NumApprovalsRequired *int `mandatory:"false" json:"numApprovalsRequired"`

	// List of pre-approved Service Provider Action names. The list of pre-defined Service Provider Actions can be obtained from the ListServiceProviderActions API. Delegated Resource Access Requests associated with a resource governed by this Delegation Control will be
	// automatically approved if the Delegated Resource Access Request only contain Service Provider Actions in the pre-approved list.
	PreApprovedServiceProviderActionNames []string `mandatory:"false" json:"preApprovedServiceProviderActionNames"`

	// List of Delegation Subscription OCID that are allowed for this Delegation Control. The allowed subscriptions will determine the available Service Provider Actions. Only support operators for the allowed subscriptions are allowed to create Delegated Resource Access Request.
	DelegationSubscriptionIds []string `mandatory:"false" json:"delegationSubscriptionIds"`

	// Set to true to allow all Delegated Resource Access Request to be approved automatically during maintenance.
	IsAutoApproveDuringMaintenance *bool `mandatory:"false" json:"isAutoApproveDuringMaintenance"`

	// The OCID of the selected resources that this Delegation Control is applicable to.
	ResourceIds []string `mandatory:"false" json:"resourceIds"`

	// The OCID of the OCI Notification topic to publish messages related to this Delegation Control.
	NotificationTopicId *string `mandatory:"false" json:"notificationTopicId"`

	// The format of the OCI Notification messages for this Delegation Control.
	NotificationMessageFormat DelegationControlNotificationMessageFormatEnum `mandatory:"false" json:"notificationMessageFormat,omitempty"`

	// The OCID of the OCI Vault that will store the secrets containing the SSH keys to access the resource governed by this Delegation Control by Delegate Access Control Service. This property is required when resourceType is CLOUDVMCLUSTER. Delegate Access Control Service will generate the SSH keys and store them as secrets in the OCI Vault.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID of the Master Encryption Key in the OCI Vault specified by vaultId. This key will be used to encrypt the SSH keys to access the resource governed by this Delegation Control by Delegate Access Control Service. This property is required when resourceType is CLOUDVMCLUSTER.
	VaultKeyId *string `mandatory:"false" json:"vaultKeyId"`

	// The current lifecycle state of the Delegation Control.
	LifecycleState DelegationControlLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Description of the current lifecycle state in more detail.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// Time when the Delegation Control was created expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z'
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time when the Delegation Control was last modified expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format, e.g. '2020-05-22T21:10:29.600Z'
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Time when the Delegation Control was deleted expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)timestamp format, e.g. '2020-05-22T21:10:29.600Z'.
	// Note a deleted Delegation Control still stays in the system, so that you can still audit Service Provider Actions associated with Delegated Resource Access Requests
	// raised on target resources governed by the deleted Delegation Control.
	TimeDeleted *common.SDKTime `mandatory:"false" json:"timeDeleted"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DelegationControl) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DelegationControl) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDelegationControlResourceTypeEnum(string(m.ResourceType)); !ok && m.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", m.ResourceType, strings.Join(GetDelegationControlResourceTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDelegationControlNotificationMessageFormatEnum(string(m.NotificationMessageFormat)); !ok && m.NotificationMessageFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NotificationMessageFormat: %s. Supported values are: %s.", m.NotificationMessageFormat, strings.Join(GetDelegationControlNotificationMessageFormatEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDelegationControlLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDelegationControlLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DelegationControlNotificationMessageFormatEnum Enum with underlying type: string
type DelegationControlNotificationMessageFormatEnum string

// Set of constants representing the allowable values for DelegationControlNotificationMessageFormatEnum
const (
	DelegationControlNotificationMessageFormatJson DelegationControlNotificationMessageFormatEnum = "JSON"
	DelegationControlNotificationMessageFormatHtml DelegationControlNotificationMessageFormatEnum = "HTML"
)

var mappingDelegationControlNotificationMessageFormatEnum = map[string]DelegationControlNotificationMessageFormatEnum{
	"JSON": DelegationControlNotificationMessageFormatJson,
	"HTML": DelegationControlNotificationMessageFormatHtml,
}

var mappingDelegationControlNotificationMessageFormatEnumLowerCase = map[string]DelegationControlNotificationMessageFormatEnum{
	"json": DelegationControlNotificationMessageFormatJson,
	"html": DelegationControlNotificationMessageFormatHtml,
}

// GetDelegationControlNotificationMessageFormatEnumValues Enumerates the set of values for DelegationControlNotificationMessageFormatEnum
func GetDelegationControlNotificationMessageFormatEnumValues() []DelegationControlNotificationMessageFormatEnum {
	values := make([]DelegationControlNotificationMessageFormatEnum, 0)
	for _, v := range mappingDelegationControlNotificationMessageFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetDelegationControlNotificationMessageFormatEnumStringValues Enumerates the set of values in String for DelegationControlNotificationMessageFormatEnum
func GetDelegationControlNotificationMessageFormatEnumStringValues() []string {
	return []string{
		"JSON",
		"HTML",
	}
}

// GetMappingDelegationControlNotificationMessageFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDelegationControlNotificationMessageFormatEnum(val string) (DelegationControlNotificationMessageFormatEnum, bool) {
	enum, ok := mappingDelegationControlNotificationMessageFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DelegationControlLifecycleStateEnum Enum with underlying type: string
type DelegationControlLifecycleStateEnum string

// Set of constants representing the allowable values for DelegationControlLifecycleStateEnum
const (
	DelegationControlLifecycleStateCreating       DelegationControlLifecycleStateEnum = "CREATING"
	DelegationControlLifecycleStateActive         DelegationControlLifecycleStateEnum = "ACTIVE"
	DelegationControlLifecycleStateUpdating       DelegationControlLifecycleStateEnum = "UPDATING"
	DelegationControlLifecycleStateDeleting       DelegationControlLifecycleStateEnum = "DELETING"
	DelegationControlLifecycleStateDeleted        DelegationControlLifecycleStateEnum = "DELETED"
	DelegationControlLifecycleStateFailed         DelegationControlLifecycleStateEnum = "FAILED"
	DelegationControlLifecycleStateNeedsAttention DelegationControlLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingDelegationControlLifecycleStateEnum = map[string]DelegationControlLifecycleStateEnum{
	"CREATING":        DelegationControlLifecycleStateCreating,
	"ACTIVE":          DelegationControlLifecycleStateActive,
	"UPDATING":        DelegationControlLifecycleStateUpdating,
	"DELETING":        DelegationControlLifecycleStateDeleting,
	"DELETED":         DelegationControlLifecycleStateDeleted,
	"FAILED":          DelegationControlLifecycleStateFailed,
	"NEEDS_ATTENTION": DelegationControlLifecycleStateNeedsAttention,
}

var mappingDelegationControlLifecycleStateEnumLowerCase = map[string]DelegationControlLifecycleStateEnum{
	"creating":        DelegationControlLifecycleStateCreating,
	"active":          DelegationControlLifecycleStateActive,
	"updating":        DelegationControlLifecycleStateUpdating,
	"deleting":        DelegationControlLifecycleStateDeleting,
	"deleted":         DelegationControlLifecycleStateDeleted,
	"failed":          DelegationControlLifecycleStateFailed,
	"needs_attention": DelegationControlLifecycleStateNeedsAttention,
}

// GetDelegationControlLifecycleStateEnumValues Enumerates the set of values for DelegationControlLifecycleStateEnum
func GetDelegationControlLifecycleStateEnumValues() []DelegationControlLifecycleStateEnum {
	values := make([]DelegationControlLifecycleStateEnum, 0)
	for _, v := range mappingDelegationControlLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDelegationControlLifecycleStateEnumStringValues Enumerates the set of values in String for DelegationControlLifecycleStateEnum
func GetDelegationControlLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingDelegationControlLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDelegationControlLifecycleStateEnum(val string) (DelegationControlLifecycleStateEnum, bool) {
	enum, ok := mappingDelegationControlLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
