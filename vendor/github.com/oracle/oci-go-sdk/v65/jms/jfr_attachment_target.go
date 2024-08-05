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

// JfrAttachmentTarget The target to collect JFR data. A target is a managed instance, with options to further limit to specific application and/or Java Runtime.
// When the applicationKey isn't specified, then all applications are selected.
// When the jreKey isn't specified, then all supported Java Runtime versions are selected.
// When the applicationInstallationKey isn't specified, then all application installations are selected.
// Keys applicationKey and applicationInstallationKey are mutually exclusive.
type JfrAttachmentTarget struct {

	// OCID of the Managed Instance to collect JFR data.
	ManagedInstanceId *string `mandatory:"true" json:"managedInstanceId"`

	// Unique key that identifies the application for JFR data collection.
	ApplicationKey *string `mandatory:"false" json:"applicationKey"`

	// Unique key that identifies the application installation for JFR data collection.
	ApplicationInstallationKey *string `mandatory:"false" json:"applicationInstallationKey"`

	// Unique key that identify the JVM for JFR data collection.
	JreKey *string `mandatory:"false" json:"jreKey"`
}

func (m JfrAttachmentTarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JfrAttachmentTarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
