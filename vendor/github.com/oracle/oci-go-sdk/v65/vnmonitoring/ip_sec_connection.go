// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IpSecConnection A connection between a DRG and CPE. This connection consists of multiple IPSec
// tunnels. Creating this connection is one of the steps required when setting up
// a Site-to-Site VPN. For more information, see Site-to-Site VPN Overview (https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/overviewIPsec.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
type IpSecConnection struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the IPSec connection.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cpe object.
	CpeId *string `mandatory:"true" json:"cpeId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.
	DrgId *string `mandatory:"true" json:"drgId"`

	// The IPSec connection's Oracle ID (OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
	Id *string `mandatory:"true" json:"id"`

	// The IPSec connection's current state.
	LifecycleState IpSecConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Static routes to the CPE. The CIDR must not be a
	// multicast address or class E address.
	//
	// Example: `10.0.1.0/24`
	StaticRoutes []string `mandatory:"true" json:"staticRoutes"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The date and time the IPSec connection was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m IpSecConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IpSecConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIpSecConnectionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetIpSecConnectionLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IpSecConnectionLifecycleStateEnum Enum with underlying type: string
type IpSecConnectionLifecycleStateEnum string

// Set of constants representing the allowable values for IpSecConnectionLifecycleStateEnum
const (
	IpSecConnectionLifecycleStateProvisioning IpSecConnectionLifecycleStateEnum = "PROVISIONING"
	IpSecConnectionLifecycleStateAvailable    IpSecConnectionLifecycleStateEnum = "AVAILABLE"
	IpSecConnectionLifecycleStateTerminating  IpSecConnectionLifecycleStateEnum = "TERMINATING"
	IpSecConnectionLifecycleStateTerminated   IpSecConnectionLifecycleStateEnum = "TERMINATED"
)

var mappingIpSecConnectionLifecycleStateEnum = map[string]IpSecConnectionLifecycleStateEnum{
	"PROVISIONING": IpSecConnectionLifecycleStateProvisioning,
	"AVAILABLE":    IpSecConnectionLifecycleStateAvailable,
	"TERMINATING":  IpSecConnectionLifecycleStateTerminating,
	"TERMINATED":   IpSecConnectionLifecycleStateTerminated,
}

var mappingIpSecConnectionLifecycleStateEnumLowerCase = map[string]IpSecConnectionLifecycleStateEnum{
	"provisioning": IpSecConnectionLifecycleStateProvisioning,
	"available":    IpSecConnectionLifecycleStateAvailable,
	"terminating":  IpSecConnectionLifecycleStateTerminating,
	"terminated":   IpSecConnectionLifecycleStateTerminated,
}

// GetIpSecConnectionLifecycleStateEnumValues Enumerates the set of values for IpSecConnectionLifecycleStateEnum
func GetIpSecConnectionLifecycleStateEnumValues() []IpSecConnectionLifecycleStateEnum {
	values := make([]IpSecConnectionLifecycleStateEnum, 0)
	for _, v := range mappingIpSecConnectionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetIpSecConnectionLifecycleStateEnumStringValues Enumerates the set of values in String for IpSecConnectionLifecycleStateEnum
func GetIpSecConnectionLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
	}
}

// GetMappingIpSecConnectionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIpSecConnectionLifecycleStateEnum(val string) (IpSecConnectionLifecycleStateEnum, bool) {
	enum, ok := mappingIpSecConnectionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
