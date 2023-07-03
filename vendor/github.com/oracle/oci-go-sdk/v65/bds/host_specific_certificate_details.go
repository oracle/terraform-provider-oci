// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HostSpecificCertificateDetails Host specific certificate details
type HostSpecificCertificateDetails struct {

	// Name of the host.
	HostName *string `mandatory:"false" json:"hostName"`

	// Type of certificate self signed or CA signed
	CertificateType HostSpecificCertificateDetailsCertificateTypeEnum `mandatory:"false" json:"certificateType,omitempty"`

	// The time the certificate expires, shown as an RFC 3339 formatted datetime string.
	TimeExpiry *common.SDKTime `mandatory:"false" json:"timeExpiry"`
}

func (m HostSpecificCertificateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostSpecificCertificateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingHostSpecificCertificateDetailsCertificateTypeEnum(string(m.CertificateType)); !ok && m.CertificateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CertificateType: %s. Supported values are: %s.", m.CertificateType, strings.Join(GetHostSpecificCertificateDetailsCertificateTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HostSpecificCertificateDetailsCertificateTypeEnum Enum with underlying type: string
type HostSpecificCertificateDetailsCertificateTypeEnum string

// Set of constants representing the allowable values for HostSpecificCertificateDetailsCertificateTypeEnum
const (
	HostSpecificCertificateDetailsCertificateTypeCustomSigned HostSpecificCertificateDetailsCertificateTypeEnum = "CUSTOM_SIGNED"
	HostSpecificCertificateDetailsCertificateTypeSelfSigned   HostSpecificCertificateDetailsCertificateTypeEnum = "SELF_SIGNED"
)

var mappingHostSpecificCertificateDetailsCertificateTypeEnum = map[string]HostSpecificCertificateDetailsCertificateTypeEnum{
	"CUSTOM_SIGNED": HostSpecificCertificateDetailsCertificateTypeCustomSigned,
	"SELF_SIGNED":   HostSpecificCertificateDetailsCertificateTypeSelfSigned,
}

var mappingHostSpecificCertificateDetailsCertificateTypeEnumLowerCase = map[string]HostSpecificCertificateDetailsCertificateTypeEnum{
	"custom_signed": HostSpecificCertificateDetailsCertificateTypeCustomSigned,
	"self_signed":   HostSpecificCertificateDetailsCertificateTypeSelfSigned,
}

// GetHostSpecificCertificateDetailsCertificateTypeEnumValues Enumerates the set of values for HostSpecificCertificateDetailsCertificateTypeEnum
func GetHostSpecificCertificateDetailsCertificateTypeEnumValues() []HostSpecificCertificateDetailsCertificateTypeEnum {
	values := make([]HostSpecificCertificateDetailsCertificateTypeEnum, 0)
	for _, v := range mappingHostSpecificCertificateDetailsCertificateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetHostSpecificCertificateDetailsCertificateTypeEnumStringValues Enumerates the set of values in String for HostSpecificCertificateDetailsCertificateTypeEnum
func GetHostSpecificCertificateDetailsCertificateTypeEnumStringValues() []string {
	return []string{
		"CUSTOM_SIGNED",
		"SELF_SIGNED",
	}
}

// GetMappingHostSpecificCertificateDetailsCertificateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHostSpecificCertificateDetailsCertificateTypeEnum(val string) (HostSpecificCertificateDetailsCertificateTypeEnum, bool) {
	enum, ok := mappingHostSpecificCertificateDetailsCertificateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
