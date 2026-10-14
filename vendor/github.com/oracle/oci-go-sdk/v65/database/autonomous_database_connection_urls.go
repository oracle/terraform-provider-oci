// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutonomousDatabaseConnectionUrls The URLs for accessing various Autonomous AI Database built-in tools (like Oracle Application Express (APEX), SQL Developer Web etc) with a browser from a Compute instance within your VCN or that has a direct connection to your VCN. Note that these URLs are also available through the Autonomous AI Databases (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) console.
// Example: `{"sqlDevWebUrl": "https://<hostname>/ords...", "apexUrl", "https://<hostname>/ords..."}`
type AutonomousDatabaseConnectionUrls struct {

	// Oracle SQL Developer Web URL.
	SqlDevWebUrl *string `mandatory:"false" json:"sqlDevWebUrl"`

	// Oracle Application Express (APEX) URL.
	ApexUrl *string `mandatory:"false" json:"apexUrl"`

	// Oracle Machine Learning user management URL. Applicable to Autonomous AI Database Serverless only.
	MachineLearningUserManagementUrl *string `mandatory:"false" json:"machineLearningUserManagementUrl"`

	// The URL of the Graph Studio for the Autonomous AI Database. Applicable to Autonomous AI Database Serverless only.
	GraphStudioUrl *string `mandatory:"false" json:"graphStudioUrl"`

	// The URL of the MongoDB API for the Autonomous AI Database.
	MongoDbUrl *string `mandatory:"false" json:"mongoDbUrl"`

	// The URL of the Oracle Machine Learning (OML) Notebook for the Autonomous AI Database. Applicable to Autonomous AI Database Serverless only.
	MachineLearningNotebookUrl *string `mandatory:"false" json:"machineLearningNotebookUrl"`

	// The Oracle REST Data Services (ORDS) URL of the Web Access for the Autonomous AI Database.
	OrdsUrl *string `mandatory:"false" json:"ordsUrl"`

	// The URL of the Database Transforms for the Autonomous AI Database. Applicable to Autonomous AI Database Serverless only.
	DatabaseTransformsUrl *string `mandatory:"false" json:"databaseTransformsUrl"`

	// Essbase URL for the Autonomous AI Database. Applicable to Autonomous AI Database Serverless only.
	EssbaseUrl *string `mandatory:"false" json:"essbaseUrl"`

	// The URL of the Spatial Studio for the Autonomous AI Database. Applicable to Autonomous AI Database Serverless only.
	SpatialStudioUrl *string `mandatory:"false" json:"spatialStudioUrl"`
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
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
