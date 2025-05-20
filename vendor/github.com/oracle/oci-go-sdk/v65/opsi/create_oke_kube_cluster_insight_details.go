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

// CreateOkeKubeClusterInsightDetails The information about the OKE (Oracle Kubernetes Engine) Kubernetes cluster to be analyzed.
type CreateOkeKubeClusterInsightDetails struct {

	// Compartment identifier of the Kubernetes cluster insight resource
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The Kubernetes cluster service account.
	ServiceAccount *string `mandatory:"true" json:"serviceAccount"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the K8s cluster token secret id.
	TokenSecretId *string `mandatory:"true" json:"tokenSecretId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the K8s cluster certificate secret id.
	CertificateSecretId *string `mandatory:"true" json:"certificateSecretId"`

	// The Kubernetes cluster API server URL.
	ApiServerUrl *string `mandatory:"true" json:"apiServerUrl"`

	// The Kubernetes cluster API Server port.
	ApiServerPort *int `mandatory:"true" json:"apiServerPort"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OKE (Oracle Kubernetes Engine) Infrastructure.
	KubeClusterInfraId *string `mandatory:"true" json:"kubeClusterInfraId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

// GetCompartmentId returns CompartmentId
func (m CreateOkeKubeClusterInsightDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetServiceAccount returns ServiceAccount
func (m CreateOkeKubeClusterInsightDetails) GetServiceAccount() *string {
	return m.ServiceAccount
}

// GetTokenSecretId returns TokenSecretId
func (m CreateOkeKubeClusterInsightDetails) GetTokenSecretId() *string {
	return m.TokenSecretId
}

// GetCertificateSecretId returns CertificateSecretId
func (m CreateOkeKubeClusterInsightDetails) GetCertificateSecretId() *string {
	return m.CertificateSecretId
}

// GetApiServerUrl returns ApiServerUrl
func (m CreateOkeKubeClusterInsightDetails) GetApiServerUrl() *string {
	return m.ApiServerUrl
}

// GetApiServerPort returns ApiServerPort
func (m CreateOkeKubeClusterInsightDetails) GetApiServerPort() *int {
	return m.ApiServerPort
}

// GetFreeformTags returns FreeformTags
func (m CreateOkeKubeClusterInsightDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateOkeKubeClusterInsightDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateOkeKubeClusterInsightDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOkeKubeClusterInsightDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateOkeKubeClusterInsightDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateOkeKubeClusterInsightDetails CreateOkeKubeClusterInsightDetails
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeCreateOkeKubeClusterInsightDetails
	}{
		"OKE_CLUSTER",
		(MarshalTypeCreateOkeKubeClusterInsightDetails)(m),
	}

	return json.Marshal(&s)
}
