// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	certNameForCertificateVersionTests                 = "test-certificate-version-cert-" + RandomString(10, charsetWithoutDigits)
	certificateVersionSingularDataSourceRepresentation = map[string]interface{}{
		"certificate_id":             Representation{RepType: Required, Create: `${oci_certificates_management_certificate.test_certificate.id}`},
		"certificate_version_number": Representation{RepType: Required, Create: `1`},
	}

	certificateVersionDataSourceRepresentation = map[string]interface{}{
		"certificate_id": Representation{RepType: Required, Create: `${oci_certificates_management_certificate.test_certificate.id}`},
		"version_number": Representation{RepType: Optional, Create: `1`},
	}

	CertificateVersionResourceConfig = GenerateResourceFromRepresentationMap("oci_certificates_management_certificate", "test_certificate", Required, Create,
		RepresentationCopyWithNewProperties(certificatesManagementCertificateRepresentation, map[string]interface{}{
			"name": Representation{RepType: Required, Create: certNameForCertificateVersionTests},
		}))
)

func TestCertificatesManagementCertificateVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCertificatesManagementCertificateVersionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_certificates_management_certificate_versions.test_certificate_versions"
	singularDatasourceName := "data.oci_certificates_management_certificate_version.test_certificate_version"

	SaveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					GenerateDataSourceFromRepresentationMap("oci_certificates_management_certificate_versions", "test_certificate_versions", Optional, Create, certificateVersionDataSourceRepresentation) +
					compartmentIdVariableStr + CertificateVersionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "certificate_id"),
					resource.TestCheckResourceAttr(datasourceName, "version_number", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "certificate_version_collection.#"),
					resource.TestCheckResourceAttr(datasourceName, "certificate_version_collection.0.items.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "certificate_version_collection.0.items.0.version_number", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					GenerateDataSourceFromRepresentationMap("oci_certificates_management_certificate_version", "test_certificate_version", Required, Create, certificateVersionSingularDataSourceRepresentation) +
					compartmentIdVariableStr + CertificateVersionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "certificate_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "certificate_version_number", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "issuer_ca_version_number"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "serial_number"),
					resource.TestCheckResourceAttr(singularDatasourceName, "stages.#", "2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttr(singularDatasourceName, "validity.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "version_number"),
				),
			},
		},
	})
}
