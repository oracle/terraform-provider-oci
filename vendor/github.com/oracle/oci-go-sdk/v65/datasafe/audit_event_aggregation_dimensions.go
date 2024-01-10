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

// AuditEventAggregationDimensions The details of the aggregation dimensions used for summarizing audit events.
type AuditEventAggregationDimensions struct {

	// The time the audit event occurred in the target database.
	AuditEventTime []common.SDKTime `mandatory:"false" json:"auditEventTime"`

	// Name of the database user whose actions were audited.
	DbUserName []string `mandatory:"false" json:"dbUserName"`

	// The OCID of the target database that was audited.
	TargetId []string `mandatory:"false" json:"targetId"`

	// The name of the target database that was audited.
	TargetName []string `mandatory:"false" json:"targetName"`

	// Class of the target that was audited.
	TargetClass []AuditEventAggregationDimensionsTargetClassEnum `mandatory:"false" json:"targetClass,omitempty"`

	// Type of object in the source database affected by the action. For example PL/SQL, SYNONYM or PACKAGE BODY.
	ObjectType []string `mandatory:"false" json:"objectType"`

	// Name of the host machine from which the session was spawned.
	ClientHostname []string `mandatory:"false" json:"clientHostname"`

	// The application from which the audit event was generated. For example SQL Plus or SQL Developer.
	ClientProgram []string `mandatory:"false" json:"clientProgram"`

	// The client identifier in each Oracle session.
	ClientId []string `mandatory:"false" json:"clientId"`

	// Type of auditing.
	AuditType []AuditEventAggregationDimensionsAuditTypeEnum `mandatory:"false" json:"auditType,omitempty"`

	// The name of the event executed by the user on the target database. For example ALTER SEQUENCE, CREATE TRIGGER or CREATE INDEX.
	EventName []string `mandatory:"false" json:"eventName"`
}

func (m AuditEventAggregationDimensions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuditEventAggregationDimensions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.TargetClass {
		if _, ok := GetMappingAuditEventAggregationDimensionsTargetClassEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetClass: %s. Supported values are: %s.", val, strings.Join(GetAuditEventAggregationDimensionsTargetClassEnumStringValues(), ",")))
		}
	}

	for _, val := range m.AuditType {
		if _, ok := GetMappingAuditEventAggregationDimensionsAuditTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuditType: %s. Supported values are: %s.", val, strings.Join(GetAuditEventAggregationDimensionsAuditTypeEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AuditEventAggregationDimensionsTargetClassEnum Enum with underlying type: string
type AuditEventAggregationDimensionsTargetClassEnum string

// Set of constants representing the allowable values for AuditEventAggregationDimensionsTargetClassEnum
const (
	AuditEventAggregationDimensionsTargetClassDatabase AuditEventAggregationDimensionsTargetClassEnum = "DATABASE"
)

var mappingAuditEventAggregationDimensionsTargetClassEnum = map[string]AuditEventAggregationDimensionsTargetClassEnum{
	"DATABASE": AuditEventAggregationDimensionsTargetClassDatabase,
}

var mappingAuditEventAggregationDimensionsTargetClassEnumLowerCase = map[string]AuditEventAggregationDimensionsTargetClassEnum{
	"database": AuditEventAggregationDimensionsTargetClassDatabase,
}

// GetAuditEventAggregationDimensionsTargetClassEnumValues Enumerates the set of values for AuditEventAggregationDimensionsTargetClassEnum
func GetAuditEventAggregationDimensionsTargetClassEnumValues() []AuditEventAggregationDimensionsTargetClassEnum {
	values := make([]AuditEventAggregationDimensionsTargetClassEnum, 0)
	for _, v := range mappingAuditEventAggregationDimensionsTargetClassEnum {
		values = append(values, v)
	}
	return values
}

// GetAuditEventAggregationDimensionsTargetClassEnumStringValues Enumerates the set of values in String for AuditEventAggregationDimensionsTargetClassEnum
func GetAuditEventAggregationDimensionsTargetClassEnumStringValues() []string {
	return []string{
		"DATABASE",
	}
}

// GetMappingAuditEventAggregationDimensionsTargetClassEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuditEventAggregationDimensionsTargetClassEnum(val string) (AuditEventAggregationDimensionsTargetClassEnum, bool) {
	enum, ok := mappingAuditEventAggregationDimensionsTargetClassEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AuditEventAggregationDimensionsAuditTypeEnum Enum with underlying type: string
type AuditEventAggregationDimensionsAuditTypeEnum string

// Set of constants representing the allowable values for AuditEventAggregationDimensionsAuditTypeEnum
const (
	AuditEventAggregationDimensionsAuditTypeStandard      AuditEventAggregationDimensionsAuditTypeEnum = "STANDARD"
	AuditEventAggregationDimensionsAuditTypeFineGrained   AuditEventAggregationDimensionsAuditTypeEnum = "FINE_GRAINED"
	AuditEventAggregationDimensionsAuditTypeXs            AuditEventAggregationDimensionsAuditTypeEnum = "XS"
	AuditEventAggregationDimensionsAuditTypeDatabaseVault AuditEventAggregationDimensionsAuditTypeEnum = "DATABASE_VAULT"
	AuditEventAggregationDimensionsAuditTypeLabelSecurity AuditEventAggregationDimensionsAuditTypeEnum = "LABEL_SECURITY"
	AuditEventAggregationDimensionsAuditTypeRman          AuditEventAggregationDimensionsAuditTypeEnum = "RMAN"
	AuditEventAggregationDimensionsAuditTypeDatapump      AuditEventAggregationDimensionsAuditTypeEnum = "DATAPUMP"
	AuditEventAggregationDimensionsAuditTypeDirectPathApi AuditEventAggregationDimensionsAuditTypeEnum = "DIRECT_PATH_API"
)

var mappingAuditEventAggregationDimensionsAuditTypeEnum = map[string]AuditEventAggregationDimensionsAuditTypeEnum{
	"STANDARD":        AuditEventAggregationDimensionsAuditTypeStandard,
	"FINE_GRAINED":    AuditEventAggregationDimensionsAuditTypeFineGrained,
	"XS":              AuditEventAggregationDimensionsAuditTypeXs,
	"DATABASE_VAULT":  AuditEventAggregationDimensionsAuditTypeDatabaseVault,
	"LABEL_SECURITY":  AuditEventAggregationDimensionsAuditTypeLabelSecurity,
	"RMAN":            AuditEventAggregationDimensionsAuditTypeRman,
	"DATAPUMP":        AuditEventAggregationDimensionsAuditTypeDatapump,
	"DIRECT_PATH_API": AuditEventAggregationDimensionsAuditTypeDirectPathApi,
}

var mappingAuditEventAggregationDimensionsAuditTypeEnumLowerCase = map[string]AuditEventAggregationDimensionsAuditTypeEnum{
	"standard":        AuditEventAggregationDimensionsAuditTypeStandard,
	"fine_grained":    AuditEventAggregationDimensionsAuditTypeFineGrained,
	"xs":              AuditEventAggregationDimensionsAuditTypeXs,
	"database_vault":  AuditEventAggregationDimensionsAuditTypeDatabaseVault,
	"label_security":  AuditEventAggregationDimensionsAuditTypeLabelSecurity,
	"rman":            AuditEventAggregationDimensionsAuditTypeRman,
	"datapump":        AuditEventAggregationDimensionsAuditTypeDatapump,
	"direct_path_api": AuditEventAggregationDimensionsAuditTypeDirectPathApi,
}

// GetAuditEventAggregationDimensionsAuditTypeEnumValues Enumerates the set of values for AuditEventAggregationDimensionsAuditTypeEnum
func GetAuditEventAggregationDimensionsAuditTypeEnumValues() []AuditEventAggregationDimensionsAuditTypeEnum {
	values := make([]AuditEventAggregationDimensionsAuditTypeEnum, 0)
	for _, v := range mappingAuditEventAggregationDimensionsAuditTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAuditEventAggregationDimensionsAuditTypeEnumStringValues Enumerates the set of values in String for AuditEventAggregationDimensionsAuditTypeEnum
func GetAuditEventAggregationDimensionsAuditTypeEnumStringValues() []string {
	return []string{
		"STANDARD",
		"FINE_GRAINED",
		"XS",
		"DATABASE_VAULT",
		"LABEL_SECURITY",
		"RMAN",
		"DATAPUMP",
		"DIRECT_PATH_API",
	}
}

// GetMappingAuditEventAggregationDimensionsAuditTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuditEventAggregationDimensionsAuditTypeEnum(val string) (AuditEventAggregationDimensionsAuditTypeEnum, bool) {
	enum, ok := mappingAuditEventAggregationDimensionsAuditTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
