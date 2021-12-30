// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	caNameForCaVersionTests = "test-ca-version-ca-" + utils.RandomString(10, utils.CharsetWithoutDigits)

	certificateAuthorityVersionSingularDataSourceRepresentation = map[string]interface{}{
		"certificate_authority_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_certificates_management_certificate_authority.test_certificate_authority.id}`},
		"certificate_authority_version_number": acctest.Representation{RepType: acctest.Required, Create: `1`},
	}

	certificateAuthorityVersionDataSourceRepresentation = map[string]interface{}{
		"certificate_authority_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_certificates_management_certificate_authority.test_certificate_authority.id}`},
		"version_number":           acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	CertificateAuthorityVersionResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_certificate_authority", acctest.Required, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(certificateAuthorityRepresentation, map[string]interface{}{
			"name": acctest.Representation{RepType: acctest.Required, Create: caNameForCaVersionTests},
		}))
)

func TestCertificatesManagementCertificateAuthorityVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCertificatesManagementCertificateAuthorityVersionResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_certificates_management_certificate_authority_versions.test_certificate_authority_versions"
	singularDatasourceName := "data.oci_certificates_management_certificate_authority_version.test_certificate_authority_version"

	acctest.SaveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_certificates_management_certificate_authority_versions", "test_certificate_authority_versions", acctest.Optional, acctest.Create, certificateAuthorityVersionDataSourceRepresentation) +
					compartmentIdVariableStr + CertificateAuthorityVersionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "certificate_authority_id"),
					resource.TestCheckResourceAttr(datasourceName, "version_number", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "certificate_authority_version_collection.#"),
					resource.TestCheckResourceAttr(datasourceName, "certificate_authority_version_collection.0.items.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "certificate_authority_version_collection.0.items.0.version_number", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_certificates_management_certificate_authority_version", "test_certificate_authority_version", acctest.Required, acctest.Create, certificateAuthorityVersionSingularDataSourceRepresentation) +
					compartmentIdVariableStr + CertificateAuthorityVersionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "certificate_authority_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "version_number", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "serial_number"),

					resource.TestCheckResourceAttr(singularDatasourceName, "stages.#", "2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "validity.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "validity.0.time_of_validity_not_after"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "validity.0.time_of_validity_not_before"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
		},
	})
}
