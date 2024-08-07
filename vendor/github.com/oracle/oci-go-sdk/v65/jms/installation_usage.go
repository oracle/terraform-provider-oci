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

// InstallationUsage Installation usage during a specified time period.
// An installation is a collection of deployed instances of a specific Java Runtime that share the same install path.
type InstallationUsage struct {

	// The vendor of the Java Runtime that is deployed with the installation.
	JreVendor *string `mandatory:"true" json:"jreVendor"`

	// The distribution of the Java Runtime that is deployed with the installation.
	JreDistribution *string `mandatory:"true" json:"jreDistribution"`

	// The version of the Java Runtime that is deployed with the installation.
	JreVersion *string `mandatory:"true" json:"jreVersion"`

	// The file system path of the Java installation.
	Path *string `mandatory:"true" json:"path"`

	// The Operating System for the installation. Deprecated, use `operatingSystem` instead.
	Os *string `mandatory:"true" json:"os"`

	// The architecture of the operating system for the installation. Deprecated, use `operatingSystem` instead.
	Architecture *string `mandatory:"true" json:"architecture"`

	// The unique identifier for the installation of a Java Runtime at a specific path on a specific operating system.
	InstallationKey *string `mandatory:"false" json:"installationKey"`

	OperatingSystem *OperatingSystem `mandatory:"false" json:"operatingSystem"`

	// The approximate count of applications running on this installation
	ApproximateApplicationCount *int `mandatory:"false" json:"approximateApplicationCount"`

	// The approximate count of managed instances reporting this installation
	ApproximateManagedInstanceCount *int `mandatory:"false" json:"approximateManagedInstanceCount"`

	// Lower bound of the specified time period filter. JMS provides a view of the data that is _per day_. The query uses only the date element of the parameter.
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// Upper bound of the specified time period filter. JMS provides a view of the data that is _per day_. The query uses only the date element of the parameter.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// The date and time the resource was _first_ reported to JMS.
	// This is potentially _before_ the specified time period provided by the filters.
	// For example, a resource can be first reported to JMS before the start of a specified time period,
	// if it is also reported during the time period.
	TimeFirstSeen *common.SDKTime `mandatory:"false" json:"timeFirstSeen"`

	// The date and time the resource was _last_ reported to JMS.
	// This is potentially _after_ the specified time period provided by the filters.
	// For example, a resource can be last reported to JMS before the start of a specified time period,
	// if it is also reported during the time period.
	TimeLastSeen *common.SDKTime `mandatory:"false" json:"timeLastSeen"`
}

func (m InstallationUsage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstallationUsage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
