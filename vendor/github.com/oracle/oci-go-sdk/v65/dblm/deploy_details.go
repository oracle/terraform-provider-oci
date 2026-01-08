// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// DeployDetails Home deployment information.
type DeployDetails struct {

	// Path to the new ORACLE_HOME location.
	NewHome *string `mandatory:"false" json:"newHome"`

	// ORACLE_BASE location of the resource.
	OracleBase *string `mandatory:"false" json:"oracleBase"`

	// Name of the image version used to deploy the Oracle home.
	ImageVersionName *string `mandatory:"false" json:"imageVersionName"`

	// The prefix for the home name used.
	HomeNamePrefix *string `mandatory:"false" json:"homeNamePrefix"`

	// Comma-separated list of Oracle groups used to configure the Oracle home.
	Groups *string `mandatory:"false" json:"groups"`
}

func (m DeployDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeployDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
