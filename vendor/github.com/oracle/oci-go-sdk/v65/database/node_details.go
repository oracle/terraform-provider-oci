// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NodeDetails Node details associated with a network.
type NodeDetails struct {

	// The node host name.
	Hostname *string `mandatory:"true" json:"hostname"`

	// The node IP address.
	Ip *string `mandatory:"true" json:"ip"`

	// The node virtual IP (VIP) host name.
	VipHostname *string `mandatory:"false" json:"vipHostname"`

	// The node virtual IP (VIP) address.
	Vip *string `mandatory:"false" json:"vip"`

	// The current state of the VM cluster network nodes.
	// CREATING - The resource is being created
	// REQUIRES_VALIDATION - The resource is created and may not be usable until it is validated.
	// VALIDATING - The resource is being validated and not available to use.
	// VALIDATED - The resource is validated and is available for consumption by VM cluster.
	// VALIDATION_FAILED - The resource validation has failed and might require user input to be corrected.
	// UPDATING - The resource is being updated and not available to use.
	// ALLOCATED - The resource is currently being used by VM cluster.
	// TERMINATING - The resource is being deleted and not available to use.
	// TERMINATED - The resource is deleted and unavailable.
	// FAILED - The resource is in a failed state due to validation or other errors.
	LifecycleState NodeDetailsLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The Db server associated with the node.
	DbServerId *string `mandatory:"false" json:"dbServerId"`
}

func (m NodeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NodeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingNodeDetailsLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetNodeDetailsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NodeDetailsLifecycleStateEnum Enum with underlying type: string
type NodeDetailsLifecycleStateEnum string

// Set of constants representing the allowable values for NodeDetailsLifecycleStateEnum
const (
	NodeDetailsLifecycleStateCreating           NodeDetailsLifecycleStateEnum = "CREATING"
	NodeDetailsLifecycleStateRequiresValidation NodeDetailsLifecycleStateEnum = "REQUIRES_VALIDATION"
	NodeDetailsLifecycleStateValidating         NodeDetailsLifecycleStateEnum = "VALIDATING"
	NodeDetailsLifecycleStateValidated          NodeDetailsLifecycleStateEnum = "VALIDATED"
	NodeDetailsLifecycleStateValidationFailed   NodeDetailsLifecycleStateEnum = "VALIDATION_FAILED"
	NodeDetailsLifecycleStateUpdating           NodeDetailsLifecycleStateEnum = "UPDATING"
	NodeDetailsLifecycleStateAllocated          NodeDetailsLifecycleStateEnum = "ALLOCATED"
	NodeDetailsLifecycleStateTerminating        NodeDetailsLifecycleStateEnum = "TERMINATING"
	NodeDetailsLifecycleStateTerminated         NodeDetailsLifecycleStateEnum = "TERMINATED"
	NodeDetailsLifecycleStateFailed             NodeDetailsLifecycleStateEnum = "FAILED"
)

var mappingNodeDetailsLifecycleStateEnum = map[string]NodeDetailsLifecycleStateEnum{
	"CREATING":            NodeDetailsLifecycleStateCreating,
	"REQUIRES_VALIDATION": NodeDetailsLifecycleStateRequiresValidation,
	"VALIDATING":          NodeDetailsLifecycleStateValidating,
	"VALIDATED":           NodeDetailsLifecycleStateValidated,
	"VALIDATION_FAILED":   NodeDetailsLifecycleStateValidationFailed,
	"UPDATING":            NodeDetailsLifecycleStateUpdating,
	"ALLOCATED":           NodeDetailsLifecycleStateAllocated,
	"TERMINATING":         NodeDetailsLifecycleStateTerminating,
	"TERMINATED":          NodeDetailsLifecycleStateTerminated,
	"FAILED":              NodeDetailsLifecycleStateFailed,
}

var mappingNodeDetailsLifecycleStateEnumLowerCase = map[string]NodeDetailsLifecycleStateEnum{
	"creating":            NodeDetailsLifecycleStateCreating,
	"requires_validation": NodeDetailsLifecycleStateRequiresValidation,
	"validating":          NodeDetailsLifecycleStateValidating,
	"validated":           NodeDetailsLifecycleStateValidated,
	"validation_failed":   NodeDetailsLifecycleStateValidationFailed,
	"updating":            NodeDetailsLifecycleStateUpdating,
	"allocated":           NodeDetailsLifecycleStateAllocated,
	"terminating":         NodeDetailsLifecycleStateTerminating,
	"terminated":          NodeDetailsLifecycleStateTerminated,
	"failed":              NodeDetailsLifecycleStateFailed,
}

// GetNodeDetailsLifecycleStateEnumValues Enumerates the set of values for NodeDetailsLifecycleStateEnum
func GetNodeDetailsLifecycleStateEnumValues() []NodeDetailsLifecycleStateEnum {
	values := make([]NodeDetailsLifecycleStateEnum, 0)
	for _, v := range mappingNodeDetailsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetNodeDetailsLifecycleStateEnumStringValues Enumerates the set of values in String for NodeDetailsLifecycleStateEnum
func GetNodeDetailsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"REQUIRES_VALIDATION",
		"VALIDATING",
		"VALIDATED",
		"VALIDATION_FAILED",
		"UPDATING",
		"ALLOCATED",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
	}
}

// GetMappingNodeDetailsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNodeDetailsLifecycleStateEnum(val string) (NodeDetailsLifecycleStateEnum, bool) {
	enum, ok := mappingNodeDetailsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
