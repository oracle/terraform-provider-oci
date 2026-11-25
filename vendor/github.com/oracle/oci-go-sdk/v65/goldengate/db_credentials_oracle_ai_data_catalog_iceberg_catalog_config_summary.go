// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DbCredentialsOracleAiDataCatalogIcebergCatalogConfigSummary Summary of the Oracle AI Data Catalog configuration based on database credentials.
type DbCredentialsOracleAiDataCatalogIcebergCatalogConfigSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database being referenced.
	DatabaseId *string `mandatory:"true" json:"databaseId"`

	// The username Oracle GoldenGate uses to connect the associated system of the given technology.
	// This username must already exist and be available by the system/application to be connected to
	// and must conform to the case sensitivty requirments defined in it.
	Username *string `mandatory:"true" json:"username"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the password is stored.
	// The password Oracle GoldenGate uses to connect the associated system of the given technology.
	// It must conform to the specific security requirements including length, case sensitivity, and so on.
	// If secretId is used plaintext field must not be provided.
	// Note: When provided, 'password' field must not be provided.
	PasswordSecretId *string `mandatory:"false" json:"passwordSecretId"`

	// The database type where Oracle AI Data Catalog is configured.
	DatabaseType OracleAiDataCatalogIcebergCatalogDatabaseTypeEnum `mandatory:"true" json:"databaseType"`
}

func (m DbCredentialsOracleAiDataCatalogIcebergCatalogConfigSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbCredentialsOracleAiDataCatalogIcebergCatalogConfigSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOracleAiDataCatalogIcebergCatalogDatabaseTypeEnum(string(m.DatabaseType)); !ok && m.DatabaseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseType: %s. Supported values are: %s.", m.DatabaseType, strings.Join(GetOracleAiDataCatalogIcebergCatalogDatabaseTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DbCredentialsOracleAiDataCatalogIcebergCatalogConfigSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDbCredentialsOracleAiDataCatalogIcebergCatalogConfigSummary DbCredentialsOracleAiDataCatalogIcebergCatalogConfigSummary
	s := struct {
		DiscriminatorParam string `json:"catalogConfigType"`
		MarshalTypeDbCredentialsOracleAiDataCatalogIcebergCatalogConfigSummary
	}{
		"DB_CREDENTIALS",
		(MarshalTypeDbCredentialsOracleAiDataCatalogIcebergCatalogConfigSummary)(m),
	}

	return json.Marshal(&s)
}
