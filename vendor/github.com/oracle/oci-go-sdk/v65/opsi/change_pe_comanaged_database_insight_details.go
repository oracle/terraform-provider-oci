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

// ChangePeComanagedDatabaseInsightDetails Details of a Private Endpoint co-managed database insight.
type ChangePeComanagedDatabaseInsightDetails struct {

	// Database service name used for connection requests.
	ServiceName *string `mandatory:"true" json:"serviceName"`

	CredentialDetails CredentialDetails `mandatory:"true" json:"credentialDetails"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OPSI private endpoint
	OpsiPrivateEndpointId *string `mandatory:"true" json:"opsiPrivateEndpointId"`

	ConnectionDetails *PeComanagedDatabaseConnectionDetails `mandatory:"false" json:"connectionDetails"`
}

func (m ChangePeComanagedDatabaseInsightDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChangePeComanagedDatabaseInsightDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ChangePeComanagedDatabaseInsightDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ConnectionDetails     *PeComanagedDatabaseConnectionDetails `json:"connectionDetails"`
		ServiceName           *string                               `json:"serviceName"`
		CredentialDetails     credentialdetails                     `json:"credentialDetails"`
		OpsiPrivateEndpointId *string                               `json:"opsiPrivateEndpointId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ConnectionDetails = model.ConnectionDetails

	m.ServiceName = model.ServiceName

	nn, e = model.CredentialDetails.UnmarshalPolymorphicJSON(model.CredentialDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CredentialDetails = nn.(CredentialDetails)
	} else {
		m.CredentialDetails = nil
	}

	m.OpsiPrivateEndpointId = model.OpsiPrivateEndpointId

	return
}
