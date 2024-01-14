// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// SecurityFeatureAnalyticsDimensions The scope of analytics data.
type SecurityFeatureAnalyticsDimensions struct {

	// The name of the security feature.
	SecurityFeature SecurityFeatureAnalyticsDimensionsSecurityFeatureEnum `mandatory:"false" json:"securityFeature,omitempty"`
}

func (m SecurityFeatureAnalyticsDimensions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecurityFeatureAnalyticsDimensions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSecurityFeatureAnalyticsDimensionsSecurityFeatureEnum(string(m.SecurityFeature)); !ok && m.SecurityFeature != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecurityFeature: %s. Supported values are: %s.", m.SecurityFeature, strings.Join(GetSecurityFeatureAnalyticsDimensionsSecurityFeatureEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SecurityFeatureAnalyticsDimensionsSecurityFeatureEnum Enum with underlying type: string
type SecurityFeatureAnalyticsDimensionsSecurityFeatureEnum string

// Set of constants representing the allowable values for SecurityFeatureAnalyticsDimensionsSecurityFeatureEnum
const (
	SecurityFeatureAnalyticsDimensionsSecurityFeatureUnifiedAudit           SecurityFeatureAnalyticsDimensionsSecurityFeatureEnum = "UNIFIED_AUDIT"
	SecurityFeatureAnalyticsDimensionsSecurityFeatureFineGrainedAudit       SecurityFeatureAnalyticsDimensionsSecurityFeatureEnum = "FINE_GRAINED_AUDIT"
	SecurityFeatureAnalyticsDimensionsSecurityFeatureTraditionalAudit       SecurityFeatureAnalyticsDimensionsSecurityFeatureEnum = "TRADITIONAL_AUDIT"
	SecurityFeatureAnalyticsDimensionsSecurityFeatureDatabaseVault          SecurityFeatureAnalyticsDimensionsSecurityFeatureEnum = "DATABASE_VAULT"
	SecurityFeatureAnalyticsDimensionsSecurityFeaturePrivilegeAnalysis      SecurityFeatureAnalyticsDimensionsSecurityFeatureEnum = "PRIVILEGE_ANALYSIS"
	SecurityFeatureAnalyticsDimensionsSecurityFeatureTablespaceEncryption   SecurityFeatureAnalyticsDimensionsSecurityFeatureEnum = "TABLESPACE_ENCRYPTION"
	SecurityFeatureAnalyticsDimensionsSecurityFeatureColumnEncryption       SecurityFeatureAnalyticsDimensionsSecurityFeatureEnum = "COLUMN_ENCRYPTION"
	SecurityFeatureAnalyticsDimensionsSecurityFeatureNetworkEncryption      SecurityFeatureAnalyticsDimensionsSecurityFeatureEnum = "NETWORK_ENCRYPTION"
	SecurityFeatureAnalyticsDimensionsSecurityFeaturePasswordAuthentication SecurityFeatureAnalyticsDimensionsSecurityFeatureEnum = "PASSWORD_AUTHENTICATION"
	SecurityFeatureAnalyticsDimensionsSecurityFeatureGlobalAuthentication   SecurityFeatureAnalyticsDimensionsSecurityFeatureEnum = "GLOBAL_AUTHENTICATION"
	SecurityFeatureAnalyticsDimensionsSecurityFeatureExternalAuthentication SecurityFeatureAnalyticsDimensionsSecurityFeatureEnum = "EXTERNAL_AUTHENTICATION"
)

var mappingSecurityFeatureAnalyticsDimensionsSecurityFeatureEnum = map[string]SecurityFeatureAnalyticsDimensionsSecurityFeatureEnum{
	"UNIFIED_AUDIT":           SecurityFeatureAnalyticsDimensionsSecurityFeatureUnifiedAudit,
	"FINE_GRAINED_AUDIT":      SecurityFeatureAnalyticsDimensionsSecurityFeatureFineGrainedAudit,
	"TRADITIONAL_AUDIT":       SecurityFeatureAnalyticsDimensionsSecurityFeatureTraditionalAudit,
	"DATABASE_VAULT":          SecurityFeatureAnalyticsDimensionsSecurityFeatureDatabaseVault,
	"PRIVILEGE_ANALYSIS":      SecurityFeatureAnalyticsDimensionsSecurityFeaturePrivilegeAnalysis,
	"TABLESPACE_ENCRYPTION":   SecurityFeatureAnalyticsDimensionsSecurityFeatureTablespaceEncryption,
	"COLUMN_ENCRYPTION":       SecurityFeatureAnalyticsDimensionsSecurityFeatureColumnEncryption,
	"NETWORK_ENCRYPTION":      SecurityFeatureAnalyticsDimensionsSecurityFeatureNetworkEncryption,
	"PASSWORD_AUTHENTICATION": SecurityFeatureAnalyticsDimensionsSecurityFeaturePasswordAuthentication,
	"GLOBAL_AUTHENTICATION":   SecurityFeatureAnalyticsDimensionsSecurityFeatureGlobalAuthentication,
	"EXTERNAL_AUTHENTICATION": SecurityFeatureAnalyticsDimensionsSecurityFeatureExternalAuthentication,
}

var mappingSecurityFeatureAnalyticsDimensionsSecurityFeatureEnumLowerCase = map[string]SecurityFeatureAnalyticsDimensionsSecurityFeatureEnum{
	"unified_audit":           SecurityFeatureAnalyticsDimensionsSecurityFeatureUnifiedAudit,
	"fine_grained_audit":      SecurityFeatureAnalyticsDimensionsSecurityFeatureFineGrainedAudit,
	"traditional_audit":       SecurityFeatureAnalyticsDimensionsSecurityFeatureTraditionalAudit,
	"database_vault":          SecurityFeatureAnalyticsDimensionsSecurityFeatureDatabaseVault,
	"privilege_analysis":      SecurityFeatureAnalyticsDimensionsSecurityFeaturePrivilegeAnalysis,
	"tablespace_encryption":   SecurityFeatureAnalyticsDimensionsSecurityFeatureTablespaceEncryption,
	"column_encryption":       SecurityFeatureAnalyticsDimensionsSecurityFeatureColumnEncryption,
	"network_encryption":      SecurityFeatureAnalyticsDimensionsSecurityFeatureNetworkEncryption,
	"password_authentication": SecurityFeatureAnalyticsDimensionsSecurityFeaturePasswordAuthentication,
	"global_authentication":   SecurityFeatureAnalyticsDimensionsSecurityFeatureGlobalAuthentication,
	"external_authentication": SecurityFeatureAnalyticsDimensionsSecurityFeatureExternalAuthentication,
}

// GetSecurityFeatureAnalyticsDimensionsSecurityFeatureEnumValues Enumerates the set of values for SecurityFeatureAnalyticsDimensionsSecurityFeatureEnum
func GetSecurityFeatureAnalyticsDimensionsSecurityFeatureEnumValues() []SecurityFeatureAnalyticsDimensionsSecurityFeatureEnum {
	values := make([]SecurityFeatureAnalyticsDimensionsSecurityFeatureEnum, 0)
	for _, v := range mappingSecurityFeatureAnalyticsDimensionsSecurityFeatureEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityFeatureAnalyticsDimensionsSecurityFeatureEnumStringValues Enumerates the set of values in String for SecurityFeatureAnalyticsDimensionsSecurityFeatureEnum
func GetSecurityFeatureAnalyticsDimensionsSecurityFeatureEnumStringValues() []string {
	return []string{
		"UNIFIED_AUDIT",
		"FINE_GRAINED_AUDIT",
		"TRADITIONAL_AUDIT",
		"DATABASE_VAULT",
		"PRIVILEGE_ANALYSIS",
		"TABLESPACE_ENCRYPTION",
		"COLUMN_ENCRYPTION",
		"NETWORK_ENCRYPTION",
		"PASSWORD_AUTHENTICATION",
		"GLOBAL_AUTHENTICATION",
		"EXTERNAL_AUTHENTICATION",
	}
}

// GetMappingSecurityFeatureAnalyticsDimensionsSecurityFeatureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityFeatureAnalyticsDimensionsSecurityFeatureEnum(val string) (SecurityFeatureAnalyticsDimensionsSecurityFeatureEnum, bool) {
	enum, ok := mappingSecurityFeatureAnalyticsDimensionsSecurityFeatureEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
