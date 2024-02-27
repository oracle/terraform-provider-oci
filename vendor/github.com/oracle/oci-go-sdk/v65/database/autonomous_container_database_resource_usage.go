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

// AutonomousContainerDatabaseResourceUsage Associated autonomous container databases usages.
type AutonomousContainerDatabaseResourceUsage struct {

	// The user-friendly name for the Autonomous Container Database. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous Container Database.
	Id *string `mandatory:"false" json:"id"`

	// Number of CPUs that are reclaimable or released to the AVMC on Autonomous Container Database restart.
	ReclaimableCpus *float32 `mandatory:"false" json:"reclaimableCpus"`

	// CPUs available for provisioning or scaling an Autonomous Database in the Autonomous Container Database.
	AvailableCpus *float32 `mandatory:"false" json:"availableCpus"`

	// Largest provisionable ADB in the Autonomous Container Database.
	LargestProvisionableAutonomousDatabaseInCpus *float32 `mandatory:"false" json:"largestProvisionableAutonomousDatabaseInCpus"`

	// CPUs / cores assigned to ADBs in the Autonomous Container Database.
	ProvisionedCpus *float32 `mandatory:"false" json:"provisionedCpus"`

	// CPUs / cores reserved for scalability, resilliency and other overheads.
	// This includes failover, autoscaling and idle instance overhead.
	ReservedCpus *float32 `mandatory:"false" json:"reservedCpus"`

	// CPUs / cores assigned to the Autonomous Container Database. Sum of provisioned,
	// reserved and reclaimable CPUs/ cores.
	UsedCpus *float32 `mandatory:"false" json:"usedCpus"`

	// Valid list of provisionable CPUs for Autonomous Database.
	ProvisionableCpus []float32 `mandatory:"false" json:"provisionableCpus"`

	// List of autonomous container database resource usage per autonomous virtual machine.
	AutonomousContainerDatabaseVmUsage []AcdAvmResourceStats `mandatory:"false" json:"autonomousContainerDatabaseVmUsage"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m AutonomousContainerDatabaseResourceUsage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousContainerDatabaseResourceUsage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
