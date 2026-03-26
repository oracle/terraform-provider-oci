// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// HostK8sPodStatus K8s Pod Status
type HostK8sPodStatus struct {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// Pod name
	Name *string `mandatory:"false" json:"name"`

	// Pod namespace
	Namespace *string `mandatory:"false" json:"namespace"`

	// Generate name
	GenerateName *string `mandatory:"false" json:"generateName"`

	// Status phase
	StatusPhase *string `mandatory:"false" json:"statusPhase"`

	// Host IPs
	HostIps *string `mandatory:"false" json:"hostIps"`

	// Pod IPs
	PodIps *string `mandatory:"false" json:"podIps"`

	// Start time
	StartTime *string `mandatory:"false" json:"startTime"`

	// Ready state
	ReadyState *string `mandatory:"false" json:"readyState"`

	// Container Id
	ContainerId *string `mandatory:"false" json:"containerId"`

	// Pod uid
	Uid *string `mandatory:"false" json:"uid"`

	// Pod parent uid
	ParentUid *string `mandatory:"false" json:"parentUid"`

	// Kubernetes OPSI ID
	KubeInsightsId *string `mandatory:"false" json:"kubeInsightsId"`
}

// GetTimeCollected returns TimeCollected
func (m HostK8sPodStatus) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m HostK8sPodStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostK8sPodStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HostK8sPodStatus) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostK8sPodStatus HostK8sPodStatus
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeHostK8sPodStatus
	}{
		"HOST_K8S_POD_STATUS",
		(MarshalTypeHostK8sPodStatus)(m),
	}

	return json.Marshal(&s)
}
