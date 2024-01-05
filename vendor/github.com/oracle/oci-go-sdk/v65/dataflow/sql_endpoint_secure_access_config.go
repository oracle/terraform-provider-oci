// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SqlEndpointSecureAccessConfig Access control rules for secure access selection.
type SqlEndpointSecureAccessConfig struct {

	// A list of SecureAccessControlRule's to which access is limited to
	AccessControlRules []SecureAccessControlRule `mandatory:"false" json:"accessControlRules"`

	// Ip Address of public endpoint
	PublicEndpointIp *string `mandatory:"false" json:"publicEndpointIp"`
}

func (m SqlEndpointSecureAccessConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlEndpointSecureAccessConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SqlEndpointSecureAccessConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSqlEndpointSecureAccessConfig SqlEndpointSecureAccessConfig
	s := struct {
		DiscriminatorParam string `json:"networkType"`
		MarshalTypeSqlEndpointSecureAccessConfig
	}{
		"SECURE_ACCESS",
		(MarshalTypeSqlEndpointSecureAccessConfig)(m),
	}

	return json.Marshal(&s)
}
