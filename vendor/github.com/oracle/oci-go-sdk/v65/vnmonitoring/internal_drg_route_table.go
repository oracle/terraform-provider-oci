// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InternalDrgRouteTable A collection of `InternalDrgRouteRule` objects. It is used to offload DRG functionality (primarily routing, but
// up-to-and-including all additional features associated with DRG attachments) onto the VCN Dataplane.
type InternalDrgRouteTable struct {

	// The label of the drg attachment.
	DrgAttachmentLabel *int64 `mandatory:"true" json:"drgAttachmentLabel"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG which contains this route table.
	DrgId *string `mandatory:"true" json:"drgId"`

	// The collection of rules which will be used by VCN Dataplane to route DRG traffic.
	Rules []InternalDrgRouteRule `mandatory:"true" json:"rules"`

	// The sequence number for the DRG Route Table update (version of the DRG Route Table). Only supported for partitioned route tables.
	SequenceNumber *int64 `mandatory:"false" json:"sequenceNumber"`

	// The total number of shards/partitions for the specified DRG Route Table. Only supported for partitioned route tables.
	ShardsTotal *int64 `mandatory:"false" json:"shardsTotal"`

	// The shard number for the DRG Route Table shard. Only supported for partitioned route tables.
	ShardId *int64 `mandatory:"false" json:"shardId"`

	// The DRG Route Table partitions's physical availability domain. This attribute will be null if this is a non-partitioned DRG Route Table.
	// Example: `PHX-AD-1`
	InternalAvailabilityDomain *string `mandatory:"false" json:"internalAvailabilityDomain"`
}

func (m InternalDrgRouteTable) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternalDrgRouteTable) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
