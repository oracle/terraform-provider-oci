// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// UpdatePolarisIcebergCatalogDetails The information to update a Polaris catalog used in the Iceberg connection.
type UpdatePolarisIcebergCatalogDetails struct {

	// The URL endpoint for the Polaris API.
	// e.g.: 'https://<your-snowflake-account>.snowflakecomputing.com/polaris/api/catalog'
	Uri *string `mandatory:"false" json:"uri"`

	// The catalog name within Polaris where Iceberg tables are registered.
	Name *string `mandatory:"false" json:"name"`

	// The OAuth client ID used for authentication.
	ClientId *string `mandatory:"false" json:"clientId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the password Oracle GoldenGate uses to connect to Snowflake platform.
	ClientSecretSecretId *string `mandatory:"false" json:"clientSecretSecretId"`

	// The Snowflake role used to access Polaris.
	PrincipalRole *string `mandatory:"false" json:"principalRole"`
}

func (m UpdatePolarisIcebergCatalogDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdatePolarisIcebergCatalogDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdatePolarisIcebergCatalogDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdatePolarisIcebergCatalogDetails UpdatePolarisIcebergCatalogDetails
	s := struct {
		DiscriminatorParam string `json:"catalogType"`
		MarshalTypeUpdatePolarisIcebergCatalogDetails
	}{
		"POLARIS",
		(MarshalTypeUpdatePolarisIcebergCatalogDetails)(m),
	}

	return json.Marshal(&s)
}
