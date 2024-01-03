// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutonomousVmResourceUsage Autonomous VM usage statistics.
type AutonomousVmResourceUsage struct {

	// The user-friendly name for the Autonomous VM cluster. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous VM Cluster.
	Id *string `mandatory:"false" json:"id"`

	// The number of CPU cores alloted to the Autonomous Container Databases in an Cloud Autonomous VM cluster.
	UsedCpus *float32 `mandatory:"false" json:"usedCpus"`

	// The number of CPU cores available.
	AvailableCpus *float32 `mandatory:"false" json:"availableCpus"`

	// CPU cores that continue to be included in the count of OCPUs available to the
	// Autonomous Container Database even after one of its Autonomous Database is
	// terminated or scaled down. You can release them to the available OCPUs at its
	// parent AVMC level by restarting the Autonomous Container Database.
	ReclaimableCpus *float32 `mandatory:"false" json:"reclaimableCpus"`

	// The number of CPUs provisioned in an Autonomous VM Cluster.
	ProvisionedCpus *float32 `mandatory:"false" json:"provisionedCpus"`

	// The number of CPUs reserved in an Autonomous VM Cluster.
	ReservedCpus *float32 `mandatory:"false" json:"reservedCpus"`

	// Associated Autonomous Container Database Usages.
	AutonomousContainerDatabaseUsage []AvmAcdResourceStats `mandatory:"false" json:"autonomousContainerDatabaseUsage"`
}

func (m AutonomousVmResourceUsage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousVmResourceUsage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
