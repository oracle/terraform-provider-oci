// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateBlueGreenDeploymentTargetDbSystemDetails Target DB System overrides for a blue/green deployment.
type CreateBlueGreenDeploymentTargetDbSystemDetails struct {

	// Target MySQL engine version.
	MysqlVersion *string `mandatory:"true" json:"mysqlVersion"`

	// The shape of the target DB System. The shape determines resources
	// allocated to the DB System - CPU cores and memory for VM shapes;
	// CPU cores, memory and storage for non-VM (or bare metal) shapes.
	// To get a list of shapes, use the
	// ListShapes operation.
	ShapeName *string `mandatory:"false" json:"shapeName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the
	// configuration to apply to the target DB System. If omitted, the
	// target DB System inherits the source DB System configuration.
	ConfigurationId *string `mandatory:"false" json:"configurationId"`

	// Initial data storage size in GiBs for the target DB System.
	// If omitted, the target DB System uses the source DB System storage size.
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`
}

func (m CreateBlueGreenDeploymentTargetDbSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateBlueGreenDeploymentTargetDbSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
