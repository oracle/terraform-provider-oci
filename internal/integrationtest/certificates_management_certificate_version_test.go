// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	certNameForCertificateVersionTests                                       = "test-certificate-version-cert-" + utils.RandomString(10, utils.CharsetWithoutDigits)
	CertificatesManagementcertificateVersionSingularDataSourceRepresentation = map[string]interface{}{
		"certificate_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_certificates_management_certificate.test_certificate.id}`},
		"certificate_version_number": acctest.Representation{RepType: acctest.Required, Create: `1`},
	}

	CertificatesManagementcertificateVersionDataSourceRepresentation = map[string]interface{}{
		"certificate_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_certificates_management_certificate.test_certificate.id}`},
		"version_number": acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	CertificatesManagementCertificateVersionResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate", "test_certificate", acctest.Required, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(certificatesManagementCertificateRepresentation, map[string]interface{}{
			"name": acctest.Representation{RepType: acctest.Required, Create: certNameForCertificateVersionTests},
		}))
)

func TestCertificatesManagementCertificateVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCertificatesManagementCertificateVersionResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_certificates_management_certificate_versions.test_certificate_versions"
	singularDatasourceName := "data.oci_certificates_management_certificate_version.test_certificate_version"

	acctest.SaveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_certificates_management_certificate_versions", "test_certificate_versions", acctest.Optional, acctest.Create, CertificatesManagementcertificateVersionDataSourceRepresentation) +
					compartmentIdVariableStr + CertificatesManagementCertificateVersionResourceConfig,
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
					acctest.GenerateDataSourceFromRepresentationMap("oci_certificates_management_certificate_version", "test_certificate_version", acctest.Required, acctest.Create, CertificatesManagementcertificateVersionSingularDataSourceRepresentation) +
					compartmentIdVariableStr + CertificatesManagementCertificateVersionResourceConfig,
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
