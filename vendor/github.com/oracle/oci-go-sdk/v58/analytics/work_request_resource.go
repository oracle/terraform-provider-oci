// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// WorkRequestResource The representation of WorkRequestResource
type WorkRequestResource struct {

	// The way in which this resource was affected by this work request.
	ActionResult WorkRequestActionResultEnum `mandatory:"true" json:"actionResult"`

	// The type of the resource the work request is affecting.
	ResourceType WorkRequestResourceTypeEnum `mandatory:"true" json:"resourceType"`

	// The OCID of the resource the work request is affecting.
	Identifier *string `mandatory:"true" json:"identifier"`

	// The URI of the affected resource.
	ResourceUri *string `mandatory:"true" json:"resourceUri"`

	// Additional metadata of the resource.
	Metadata map[string]string `mandatory:"false" json:"metadata"`
}

func (m WorkRequestResource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkRequestResource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWorkRequestActionResultEnum(string(m.ActionResult)); !ok && m.ActionResult != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionResult: %s. Supported values are: %s.", m.ActionResult, strings.Join(GetWorkRequestActionResultEnumStringValues(), ",")))
	}
	if _, ok := GetMappingWorkRequestResourceTypeEnum(string(m.ResourceType)); !ok && m.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", m.ResourceType, strings.Join(GetWorkRequestResourceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
