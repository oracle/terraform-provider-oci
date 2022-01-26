// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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
	// Autonomous Databases on shared Exadata infrastructure always use the long format.
	SyntaxFormat DatabaseConnectionStringProfileSyntaxFormatEnum `mandatory:"true" json:"syntaxFormat"`

	// Consumer group used by the connection.
	ConsumerGroup DatabaseConnectionStringProfileConsumerGroupEnum `mandatory:"false" json:"consumerGroup,omitempty"`

	// Specifies whether the TLS handshake is using one-way (`SERVER`) or mutual (`MUTUAL`) authentication.
	TlsAuthentication DatabaseConnectionStringProfileTlsAuthenticationEnum `mandatory:"false" json:"tlsAuthentication,omitempty"`
}

func (m DatabaseConnectionStringProfile) String() string {
	return common.PointerString(m)
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

var mappingDatabaseConnectionStringProfileConsumerGroup = map[string]DatabaseConnectionStringProfileConsumerGroupEnum{
	"HIGH":     DatabaseConnectionStringProfileConsumerGroupHigh,
	"MEDIUM":   DatabaseConnectionStringProfileConsumerGroupMedium,
	"LOW":      DatabaseConnectionStringProfileConsumerGroupLow,
	"TP":       DatabaseConnectionStringProfileConsumerGroupTp,
	"TPURGENT": DatabaseConnectionStringProfileConsumerGroupTpurgent,
}

// GetDatabaseConnectionStringProfileConsumerGroupEnumValues Enumerates the set of values for DatabaseConnectionStringProfileConsumerGroupEnum
func GetDatabaseConnectionStringProfileConsumerGroupEnumValues() []DatabaseConnectionStringProfileConsumerGroupEnum {
	values := make([]DatabaseConnectionStringProfileConsumerGroupEnum, 0)
	for _, v := range mappingDatabaseConnectionStringProfileConsumerGroup {
		values = append(values, v)
	}
	return values
}

// DatabaseConnectionStringProfileProtocolEnum Enum with underlying type: string
type DatabaseConnectionStringProfileProtocolEnum string

// Set of constants representing the allowable values for DatabaseConnectionStringProfileProtocolEnum
const (
	DatabaseConnectionStringProfileProtocolTcp  DatabaseConnectionStringProfileProtocolEnum = "TCP"
	DatabaseConnectionStringProfileProtocolTcps DatabaseConnectionStringProfileProtocolEnum = "TCPS"
)

var mappingDatabaseConnectionStringProfileProtocol = map[string]DatabaseConnectionStringProfileProtocolEnum{
	"TCP":  DatabaseConnectionStringProfileProtocolTcp,
	"TCPS": DatabaseConnectionStringProfileProtocolTcps,
}

// GetDatabaseConnectionStringProfileProtocolEnumValues Enumerates the set of values for DatabaseConnectionStringProfileProtocolEnum
func GetDatabaseConnectionStringProfileProtocolEnumValues() []DatabaseConnectionStringProfileProtocolEnum {
	values := make([]DatabaseConnectionStringProfileProtocolEnum, 0)
	for _, v := range mappingDatabaseConnectionStringProfileProtocol {
		values = append(values, v)
	}
	return values
}

// DatabaseConnectionStringProfileTlsAuthenticationEnum Enum with underlying type: string
type DatabaseConnectionStringProfileTlsAuthenticationEnum string

// Set of constants representing the allowable values for DatabaseConnectionStringProfileTlsAuthenticationEnum
const (
	DatabaseConnectionStringProfileTlsAuthenticationServer DatabaseConnectionStringProfileTlsAuthenticationEnum = "SERVER"
	DatabaseConnectionStringProfileTlsAuthenticationMutual DatabaseConnectionStringProfileTlsAuthenticationEnum = "MUTUAL"
)

var mappingDatabaseConnectionStringProfileTlsAuthentication = map[string]DatabaseConnectionStringProfileTlsAuthenticationEnum{
	"SERVER": DatabaseConnectionStringProfileTlsAuthenticationServer,
	"MUTUAL": DatabaseConnectionStringProfileTlsAuthenticationMutual,
}

// GetDatabaseConnectionStringProfileTlsAuthenticationEnumValues Enumerates the set of values for DatabaseConnectionStringProfileTlsAuthenticationEnum
func GetDatabaseConnectionStringProfileTlsAuthenticationEnumValues() []DatabaseConnectionStringProfileTlsAuthenticationEnum {
	values := make([]DatabaseConnectionStringProfileTlsAuthenticationEnum, 0)
	for _, v := range mappingDatabaseConnectionStringProfileTlsAuthentication {
		values = append(values, v)
	}
	return values
}

// DatabaseConnectionStringProfileHostFormatEnum Enum with underlying type: string
type DatabaseConnectionStringProfileHostFormatEnum string

// Set of constants representing the allowable values for DatabaseConnectionStringProfileHostFormatEnum
const (
	DatabaseConnectionStringProfileHostFormatFqdn DatabaseConnectionStringProfileHostFormatEnum = "FQDN"
	DatabaseConnectionStringProfileHostFormatIp   DatabaseConnectionStringProfileHostFormatEnum = "IP"
)

var mappingDatabaseConnectionStringProfileHostFormat = map[string]DatabaseConnectionStringProfileHostFormatEnum{
	"FQDN": DatabaseConnectionStringProfileHostFormatFqdn,
	"IP":   DatabaseConnectionStringProfileHostFormatIp,
}

// GetDatabaseConnectionStringProfileHostFormatEnumValues Enumerates the set of values for DatabaseConnectionStringProfileHostFormatEnum
func GetDatabaseConnectionStringProfileHostFormatEnumValues() []DatabaseConnectionStringProfileHostFormatEnum {
	values := make([]DatabaseConnectionStringProfileHostFormatEnum, 0)
	for _, v := range mappingDatabaseConnectionStringProfileHostFormat {
		values = append(values, v)
	}
	return values
}

// DatabaseConnectionStringProfileSessionModeEnum Enum with underlying type: string
type DatabaseConnectionStringProfileSessionModeEnum string

// Set of constants representing the allowable values for DatabaseConnectionStringProfileSessionModeEnum
const (
	DatabaseConnectionStringProfileSessionModeDirect   DatabaseConnectionStringProfileSessionModeEnum = "DIRECT"
	DatabaseConnectionStringProfileSessionModeRedirect DatabaseConnectionStringProfileSessionModeEnum = "REDIRECT"
)

var mappingDatabaseConnectionStringProfileSessionMode = map[string]DatabaseConnectionStringProfileSessionModeEnum{
	"DIRECT":   DatabaseConnectionStringProfileSessionModeDirect,
	"REDIRECT": DatabaseConnectionStringProfileSessionModeRedirect,
}

// GetDatabaseConnectionStringProfileSessionModeEnumValues Enumerates the set of values for DatabaseConnectionStringProfileSessionModeEnum
func GetDatabaseConnectionStringProfileSessionModeEnumValues() []DatabaseConnectionStringProfileSessionModeEnum {
	values := make([]DatabaseConnectionStringProfileSessionModeEnum, 0)
	for _, v := range mappingDatabaseConnectionStringProfileSessionMode {
		values = append(values, v)
	}
	return values
}

// DatabaseConnectionStringProfileSyntaxFormatEnum Enum with underlying type: string
type DatabaseConnectionStringProfileSyntaxFormatEnum string

// Set of constants representing the allowable values for DatabaseConnectionStringProfileSyntaxFormatEnum
const (
	DatabaseConnectionStringProfileSyntaxFormatLong          DatabaseConnectionStringProfileSyntaxFormatEnum = "LONG"
	DatabaseConnectionStringProfileSyntaxFormatEzconnect     DatabaseConnectionStringProfileSyntaxFormatEnum = "EZCONNECT"
	DatabaseConnectionStringProfileSyntaxFormatEzconnectplus DatabaseConnectionStringProfileSyntaxFormatEnum = "EZCONNECTPLUS"
)

var mappingDatabaseConnectionStringProfileSyntaxFormat = map[string]DatabaseConnectionStringProfileSyntaxFormatEnum{
	"LONG":          DatabaseConnectionStringProfileSyntaxFormatLong,
	"EZCONNECT":     DatabaseConnectionStringProfileSyntaxFormatEzconnect,
	"EZCONNECTPLUS": DatabaseConnectionStringProfileSyntaxFormatEzconnectplus,
}

// GetDatabaseConnectionStringProfileSyntaxFormatEnumValues Enumerates the set of values for DatabaseConnectionStringProfileSyntaxFormatEnum
func GetDatabaseConnectionStringProfileSyntaxFormatEnumValues() []DatabaseConnectionStringProfileSyntaxFormatEnum {
	values := make([]DatabaseConnectionStringProfileSyntaxFormatEnum, 0)
	for _, v := range mappingDatabaseConnectionStringProfileSyntaxFormat {
		values = append(values, v)
	}
	return values
}
