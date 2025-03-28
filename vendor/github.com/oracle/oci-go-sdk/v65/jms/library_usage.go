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

// LibraryUsage Library usage during a specified time period.
type LibraryUsage struct {

	// The internal identifier of the library.
	LibraryKey *string `mandatory:"true" json:"libraryKey"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related fleet.
	FleetId *string `mandatory:"true" json:"fleetId"`

	// The name of the library.
	LibraryName *string `mandatory:"true" json:"libraryName"`

	// The version of the library.
	LibraryVersion *string `mandatory:"false" json:"libraryVersion"`

	// The Common Vulnerabilities and Exposures (CVE) ID.
	CveId *string `mandatory:"false" json:"cveId"`

	// The Common Vulnerability Scoring System (CVSS) score.
	CvssScore *float32 `mandatory:"false" json:"cvssScore"`

	// The approximate count of applications using the library.
	ApproximateApplicationCount *int `mandatory:"false" json:"approximateApplicationCount"`

	// The approximate count of Java Server instances using the library.
	ApproximateJavaServerInstanceCount *int `mandatory:"false" json:"approximateJavaServerInstanceCount"`

	// The approximate count of deployed applications using the library.
	ApproximateDeployedApplicationCount *int `mandatory:"false" json:"approximateDeployedApplicationCount"`

	// The approximate count of managed instances using the library.
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

	// The date and time of the last CVEs refresh was completed.
	TimeLastCveRefreshed *common.SDKTime `mandatory:"false" json:"timeLastCveRefreshed"`
}

func (m LibraryUsage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LibraryUsage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
