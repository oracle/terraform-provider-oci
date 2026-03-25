// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// OkeKubeClusterInsightSummary Summary of an OKE (Oracle Kubernetes Engine) Kubernetes cluster insight resource.
type OkeKubeClusterInsightSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Kubernetes cluster insight resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The Kubernetes cluster name.
	Name *string `mandatory:"true" json:"name"`

	// The Kubernetes cluster API server URL.
	ApiServerUrl *string `mandatory:"true" json:"apiServerUrl"`

	// The Kubernetes cluster API Server port.
	ApiServerPort *int `mandatory:"true" json:"apiServerPort"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The time the the Kubernetes cluster insight was first enabled. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OKE (Oracle Kubernetes Engine) Infrastructure.
	KubeClusterInfraId *string `mandatory:"true" json:"kubeClusterInfraId"`

	// The Kubernetes cluster service account.
	ServiceAccount *string `mandatory:"false" json:"serviceAccount"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the K8s cluster token secret id.
	TokenSecretId *string `mandatory:"false" json:"tokenSecretId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the K8s cluster certificate secret id.
	CertificateSecretId *string `mandatory:"false" json:"certificateSecretId"`

	// List of K8s cluster node pools registered within Ops Insights.
	NodePools []KubeClusterNodePool `mandatory:"false" json:"nodePools"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The time the Kubernetes cluster insight was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Indicates the status of a Kubernetes cluster insight in Ops Insights
	Status ResourceStatusEnum `mandatory:"true" json:"status"`

	// The current state of the Kubernetes cluster insight.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m OkeKubeClusterInsightSummary) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m OkeKubeClusterInsightSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetName returns Name
func (m OkeKubeClusterInsightSummary) GetName() *string {
	return m.Name
}

// GetServiceAccount returns ServiceAccount
func (m OkeKubeClusterInsightSummary) GetServiceAccount() *string {
	return m.ServiceAccount
}

// GetTokenSecretId returns TokenSecretId
func (m OkeKubeClusterInsightSummary) GetTokenSecretId() *string {
	return m.TokenSecretId
}

// GetCertificateSecretId returns CertificateSecretId
func (m OkeKubeClusterInsightSummary) GetCertificateSecretId() *string {
	return m.CertificateSecretId
}

// GetApiServerUrl returns ApiServerUrl
func (m OkeKubeClusterInsightSummary) GetApiServerUrl() *string {
	return m.ApiServerUrl
}

// GetApiServerPort returns ApiServerPort
func (m OkeKubeClusterInsightSummary) GetApiServerPort() *int {
	return m.ApiServerPort
}

// GetNodePools returns NodePools
func (m OkeKubeClusterInsightSummary) GetNodePools() []KubeClusterNodePool {
	return m.NodePools
}

// GetFreeformTags returns FreeformTags
func (m OkeKubeClusterInsightSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m OkeKubeClusterInsightSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m OkeKubeClusterInsightSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetStatus returns Status
func (m OkeKubeClusterInsightSummary) GetStatus() ResourceStatusEnum {
	return m.Status
}

// GetTimeCreated returns TimeCreated
func (m OkeKubeClusterInsightSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m OkeKubeClusterInsightSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m OkeKubeClusterInsightSummary) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m OkeKubeClusterInsightSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

func (m OkeKubeClusterInsightSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OkeKubeClusterInsightSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingResourceStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetResourceStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OkeKubeClusterInsightSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOkeKubeClusterInsightSummary OkeKubeClusterInsightSummary
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeOkeKubeClusterInsightSummary
	}{
		"OKE_CLUSTER",
		(MarshalTypeOkeKubeClusterInsightSummary)(m),
	}

	return json.Marshal(&s)
}
