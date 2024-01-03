// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AuditTrailDimensions Details of aggregation dimensions used for summarizing audit trails.
type AuditTrailDimensions struct {

	// The location represents the source of audit records that provides documentary evidence of the sequence of activities in the target database.
	Location *string `mandatory:"false" json:"location"`

	// The current state of the audit trail.
	LifecycleState AuditTrailLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The current sub-state of the audit trail..
	Status *string `mandatory:"false" json:"status"`

	// The OCID of the Data Safe target for which the audit trail is created.
	TargetId *string `mandatory:"false" json:"targetId"`
}

func (m AuditTrailDimensions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuditTrailDimensions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAuditTrailLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAuditTrailLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
