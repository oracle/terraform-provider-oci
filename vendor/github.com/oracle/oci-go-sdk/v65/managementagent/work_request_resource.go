// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.cloud.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WorkRequestResource A resource created or operated on by a work request.
type WorkRequestResource struct {

	// The resource type the work request affects.
	EntityType *string `mandatory:"true" json:"entityType"`

	// The way in which this resource is affected by the work tracked in the work request.
	// A resource being created, updated, or deleted will remain in the IN_PROGRESS state until
	// work is complete for that resource at which point it will transition to CREATED, UPDATED,
	// or DELETED, respectively.
	ActionType ActionTypesEnum `mandatory:"true" json:"actionType"`

	// The identifier of the resource the work request affects.
	Identifier *string `mandatory:"true" json:"identifier"`

	// The identifier of the source the work request is requesting.
	SourceId *string `mandatory:"false" json:"sourceId"`

	// The name of the source the work request is requesting.
	SourceName *string `mandatory:"false" json:"sourceName"`

	// The version of the source the work request is requesting.
	SourceVersion *string `mandatory:"false" json:"sourceVersion"`

	// The URI path that the user can do a GET on to access the resource metadata
	EntityUri *string `mandatory:"false" json:"entityUri"`

	// The date and time the request was created, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339),
	// section 5.6.
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`

	// The date and time the request was started, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339),
	// section 5.6.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the request was finished, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339),
	// section 5.6.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// Additional metadata about the resource that has been operated upon by
	// this work request. For WorkRequests operationType WORK_DELIVERY the metadata will contain: workDeliveryStatus
	// indicating the status of the work delivery item as a WorkDeliveryStatus value, workSubmissionKey the WorkSubmission request id,
	//  and workSubmissionDetails containing any details of result
	Metadata *interface{} `mandatory:"false" json:"metadata"`
}

func (m WorkRequestResource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkRequestResource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingActionTypesEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetActionTypesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
