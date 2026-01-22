// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GenerateDbSystemStatusDetails Parameters for selecting which attributes to include in the DB System status.
type GenerateDbSystemStatusDetails struct {

	// Specifies whether to collect the full set of GTIDs executed on the DB System
	// (gtid_executed (https://dev.mysql.com/doc/en/replication-options-gtids.html#sysvar_gtid_executed))
	// or the set of GTIDs available in the DB System binary logs, not part of
	// (gtid_purged (https://dev.mysql.com/doc/en/replication-options-gtids.html#sysvar_gtid_purged))
	// (GTID_AVAILABLE, which is gtid_executed - gtid_purged).
	// Use GTID_AVAILABLE if gtid_executed is too large to fit in the response.
	GtidSetType GenerateDbSystemStatusDetailsGtidSetTypeEnum `mandatory:"false" json:"gtidSetType,omitempty"`

	// Specifies the GTID set to be checked on the DB System to determine whether it has been applied. The result of
	// the check is isGtidSetApplied in the response.
	GtidSetToApply *string `mandatory:"false" json:"gtidSetToApply"`

	// List of Channel IDs for which to collect status information.
	ChannelIds []string `mandatory:"false" json:"channelIds"`
}

func (m GenerateDbSystemStatusDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenerateDbSystemStatusDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingGenerateDbSystemStatusDetailsGtidSetTypeEnum(string(m.GtidSetType)); !ok && m.GtidSetType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GtidSetType: %s. Supported values are: %s.", m.GtidSetType, strings.Join(GetGenerateDbSystemStatusDetailsGtidSetTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GenerateDbSystemStatusDetailsGtidSetTypeEnum Enum with underlying type: string
type GenerateDbSystemStatusDetailsGtidSetTypeEnum string

// Set of constants representing the allowable values for GenerateDbSystemStatusDetailsGtidSetTypeEnum
const (
	GenerateDbSystemStatusDetailsGtidSetTypeExecuted  GenerateDbSystemStatusDetailsGtidSetTypeEnum = "GTID_EXECUTED"
	GenerateDbSystemStatusDetailsGtidSetTypeAvailable GenerateDbSystemStatusDetailsGtidSetTypeEnum = "GTID_AVAILABLE"
)

var mappingGenerateDbSystemStatusDetailsGtidSetTypeEnum = map[string]GenerateDbSystemStatusDetailsGtidSetTypeEnum{
	"GTID_EXECUTED":  GenerateDbSystemStatusDetailsGtidSetTypeExecuted,
	"GTID_AVAILABLE": GenerateDbSystemStatusDetailsGtidSetTypeAvailable,
}

var mappingGenerateDbSystemStatusDetailsGtidSetTypeEnumLowerCase = map[string]GenerateDbSystemStatusDetailsGtidSetTypeEnum{
	"gtid_executed":  GenerateDbSystemStatusDetailsGtidSetTypeExecuted,
	"gtid_available": GenerateDbSystemStatusDetailsGtidSetTypeAvailable,
}

// GetGenerateDbSystemStatusDetailsGtidSetTypeEnumValues Enumerates the set of values for GenerateDbSystemStatusDetailsGtidSetTypeEnum
func GetGenerateDbSystemStatusDetailsGtidSetTypeEnumValues() []GenerateDbSystemStatusDetailsGtidSetTypeEnum {
	values := make([]GenerateDbSystemStatusDetailsGtidSetTypeEnum, 0)
	for _, v := range mappingGenerateDbSystemStatusDetailsGtidSetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGenerateDbSystemStatusDetailsGtidSetTypeEnumStringValues Enumerates the set of values in String for GenerateDbSystemStatusDetailsGtidSetTypeEnum
func GetGenerateDbSystemStatusDetailsGtidSetTypeEnumStringValues() []string {
	return []string{
		"GTID_EXECUTED",
		"GTID_AVAILABLE",
	}
}

// GetMappingGenerateDbSystemStatusDetailsGtidSetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenerateDbSystemStatusDetailsGtidSetTypeEnum(val string) (GenerateDbSystemStatusDetailsGtidSetTypeEnum, bool) {
	enum, ok := mappingGenerateDbSystemStatusDetailsGtidSetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
