// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OsManagementHubManagedInstanceGroupAttachManagedInstancesManagementRequiredOnlyResource = OsManagementHubManagedInstanceGroupAttachManagedInstancesManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group_attach_managed_instances_management", "test_managed_instance_group_attach_managed_instances_management", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupAttachManagedInstancesManagementRepresentation)

	OsManagementHubManagedInstanceGroupAttachManagedInstancesManagementRepresentation = map[string]interface{}{
		"managed_instance_group_id": acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("managed_instance_group_ocid")},
		"managed_instances":         acctest.Representation{RepType: acctest.Required, Create: []string{utils.GetEnvSettingWithBlankDefault("managed_instance_for_mig_ocid")}},
		"work_request_details":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubManagedInstanceGroupAttachManagedInstancesManagementWorkRequestDetailsRepresentation},
	}
	OsManagementHubManagedInstanceGroupAttachManagedInstancesManagementWorkRequestDetailsRepresentation = map[string]interface{}{
		"description":  acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}

	OsManagementHubManagedInstanceGroupAttachManagedInstancesManagementResourceDependencies = ""
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagedInstanceGroupAttachManagedInstancesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagedInstanceGroupAttachManagedInstancesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_os_management_hub_managed_instance_group_attach_managed_instances_management.test_managed_instance_group_attach_managed_instances_management"
	resourceNameDetach := "oci_os_management_hub_managed_instance_group_detach_managed_instances_management.test_managed_instance_group_detach_managed_instances_management"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsManagementHubManagedInstanceGroupAttachManagedInstancesManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group_attach_managed_instances_management", "test_managed_instance_group_attach_managed_instances_management", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceGroupAttachManagedInstancesManagementRepresentation), "osmanagementhub", "managedInstanceGroupAttachManagedInstancesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceGroupAttachManagedInstancesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group_attach_managed_instances_management", "test_managed_instance_group_attach_managed_instances_management", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupAttachManagedInstancesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_group_id"),
				resource.TestCheckResourceAttr(resourceName, "managed_instances.#", "1"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceGroupAttachManagedInstancesManagementResourceDependencies,
		},
		// Detach managed instance
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceGroupDetachManagedInstancesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group_detach_managed_instances_management", "test_managed_instance_group_detach_managed_instances_management", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupDetachManagedInstancesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameDetach, "managed_instance_group_id"),
				resource.TestCheckResourceAttr(resourceNameDetach, "managed_instances.#", "1"),
			),
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceGroupAttachManagedInstancesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group_attach_managed_instances_management", "test_managed_instance_group_attach_managed_instances_management", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceGroupAttachManagedInstancesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_group_id"),
				resource.TestCheckResourceAttr(resourceName, "managed_instances.#", "1"),
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
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceGroupAttachManagedInstancesManagementResourceDependencies,
		},
		// Detach managed instance
		{
			Config: config + compartmentIdVariableStr + OsManagementHubManagedInstanceGroupDetachManagedInstancesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group_detach_managed_instances_management", "test_managed_instance_group_detach_managed_instances_management", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupDetachManagedInstancesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameDetach, "managed_instance_group_id"),
				resource.TestCheckResourceAttr(resourceNameDetach, "managed_instances.#", "1"),
			),
		},
	})
}
