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

// UpdatePrivilegedApiControlDetails The data to update a PrivilegedApiControl.
type UpdatePrivilegedApiControlDetails struct {

	// Name of the privilegedApi control. Needs to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the privilegedApi control.
	Description *string `mandatory:"false" json:"description"`

	// resourceType for which the PrivilegedApiControl is applicable
	ResourceType *string `mandatory:"false" json:"resourceType"`

	// contains Resource details
	Resources []string `mandatory:"false" json:"resources"`

	// The OCID of the OCI Notification topic to publish messages related to this Delegation Control.
	NotificationTopicId *string `mandatory:"false" json:"notificationTopicId"`

	// List of user IAM group ids who can approve an privilegedApi request associated with a target resource under the governance of this privilegedApi control.
	ApproverGroupIdList []string `mandatory:"false" json:"approverGroupIdList"`

	// List of privileged operator operations. If Privileged API Managment is enabled for a resource it will be validated whether the operation done by the operator is a part of privileged operation.
	PrivilegedOperationList []PrivilegedApiDetails `mandatory:"false" json:"privilegedOperationList"`

	// Number of approvers required to approve an privilegedApi request.
	NumberOfApprovers *int `mandatory:"false" json:"numberOfApprovers"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdatePrivilegedApiControlDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdatePrivilegedApiControlDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
