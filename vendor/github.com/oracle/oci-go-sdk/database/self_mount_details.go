// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// SelfMountDetails Used for creating NFS Self mount backup destinations for non-autonomous ExaCC.
type SelfMountDetails struct {

	// The local directory path on each VM cluster node where the NFS server location is mounted. The local directory path and the NFS server location must each be the same across all of the VM cluster nodes. Ensure that the NFS mount is maintained continuously on all of the VM cluster nodes.
	LocalMountPointPath *string `mandatory:"true" json:"localMountPointPath"`
}

func (m SelfMountDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m SelfMountDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSelfMountDetails SelfMountDetails
	s := struct {
		DiscriminatorParam string `json:"mountType"`
		MarshalTypeSelfMountDetails
	}{
		"SELF_MOUNT",
		(MarshalTypeSelfMountDetails)(m),
	}

	return json.Marshal(&s)
}
