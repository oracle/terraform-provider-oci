// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateDrPlanGroupDetails The details for updating a group in a DR plan.
type UpdateDrPlanGroupDetails struct {

	// The unique id of the group. Must not be modified by user.
	// Example: `sgid1.group..uniqueID`
	Id *string `mandatory:"false" json:"id"`

	// The display name of the group.
	// Example: `My_GROUP_3 - EBS Start`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The group type.
	// Example: `BUILT_IN`
	Type DrPlanGroupTypeEnum `mandatory:"false" json:"type,omitempty"`

	// A flag indicating whether this group should be enabled for execution.
	// This flag is only applicable to the `USER_DEFINED_PAUSE` group. The flag should be null for the remaining group types.
	// Example: `true`
	IsPauseEnabled *bool `mandatory:"false" json:"isPauseEnabled"`

	// The list of steps in this group.
	Steps []UpdateDrPlanStepDetails `mandatory:"false" json:"steps"`
}

func (m UpdateDrPlanGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDrPlanGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDrPlanGroupTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetDrPlanGroupTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
