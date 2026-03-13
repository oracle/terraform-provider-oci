// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateCatalogPeerWithExadbXsNewVaultAndClusterDetails Details required for creation of ExaDbXs based catalog peer.
type CreateCatalogPeerWithExadbXsNewVaultAndClusterDetails struct {

	// The name of the availability domain that the distributed database shard will be located in.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	DbStorageVaultDetails *DbStorageVaultDetails `mandatory:"true" json:"dbStorageVaultDetails"`

	VmClusterDetails *VmClusterDetails `mandatory:"true" json:"vmClusterDetails"`

	// The protectionMode for the catalog peer.
	ProtectionMode DistributedDbProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// The redo transport type to use for this Data Guard association.
	TransportType DistributedDbTransportTypeEnum `mandatory:"false" json:"transportType,omitempty"`
}

func (m CreateCatalogPeerWithExadbXsNewVaultAndClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCatalogPeerWithExadbXsNewVaultAndClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDistributedDbProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetDistributedDbProtectionModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDistributedDbTransportTypeEnum(string(m.TransportType)); !ok && m.TransportType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TransportType: %s. Supported values are: %s.", m.TransportType, strings.Join(GetDistributedDbTransportTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
