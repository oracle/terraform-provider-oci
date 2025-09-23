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

// LibraryManagedInstanceUsageSummary Summary of a managed instance where a given library was detected statically and/or dynamically. Contains the count of applications where the library was detected.
type LibraryManagedInstanceUsageSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related managed instance.
	ManagedInstanceId *string `mandatory:"true" json:"managedInstanceId"`

	// The hostname of the managed instance.
	Hostname *string `mandatory:"true" json:"hostname"`

	// The count of applications where the specified library was detected.
	ApplicationCount *int `mandatory:"true" json:"applicationCount"`

	// The date and time a library or Java package was _last_ detected in a dynamic library scan.
	LastDetectedDynamically *common.SDKTime `mandatory:"true" json:"lastDetectedDynamically"`

	// The timestamp of the first time the specified library was detected in classpath.
	FirstSeenInClasspath *common.SDKTime `mandatory:"true" json:"firstSeenInClasspath"`

	// The timestamp of the last time the specified library was detected in classpath.
	LastSeenInClasspath *common.SDKTime `mandatory:"true" json:"lastSeenInClasspath"`
}

func (m LibraryManagedInstanceUsageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LibraryManagedInstanceUsageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
