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
	instanceMaintenanceRebootSingularDataSourceRepresentation = map[string]interface{}{
		"instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
	}

	InstanceMaintenanceRebootResourceConfig = utils.OciImageIdsVariable
)

// issue-routing-tag: core/default
func TestCoreInstanceMaintenanceRebootResource_basic(t *testing.T) {
	t.Skip("Skipping this test as the instance ocid is hardcoded and may not exist when the test runs")

	httpreplay.SetScenario("TestCoreInstanceMaintenanceRebootResource_basic")
	defer httpreplay.SaveScenario()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	config := acctest.ProviderTestConfig() + acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance_maintenance_reboot", "test_instance_maintenance_reboot",
		acctest.Required, acctest.Create, instanceMaintenanceRebootSingularDataSourceRepresentation) + compartmentIdVariableStr + InstanceMaintenanceRebootResourceConfig

	singularDatasourceName := "data.oci_core_instance_maintenance_reboot.test_instance_maintenance_reboot"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_maintenance_reboot_due_max"),
			),
		},
	})
}
