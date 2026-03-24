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

// RenewBdsCertificateDetails Request body for renewing certificates.
type RenewBdsCertificateDetails struct {
	ManageCertificateLevelTypeDetails ManageBdsCertificateLevelTypeDetails `mandatory:"true" json:"manageCertificateLevelTypeDetails"`

	// Base-64 encoded password for the cluster admin user.
	ClusterAdminPassword *string `mandatory:"false" json:"clusterAdminPassword"`

	// The secretId for the clusterAdminPassword.
	SecretId *string `mandatory:"false" json:"secretId"`
}

func (m RenewBdsCertificateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RenewBdsCertificateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *RenewBdsCertificateDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ClusterAdminPassword              *string                              `json:"clusterAdminPassword"`
		SecretId                          *string                              `json:"secretId"`
		ManageCertificateLevelTypeDetails managebdscertificateleveltypedetails `json:"manageCertificateLevelTypeDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ClusterAdminPassword = model.ClusterAdminPassword

	m.SecretId = model.SecretId

	nn, e = model.ManageCertificateLevelTypeDetails.UnmarshalPolymorphicJSON(model.ManageCertificateLevelTypeDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ManageCertificateLevelTypeDetails = nn.(ManageBdsCertificateLevelTypeDetails)
	} else {
		m.ManageCertificateLevelTypeDetails = nil
	}

	return
}
