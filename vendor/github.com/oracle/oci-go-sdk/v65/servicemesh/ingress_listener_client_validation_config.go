// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IngressListenerClientValidationConfig Resource representing the TLS configuration used for validating client certificates.
type IngressListenerClientValidationConfig struct {
	TrustedCaBundle CaBundle `mandatory:"false" json:"trustedCaBundle"`

	// A list of alternate names to verify the subject identity in the certificate presented by the client.
	SubjectAlternateNames []string `mandatory:"false" json:"subjectAlternateNames"`
}

func (m IngressListenerClientValidationConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IngressListenerClientValidationConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *IngressListenerClientValidationConfig) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TrustedCaBundle       cabundle `json:"trustedCaBundle"`
		SubjectAlternateNames []string `json:"subjectAlternateNames"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.TrustedCaBundle.UnmarshalPolymorphicJSON(model.TrustedCaBundle.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TrustedCaBundle = nn.(CaBundle)
	} else {
		m.TrustedCaBundle = nil
	}

	m.SubjectAlternateNames = make([]string, len(model.SubjectAlternateNames))
	for i, n := range model.SubjectAlternateNames {
		m.SubjectAlternateNames[i] = n
	}

	return
}
