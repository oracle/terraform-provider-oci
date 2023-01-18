// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// AutonomousDatabaseConnectionStrings Connection strings to connect to an Oracle Autonomous Database.
// Example output for connection strings. See DatabaseConnectionStringProfile for additional details:
//     	"connectionStrings": {
//         "allConnectionStrings": {
//           "HIGH": "adb.region.oraclecloud.com:1522/unique_id_databasename_high.adwc.oraclecloud.com",
//           "LOW": "adb.region.oraclecloud.com:1522/unique_id_databasename_low.adwc.oraclecloud.com",
//           "MEDIUM": "adb.region.oraclecloud.com:1522/unique_id_databasename_medium.adwc.oraclecloud.com"
//         },
//         "profiles": [
//           {
//             "displayName": "databasename_high",
//             "value": "(description= (retry_count=20)(retry_delay=3)(address=(protocol=tcps)(port=1522)(host=adb.region.oraclecloud.com))(connect_data=(service_name=unique_id_databasename_high.adwc.oraclecloud.com))(security=(ssl_server_cert_dn="CN=adwc.uscom-east-1.oraclecloud.com,OU=Oracle BMCS US,O=Oracle Corporation,L=Redwood City,ST=California,C=US")))",
//             "consumerGroup": "HIGH",
//             "protocol": "TCPS",
//             "tlsAuthentication": "MUTUAL",
//             "hostFormat": "FQDN",
//             "sessionMode": "DIRECT",
//             "syntaxFormat": "LONG"
//           },
//           {
//             "displayName": "databasename_low",
//             "value": "(description= (retry_count=20)(retry_delay=3)(address=(protocol=tcps)(port=1522)(host=adb.region.oraclecloud.com))(connect_data=(service_name=unique_id_databasename_low.adwc.oraclecloud.com))(security=(ssl_server_cert_dn="CN=adwc.uscom-east-1.oraclecloud.com,OU=Oracle BMCS US,O=Oracle Corporation,L=Redwood City,ST=California,C=US")))",
//             "consumerGroup": "LOW",
//             "protocol": "TCPS",
//             "tlsAuthentication": "MUTUAL",
//             "hostFormat": "FQDN",
//             "sessionMode": "DIRECT",
//             "syntaxFormat": "LONG"
//           },
//           {
//             "displayName": "databasename_medium",
//             "value": "(description= (retry_count=20)(retry_delay=3)(address=(protocol=tcps)(port=1522)(host=adb.region.oraclecloud.com))(connect_data=(service_name=unique_id_databasename_medium.adwc.oraclecloud.com))(security=(ssl_server_cert_dn="CN=adwc.uscom-east-1.oraclecloud.com,OU=Oracle BMCS US,O=Oracle Corporation,L=Redwood City,ST=California,C=US")))",
//             "consumerGroup": "MEDIUM",
//             "protocol": "TCPS",
//             "tlsAuthentication": "MUTUAL",
//             "hostFormat": "FQDN",
//             "sessionMode": "DIRECT",
//             "syntaxFormat": "LONG"
//           }
//         ],
//         "dedicated": null,
//         "high": "adb.region.oraclecloud.com:1522/unique_id_databasename_high.adwc.oraclecloud.com",
//         "low": "adb.region.oraclecloud.com:1522/unique_id_databasename_low.adwc.oraclecloud.com",
//         "medium": "adb.region.oraclecloud.com:1522/unique_id_databasename_medium.adwc.oraclecloud.com"
//       }
type AutonomousDatabaseConnectionStrings struct {

	// The High database service provides the highest level of resources to each SQL statement resulting in the highest performance, but supports the fewest number of concurrent SQL statements.
	High *string `mandatory:"false" json:"high"`

	// The Medium database service provides a lower level of resources to each SQL statement potentially resulting a lower level of performance, but supports more concurrent SQL statements.
	Medium *string `mandatory:"false" json:"medium"`

	// The Low database service provides the least level of resources to each SQL statement, but supports the most number of concurrent SQL statements.
	Low *string `mandatory:"false" json:"low"`

	// The database service provides the least level of resources to each SQL statement, but supports the most number of concurrent SQL statements.
	Dedicated *string `mandatory:"false" json:"dedicated"`

	// Returns all connection strings that can be used to connect to the Autonomous Database.
	// For more information, please see Predefined Database Service Names for Autonomous Transaction Processing (https://docs.oracle.com/en/cloud/paas/atp-cloud/atpug/connect-predefined.html#GUID-9747539B-FD46-44F1-8FF8-F5AC650F15BE)
	AllConnectionStrings map[string]string `mandatory:"false" json:"allConnectionStrings"`

	// A list of connection string profiles to allow clients to group, filter and select connection string values based on structured metadata.
	Profiles []DatabaseConnectionStringProfile `mandatory:"false" json:"profiles"`
}

func (m AutonomousDatabaseConnectionStrings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousDatabaseConnectionStrings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
