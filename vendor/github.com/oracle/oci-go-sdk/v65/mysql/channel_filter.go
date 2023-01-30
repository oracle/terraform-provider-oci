// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// ChannelFilter Replication filter rule for a channel.
type ChannelFilter struct {

	// The type of the filter rule.
	// For details on each type, see
	// Replication Filtering Rules (https://dev.mysql.com/doc/refman/8.0/en/replication-rules.html)
	Type ChannelFilterTypeEnum `mandatory:"true" json:"type"`

	// The body of the filter rule. This can represent a database, a table, or a database pair (represented as
	// "db1->db2"). For more information, see
	// Replication Filtering Rules (https://dev.mysql.com/doc/refman/8.0/en/replication-rules.html).
	Value *string `mandatory:"true" json:"value"`
}

func (m ChannelFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChannelFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingChannelFilterTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetChannelFilterTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ChannelFilterTypeEnum Enum with underlying type: string
type ChannelFilterTypeEnum string

// Set of constants representing the allowable values for ChannelFilterTypeEnum
const (
	ChannelFilterTypeDoDb            ChannelFilterTypeEnum = "REPLICATE_DO_DB"
	ChannelFilterTypeIgnoreDb        ChannelFilterTypeEnum = "REPLICATE_IGNORE_DB"
	ChannelFilterTypeDoTable         ChannelFilterTypeEnum = "REPLICATE_DO_TABLE"
	ChannelFilterTypeIgnoreTable     ChannelFilterTypeEnum = "REPLICATE_IGNORE_TABLE"
	ChannelFilterTypeWildDoTable     ChannelFilterTypeEnum = "REPLICATE_WILD_DO_TABLE"
	ChannelFilterTypeWildIgnoreTable ChannelFilterTypeEnum = "REPLICATE_WILD_IGNORE_TABLE"
	ChannelFilterTypeRewriteDb       ChannelFilterTypeEnum = "REPLICATE_REWRITE_DB"
)

var mappingChannelFilterTypeEnum = map[string]ChannelFilterTypeEnum{
	"REPLICATE_DO_DB":             ChannelFilterTypeDoDb,
	"REPLICATE_IGNORE_DB":         ChannelFilterTypeIgnoreDb,
	"REPLICATE_DO_TABLE":          ChannelFilterTypeDoTable,
	"REPLICATE_IGNORE_TABLE":      ChannelFilterTypeIgnoreTable,
	"REPLICATE_WILD_DO_TABLE":     ChannelFilterTypeWildDoTable,
	"REPLICATE_WILD_IGNORE_TABLE": ChannelFilterTypeWildIgnoreTable,
	"REPLICATE_REWRITE_DB":        ChannelFilterTypeRewriteDb,
}

var mappingChannelFilterTypeEnumLowerCase = map[string]ChannelFilterTypeEnum{
	"replicate_do_db":             ChannelFilterTypeDoDb,
	"replicate_ignore_db":         ChannelFilterTypeIgnoreDb,
	"replicate_do_table":          ChannelFilterTypeDoTable,
	"replicate_ignore_table":      ChannelFilterTypeIgnoreTable,
	"replicate_wild_do_table":     ChannelFilterTypeWildDoTable,
	"replicate_wild_ignore_table": ChannelFilterTypeWildIgnoreTable,
	"replicate_rewrite_db":        ChannelFilterTypeRewriteDb,
}

// GetChannelFilterTypeEnumValues Enumerates the set of values for ChannelFilterTypeEnum
func GetChannelFilterTypeEnumValues() []ChannelFilterTypeEnum {
	values := make([]ChannelFilterTypeEnum, 0)
	for _, v := range mappingChannelFilterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetChannelFilterTypeEnumStringValues Enumerates the set of values in String for ChannelFilterTypeEnum
func GetChannelFilterTypeEnumStringValues() []string {
	return []string{
		"REPLICATE_DO_DB",
		"REPLICATE_IGNORE_DB",
		"REPLICATE_DO_TABLE",
		"REPLICATE_IGNORE_TABLE",
		"REPLICATE_WILD_DO_TABLE",
		"REPLICATE_WILD_IGNORE_TABLE",
		"REPLICATE_REWRITE_DB",
	}
}

// GetMappingChannelFilterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingChannelFilterTypeEnum(val string) (ChannelFilterTypeEnum, bool) {
	enum, ok := mappingChannelFilterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
