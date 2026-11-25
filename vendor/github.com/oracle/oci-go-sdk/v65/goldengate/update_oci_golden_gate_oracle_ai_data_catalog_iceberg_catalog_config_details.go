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

// UpdateOciGoldenGateOracleAiDataCatalogIcebergCatalogConfigDetails The information to update an Oracle AI Data Catalog configuration based on an OCI GoldenGate deployment.
type UpdateOciGoldenGateOracleAiDataCatalogIcebergCatalogConfigDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the deployment being referenced.
	DeploymentId *string `mandatory:"false" json:"deploymentId"`
}

func (m UpdateOciGoldenGateOracleAiDataCatalogIcebergCatalogConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOciGoldenGateOracleAiDataCatalogIcebergCatalogConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateOciGoldenGateOracleAiDataCatalogIcebergCatalogConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateOciGoldenGateOracleAiDataCatalogIcebergCatalogConfigDetails UpdateOciGoldenGateOracleAiDataCatalogIcebergCatalogConfigDetails
	s := struct {
		DiscriminatorParam string `json:"catalogConfigType"`
		MarshalTypeUpdateOciGoldenGateOracleAiDataCatalogIcebergCatalogConfigDetails
	}{
		"OCI_GOLDENGATE",
		(MarshalTypeUpdateOciGoldenGateOracleAiDataCatalogIcebergCatalogConfigDetails)(m),
	}

	return json.Marshal(&s)
}
