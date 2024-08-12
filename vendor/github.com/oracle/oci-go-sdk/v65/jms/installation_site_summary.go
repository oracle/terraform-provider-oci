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

// InstallationSiteSummary Installation site of a Java Runtime.
// An installation site is a Java Runtime installed at a specific path on a managed instance.
type InstallationSiteSummary struct {

	// The unique identifier for the installation of Java Runtime at a specific path on a specific operating system.
	InstallationKey *string `mandatory:"true" json:"installationKey"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the related managed instance.
	ManagedInstanceId *string `mandatory:"true" json:"managedInstanceId"`

	Jre *JavaRuntimeId `mandatory:"false" json:"jre"`

	// The security status of the Java Runtime.
	SecurityStatus JreSecurityStatusEnum `mandatory:"false" json:"securityStatus,omitempty"`

	// The file system path of the installation.
	Path *string `mandatory:"false" json:"path"`

	OperatingSystem *OperatingSystem `mandatory:"false" json:"operatingSystem"`

	// The approximate count of applications running on this installation
	ApproximateApplicationCount *int `mandatory:"false" json:"approximateApplicationCount"`

	// The date and time the resource was _last_ reported to JMS.
	// This is potentially _after_ the specified time period provided by the filters.
	// For example, a resource can be last reported to JMS before the start of a specified time period,
	// if it is also reported during the time period.
	TimeLastSeen *common.SDKTime `mandatory:"false" json:"timeLastSeen"`

	// The list of operations that are blocklisted.
	Blocklist []BlocklistEntry `mandatory:"false" json:"blocklist"`

	// The lifecycle state of the installation site.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m InstallationSiteSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstallationSiteSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingJreSecurityStatusEnum(string(m.SecurityStatus)); !ok && m.SecurityStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityStatus: %s. Supported values are: %s.", m.SecurityStatus, strings.Join(GetJreSecurityStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
