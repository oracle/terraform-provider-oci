// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// DbSystemInternalDetails DbSystem internal details.
type DbSystemInternalDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Database system.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// The display name of the Database system.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of compartment of the Database system.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	Databases *ExternalDatabaseCollection `mandatory:"false" json:"databases"`

	Cluster *ExternalCluster `mandatory:"false" json:"cluster"`

	ClusterInstances *ExternalClusterInstanceCollection `mandatory:"false" json:"clusterInstances"`

	Asm *ExternalAsm `mandatory:"false" json:"asm"`

	AsmInstances *ExternalAsmInstanceCollection `mandatory:"false" json:"asmInstances"`

	Listeners *ExternalListenerCollection `mandatory:"false" json:"listeners"`

	DbNodes *ExternalDbNodeCollection `mandatory:"false" json:"dbNodes"`

	DbHomes *ExternalDbHomeCollection `mandatory:"false" json:"dbHomes"`
}

func (m DbSystemInternalDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbSystemInternalDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
