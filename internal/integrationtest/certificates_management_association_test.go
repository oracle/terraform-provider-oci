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
	// ENV VARIABLES that can be configured to query data sources by: name, associationId, certificatesResourceId, or associatedResourceId
	associationName            = utils.GetEnvSettingWithBlankDefault("association_name")
	associationNameVariableStr = fmt.Sprintf("variable \"association_name\" { default = \"%s\" }\n", associationId)

	certificatesResourceId            = utils.GetEnvSettingWithBlankDefault("certificates_resource_ocid")
	certificatesResourceIdVariableStr = fmt.Sprintf("variable \"certificates_resource_ocid\" { default = \"%s\" }\n", associationId)

	associatedResourceId            = utils.GetEnvSettingWithBlankDefault("associated_resource_ocid")
	associatedResourceIdVariableStr = fmt.Sprintf("variable \"associated_resource_ocid\" { default = \"%s\" }\n", associationId)

	associationId            = utils.GetEnvSettingWithBlankDefault("association_ocid")
	associationIdVariableStr = fmt.Sprintf("variable \"association_ocid\" { default = \"%s\" }\n", associationId)

	CertificatesManagementassociationSingularDataSourceRepresentation = map[string]interface{}{
		"association_id": acctest.Representation{RepType: acctest.Required, Create: associationId},
	}

	CertificatesManagementassociationDataSourceRepresentation = map[string]interface{}{
		"association_id":   acctest.Representation{RepType: acctest.Optional, Create: associationId},
		"association_type": acctest.Representation{RepType: acctest.Optional, Create: `CERTIFICATE`},
	}
)

func TestCertificatesManagementAssociationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCertificatesManagementAssociationResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_certificates_management_associations.test_associations"
	singularDatasourceName := "data.oci_certificates_management_association.test_association"

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
					acctest.GenerateDataSourceFromRepresentationMap("oci_certificates_management_associations", "test_associations", acctest.Optional, acctest.Create, CertificatesManagementassociationDataSourceRepresentation) +
					compartmentIdVariableStr,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "association_id"),
					resource.TestCheckResourceAttr(datasourceName, "association_type", "CERTIFICATE"),
					resource.TestCheckResourceAttrSet(datasourceName, "association_collection.#"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_certificates_management_association", "test_association", acctest.Required, acctest.Create, CertificatesManagementassociationSingularDataSourceRepresentation) +
					compartmentIdVariableStr,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "association_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "association_type"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
		},
	})
}
