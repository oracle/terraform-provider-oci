// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OsManagementHubIgnoreDefinedTagsRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	OsManagementHubSoftwareSourceSingularDataSourceRepresentation = map[string]interface{}{
		"software_source_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id}`},
	}

	OsManagementHubSoftwareSourceDataSourceRepresentation = map[string]interface{}{
		"arch_type":                 acctest.Representation{RepType: acctest.Optional, Create: []string{`X86_64`}},
		"availability":              acctest.Representation{RepType: acctest.Optional, Create: []string{`SELECTED`}},
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `ol8_baseos_latest-x86_64`},
		"display_name_contains":     acctest.Representation{RepType: acctest.Optional, Create: `ol8_baseos_latest-x86_64`},
		"display_name_not_equal_to": acctest.Representation{RepType: acctest.Optional, Create: []string{`displayNameNotEqualTo`}},
		"os_family":                 acctest.Representation{RepType: acctest.Optional, Create: []string{`ORACLE_LINUX_8`}},
		"software_source_id":        acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id}`},
		"software_source_type":      acctest.Representation{RepType: acctest.Optional, Create: []string{`VENDOR`}},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"vendor_name":               acctest.Representation{RepType: acctest.Optional, Create: `ORACLE`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubSoftwareSourceDataSourceFilterRepresentation},
	}
	OsManagementHubSoftwareSourceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id}`}},
	}

	OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Representation = map[string]interface{}{
		"arch_type":            acctest.Representation{RepType: acctest.Optional, Create: []string{`X86_64`}},
		"availability":         acctest.Representation{RepType: acctest.Optional, Create: []string{`SELECTED`}},
		"compartment_id":       acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":         acctest.Representation{RepType: acctest.Optional, Create: `ol8_baseos_latest-x86_64`},
		"os_family":            acctest.Representation{RepType: acctest.Optional, Create: []string{`ORACLE_LINUX_8`}},
		"software_source_type": acctest.Representation{RepType: acctest.Optional, Create: []string{`VENDOR`}},
		"state":                acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"vendor_name":          acctest.Representation{RepType: acctest.Optional, Create: `ORACLE`},
	}

	OsManagementHubVendorSoftwareSourceOl8AppstreamX8664Representation = map[string]interface{}{
		"arch_type":            acctest.Representation{RepType: acctest.Optional, Create: []string{`X86_64`}},
		"availability":         acctest.Representation{RepType: acctest.Optional, Create: []string{`SELECTED`}},
		"compartment_id":       acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":         acctest.Representation{RepType: acctest.Optional, Create: `ol8_appstream-x86_64`},
		"os_family":            acctest.Representation{RepType: acctest.Optional, Create: []string{`ORACLE_LINUX_8`}},
		"software_source_type": acctest.Representation{RepType: acctest.Optional, Create: []string{`VENDOR`}},
		"state":                acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"vendor_name":          acctest.Representation{RepType: acctest.Optional, Create: `ORACLE`},
	}

	OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config = acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_software_sources", "ol8_baseos_latest_x86_64", acctest.Optional, acctest.Create, OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Representation)
	OsManagementHubVendorSoftwareSourceOl8AppstreamX8664Config    = acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_software_sources", "ol8_appstream_x86_64", acctest.Optional, acctest.Create, OsManagementHubVendorSoftwareSourceOl8AppstreamX8664Representation)

	OsManagementHubSoftwareSourceRepresentation = map[string]interface{}{
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                  acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"software_source_type":          acctest.Representation{RepType: acctest.Required, Create: `CUSTOM`},
		"vendor_software_sources":       []acctest.RepresentationGroup{{RepType: acctest.Required, Group: OsManagementHubSoftwareSourceVendorSoftwareSourcesRepresentation}, {RepType: acctest.Required, Group: OsManagementHubSoftwareSourceVendorSoftwareSourcesRepresentation2}},
		"custom_software_source_filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubSoftwareSourceCustomSoftwareSourceFilterRepresentation},
		"description":                   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		"is_automatically_updated":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"defined_tags":                  acctest.Representation{RepType: acctest.Optional, Create: OsManagementHubIgnoreDefinedTagsRepresentation},
	}
	OsManagementHubSoftwareSourceVendorSoftwareSourcesRepresentation = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `ol8_appstream-x86_64`},
		"id":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_os_management_hub_software_sources.ol8_appstream_x86_64.software_source_collection[0].items[0].id}`},
	}
	OsManagementHubSoftwareSourceVendorSoftwareSourcesRepresentation2 = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `ol8_baseos_latest-x86_64`},
		"id":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id}`},
	}
	OsManagementHubSoftwareSourceCustomSoftwareSourceFilterRepresentation = map[string]interface{}{
		"module_stream_profile_filters": acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubSoftwareSourceCustomSoftwareSourceFilterModuleStreamProfileFiltersRepresentation},
		"package_filters":               acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubSoftwareSourceCustomSoftwareSourceFilterPackageFiltersRepresentation},
		"package_group_filters":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: OsManagementHubSoftwareSourceCustomSoftwareSourceFilterPackageGroupFiltersRepresentation},
	}
	OsManagementHubSoftwareSourceCustomSoftwareSourceFilterModuleStreamProfileFiltersRepresentation = map[string]interface{}{
		"filter_type":  acctest.Representation{RepType: acctest.Required, Create: `INCLUDE`},
		"module_name":  acctest.Representation{RepType: acctest.Required, Create: `php`},
		"profile_name": acctest.Representation{RepType: acctest.Optional, Create: `common`},
		"stream_name":  acctest.Representation{RepType: acctest.Optional, Create: `8.0`},
	}
	OsManagementHubSoftwareSourceCustomSoftwareSourceFilterPackageFiltersRepresentation = map[string]interface{}{
		"filter_type":  acctest.Representation{RepType: acctest.Required, Create: `INCLUDE`},
		"package_name": acctest.Representation{RepType: acctest.Required, Create: `ed`},
		// "package_name_pattern": acctest.Representation{RepType: acctest.Optional, Create: `ed`},
		// "package_version":      acctest.Representation{RepType: acctest.Optional, Create: `packageVersion`, Update: `packageVersion2`},
	}
	OsManagementHubSoftwareSourceCustomSoftwareSourceFilterPackageGroupFiltersRepresentation = map[string]interface{}{
		"filter_type":    acctest.Representation{RepType: acctest.Required, Create: `INCLUDE`},
		"package_groups": acctest.Representation{RepType: acctest.Optional, Create: []string{`base`}},
	}
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubSoftwareSourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubSoftwareSourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_os_management_hub_software_source.test_software_source"
	datasourceName := "data.oci_os_management_hub_software_sources.test_software_sources"
	singularDatasourceName := "data.oci_os_management_hub_software_source.test_software_source"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source", "test_software_source", acctest.Optional, acctest.Create, OsManagementHubSoftwareSourceRepresentation), "osmanagementhub", "softwareSource", t)

	acctest.ResourceTest(t, testAccCheckOsManagementHubSoftwareSourceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config + OsManagementHubVendorSoftwareSourceOl8AppstreamX8664Config + acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source", "test_software_source", acctest.Required, acctest.Create, OsManagementHubSoftwareSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "software_source_type", "CUSTOM"),
				resource.TestCheckResourceAttr(resourceName, "vendor_software_sources.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "vendor_software_sources.0.display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "vendor_software_sources.0.id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config + OsManagementHubVendorSoftwareSourceOl8AppstreamX8664Config + acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source", "test_software_source", acctest.Optional, acctest.Create, OsManagementHubSoftwareSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "arch_type"),
				resource.TestCheckResourceAttrSet(resourceName, "availability"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "custom_software_source_filter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "custom_software_source_filter.0.module_stream_profile_filters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "custom_software_source_filter.0.module_stream_profile_filters.0.filter_type", "INCLUDE"),
				resource.TestCheckResourceAttr(resourceName, "custom_software_source_filter.0.module_stream_profile_filters.0.module_name", "php"),
				resource.TestCheckResourceAttrSet(resourceName, "custom_software_source_filter.0.module_stream_profile_filters.0.profile_name"),
				resource.TestCheckResourceAttrSet(resourceName, "custom_software_source_filter.0.module_stream_profile_filters.0.stream_name"),
				resource.TestCheckResourceAttr(resourceName, "custom_software_source_filter.0.package_filters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "custom_software_source_filter.0.package_filters.0.filter_type", "INCLUDE"),
				resource.TestCheckResourceAttr(resourceName, "custom_software_source_filter.0.package_filters.0.package_name", "ed"),
				resource.TestCheckResourceAttr(resourceName, "custom_software_source_filter.0.package_group_filters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "custom_software_source_filter.0.package_group_filters.0.filter_type", "INCLUDE"),
				resource.TestCheckResourceAttr(resourceName, "custom_software_source_filter.0.package_group_filters.0.package_groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_automatically_updated", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "os_family"),
				resource.TestCheckResourceAttrSet(resourceName, "repo_id"),
				resource.TestCheckResourceAttr(resourceName, "software_source_type", "CUSTOM"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "url"),
				resource.TestCheckResourceAttr(resourceName, "vendor_software_sources.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "vendor_software_sources.0.display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "vendor_software_sources.0.id"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config + OsManagementHubVendorSoftwareSourceOl8AppstreamX8664Config + acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source", "test_software_source", acctest.Optional, acctest.Update, OsManagementHubSoftwareSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "arch_type"),
				resource.TestCheckResourceAttrSet(resourceName, "availability"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "custom_software_source_filter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "custom_software_source_filter.0.module_stream_profile_filters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "custom_software_source_filter.0.module_stream_profile_filters.0.filter_type", "INCLUDE"),
				resource.TestCheckResourceAttr(resourceName, "custom_software_source_filter.0.module_stream_profile_filters.0.module_name", "php"),
				resource.TestCheckResourceAttrSet(resourceName, "custom_software_source_filter.0.module_stream_profile_filters.0.profile_name"),
				resource.TestCheckResourceAttrSet(resourceName, "custom_software_source_filter.0.module_stream_profile_filters.0.stream_name"),
				resource.TestCheckResourceAttr(resourceName, "custom_software_source_filter.0.package_filters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "custom_software_source_filter.0.package_filters.0.filter_type", "INCLUDE"),
				resource.TestCheckResourceAttr(resourceName, "custom_software_source_filter.0.package_filters.0.package_name", "ed"),
				resource.TestCheckResourceAttr(resourceName, "custom_software_source_filter.0.package_group_filters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "custom_software_source_filter.0.package_group_filters.0.filter_type", "INCLUDE"),
				resource.TestCheckResourceAttr(resourceName, "custom_software_source_filter.0.package_group_filters.0.package_groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_automatically_updated", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "os_family"),
				resource.TestCheckResourceAttrSet(resourceName, "repo_id"),
				resource.TestCheckResourceAttr(resourceName, "software_source_type", "CUSTOM"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "url"),
				resource.TestCheckResourceAttr(resourceName, "vendor_software_sources.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "vendor_software_sources.0.display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "vendor_software_sources.0.id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify resource import
		{
			Config:                  config + acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source", "test_software_source", acctest.Required, acctest.Create, OsManagementHubSoftwareSourceRepresentation),
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"state", "package_count"},
			ResourceName:            resourceName,
		},
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config + OsManagementHubVendorSoftwareSourceOl8AppstreamX8664Config + acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_software_sources", "test_software_sources", acctest.Optional, acctest.Update, OsManagementHubSoftwareSourceDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "arch_type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "availability.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "ol8_baseos_latest-x86_64"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "ol8_baseos_latest-x86_64"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_not_equal_to.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "os_family.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "software_source_id"),
				resource.TestCheckResourceAttr(datasourceName, "software_source_type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "state.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "vendor_name", "ORACLE"),

				resource.TestCheckResourceAttr(datasourceName, "software_source_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config + OsManagementHubVendorSoftwareSourceOl8AppstreamX8664Config + acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_software_source", "test_software_source", acctest.Required, acctest.Create, OsManagementHubSoftwareSourceSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "software_source_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "arch_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "checksum_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "ol8_baseos_latest-x86_64"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "gpg_key_fingerprint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "gpg_key_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "gpg_key_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "os_family"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "package_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "repo_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "software_source_type", "VENDOR"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vendor_name"),
			),
		},
	})
}

func testAccCheckOsManagementHubSoftwareSourceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).SoftwareSourceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_os_management_hub_software_source" {
			noResourceFound = false
			request := oci_os_management_hub.GetSoftwareSourceRequest{}

			tmp := rs.Primary.ID
			request.SoftwareSourceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "os_management_hub")

			response, err := client.GetSoftwareSource(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_os_management_hub.SoftwareSourceLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					fmt.Println("resource lifecycle state is not in expected deleted lifecycle states", response.GetLifecycleState())
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("OsManagementHubSoftwareSource") {
		resource.AddTestSweepers("OsManagementHubSoftwareSource", &resource.Sweeper{
			Name:         "OsManagementHubSoftwareSource",
			Dependencies: acctest.DependencyGraph["softwareSource"],
			F:            sweepOsManagementHubSoftwareSourceResource,
		})
	}
}

func sweepOsManagementHubSoftwareSourceResource(compartment string) error {
	softwareSourceClient := acctest.GetTestClients(&schema.ResourceData{}).SoftwareSourceClient()
	softwareSourceIds, err := getOsManagementHubSoftwareSourceIds(compartment)
	if err != nil {
		return err
	}
	for _, softwareSourceId := range softwareSourceIds {
		if ok := acctest.SweeperDefaultResourceId[softwareSourceId]; !ok {
			deleteSoftwareSourceRequest := oci_os_management_hub.DeleteSoftwareSourceRequest{}

			deleteSoftwareSourceRequest.SoftwareSourceId = &softwareSourceId

			deleteSoftwareSourceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "os_management_hub")
			_, error := softwareSourceClient.DeleteSoftwareSource(context.Background(), deleteSoftwareSourceRequest)
			if error != nil {
				fmt.Printf("Error deleting SoftwareSource %s %s, It is possible that the resource is already deleted. Please verify manually \n", softwareSourceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &softwareSourceId, OsManagementHubSoftwareSourceSweepWaitCondition, time.Duration(3*time.Minute),
				OsManagementHubSoftwareSourceSweepResponseFetchOperation, "os_management_hub", true)
		}
	}
	return nil
}

func getOsManagementHubSoftwareSourceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SoftwareSourceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	softwareSourceClient := acctest.GetTestClients(&schema.ResourceData{}).SoftwareSourceClient()

	listSoftwareSourcesRequest := oci_os_management_hub.ListSoftwareSourcesRequest{}
	listSoftwareSourcesRequest.CompartmentId = &compartmentId
	listSoftwareSourcesRequest.LifecycleState = []oci_os_management_hub.SoftwareSourceLifecycleStateEnum{oci_os_management_hub.SoftwareSourceLifecycleStateActive}
	listSoftwareSourcesResponse, err := softwareSourceClient.ListSoftwareSources(context.Background(), listSoftwareSourcesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SoftwareSource list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, softwareSource := range listSoftwareSourcesResponse.Items {
		id := *softwareSource.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SoftwareSourceId", id)
	}
	return resourceIds, nil
}

func OsManagementHubSoftwareSourceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if softwareSourceResponse, ok := response.Response.(oci_os_management_hub.GetSoftwareSourceResponse); ok {
		return softwareSourceResponse.GetLifecycleState() != oci_os_management_hub.SoftwareSourceLifecycleStateDeleted
	}
	return false
}

func OsManagementHubSoftwareSourceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.SoftwareSourceClient().GetSoftwareSource(context.Background(), oci_os_management_hub.GetSoftwareSourceRequest{
		SoftwareSourceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
