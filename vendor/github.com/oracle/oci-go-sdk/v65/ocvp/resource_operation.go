// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResourceOperation Status of the operation performed on the SDDC / Clster / ESXi host.
type ResourceOperation struct {

	// Name of the operation.
	OperationName *string `mandatory:"true" json:"operationName"`

	// The operation status.
	OperationStatus ResourceOperationStatusEnum `mandatory:"true" json:"operationStatus"`

	// The date and time the operation started at `IN_PROGRESS` state,
	// in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeOperationStarted *common.SDKTime `mandatory:"true" json:"timeOperationStarted"`

	// The date and time the operation reached a terminal state, either `FAILED` OR
	// `SUCCEEDED`. Format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeOperationFinished *common.SDKTime `mandatory:"false" json:"timeOperationFinished"`
}

func (m ResourceOperation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResourceOperation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResourceOperationStatusEnum(string(m.OperationStatus)); !ok && m.OperationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationStatus: %s. Supported values are: %s.", m.OperationStatus, strings.Join(GetResourceOperationStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
