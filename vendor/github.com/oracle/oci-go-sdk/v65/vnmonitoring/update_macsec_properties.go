// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateMacsecProperties Properties used to update MACsec settings.
type UpdateMacsecProperties struct {

	// Indicates whether or not MACsec is enabled.
	State MacsecStateEnum `mandatory:"true" json:"state"`

	PrimaryKey *UpdateMacsecKey `mandatory:"false" json:"primaryKey"`

	// Type of encryption cipher suite to use for the MACsec connection.
	EncryptionCipher MacsecEncryptionCipherEnum `mandatory:"false" json:"encryptionCipher,omitempty"`
}

func (m UpdateMacsecProperties) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateMacsecProperties) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMacsecStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetMacsecStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingMacsecEncryptionCipherEnum(string(m.EncryptionCipher)); !ok && m.EncryptionCipher != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EncryptionCipher: %s. Supported values are: %s.", m.EncryptionCipher, strings.Join(GetMacsecEncryptionCipherEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
