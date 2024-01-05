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

// DatabaseConnectionStringProfile The connection string profile to allow clients to group, filter and select connection string values based on structured metadata.
type DatabaseConnectionStringProfile struct {

	// A user-friendly name for the connection.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Connection string value.
	Value *string `mandatory:"true" json:"value"`

	// Protocol used by the connection.
	Protocol DatabaseConnectionStringProfileProtocolEnum `mandatory:"true" json:"protocol"`

	// Host format used in connection string.
	HostFormat DatabaseConnectionStringProfileHostFormatEnum `mandatory:"true" json:"hostFormat"`

	// Specifies whether the listener performs a direct hand-off of the session, or redirects the session. In RAC deployments where SCAN is used, sessions are redirected to a Node VIP. Use `DIRECT` for direct hand-offs. Use `REDIRECT` to redirect the session.
	SessionMode DatabaseConnectionStringProfileSessionModeEnum `mandatory:"true" json:"sessionMode"`

	// Specifies whether the connection string is using the long (`LONG`), Easy Connect (`EZCONNECT`), or Easy Connect Plus (`EZCONNECTPLUS`) format.
	// Autonomous Database Serverless instances always use the long format.
	SyntaxFormat DatabaseConnectionStringProfileSyntaxFormatEnum `mandatory:"true" json:"syntaxFormat"`

	// Consumer group used by the connection.
	ConsumerGroup DatabaseConnectionStringProfileConsumerGroupEnum `mandatory:"false" json:"consumerGroup,omitempty"`

	// Specifies whether the TLS handshake is using one-way (`SERVER`) or mutual (`MUTUAL`) authentication.
	TlsAuthentication DatabaseConnectionStringProfileTlsAuthenticationEnum `mandatory:"false" json:"tlsAuthentication,omitempty"`

	// True for a regional connection string, applicable to cross-region DG only.
	IsRegional *bool `mandatory:"false" json:"isRegional"`
}

func (m DatabaseConnectionStringProfile) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseConnectionStringProfile) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseConnectionStringProfileProtocolEnum(string(m.Protocol)); !ok && m.Protocol != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Protocol: %s. Supported values are: %s.", m.Protocol, strings.Join(GetDatabaseConnectionStringProfileProtocolEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseConnectionStringProfileHostFormatEnum(string(m.HostFormat)); !ok && m.HostFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HostFormat: %s. Supported values are: %s.", m.HostFormat, strings.Join(GetDatabaseConnectionStringProfileHostFormatEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseConnectionStringProfileSessionModeEnum(string(m.SessionMode)); !ok && m.SessionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SessionMode: %s. Supported values are: %s.", m.SessionMode, strings.Join(GetDatabaseConnectionStringProfileSessionModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseConnectionStringProfileSyntaxFormatEnum(string(m.SyntaxFormat)); !ok && m.SyntaxFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SyntaxFormat: %s. Supported values are: %s.", m.SyntaxFormat, strings.Join(GetDatabaseConnectionStringProfileSyntaxFormatEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDatabaseConnectionStringProfileConsumerGroupEnum(string(m.ConsumerGroup)); !ok && m.ConsumerGroup != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConsumerGroup: %s. Supported values are: %s.", m.ConsumerGroup, strings.Join(GetDatabaseConnectionStringProfileConsumerGroupEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseConnectionStringProfileTlsAuthenticationEnum(string(m.TlsAuthentication)); !ok && m.TlsAuthentication != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TlsAuthentication: %s. Supported values are: %s.", m.TlsAuthentication, strings.Join(GetDatabaseConnectionStringProfileTlsAuthenticationEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseConnectionStringProfileConsumerGroupEnum Enum with underlying type: string
type DatabaseConnectionStringProfileConsumerGroupEnum string

// Set of constants representing the allowable values for DatabaseConnectionStringProfileConsumerGroupEnum
const (
	DatabaseConnectionStringProfileConsumerGroupHigh     DatabaseConnectionStringProfileConsumerGroupEnum = "HIGH"
	DatabaseConnectionStringProfileConsumerGroupMedium   DatabaseConnectionStringProfileConsumerGroupEnum = "MEDIUM"
	DatabaseConnectionStringProfileConsumerGroupLow      DatabaseConnectionStringProfileConsumerGroupEnum = "LOW"
	DatabaseConnectionStringProfileConsumerGroupTp       DatabaseConnectionStringProfileConsumerGroupEnum = "TP"
	DatabaseConnectionStringProfileConsumerGroupTpurgent DatabaseConnectionStringProfileConsumerGroupEnum = "TPURGENT"
)

var mappingDatabaseConnectionStringProfileConsumerGroupEnum = map[string]DatabaseConnectionStringProfileConsumerGroupEnum{
	"HIGH":     DatabaseConnectionStringProfileConsumerGroupHigh,
	"MEDIUM":   DatabaseConnectionStringProfileConsumerGroupMedium,
	"LOW":      DatabaseConnectionStringProfileConsumerGroupLow,
	"TP":       DatabaseConnectionStringProfileConsumerGroupTp,
	"TPURGENT": DatabaseConnectionStringProfileConsumerGroupTpurgent,
}

var mappingDatabaseConnectionStringProfileConsumerGroupEnumLowerCase = map[string]DatabaseConnectionStringProfileConsumerGroupEnum{
	"high":     DatabaseConnectionStringProfileConsumerGroupHigh,
	"medium":   DatabaseConnectionStringProfileConsumerGroupMedium,
	"low":      DatabaseConnectionStringProfileConsumerGroupLow,
	"tp":       DatabaseConnectionStringProfileConsumerGroupTp,
	"tpurgent": DatabaseConnectionStringProfileConsumerGroupTpurgent,
}

// GetDatabaseConnectionStringProfileConsumerGroupEnumValues Enumerates the set of values for DatabaseConnectionStringProfileConsumerGroupEnum
func GetDatabaseConnectionStringProfileConsumerGroupEnumValues() []DatabaseConnectionStringProfileConsumerGroupEnum {
	values := make([]DatabaseConnectionStringProfileConsumerGroupEnum, 0)
	for _, v := range mappingDatabaseConnectionStringProfileConsumerGroupEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseConnectionStringProfileConsumerGroupEnumStringValues Enumerates the set of values in String for DatabaseConnectionStringProfileConsumerGroupEnum
func GetDatabaseConnectionStringProfileConsumerGroupEnumStringValues() []string {
	return []string{
		"HIGH",
		"MEDIUM",
		"LOW",
		"TP",
		"TPURGENT",
	}
}

// GetMappingDatabaseConnectionStringProfileConsumerGroupEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseConnectionStringProfileConsumerGroupEnum(val string) (DatabaseConnectionStringProfileConsumerGroupEnum, bool) {
	enum, ok := mappingDatabaseConnectionStringProfileConsumerGroupEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DatabaseConnectionStringProfileProtocolEnum Enum with underlying type: string
type DatabaseConnectionStringProfileProtocolEnum string

// Set of constants representing the allowable values for DatabaseConnectionStringProfileProtocolEnum
const (
	DatabaseConnectionStringProfileProtocolTcp  DatabaseConnectionStringProfileProtocolEnum = "TCP"
	DatabaseConnectionStringProfileProtocolTcps DatabaseConnectionStringProfileProtocolEnum = "TCPS"
)

var mappingDatabaseConnectionStringProfileProtocolEnum = map[string]DatabaseConnectionStringProfileProtocolEnum{
	"TCP":  DatabaseConnectionStringProfileProtocolTcp,
	"TCPS": DatabaseConnectionStringProfileProtocolTcps,
}

var mappingDatabaseConnectionStringProfileProtocolEnumLowerCase = map[string]DatabaseConnectionStringProfileProtocolEnum{
	"tcp":  DatabaseConnectionStringProfileProtocolTcp,
	"tcps": DatabaseConnectionStringProfileProtocolTcps,
}

// GetDatabaseConnectionStringProfileProtocolEnumValues Enumerates the set of values for DatabaseConnectionStringProfileProtocolEnum
func GetDatabaseConnectionStringProfileProtocolEnumValues() []DatabaseConnectionStringProfileProtocolEnum {
	values := make([]DatabaseConnectionStringProfileProtocolEnum, 0)
	for _, v := range mappingDatabaseConnectionStringProfileProtocolEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseConnectionStringProfileProtocolEnumStringValues Enumerates the set of values in String for DatabaseConnectionStringProfileProtocolEnum
func GetDatabaseConnectionStringProfileProtocolEnumStringValues() []string {
	return []string{
		"TCP",
		"TCPS",
	}
}

// GetMappingDatabaseConnectionStringProfileProtocolEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseConnectionStringProfileProtocolEnum(val string) (DatabaseConnectionStringProfileProtocolEnum, bool) {
	enum, ok := mappingDatabaseConnectionStringProfileProtocolEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DatabaseConnectionStringProfileTlsAuthenticationEnum Enum with underlying type: string
type DatabaseConnectionStringProfileTlsAuthenticationEnum string

// Set of constants representing the allowable values for DatabaseConnectionStringProfileTlsAuthenticationEnum
const (
	DatabaseConnectionStringProfileTlsAuthenticationServer DatabaseConnectionStringProfileTlsAuthenticationEnum = "SERVER"
	DatabaseConnectionStringProfileTlsAuthenticationMutual DatabaseConnectionStringProfileTlsAuthenticationEnum = "MUTUAL"
)

var mappingDatabaseConnectionStringProfileTlsAuthenticationEnum = map[string]DatabaseConnectionStringProfileTlsAuthenticationEnum{
	"SERVER": DatabaseConnectionStringProfileTlsAuthenticationServer,
	"MUTUAL": DatabaseConnectionStringProfileTlsAuthenticationMutual,
}

var mappingDatabaseConnectionStringProfileTlsAuthenticationEnumLowerCase = map[string]DatabaseConnectionStringProfileTlsAuthenticationEnum{
	"server": DatabaseConnectionStringProfileTlsAuthenticationServer,
	"mutual": DatabaseConnectionStringProfileTlsAuthenticationMutual,
}

// GetDatabaseConnectionStringProfileTlsAuthenticationEnumValues Enumerates the set of values for DatabaseConnectionStringProfileTlsAuthenticationEnum
func GetDatabaseConnectionStringProfileTlsAuthenticationEnumValues() []DatabaseConnectionStringProfileTlsAuthenticationEnum {
	values := make([]DatabaseConnectionStringProfileTlsAuthenticationEnum, 0)
	for _, v := range mappingDatabaseConnectionStringProfileTlsAuthenticationEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseConnectionStringProfileTlsAuthenticationEnumStringValues Enumerates the set of values in String for DatabaseConnectionStringProfileTlsAuthenticationEnum
func GetDatabaseConnectionStringProfileTlsAuthenticationEnumStringValues() []string {
	return []string{
		"SERVER",
		"MUTUAL",
	}
}

// GetMappingDatabaseConnectionStringProfileTlsAuthenticationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseConnectionStringProfileTlsAuthenticationEnum(val string) (DatabaseConnectionStringProfileTlsAuthenticationEnum, bool) {
	enum, ok := mappingDatabaseConnectionStringProfileTlsAuthenticationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DatabaseConnectionStringProfileHostFormatEnum Enum with underlying type: string
type DatabaseConnectionStringProfileHostFormatEnum string

// Set of constants representing the allowable values for DatabaseConnectionStringProfileHostFormatEnum
const (
	DatabaseConnectionStringProfileHostFormatFqdn DatabaseConnectionStringProfileHostFormatEnum = "FQDN"
	DatabaseConnectionStringProfileHostFormatIp   DatabaseConnectionStringProfileHostFormatEnum = "IP"
)

var mappingDatabaseConnectionStringProfileHostFormatEnum = map[string]DatabaseConnectionStringProfileHostFormatEnum{
	"FQDN": DatabaseConnectionStringProfileHostFormatFqdn,
	"IP":   DatabaseConnectionStringProfileHostFormatIp,
}

var mappingDatabaseConnectionStringProfileHostFormatEnumLowerCase = map[string]DatabaseConnectionStringProfileHostFormatEnum{
	"fqdn": DatabaseConnectionStringProfileHostFormatFqdn,
	"ip":   DatabaseConnectionStringProfileHostFormatIp,
}

// GetDatabaseConnectionStringProfileHostFormatEnumValues Enumerates the set of values for DatabaseConnectionStringProfileHostFormatEnum
func GetDatabaseConnectionStringProfileHostFormatEnumValues() []DatabaseConnectionStringProfileHostFormatEnum {
	values := make([]DatabaseConnectionStringProfileHostFormatEnum, 0)
	for _, v := range mappingDatabaseConnectionStringProfileHostFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseConnectionStringProfileHostFormatEnumStringValues Enumerates the set of values in String for DatabaseConnectionStringProfileHostFormatEnum
func GetDatabaseConnectionStringProfileHostFormatEnumStringValues() []string {
	return []string{
		"FQDN",
		"IP",
	}
}

// GetMappingDatabaseConnectionStringProfileHostFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseConnectionStringProfileHostFormatEnum(val string) (DatabaseConnectionStringProfileHostFormatEnum, bool) {
	enum, ok := mappingDatabaseConnectionStringProfileHostFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DatabaseConnectionStringProfileSessionModeEnum Enum with underlying type: string
type DatabaseConnectionStringProfileSessionModeEnum string

// Set of constants representing the allowable values for DatabaseConnectionStringProfileSessionModeEnum
const (
	DatabaseConnectionStringProfileSessionModeDirect   DatabaseConnectionStringProfileSessionModeEnum = "DIRECT"
	DatabaseConnectionStringProfileSessionModeRedirect DatabaseConnectionStringProfileSessionModeEnum = "REDIRECT"
)

var mappingDatabaseConnectionStringProfileSessionModeEnum = map[string]DatabaseConnectionStringProfileSessionModeEnum{
	"DIRECT":   DatabaseConnectionStringProfileSessionModeDirect,
	"REDIRECT": DatabaseConnectionStringProfileSessionModeRedirect,
}

var mappingDatabaseConnectionStringProfileSessionModeEnumLowerCase = map[string]DatabaseConnectionStringProfileSessionModeEnum{
	"direct":   DatabaseConnectionStringProfileSessionModeDirect,
	"redirect": DatabaseConnectionStringProfileSessionModeRedirect,
}

// GetDatabaseConnectionStringProfileSessionModeEnumValues Enumerates the set of values for DatabaseConnectionStringProfileSessionModeEnum
func GetDatabaseConnectionStringProfileSessionModeEnumValues() []DatabaseConnectionStringProfileSessionModeEnum {
	values := make([]DatabaseConnectionStringProfileSessionModeEnum, 0)
	for _, v := range mappingDatabaseConnectionStringProfileSessionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseConnectionStringProfileSessionModeEnumStringValues Enumerates the set of values in String for DatabaseConnectionStringProfileSessionModeEnum
func GetDatabaseConnectionStringProfileSessionModeEnumStringValues() []string {
	return []string{
		"DIRECT",
		"REDIRECT",
	}
}

// GetMappingDatabaseConnectionStringProfileSessionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseConnectionStringProfileSessionModeEnum(val string) (DatabaseConnectionStringProfileSessionModeEnum, bool) {
	enum, ok := mappingDatabaseConnectionStringProfileSessionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DatabaseConnectionStringProfileSyntaxFormatEnum Enum with underlying type: string
type DatabaseConnectionStringProfileSyntaxFormatEnum string

// Set of constants representing the allowable values for DatabaseConnectionStringProfileSyntaxFormatEnum
const (
	DatabaseConnectionStringProfileSyntaxFormatLong          DatabaseConnectionStringProfileSyntaxFormatEnum = "LONG"
	DatabaseConnectionStringProfileSyntaxFormatEzconnect     DatabaseConnectionStringProfileSyntaxFormatEnum = "EZCONNECT"
	DatabaseConnectionStringProfileSyntaxFormatEzconnectplus DatabaseConnectionStringProfileSyntaxFormatEnum = "EZCONNECTPLUS"
)

var mappingDatabaseConnectionStringProfileSyntaxFormatEnum = map[string]DatabaseConnectionStringProfileSyntaxFormatEnum{
	"LONG":          DatabaseConnectionStringProfileSyntaxFormatLong,
	"EZCONNECT":     DatabaseConnectionStringProfileSyntaxFormatEzconnect,
	"EZCONNECTPLUS": DatabaseConnectionStringProfileSyntaxFormatEzconnectplus,
}

var mappingDatabaseConnectionStringProfileSyntaxFormatEnumLowerCase = map[string]DatabaseConnectionStringProfileSyntaxFormatEnum{
	"long":          DatabaseConnectionStringProfileSyntaxFormatLong,
	"ezconnect":     DatabaseConnectionStringProfileSyntaxFormatEzconnect,
	"ezconnectplus": DatabaseConnectionStringProfileSyntaxFormatEzconnectplus,
}

// GetDatabaseConnectionStringProfileSyntaxFormatEnumValues Enumerates the set of values for DatabaseConnectionStringProfileSyntaxFormatEnum
func GetDatabaseConnectionStringProfileSyntaxFormatEnumValues() []DatabaseConnectionStringProfileSyntaxFormatEnum {
	values := make([]DatabaseConnectionStringProfileSyntaxFormatEnum, 0)
	for _, v := range mappingDatabaseConnectionStringProfileSyntaxFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseConnectionStringProfileSyntaxFormatEnumStringValues Enumerates the set of values in String for DatabaseConnectionStringProfileSyntaxFormatEnum
func GetDatabaseConnectionStringProfileSyntaxFormatEnumStringValues() []string {
	return []string{
		"LONG",
		"EZCONNECT",
		"EZCONNECTPLUS",
	}
}

// GetMappingDatabaseConnectionStringProfileSyntaxFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseConnectionStringProfileSyntaxFormatEnum(val string) (DatabaseConnectionStringProfileSyntaxFormatEnum, bool) {
	enum, ok := mappingDatabaseConnectionStringProfileSyntaxFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
