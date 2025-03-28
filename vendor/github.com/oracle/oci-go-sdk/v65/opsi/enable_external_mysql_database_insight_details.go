// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EnableExternalMysqlDatabaseInsightDetails The information about database to be analyzed.
type EnableExternalMysqlDatabaseInsightDetails struct {

	// The DBM owned database connector OCID (https://docs.oracle.com/iaas/database-management/doc/view-connector-details.html) mapping to the database credentials and connection details.
	DatabaseConnectorId *string `mandatory:"true" json:"databaseConnectorId"`
}

func (m EnableExternalMysqlDatabaseInsightDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EnableExternalMysqlDatabaseInsightDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m EnableExternalMysqlDatabaseInsightDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeEnableExternalMysqlDatabaseInsightDetails EnableExternalMysqlDatabaseInsightDetails
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeEnableExternalMysqlDatabaseInsightDetails
	}{
		"EXTERNAL_MYSQL_DATABASE_SYSTEM",
		(MarshalTypeEnableExternalMysqlDatabaseInsightDetails)(m),
	}

	return json.Marshal(&s)
}
