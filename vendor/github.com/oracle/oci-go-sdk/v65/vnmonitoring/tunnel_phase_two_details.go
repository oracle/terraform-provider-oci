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

// TunnelPhaseTwoDetails IPsec tunnel detail information specific to phase two.
type TunnelPhaseTwoDetails struct {

	// Indicates whether custom phase two configuration is enabled.
	// If this option is not enabled, default settings are proposed.
	IsCustomPhaseTwoConfig *bool `mandatory:"false" json:"isCustomPhaseTwoConfig"`

	// The total configured lifetime of the IKE security association.
	Lifetime *int64 `mandatory:"false" json:"lifetime"`

	// The remaining lifetime before the key is refreshed.
	RemainingLifetime *int64 `mandatory:"false" json:"remainingLifetime"`

	// Phase two authentication algorithm proposed during tunnel negotiation.
	CustomAuthenticationAlgorithm *string `mandatory:"false" json:"customAuthenticationAlgorithm"`

	// The negotiated phase two authentication algorithm.
	NegotiatedAuthenticationAlgorithm *string `mandatory:"false" json:"negotiatedAuthenticationAlgorithm"`

	// The proposed custom phase two encryption algorithm.
	CustomEncryptionAlgorithm *string `mandatory:"false" json:"customEncryptionAlgorithm"`

	// The negotiated encryption algorithm.
	NegotiatedEncryptionAlgorithm *string `mandatory:"false" json:"negotiatedEncryptionAlgorithm"`

	// The proposed Diffie-Hellman group.
	DhGroup *string `mandatory:"false" json:"dhGroup"`

	// The negotiated Diffie-Hellman group.
	NegotiatedDhGroup *string `mandatory:"false" json:"negotiatedDhGroup"`

	// Indicates that ESP phase two is established.
	IsEspEstablished *bool `mandatory:"false" json:"isEspEstablished"`

	// Indicates that PFS (perfect forward secrecy) is enabled.
	IsPfsEnabled *bool `mandatory:"false" json:"isPfsEnabled"`

	// The date and time the remaining lifetime was last retrieved, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	RemainingLifetimeLastRetrieved *common.SDKTime `mandatory:"false" json:"remainingLifetimeLastRetrieved"`
}

func (m TunnelPhaseTwoDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TunnelPhaseTwoDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
