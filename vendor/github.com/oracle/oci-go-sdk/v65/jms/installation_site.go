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

// InstallationSite Installation site of a Java Runtime.
// An installation site is a Java Runtime installed at a specific path on a managed instance.
type InstallationSite struct {

	// The unique identifier for the installation of a Java Runtime at a specific path on a specific operating system.
	InstallationKey *string `mandatory:"true" json:"installationKey"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the related managed instance.
	ManagedInstanceId *string `mandatory:"true" json:"managedInstanceId"`

	Jre *JavaRuntimeId `mandatory:"true" json:"jre"`

	// The file system path of the installation.
	Path *string `mandatory:"true" json:"path"`

	OperatingSystem *OperatingSystem `mandatory:"true" json:"operatingSystem"`

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

	// The type of the source of events.
	ManagedInstanceType ManagedInstanceTypeEnum `mandatory:"false" json:"managedInstanceType,omitempty"`

	// The hostname of the managed instance (if applicable).
	Hostname *string `mandatory:"false" json:"hostname"`
}

func (m InstallationSite) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstallationSite) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingManagedInstanceTypeEnum(string(m.ManagedInstanceType)); !ok && m.ManagedInstanceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagedInstanceType: %s. Supported values are: %s.", m.ManagedInstanceType, strings.Join(GetManagedInstanceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
