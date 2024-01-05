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

// DeployedApplicationInstallationUsageSummary Summarize usage information about an application deployed on Java servers including installation information during a specified time period. The main difference between DeployedApplicationInstallationUsageSummary and DeployedApplicationUsage is the presence of the applicationSourcePath. DeployedApplicationUsage provides only an aggregated view to the deployed applications without installation information. It therefore doesn’t distinguish between applications with the identical deployment information deployed to different paths. DeployedApplicationInstallationUsageSummary contains installation information, and it’s therefore possible to target actions.
type DeployedApplicationInstallationUsageSummary struct {

	// The internal identifier of the deployed application installation.
	ApplicationInstallationKey *string `mandatory:"true" json:"applicationInstallationKey"`

	// The internal identifier of the deployed application.
	// ApplicationKey will be identical for deployed applications with different applicationSourcePaths.
	ApplicationKey *string `mandatory:"true" json:"applicationKey"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the related fleet.
	FleetId *string `mandatory:"true" json:"fleetId"`

	// The name of the deployed application.
	ApplicationName *string `mandatory:"true" json:"applicationName"`

	// The type of the deployed application.
	ApplicationType *string `mandatory:"false" json:"applicationType"`

	// The full path to source WAR or EAR file for deployed application.
	ApplicationSourcePath *string `mandatory:"false" json:"applicationSourcePath"`

	// Whether or not the deployed application is clustered.
	IsClustered *bool `mandatory:"false" json:"isClustered"`

	// The approximate count of Java Server instances running the deployed application installations.
	ApproximateJavaServerInstanceCount *int `mandatory:"false" json:"approximateJavaServerInstanceCount"`

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

func (m DeployedApplicationInstallationUsageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeployedApplicationInstallationUsageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
