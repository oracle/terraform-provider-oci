// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle API Access Control
//
// This service is used to restrict the control plane service apis; so that everybody won't be
// able to access those apis.
// There are two main resouces defined as a part of this service
// 1. PrivilegedApiControl: This is created by the customer which defines which service apis are
//    controlled and who can access it.
// 2. PrivilegedApiRequest: This is a request object again created by the customer operators who           seek access to those privileged apis. After a request is obtained based on the                       PrivilegedAccessControl for which the api belongs to, either it can be approved so that the          requested person can execute the service apis or it will wait for the customer to approve it.
//

package apiaccesscontrol

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PrivilegedApiRequestSummary Summary of access request.
type PrivilegedApiRequestSummary struct {

	// The OCID of the access request.
	Id *string `mandatory:"true" json:"id"`

	// Comment associated with the privilegedApi request.
	ReasonSummary *string `mandatory:"true" json:"reasonSummary"`

	// The OCID of the target resource associated with the privilegedApi request. The operator raises an privilegedApi request to get approval to access the target resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// Name of the privilegedApi control. The name must be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// This is a system-generated identifier to identity a Request in human readable form in the form of REQYYYYMMDD<number>.
	RequestId *string `mandatory:"false" json:"requestId"`

	// The OCID of the compartment that contains the privilegedApi request.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// resourceName for which the PrivilegedApiRequest is applicable
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// The subresource names requested for approval.
	SubResourceNameList []string `mandatory:"false" json:"subResourceNameList"`

	// resourceType for which the PrivilegedApiRequest is applicable
	ResourceType *string `mandatory:"false" json:"resourceType"`

	// List of api names, attributes for which approval is sought by the user.
	PrivilegedOperationList []PrivilegedApiRequestOperationDetails `mandatory:"false" json:"privilegedOperationList"`

	// The current state of the Access Request.
	State PrivilegedApiRequestStateEnum `mandatory:"false" json:"state,omitempty"`

	// The current state of the PrivilegedApiRequest.
	LifecycleState PrivilegedApiRequestLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// More in detail about the lifeCycleState.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Time when the privilegedApi request was created by the operator user in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.Example: '2020-05-22T21:10:29.600Z'
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time when the privilegedApi request was last modified in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.Example: '2020-05-22T21:10:29.600Z'
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Duration in hours for which access is sought on the target resource.
	DurationInHrs *int `mandatory:"false" json:"durationInHrs"`

	// Priority assigned to the privilegedApi request by the operator
	Severity PrivilegedApiRequestSeverityEnum `mandatory:"false" json:"severity,omitempty"`

	// Time in future when the user for the access request needs to be created in RFC 3339 (https://tools.ietf.org/html/rfc3339)timestamp format. Example: '2020-05-22T21:10:29.600Z'
	TimeRequestedForFutureAccess *common.SDKTime `mandatory:"false" json:"timeRequestedForFutureAccess"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m PrivilegedApiRequestSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrivilegedApiRequestSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPrivilegedApiRequestStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetPrivilegedApiRequestStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPrivilegedApiRequestLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPrivilegedApiRequestLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPrivilegedApiRequestSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetPrivilegedApiRequestSeverityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
