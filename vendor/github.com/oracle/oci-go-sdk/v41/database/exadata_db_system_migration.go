// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// ExadataDbSystemMigration Information about the Exadata DB system migration. The migration is used to move the Exadata Cloud Service instance from the DB system resource model to the new cloud Exadata infrastructure resource model.
type ExadataDbSystemMigration struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DB system.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the cloud VM cluster.
	CloudVmClusterId *string `mandatory:"true" json:"cloudVmClusterId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the cloud Exadata infrastructure.
	CloudExadataInfrastructureId *string `mandatory:"true" json:"cloudExadataInfrastructureId"`

	// The details of addtional resources related to the migration.
	AdditionalMigrations []ExadataDbSystemMigrationSummary `mandatory:"false" json:"additionalMigrations"`
}

func (m ExadataDbSystemMigration) String() string {
	return common.PointerString(m)
}
