// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// RemoveVirtualMachineFromVmClusterDetails Details of removing Virtual Machines from the VM Cluster. Applies to Exadata Cloud@Customer instances only.
type RemoveVirtualMachineFromVmClusterDetails struct {

	// The list of Exacc DB servers for the cluster to be removed.
	DbServers []DbServerDetails `mandatory:"true" json:"dbServers"`
}

func (m RemoveVirtualMachineFromVmClusterDetails) String() string {
	return common.PointerString(m)
}
