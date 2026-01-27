// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ManagementApplianceConnection Management appliance connection.
type ManagementApplianceConnection struct {

	// Type of connection.
	Type ManagementApplianceConnectionTypeEnum `mandatory:"true" json:"type"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of secret in OCI vault, that is used for storage of username and password in JSON format.
	CredentialsSecretId *string `mandatory:"true" json:"credentialsSecretId"`
}

func (m ManagementApplianceConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagementApplianceConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingManagementApplianceConnectionTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetManagementApplianceConnectionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
