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
	// ENV VARIABLES that can be configured to query data sources by: name, associationId, certificatesResourceId, or associatedResourceId
	associationName            = getEnvSettingWithBlankDefault("association_name")
	associationNameVariableStr = fmt.Sprintf("variable \"association_name\" { default = \"%s\" }\n", associationId)

	certificatesResourceId            = getEnvSettingWithBlankDefault("certificates_resource_ocid")
	certificatesResourceIdVariableStr = fmt.Sprintf("variable \"certificates_resource_ocid\" { default = \"%s\" }\n", associationId)

	associatedResourceId            = getEnvSettingWithBlankDefault("associated_resource_ocid")
	associatedResourceIdVariableStr = fmt.Sprintf("variable \"associated_resource_ocid\" { default = \"%s\" }\n", associationId)

	associationId            = getEnvSettingWithBlankDefault("association_ocid")
	associationIdVariableStr = fmt.Sprintf("variable \"association_ocid\" { default = \"%s\" }\n", associationId)

	associationSingularDataSourceRepresentation = map[string]interface{}{
		"association_id": Representation{RepType: Required, Create: associationId},
	}

	associationDataSourceRepresentation = map[string]interface{}{
		"association_id":   Representation{RepType: Optional, Create: associationId},
		"association_type": Representation{RepType: Optional, Create: `CERTIFICATE`},
	}
)

func TestCertificatesManagementAssociationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCertificatesManagementAssociationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_certificates_management_associations.test_associations"
	singularDatasourceName := "data.oci_certificates_management_association.test_association"

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
					GenerateDataSourceFromRepresentationMap("oci_certificates_management_associations", "test_associations", Optional, Create, associationDataSourceRepresentation) +
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
					GenerateDataSourceFromRepresentationMap("oci_certificates_management_association", "test_association", Required, Create, associationSingularDataSourceRepresentation) +
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
