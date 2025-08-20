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

// HostK8sWorkload K8s Workload
type HostK8sWorkload struct {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// Workload name
	Name *string `mandatory:"false" json:"name"`

	// Workload kind
	Kind *string `mandatory:"false" json:"kind"`

	// Namespace
	Namespace *string `mandatory:"false" json:"namespace"`

	// Workload generation
	Generation *string `mandatory:"false" json:"generation"`

	// Workload creation timestamp
	CreationTimestamp *string `mandatory:"false" json:"creationTimestamp"`

	// Workload labels
	Labels *string `mandatory:"false" json:"labels"`

	// Replicas
	Replicas *string `mandatory:"false" json:"replicas"`

	// Workload selector match labels
	SelectorMatchLabels *string `mandatory:"false" json:"selectorMatchLabels"`

	// Workload selector match labels
	TemplateMetadata *string `mandatory:"false" json:"templateMetadata"`

	// Workload UID
	Uid *string `mandatory:"false" json:"uid"`

	// Kubernetes OPSI ID
	KubeInsightsId *string `mandatory:"false" json:"kubeInsightsId"`
}

// GetTimeCollected returns TimeCollected
func (m HostK8sWorkload) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m HostK8sWorkload) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostK8sWorkload) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HostK8sWorkload) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostK8sWorkload HostK8sWorkload
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeHostK8sWorkload
	}{
		"HOST_K8S_WORKLOAD",
		(MarshalTypeHostK8sWorkload)(m),
	}

	return json.Marshal(&s)
}
