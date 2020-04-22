// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AutonomousDatabaseConnectionUrls The URLs for accessing Oracle Application Express (APEX) and SQL Developer Web with a browser from a Compute instance within your VCN or that has a direct connection to your VCN. Note that these URLs are provided by the console only for databases on dedicated Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adbddoverview.htm).
// Example: `{"sqlDevWebUrl": "https://<hostname>/ords...", "apexUrl", "https://<hostname>/ords..."}`
type AutonomousDatabaseConnectionUrls struct {

	// Oracle SQL Developer Web URL.
	SqlDevWebUrl *string `mandatory:"false" json:"sqlDevWebUrl"`

	// Oracle Application Express (APEX) URL.
	ApexUrl *string `mandatory:"false" json:"apexUrl"`

	// Oracle Machine Learning user management URL.
	MachineLearningUserManagementUrl *string `mandatory:"false" json:"machineLearningUserManagementUrl"`
}

func (m AutonomousDatabaseConnectionUrls) String() string {
	return common.PointerString(m)
}
