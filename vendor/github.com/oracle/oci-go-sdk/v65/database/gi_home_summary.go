// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GiHomeSummary A Grid Infrastructure Home summary.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an
// administrator. If you're an administrator who needs to write policies to give users access,
// see Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type GiHomeSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Grid Infrastructure Home.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-provided name for the Grid Infrastructure Home. The name does not need to be unique and gets dynamically generated if null.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the Grid Infrastructure Home.
	LifecycleState GiHomeSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The path of the Grid Infrastructure Home.
	HomePath *string `mandatory:"true" json:"homePath"`

	// A valid Oracle Grid Infrastructure (GI) software version.
	GiVersion *string `mandatory:"true" json:"giVersion"`

	// Indicates whether the grid infrastructure home is created by default.
	IsDefaultCreated *bool `mandatory:"false" json:"isDefaultCreated"`

	// The time and date as an RFC3339 formatted string, e.g., 2024-04-11T01:59:07.032Z, when the grid infrastructure home was created
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The grid infrastructure home image OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
	GiImageId *string `mandatory:"false" json:"giImageId"`

	// The time and date as an RFC3339 formatted string, e.g., 2024-04-11T01:59:07.032Z, when the grid infrastructure home was updated
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m GiHomeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GiHomeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGiHomeSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetGiHomeSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GiHomeSummaryLifecycleStateEnum Enum with underlying type: string
type GiHomeSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for GiHomeSummaryLifecycleStateEnum
const (
	GiHomeSummaryLifecycleStateProvisioning GiHomeSummaryLifecycleStateEnum = "PROVISIONING"
	GiHomeSummaryLifecycleStateAvailable    GiHomeSummaryLifecycleStateEnum = "AVAILABLE"
	GiHomeSummaryLifecycleStateUpdating     GiHomeSummaryLifecycleStateEnum = "UPDATING"
	GiHomeSummaryLifecycleStateTerminating  GiHomeSummaryLifecycleStateEnum = "TERMINATING"
	GiHomeSummaryLifecycleStateTerminated   GiHomeSummaryLifecycleStateEnum = "TERMINATED"
	GiHomeSummaryLifecycleStateFailed       GiHomeSummaryLifecycleStateEnum = "FAILED"
	GiHomeSummaryLifecycleStateInactive     GiHomeSummaryLifecycleStateEnum = "INACTIVE"
)

var mappingGiHomeSummaryLifecycleStateEnum = map[string]GiHomeSummaryLifecycleStateEnum{
	"PROVISIONING": GiHomeSummaryLifecycleStateProvisioning,
	"AVAILABLE":    GiHomeSummaryLifecycleStateAvailable,
	"UPDATING":     GiHomeSummaryLifecycleStateUpdating,
	"TERMINATING":  GiHomeSummaryLifecycleStateTerminating,
	"TERMINATED":   GiHomeSummaryLifecycleStateTerminated,
	"FAILED":       GiHomeSummaryLifecycleStateFailed,
	"INACTIVE":     GiHomeSummaryLifecycleStateInactive,
}

var mappingGiHomeSummaryLifecycleStateEnumLowerCase = map[string]GiHomeSummaryLifecycleStateEnum{
	"provisioning": GiHomeSummaryLifecycleStateProvisioning,
	"available":    GiHomeSummaryLifecycleStateAvailable,
	"updating":     GiHomeSummaryLifecycleStateUpdating,
	"terminating":  GiHomeSummaryLifecycleStateTerminating,
	"terminated":   GiHomeSummaryLifecycleStateTerminated,
	"failed":       GiHomeSummaryLifecycleStateFailed,
	"inactive":     GiHomeSummaryLifecycleStateInactive,
}

// GetGiHomeSummaryLifecycleStateEnumValues Enumerates the set of values for GiHomeSummaryLifecycleStateEnum
func GetGiHomeSummaryLifecycleStateEnumValues() []GiHomeSummaryLifecycleStateEnum {
	values := make([]GiHomeSummaryLifecycleStateEnum, 0)
	for _, v := range mappingGiHomeSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetGiHomeSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for GiHomeSummaryLifecycleStateEnum
func GetGiHomeSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"UPDATING",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
		"INACTIVE",
	}
}

// GetMappingGiHomeSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGiHomeSummaryLifecycleStateEnum(val string) (GiHomeSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingGiHomeSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
