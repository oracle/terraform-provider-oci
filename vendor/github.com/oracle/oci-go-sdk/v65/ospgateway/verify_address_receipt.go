// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OSP Gateway API
//
// This site describes all the Rest endpoints of OSP Gateway.
//

package ospgateway

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VerifyAddressReceipt Address verficiation result
type VerifyAddressReceipt struct {
	Address *Address `mandatory:"true" json:"address"`

	// Address quality type.
	Quality AddressQualityTypeEnum `mandatory:"true" json:"quality"`

	// Address verification code.
	VerificationCode AddressVerificationCodeEnum `mandatory:"true" json:"verificationCode"`
}

func (m VerifyAddressReceipt) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VerifyAddressReceipt) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAddressQualityTypeEnum(string(m.Quality)); !ok && m.Quality != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Quality: %s. Supported values are: %s.", m.Quality, strings.Join(GetAddressQualityTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAddressVerificationCodeEnum(string(m.VerificationCode)); !ok && m.VerificationCode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VerificationCode: %s. Supported values are: %s.", m.VerificationCode, strings.Join(GetAddressVerificationCodeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
