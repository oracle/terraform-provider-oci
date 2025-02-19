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
	DatabaseScheduledActionParamDataSourceRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `DB_SERVER_FULL_SOFTWARE_UPDATE`},
	}

	DatabaseScheduledActionParamResourceConfig = ""
)

// issue-routing-tag: database/default
func TestDatabaseScheduledActionParamResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseScheduledActionParamResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_scheduled_action_params.test_scheduled_action_params"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_scheduled_action_params", "test_scheduled_action_params", acctest.Required, acctest.Create, DatabaseScheduledActionParamDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseScheduledActionParamResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "type", "DB_SERVER_FULL_SOFTWARE_UPDATE"),

				resource.TestCheckResourceAttrSet(datasourceName, "action_param_values_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "action_param_values_collection.0.items.#", "4"),
			),
		},
	})
}
