// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	//metadataid                                                  = `ocid1.pactlapimetadata.oc1.iad.` + utils.RandomString(60, utils.CharsetWithoutDigits)
	ApiaccesscontrolApiMetadataSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"api_metadata_id": acctest.Representation{RepType: acctest.Required, Create: `${var.apimetadata_id}`},
	}

	ApiaccesscontrolApiMetadataDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"resource_type":  acctest.Representation{RepType: acctest.Optional, Create: `CLOUDEXADATAINFRASTRUCTURE`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}

	ApiaccesscontrolApiMetadataResourceConfig = ""
)

// issue-routing-tag: apiaccesscontrol/default
func TestApiaccesscontrolApiMetadataResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApiaccesscontrolApiMetadataResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	//apimetadataId := utils.GetEnvSettingWithBlankDefault("apimetadata_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	//apimetadataIdVariableStr := fmt.Sprintf("variable \"apimetadata_id\" { default = \"%s\" }\n", apimetadataId)

	datasourceName := "data.oci_apiaccesscontrol_api_metadatas.test_api_metadatas"
	//singularDatasourceName := "data.oci_apiaccesscontrol_api_metadata.test_api_metadata"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apiaccesscontrol_api_metadatas", "test_api_metadatas", acctest.Optional, acctest.Create, ApiaccesscontrolApiMetadataDataSourceRepresentation) +
				compartmentIdVariableStr + ApiaccesscontrolApiMetadataResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "resource_type", "CLOUDEXADATAINFRASTRUCTURE"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttrSet(datasourceName, "api_metadata_collection.#"),
			),
		},
		//This is not required yet to be removed from spec.
		// verify singular datasource
		//{
		//Config: config +
		//	acctest.GenerateDataSourceFromRepresentationMap("oci_apiaccesscontrol_api_metadata", "test_api_metadata", acctest.Required, acctest.Create, ApiaccesscontrolApiMetadataSingularDataSourceRepresentation) +
		//	apimetadataIdVariableStr + ApiaccesscontrolApiMetadataResourceConfig,
		//Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		//	resource.TestCheckResourceAttrSet(singularDatasourceName, "api_metadata_id"),
		//
		//	resource.TestCheckResourceAttrSet(singularDatasourceName, "api_name"),
		//	resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
		//	resource.TestCheckResourceAttrSet(singularDatasourceName, "entity_type"),
		//	resource.TestCheckResourceAttr(singularDatasourceName, "fields.#", "1"),
		//	resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
		//	resource.TestCheckResourceAttrSet(singularDatasourceName, "path"),
		//	resource.TestCheckResourceAttrSet(singularDatasourceName, "service_name"),
		//	resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
		//	resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
		//	resource.TestCheckResourceAttrSet(singularDatasourceName, "time_deleted"),
		//	resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
		//),
		//},
	})
}
