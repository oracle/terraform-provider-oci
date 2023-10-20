// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.cloud.oracle.com/Content/application-dependency-management/home.htm).
//

package adm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NetworkConfiguration A network configuration defines the required network characteristics for an ADM remediation recipe.
// A network configuration is required if the build service is one of: GitHub Actions, GitLab Pipeline, or Jenkins Pipeline.
type NetworkConfiguration struct {

	// The Oracle Cloud Identifier (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)) of the subnet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The list of Oracle Cloud Identifiers (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)) corresponding to Network Security Groups.
	NsgIds []string `mandatory:"false" json:"nsgIds"`
}

func (m NetworkConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NetworkConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
