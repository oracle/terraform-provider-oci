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

// AvmAcdResourceStats Associated autonomous container databases usages.
type AvmAcdResourceStats struct {

	// The user-friendly name for the Autonomous Container Database. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous Container Database.
	Id *string `mandatory:"false" json:"id"`

	// CPUs/cores assigned to Autonomous Databases in the ACD instances.
	ProvisionedCpus *float32 `mandatory:"false" json:"provisionedCpus"`

	// The number of CPU cores available.
	AvailableCpus *float32 `mandatory:"false" json:"availableCpus"`

	// CPUs/cores assigned to the ACD instance. Sum of provisioned, reserved and reclaimable CPUs/ cores
	// to the ACD instance.
	UsedCpus *float32 `mandatory:"false" json:"usedCpus"`

	// CPUs/cores reserved for scalability, resilliency and other overheads. This includes failover,
	// autoscaling and idle instance overhead.
	ReservedCpus *float32 `mandatory:"false" json:"reservedCpus"`

	// CPUs/cores that continue to be included in the count of OCPUs available to the
	// Autonomous Container Database even after one of its Autonomous Database is terminated
	// or scaled down. You can release them to the available OCPUs at its parent AVMC level by
	// restarting the Autonomous Container Database.
	ReclaimableCpus *float32 `mandatory:"false" json:"reclaimableCpus"`
}

func (m AvmAcdResourceStats) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AvmAcdResourceStats) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
