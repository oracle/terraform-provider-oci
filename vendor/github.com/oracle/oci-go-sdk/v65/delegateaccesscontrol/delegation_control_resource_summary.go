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

// DelegationControlResourceSummary Details of the resources that this Delegation Control is applicable to.
type DelegationControlResourceSummary struct {

	// OCID of the resource.
	Id *string `mandatory:"false" json:"id"`

	// The current status of the resource in Delegation Control.
	ResourceStatus DelegationControlResourceSummaryResourceStatusEnum `mandatory:"false" json:"resourceStatus,omitempty"`
}

func (m DelegationControlResourceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DelegationControlResourceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDelegationControlResourceSummaryResourceStatusEnum(string(m.ResourceStatus)); !ok && m.ResourceStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceStatus: %s. Supported values are: %s.", m.ResourceStatus, strings.Join(GetDelegationControlResourceSummaryResourceStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DelegationControlResourceSummaryResourceStatusEnum Enum with underlying type: string
type DelegationControlResourceSummaryResourceStatusEnum string

// Set of constants representing the allowable values for DelegationControlResourceSummaryResourceStatusEnum
const (
	DelegationControlResourceSummaryResourceStatusCreated        DelegationControlResourceSummaryResourceStatusEnum = "CREATED"
	DelegationControlResourceSummaryResourceStatusApplying       DelegationControlResourceSummaryResourceStatusEnum = "APPLYING"
	DelegationControlResourceSummaryResourceStatusApplied        DelegationControlResourceSummaryResourceStatusEnum = "APPLIED"
	DelegationControlResourceSummaryResourceStatusApplyFailed    DelegationControlResourceSummaryResourceStatusEnum = "APPLY_FAILED"
	DelegationControlResourceSummaryResourceStatusUpdating       DelegationControlResourceSummaryResourceStatusEnum = "UPDATING"
	DelegationControlResourceSummaryResourceStatusUpdateFailed   DelegationControlResourceSummaryResourceStatusEnum = "UPDATE_FAILED"
	DelegationControlResourceSummaryResourceStatusDeleting       DelegationControlResourceSummaryResourceStatusEnum = "DELETING"
	DelegationControlResourceSummaryResourceStatusDeleted        DelegationControlResourceSummaryResourceStatusEnum = "DELETED"
	DelegationControlResourceSummaryResourceStatusDeletionFailed DelegationControlResourceSummaryResourceStatusEnum = "DELETION_FAILED"
)

var mappingDelegationControlResourceSummaryResourceStatusEnum = map[string]DelegationControlResourceSummaryResourceStatusEnum{
	"CREATED":         DelegationControlResourceSummaryResourceStatusCreated,
	"APPLYING":        DelegationControlResourceSummaryResourceStatusApplying,
	"APPLIED":         DelegationControlResourceSummaryResourceStatusApplied,
	"APPLY_FAILED":    DelegationControlResourceSummaryResourceStatusApplyFailed,
	"UPDATING":        DelegationControlResourceSummaryResourceStatusUpdating,
	"UPDATE_FAILED":   DelegationControlResourceSummaryResourceStatusUpdateFailed,
	"DELETING":        DelegationControlResourceSummaryResourceStatusDeleting,
	"DELETED":         DelegationControlResourceSummaryResourceStatusDeleted,
	"DELETION_FAILED": DelegationControlResourceSummaryResourceStatusDeletionFailed,
}

var mappingDelegationControlResourceSummaryResourceStatusEnumLowerCase = map[string]DelegationControlResourceSummaryResourceStatusEnum{
	"created":         DelegationControlResourceSummaryResourceStatusCreated,
	"applying":        DelegationControlResourceSummaryResourceStatusApplying,
	"applied":         DelegationControlResourceSummaryResourceStatusApplied,
	"apply_failed":    DelegationControlResourceSummaryResourceStatusApplyFailed,
	"updating":        DelegationControlResourceSummaryResourceStatusUpdating,
	"update_failed":   DelegationControlResourceSummaryResourceStatusUpdateFailed,
	"deleting":        DelegationControlResourceSummaryResourceStatusDeleting,
	"deleted":         DelegationControlResourceSummaryResourceStatusDeleted,
	"deletion_failed": DelegationControlResourceSummaryResourceStatusDeletionFailed,
}

// GetDelegationControlResourceSummaryResourceStatusEnumValues Enumerates the set of values for DelegationControlResourceSummaryResourceStatusEnum
func GetDelegationControlResourceSummaryResourceStatusEnumValues() []DelegationControlResourceSummaryResourceStatusEnum {
	values := make([]DelegationControlResourceSummaryResourceStatusEnum, 0)
	for _, v := range mappingDelegationControlResourceSummaryResourceStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDelegationControlResourceSummaryResourceStatusEnumStringValues Enumerates the set of values in String for DelegationControlResourceSummaryResourceStatusEnum
func GetDelegationControlResourceSummaryResourceStatusEnumStringValues() []string {
	return []string{
		"CREATED",
		"APPLYING",
		"APPLIED",
		"APPLY_FAILED",
		"UPDATING",
		"UPDATE_FAILED",
		"DELETING",
		"DELETED",
		"DELETION_FAILED",
	}
}

// GetMappingDelegationControlResourceSummaryResourceStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDelegationControlResourceSummaryResourceStatusEnum(val string) (DelegationControlResourceSummaryResourceStatusEnum, bool) {
	enum, ok := mappingDelegationControlResourceSummaryResourceStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
