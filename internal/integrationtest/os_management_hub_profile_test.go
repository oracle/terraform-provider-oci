// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OsManagementHubSoftwareSourceProfileRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Required, acctest.Create, OsManagementHubSoftwareSourceProfileRepresentation)
	OsManagementHubSoftwareSourceProfileResourceConfig       = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Optional, acctest.Update, OsManagementHubSoftwareSourceProfileRepresentation)
	OsManagementHubGroupProfileRequiredOnlyResource          = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Required, acctest.Create, OsManagementHubGroupProfileRepresentation)
	OsManagementHubGroupProfileResourceConfig                = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Optional, acctest.Update, OsManagementHubGroupProfileRepresentation)
	OsManagementHubLifecycleProfileRequiredOnlyResource      = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Required, acctest.Create, OsManagementHubLifecycleProfileRepresentation)
	OsManagementHubLifecycleProfileResourceConfig            = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Optional, acctest.Update, OsManagementHubLifecycleProfileRepresentation)

	OsManagementHubProfileSingularDataSourceRepresentation = map[string]interface{}{
		"profile_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_profile.test_profile.id}`},
	}

	OsManagementHubSoftwareSourceProfileDataSourceRepresentation = map[string]interface{}{
		"arch_type":             acctest.Representation{RepType: acctest.Optional, Create: `X86_64`},
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: []string{`displayName`}, Update: []string{`displayName2`}},
		"display_name_contains": acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"os_family":             acctest.Representation{RepType: acctest.Optional, Create: `ORACLE_LINUX_8`},
		"profile_id":            acctest.Representation{RepType: acctest.Optional, Create: `${oci_os_management_hub_profile.test_profile.id}`},
		"profile_type":          acctest.Representation{RepType: acctest.Optional, Create: []string{`SOFTWARESOURCE`}},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"vendor_name":           acctest.Representation{RepType: acctest.Optional, Create: `ORACLE`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubProfileDataSourceFilterRepresentation}}
	OsManagementHubGroupProfileDataSourceRepresentation = map[string]interface{}{
		"arch_type":             acctest.Representation{RepType: acctest.Optional, Create: `X86_64`},
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: []string{`displayName`}, Update: []string{`displayName2`}},
		"display_name_contains": acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"os_family":             acctest.Representation{RepType: acctest.Optional, Create: `ORACLE_LINUX_8`},
		"profile_id":            acctest.Representation{RepType: acctest.Optional, Create: `${oci_os_management_hub_profile.test_profile.id}`},
		"profile_type":          acctest.Representation{RepType: acctest.Optional, Create: []string{`GROUP`}},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"vendor_name":           acctest.Representation{RepType: acctest.Optional, Create: `ORACLE`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubProfileDataSourceFilterRepresentation}}
	OsManagementHubLifecycleProfileDataSourceRepresentation = map[string]interface{}{
		"arch_type":             acctest.Representation{RepType: acctest.Optional, Create: `X86_64`},
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: []string{`displayName`}, Update: []string{`displayName2`}},
		"display_name_contains": acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"os_family":             acctest.Representation{RepType: acctest.Optional, Create: `ORACLE_LINUX_8`},
		"profile_id":            acctest.Representation{RepType: acctest.Optional, Create: `${oci_os_management_hub_profile.test_profile.id}`},
		"profile_type":          acctest.Representation{RepType: acctest.Optional, Create: []string{`LIFECYCLE`}},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"vendor_name":           acctest.Representation{RepType: acctest.Optional, Create: `ORACLE`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubProfileDataSourceFilterRepresentation}}
	OsManagementHubProfileDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_os_management_hub_profile.test_profile.id}`}},
	}

	OsManagementHubProfileRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":              acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"profile_type":              acctest.Representation{RepType: acctest.Required, Create: `SOFTWARESOURCE`},
		"arch_type":                 acctest.Representation{RepType: acctest.Required, Create: `X86_64`},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: OsManagementHubProfileIgnoreDefinedTagsRepresentation},
		"description":               acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle_stage_id":        acctest.Representation{RepType: acctest.Optional, Create: `${oci_os_management_hub_lifecycle_stage.test_lifecycle_stage.id}`},
		"managed_instance_group_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_os_management_hub_managed_instance_group.test_managed_instance_group.id}`},
		"management_station_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_os_management_hub_management_station.test_management_station.id}`},
		"os_family":                 acctest.Representation{RepType: acctest.Optional, Create: `ORACLE_LINUX_8`},
		"software_source_ids":       acctest.Representation{RepType: acctest.Optional, Create: []string{`softwareSourceIds`}},
		"vendor_name":               acctest.Representation{RepType: acctest.Optional, Create: `ORACLE`},
	}

	OsManagementHubSoftwareSourceProfileRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"profile_type":          acctest.Representation{RepType: acctest.Required, Create: `SOFTWARESOURCE`},
		"software_source_ids":   acctest.Representation{RepType: acctest.Required, Create: []string{`${data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id}`}},
		"management_station_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_management_station.test_management_station.id}`},
		"arch_type":             acctest.Representation{RepType: acctest.Required, Create: `X86_64`},
		"os_family":             acctest.Representation{RepType: acctest.Required, Create: `ORACLE_LINUX_8`},
		"vendor_name":           acctest.Representation{RepType: acctest.Required, Create: `ORACLE`},
		"description":           acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"defined_tags":          acctest.Representation{RepType: acctest.Optional, Create: OsManagementHubProfileIgnoreDefinedTagsRepresentation},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	OsManagementHubGroupProfileRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":              acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"profile_type":              acctest.Representation{RepType: acctest.Required, Create: `GROUP`},
		"managed_instance_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_managed_instance_group.test_managed_instance_group.id}`},
		"management_station_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_management_station.test_management_station.id}`},
		"arch_type":                 acctest.Representation{RepType: acctest.Optional, Create: `X86_64`},
		"description":               acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"os_family":                 acctest.Representation{RepType: acctest.Optional, Create: `ORACLE_LINUX_8`},
		"vendor_name":               acctest.Representation{RepType: acctest.Optional, Create: `ORACLE`},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: OsManagementHubProfileIgnoreDefinedTagsRepresentation},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	OsManagementHubLifecycleProfileRepresentation = map[string]interface{}{
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
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	OsManagementHubProfileResourceDependencies = OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config + acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_management_station", "test_management_station", acctest.Required, acctest.Create, OsManagementHubManagementStationRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group", "test_managed_instance_group", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_lifecycle_environment", "test_lifecycle_environment", acctest.Required, acctest.Create, OsManagementHubLifecycleEnvironmentRepresentation)

	OsManagementHubProfileIgnoreDefinedTagsRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubProfileResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubProfileResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_os_management_hub_profile.test_profile"
	datasourceName := "data.oci_os_management_hub_profiles.test_profiles"
	singularDatasourceName := "data.oci_os_management_hub_profile.test_profile"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Optional, acctest.Create, OsManagementHubSoftwareSourceProfileRepresentation), "osmanagementhub", "profile", t)
	acctest.SaveConfigContent(config+compartmentIdVariableStr+acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Optional, acctest.Create, OsManagementHubGroupProfileRepresentation), "osmanagementhub", "profile", t)
	acctest.SaveConfigContent(config+compartmentIdVariableStr+acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Optional, acctest.Create, OsManagementHubLifecycleProfileRepresentation), "osmanagementhub", "profile", t)

	acctest.ResourceTest(t, testAccCheckOsManagementHubProfileDestroy, []resource.TestStep{
		// verify Create software source profile
		{
			Config: config + compartmentIdVariableStr + OsManagementHubProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Required, acctest.Create, OsManagementHubSoftwareSourceProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckNoResourceAttr(resourceName, "lifecycle_stage_id"),
				resource.TestCheckNoResourceAttr(resourceName, "managed_instance_group_id"),
				resource.TestCheckResourceAttr(resourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttr(resourceName, "profile_type", "SOFTWARESOURCE"),
				resource.TestCheckResourceAttr(resourceName, "software_source_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vendor_name", "ORACLE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubProfileResourceDependencies,
		},
		// verify Create software source profile with optionals
		{
			Config: config + compartmentIdVariableStr + OsManagementHubProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Optional, acctest.Create, OsManagementHubSoftwareSourceProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckNoResourceAttr(resourceName, "lifecycle_stage_id"),
				resource.TestCheckNoResourceAttr(resourceName, "managed_instance_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_station_id"),
				resource.TestCheckResourceAttr(resourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttr(resourceName, "profile_type", "SOFTWARESOURCE"),
				resource.TestCheckResourceAttr(resourceName, "software_source_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "software_sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vendor_name", "ORACLE"),

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

		// verify updates to updatable parameters in software source profile
		{
			Config: config + compartmentIdVariableStr + OsManagementHubProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Optional, acctest.Update, OsManagementHubSoftwareSourceProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckNoResourceAttr(resourceName, "lifecycle_stage_id"),
				resource.TestCheckNoResourceAttr(resourceName, "managed_instance_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_station_id"),
				resource.TestCheckResourceAttr(resourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttr(resourceName, "profile_type", "SOFTWARESOURCE"),
				resource.TestCheckResourceAttr(resourceName, "software_source_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "software_sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vendor_name", "ORACLE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify software source profile datasource
		{
			Config: config + compartmentIdVariableStr + OsManagementHubProfileResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_profiles", "test_profiles", acctest.Optional, acctest.Update, OsManagementHubSoftwareSourceProfileDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Optional, acctest.Update, OsManagementHubSoftwareSourceProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "displayName"),
				resource.TestCheckResourceAttr(datasourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttrSet(datasourceName, "profile_id"),
				resource.TestCheckResourceAttr(datasourceName, "profile_type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "vendor_name", "ORACLE"),

				resource.TestCheckResourceAttr(datasourceName, "profile_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "profile_collection.0.items.#", "1"),
			),
		},
		// verify singular software source profile datasource
		{
			Config: config + compartmentIdVariableStr + OsManagementHubProfileResourceDependencies + OsManagementHubSoftwareSourceProfileResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Required, acctest.Create, OsManagementHubProfileSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "profile_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttr(singularDatasourceName, "profile_type", "SOFTWARESOURCE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "software_sources.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vendor_name", "ORACLE"),
			),
		},
		// verify software source profile resource import
		{
			Config:            config + OsManagementHubSoftwareSourceProfileRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"lifecycle_stage_id",
				"managed_instance_group_id",
				"software_source_ids",
			},
			ResourceName: resourceName,
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubProfileResourceDependencies,
		},
		// verify Create group profile
		{
			Config: config + compartmentIdVariableStr + OsManagementHubProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Required, acctest.Create, OsManagementHubGroupProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_group_id"),
				resource.TestCheckResourceAttr(resourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttr(resourceName, "profile_type", "GROUP"),
				resource.TestCheckResourceAttr(resourceName, "vendor_name", "ORACLE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubProfileResourceDependencies,
		},
		// verify Create group profile with optionals
		{
			Config: config + compartmentIdVariableStr + OsManagementHubProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Optional, acctest.Create, OsManagementHubGroupProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "managed_instance_group.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_station_id"),
				resource.TestCheckResourceAttr(resourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttr(resourceName, "profile_type", "GROUP"),
				resource.TestCheckResourceAttr(resourceName, "vendor_name", "ORACLE"),

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

		// verify updates to updatable parameters in group profile
		{
			Config: config + compartmentIdVariableStr + OsManagementHubProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Optional, acctest.Update, OsManagementHubGroupProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "managed_instance_group.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "managed_instance_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_station_id"),
				resource.TestCheckResourceAttr(resourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttr(resourceName, "profile_type", "GROUP"),
				resource.TestCheckResourceAttr(resourceName, "vendor_name", "ORACLE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify group profile datasource
		{
			Config: config + compartmentIdVariableStr + OsManagementHubProfileResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_profiles", "test_profiles", acctest.Optional, acctest.Update, OsManagementHubGroupProfileDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Optional, acctest.Update, OsManagementHubGroupProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "displayName"),
				resource.TestCheckResourceAttr(datasourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttrSet(datasourceName, "profile_id"),
				resource.TestCheckResourceAttr(datasourceName, "profile_type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "vendor_name", "ORACLE"),

				resource.TestCheckResourceAttr(datasourceName, "profile_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "profile_collection.0.items.#", "1"),
			),
		},
		// verify singular group profile datasource
		{
			Config: config + compartmentIdVariableStr + OsManagementHubProfileResourceDependencies + OsManagementHubGroupProfileResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Required, acctest.Create, OsManagementHubProfileSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "profile_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "managed_instance_group.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttr(singularDatasourceName, "profile_type", "GROUP"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vendor_name", "ORACLE"),
			),
		},
		// verify group profile resource import
		{
			Config:            config + OsManagementHubGroupProfileRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"lifecycle_stage_id",
				"managed_instance_group_id",
				"software_source_ids",
			},
			ResourceName: resourceName,
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubProfileResourceDependencies,
		},
		// verify Create lifecycle profile
		{
			Config: config + compartmentIdVariableStr + OsManagementHubProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Required, acctest.Create, OsManagementHubLifecycleProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttr(resourceName, "profile_type", "LIFECYCLE"),
				resource.TestCheckResourceAttrSet(resourceName, "lifecycle_stage_id"),
				resource.TestCheckResourceAttr(resourceName, "vendor_name", "ORACLE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OsManagementHubProfileResourceDependencies,
		},
		// verify Create lifecycle profile with optionals
		{
			Config: config + compartmentIdVariableStr + OsManagementHubProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Optional, acctest.Create, OsManagementHubLifecycleProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "lifecycle_stage.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "lifecycle_stage_id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_station_id"),
				resource.TestCheckResourceAttr(resourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttr(resourceName, "profile_type", "LIFECYCLE"),
				resource.TestCheckResourceAttr(resourceName, "vendor_name", "ORACLE"),

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

		// verify updates to updatable parameters in lifecycle profile
		{
			Config: config + compartmentIdVariableStr + OsManagementHubProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Optional, acctest.Update, OsManagementHubLifecycleProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "lifecycle_stage.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "lifecycle_stage_id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_station_id"),
				resource.TestCheckResourceAttr(resourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttr(resourceName, "profile_type", "LIFECYCLE"),
				resource.TestCheckResourceAttr(resourceName, "vendor_name", "ORACLE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify lifecycle profile datasource
		{
			Config: config + compartmentIdVariableStr + OsManagementHubProfileResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_profiles", "test_profiles", acctest.Optional, acctest.Update, OsManagementHubLifecycleProfileDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Optional, acctest.Update, OsManagementHubLifecycleProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "displayName"),
				resource.TestCheckResourceAttr(datasourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttrSet(datasourceName, "profile_id"),
				resource.TestCheckResourceAttr(datasourceName, "profile_type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "vendor_name", "ORACLE"),

				resource.TestCheckResourceAttr(datasourceName, "profile_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "profile_collection.0.items.#", "1"),
			),
		},
		// verify singular lifecycle profile datasource
		{
			Config: config + compartmentIdVariableStr + OsManagementHubProfileResourceDependencies + OsManagementHubLifecycleProfileResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_profile", "test_profile", acctest.Required, acctest.Create, OsManagementHubProfileSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "profile_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "arch_type", "X86_64"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lifecycle_environment.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lifecycle_stage.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "os_family", "ORACLE_LINUX_8"),
				resource.TestCheckResourceAttr(singularDatasourceName, "profile_type", "LIFECYCLE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vendor_name", "ORACLE"),
			),
		},
		// verify software source profile resource import
		{
			Config:            config + OsManagementHubLifecycleProfileRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"lifecycle_stage_id",
				"managed_instance_group_id",
				"software_source_ids",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckOsManagementHubProfileDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OnboardingClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_os_management_hub_profile" {
			noResourceFound = false
			request := oci_os_management_hub.GetProfileRequest{}

			tmp := rs.Primary.ID
			request.ProfileId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "os_management_hub")

			response, err := client.GetProfile(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_os_management_hub.ProfileLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
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
	if !acctest.InSweeperExcludeList("OsManagementHubProfile") {
		resource.AddTestSweepers("OsManagementHubProfile", &resource.Sweeper{
			Name:         "OsManagementHubProfile",
			Dependencies: acctest.DependencyGraph["profile"],
			F:            sweepOsManagementHubProfileResource,
		})
	}
}

func sweepOsManagementHubProfileResource(compartment string) error {
	onboardingClient := acctest.GetTestClients(&schema.ResourceData{}).OnboardingClient()
	profileIds, err := getOsManagementHubProfileIds(compartment)
	if err != nil {
		return err
	}
	for _, profileId := range profileIds {
		if ok := acctest.SweeperDefaultResourceId[profileId]; !ok {
			deleteProfileRequest := oci_os_management_hub.DeleteProfileRequest{}

			deleteProfileRequest.ProfileId = &profileId

			deleteProfileRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "os_management_hub")
			_, error := onboardingClient.DeleteProfile(context.Background(), deleteProfileRequest)
			if error != nil {
				fmt.Printf("Error deleting Profile %s %s, It is possible that the resource is already deleted. Please verify manually \n", profileId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &profileId, OsManagementHubProfileSweepWaitCondition, time.Duration(3*time.Minute),
				OsManagementHubProfileSweepResponseFetchOperation, "os_management_hub", true)
		}
	}
	return nil
}

func getOsManagementHubProfileIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ProfileId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	onboardingClient := acctest.GetTestClients(&schema.ResourceData{}).OnboardingClient()

	listProfilesRequest := oci_os_management_hub.ListProfilesRequest{}
	listProfilesRequest.CompartmentId = &compartmentId
	listProfilesRequest.LifecycleState = oci_os_management_hub.ProfileLifecycleStateActive
	listProfilesResponse, err := onboardingClient.ListProfiles(context.Background(), listProfilesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Profile list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, profile := range listProfilesResponse.Items {
		id := *profile.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ProfileId", id)
	}
	return resourceIds, nil
}

func OsManagementHubProfileSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if profileResponse, ok := response.Response.(oci_os_management_hub.GetProfileResponse); ok {
		return profileResponse.GetLifecycleState() != oci_os_management_hub.ProfileLifecycleStateDeleted
	}
	return false
}

func OsManagementHubProfileSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OnboardingClient().GetProfile(context.Background(), oci_os_management_hub.GetProfileRequest{
		ProfileId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
