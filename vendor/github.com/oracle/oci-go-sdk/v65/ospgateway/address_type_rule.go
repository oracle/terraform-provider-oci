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

// AddressTypeRule Address type rule information
type AddressTypeRule struct {

	// Address type rule fields
	Fields []Field `mandatory:"true" json:"fields"`

	// Third party validation.
	ThirdPartyValidation ThirdPartyValidationTypeEnum `mandatory:"false" json:"thirdPartyValidation,omitempty"`
}

func (m AddressTypeRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddressTypeRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingThirdPartyValidationTypeEnum(string(m.ThirdPartyValidation)); !ok && m.ThirdPartyValidation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ThirdPartyValidation: %s. Supported values are: %s.", m.ThirdPartyValidation, strings.Join(GetThirdPartyValidationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
