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
	OsManagementHubSoftwareSourceChangeAvailabilityManagementRequiredOnlyResource = OsManagementHubSoftwareSourceChangeAvailabilityManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source_change_availability_management", "test_software_source_change_availability_management", acctest.Required, acctest.Create, OsManagementHubSoftwareSourceChangeAvailabilityManagementRepresentation)

	OsManagementHubSoftwareSourceChangeAvailabilityManagementRepresentation = map[string]interface{}{
		"software_source_availabilities": acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubSoftwareSourceChangeAvailabilityManagementSoftwareSourceAvailabilitiesRepresentation},
	}
	OsManagementHubSoftwareSourceChangeAvailabilityManagementSoftwareSourceAvailabilitiesRepresentation = map[string]interface{}{
		"software_source_id":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_os_management_hub_software_sources.test_software_sources.software_source_collection[0].items[0].id}`},
		"availability":        acctest.Representation{RepType: acctest.Required, Create: `AVAILABLE`},
		"availability_at_oci": acctest.Representation{RepType: acctest.Required, Create: `AVAILABLE`},
	}

	OsManagementHubSoftwareSourceChangeAvailabilityManagementRepresentation2 = map[string]interface{}{
		"software_source_availabilities": acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubSoftwareSourceChangeAvailabilityManagementSoftwareSourceAvailabilitiesRepresentation2},
	}
	OsManagementHubSoftwareSourceChangeAvailabilityManagementSoftwareSourceAvailabilitiesRepresentation2 = map[string]interface{}{
		"software_source_id":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_os_management_hub_software_sources.test_software_sources.software_source_collection[0].items[0].id}`},
		"availability":        acctest.Representation{RepType: acctest.Required, Create: `SELECTED`},
		"availability_at_oci": acctest.Representation{RepType: acctest.Required, Create: `SELECTED`},
	}

	OsManagementHubSoftwareSourceChangeAvailabilityManagementRepresentation3 = map[string]interface{}{
		"software_source_availabilities": acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubSoftwareSourceChangeAvailabilityManagementSoftwareSourceAvailabilitiesRepresentation3},
	}
	OsManagementHubSoftwareSourceChangeAvailabilityManagementSoftwareSourceAvailabilitiesRepresentation3 = map[string]interface{}{
		"software_source_id":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_os_management_hub_software_sources.test_software_sources.software_source_collection[0].items[0].id}`},
		"availability":        acctest.Representation{RepType: acctest.Required, Create: `SELECTED`},
		"availability_at_oci": acctest.Representation{RepType: acctest.Required, Create: `AVAILABLE`},
	}

	OsManagementHubSoftwareSourceChangeAvailabilityManagementRepresentation4 = map[string]interface{}{
		"software_source_availabilities": acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubSoftwareSourceChangeAvailabilityManagementSoftwareSourceAvailabilitiesRepresentation4},
	}
	OsManagementHubSoftwareSourceChangeAvailabilityManagementSoftwareSourceAvailabilitiesRepresentation4 = map[string]interface{}{
		"software_source_id":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_os_management_hub_software_sources.test_software_sources.software_source_collection[0].items[0].id}`},
		"availability":        acctest.Representation{RepType: acctest.Required, Create: `AVAILABLE`},
		"availability_at_oci": acctest.Representation{RepType: acctest.Required, Create: `SELECTED`},
	}

	OsManagementHubSoftwareSourceChangeAvailabilityManagementRepresentation5 = map[string]interface{}{
		"software_source_availabilities": acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubSoftwareSourceChangeAvailabilityManagementSoftwareSourceAvailabilitiesRepresentation5},
	}
	OsManagementHubSoftwareSourceChangeAvailabilityManagementSoftwareSourceAvailabilitiesRepresentation5 = map[string]interface{}{
		"software_source_id":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_os_management_hub_software_sources.test_software_sources.software_source_collection[0].items[0].id}`},
		"availability":        acctest.Representation{RepType: acctest.Required, Create: `AVAILABLE`},
		"availability_at_oci": acctest.Representation{RepType: acctest.Required, Create: `AVAILABLE`},
	}

	OsManagementHubSoftwareSourceChangeAvailbalityRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":         acctest.Representation{RepType: acctest.Required, Create: `${var.software_source_name}`, Update: `${var.software_source_name}`},
		"software_source_type": acctest.Representation{RepType: acctest.Required, Create: []string{`VENDOR`}},
	}

	OsManagementHubSoftwareSourceChangeAvailabilityManagementResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_software_sources", "test_software_sources", acctest.Required, acctest.Create, OsManagementHubSoftwareSourceChangeAvailbalityRepresentation)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubSoftwareSourceChangeAvailabilityManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubSoftwareSourceChangeAvailabilityManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	softwareSourceName := utils.GetEnvSettingWithBlankDefault("software_source_name")
	softwareSourceNameVariableStr := fmt.Sprintf("variable \"software_source_name\" { default = \"%s\" }\n", softwareSourceName)

	resourceName := "oci_os_management_hub_software_source_change_availability_management.test_software_source_change_availability_management"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+softwareSourceNameVariableStr+OsManagementHubSoftwareSourceChangeAvailabilityManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source_change_availability_management", "test_software_source_change_availability_management", acctest.Optional, acctest.Create, OsManagementHubSoftwareSourceChangeAvailabilityManagementRepresentation), "osmanagementhub", "softwareSourceChangeAvailabilityManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + softwareSourceNameVariableStr + OsManagementHubSoftwareSourceChangeAvailabilityManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source_change_availability_management", "test_software_source_change_availability_management", acctest.Required, acctest.Create, OsManagementHubSoftwareSourceChangeAvailabilityManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(resource.TestCheckResourceAttr(resourceName, "software_source_availabilities.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "software_source_availabilities.0.availability", "AVAILABLE"),
				resource.TestCheckResourceAttr(resourceName, "software_source_availabilities.0.availability_at_oci", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(resourceName, "software_source_availabilities.0.software_source_id")),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + softwareSourceNameVariableStr + OsManagementHubSoftwareSourceChangeAvailabilityManagementResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + softwareSourceNameVariableStr + OsManagementHubSoftwareSourceChangeAvailabilityManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source_change_availability_management", "test_software_source_change_availability_management", acctest.Optional, acctest.Create, OsManagementHubSoftwareSourceChangeAvailabilityManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "software_source_availabilities.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "software_source_availabilities.0.availability", "AVAILABLE"),
				resource.TestCheckResourceAttr(resourceName, "software_source_availabilities.0.availability_at_oci", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(resourceName, "software_source_availabilities.0.software_source_id"),
			),
		},
		// verify Create
		{
			Config: config + compartmentIdVariableStr + softwareSourceNameVariableStr + OsManagementHubSoftwareSourceChangeAvailabilityManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source_change_availability_management", "test_software_source_change_availability_management", acctest.Required, acctest.Create, OsManagementHubSoftwareSourceChangeAvailabilityManagementRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "software_source_availabilities.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "software_source_availabilities.0.availability", "SELECTED"),
				resource.TestCheckResourceAttr(resourceName, "software_source_availabilities.0.availability_at_oci", "SELECTED"),
				resource.TestCheckResourceAttrSet(resourceName, "software_source_availabilities.0.software_source_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + softwareSourceNameVariableStr + OsManagementHubSoftwareSourceChangeAvailabilityManagementResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + softwareSourceNameVariableStr + OsManagementHubSoftwareSourceChangeAvailabilityManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source_change_availability_management", "test_software_source_change_availability_management", acctest.Optional, acctest.Create, OsManagementHubSoftwareSourceChangeAvailabilityManagementRepresentation3),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "software_source_availabilities.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "software_source_availabilities.0.availability", "SELECTED"),
				resource.TestCheckResourceAttr(resourceName, "software_source_availabilities.0.availability_at_oci", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(resourceName, "software_source_availabilities.0.software_source_id"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + softwareSourceNameVariableStr + OsManagementHubSoftwareSourceChangeAvailabilityManagementResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + softwareSourceNameVariableStr + OsManagementHubSoftwareSourceChangeAvailabilityManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source_change_availability_management", "test_software_source_change_availability_management", acctest.Optional, acctest.Create, OsManagementHubSoftwareSourceChangeAvailabilityManagementRepresentation4),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "software_source_availabilities.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "software_source_availabilities.0.availability", "AVAILABLE"),
				resource.TestCheckResourceAttr(resourceName, "software_source_availabilities.0.availability_at_oci", "SELECTED"),
				resource.TestCheckResourceAttrSet(resourceName, "software_source_availabilities.0.software_source_id"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + softwareSourceNameVariableStr + OsManagementHubSoftwareSourceChangeAvailabilityManagementResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + softwareSourceNameVariableStr + OsManagementHubSoftwareSourceChangeAvailabilityManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_software_source_change_availability_management", "test_software_source_change_availability_management", acctest.Optional, acctest.Create, OsManagementHubSoftwareSourceChangeAvailabilityManagementRepresentation5),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "software_source_availabilities.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "software_source_availabilities.0.availability", "AVAILABLE"),
				resource.TestCheckResourceAttr(resourceName, "software_source_availabilities.0.availability_at_oci", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(resourceName, "software_source_availabilities.0.software_source_id"),
			),
		},
	})
}
