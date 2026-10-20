// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ManagePrivilegesDetails The details used to manage the features of a Data Safe target database.
type ManagePrivilegesDetails struct {

	// The Data Safe features granted to the databases registering under the registration policy.
	Features []ManagePrivilegesDetailsFeaturesEnum `mandatory:"true" json:"features"`
}

func (m ManagePrivilegesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagePrivilegesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.Features {
		if _, ok := GetMappingManagePrivilegesDetailsFeaturesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Features: %s. Supported values are: %s.", val, strings.Join(GetManagePrivilegesDetailsFeaturesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ManagePrivilegesDetailsFeaturesEnum Enum with underlying type: string
type ManagePrivilegesDetailsFeaturesEnum string

// Set of constants representing the allowable values for ManagePrivilegesDetailsFeaturesEnum
const (
	ManagePrivilegesDetailsFeaturesAssessment      ManagePrivilegesDetailsFeaturesEnum = "ASSESSMENT"
	ManagePrivilegesDetailsFeaturesAuditCollection ManagePrivilegesDetailsFeaturesEnum = "AUDIT_COLLECTION"
	ManagePrivilegesDetailsFeaturesAuditSetting    ManagePrivilegesDetailsFeaturesEnum = "AUDIT_SETTING"
	ManagePrivilegesDetailsFeaturesDataDiscovery   ManagePrivilegesDetailsFeaturesEnum = "DATA_DISCOVERY"
	ManagePrivilegesDetailsFeaturesMasking         ManagePrivilegesDetailsFeaturesEnum = "MASKING"
	ManagePrivilegesDetailsFeaturesSqlFirewall     ManagePrivilegesDetailsFeaturesEnum = "SQL_FIREWALL"
	ManagePrivilegesDetailsFeaturesAll             ManagePrivilegesDetailsFeaturesEnum = "ALL"
)

var mappingManagePrivilegesDetailsFeaturesEnum = map[string]ManagePrivilegesDetailsFeaturesEnum{
	"ASSESSMENT":       ManagePrivilegesDetailsFeaturesAssessment,
	"AUDIT_COLLECTION": ManagePrivilegesDetailsFeaturesAuditCollection,
	"AUDIT_SETTING":    ManagePrivilegesDetailsFeaturesAuditSetting,
	"DATA_DISCOVERY":   ManagePrivilegesDetailsFeaturesDataDiscovery,
	"MASKING":          ManagePrivilegesDetailsFeaturesMasking,
	"SQL_FIREWALL":     ManagePrivilegesDetailsFeaturesSqlFirewall,
	"ALL":              ManagePrivilegesDetailsFeaturesAll,
}

var mappingManagePrivilegesDetailsFeaturesEnumLowerCase = map[string]ManagePrivilegesDetailsFeaturesEnum{
	"assessment":       ManagePrivilegesDetailsFeaturesAssessment,
	"audit_collection": ManagePrivilegesDetailsFeaturesAuditCollection,
	"audit_setting":    ManagePrivilegesDetailsFeaturesAuditSetting,
	"data_discovery":   ManagePrivilegesDetailsFeaturesDataDiscovery,
	"masking":          ManagePrivilegesDetailsFeaturesMasking,
	"sql_firewall":     ManagePrivilegesDetailsFeaturesSqlFirewall,
	"all":              ManagePrivilegesDetailsFeaturesAll,
}

// GetManagePrivilegesDetailsFeaturesEnumValues Enumerates the set of values for ManagePrivilegesDetailsFeaturesEnum
func GetManagePrivilegesDetailsFeaturesEnumValues() []ManagePrivilegesDetailsFeaturesEnum {
	values := make([]ManagePrivilegesDetailsFeaturesEnum, 0)
	for _, v := range mappingManagePrivilegesDetailsFeaturesEnum {
		values = append(values, v)
	}
	return values
}

// GetManagePrivilegesDetailsFeaturesEnumStringValues Enumerates the set of values in String for ManagePrivilegesDetailsFeaturesEnum
func GetManagePrivilegesDetailsFeaturesEnumStringValues() []string {
	return []string{
		"ASSESSMENT",
		"AUDIT_COLLECTION",
		"AUDIT_SETTING",
		"DATA_DISCOVERY",
		"MASKING",
		"SQL_FIREWALL",
		"ALL",
	}
}

// GetMappingManagePrivilegesDetailsFeaturesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagePrivilegesDetailsFeaturesEnum(val string) (ManagePrivilegesDetailsFeaturesEnum, bool) {
	enum, ok := mappingManagePrivilegesDetailsFeaturesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
