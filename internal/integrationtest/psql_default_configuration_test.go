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
	PsqlDefaultConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"default_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${var.default_config_id}`},
	}
	PsqlDefaultConfigurationDataSourceRepresentation = map[string]interface{}{
		//"configuration_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.default_config_id}`},
		"db_version":   acctest.Representation{RepType: acctest.Optional, Create: `14`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `PostgreSQL.VM.Standard.E4.Flex.2.32GB-14-0_45`},
		"shape":        acctest.Representation{RepType: acctest.Optional, Create: `VM.Standard.E4.Flex.2.32GB`},
		"state":        acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}

	PsqlDefaultConfigurationDataSourceRepresentation2 = map[string]interface{}{
		"configuration_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.default_config_id}`},
	}

	PsqlDefaultConfigurationResourceConfig = ""
)

// issue-routing-tag: psql/default
func TestPsqlDefaultConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestPsqlDefaultConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	defaultConfigId := utils.GetEnvSettingWithBlankDefault("default_config_ocid")
	defaultConfigIdVariableStr := fmt.Sprintf("variable \"default_config_id\" { default = \"%s\" }\n", defaultConfigId)

	datasourceName := "data.oci_psql_default_configurations.test_default_configurations"
	singularDatasourceName := "data.oci_psql_default_configuration.test_default_configuration"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + defaultConfigIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_psql_default_configurations", "test_default_configurations", acctest.Optional, acctest.Create, PsqlDefaultConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + PsqlDefaultConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttr(datasourceName, "db_version", "14"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "PostgreSQL.VM.Standard.E4.Flex.2.32GB-14-0_45"),
				resource.TestCheckResourceAttr(datasourceName, "shape", "VM.Standard.E4.Flex.2.32GB"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "default_configuration_collection.0.items.0.id", defaultConfigId),
				resource.TestCheckResourceAttr(datasourceName, "default_configuration_collection.0.items.0.db_version", "14"),
				resource.TestCheckResourceAttr(datasourceName, "default_configuration_collection.0.items.0.display_name", "PostgreSQL.VM.Standard.E4.Flex.2.32GB-14-0_45"),
				resource.TestCheckResourceAttr(datasourceName, "default_configuration_collection.0.items.0.shape", "VM.Standard.E4.Flex.2.32GB"),
				resource.TestCheckResourceAttr(datasourceName, "default_configuration_collection.0.items.0.state", "ACTIVE"),

				resource.TestCheckResourceAttrSet(datasourceName, "default_configuration_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "default_configuration_collection.0.items.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "default_configuration_collection.0.items.0.db_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "default_configuration_collection.0.items.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "default_configuration_collection.0.items.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "default_configuration_collection.0.items.0.instance_memory_size_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "default_configuration_collection.0.items.0.instance_ocpu_count"),
				resource.TestCheckResourceAttrSet(datasourceName, "default_configuration_collection.0.items.0.shape"),
				resource.TestCheckResourceAttrSet(datasourceName, "default_configuration_collection.0.items.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "default_configuration_collection.0.items.0.time_created"),
			),
		},
		// verify datasource configuration_id
		{
			Config: config + defaultConfigIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_psql_default_configurations", "test_default_configurations", acctest.Optional, acctest.Create, PsqlDefaultConfigurationDataSourceRepresentation2) +
				compartmentIdVariableStr + PsqlDefaultConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "configuration_id"),

				resource.TestCheckResourceAttr(datasourceName, "default_configuration_collection.0.items.0.id", defaultConfigId),
				resource.TestCheckResourceAttr(datasourceName, "default_configuration_collection.0.items.0.db_version", "14"),
				resource.TestCheckResourceAttr(datasourceName, "default_configuration_collection.0.items.0.display_name", "PostgreSQL.VM.Standard.E4.Flex.2.32GB-14-0_45"),
				resource.TestCheckResourceAttr(datasourceName, "default_configuration_collection.0.items.0.shape", "VM.Standard.E4.Flex.2.32GB"),
				resource.TestCheckResourceAttr(datasourceName, "default_configuration_collection.0.items.0.state", "ACTIVE"),

				resource.TestCheckResourceAttrSet(datasourceName, "default_configuration_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "default_configuration_collection.0.items.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "default_configuration_collection.0.items.0.db_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "default_configuration_collection.0.items.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "default_configuration_collection.0.items.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "default_configuration_collection.0.items.0.instance_memory_size_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "default_configuration_collection.0.items.0.instance_ocpu_count"),
				resource.TestCheckResourceAttrSet(datasourceName, "default_configuration_collection.0.items.0.shape"),
				resource.TestCheckResourceAttrSet(datasourceName, "default_configuration_collection.0.items.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "default_configuration_collection.0.items.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config + defaultConfigIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_psql_default_configuration", "test_default_configuration", acctest.Required, acctest.Create, PsqlDefaultConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + PsqlDefaultConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "default_configuration_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "configuration_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_memory_size_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_ocpu_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_flexible"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
	})
}
