// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// AutonomousDatabaseConnectionUrls The URLs for accessing Oracle Application Express (APEX) and SQL Developer Web with a browser from a Compute instance within your VCN or that has a direct connection to your VCN. Note that these URLs are provided by the console only for databases on dedicated Exadata infrastructure (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html).
// Example: `{"sqlDevWebUrl": "https://<hostname>/ords...", "apexUrl", "https://<hostname>/ords..."}`
type AutonomousDatabaseConnectionUrls struct {

	// Oracle SQL Developer Web URL.
	SqlDevWebUrl *string `mandatory:"false" json:"sqlDevWebUrl"`

	// Oracle Application Express (APEX) URL.
	ApexUrl *string `mandatory:"false" json:"apexUrl"`

	// Oracle Machine Learning user management URL.
	MachineLearningUserManagementUrl *string `mandatory:"false" json:"machineLearningUserManagementUrl"`

	// The URL of the Graph Studio for the Autonomous Database.
	GraphStudioUrl *string `mandatory:"false" json:"graphStudioUrl"`
}

func (m AutonomousDatabaseConnectionUrls) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousDatabaseConnectionUrls) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
