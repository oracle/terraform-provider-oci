// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DbClusterBackupSourceFromDbCluster A backup created from a shared-storage DB cluster.
type DbClusterBackupSourceFromDbCluster struct {

	// The OCID of the shared-storage DB cluster from which the backup was created.
	DbClusterId *string `mandatory:"true" json:"dbClusterId"`
}

func (m DbClusterBackupSourceFromDbCluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbClusterBackupSourceFromDbCluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DbClusterBackupSourceFromDbCluster) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDbClusterBackupSourceFromDbCluster DbClusterBackupSourceFromDbCluster
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeDbClusterBackupSourceFromDbCluster
	}{
		"DBCLUSTER",
		(MarshalTypeDbClusterBackupSourceFromDbCluster)(m),
	}

	return json.Marshal(&s)
}
