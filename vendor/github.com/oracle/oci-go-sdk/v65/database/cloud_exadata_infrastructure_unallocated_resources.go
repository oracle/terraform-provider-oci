// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CloudExadataInfrastructureUnallocatedResources Details of unallocated resources of the Cloud Exadata infrastructure. Applies to Cloud Exadata infrastructure instances only.
type CloudExadataInfrastructureUnallocatedResources struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud Exadata infrastructure.
	CloudExadataInfrastructureId *string `mandatory:"true" json:"cloudExadataInfrastructureId"`

	// The user-friendly name for the Cloud Exadata infrastructure. The name does not need to be unique.
	CloudExadataInfrastructureDisplayName *string `mandatory:"true" json:"cloudExadataInfrastructureDisplayName"`

	// The minimum amount of unallocated storage available across all nodes in the infrastructure.
	LocalStorageInGbs *int `mandatory:"false" json:"localStorageInGbs"`

	// The minimum amount of unallocated ocpus available across all nodes in the infrastructure.
	Ocpus *int `mandatory:"false" json:"ocpus"`

	// The minimum amount of unallocated memory available across all nodes in the infrastructure.
	MemoryInGBs *int `mandatory:"false" json:"memoryInGBs"`

	// Total unallocated exadata storage in the infrastructure in TBs.
	ExadataStorageInTBs *float64 `mandatory:"false" json:"exadataStorageInTBs"`

	// The list of Cloud Autonomous VM Clusters on the Infrastructure and their associated unallocated resources details.
	CloudAutonomousVmClusters []CloudAutonomousVmClusterResourceDetails `mandatory:"false" json:"cloudAutonomousVmClusters"`
}

func (m CloudExadataInfrastructureUnallocatedResources) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudExadataInfrastructureUnallocatedResources) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
