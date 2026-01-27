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
	OsManagementHubManagedInstanceInstallPackagesManagementRequiredOnlyResource = OsManagementHubManagedInstanceInstallPackagesManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_install_packages_management", "test_managed_instance_install_packages_management", acctest.Required, acctest.Create, OsManagementHubManagedInstanceInstallPackagesManagementRepresentation)

	OsManagementHubManagedInstanceInstallPackagesManagementRepresentation = map[string]interface{}{
		"managed_instance_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_managed_instance.test_managed_instance.id}`},
		"package_names":        acctest.Representation{RepType: acctest.Required, Create: []string{`ModemManager-1.20.2-1.el8.x86_64`}},
		"work_request_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubManagedInstanceInstallPackagesManagementWorkRequestDetailsRepresentation},
	}
	OsManagementHubManagedInstanceInstallPackagesManagementWorkRequestDetailsRepresentation = map[string]interface{}{
		"description":  acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}

	OsManagementHubManagedInstanceInstallPackagesManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance", "test_managed_instance", acctest.Required, acctest.Create, OsManagementHubManagedInstanceRepresentation)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagedInstanceInstallPackagesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagedInstanceInstallPackagesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_os_management_hub_managed_instance_install_packages_management.test_managed_instance_install_packages_management"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsManagementHubManagedInstanceInstallPackagesManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_install_packages_management", "test_managed_instance_install_packages_management", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceInstallPackagesManagementRepresentation), "osmanagementhub", "managedInstanceInstallPackagesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceInstallPackagesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_install_packages_management", "test_managed_instance_install_packages_management", acctest.Required, acctest.Create, OsManagementHubManagedInstanceInstallPackagesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "package_names.#", "1"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceInstallPackagesManagementResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceInstallPackagesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_install_packages_management", "test_managed_instance_install_packages_management", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceInstallPackagesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "package_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "work_request_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "work_request_details.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "work_request_details.0.display_name", "displayName"),

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
