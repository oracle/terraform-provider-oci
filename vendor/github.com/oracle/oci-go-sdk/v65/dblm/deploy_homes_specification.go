// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DeployHomesSpecification Specification for deploying an Oracle home.
type DeployHomesSpecification struct {

	// If true, run the CVU validations.
	IsSkipCvuChecks *bool `mandatory:"false" json:"isSkipCvuChecks"`

	// Location of the new ORACLE_HOME. E.g., "/u01/app/oracle/product/19.0.0/dbhome"
	//  Following variables are supported in this field:
	//  %ORACLE_BASE% Oracle base location of the resource
	//  %OH_PARENT_DIR% Parent directory of the current Oracle home location of the resource
	//  %VERSION% Base version of the Image (E.g., "19.0.0")
	//  %VERSION_NAME% Version of the image
	//  %U% Unique identifier to install more than one Oracle home in the same host(s)
	//  E.g., "%ORACLE_BASE%/%VERSION%/dbhome_%U%"
	NewOracleHome *string `mandatory:"false" json:"newOracleHome"`

	Schedule *Schedule `mandatory:"false" json:"schedule"`

	Resources *Resources `mandatory:"false" json:"resources"`

	// Prefix to use when deploying an Oracle home.
	HomeNamePrefix *string `mandatory:"false" json:"homeNamePrefix"`
}

func (m DeployHomesSpecification) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeployHomesSpecification) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
