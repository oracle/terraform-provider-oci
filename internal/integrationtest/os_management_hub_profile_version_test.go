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
	OsManagementHubProfileVersionSingularDataSourceRepresentation = map[string]interface{}{
		"profile_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_profile.test_profile.id}`},
		"profile_version": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_profile.test_profile.profile_version}`},
	}

	OsManagementHubProfileVersionResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_lifecycle_stages", "test_lifecycle_stages", acctest.Required, acctest.Create, OsManagementHubLifecycleStageDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group", "test_managed_instance_group", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Required, acctest.Create, OsManagementHubProfileRepresentation)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubProfileVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubProfileVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_os_management_hub_profile_version.test_profile_version"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_profile_version", "test_profile_version", acctest.Required, acctest.Create, OsManagementHubProfileVersionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OsManagementHubProfileVersionResourceConfig + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config + OsManagementHubLifecycleEnvironmentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "profile_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "profile_version"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "arch_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_default_profile"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_service_provided_profile"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lifecycle_environment.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lifecycle_stage.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "managed_instance_group.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "os_family"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "profile_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "profile_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "registration_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "software_sources.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_modified"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vendor_name"),
			),
		},
	})
}
