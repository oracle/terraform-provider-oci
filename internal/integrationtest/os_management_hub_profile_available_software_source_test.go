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
	OsManagementHubProfileAvailableSoftwareSourceDataSourceRepresentation = map[string]interface{}{
		"profile_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_profile.test_profile.id}`},
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: []string{`displayName`}},
		"display_name_contains": acctest.Representation{RepType: acctest.Optional, Create: `displayNameContains`},
	}

	OsManagementHubProfileAvailableSoftwareSourceResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_lifecycle_stages", "test_lifecycle_stages", acctest.Required, acctest.Create, OsManagementHubLifecycleStageDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group", "test_managed_instance_group", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Required, acctest.Create, OsManagementHubProfileRepresentation)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubProfileAvailableSoftwareSourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubProfileAvailableSoftwareSourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_os_management_hub_profile_available_software_sources.test_profile_available_software_sources"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// required parameters
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_profile_available_software_sources", "test_profile_available_software_sources", acctest.Required, acctest.Create, OsManagementHubProfileAvailableSoftwareSourceDataSourceRepresentation) +
				compartmentIdVariableStr + OsManagementHubProfileAvailableSoftwareSourceResourceConfig + OsManagementHubLifecycleEnvironmentRequiredOnlyResource + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "profile_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "available_software_source_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "available_software_source_collection.0.items.#"),
			),
		},
		// optional parameters
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_profile_available_software_sources", "test_profile_available_software_sources", acctest.Optional, acctest.Create, OsManagementHubProfileAvailableSoftwareSourceDataSourceRepresentation) +
				compartmentIdVariableStr + OsManagementHubProfileAvailableSoftwareSourceResourceConfig + OsManagementHubLifecycleEnvironmentRequiredOnlyResource + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "displayNameContains"),
				resource.TestCheckResourceAttrSet(datasourceName, "display_name.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "profile_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "available_software_source_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "available_software_source_collection.0.items.#"),
			),
		},
	})
}
