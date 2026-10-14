// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SecurityUserDetails The username and password source details for the security master user.
type SecurityUserDetails struct {

	// The name of the master user that is used to manage security config.
	Username *string `mandatory:"true" json:"username"`

	PasswordDetails SecurityUserPasswordDetails `mandatory:"true" json:"passwordDetails"`
}

func (m SecurityUserDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecurityUserDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *SecurityUserDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Username        *string                     `json:"username"`
		PasswordDetails securityuserpassworddetails `json:"passwordDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Username = model.Username

	nn, e = model.PasswordDetails.UnmarshalPolymorphicJSON(model.PasswordDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PasswordDetails = nn.(SecurityUserPasswordDetails)
	} else {
		m.PasswordDetails = nil
	}

	return
}
