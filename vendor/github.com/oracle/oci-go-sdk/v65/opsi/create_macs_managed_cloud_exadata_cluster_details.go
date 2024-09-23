// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// CreateMacsManagedCloudExadataClusterDetails The information of the VM Cluster which contains databases.
type CreateMacsManagedCloudExadataClusterDetails interface {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM Cluster.
	GetVmclusterId() *string

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string
}

type createmacsmanagedcloudexadataclusterdetails struct {
	JsonData      []byte
	VmclusterId   *string `mandatory:"true" json:"vmclusterId"`
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
	VmClusterType string  `json:"vmClusterType"`
}

// UnmarshalJSON unmarshals json
func (m *createmacsmanagedcloudexadataclusterdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatemacsmanagedcloudexadataclusterdetails createmacsmanagedcloudexadataclusterdetails
	s := struct {
		Model Unmarshalercreatemacsmanagedcloudexadataclusterdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.VmclusterId = s.Model.VmclusterId
	m.CompartmentId = s.Model.CompartmentId
	m.VmClusterType = s.Model.VmClusterType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createmacsmanagedcloudexadataclusterdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.VmClusterType {
	case "vmCluster":
		mm := CreateMacsManagedCloudExadataVmclusterDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateMacsManagedCloudExadataClusterDetails: %s.", m.VmClusterType)
		return *m, nil
	}
}

// GetVmclusterId returns VmclusterId
func (m createmacsmanagedcloudexadataclusterdetails) GetVmclusterId() *string {
	return m.VmclusterId
}

// GetCompartmentId returns CompartmentId
func (m createmacsmanagedcloudexadataclusterdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m createmacsmanagedcloudexadataclusterdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createmacsmanagedcloudexadataclusterdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
