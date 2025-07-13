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

// issue-routing-tag: datascience/default
func TestDatascienceModelDeploymentModelStateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceModelDeploymentModelStateResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_datascience_model_group_models.test_model_group_models"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_model_group_models", "test_model_group_models", acctest.Required, acctest.Create, DatascienceModelGroupModelDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceModelGroupModelResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "model_group_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "model_group_models.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_group_models.0.category"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_group_models.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_group_models.0.created_by"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_group_models.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_group_models.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_group_models.0.is_model_by_reference"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_group_models.0.model_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_group_models.0.project_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_group_models.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_group_models.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_group_models.0.time_updated"),
			),
		},
	})
}
