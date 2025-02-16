// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LibraryDetails The details required to patch a library
type LibraryDetails struct {

	// The OCID of the Compute Cluster.
	ClusterId *string `mandatory:"true" json:"clusterId"`

	// The OCID of a compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The Library type.
	Type LibraryTypeEnum `mandatory:"true" json:"type"`

	// The operation (install/uninstall) for patch library update.
	Operation LibraryOperationEnum `mandatory:"true" json:"operation"`

	// A user-friendly name. It does not have to be unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The repository for referenced Library.
	Repository *string `mandatory:"false" json:"repository"`

	// The coordinate for referenced Library.
	Coordinate *string `mandatory:"false" json:"coordinate"`

	// List of exclusions for library installation.
	Exclusions map[string]string `mandatory:"false" json:"exclusions"`

	// The pipIndexUrl for referenced Library.
	PipIndexUrl *string `mandatory:"false" json:"pipIndexUrl"`

	// The pythonPackage for referenced Library.
	PythonPackage *string `mandatory:"false" json:"pythonPackage"`

	// The path for referenced Library.
	Path *string `mandatory:"false" json:"path"`

	// The date and time the resource was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2018-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the resource was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2018-04-03T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m LibraryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LibraryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLibraryTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetLibraryTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLibraryOperationEnum(string(m.Operation)); !ok && m.Operation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operation: %s. Supported values are: %s.", m.Operation, strings.Join(GetLibraryOperationEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
