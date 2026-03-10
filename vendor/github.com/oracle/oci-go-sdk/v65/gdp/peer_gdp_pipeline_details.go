// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Guarded Data Pipelines API
//
// Use Guarded Data Pipelines to facilitate data transfer between different security domains. The service provides physical, network, and logistical isolation between security domains, malware and vulnerability scanning, auditing, and logging, with deep content inspection capabilities.
//

package gdp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PeerGdpPipelineDetails Details about the peering pipeline.
type PeerGdpPipelineDetails struct {

	// OCID of the peering pipeline.
	PeeringGdpPipelineId *string `mandatory:"true" json:"peeringGdpPipelineId"`

	// Unique one-time-token generated to validate the identity of the sender and prevent the peer request from being tampered with.
	PeerValidationNonce *string `mandatory:"false" json:"peerValidationNonce"`
}

func (m PeerGdpPipelineDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PeerGdpPipelineDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
