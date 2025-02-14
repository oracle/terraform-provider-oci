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
	OsManagementHubProfileAttachManagedInstanceGroupManagementRepresentation = map[string]interface{}{
		"managed_instance_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_managed_instance_group.test_managed_instance_group.id}`},
		"profile_id":                acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_profile.test_profile.id}`},
	}

	OsManagementHubGrpProfileRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":              acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"profile_type":              acctest.Representation{RepType: acctest.Required, Create: `GROUP`},
		"managed_instance_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_managed_instance_group.test_managed_instance_group.id}`},
		"arch_type":                 acctest.Representation{RepType: acctest.Optional, Create: `X86_64`},
		"registration_type":         acctest.Representation{RepType: acctest.Required, Create: `OCI_LINUX`},
		"description":               acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"os_family":                 acctest.Representation{RepType: acctest.Optional, Create: `ORACLE_LINUX_8`},
		"vendor_name":               acctest.Representation{RepType: acctest.Optional, Create: `ORACLE`},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: OsManagementHubProfileIgnoreDefinedTagsRepresentation},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Finance"}},
	}

	OsManagementHubManagedInstanceGroupRepresentation2 = map[string]interface{}{
		"arch_type":            acctest.Representation{RepType: acctest.Required, Create: `X86_64`},
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":         acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"os_family":            acctest.Representation{RepType: acctest.Required, Create: `ORACLE_LINUX_8`},
		"vendor_name":          acctest.Representation{RepType: acctest.Required, Create: `ORACLE`},
		"defined_tags":         acctest.Representation{RepType: acctest.Optional, Create: OsManagementHubManagedInstanceGroupIgnoreDefinedTagsRepresentation},
		"description":          acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"location":             acctest.Representation{RepType: acctest.Required, Create: `OCI_COMPUTE`},
		"managed_instance_ids": acctest.Representation{RepType: acctest.Optional, Create: []string{}},
		"software_source_ids":  acctest.Representation{RepType: acctest.Required, Create: []string{`${data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id}`}},
	}

	OsManagementHubProfileAttachManagedInstanceGroupManagementResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_lifecycle_stages", "test_lifecycle_stages", acctest.Required, acctest.Create, OsManagementHubLifecycleStageDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group", "test_managed_instance_group", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupRepresentation2) +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Required, acctest.Create, OsManagementHubGrpProfileRepresentation)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubProfileAttachManagedInstanceGroupManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubProfileAttachManagedInstanceGroupManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_os_management_hub_profile_attach_managed_instance_group_management.test_profile_attach_managed_instance_group_management"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OsManagementHubProfileAttachManagedInstanceGroupManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile_attach_managed_instance_group_management", "test_profile_attach_managed_instance_group_management", acctest.Required, acctest.Create, OsManagementHubProfileAttachManagedInstanceGroupManagementRepresentation), "osmanagementhub", "profileAttachManagedInstanceGroupManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubProfileAttachManagedInstanceGroupManagementResourceDependencies +
				OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config + OsManagementHubLifecycleEnvironmentRequiredOnlyResource +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile_attach_managed_instance_group_management", "test_profile_attach_managed_instance_group_management", acctest.Required, acctest.Create, OsManagementHubProfileAttachManagedInstanceGroupManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_group_id"),
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
