// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConfigLevelManageBdsCertificateDetails Details of certificate configuration / certificate authority level used to trigger the BDS certificate generation or renewal.
type ConfigLevelManageBdsCertificateDetails struct {

	// The id of the BDS certificate configuration used to generate or renew BDS certificate(s).
	CertificateConfigurationId *string `mandatory:"true" json:"certificateConfigurationId"`

	// Boolean flag specifying whether the request will only generate certificates for nodes which do not have the same certificate authority as the certificate configuration or not. The flag is only used for generating certificates from CONFIG_LEVEL.
	IsMissingNodesOnly *bool `mandatory:"false" json:"isMissingNodesOnly"`
}

func (m ConfigLevelManageBdsCertificateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConfigLevelManageBdsCertificateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ConfigLevelManageBdsCertificateDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeConfigLevelManageBdsCertificateDetails ConfigLevelManageBdsCertificateDetails
	s := struct {
		DiscriminatorParam string `json:"triggerType"`
		MarshalTypeConfigLevelManageBdsCertificateDetails
	}{
		"CONFIG_LEVEL",
		(MarshalTypeConfigLevelManageBdsCertificateDetails)(m),
	}

	return json.Marshal(&s)
}
