// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.cloud.oracle.com/Content/application-dependency-management/home.htm).
//

package adm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DetectConfiguration A configuration to define the constraints when detecting vulnerable dependencies.
type DetectConfiguration struct {

	// The list of dependencies to be ignored by the recommendation algorithm. The dependency pattern is matched against the 'group:artifact:version' or the purl of a dependency.
	// An asterisk (*) at the end in the dependency pattern acts as a wildcard and matches zero or more characters.
	Exclusions []string `mandatory:"false" json:"exclusions"`

	// The upgrade policy for recommendations.
	// The `Nearest` upgrade policy upgrades a dependency to the oldest version that meets both of the following criteria: it is newer than the current version and it is not affected by a vulnerability.
	UpgradePolicy DetectConfigurationUpgradePolicyEnum `mandatory:"false" json:"upgradePolicy,omitempty"`

	// The maximum Common Vulnerability Scoring System Version 2 (CVSS V2) score. An artifact with a CVSS V2 score below this value is not considered for patching.
	MaxPermissibleCvssV2Score *float32 `mandatory:"false" json:"maxPermissibleCvssV2Score"`

	// The maximum Common Vulnerability Scoring System Version 3 (CVSS V3) score. An artifact with a CVSS V3 score below this value is not considered for patching.
	MaxPermissibleCvssV3Score *float32 `mandatory:"false" json:"maxPermissibleCvssV3Score"`
}

func (m DetectConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DetectConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDetectConfigurationUpgradePolicyEnum(string(m.UpgradePolicy)); !ok && m.UpgradePolicy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpgradePolicy: %s. Supported values are: %s.", m.UpgradePolicy, strings.Join(GetDetectConfigurationUpgradePolicyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DetectConfigurationUpgradePolicyEnum Enum with underlying type: string
type DetectConfigurationUpgradePolicyEnum string

// Set of constants representing the allowable values for DetectConfigurationUpgradePolicyEnum
const (
	DetectConfigurationUpgradePolicyNearest DetectConfigurationUpgradePolicyEnum = "NEAREST"
)

var mappingDetectConfigurationUpgradePolicyEnum = map[string]DetectConfigurationUpgradePolicyEnum{
	"NEAREST": DetectConfigurationUpgradePolicyNearest,
}

var mappingDetectConfigurationUpgradePolicyEnumLowerCase = map[string]DetectConfigurationUpgradePolicyEnum{
	"nearest": DetectConfigurationUpgradePolicyNearest,
}

// GetDetectConfigurationUpgradePolicyEnumValues Enumerates the set of values for DetectConfigurationUpgradePolicyEnum
func GetDetectConfigurationUpgradePolicyEnumValues() []DetectConfigurationUpgradePolicyEnum {
	values := make([]DetectConfigurationUpgradePolicyEnum, 0)
	for _, v := range mappingDetectConfigurationUpgradePolicyEnum {
		values = append(values, v)
	}
	return values
}

// GetDetectConfigurationUpgradePolicyEnumStringValues Enumerates the set of values in String for DetectConfigurationUpgradePolicyEnum
func GetDetectConfigurationUpgradePolicyEnumStringValues() []string {
	return []string{
		"NEAREST",
	}
}

// GetMappingDetectConfigurationUpgradePolicyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDetectConfigurationUpgradePolicyEnum(val string) (DetectConfigurationUpgradePolicyEnum, bool) {
	enum, ok := mappingDetectConfigurationUpgradePolicyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
