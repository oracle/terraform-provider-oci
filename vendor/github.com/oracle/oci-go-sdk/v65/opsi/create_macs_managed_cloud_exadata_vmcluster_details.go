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

// CreateMacsManagedCloudExadataVmclusterDetails The information of the VM Cluster which contains databases.
type CreateMacsManagedCloudExadataVmclusterDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM Cluster.
	VmclusterId *string `mandatory:"true" json:"vmclusterId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The databases that belong to the VM Cluster
	MemberDatabaseDetails []CreateMacsManagedCloudDatabaseInsightDetails `mandatory:"false" json:"memberDatabaseDetails"`
}

// GetVmclusterId returns VmclusterId
func (m CreateMacsManagedCloudExadataVmclusterDetails) GetVmclusterId() *string {
	return m.VmclusterId
}

// GetCompartmentId returns CompartmentId
func (m CreateMacsManagedCloudExadataVmclusterDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m CreateMacsManagedCloudExadataVmclusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMacsManagedCloudExadataVmclusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateMacsManagedCloudExadataVmclusterDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateMacsManagedCloudExadataVmclusterDetails CreateMacsManagedCloudExadataVmclusterDetails
	s := struct {
		DiscriminatorParam string `json:"vmClusterType"`
		MarshalTypeCreateMacsManagedCloudExadataVmclusterDetails
	}{
		"vmCluster",
		(MarshalTypeCreateMacsManagedCloudExadataVmclusterDetails)(m),
	}

	return json.Marshal(&s)
}
