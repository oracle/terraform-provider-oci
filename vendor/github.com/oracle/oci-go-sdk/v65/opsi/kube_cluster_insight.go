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

// KubeClusterInsight Kubernetes cluster insight resource.
type KubeClusterInsight interface {

	// Kubernetes cluster insight identifier
	GetId() *string

	// Compartment identifier of the Kubernetes cluster insight resource
	GetCompartmentId() *string

	// The Kubernetes cluster name.
	GetName() *string

	// Indicates the status of a Kubernetes cluster insight in Ops Insights
	GetStatus() ResourceStatusEnum

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

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

	// The Kubernetes cluster API server URL.
	GetApiServerUrl() *string

	// The Kubernetes cluster API Server port.
	GetApiServerPort() *int

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

type kubeclusterinsight struct {
	JsonData            []byte
	ServiceAccount      *string                           `mandatory:"false" json:"serviceAccount"`
	TokenSecretId       *string                           `mandatory:"false" json:"tokenSecretId"`
	CertificateSecretId *string                           `mandatory:"false" json:"certificateSecretId"`
	ApiServerUrl        *string                           `mandatory:"false" json:"apiServerUrl"`
	ApiServerPort       *int                              `mandatory:"false" json:"apiServerPort"`
	NodePools           []KubeClusterNodePool             `mandatory:"false" json:"nodePools"`
	SystemTags          map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	TimeUpdated         *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	LifecycleDetails    *string                           `mandatory:"false" json:"lifecycleDetails"`
	Id                  *string                           `mandatory:"true" json:"id"`
	CompartmentId       *string                           `mandatory:"true" json:"compartmentId"`
	Name                *string                           `mandatory:"true" json:"name"`
	Status              ResourceStatusEnum                `mandatory:"true" json:"status"`
	FreeformTags        map[string]string                 `mandatory:"true" json:"freeformTags"`
	DefinedTags         map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`
	TimeCreated         *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	LifecycleState      LifecycleStateEnum                `mandatory:"true" json:"lifecycleState"`
	EntitySource        string                            `json:"entitySource"`
}

// UnmarshalJSON unmarshals json
func (m *kubeclusterinsight) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerkubeclusterinsight kubeclusterinsight
	s := struct {
		Model Unmarshalerkubeclusterinsight
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.CompartmentId = s.Model.CompartmentId
	m.Name = s.Model.Name
	m.Status = s.Model.Status
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.TimeCreated = s.Model.TimeCreated
	m.LifecycleState = s.Model.LifecycleState
	m.ServiceAccount = s.Model.ServiceAccount
	m.TokenSecretId = s.Model.TokenSecretId
	m.CertificateSecretId = s.Model.CertificateSecretId
	m.ApiServerUrl = s.Model.ApiServerUrl
	m.ApiServerPort = s.Model.ApiServerPort
	m.NodePools = s.Model.NodePools
	m.SystemTags = s.Model.SystemTags
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.EntitySource = s.Model.EntitySource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *kubeclusterinsight) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntitySource {
	case "OKE_CLUSTER":
		mm := OkeKubeClusterInsight{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for KubeClusterInsight: %s.", m.EntitySource)
		return *m, nil
	}
}

// GetServiceAccount returns ServiceAccount
func (m kubeclusterinsight) GetServiceAccount() *string {
	return m.ServiceAccount
}

// GetTokenSecretId returns TokenSecretId
func (m kubeclusterinsight) GetTokenSecretId() *string {
	return m.TokenSecretId
}

// GetCertificateSecretId returns CertificateSecretId
func (m kubeclusterinsight) GetCertificateSecretId() *string {
	return m.CertificateSecretId
}

// GetApiServerUrl returns ApiServerUrl
func (m kubeclusterinsight) GetApiServerUrl() *string {
	return m.ApiServerUrl
}

// GetApiServerPort returns ApiServerPort
func (m kubeclusterinsight) GetApiServerPort() *int {
	return m.ApiServerPort
}

// GetNodePools returns NodePools
func (m kubeclusterinsight) GetNodePools() []KubeClusterNodePool {
	return m.NodePools
}

// GetSystemTags returns SystemTags
func (m kubeclusterinsight) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetTimeUpdated returns TimeUpdated
func (m kubeclusterinsight) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleDetails returns LifecycleDetails
func (m kubeclusterinsight) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetId returns Id
func (m kubeclusterinsight) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m kubeclusterinsight) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetName returns Name
func (m kubeclusterinsight) GetName() *string {
	return m.Name
}

// GetStatus returns Status
func (m kubeclusterinsight) GetStatus() ResourceStatusEnum {
	return m.Status
}

// GetFreeformTags returns FreeformTags
func (m kubeclusterinsight) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m kubeclusterinsight) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetTimeCreated returns TimeCreated
func (m kubeclusterinsight) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetLifecycleState returns LifecycleState
func (m kubeclusterinsight) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

func (m kubeclusterinsight) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m kubeclusterinsight) ValidateEnumValue() (bool, error) {
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
