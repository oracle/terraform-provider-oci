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

// OkeKubeClusterInsight OKE (Oracle Kubernetes Engine) Kubernetes cluster insight resource.
type OkeKubeClusterInsight struct {

	// Kubernetes cluster insight identifier
	Id *string `mandatory:"true" json:"id"`

	// Compartment identifier of the Kubernetes cluster insight resource
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The Kubernetes cluster name.
	Name *string `mandatory:"true" json:"name"`

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

	// The Kubernetes cluster API server URL.
	ApiServerUrl *string `mandatory:"false" json:"apiServerUrl"`

	// The Kubernetes cluster API Server port.
	ApiServerPort *int `mandatory:"false" json:"apiServerPort"`

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
func (m OkeKubeClusterInsight) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m OkeKubeClusterInsight) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetName returns Name
func (m OkeKubeClusterInsight) GetName() *string {
	return m.Name
}

// GetServiceAccount returns ServiceAccount
func (m OkeKubeClusterInsight) GetServiceAccount() *string {
	return m.ServiceAccount
}

// GetTokenSecretId returns TokenSecretId
func (m OkeKubeClusterInsight) GetTokenSecretId() *string {
	return m.TokenSecretId
}

// GetCertificateSecretId returns CertificateSecretId
func (m OkeKubeClusterInsight) GetCertificateSecretId() *string {
	return m.CertificateSecretId
}

// GetApiServerUrl returns ApiServerUrl
func (m OkeKubeClusterInsight) GetApiServerUrl() *string {
	return m.ApiServerUrl
}

// GetApiServerPort returns ApiServerPort
func (m OkeKubeClusterInsight) GetApiServerPort() *int {
	return m.ApiServerPort
}

// GetNodePools returns NodePools
func (m OkeKubeClusterInsight) GetNodePools() []KubeClusterNodePool {
	return m.NodePools
}

// GetStatus returns Status
func (m OkeKubeClusterInsight) GetStatus() ResourceStatusEnum {
	return m.Status
}

// GetFreeformTags returns FreeformTags
func (m OkeKubeClusterInsight) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m OkeKubeClusterInsight) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m OkeKubeClusterInsight) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetTimeCreated returns TimeCreated
func (m OkeKubeClusterInsight) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m OkeKubeClusterInsight) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m OkeKubeClusterInsight) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m OkeKubeClusterInsight) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

func (m OkeKubeClusterInsight) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OkeKubeClusterInsight) ValidateEnumValue() (bool, error) {
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
func (m OkeKubeClusterInsight) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOkeKubeClusterInsight OkeKubeClusterInsight
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeOkeKubeClusterInsight
	}{
		"OKE_CLUSTER",
		(MarshalTypeOkeKubeClusterInsight)(m),
	}

	return json.Marshal(&s)
}
