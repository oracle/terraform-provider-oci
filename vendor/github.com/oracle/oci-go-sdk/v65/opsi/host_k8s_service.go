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

// HostK8sService K8s Services
type HostK8sService struct {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// Service name
	Name *string `mandatory:"false" json:"name"`

	// Namespace
	Namespace *string `mandatory:"false" json:"namespace"`

	// Workload creation timestamp
	CreationTimestamp *string `mandatory:"false" json:"creationTimestamp"`

	// Cluster IPs
	ClusterIps *string `mandatory:"false" json:"clusterIps"`

	// Internal traffic policy
	InternalTrafficPolicy *string `mandatory:"false" json:"internalTrafficPolicy"`

	// IP families
	IpFamilies *string `mandatory:"false" json:"ipFamilies"`

	// Family policy
	FamilyPolicy *string `mandatory:"false" json:"familyPolicy"`

	// Ports
	Ports *string `mandatory:"false" json:"ports"`

	// Selector
	Selector *string `mandatory:"false" json:"selector"`

	// Session affinity
	SessionAffinity *string `mandatory:"false" json:"sessionAffinity"`

	// Service type
	Type *string `mandatory:"false" json:"type"`

	// Load balancer status
	LbStatus *string `mandatory:"false" json:"lbStatus"`

	// Service UID
	Uid *string `mandatory:"false" json:"uid"`

	// Kubernetes OPSI ID
	KubeInsightsId *string `mandatory:"false" json:"kubeInsightsId"`
}

// GetTimeCollected returns TimeCollected
func (m HostK8sService) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m HostK8sService) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostK8sService) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HostK8sService) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostK8sService HostK8sService
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeHostK8sService
	}{
		"HOST_K8S_SERVICE",
		(MarshalTypeHostK8sService)(m),
	}

	return json.Marshal(&s)
}
