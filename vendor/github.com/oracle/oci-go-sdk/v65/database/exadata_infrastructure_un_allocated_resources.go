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

// ExadataInfrastructureUnAllocatedResources Un allocated resources details of the Exadata Cloud@Customer infrastructure. Applies to Exadata Cloud@Customer instances only.
type ExadataInfrastructureUnAllocatedResources struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the Exadata Cloud@Customer infrastructure. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The minimum amount of un allocated storage that is available across all nodes in the infrastructure.
	LocalStorageInGbs *int `mandatory:"false" json:"localStorageInGbs"`

	// The minimum amount of un allocated ocpus that is available across all nodes in the infrastructure.
	Ocpus *int `mandatory:"false" json:"ocpus"`

	// The minimum amount of un allocated memory that is available across all nodes in the infrastructure.
	MemoryInGBs *int `mandatory:"false" json:"memoryInGBs"`

	// Total unallocated exadata storage in the infrastructure in TBs.
	ExadataStorageInTBs *float64 `mandatory:"false" json:"exadataStorageInTBs"`

	// The list of Autonomous VM Clusters on the Infra and their associated unallocated resources details
	AutonomousVmClusters []AutonomousVmClusterResourceDetails `mandatory:"false" json:"autonomousVmClusters"`
}

func (m ExadataInfrastructureUnAllocatedResources) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadataInfrastructureUnAllocatedResources) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
