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

// ManagedInstanceUsage Managed instance usage during a specified time period.
// An entity that emits usage events to Java Management Service (JMS) is represented as a managed instance.
// A managed instance has a unique identity which is used by JMS to distinguish it from other managed instances.
// Currently, JMS supports only one kind of managed instance, a Management Agent.
type ManagedInstanceUsage struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related managed instance.
	ManagedInstanceId *string `mandatory:"true" json:"managedInstanceId"`

	// The type of the source of events.
	ManagedInstanceType ManagedInstanceTypeEnum `mandatory:"true" json:"managedInstanceType"`

	// The hostname of the managed instance (if applicable).
	Hostname *string `mandatory:"false" json:"hostname"`

	// The host OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related managed instance.
	HostId *string `mandatory:"false" json:"hostId"`

	OperatingSystem *OperatingSystem `mandatory:"false" json:"operatingSystem"`

	Agent *Agent `mandatory:"false" json:"agent"`

	// The approximate count of applications reported by this managed instance.
	ApproximateApplicationCount *int `mandatory:"false" json:"approximateApplicationCount"`

	// The approximate count of installations reported by this managed instance.
	ApproximateInstallationCount *int `mandatory:"false" json:"approximateInstallationCount"`

	// The approximate count of Java Runtimes reported by this managed instance.
	ApproximateJreCount *int `mandatory:"false" json:"approximateJreCount"`

	// DRS file status
	DrsFileStatus DrsFileStatusEnum `mandatory:"false" json:"drsFileStatus,omitempty"`

	// Comma separated list of user names that invoked applications within this managed instance.
	ApplicationInvokedBy *string `mandatory:"false" json:"applicationInvokedBy"`

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

func (m ManagedInstanceUsage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedInstanceUsage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingManagedInstanceTypeEnum(string(m.ManagedInstanceType)); !ok && m.ManagedInstanceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagedInstanceType: %s. Supported values are: %s.", m.ManagedInstanceType, strings.Join(GetManagedInstanceTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDrsFileStatusEnum(string(m.DrsFileStatus)); !ok && m.DrsFileStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DrsFileStatus: %s. Supported values are: %s.", m.DrsFileStatus, strings.Join(GetDrsFileStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
