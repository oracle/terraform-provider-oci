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

// UpdateKubeClusterInsightDetails The information to be updated.
type UpdateKubeClusterInsightDetails interface {

	// The Kubernetes cluster service account.
	GetServiceAccount() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the K8s cluster token secret id.
	GetTokenSecretId() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the K8s cluster certificate secret id.
	GetCertificateSecretId() *string

	// The Kubernetes cluster API server URL.
	GetApiServerUrl() *string

	// The Kubernetes cluster API Server port.
	GetApiServerPort() *int

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type updatekubeclusterinsightdetails struct {
	JsonData            []byte
	ServiceAccount      *string                           `mandatory:"false" json:"serviceAccount"`
	TokenSecretId       *string                           `mandatory:"false" json:"tokenSecretId"`
	CertificateSecretId *string                           `mandatory:"false" json:"certificateSecretId"`
	ApiServerUrl        *string                           `mandatory:"false" json:"apiServerUrl"`
	ApiServerPort       *int                              `mandatory:"false" json:"apiServerPort"`
	FreeformTags        map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags         map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	EntitySource        string                            `json:"entitySource"`
}

// UnmarshalJSON unmarshals json
func (m *updatekubeclusterinsightdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatekubeclusterinsightdetails updatekubeclusterinsightdetails
	s := struct {
		Model Unmarshalerupdatekubeclusterinsightdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ServiceAccount = s.Model.ServiceAccount
	m.TokenSecretId = s.Model.TokenSecretId
	m.CertificateSecretId = s.Model.CertificateSecretId
	m.ApiServerUrl = s.Model.ApiServerUrl
	m.ApiServerPort = s.Model.ApiServerPort
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.EntitySource = s.Model.EntitySource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatekubeclusterinsightdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntitySource {
	case "OKE_CLUSTER":
		mm := UpdateOkeKubeClusterInsightDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateKubeClusterInsightDetails: %s.", m.EntitySource)
		return *m, nil
	}
}

// GetServiceAccount returns ServiceAccount
func (m updatekubeclusterinsightdetails) GetServiceAccount() *string {
	return m.ServiceAccount
}

// GetTokenSecretId returns TokenSecretId
func (m updatekubeclusterinsightdetails) GetTokenSecretId() *string {
	return m.TokenSecretId
}

// GetCertificateSecretId returns CertificateSecretId
func (m updatekubeclusterinsightdetails) GetCertificateSecretId() *string {
	return m.CertificateSecretId
}

// GetApiServerUrl returns ApiServerUrl
func (m updatekubeclusterinsightdetails) GetApiServerUrl() *string {
	return m.ApiServerUrl
}

// GetApiServerPort returns ApiServerPort
func (m updatekubeclusterinsightdetails) GetApiServerPort() *int {
	return m.ApiServerPort
}

// GetFreeformTags returns FreeformTags
func (m updatekubeclusterinsightdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m updatekubeclusterinsightdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m updatekubeclusterinsightdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatekubeclusterinsightdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
