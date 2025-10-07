// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CertificateConfig Custom certificate config for customer provided certs.
type CertificateConfig struct {

	// Specifies whether the certificate to be used in cluster is managed by OpenSearch or OCI Certificates service.
	ClusterCertificateMode CertificateModeEnum `mandatory:"false" json:"clusterCertificateMode,omitempty"`

	// Specifies whether the certificate to be used in dashboard is managed by OpenSearch or OCI Certificates service.
	DashboardCertificateMode CertificateModeEnum `mandatory:"false" json:"dashboardCertificateMode,omitempty"`

	// certificate to be used for OpenSearch cluster api communication
	OpenSearchApiCertificateId *string `mandatory:"false" json:"openSearchApiCertificateId"`

	// certificate to be used for OpenSearch dashboard api communication
	OpenSearchDashboardCertificateId *string `mandatory:"false" json:"openSearchDashboardCertificateId"`
}

func (m CertificateConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CertificateConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCertificateModeEnum(string(m.ClusterCertificateMode)); !ok && m.ClusterCertificateMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClusterCertificateMode: %s. Supported values are: %s.", m.ClusterCertificateMode, strings.Join(GetCertificateModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCertificateModeEnum(string(m.DashboardCertificateMode)); !ok && m.DashboardCertificateMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DashboardCertificateMode: %s. Supported values are: %s.", m.DashboardCertificateMode, strings.Join(GetCertificateModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
