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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreatePeComanagedExadataVmclusterDetails The information of the VM Cluster which contains databases. Either an opsiPrivateEndpointId or dbmPrivateEndpointId must be specified. If the dbmPrivateEndpointId is specified, a new Operations Insights private endpoint will be created.
type CreatePeComanagedExadataVmclusterDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM Cluster.
	VmclusterId *string `mandatory:"true" json:"vmclusterId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OPSI private endpoint
	OpsiPrivateEndpointId *string `mandatory:"false" json:"opsiPrivateEndpointId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Management private endpoint
	DbmPrivateEndpointId *string `mandatory:"false" json:"dbmPrivateEndpointId"`

	// The databases that belong to the VM Cluster
	MemberDatabaseDetails []CreatePeComanagedDatabaseInsightDetails `mandatory:"false" json:"memberDatabaseDetails"`

	// Exadata VMCluster type
	VmClusterType ExadataVmClusterTypeEnum `mandatory:"false" json:"vmClusterType,omitempty"`

	// The autonomous databases that belong to the Autonomous VM Cluster
	MemberAutonomousDetails []CreateAutonomousDatabaseInsightDetails `mandatory:"false" json:"memberAutonomousDetails"`
}

func (m CreatePeComanagedExadataVmclusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePeComanagedExadataVmclusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExadataVmClusterTypeEnum(string(m.VmClusterType)); !ok && m.VmClusterType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VmClusterType: %s. Supported values are: %s.", m.VmClusterType, strings.Join(GetExadataVmClusterTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
