// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ChannelTargetDbSystem Core properties of a DB System Channel target.
type ChannelTargetDbSystem struct {

	// The OCID of the source DB System.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// The case-insensitive name that identifies the replication channel. Channel names
	// must follow the rules defined for MySQL identifiers (https://dev.mysql.com/doc/refman/8.0/en/identifiers.html).
	// The names of non-Deleted Channels must be unique for each DB System.
	ChannelName *string `mandatory:"true" json:"channelName"`

	// The username for the replication applier of the target MySQL DB System.
	ApplierUsername *string `mandatory:"true" json:"applierUsername"`

	// Specifies the amount of time, in seconds, that the channel waits before
	// applying a transaction received from the source.
	DelayInSeconds *int `mandatory:"true" json:"delayInSeconds"`

	// Replication filter rules to be applied at the DB System Channel target.
	Filters []ChannelFilter `mandatory:"false" json:"filters"`

	// Specifies how a replication channel handles the creation and alteration of tables
	// that do not have a primary key.
	TablesWithoutPrimaryKeyHandling ChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnum `mandatory:"true" json:"tablesWithoutPrimaryKeyHandling"`
}

func (m ChannelTargetDbSystem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChannelTargetDbSystem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnum(string(m.TablesWithoutPrimaryKeyHandling)); !ok && m.TablesWithoutPrimaryKeyHandling != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TablesWithoutPrimaryKeyHandling: %s. Supported values are: %s.", m.TablesWithoutPrimaryKeyHandling, strings.Join(GetChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ChannelTargetDbSystem) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeChannelTargetDbSystem ChannelTargetDbSystem
	s := struct {
		DiscriminatorParam string `json:"targetType"`
		MarshalTypeChannelTargetDbSystem
	}{
		"DBSYSTEM",
		(MarshalTypeChannelTargetDbSystem)(m),
	}

	return json.Marshal(&s)
}

// ChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnum Enum with underlying type: string
type ChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnum string

// Set of constants representing the allowable values for ChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnum
const (
	ChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingRaiseError                 ChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnum = "RAISE_ERROR"
	ChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingAllow                      ChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnum = "ALLOW"
	ChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingGenerateImplicitPrimaryKey ChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnum = "GENERATE_IMPLICIT_PRIMARY_KEY"
)

var mappingChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnum = map[string]ChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnum{
	"RAISE_ERROR":                   ChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingRaiseError,
	"ALLOW":                         ChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingAllow,
	"GENERATE_IMPLICIT_PRIMARY_KEY": ChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingGenerateImplicitPrimaryKey,
}

var mappingChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnumLowerCase = map[string]ChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnum{
	"raise_error":                   ChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingRaiseError,
	"allow":                         ChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingAllow,
	"generate_implicit_primary_key": ChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingGenerateImplicitPrimaryKey,
}

// GetChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnumValues Enumerates the set of values for ChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnum
func GetChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnumValues() []ChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnum {
	values := make([]ChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnum, 0)
	for _, v := range mappingChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnum {
		values = append(values, v)
	}
	return values
}

// GetChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnumStringValues Enumerates the set of values in String for ChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnum
func GetChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnumStringValues() []string {
	return []string{
		"RAISE_ERROR",
		"ALLOW",
		"GENERATE_IMPLICIT_PRIMARY_KEY",
	}
}

// GetMappingChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnum(val string) (ChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnum, bool) {
	enum, ok := mappingChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
