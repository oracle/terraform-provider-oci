// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TestMacsManagedCloudDatabaseInsightConnectionDetails Connection details of a MACS-managed cloud database .
type TestMacsManagedCloudDatabaseInsightConnectionDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Management Agent
	ManagementAgentId *string `mandatory:"true" json:"managementAgentId"`

	ConnectionDetails *ConnectionDetails `mandatory:"true" json:"connectionDetails"`

	ConnectionCredentialDetails CredentialDetails `mandatory:"true" json:"connectionCredentialDetails"`
}

func (m TestMacsManagedCloudDatabaseInsightConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TestMacsManagedCloudDatabaseInsightConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *TestMacsManagedCloudDatabaseInsightConnectionDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ManagementAgentId           *string            `json:"managementAgentId"`
		ConnectionDetails           *ConnectionDetails `json:"connectionDetails"`
		ConnectionCredentialDetails credentialdetails  `json:"connectionCredentialDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ManagementAgentId = model.ManagementAgentId

	m.ConnectionDetails = model.ConnectionDetails

	nn, e = model.ConnectionCredentialDetails.UnmarshalPolymorphicJSON(model.ConnectionCredentialDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ConnectionCredentialDetails = nn.(CredentialDetails)
	} else {
		m.ConnectionCredentialDetails = nil
	}

	return
}
