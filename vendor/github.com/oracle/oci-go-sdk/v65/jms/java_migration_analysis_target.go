// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JavaMigrationAnalysisTarget The target describes the input data for Java migration analysis.
// A target contains a managed instance, application Installation Key, sourceJdkVersion, and targetJdkVersion.
type JavaMigrationAnalysisTarget struct {

	// The OCID of the managed instance that hosts the application for which the Java migration analysis was performed.
	ManagedInstanceId *string `mandatory:"true" json:"managedInstanceId"`

	// The unique key that identifies the application's installation path that is to be used for the Java migration analysis.
	ApplicationInstallationKey *string `mandatory:"true" json:"applicationInstallationKey"`

	// The JDK version the application is currently running on.
	SourceJdkVersion *string `mandatory:"true" json:"sourceJdkVersion"`

	// The JDK version against which the migration analysis was performed to identify effort required to move from source JDK.
	TargetJdkVersion *string `mandatory:"true" json:"targetJdkVersion"`
}

func (m JavaMigrationAnalysisTarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JavaMigrationAnalysisTarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
