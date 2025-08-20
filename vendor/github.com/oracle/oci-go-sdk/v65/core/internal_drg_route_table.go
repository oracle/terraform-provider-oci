// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InternalDrgRouteTable A collection of `InternalDrgRouteRule` objects. It is used to offload DRG functionality (primarily routing, but
// up-to-and-including all additional features associated with DRG attachments) onto the VCN Dataplane.
type InternalDrgRouteTable struct {

	// The label of the DRG attachment.
	DrgAttachmentLabel *int64 `mandatory:"true" json:"drgAttachmentLabel"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG which contains this route table.
	DrgId *string `mandatory:"true" json:"drgId"`

	// The collection of rules which will be used by VCN Dataplane to route DRG traffic.
	Rules []InternalDrgRouteRule `mandatory:"true" json:"rules"`

	// The sequence number for the DRG route table update (version of the DRG route table). Only supported for partitioned route tables.
	SequenceNumber *int64 `mandatory:"false" json:"sequenceNumber"`

	// The total number of shards/partitions for the specified DRG route table. Only supported for partitioned route tables.
	ShardsTotal *int64 `mandatory:"false" json:"shardsTotal"`

	// The shard number for the DRG route table shard. Only supported for partitioned route tables.
	ShardId *int64 `mandatory:"false" json:"shardId"`

	// The DRG route table partitions's physical availability domain. This attribute will be null if this is a non-partitioned DRG route table.
	// Example: `PHX-AD-1`
	InternalAvailabilityDomain *string `mandatory:"false" json:"internalAvailabilityDomain"`

	// The type of DrgRouteTable
	DrgRouteTableType InternalDrgRouteTableDrgRouteTableTypeEnum `mandatory:"false" json:"drgRouteTableType,omitempty"`
}

func (m InternalDrgRouteTable) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternalDrgRouteTable) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingInternalDrgRouteTableDrgRouteTableTypeEnum(string(m.DrgRouteTableType)); !ok && m.DrgRouteTableType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DrgRouteTableType: %s. Supported values are: %s.", m.DrgRouteTableType, strings.Join(GetInternalDrgRouteTableDrgRouteTableTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InternalDrgRouteTableDrgRouteTableTypeEnum Enum with underlying type: string
type InternalDrgRouteTableDrgRouteTableTypeEnum string

// Set of constants representing the allowable values for InternalDrgRouteTableDrgRouteTableTypeEnum
const (
	InternalDrgRouteTableDrgRouteTableTypeDefault            InternalDrgRouteTableDrgRouteTableTypeEnum = "DEFAULT"
	InternalDrgRouteTableDrgRouteTableTypeUnderlayAccessProd InternalDrgRouteTableDrgRouteTableTypeEnum = "UNDERLAY_ACCESS_PROD"
	InternalDrgRouteTableDrgRouteTableTypeUnderlayAccessTest InternalDrgRouteTableDrgRouteTableTypeEnum = "UNDERLAY_ACCESS_TEST"
)

var mappingInternalDrgRouteTableDrgRouteTableTypeEnum = map[string]InternalDrgRouteTableDrgRouteTableTypeEnum{
	"DEFAULT":              InternalDrgRouteTableDrgRouteTableTypeDefault,
	"UNDERLAY_ACCESS_PROD": InternalDrgRouteTableDrgRouteTableTypeUnderlayAccessProd,
	"UNDERLAY_ACCESS_TEST": InternalDrgRouteTableDrgRouteTableTypeUnderlayAccessTest,
}

var mappingInternalDrgRouteTableDrgRouteTableTypeEnumLowerCase = map[string]InternalDrgRouteTableDrgRouteTableTypeEnum{
	"default":              InternalDrgRouteTableDrgRouteTableTypeDefault,
	"underlay_access_prod": InternalDrgRouteTableDrgRouteTableTypeUnderlayAccessProd,
	"underlay_access_test": InternalDrgRouteTableDrgRouteTableTypeUnderlayAccessTest,
}

// GetInternalDrgRouteTableDrgRouteTableTypeEnumValues Enumerates the set of values for InternalDrgRouteTableDrgRouteTableTypeEnum
func GetInternalDrgRouteTableDrgRouteTableTypeEnumValues() []InternalDrgRouteTableDrgRouteTableTypeEnum {
	values := make([]InternalDrgRouteTableDrgRouteTableTypeEnum, 0)
	for _, v := range mappingInternalDrgRouteTableDrgRouteTableTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInternalDrgRouteTableDrgRouteTableTypeEnumStringValues Enumerates the set of values in String for InternalDrgRouteTableDrgRouteTableTypeEnum
func GetInternalDrgRouteTableDrgRouteTableTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"UNDERLAY_ACCESS_PROD",
		"UNDERLAY_ACCESS_TEST",
	}
}

// GetMappingInternalDrgRouteTableDrgRouteTableTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInternalDrgRouteTableDrgRouteTableTypeEnum(val string) (InternalDrgRouteTableDrgRouteTableTypeEnum, bool) {
	enum, ok := mappingInternalDrgRouteTableDrgRouteTableTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
