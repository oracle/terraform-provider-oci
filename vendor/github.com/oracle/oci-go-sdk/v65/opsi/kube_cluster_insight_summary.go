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

// KubeClusterInsightSummary Summary of a Kubernetes cluster insight resource.
type KubeClusterInsightSummary interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Kubernetes cluster insight resource.
	GetId() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// The Kubernetes cluster name.
	GetName() *string

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

	// Indicates the status of a Kubernetes cluster insight in Ops Insights
	GetStatus() ResourceStatusEnum

	// The time the the Kubernetes cluster insight was first enabled. An RFC3339 formatted datetime string
	GetTimeCreated() *common.SDKTime

	// The current state of the Kubernetes cluster insight.
	GetLifecycleState() LifecycleStateEnum

	// The Kubernetes cluster service account.
	GetServiceAccount() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the K8s cluster token secret id.
	GetTokenSecretId() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the K8s cluster certificate secret id.
	GetCertificateSecretId() *string

	// List of K8s cluster node pools registered within Ops Insights.
	GetNodePools() []KubeClusterNodePool

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}

	// The time the Kubernetes cluster insight was updated. An RFC3339 formatted datetime string
	GetTimeUpdated() *common.SDKTime

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecycleDetails() *string
}

type kubeclusterinsightsummary struct {
	JsonData            []byte
	ServiceAccount      *string                           `mandatory:"false" json:"serviceAccount"`
	TokenSecretId       *string                           `mandatory:"false" json:"tokenSecretId"`
	CertificateSecretId *string                           `mandatory:"false" json:"certificateSecretId"`
	NodePools           []KubeClusterNodePool             `mandatory:"false" json:"nodePools"`
	SystemTags          map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	TimeUpdated         *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	LifecycleDetails    *string                           `mandatory:"false" json:"lifecycleDetails"`
	Id                  *string                           `mandatory:"true" json:"id"`
	CompartmentId       *string                           `mandatory:"true" json:"compartmentId"`
	Name                *string                           `mandatory:"true" json:"name"`
	ApiServerUrl        *string                           `mandatory:"true" json:"apiServerUrl"`
	ApiServerPort       *int                              `mandatory:"true" json:"apiServerPort"`
	FreeformTags        map[string]string                 `mandatory:"true" json:"freeformTags"`
	DefinedTags         map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`
	Status              ResourceStatusEnum                `mandatory:"true" json:"status"`
	TimeCreated         *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	LifecycleState      LifecycleStateEnum                `mandatory:"true" json:"lifecycleState"`
	EntitySource        string                            `json:"entitySource"`
}

// UnmarshalJSON unmarshals json
func (m *kubeclusterinsightsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerkubeclusterinsightsummary kubeclusterinsightsummary
	s := struct {
		Model Unmarshalerkubeclusterinsightsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.Name = s.Model.Name
	m.ApiServerUrl = s.Model.ApiServerUrl
	m.ApiServerPort = s.Model.ApiServerPort
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.Status = s.Model.Status
	m.TimeCreated = s.Model.TimeCreated
	m.LifecycleState = s.Model.LifecycleState
	m.ServiceAccount = s.Model.ServiceAccount
	m.TokenSecretId = s.Model.TokenSecretId
	m.CertificateSecretId = s.Model.CertificateSecretId
	m.NodePools = s.Model.NodePools
	m.SystemTags = s.Model.SystemTags
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.EntitySource = s.Model.EntitySource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *kubeclusterinsightsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntitySource {
	case "OKE_CLUSTER":
		mm := OkeKubeClusterInsightSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for KubeClusterInsightSummary: %s.", m.EntitySource)
		return *m, nil
	}
}

// GetServiceAccount returns ServiceAccount
func (m kubeclusterinsightsummary) GetServiceAccount() *string {
	return m.ServiceAccount
}

// GetTokenSecretId returns TokenSecretId
func (m kubeclusterinsightsummary) GetTokenSecretId() *string {
	return m.TokenSecretId
}

// GetCertificateSecretId returns CertificateSecretId
func (m kubeclusterinsightsummary) GetCertificateSecretId() *string {
	return m.CertificateSecretId
}

// GetNodePools returns NodePools
func (m kubeclusterinsightsummary) GetNodePools() []KubeClusterNodePool {
	return m.NodePools
}

// GetSystemTags returns SystemTags
func (m kubeclusterinsightsummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetTimeUpdated returns TimeUpdated
func (m kubeclusterinsightsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleDetails returns LifecycleDetails
func (m kubeclusterinsightsummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetId returns Id
func (m kubeclusterinsightsummary) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m kubeclusterinsightsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetName returns Name
func (m kubeclusterinsightsummary) GetName() *string {
	return m.Name
}

// GetApiServerUrl returns ApiServerUrl
func (m kubeclusterinsightsummary) GetApiServerUrl() *string {
	return m.ApiServerUrl
}

// GetApiServerPort returns ApiServerPort
func (m kubeclusterinsightsummary) GetApiServerPort() *int {
	return m.ApiServerPort
}

// GetFreeformTags returns FreeformTags
func (m kubeclusterinsightsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m kubeclusterinsightsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetStatus returns Status
func (m kubeclusterinsightsummary) GetStatus() ResourceStatusEnum {
	return m.Status
}

// GetTimeCreated returns TimeCreated
func (m kubeclusterinsightsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetLifecycleState returns LifecycleState
func (m kubeclusterinsightsummary) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

func (m kubeclusterinsightsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m kubeclusterinsightsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResourceStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetResourceStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
