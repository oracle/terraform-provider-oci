// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SshUser An entry which includes the user's OCID and their respective full public key in PEM format.
type SshUser struct {

	// The ocid of the user allowed to ssh Must be a string of a valid user ocid.
	UserId *string `mandatory:"true" json:"userId"`

	// The full public key associated with the user. Must be a full public key in PEM format, and match an API key associated with the user.
	PubKey *string `mandatory:"true" json:"pubKey"`

	// The tenancy OCID for a cross-tenancy user. If absent, the tenancy of the notebook will be used. If used, must be a string of a valid tenancy ocid.
	UserTenancy *string `mandatory:"false" json:"userTenancy"`
}

func (m SshUser) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SshUser) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
