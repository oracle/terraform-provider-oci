// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateMigrationAssetDetails Details of the new migration asset.
type CreateMigrationAssetDetails struct {

	// OCID of an asset for an inventory.
	InventoryAssetId *string `mandatory:"true" json:"inventoryAssetId"`

	// OCID of the associated migration.
	MigrationId *string `mandatory:"true" json:"migrationId"`

	// Availability domain
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// Replication compartment identifier
	ReplicationCompartmentId *string `mandatory:"true" json:"replicationCompartmentId"`

	// Name of snapshot bucket
	SnapShotBucketName *string `mandatory:"true" json:"snapShotBucketName"`

	// A user-friendly name. If empty, then source asset name will be used. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Replication schedule identifier
	ReplicationScheduleId *string `mandatory:"false" json:"replicationScheduleId"`

	// List of migration assets that depends on this asset.
	DependsOn []string `mandatory:"false" json:"dependsOn"`
}

func (m CreateMigrationAssetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMigrationAssetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
