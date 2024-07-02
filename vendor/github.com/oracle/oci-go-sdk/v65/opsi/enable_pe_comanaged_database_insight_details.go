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

// EnablePeComanagedDatabaseInsightDetails The information about database to be analyzed.
type EnablePeComanagedDatabaseInsightDetails struct {

	// The compartment OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Private service accessed database.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OPSI private endpoint
	OpsiPrivateEndpointId *string `mandatory:"true" json:"opsiPrivateEndpointId"`

	// Database service name used for connection requests.
	ServiceName *string `mandatory:"true" json:"serviceName"`

	CredentialDetails CredentialDetails `mandatory:"true" json:"credentialDetails"`

	ConnectionDetails *PeComanagedDatabaseConnectionDetails `mandatory:"false" json:"connectionDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m EnablePeComanagedDatabaseInsightDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EnablePeComanagedDatabaseInsightDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m EnablePeComanagedDatabaseInsightDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeEnablePeComanagedDatabaseInsightDetails EnablePeComanagedDatabaseInsightDetails
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeEnablePeComanagedDatabaseInsightDetails
	}{
		"PE_COMANAGED_DATABASE",
		(MarshalTypeEnablePeComanagedDatabaseInsightDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *EnablePeComanagedDatabaseInsightDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ConnectionDetails     *PeComanagedDatabaseConnectionDetails `json:"connectionDetails"`
		FreeformTags          map[string]string                     `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{}     `json:"definedTags"`
		SystemTags            map[string]map[string]interface{}     `json:"systemTags"`
		CompartmentId         *string                               `json:"compartmentId"`
		OpsiPrivateEndpointId *string                               `json:"opsiPrivateEndpointId"`
		ServiceName           *string                               `json:"serviceName"`
		CredentialDetails     credentialdetails                     `json:"credentialDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ConnectionDetails = model.ConnectionDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.CompartmentId = model.CompartmentId

	m.OpsiPrivateEndpointId = model.OpsiPrivateEndpointId

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

	return
}
