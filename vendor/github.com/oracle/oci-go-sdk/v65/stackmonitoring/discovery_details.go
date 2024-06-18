// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DiscoveryDetails The request of DiscoveryJob Resource details.
type DiscoveryDetails struct {

	// The OCID of Management Agent
	AgentId *string `mandatory:"true" json:"agentId"`

	// Resource Type.
	ResourceType DiscoveryDetailsResourceTypeEnum `mandatory:"true" json:"resourceType"`

	// The Name of resource type
	ResourceName *string `mandatory:"true" json:"resourceName"`

	Properties *PropertyDetails `mandatory:"true" json:"properties"`

	// License edition of the monitored resource.
	License LicenseTypeEnum `mandatory:"false" json:"license,omitempty"`

	Credentials *CredentialCollection `mandatory:"false" json:"credentials"`

	Tags *PropertyDetails `mandatory:"false" json:"tags"`
}

func (m DiscoveryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiscoveryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDiscoveryDetailsResourceTypeEnum(string(m.ResourceType)); !ok && m.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", m.ResourceType, strings.Join(GetDiscoveryDetailsResourceTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLicenseTypeEnum(string(m.License)); !ok && m.License != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for License: %s. Supported values are: %s.", m.License, strings.Join(GetLicenseTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DiscoveryDetailsResourceTypeEnum Enum with underlying type: string
type DiscoveryDetailsResourceTypeEnum string

// Set of constants representing the allowable values for DiscoveryDetailsResourceTypeEnum
const (
	DiscoveryDetailsResourceTypeWeblogicDomain   DiscoveryDetailsResourceTypeEnum = "WEBLOGIC_DOMAIN"
	DiscoveryDetailsResourceTypeEbsInstance      DiscoveryDetailsResourceTypeEnum = "EBS_INSTANCE"
	DiscoveryDetailsResourceTypeSqlServer        DiscoveryDetailsResourceTypeEnum = "SQL_SERVER"
	DiscoveryDetailsResourceTypeApacheTomcat     DiscoveryDetailsResourceTypeEnum = "APACHE_TOMCAT"
	DiscoveryDetailsResourceTypeOracleDatabase   DiscoveryDetailsResourceTypeEnum = "ORACLE_DATABASE"
	DiscoveryDetailsResourceTypeOciOracleDb      DiscoveryDetailsResourceTypeEnum = "OCI_ORACLE_DB"
	DiscoveryDetailsResourceTypeOciOracleCdb     DiscoveryDetailsResourceTypeEnum = "OCI_ORACLE_CDB"
	DiscoveryDetailsResourceTypeOciOraclePdb     DiscoveryDetailsResourceTypeEnum = "OCI_ORACLE_PDB"
	DiscoveryDetailsResourceTypeHost             DiscoveryDetailsResourceTypeEnum = "HOST"
	DiscoveryDetailsResourceTypeOraclePsft       DiscoveryDetailsResourceTypeEnum = "ORACLE_PSFT"
	DiscoveryDetailsResourceTypeOracleMft        DiscoveryDetailsResourceTypeEnum = "ORACLE_MFT"
	DiscoveryDetailsResourceTypeApacheHttpServer DiscoveryDetailsResourceTypeEnum = "APACHE_HTTP_SERVER"
	DiscoveryDetailsResourceTypeOracleGoldengate DiscoveryDetailsResourceTypeEnum = "ORACLE_GOLDENGATE"
	DiscoveryDetailsResourceTypeCustomResource   DiscoveryDetailsResourceTypeEnum = "CUSTOM_RESOURCE"
	DiscoveryDetailsResourceTypeOracleHttpServer DiscoveryDetailsResourceTypeEnum = "ORACLE_HTTP_SERVER"
	DiscoveryDetailsResourceTypeMicrosoftIis     DiscoveryDetailsResourceTypeEnum = "MICROSOFT_IIS"
)

var mappingDiscoveryDetailsResourceTypeEnum = map[string]DiscoveryDetailsResourceTypeEnum{
	"WEBLOGIC_DOMAIN":    DiscoveryDetailsResourceTypeWeblogicDomain,
	"EBS_INSTANCE":       DiscoveryDetailsResourceTypeEbsInstance,
	"SQL_SERVER":         DiscoveryDetailsResourceTypeSqlServer,
	"APACHE_TOMCAT":      DiscoveryDetailsResourceTypeApacheTomcat,
	"ORACLE_DATABASE":    DiscoveryDetailsResourceTypeOracleDatabase,
	"OCI_ORACLE_DB":      DiscoveryDetailsResourceTypeOciOracleDb,
	"OCI_ORACLE_CDB":     DiscoveryDetailsResourceTypeOciOracleCdb,
	"OCI_ORACLE_PDB":     DiscoveryDetailsResourceTypeOciOraclePdb,
	"HOST":               DiscoveryDetailsResourceTypeHost,
	"ORACLE_PSFT":        DiscoveryDetailsResourceTypeOraclePsft,
	"ORACLE_MFT":         DiscoveryDetailsResourceTypeOracleMft,
	"APACHE_HTTP_SERVER": DiscoveryDetailsResourceTypeApacheHttpServer,
	"ORACLE_GOLDENGATE":  DiscoveryDetailsResourceTypeOracleGoldengate,
	"CUSTOM_RESOURCE":    DiscoveryDetailsResourceTypeCustomResource,
	"ORACLE_HTTP_SERVER": DiscoveryDetailsResourceTypeOracleHttpServer,
	"MICROSOFT_IIS":      DiscoveryDetailsResourceTypeMicrosoftIis,
}

var mappingDiscoveryDetailsResourceTypeEnumLowerCase = map[string]DiscoveryDetailsResourceTypeEnum{
	"weblogic_domain":    DiscoveryDetailsResourceTypeWeblogicDomain,
	"ebs_instance":       DiscoveryDetailsResourceTypeEbsInstance,
	"sql_server":         DiscoveryDetailsResourceTypeSqlServer,
	"apache_tomcat":      DiscoveryDetailsResourceTypeApacheTomcat,
	"oracle_database":    DiscoveryDetailsResourceTypeOracleDatabase,
	"oci_oracle_db":      DiscoveryDetailsResourceTypeOciOracleDb,
	"oci_oracle_cdb":     DiscoveryDetailsResourceTypeOciOracleCdb,
	"oci_oracle_pdb":     DiscoveryDetailsResourceTypeOciOraclePdb,
	"host":               DiscoveryDetailsResourceTypeHost,
	"oracle_psft":        DiscoveryDetailsResourceTypeOraclePsft,
	"oracle_mft":         DiscoveryDetailsResourceTypeOracleMft,
	"apache_http_server": DiscoveryDetailsResourceTypeApacheHttpServer,
	"oracle_goldengate":  DiscoveryDetailsResourceTypeOracleGoldengate,
	"custom_resource":    DiscoveryDetailsResourceTypeCustomResource,
	"oracle_http_server": DiscoveryDetailsResourceTypeOracleHttpServer,
	"microsoft_iis":      DiscoveryDetailsResourceTypeMicrosoftIis,
}

// GetDiscoveryDetailsResourceTypeEnumValues Enumerates the set of values for DiscoveryDetailsResourceTypeEnum
func GetDiscoveryDetailsResourceTypeEnumValues() []DiscoveryDetailsResourceTypeEnum {
	values := make([]DiscoveryDetailsResourceTypeEnum, 0)
	for _, v := range mappingDiscoveryDetailsResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveryDetailsResourceTypeEnumStringValues Enumerates the set of values in String for DiscoveryDetailsResourceTypeEnum
func GetDiscoveryDetailsResourceTypeEnumStringValues() []string {
	return []string{
		"WEBLOGIC_DOMAIN",
		"EBS_INSTANCE",
		"SQL_SERVER",
		"APACHE_TOMCAT",
		"ORACLE_DATABASE",
		"OCI_ORACLE_DB",
		"OCI_ORACLE_CDB",
		"OCI_ORACLE_PDB",
		"HOST",
		"ORACLE_PSFT",
		"ORACLE_MFT",
		"APACHE_HTTP_SERVER",
		"ORACLE_GOLDENGATE",
		"CUSTOM_RESOURCE",
		"ORACLE_HTTP_SERVER",
		"MICROSOFT_IIS",
	}
}

// GetMappingDiscoveryDetailsResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveryDetailsResourceTypeEnum(val string) (DiscoveryDetailsResourceTypeEnum, bool) {
	enum, ok := mappingDiscoveryDetailsResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
