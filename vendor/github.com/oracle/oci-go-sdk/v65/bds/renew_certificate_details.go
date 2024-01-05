// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// RenewCertificateDetails The request body info about renew certificate service list.
type RenewCertificateDetails struct {

	// Base-64 encoded password for the cluster admin user.
	ClusterAdminPassword *string `mandatory:"true" json:"clusterAdminPassword"`

	// List of services for which certificate needs to be renewed. If no services provided renew will happen only for default services - AMBARI,RANGER,HUE,LIVY.
	Services []ServiceEnum `mandatory:"false" json:"services"`

	// Plain text certificate/s in order, separated by new line character. If not provided in request a self-signed root certificate is generated inside the cluster. In case hostCertDetails is provided, root certificate is mandatory.
	RootCertificate *string `mandatory:"false" json:"rootCertificate"`

	// List of leaf certificates to use for services on each host. If custom host certificate is provided the root certificate becomes required.
	HostCertDetails []HostCertDetails `mandatory:"false" json:"hostCertDetails"`

	// Base-64 encoded password for CA certificate's private key. This value can be empty.
	ServerKeyPassword *string `mandatory:"false" json:"serverKeyPassword"`
}

func (m RenewCertificateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RenewCertificateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
