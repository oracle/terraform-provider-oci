// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateFusionEnvironmentAdminUserDetails The credentials for the Fusion Applications service administrator.
type CreateFusionEnvironmentAdminUserDetails struct {

	// The username for the administrator.
	Username *string `mandatory:"true" json:"username"`

	// The email address for the administrator.
	EmailAddress *string `mandatory:"true" json:"emailAddress"`

	// The administrator's first name.
	FirstName *string `mandatory:"true" json:"firstName"`

	// The administrator's last name.
	LastName *string `mandatory:"true" json:"lastName"`

	// The password for the administrator.
	Password *string `mandatory:"false" json:"password"`
}

func (m CreateFusionEnvironmentAdminUserDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateFusionEnvironmentAdminUserDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
