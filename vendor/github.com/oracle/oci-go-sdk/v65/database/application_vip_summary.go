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

// ApplicationVipSummary Details of an application virtual IP (VIP) address.
type ApplicationVipSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the application virtual IP (VIP) address.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the cloud VM cluster associated with the application virtual IP (VIP) address.
	CloudVmClusterId *string `mandatory:"true" json:"cloudVmClusterId"`

	// The hostname of the application virtual IP (VIP) address.
	HostnameLabel *string `mandatory:"true" json:"hostnameLabel"`

	// The current lifecycle state of the application virtual IP (VIP) address.
	LifecycleState ApplicationVipSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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

func (m ApplicationVipSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApplicationVipSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingApplicationVipSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetApplicationVipSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ApplicationVipSummaryLifecycleStateEnum Enum with underlying type: string
type ApplicationVipSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ApplicationVipSummaryLifecycleStateEnum
const (
	ApplicationVipSummaryLifecycleStateProvisioning ApplicationVipSummaryLifecycleStateEnum = "PROVISIONING"
	ApplicationVipSummaryLifecycleStateAvailable    ApplicationVipSummaryLifecycleStateEnum = "AVAILABLE"
	ApplicationVipSummaryLifecycleStateTerminating  ApplicationVipSummaryLifecycleStateEnum = "TERMINATING"
	ApplicationVipSummaryLifecycleStateTerminated   ApplicationVipSummaryLifecycleStateEnum = "TERMINATED"
	ApplicationVipSummaryLifecycleStateFailed       ApplicationVipSummaryLifecycleStateEnum = "FAILED"
)

var mappingApplicationVipSummaryLifecycleStateEnum = map[string]ApplicationVipSummaryLifecycleStateEnum{
	"PROVISIONING": ApplicationVipSummaryLifecycleStateProvisioning,
	"AVAILABLE":    ApplicationVipSummaryLifecycleStateAvailable,
	"TERMINATING":  ApplicationVipSummaryLifecycleStateTerminating,
	"TERMINATED":   ApplicationVipSummaryLifecycleStateTerminated,
	"FAILED":       ApplicationVipSummaryLifecycleStateFailed,
}

var mappingApplicationVipSummaryLifecycleStateEnumLowerCase = map[string]ApplicationVipSummaryLifecycleStateEnum{
	"provisioning": ApplicationVipSummaryLifecycleStateProvisioning,
	"available":    ApplicationVipSummaryLifecycleStateAvailable,
	"terminating":  ApplicationVipSummaryLifecycleStateTerminating,
	"terminated":   ApplicationVipSummaryLifecycleStateTerminated,
	"failed":       ApplicationVipSummaryLifecycleStateFailed,
}

// GetApplicationVipSummaryLifecycleStateEnumValues Enumerates the set of values for ApplicationVipSummaryLifecycleStateEnum
func GetApplicationVipSummaryLifecycleStateEnumValues() []ApplicationVipSummaryLifecycleStateEnum {
	values := make([]ApplicationVipSummaryLifecycleStateEnum, 0)
	for _, v := range mappingApplicationVipSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetApplicationVipSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ApplicationVipSummaryLifecycleStateEnum
func GetApplicationVipSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
	}
}

// GetMappingApplicationVipSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApplicationVipSummaryLifecycleStateEnum(val string) (ApplicationVipSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingApplicationVipSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
