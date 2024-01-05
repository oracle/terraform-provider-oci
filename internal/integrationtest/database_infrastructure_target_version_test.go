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
	infrastructureTargetVersionSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"target_resource_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.target_resource_id}`},
		"target_resource_type": acctest.Representation{RepType: acctest.Required, Create: `EXACC_INFRASTRUCTURE`},
	}

	InfrastructureTargetVersionResourceConfig = ""
)

// issue-routing-tag: database/ExaCC
func TestDatabaseInfrastructureTargetVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseInfrastructureTargetVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetResourceId := utils.GetEnvSettingWithBlankDefault("target_resource_id")
	targetResourceIdVariableStr := fmt.Sprintf("variable \"target_resource_id\" { default = \"%s\" }\n", targetResourceId)

	singularDatasourceName := "data.oci_database_infrastructure_target_version.test_infrastructure_target_version"

	acctest.SaveConfigContent("", "", "", t)
	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_infrastructure_target_version", "test_infrastructure_target_version", acctest.Required, acctest.Create, infrastructureTargetVersionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + targetResourceIdVariableStr + InfrastructureTargetVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_resource_type"),

				resource.TestCheckResourceAttr(singularDatasourceName, "target_db_version_history_entry.#", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_resource_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_storage_version_history_entry.#", "1"),
			),
		},
	})
}
