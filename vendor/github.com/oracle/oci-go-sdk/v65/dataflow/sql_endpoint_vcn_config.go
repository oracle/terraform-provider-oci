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

// SqlEndpointVcnConfig The VCN configuration for VCN network type selection.
type SqlEndpointVcnConfig struct {

	// The VCN OCID.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The VCN Subnet OCID.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The host name prefix.
	HostNamePrefix *string `mandatory:"false" json:"hostNamePrefix"`

	// The OCIDs of Network Security Groups (NSGs).
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Ip Address of private endpoint
	PrivateEndpointIp *string `mandatory:"false" json:"privateEndpointIp"`
}

func (m SqlEndpointVcnConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlEndpointVcnConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SqlEndpointVcnConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSqlEndpointVcnConfig SqlEndpointVcnConfig
	s := struct {
		DiscriminatorParam string `json:"networkType"`
		MarshalTypeSqlEndpointVcnConfig
	}{
		"VCN",
		(MarshalTypeSqlEndpointVcnConfig)(m),
	}

	return json.Marshal(&s)
}
