// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NoSQL Database API
//
// The control plane API for NoSQL Database Cloud Service HTTPS
// provides endpoints to perform NDCS operations, including creation
// and deletion of tables and indexes; population and access of data
// in tables; and access of table usage metrics.
//

package nosql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TableCollection Results of ListTables.
type TableCollection struct {

	// A page of TableSummary objects.
	Items []TableSummary `mandatory:"false" json:"items"`

	// The maximum number of reclaimable tables allowed in the tenancy.
	MaxAutoReclaimableTables *int `mandatory:"false" json:"maxAutoReclaimableTables"`

	// The current number of reclaimable tables in the tenancy.
	AutoReclaimableTables *int `mandatory:"false" json:"autoReclaimableTables"`

	// The current number of on demand capacity tables in the tenancy.
	OnDemandCapacityTables *int `mandatory:"false" json:"onDemandCapacityTables"`

	// The maximum number of on demand capacity tables allowed in the tenancy.
	MaxOnDemandCapacityTables *int `mandatory:"false" json:"maxOnDemandCapacityTables"`

	// An array of regions that are available for replication.
	AvailableReplicationRegions []string `mandatory:"false" json:"availableReplicationRegions"`
}

func (m TableCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TableCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
