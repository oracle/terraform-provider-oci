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

// CreateCatalogPeerWithBaseDbDetails Details required for creation of peer catalog.
type CreateCatalogPeerWithBaseDbDetails struct {

	// The name of the availability domain that the peer base database system will be located in.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The identifier of the subnet for the peer Dbsystem instance.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The protection mode for the Data Guard association's primary and standby Base database based catalog.
	ProtectionMode BaseDbProtectionModeEnum `mandatory:"true" json:"protectionMode"`

	// The redo transport type to use for Data Guard association for Base database based catalog.
	TransportType BaseDbTransportTypeEnum `mandatory:"true" json:"transportType"`

	// Fault Domain in which this base database is provisioned.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// Flag to enable active Data Guard.
	IsActiveDataGuardEnabled *bool `mandatory:"false" json:"isActiveDataGuardEnabled"`
}

func (m CreateCatalogPeerWithBaseDbDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCatalogPeerWithBaseDbDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBaseDbProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetBaseDbProtectionModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseDbTransportTypeEnum(string(m.TransportType)); !ok && m.TransportType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TransportType: %s. Supported values are: %s.", m.TransportType, strings.Join(GetBaseDbTransportTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
