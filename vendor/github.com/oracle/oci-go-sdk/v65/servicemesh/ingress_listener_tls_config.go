// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// IngressListenerTlsConfig TLS enforcement config for the ingress listener.
type IngressListenerTlsConfig struct {

	// DISABLED: Connection can only be plaintext.
	// PERMISSIVE: Connection can be either plaintext or TLS/mTLS. If the clientValidation.trustedCaBundle property is configured for the listener, mTLS is performed and the client's certificates are validated by the gateway.
	// TLS: Connection can only be TLS.
	// MUTUAL_TLS: Connection can only be MTLS.
	Mode IngressListenerTlsConfigModeEnum `mandatory:"true" json:"mode"`

	ServerCertificate TlsCertificate `mandatory:"false" json:"serverCertificate"`

	ClientValidation *IngressListenerClientValidationConfig `mandatory:"false" json:"clientValidation"`
}

func (m IngressListenerTlsConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IngressListenerTlsConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIngressListenerTlsConfigModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetIngressListenerTlsConfigModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *IngressListenerTlsConfig) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ServerCertificate tlscertificate                         `json:"serverCertificate"`
		ClientValidation  *IngressListenerClientValidationConfig `json:"clientValidation"`
		Mode              IngressListenerTlsConfigModeEnum       `json:"mode"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.ServerCertificate.UnmarshalPolymorphicJSON(model.ServerCertificate.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ServerCertificate = nn.(TlsCertificate)
	} else {
		m.ServerCertificate = nil
	}

	m.ClientValidation = model.ClientValidation

	m.Mode = model.Mode

	return
}

// IngressListenerTlsConfigModeEnum Enum with underlying type: string
type IngressListenerTlsConfigModeEnum string

// Set of constants representing the allowable values for IngressListenerTlsConfigModeEnum
const (
	IngressListenerTlsConfigModeDisabled   IngressListenerTlsConfigModeEnum = "DISABLED"
	IngressListenerTlsConfigModePermissive IngressListenerTlsConfigModeEnum = "PERMISSIVE"
	IngressListenerTlsConfigModeTls        IngressListenerTlsConfigModeEnum = "TLS"
	IngressListenerTlsConfigModeMutualTls  IngressListenerTlsConfigModeEnum = "MUTUAL_TLS"
)

var mappingIngressListenerTlsConfigModeEnum = map[string]IngressListenerTlsConfigModeEnum{
	"DISABLED":   IngressListenerTlsConfigModeDisabled,
	"PERMISSIVE": IngressListenerTlsConfigModePermissive,
	"TLS":        IngressListenerTlsConfigModeTls,
	"MUTUAL_TLS": IngressListenerTlsConfigModeMutualTls,
}

var mappingIngressListenerTlsConfigModeEnumLowerCase = map[string]IngressListenerTlsConfigModeEnum{
	"disabled":   IngressListenerTlsConfigModeDisabled,
	"permissive": IngressListenerTlsConfigModePermissive,
	"tls":        IngressListenerTlsConfigModeTls,
	"mutual_tls": IngressListenerTlsConfigModeMutualTls,
}

// GetIngressListenerTlsConfigModeEnumValues Enumerates the set of values for IngressListenerTlsConfigModeEnum
func GetIngressListenerTlsConfigModeEnumValues() []IngressListenerTlsConfigModeEnum {
	values := make([]IngressListenerTlsConfigModeEnum, 0)
	for _, v := range mappingIngressListenerTlsConfigModeEnum {
		values = append(values, v)
	}
	return values
}

// GetIngressListenerTlsConfigModeEnumStringValues Enumerates the set of values in String for IngressListenerTlsConfigModeEnum
func GetIngressListenerTlsConfigModeEnumStringValues() []string {
	return []string{
		"DISABLED",
		"PERMISSIVE",
		"TLS",
		"MUTUAL_TLS",
	}
}

// GetMappingIngressListenerTlsConfigModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIngressListenerTlsConfigModeEnum(val string) (IngressListenerTlsConfigModeEnum, bool) {
	enum, ok := mappingIngressListenerTlsConfigModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
