// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ApplicationVip Details of an application virtual IP (VIP) address.
type ApplicationVip struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the application virtual IP (VIP) address.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the cloud VM cluster associated with the application virtual IP (VIP) address.
	CloudVmClusterId *string `mandatory:"true" json:"cloudVmClusterId"`

	// The hostname of the application virtual IP (VIP) address.
	HostnameLabel *string `mandatory:"true" json:"hostnameLabel"`

	// The current lifecycle state of the application virtual IP (VIP) address.
	LifecycleState ApplicationVipLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the create operation for the application virtual IP (VIP) address completed.
	TimeAssigned *common.SDKTime `mandatory:"true" json:"timeAssigned"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet associated with the application virtual IP (VIP) address.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The application virtual IP (VIP) address.
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// Additional information about the current lifecycle state of the application virtual IP (VIP) address.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ApplicationVip) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApplicationVip) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingApplicationVipLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetApplicationVipLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApplicationVipLifecycleStateEnum Enum with underlying type: string
type ApplicationVipLifecycleStateEnum string

// Set of constants representing the allowable values for ApplicationVipLifecycleStateEnum
const (
	ApplicationVipLifecycleStateProvisioning ApplicationVipLifecycleStateEnum = "PROVISIONING"
	ApplicationVipLifecycleStateAvailable    ApplicationVipLifecycleStateEnum = "AVAILABLE"
	ApplicationVipLifecycleStateTerminating  ApplicationVipLifecycleStateEnum = "TERMINATING"
	ApplicationVipLifecycleStateTerminated   ApplicationVipLifecycleStateEnum = "TERMINATED"
	ApplicationVipLifecycleStateFailed       ApplicationVipLifecycleStateEnum = "FAILED"
)

var mappingApplicationVipLifecycleStateEnum = map[string]ApplicationVipLifecycleStateEnum{
	"PROVISIONING": ApplicationVipLifecycleStateProvisioning,
	"AVAILABLE":    ApplicationVipLifecycleStateAvailable,
	"TERMINATING":  ApplicationVipLifecycleStateTerminating,
	"TERMINATED":   ApplicationVipLifecycleStateTerminated,
	"FAILED":       ApplicationVipLifecycleStateFailed,
}

var mappingApplicationVipLifecycleStateEnumLowerCase = map[string]ApplicationVipLifecycleStateEnum{
	"provisioning": ApplicationVipLifecycleStateProvisioning,
	"available":    ApplicationVipLifecycleStateAvailable,
	"terminating":  ApplicationVipLifecycleStateTerminating,
	"terminated":   ApplicationVipLifecycleStateTerminated,
	"failed":       ApplicationVipLifecycleStateFailed,
}

// GetApplicationVipLifecycleStateEnumValues Enumerates the set of values for ApplicationVipLifecycleStateEnum
func GetApplicationVipLifecycleStateEnumValues() []ApplicationVipLifecycleStateEnum {
	values := make([]ApplicationVipLifecycleStateEnum, 0)
	for _, v := range mappingApplicationVipLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetApplicationVipLifecycleStateEnumStringValues Enumerates the set of values in String for ApplicationVipLifecycleStateEnum
func GetApplicationVipLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
	}
}

// GetMappingApplicationVipLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApplicationVipLifecycleStateEnum(val string) (ApplicationVipLifecycleStateEnum, bool) {
	enum, ok := mappingApplicationVipLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
