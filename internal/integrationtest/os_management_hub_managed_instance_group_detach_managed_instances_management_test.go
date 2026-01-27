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
	OsManagementHubManagedInstanceGroupDetachManagedInstancesManagementRepresentation = map[string]interface{}{
		"managed_instance_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_managed_instance_group.test_managed_instance_group.id}`},
		"managed_instances":         acctest.Representation{RepType: acctest.Required, Create: []string{utils.GetEnvSettingWithBlankDefault("osmh_managed_instance_ocid")}},
		"depends_on":                acctest.Representation{RepType: acctest.Required, Create: []string{"oci_os_management_hub_managed_instance_group_attach_managed_instances_management.test_managed_instance_group_attach_managed_instances_management"}},
	}

	OsManagementHubManagedInstanceGroupDetachManagedInstancesManagementResourceDependencies = ""
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagedInstanceGroupDetachManagedInstancesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagedInstanceGroupDetachManagedInstancesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_os_management_hub_managed_instance_group_detach_managed_instances_management.test_managed_instance_group_detach_managed_instances_management"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsManagementHubManagedInstanceGroupDetachManagedInstancesManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group_detach_managed_instances_management", "test_managed_instance_group_detach_managed_instances_management", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupDetachManagedInstancesManagementRepresentation), "osmanagementhub", "managedInstanceGroupDetachManagedInstancesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			// Config includes Dependencies (creating the group), the Detach Resource, and the Attach Resource.
			// The `depends_on` in the Representation map ensures Attach happens first.
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceGroupAttachManagedInstancesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group_detach_managed_instances_management", "test_managed_instance_group_detach_managed_instances_management", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupDetachManagedInstancesManagementRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group", "test_managed_instance_group", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceGroupRepresentation) + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group_attach_managed_instances_management", "test_managed_instance_group_attach_managed_instances_management", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupAttachManagedInstancesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_group_id"),
				resource.TestCheckResourceAttr(resourceName, "managed_instances.#", "1"),

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
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceGroupDetachManagedInstancesManagementResourceDependencies,
		},
	})
}
