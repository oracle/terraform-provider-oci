// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KubernetesClusterDataSourceSummaryItem A Kubernetes collection data source summary.
type KubernetesClusterDataSourceSummaryItem struct {

	// Data source type and name identifier.
	Key *string `mandatory:"true" json:"key"`

	// Unique name of the dataSource.
	Name *string `mandatory:"true" json:"name"`

	// true if the Kubernetes cluster type is Daemon set
	IsDaemonSet *bool `mandatory:"false" json:"isDaemonSet"`
}

// GetKey returns Key
func (m KubernetesClusterDataSourceSummaryItem) GetKey() *string {
	return m.Key
}

// GetName returns Name
func (m KubernetesClusterDataSourceSummaryItem) GetName() *string {
	return m.Name
}

func (m KubernetesClusterDataSourceSummaryItem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KubernetesClusterDataSourceSummaryItem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m KubernetesClusterDataSourceSummaryItem) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeKubernetesClusterDataSourceSummaryItem KubernetesClusterDataSourceSummaryItem
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeKubernetesClusterDataSourceSummaryItem
	}{
		"KUBERNETES_CLUSTER",
		(MarshalTypeKubernetesClusterDataSourceSummaryItem)(m),
	}

	return json.Marshal(&s)
}
