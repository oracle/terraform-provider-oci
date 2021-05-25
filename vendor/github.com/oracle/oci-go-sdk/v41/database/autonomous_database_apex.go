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

// AutonomousDatabaseApex Oracle APEX Application Development is a low-code development platform that enables you to build scalable, secure enterprise apps, with world-class features. Autonomous Database with the APEX workload type is optimized to support APEX development.
type AutonomousDatabaseApex struct {

	// The Oracle APEX Application Development version.
	ApexVersion *string `mandatory:"false" json:"apexVersion"`

	// The Oracle REST Data Services (ORDS) version.
	OrdsVersion *string `mandatory:"false" json:"ordsVersion"`
}

func (m AutonomousDatabaseApex) String() string {
	return common.PointerString(m)
}
