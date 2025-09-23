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

// UncorrelatedPackageApplicationUsageSummary Summary of an application where a given libary was detected. Contains the count of managed instances where the package was detected with this application.
type UncorrelatedPackageApplicationUsageSummary struct {

	// The internal identifier of a Java application.
	ApplicationKey *string `mandatory:"true" json:"applicationKey"`

	// The displayed name of the Java application.
	ApplicationName *string `mandatory:"true" json:"applicationName"`

	// The count of managed instances wherein the specified library was detected.
	ManagedInstanceCount *int `mandatory:"true" json:"managedInstanceCount"`

	// The date and time a library or Java package was _last_ detected in a dynamic library scan.
	LastDetectedDynamically *common.SDKTime `mandatory:"true" json:"lastDetectedDynamically"`
}

func (m UncorrelatedPackageApplicationUsageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UncorrelatedPackageApplicationUsageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
