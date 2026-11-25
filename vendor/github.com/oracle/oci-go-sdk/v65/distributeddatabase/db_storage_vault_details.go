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

// DbStorageVaultDetails Details of the request to create exascale db vault storage for shard or catalog of the distributed database.
type DbStorageVaultDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for Db Storage Vault.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Total storage capacity in GB for vault storage.
	HighCapacityDatabaseStorage *int `mandatory:"true" json:"highCapacityDatabaseStorage"`

	// The size of additional Flash Cache in percentage of High Capacity database storage.
	AdditionalFlashCacheInPercent *int `mandatory:"false" json:"additionalFlashCacheInPercent"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	// Indicates if autoscale feature is enabled for the Database Storage Vault. The default value is `FALSE`.
	IsAutoscaleEnabled *bool `mandatory:"false" json:"isAutoscaleEnabled"`

	// The maximum limit, in gigabytes, to which the Vault storage size can automatically scale when auto scaling is enabled for the Database Storage Vault.
	AutoscaleLimitInGBs *int `mandatory:"false" json:"autoscaleLimitInGBs"`
}

func (m DbStorageVaultDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbStorageVaultDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
