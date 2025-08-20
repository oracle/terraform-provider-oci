// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// WebLogic Management Service API
//
// WebLogic Management Service is an OCI service that enables a unified view and management of WebLogic domains
// in Oracle Cloud Infrastructure. Features include on-demand patching of WebLogic domains, rollback of the
// last applied patch, discovery and management of WebLogic instances on a compute host.
//

package wlms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateWlsDomainCredentialDetails The WebLogic domain credential.
type UpdateWlsDomainCredentialDetails struct {

	// The strategy for passing new WebLogic credential.
	Strategy WlsDomainCredentialStrategyEnum `mandatory:"false" json:"strategy,omitempty"`

	// The OCID for WebLogic user secret.
	UserSecretId *string `mandatory:"false" json:"userSecretId"`

	// The OCID for WebLogic password secret.
	PasswordSecretId *string `mandatory:"false" json:"passwordSecretId"`
}

func (m UpdateWlsDomainCredentialDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateWlsDomainCredentialDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingWlsDomainCredentialStrategyEnum(string(m.Strategy)); !ok && m.Strategy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Strategy: %s. Supported values are: %s.", m.Strategy, strings.Join(GetWlsDomainCredentialStrategyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
