// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDbManagementPrivateEndpointDetails The details used to create a new Database Management private endpoint.
type CreateDbManagementPrivateEndpointDetails struct {

	// The display name of the Database Management private endpoint.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Specifies whether the Database Management private endpoint will be used for Oracle Databases in a cluster.
	IsCluster *bool `mandatory:"false" json:"isCluster"`

	// The description of the private endpoint.
	Description *string `mandatory:"false" json:"description"`

	// The OCIDs of the Network Security Groups to which the Database Management private endpoint belongs.
	NsgIds []string `mandatory:"false" json:"nsgIds"`
}

func (m CreateDbManagementPrivateEndpointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDbManagementPrivateEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
