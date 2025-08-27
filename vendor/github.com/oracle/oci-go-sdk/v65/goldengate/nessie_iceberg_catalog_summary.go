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

// NessieIcebergCatalogSummary Summary of the Nessie catalog used in the Iceberg connection.
type NessieIcebergCatalogSummary struct {

	// Nessie URI.
	// e.g.: 'http://<nessie-server>.com:10001/api/v2'
	Uri *string `mandatory:"true" json:"uri"`

	// The active branch of the Nessie catalog from which Iceberg reads and writes table metadata.
	Branch *string `mandatory:"true" json:"branch"`
}

func (m NessieIcebergCatalogSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NessieIcebergCatalogSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m NessieIcebergCatalogSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeNessieIcebergCatalogSummary NessieIcebergCatalogSummary
	s := struct {
		DiscriminatorParam string `json:"catalogType"`
		MarshalTypeNessieIcebergCatalogSummary
	}{
		"NESSIE",
		(MarshalTypeNessieIcebergCatalogSummary)(m),
	}

	return json.Marshal(&s)
}
