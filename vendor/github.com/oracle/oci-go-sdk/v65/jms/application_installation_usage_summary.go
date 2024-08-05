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

// ApplicationInstallationUsageSummary Summarizes application installation usage information during a specified time period. The main difference between ApplicationUsage and ApplicationInstallationUsageSummary is the presence of installation information. ApplicationUsage provides only aggregated information for an application regardless of the installation paths. Therefore, two different applications with the same application name installed in two different paths will be aggregated to a single application. This aggregation makes it difficult to focus actions to single application installed on a known path.
// An application installation is independent of the Java Runtime on which it's running or the Managed Instance where it's installed.
type ApplicationInstallationUsageSummary struct {

	// An internal identifier for the application installation that is unique to a Fleet.
	ApplicationInstallationKey *string `mandatory:"true" json:"applicationInstallationKey"`

	// An internal identifier for the application that is unique to a Fleet.
	// ApplicationKey will be identical for applications with different installation information.
	ApplicationKey *string `mandatory:"true" json:"applicationKey"`

	// The name of the application.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The type of the application, denoted by how the application was started.
	ApplicationType *string `mandatory:"true" json:"applicationType"`

	// The full path on which the application installation was detected.
	InstallationPath *string `mandatory:"false" json:"installationPath"`

	// List of full paths where the application last searched for classes.
	// Contains full paths to all items from module-list and class path list.
	FullClassPath []string `mandatory:"false" json:"fullClassPath"`

	// The operating systems running this application.
	OperatingSystems []OperatingSystem `mandatory:"false" json:"operatingSystems"`

	// The approximate count of installations running this application.
	ApproximateInstallationCount *int `mandatory:"false" json:"approximateInstallationCount"`

	// The approximate count of Java Runtimes running this application.
	ApproximateJreCount *int `mandatory:"false" json:"approximateJreCount"`

	// The approximate count of managed instances reporting this application.
	ApproximateManagedInstanceCount *int `mandatory:"false" json:"approximateManagedInstanceCount"`

	// The approximate count of libraries in this application.
	ApproximateLibraryCount *int `mandatory:"false" json:"approximateLibraryCount"`

	// Comma separated list of user names that invoked application installations.
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

func (m ApplicationInstallationUsageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApplicationInstallationUsageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
