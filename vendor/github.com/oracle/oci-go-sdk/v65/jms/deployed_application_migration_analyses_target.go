// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// DeployedApplicationMigrationAnalysesTarget The target describes the input data for deployed Java migration analyses.
// A target contains a managed instance, deployed application installation Key, sourceJdkVersion, targetJdkVersion and optional excludePackagePrefixes.
type DeployedApplicationMigrationAnalysesTarget struct {

	// The OCID of the managed instance that hosts the application for which the Java migration analyses was performed.
	ManagedInstanceId *string `mandatory:"true" json:"managedInstanceId"`

	// The unique key that identifies the deployed application's installation path that is to be used for the Java migration analyses.
	DeployedApplicationInstallationKey *string `mandatory:"true" json:"deployedApplicationInstallationKey"`

	// The JDK version the application is currently running on.
	SourceJdkVersion *string `mandatory:"true" json:"sourceJdkVersion"`

	// The JDK version against which the migration analyses was performed to identify effort required to move from source JDK.
	TargetJdkVersion *string `mandatory:"true" json:"targetJdkVersion"`

	// Excludes the packages that starts with the prefix from the migration analyses result.
	ExcludePackagePrefixes []string `mandatory:"false" json:"excludePackagePrefixes"`
}

func (m DeployedApplicationMigrationAnalysesTarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeployedApplicationMigrationAnalysesTarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
