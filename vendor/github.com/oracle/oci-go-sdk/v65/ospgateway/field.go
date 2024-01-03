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

// Field Field information
type Field struct {

	// The field name
	Name *string `mandatory:"true" json:"name"`

	// The given field is requeired or not
	IsRequired *bool `mandatory:"true" json:"isRequired"`

	Format *Format `mandatory:"false" json:"format"`

	Label *Label `mandatory:"false" json:"label"`

	// Locale code (rfc4646 format) of a forced language (e.g.: jp addresses require jp always)
	Language *string `mandatory:"false" json:"language"`
}

func (m Field) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Field) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
