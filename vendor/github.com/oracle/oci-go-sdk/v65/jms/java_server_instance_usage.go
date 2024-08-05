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

// JavaServerInstanceUsage Java Server instance usage during a specified time period.
type JavaServerInstanceUsage struct {

	// The internal identifier of the Java Server instance.
	ServerInstanceKey *string `mandatory:"true" json:"serverInstanceKey"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the related Fleet.
	FleetId *string `mandatory:"true" json:"fleetId"`

	// The name of the Java Server instance.
	ServerInstanceName *string `mandatory:"true" json:"serverInstanceName"`

	// The internal identifier of the related Java Server.
	ServerKey *string `mandatory:"true" json:"serverKey"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the related managed instance.
	ManagedInstanceId *string `mandatory:"true" json:"managedInstanceId"`

	// The port of the Java Server instance.
	ServerInstancePort *int `mandatory:"false" json:"serverInstancePort"`

	// The name of the Java Server.
	ServerName *string `mandatory:"false" json:"serverName"`

	// The version of the Java Server.
	ServerVersion *string `mandatory:"false" json:"serverVersion"`

	// The host name of the related managed instance.
	HostName *string `mandatory:"false" json:"hostName"`

	// The internal identifier of the related Java Runtime.
	JvmKey *string `mandatory:"false" json:"jvmKey"`

	// The vendor of the Java Runtime.
	JvmVendor *string `mandatory:"false" json:"jvmVendor"`

	// The distribution of the Java Runtime.
	JvmDistribution *string `mandatory:"false" json:"jvmDistribution"`

	// The version of the Java Runtime.
	JvmVersion *string `mandatory:"false" json:"jvmVersion"`

	// The security status of the Java Runtime.
	JvmSecurityStatus JreSecurityStatusEnum `mandatory:"false" json:"jvmSecurityStatus,omitempty"`

	// The approximate count of deployed applications in the Java Server instance.
	ApproximateDeployedApplicationCount *int `mandatory:"false" json:"approximateDeployedApplicationCount"`

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

func (m JavaServerInstanceUsage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JavaServerInstanceUsage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingJreSecurityStatusEnum(string(m.JvmSecurityStatus)); !ok && m.JvmSecurityStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JvmSecurityStatus: %s. Supported values are: %s.", m.JvmSecurityStatus, strings.Join(GetJreSecurityStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
