// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// TunnelSecurityAssociationSummary Detailed Tunnel SA
type TunnelSecurityAssociationSummary struct {

	// IP and mask of the Partner Subnet for Policy Based VPNs or Static Routes
	CpeSubnet *string `mandatory:"false" json:"cpeSubnet"`

	// IP and mask of the Local Subnet for Policy Based VPNs or Static Routes
	OracleSubnet *string `mandatory:"false" json:"oracleSubnet"`

	// Phase 1 Status of the Tunnel
	TunnelSaStatus TunnelSecurityAssociationSummaryTunnelSaStatusEnum `mandatory:"false" json:"tunnelSaStatus,omitempty"`

	// Current state if status is not up, including phase1/phase2 and possible reason for tunnel not up
	TunnelSaErrorInfo *string `mandatory:"false" json:"tunnelSaErrorInfo"`

	// Seconds in current state
	Time *string `mandatory:"false" json:"time"`
}

func (m TunnelSecurityAssociationSummary) String() string {
	return common.PointerString(m)
}

// TunnelSecurityAssociationSummaryTunnelSaStatusEnum Enum with underlying type: string
type TunnelSecurityAssociationSummaryTunnelSaStatusEnum string

// Set of constants representing the allowable values for TunnelSecurityAssociationSummaryTunnelSaStatusEnum
const (
	TunnelSecurityAssociationSummaryTunnelSaStatusInitiating TunnelSecurityAssociationSummaryTunnelSaStatusEnum = "INITIATING"
	TunnelSecurityAssociationSummaryTunnelSaStatusListening  TunnelSecurityAssociationSummaryTunnelSaStatusEnum = "LISTENING"
	TunnelSecurityAssociationSummaryTunnelSaStatusUp         TunnelSecurityAssociationSummaryTunnelSaStatusEnum = "UP"
	TunnelSecurityAssociationSummaryTunnelSaStatusDown       TunnelSecurityAssociationSummaryTunnelSaStatusEnum = "DOWN"
	TunnelSecurityAssociationSummaryTunnelSaStatusError      TunnelSecurityAssociationSummaryTunnelSaStatusEnum = "ERROR"
	TunnelSecurityAssociationSummaryTunnelSaStatusUnknown    TunnelSecurityAssociationSummaryTunnelSaStatusEnum = "UNKNOWN"
)

var mappingTunnelSecurityAssociationSummaryTunnelSaStatus = map[string]TunnelSecurityAssociationSummaryTunnelSaStatusEnum{
	"INITIATING": TunnelSecurityAssociationSummaryTunnelSaStatusInitiating,
	"LISTENING":  TunnelSecurityAssociationSummaryTunnelSaStatusListening,
	"UP":         TunnelSecurityAssociationSummaryTunnelSaStatusUp,
	"DOWN":       TunnelSecurityAssociationSummaryTunnelSaStatusDown,
	"ERROR":      TunnelSecurityAssociationSummaryTunnelSaStatusError,
	"UNKNOWN":    TunnelSecurityAssociationSummaryTunnelSaStatusUnknown,
}

// GetTunnelSecurityAssociationSummaryTunnelSaStatusEnumValues Enumerates the set of values for TunnelSecurityAssociationSummaryTunnelSaStatusEnum
func GetTunnelSecurityAssociationSummaryTunnelSaStatusEnumValues() []TunnelSecurityAssociationSummaryTunnelSaStatusEnum {
	values := make([]TunnelSecurityAssociationSummaryTunnelSaStatusEnum, 0)
	for _, v := range mappingTunnelSecurityAssociationSummaryTunnelSaStatus {
		values = append(values, v)
	}
	return values
}
