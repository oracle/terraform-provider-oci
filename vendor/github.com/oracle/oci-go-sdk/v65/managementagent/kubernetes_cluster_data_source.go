// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.cloud.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KubernetesClusterDataSource A Kubernetes cluster data source.
type KubernetesClusterDataSource struct {

	// Identifier for DataSource. This represents the type and name for the data source associated with the Management Agent.
	Key *string `mandatory:"true" json:"key"`

	// Unique name of the DataSource.
	Name *string `mandatory:"true" json:"name"`

	// Compartment owning this DataSource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the DataSource was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the DataSource data was last received. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The Kubernetes namespace
	Namespace *string `mandatory:"true" json:"namespace"`

	// If the Kubernetes cluster type is Daemon set then this will be set to true.
	IsDaemonSet *bool `mandatory:"false" json:"isDaemonSet"`

	// State of the DataSource.
	State LifecycleStatesEnum `mandatory:"true" json:"state"`
}

// GetKey returns Key
func (m KubernetesClusterDataSource) GetKey() *string {
	return m.Key
}

// GetName returns Name
func (m KubernetesClusterDataSource) GetName() *string {
	return m.Name
}

// GetCompartmentId returns CompartmentId
func (m KubernetesClusterDataSource) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetState returns State
func (m KubernetesClusterDataSource) GetState() LifecycleStatesEnum {
	return m.State
}

// GetTimeCreated returns TimeCreated
func (m KubernetesClusterDataSource) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m KubernetesClusterDataSource) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m KubernetesClusterDataSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KubernetesClusterDataSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLifecycleStatesEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m KubernetesClusterDataSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeKubernetesClusterDataSource KubernetesClusterDataSource
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeKubernetesClusterDataSource
	}{
		"KUBERNETES_CLUSTER",
		(MarshalTypeKubernetesClusterDataSource)(m),
	}

	return json.Marshal(&s)
}
