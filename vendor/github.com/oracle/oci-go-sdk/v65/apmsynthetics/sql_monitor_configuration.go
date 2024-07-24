// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SqlMonitorConfiguration Request configuration details for the SQL monitor type.
type SqlMonitorConfiguration struct {

	// If isFailureRetried is enabled, then a failed call will be retried.
	IsFailureRetried *bool `mandatory:"false" json:"isFailureRetried"`

	DnsConfiguration *DnsConfiguration `mandatory:"false" json:"dnsConfiguration"`

	// SQL query to be executed.
	Query *string `mandatory:"false" json:"query"`

	DatabaseAuthenticationDetails *BasicAuthenticationDetails `mandatory:"false" json:"databaseAuthenticationDetails"`

	// Database role.
	DatabaseRole *string `mandatory:"false" json:"databaseRole"`

	// Database connection string.
	ConnectionString *string `mandatory:"false" json:"connectionString"`

	DatabaseWalletDetails *DatabaseWalletDetails `mandatory:"false" json:"databaseWalletDetails"`

	// Database type.
	DatabaseType DatabaseTypeEnum `mandatory:"false" json:"databaseType,omitempty"`

	// Database connection type. Only CUSTOM_JDBC is supported for MYSQL database type.
	DatabaseConnectionType DatabaseConnectionTypeEnum `mandatory:"false" json:"databaseConnectionType,omitempty"`
}

// GetIsFailureRetried returns IsFailureRetried
func (m SqlMonitorConfiguration) GetIsFailureRetried() *bool {
	return m.IsFailureRetried
}

// GetDnsConfiguration returns DnsConfiguration
func (m SqlMonitorConfiguration) GetDnsConfiguration() *DnsConfiguration {
	return m.DnsConfiguration
}

func (m SqlMonitorConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlMonitorConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDatabaseTypeEnum(string(m.DatabaseType)); !ok && m.DatabaseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseType: %s. Supported values are: %s.", m.DatabaseType, strings.Join(GetDatabaseTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseConnectionTypeEnum(string(m.DatabaseConnectionType)); !ok && m.DatabaseConnectionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseConnectionType: %s. Supported values are: %s.", m.DatabaseConnectionType, strings.Join(GetDatabaseConnectionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SqlMonitorConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSqlMonitorConfiguration SqlMonitorConfiguration
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeSqlMonitorConfiguration
	}{
		"SQL_CONFIG",
		(MarshalTypeSqlMonitorConfiguration)(m),
	}

	return json.Marshal(&s)
}
