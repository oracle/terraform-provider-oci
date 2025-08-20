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

// CreatePrivilegedApiRequestDetails An Oracle operator raises privilegedApi request when they need access to any infrastructure resource governed by PrivilegedApi Access Control.
//
//	The privilegedApi request identifies the target resource and the set of operator actions. Access request handling depends upon the Operator Control
//	that governs the target resource, and the set of operator actions listed for approval in the access request. If all of the operator actions
//	listed in the privilegedApi request are in the pre-approved list in the PrivilegedApi Control that governs the target resource, then the privilegedApi request is
//	automatically approved. If not, then the privilegedApi request requires explicit approval from the approver group specified by the PrivilegedApi Control governing the target resource.
//
// You can approve or reject an privilegedApi request. You can also revoke the approval of an already approved privilegedApi request. While creating an access request,
//
//	the operator specifies the duration of access. You have the option to approve the entire duration or reduce or even increase the time duration.
//	An operator can also request for an extension. The approval for such an extension is processed the same way the original privilegedApi request was processed.
type CreatePrivilegedApiRequestDetails struct {

	// Summary comment by the operator creating the access request.
	ReasonSummary *string `mandatory:"true" json:"reasonSummary"`

	// The OCID of the target resource associated with the access request. The operator raises an access request to get approval to access the target resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// List of api names, attributes for which approval is sought by the user.
	PrivilegedOperationList []PrivilegedApiRequestOperationDetails `mandatory:"true" json:"privilegedOperationList"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The subresource names requested for approval.
	SubResourceNameList []string `mandatory:"false" json:"subResourceNameList"`

	// The OCID of the OCI Notification topic to publish messages related to this Privileged Api Request.
	NotificationTopicId *string `mandatory:"false" json:"notificationTopicId"`

	// Reason in detail for which the operator is requesting access on the target resource.
	ReasonDetail *string `mandatory:"false" json:"reasonDetail"`

	// Priority assigned to the access request by the operator
	Severity PrivilegedApiRequestSeverityEnum `mandatory:"false" json:"severity,omitempty"`

	// Duration in hours for which access is sought on the target resource.
	DurationInHrs *int `mandatory:"false" json:"durationInHrs"`

	// A list of ticket numbers related to this Privileged Api Access Request, e.g. Service Request (SR) number and JIRA ticket number.
	TicketNumbers []string `mandatory:"false" json:"ticketNumbers"`

	// Time in future when the user for the privilegedApi request needs to be created in RFC 3339 (https://tools.ietf.org/html/rfc3339)timestamp format. Example: '2020-05-22T21:10:29.600Z'
	TimeRequestedForFutureAccess *common.SDKTime `mandatory:"false" json:"timeRequestedForFutureAccess"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreatePrivilegedApiRequestDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePrivilegedApiRequestDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPrivilegedApiRequestSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetPrivilegedApiRequestSeverityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
