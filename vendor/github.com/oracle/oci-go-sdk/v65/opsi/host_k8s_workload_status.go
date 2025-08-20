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

// HostK8sWorkloadStatus K8s Workload Status
type HostK8sWorkloadStatus struct {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// Workload name
	Name *string `mandatory:"false" json:"name"`

	// Observed generation
	ObservedGeneration *string `mandatory:"false" json:"observedGeneration"`

	// Replicas
	Replicas *string `mandatory:"false" json:"replicas"`

	// Updated replicas
	UpdatedReplicas *string `mandatory:"false" json:"updatedReplicas"`

	// Ready replicas
	ReadyReplicas *string `mandatory:"false" json:"readyReplicas"`

	// Available replicas
	AvailableReplicas *string `mandatory:"false" json:"availableReplicas"`

	// Workload uid
	Uid *string `mandatory:"false" json:"uid"`

	// Kubernetes OPSI ID
	KubeInsightsId *string `mandatory:"false" json:"kubeInsightsId"`
}

// GetTimeCollected returns TimeCollected
func (m HostK8sWorkloadStatus) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m HostK8sWorkloadStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostK8sWorkloadStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HostK8sWorkloadStatus) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostK8sWorkloadStatus HostK8sWorkloadStatus
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeHostK8sWorkloadStatus
	}{
		"HOST_K8S_WORKLOAD_STATUS",
		(MarshalTypeHostK8sWorkloadStatus)(m),
	}

	return json.Marshal(&s)
}
