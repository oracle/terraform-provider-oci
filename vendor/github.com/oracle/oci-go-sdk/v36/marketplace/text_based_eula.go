// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v36/common"
)

// TextBasedEula An EULA that is provided as text
type TextBasedEula struct {

	// text of the eula
	LicenseText *string `mandatory:"false" json:"licenseText"`
}

func (m TextBasedEula) String() string {
	return common.PointerString(m)
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
