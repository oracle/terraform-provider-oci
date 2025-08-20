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

// CreateInternalDrgRouteTableDetails Details to create and update internal DRG route table. Partitioned DRG Route Tables are supported when specifying the sharding information.
type CreateInternalDrgRouteTableDetails struct {

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

	// Timings info for upstream services to track route propagation latency
	UpstreamTimings []UpstreamTiming `mandatory:"false" json:"upstreamTimings"`

	// The type of DrgRouteTable
	DrgRouteTableType CreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnum `mandatory:"false" json:"drgRouteTableType,omitempty"`
}

func (m CreateInternalDrgRouteTableDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateInternalDrgRouteTableDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnum(string(m.DrgRouteTableType)); !ok && m.DrgRouteTableType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DrgRouteTableType: %s. Supported values are: %s.", m.DrgRouteTableType, strings.Join(GetCreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnum Enum with underlying type: string
type CreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnum string

// Set of constants representing the allowable values for CreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnum
const (
	CreateInternalDrgRouteTableDetailsDrgRouteTableTypeDefault            CreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnum = "DEFAULT"
	CreateInternalDrgRouteTableDetailsDrgRouteTableTypeUnderlayAccessProd CreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnum = "UNDERLAY_ACCESS_PROD"
	CreateInternalDrgRouteTableDetailsDrgRouteTableTypeUnderlayAccessTest CreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnum = "UNDERLAY_ACCESS_TEST"
)

var mappingCreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnum = map[string]CreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnum{
	"DEFAULT":              CreateInternalDrgRouteTableDetailsDrgRouteTableTypeDefault,
	"UNDERLAY_ACCESS_PROD": CreateInternalDrgRouteTableDetailsDrgRouteTableTypeUnderlayAccessProd,
	"UNDERLAY_ACCESS_TEST": CreateInternalDrgRouteTableDetailsDrgRouteTableTypeUnderlayAccessTest,
}

var mappingCreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnumLowerCase = map[string]CreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnum{
	"default":              CreateInternalDrgRouteTableDetailsDrgRouteTableTypeDefault,
	"underlay_access_prod": CreateInternalDrgRouteTableDetailsDrgRouteTableTypeUnderlayAccessProd,
	"underlay_access_test": CreateInternalDrgRouteTableDetailsDrgRouteTableTypeUnderlayAccessTest,
}

// GetCreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnumValues Enumerates the set of values for CreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnum
func GetCreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnumValues() []CreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnum {
	values := make([]CreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnum, 0)
	for _, v := range mappingCreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnumStringValues Enumerates the set of values in String for CreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnum
func GetCreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"UNDERLAY_ACCESS_PROD",
		"UNDERLAY_ACCESS_TEST",
	}
}

// GetMappingCreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnum(val string) (CreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnum, bool) {
	enum, ok := mappingCreateInternalDrgRouteTableDetailsDrgRouteTableTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
