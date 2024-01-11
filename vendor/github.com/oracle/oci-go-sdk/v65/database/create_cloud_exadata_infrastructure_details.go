// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateCloudExadataInfrastructureDetails Request to create cloud Exadata infrastructure. Applies to Exadata Cloud Service instances only.
type CreateCloudExadataInfrastructureDetails struct {

	// The availability domain where the cloud Exadata infrastructure is located.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the cloud Exadata infrastructure resource. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The shape of the cloud Exadata infrastructure resource.
	Shape *string `mandatory:"true" json:"shape"`

	// The number of compute servers for the cloud Exadata infrastructure.
	ComputeCount *int `mandatory:"false" json:"computeCount"`

	// The number of storage servers for the cloud Exadata infrastructure.
	StorageCount *int `mandatory:"false" json:"storageCount"`

	MaintenanceWindow *MaintenanceWindow `mandatory:"false" json:"maintenanceWindow"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the cluster placement group of the Exadata Infrastructure.
	ClusterPlacementGroupId *string `mandatory:"false" json:"clusterPlacementGroupId"`

	// Customer contacts.
	CustomerContacts []CustomerContact `mandatory:"false" json:"customerContacts"`
}

func (m CreateCloudExadataInfrastructureDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCloudExadataInfrastructureDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
