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

// MigrateDrgDetails The request to update a `DrgAttachment` and migrate the `Drg` to the destination DRG type.
type MigrateDrgDetails struct {

	// Type of the `Drg` to be migrated to.
	DestinationDrgType MigrateDrgDetailsDestinationDrgTypeEnum `mandatory:"true" json:"destinationDrgType"`

	// The DRG attachment's Oracle ID (OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
	DrgAttachmentId *string `mandatory:"false" json:"drgAttachmentId"`

	// NextHop target's MPLS label.
	MplsLabel *string `mandatory:"false" json:"mplsLabel"`

	// The string in the form ASN:mplsLabel.
	RouteTarget *string `mandatory:"false" json:"routeTarget"`
}

func (m MigrateDrgDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MigrateDrgDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMigrateDrgDetailsDestinationDrgTypeEnum(string(m.DestinationDrgType)); !ok && m.DestinationDrgType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DestinationDrgType: %s. Supported values are: %s.", m.DestinationDrgType, strings.Join(GetMigrateDrgDetailsDestinationDrgTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MigrateDrgDetailsDestinationDrgTypeEnum Enum with underlying type: string
type MigrateDrgDetailsDestinationDrgTypeEnum string

// Set of constants representing the allowable values for MigrateDrgDetailsDestinationDrgTypeEnum
const (
	MigrateDrgDetailsDestinationDrgTypeClassical  MigrateDrgDetailsDestinationDrgTypeEnum = "DRG_CLASSICAL"
	MigrateDrgDetailsDestinationDrgTypeTransitHub MigrateDrgDetailsDestinationDrgTypeEnum = "DRG_TRANSIT_HUB"
)

var mappingMigrateDrgDetailsDestinationDrgTypeEnum = map[string]MigrateDrgDetailsDestinationDrgTypeEnum{
	"DRG_CLASSICAL":   MigrateDrgDetailsDestinationDrgTypeClassical,
	"DRG_TRANSIT_HUB": MigrateDrgDetailsDestinationDrgTypeTransitHub,
}

var mappingMigrateDrgDetailsDestinationDrgTypeEnumLowerCase = map[string]MigrateDrgDetailsDestinationDrgTypeEnum{
	"drg_classical":   MigrateDrgDetailsDestinationDrgTypeClassical,
	"drg_transit_hub": MigrateDrgDetailsDestinationDrgTypeTransitHub,
}

// GetMigrateDrgDetailsDestinationDrgTypeEnumValues Enumerates the set of values for MigrateDrgDetailsDestinationDrgTypeEnum
func GetMigrateDrgDetailsDestinationDrgTypeEnumValues() []MigrateDrgDetailsDestinationDrgTypeEnum {
	values := make([]MigrateDrgDetailsDestinationDrgTypeEnum, 0)
	for _, v := range mappingMigrateDrgDetailsDestinationDrgTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMigrateDrgDetailsDestinationDrgTypeEnumStringValues Enumerates the set of values in String for MigrateDrgDetailsDestinationDrgTypeEnum
func GetMigrateDrgDetailsDestinationDrgTypeEnumStringValues() []string {
	return []string{
		"DRG_CLASSICAL",
		"DRG_TRANSIT_HUB",
	}
}

// GetMappingMigrateDrgDetailsDestinationDrgTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMigrateDrgDetailsDestinationDrgTypeEnum(val string) (MigrateDrgDetailsDestinationDrgTypeEnum, bool) {
	enum, ok := mappingMigrateDrgDetailsDestinationDrgTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
