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

// JreUsage Java Runtime usage during a specified time period. A Java Runtime is identified by its vendor and version.
type JreUsage struct {

	// The vendor of the Java Runtime.
	Vendor *string `mandatory:"true" json:"vendor"`

	// The distribution of a Java Runtime is the name of the lineage of product to which it belongs, for example _Java(TM) SE Runtime Environment_.
	Distribution *string `mandatory:"true" json:"distribution"`

	// The version of the Java Runtime.
	Version *string `mandatory:"true" json:"version"`

	// The internal identifier of the Java Runtime.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related fleet.
	FleetId *string `mandatory:"false" json:"fleetId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related managed instance. This property value is present only for /listJreUsage.
	ManagedInstanceId *string `mandatory:"false" json:"managedInstanceId"`

	// The security status of the Java Runtime.
	SecurityStatus JreSecurityStatusEnum `mandatory:"false" json:"securityStatus,omitempty"`

	// The release date of the Java Runtime (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	ReleaseDate *common.SDKTime `mandatory:"false" json:"releaseDate"`

	// The End of Support Life (EOSL) date of the Java Runtime (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	EndOfSupportLifeDate *common.SDKTime `mandatory:"false" json:"endOfSupportLifeDate"`

	// The number of days since this release has been under the security baseline.
	DaysUnderSecurityBaseline *int `mandatory:"false" json:"daysUnderSecurityBaseline"`

	// The operating systems that have this Java Runtime installed.
	OperatingSystems []OperatingSystem `mandatory:"false" json:"operatingSystems"`

	// The approximate count of installations that are installations of this Java Runtime.
	ApproximateInstallationCount *int `mandatory:"false" json:"approximateInstallationCount"`

	// The approximate count of the applications running on this Java Runtime.
	ApproximateApplicationCount *int `mandatory:"false" json:"approximateApplicationCount"`

	// The approximate count of the managed instances that report this Java Runtime.
	ApproximateManagedInstanceCount *int `mandatory:"false" json:"approximateManagedInstanceCount"`

	// The approximate count of work requests working on this Java Runtime.
	ApproximatePendingWorkRequestCount *int `mandatory:"false" json:"approximatePendingWorkRequestCount"`

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

func (m JreUsage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JreUsage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingJreSecurityStatusEnum(string(m.SecurityStatus)); !ok && m.SecurityStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityStatus: %s. Supported values are: %s.", m.SecurityStatus, strings.Join(GetJreSecurityStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
