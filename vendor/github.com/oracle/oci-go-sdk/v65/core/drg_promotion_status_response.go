// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrgPromotionStatusResponse The promotion/unpromotion status of a DRG
type DrgPromotionStatusResponse struct {

	// OCID of the DRG
	DrgId *string `mandatory:"true" json:"drgId"`

	// The promotion status of the DRG
	DrgPromotionStatus DrgPromotionStatusResponseDrgPromotionStatusEnum `mandatory:"false" json:"drgPromotionStatus,omitempty"`

	// A map of the promotion status of each RPC connection on this DRG {conn_id -> promo_status}
	RpcPromotionStatus map[string]string `mandatory:"false" json:"rpcPromotionStatus"`

	// A map of the promotion status of each VC on this DRG {conn_id -> promo_status}
	VcPromotionStatus map[string]string `mandatory:"false" json:"vcPromotionStatus"`

	// A map of the promotion status of each IPSec connection on this DRG {conn_id -> promo_status}
	IpsecPromotionStatus map[string]string `mandatory:"false" json:"ipsecPromotionStatus"`
}

func (m DrgPromotionStatusResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrgPromotionStatusResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDrgPromotionStatusResponseDrgPromotionStatusEnum(string(m.DrgPromotionStatus)); !ok && m.DrgPromotionStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DrgPromotionStatus: %s. Supported values are: %s.", m.DrgPromotionStatus, strings.Join(GetDrgPromotionStatusResponseDrgPromotionStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DrgPromotionStatusResponseDrgPromotionStatusEnum Enum with underlying type: string
type DrgPromotionStatusResponseDrgPromotionStatusEnum string

// Set of constants representing the allowable values for DrgPromotionStatusResponseDrgPromotionStatusEnum
const (
	DrgPromotionStatusResponseDrgPromotionStatusUnpromoted  DrgPromotionStatusResponseDrgPromotionStatusEnum = "UNPROMOTED"
	DrgPromotionStatusResponseDrgPromotionStatusPromoting   DrgPromotionStatusResponseDrgPromotionStatusEnum = "PROMOTING"
	DrgPromotionStatusResponseDrgPromotionStatusPromoted    DrgPromotionStatusResponseDrgPromotionStatusEnum = "PROMOTED"
	DrgPromotionStatusResponseDrgPromotionStatusUnpromoting DrgPromotionStatusResponseDrgPromotionStatusEnum = "UNPROMOTING"
)

var mappingDrgPromotionStatusResponseDrgPromotionStatusEnum = map[string]DrgPromotionStatusResponseDrgPromotionStatusEnum{
	"UNPROMOTED":  DrgPromotionStatusResponseDrgPromotionStatusUnpromoted,
	"PROMOTING":   DrgPromotionStatusResponseDrgPromotionStatusPromoting,
	"PROMOTED":    DrgPromotionStatusResponseDrgPromotionStatusPromoted,
	"UNPROMOTING": DrgPromotionStatusResponseDrgPromotionStatusUnpromoting,
}

var mappingDrgPromotionStatusResponseDrgPromotionStatusEnumLowerCase = map[string]DrgPromotionStatusResponseDrgPromotionStatusEnum{
	"unpromoted":  DrgPromotionStatusResponseDrgPromotionStatusUnpromoted,
	"promoting":   DrgPromotionStatusResponseDrgPromotionStatusPromoting,
	"promoted":    DrgPromotionStatusResponseDrgPromotionStatusPromoted,
	"unpromoting": DrgPromotionStatusResponseDrgPromotionStatusUnpromoting,
}

// GetDrgPromotionStatusResponseDrgPromotionStatusEnumValues Enumerates the set of values for DrgPromotionStatusResponseDrgPromotionStatusEnum
func GetDrgPromotionStatusResponseDrgPromotionStatusEnumValues() []DrgPromotionStatusResponseDrgPromotionStatusEnum {
	values := make([]DrgPromotionStatusResponseDrgPromotionStatusEnum, 0)
	for _, v := range mappingDrgPromotionStatusResponseDrgPromotionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDrgPromotionStatusResponseDrgPromotionStatusEnumStringValues Enumerates the set of values in String for DrgPromotionStatusResponseDrgPromotionStatusEnum
func GetDrgPromotionStatusResponseDrgPromotionStatusEnumStringValues() []string {
	return []string{
		"UNPROMOTED",
		"PROMOTING",
		"PROMOTED",
		"UNPROMOTING",
	}
}

// GetMappingDrgPromotionStatusResponseDrgPromotionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrgPromotionStatusResponseDrgPromotionStatusEnum(val string) (DrgPromotionStatusResponseDrgPromotionStatusEnum, bool) {
	enum, ok := mappingDrgPromotionStatusResponseDrgPromotionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
