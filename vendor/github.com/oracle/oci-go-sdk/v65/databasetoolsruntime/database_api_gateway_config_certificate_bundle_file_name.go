// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseApiGatewayConfigCertificateBundleFileName The certificate bundle that describes the SSL certicicate files.
type DatabaseApiGatewayConfigCertificateBundleFileName struct {
	CertificatePublic *DatabaseApiGatewayConfigCertificatePublicFileName `mandatory:"false" json:"certificatePublic"`

	CertificatePrivateKey *DatabaseApiGatewayConfigCertificatePrivateKeyFileName `mandatory:"false" json:"certificatePrivateKey"`
}

func (m DatabaseApiGatewayConfigCertificateBundleFileName) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseApiGatewayConfigCertificateBundleFileName) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DatabaseApiGatewayConfigCertificateBundleFileName) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseApiGatewayConfigCertificateBundleFileName DatabaseApiGatewayConfigCertificateBundleFileName
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDatabaseApiGatewayConfigCertificateBundleFileName
	}{
		"FILENAME",
		(MarshalTypeDatabaseApiGatewayConfigCertificateBundleFileName)(m),
	}

	return json.Marshal(&s)
}
