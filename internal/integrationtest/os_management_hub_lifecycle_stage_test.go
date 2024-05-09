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
	OsManagementHubLifecycleStageSingularDataSourceRepresentation = map[string]interface{}{
		"lifecycle_stage_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_lifecycle_environment.test_lifecycle_environment.stages[0].id}`},
	}

	OsManagementHubLifecycleStageDataSourceRepresentation = map[string]interface{}{
		"arch_type":             acctest.Representation{RepType: acctest.Optional, Create: `X86_64`},
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: []string{`displayName`}},
		"display_name_contains": acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"lifecycle_stage_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_lifecycle_environment.test_lifecycle_environment.stages[0].id}`},
		"location":              acctest.Representation{RepType: acctest.Optional, Create: []string{`ON_PREMISE`}},
		"location_not_equal_to": acctest.Representation{RepType: acctest.Optional, Create: []string{`OCI_COMPUTE`}},
		"os_family":             acctest.Representation{RepType: acctest.Optional, Create: `ORACLE_LINUX_8`},
		"software_source_id":    acctest.Representation{RepType: acctest.Optional, Create: nil},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubLifecycleStageGroupDataSourceFilterRepresentation},
	}

	OsManagementHubLifecycleStageGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_os_management_hub_lifecycle_environment.test_lifecycle_environment.stages[0].id}`}},
	}

	OsManagementHubLifecycleStageSoftwareSourceIdDataSourceRepresentation = map[string]interface{}{}
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubLifecycleStageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubLifecycleStageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_os_management_hub_lifecycle_stages.test_lifecycle_stages"
	singularDatasourceName := "data.oci_os_management_hub_lifecycle_stage.test_lifecycle_stage"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + OsManagementHubLifecycleEnvironmentRequiredOnlyResource +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_lifecycle_stages", "test_lifecycle_stages", acctest.Optional, acctest.Create, OsManagementHubLifecycleStageDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "lifecycle_stage_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "lifecycle_stage_id"),
				resource.TestCheckResourceAttr(datasourceName, "location.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "location_not_equal_to.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttr(datasourceName, "software_source_id.#", "0"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "lifecycle_stage_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + OsManagementHubLifecycleEnvironmentRequiredOnlyResource +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_lifecycle_stage", "test_lifecycle_stage", acctest.Required, acctest.Create, OsManagementHubLifecycleStageSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "lifecycle_stage_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "arch_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "lifecycle_environment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "location"),
				resource.TestCheckResourceAttr(singularDatasourceName, "managed_instance_ids.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "os_family"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "rank"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_modified"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vendor_name"),
			),
		},
	})
}
