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

// DelegatedResourceAccessRequestHistorySummary Summary of Delegated Resource Access Request status changes.
type DelegatedResourceAccessRequestHistorySummary struct {

	// The current lifecycle state of the Delegated Resource Access Request.
	LifecycleState DelegatedResourceAccessRequestLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The current status of the Delegated Resource Access Request.
	RequestStatus DelegatedResourceAccessRequestStatusEnum `mandatory:"false" json:"requestStatus,omitempty"`

	// ID of user who modified the Delegated Resource Access Request. For operator, this field is "Operator".
	UserId *string `mandatory:"false" json:"userId"`

	// Comment about the status change.
	Comment *string `mandatory:"false" json:"comment"`

	// Time when the respective action happened in RFC 3339 (https://tools.ietf.org/html/rfc3339)timestamp format, e.g. '2020-05-22T21:10:29.600Z'.
	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp"`
}

func (m DelegatedResourceAccessRequestHistorySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DelegatedResourceAccessRequestHistorySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDelegatedResourceAccessRequestLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDelegatedResourceAccessRequestLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDelegatedResourceAccessRequestStatusEnum(string(m.RequestStatus)); !ok && m.RequestStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RequestStatus: %s. Supported values are: %s.", m.RequestStatus, strings.Join(GetDelegatedResourceAccessRequestStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
