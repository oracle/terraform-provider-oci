// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// OnPremiseConnector The details required to establish a connection to the database using an on-premises connector.
type OnPremiseConnector struct {

	// The OCID of the on-premises connector.
	OnPremConnectorId *string `mandatory:"false" json:"onPremConnectorId"`
}

func (m OnPremiseConnector) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OnPremiseConnector) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OnPremiseConnector) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOnPremiseConnector OnPremiseConnector
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeOnPremiseConnector
	}{
		"ONPREM_CONNECTOR",
		(MarshalTypeOnPremiseConnector)(m),
	}

	return json.Marshal(&s)
}
