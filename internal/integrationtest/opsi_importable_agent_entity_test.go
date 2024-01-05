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
	OpsiImportableAgentEntitySingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	OpsiImportableAgentEntityDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	OpsiImportableAgentEntityResourceConfig = ""
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiImportableAgentEntityResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiImportableAgentEntityResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_opsi_importable_agent_entities.test_importable_agent_entities"
	singularDatasourceName := "data.oci_opsi_importable_agent_entity.test_importable_agent_entity"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_importable_agent_entities", "test_importable_agent_entities", acctest.Required, acctest.Create, OpsiImportableAgentEntityDataSourceRepresentation) +
				compartmentIdVariableStr + OpsiImportableAgentEntityResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_importable_agent_entity", "test_importable_agent_entity", acctest.Required, acctest.Create, OpsiImportableAgentEntitySingularDataSourceRepresentation) +
				compartmentIdVariableStr + OpsiImportableAgentEntityResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
			),
		},
	})
}
