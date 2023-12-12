// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseSecurityConfigSummary Database Security Configurations resource represents the target database configurations.
// Included in the Database Security Configurations are the SQL Firewall configurations such as
// the status of the firewall, the time that the firewall status was last updated, violation log auto purge settings, etc.
type DatabaseSecurityConfigSummary struct {

	// The OCID of the database security config.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the database security config.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the database security config.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The target OCID corresponding to the database security config.
	TargetId *string `mandatory:"true" json:"targetId"`

	// The time that the database security config was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the database security config.
	LifecycleState DatabaseSecurityConfigLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The description of the database security config.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the database security configuration was last updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The last date and time the database security config was refreshed, in the format defined by RFC3339.
	TimeLastRefreshed *common.SDKTime `mandatory:"false" json:"timeLastRefreshed"`

	// Details about the current state of the database security config in Data Safe.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	SqlFirewallConfig *SqlFirewallConfig `mandatory:"false" json:"sqlFirewallConfig"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m DatabaseSecurityConfigSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseSecurityConfigSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseSecurityConfigLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDatabaseSecurityConfigLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
