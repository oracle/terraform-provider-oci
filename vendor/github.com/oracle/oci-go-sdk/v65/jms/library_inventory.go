// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Fleets API
//
// The APIs for the Fleet Management (https://docs.oracle.com/en-us/iaas/jms/doc/fleet-management.html) feature of Java Management Service to monitor and manage the usage of Java in your enterprise. Use these APIs to manage fleets, configure managed instances to report to fleets, and gain insights into the Java workloads running on these instances by carrying out basic and advanced features.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LibraryInventory Inventory of libraries in a fleet during a specified time period.
type LibraryInventory struct {

	// The count of libraries which are detected statically.
	StaticallyDetectedLibraryCount *int `mandatory:"true" json:"staticallyDetectedLibraryCount"`

	// The count of libraries which are detected dynamically.
	DynamicallyDetectedLibraryCount *int `mandatory:"true" json:"dynamicallyDetectedLibraryCount"`

	// The count of packages which are detected but not correlated to any libraries.
	UncorrelatedPackageCount *int `mandatory:"true" json:"uncorrelatedPackageCount"`

	// The count of libraries with high severity vulnerabilities.
	HighSeverityLibraryCount *int `mandatory:"true" json:"highSeverityLibraryCount"`

	// The count of libraries with medium severity vulnerabilities.
	MediumSeverityLibraryCount *int `mandatory:"true" json:"mediumSeverityLibraryCount"`

	// The count of libraries with low severity vulnerabilities.
	LowSeverityLibraryCount *int `mandatory:"true" json:"lowSeverityLibraryCount"`
}

func (m LibraryInventory) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LibraryInventory) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
