// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OsManagementHubProfileAttachLifecycleStageManagementRepresentation = map[string]interface{}{
		"lifecycle_stage_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_os_management_hub_lifecycle_stage.test_lifecycle_stage.id}`},
		"profile_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_profile.test_profile.id}`},
	}

	OsManagementHubLCProfileRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"profile_type":          acctest.Representation{RepType: acctest.Required, Create: `LIFECYCLE`},
		"lifecycle_stage_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_lifecycle_environment.test_lifecycle_environment.stages[0].id}`},
		"management_station_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_management_station.test_management_station.id}`},
		"arch_type":             acctest.Representation{RepType: acctest.Optional, Create: `X86_64`},
		"description":           acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"os_family":             acctest.Representation{RepType: acctest.Optional, Create: `ORACLE_LINUX_8`},
		"vendor_name":           acctest.Representation{RepType: acctest.Optional, Create: `ORACLE`},
		"defined_tags":          acctest.Representation{RepType: acctest.Optional, Create: OsManagementHubProfileIgnoreDefinedTagsRepresentation},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Finance"}},
	}

	OsManagementHubProfileAttachLifecycleStageManagementResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_lifecycle_stages", "test_lifecycle_stages", acctest.Required, acctest.Create, OsManagementHubLifecycleStageDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group", "test_managed_instance_group", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Required, acctest.Create, OsManagementHubLCProfileRepresentation)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubProfileAttachLifecycleStageManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubProfileAttachLifecycleStageManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_os_management_hub_profile_attach_lifecycle_stage_management.test_profile_attach_lifecycle_stage_management"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsManagementHubProfileAttachLifecycleStageManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile_attach_lifecycle_stage_management", "test_profile_attach_lifecycle_stage_management", acctest.Required, acctest.Create, OsManagementHubProfileAttachLifecycleStageManagementRepresentation), "osmanagementhub", "profileAttachLifecycleStageManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubProfileAttachLifecycleStageManagementResourceDependencies + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile_attach_lifecycle_stage_management", "test_profile_attach_lifecycle_stage_management", acctest.Required, acctest.Create, OsManagementHubProfileAttachLifecycleStageManagementRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_lifecycle_stage", "test_lifecycle_stage", acctest.Required, acctest.Create, OsManagementHubLifecycleStageSingularDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_environment", "test_lifecycle_environment", acctest.Required, acctest.Create, OsManagementHubLifecycleEnvironmentRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_management_station", "test_management_station", acctest.Required, acctest.Create, OsManagementHubManagementStationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "lifecycle_stage_id"),
				resource.TestCheckResourceAttrSet(resourceName, "profile_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
	})
}
