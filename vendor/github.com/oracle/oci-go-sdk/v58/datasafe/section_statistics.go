// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// SectionStatistics Statistics showing the number of findings with a particular risk level for each category.
type SectionStatistics struct {

	// The number of targets that contributed to the counts at this risk level.
	TargetsCount *int `mandatory:"false" json:"targetsCount"`

	// The number of findings in the Auditing category.
	AuditingFindingsCount *int `mandatory:"false" json:"auditingFindingsCount"`

	// The number of findings in the Authorization Control category.
	AuthorizationControlFindingsCount *int `mandatory:"false" json:"authorizationControlFindingsCount"`

	// The number of findings in the Data Encryption category.
	DataEncryptionFindingsCount *int `mandatory:"false" json:"dataEncryptionFindingsCount"`

	// The number of findings in the Database Configuration category.
	DbConfigurationFindingsCount *int `mandatory:"false" json:"dbConfigurationFindingsCount"`

	// The number of findings in the Fine-Grained Access Control category.
	FineGrainedAccessControlFindingsCount *int `mandatory:"false" json:"fineGrainedAccessControlFindingsCount"`

	// The number of findings in the Privileges and Roles category.
	PrivilegesAndRolesFindingsCount *int `mandatory:"false" json:"privilegesAndRolesFindingsCount"`

	// The number of findings in the User Accounts category.
	UserAccountsFindingsCount *int `mandatory:"false" json:"userAccountsFindingsCount"`
}

func (m SectionStatistics) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SectionStatistics) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
