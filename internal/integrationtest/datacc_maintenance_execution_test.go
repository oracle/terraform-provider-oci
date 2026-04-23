// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
)

var (
	DataccMaintenanceExecutionSingularDataSourceRepresentation = map[string]interface{}{
		"maintenance_execution_id": acctest.Representation{RepType: acctest.Required, Create: `${var.maintenance_execution_id}`},
	}

	DataccMaintenanceExecutionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.maintenance_execution_compartment_id}`},
		// "display_name":         acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		// "infrastructure_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.maintenance_execution_infrastructure_id}`},
		// "maintenance_run_id":   acctest.Representation{RepType: acctest.Optional, Create: `${var.maintenance_execution_maintenance_run_id}`},
		"maintenance_subtype":  acctest.Representation{RepType: acctest.Optional, Create: `YEARLY`},
		"maintenance_type":     acctest.Representation{RepType: acctest.Optional, Create: `PLANNED`},
		"state":                acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"target_resource_type": acctest.Representation{RepType: acctest.Optional, Create: `DB_CC_INFRASTRUCTURE`},
		// "time_accepted_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `timeAcceptedGreaterThanOrEqualTo`},
		// "time_accepted_less_than_or_equal_to":    acctest.Representation{RepType: acctest.Optional, Create: `timeAcceptedLessThanOrEqualTo`},
		"type": acctest.Representation{RepType: acctest.Optional, Create: `NOTIFY`},
	}

	// DataccMaintenanceExecutionResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_datacc_infrastructure", "test_infrastructure", acctest.Required, acctest.Create, DataccInfrastructureRepresentation) +
	// 	acctest.GenerateResourceFromRepresentationMap("oci_datacc_maintenance_run", "test_maintenance_run", acctest.Required, acctest.Create, DataccMaintenanceRunRepresentation)
)

// issue-routing-tag: datacc/default
func TestDataccMaintenanceExecutionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataccMaintenanceExecutionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	// override terraform-federation-test profile with our own user profile
	if overrideProfile := os.Getenv("datacc_custom_config_file_profile_override"); overrideProfile != "" {
		t.Setenv(globalvar.TfEnvPrefix+globalvar.ConfigFileProfileAttrName, overrideProfile)
		t.Setenv(globalvar.TfEnvPrefix+globalvar.AuthAttrName, "")
		t.Setenv(globalvar.AuthAttrName, globalvar.AuthSecurityToken)
	}

	const testResourceType = "maintenance_execution"
	tfVariableStr := GenerateTFVariableStrings(testResourceType)
	getTFVar := func(variableName string) string {
		return os.Getenv(globalvar.TfEnvPrefix + testResourceType + "_" + variableName)
	}

	compartmentId := getTFVar("compartment_id")

	datasourceName := "data.oci_datacc_maintenance_executions.test_maintenance_executions"
	singularDatasourceName := "data.oci_datacc_maintenance_execution.test_maintenance_execution"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + tfVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datacc_maintenance_executions", "test_maintenance_executions", acctest.Required, acctest.Create, DataccMaintenanceExecutionDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "maintenance_execution_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_execution_collection.0.items.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_execution_collection.0.items.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_execution_collection.0.items.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_execution_collection.0.items.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_execution_collection.0.items.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_execution_collection.0.items.0.infrastructure_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_execution_collection.0.items.0.maintenance_run_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_execution_collection.0.items.0.source_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_execution_collection.0.items.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_execution_collection.0.items.0.target_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_execution_collection.0.items.0.time_ended"),
				resource.TestCheckResourceAttrSet(datasourceName, "maintenance_execution_collection.0.items.0.time_started"),
			),
		},
		// verify singular datasource
		{
			Config: config + tfVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datacc_maintenance_execution", "test_maintenance_execution", acctest.Required, acctest.Create, DataccMaintenanceExecutionSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_execution_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "custom_action_timeout_in_mins"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "infrastructure_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_custom_action_timeout_enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "maintenance_run_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "source_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_ended"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
			),
		},
	})
}
