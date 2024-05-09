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
	OsManagementHubManagedInstanceAvailablePackageDataSourceRepresentation = map[string]interface{}{
		"managed_instance_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_managed_instance.test_managed_instance.id}`},
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: []string{`displayName`}},
		"display_name_contains": acctest.Representation{RepType: acctest.Optional, Create: `displayNameContains`},
	}

	OsManagementHubManagedInstanceAvailablePackageResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance", "test_managed_instance", acctest.Required, acctest.Create, OsManagementHubManagedInstanceRepresentation)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagedInstanceAvailablePackageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagedInstanceAvailablePackageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_os_management_hub_managed_instance_available_packages.test_managed_instance_available_packages"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_managed_instance_available_packages", "test_managed_instance_available_packages", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceAvailablePackageDataSourceRepresentation) +
				compartmentIdVariableStr + OsManagementHubManagedInstanceAvailablePackageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "available_package_collection.#"),
			),
		},
	})
}
