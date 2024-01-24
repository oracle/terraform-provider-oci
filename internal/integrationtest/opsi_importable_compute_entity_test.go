// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OpsiImportableComputeEntitySingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	OpsiImportableComputeEntityDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	OpsiImportableComputeEntityResourceConfig = ""
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiImportableComputeEntityResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiImportableComputeEntityResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	//Only doing singular test as that is what we use in test tenancy
	//datasourceName := "data.oci_opsi_importable_compute_entities.test_importable_compute_entities"
	singularDatasourceName := "data.oci_opsi_importable_compute_entity.test_importable_compute_entity"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource - Only doing singular test as that is what we use in test tenancy
		//{
		//	Config: config +
		//		acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_importable_compute_entities", "test_importable_compute_entities", acctest.Required, acctest.Create, OpsiImportableComputeEntityDataSourceRepresentation) +

		//		compartmentIdVariableStr + OpsiImportableComputeEntityResourceConfig,
		//	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		//		resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

		//		resource.TestCheckResourceAttr(datasourceName, "items.#", "2"),
		//	),
		//},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_importable_compute_entity", "test_importable_compute_entity", acctest.Required, acctest.Create, OpsiImportableComputeEntitySingularDataSourceRepresentation) +
				compartmentIdVariableStr + OpsiImportableComputeEntityResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
			),
		},
	})
}
