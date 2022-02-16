// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperatorAccessControl API
//
// Operator Access Control enables you to control the time duration and the actions an Oracle operator can perform on your Exadata Cloud@Customer infrastructure.
// Using logging service, you can view a near real-time audit report of all actions performed by an Oracle operator.
// Use the table of contents and search tool to explore the OperatorAccessControl API.
//

package operatoraccesscontrol

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// AccessRequestSummary Summary of access request.
type AccessRequestSummary struct {

	// The OCID of the access request.
	Id *string `mandatory:"true" json:"id"`

	// Comment associated with the access request.
	AccessReasonSummary *string `mandatory:"true" json:"accessReasonSummary"`

	// The OCID of the target resource associated with the access request. The operator raises an access request to get approval to
	// access the target resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// This is a system-generated identifier.
	RequestId *string `mandatory:"false" json:"requestId"`

	// The OCID of the compartment that contains the access request.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The name of the target resource.
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// resourceType for which the AccessRequest is applicable
	ResourceType ResourceTypesEnum `mandatory:"false" json:"resourceType,omitempty"`

	// The current state of the AccessRequest.
	LifecycleState AccessRequestLifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Time when the access request was created by the operator user in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.Example: '2020-05-22T21:10:29.600Z'
	TimeOfCreation *common.SDKTime `mandatory:"false" json:"timeOfCreation"`

	// Time when the access request was last modified in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.Example: '2020-05-22T21:10:29.600Z'
	TimeOfModification *common.SDKTime `mandatory:"false" json:"timeOfModification"`

	// The time when access request is scheduled to be approved in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.Example: '2020-05-22T21:10:29.600Z'
	TimeOfUserCreation *common.SDKTime `mandatory:"false" json:"timeOfUserCreation"`

	// Duration in hours for which access is sought on the target resource.
	Duration *int `mandatory:"false" json:"duration"`

	// Duration in hours for which extension access is sought on the target resource.
	ExtendDuration *int `mandatory:"false" json:"extendDuration"`

	// Priority assigned to the access request by the operator
	Severity AccessRequestSeveritiesEnum `mandatory:"false" json:"severity,omitempty"`

	// Whether the access request was automatically approved.
	IsAutoApproved *bool `mandatory:"false" json:"isAutoApproved"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m AccessRequestSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AccessRequestSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingResourceTypesEnum(string(m.ResourceType)); !ok && m.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", m.ResourceType, strings.Join(GetResourceTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAccessRequestLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAccessRequestLifecycleStatesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAccessRequestSeveritiesEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetAccessRequestSeveritiesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
