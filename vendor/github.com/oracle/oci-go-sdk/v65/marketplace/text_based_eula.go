// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TextBasedEula An end user license agreement that is provided as text.
type TextBasedEula struct {

	// The text of the end user license agreement.
	LicenseText *string `mandatory:"false" json:"licenseText"`
}

func (m TextBasedEula) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TextBasedEula) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TextBasedEula) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTextBasedEula TextBasedEula
	s := struct {
		DiscriminatorParam string `json:"eulaType"`
		MarshalTypeTextBasedEula
	}{
		"TEXT",
		(MarshalTypeTextBasedEula)(m),
	}

	return json.Marshal(&s)
}
