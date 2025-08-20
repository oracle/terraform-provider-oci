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

// HostK8sContainer K8s Containers
type HostK8sContainer struct {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// Container name
	Name *string `mandatory:"false" json:"name"`

	// Container image
	Image *string `mandatory:"false" json:"image"`

	// Container start command
	Command *string `mandatory:"false" json:"command"`

	// Container command arguments
	Arguments *string `mandatory:"false" json:"arguments"`

	// Container Ports
	Ports *string `mandatory:"false" json:"ports"`

	// Container resource limits
	ResourceLimits *string `mandatory:"false" json:"resourceLimits"`

	// Container resource requests
	ResourceRequests *string `mandatory:"false" json:"resourceRequests"`

	// Container volume mounts
	VolumeMounts *string `mandatory:"false" json:"volumeMounts"`

	// Container Restart policy
	RestartPolicy *string `mandatory:"false" json:"restartPolicy"`

	// Scheduler Name
	SchedulerName *string `mandatory:"false" json:"schedulerName"`

	// Container Tolerations
	Tolerations *string `mandatory:"false" json:"tolerations"`

	// Container strategy
	Strategy *string `mandatory:"false" json:"strategy"`

	// Parent UID specifies parent workload definition for this container
	ParentUid *string `mandatory:"false" json:"parentUid"`

	// Kubernetes OPSI ID
	KubeInsightsId *string `mandatory:"false" json:"kubeInsightsId"`
}

// GetTimeCollected returns TimeCollected
func (m HostK8sContainer) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m HostK8sContainer) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostK8sContainer) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HostK8sContainer) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostK8sContainer HostK8sContainer
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeHostK8sContainer
	}{
		"HOST_K8S_CONTAINER",
		(MarshalTypeHostK8sContainer)(m),
	}

	return json.Marshal(&s)
}
