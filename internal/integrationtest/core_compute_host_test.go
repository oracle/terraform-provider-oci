// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	CoreComputeHostResourceWithRequired = acctest.GenerateResourceFromRepresentationMap("oci_core_compute_host", "test_compute_host", acctest.Required, acctest.Create, CoreComputeHostRepresentation)

	CoreComputeHostResourceWithOptional = acctest.GenerateResourceFromRepresentationMap("oci_core_compute_host", "test_compute_host", acctest.Optional, acctest.Update, CoreComputeHostRepresentation)

	CoreComputeHostSingularDataSourceRepresentation = map[string]interface{}{
		"compute_host_id": acctest.Representation{RepType: acctest.Required, Create: `${var.baremetalhost_id}`},
	}

	CoreComputeHostDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain":          acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compute_host_group_id":        acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_compute_host_group.test_compute_host_group.id}`},
		"compute_host_health":          acctest.Representation{RepType: acctest.Optional, Create: `computeHostHealth`},
		"compute_host_lifecycle_state": acctest.Representation{RepType: acctest.Optional, Create: `computeHostLifecycleState`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"network_resource_id":          acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_guard_resource.test_resource.id}`},
		"filter":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreComputeHostDataSourceFilterRepresentation}}
	CoreComputeHostDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_compute_host.test_compute_host.id}`}},
	}

	CoreComputeHostDataSourceRepresentationWithNoFilters = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compute_host_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	CoreComputeHostRepresentation = map[string]interface{}{
		"compute_host_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.baremetalhost_id}`},
		"compute_host_group_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_compute_host_group.test_compute_host_group.id}`},
	}

	CoreComputeHostResourceWithOptionalSubCompartment = acctest.GenerateResourceFromRepresentationMap("oci_core_compute_host", "test_compute_host_list2", acctest.Optional, acctest.Update, CoreComputeHostRepresentationSubCompartment)

	CoreComputeHostRepresentationSubCompartment = map[string]interface{}{
		"compute_host_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.baremetalhost_id2}`},
		"compute_host_group_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_compute_host_group.test_compute_host_group.id}`},
	}

	CoreComputeHostResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_resources", "test_resources", acctest.Required, acctest.Create, CloudGuardResourceDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_compute_host_group", "test_compute_host_group", acctest.Required, acctest.Create, CoreComputeHostGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_compute_host_group", "test_compute_host_group2", acctest.Required, acctest.Create, CoreComputeHostGroupRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies

	CloudGuardResourceDataSourceRepresentation = map[string]interface{}{}
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreComputeHostResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreComputeHostResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig() + `
	variable "baremetalhost_id" {
		default = "` + utils.GetEnvSettingWithBlankDefault("baremetalhost_id") + `"
	}
	`

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_compute_hosts.test_compute_hosts"
	singularDatasourceName := "data.oci_core_compute_host.test_compute_host"

	var resId, resId2 string

	resourceName := "oci_core_compute_host.test_compute_host"
	hostGroup1ResourceName := "oci_core_compute_host_group.test_compute_host_group"
	hostGroup2ResourceName := "oci_core_compute_host_group.test_compute_host_group2"

	step1Config := config + compartmentIdVariableStr + CoreComputeHostResourceDependencies + CoreComputeHostResourceWithRequired

	acctest.SaveConfigContent(step1Config, "core", "computeHost", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: step1Config,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compute_host_id"),
				resource.TestCheckNoResourceAttr(resourceName, "compute_host_group_id"),

				func(s *terraform.State) (err error) {
					//resId, err = acctest.FromInstanceState(s, resourceName, "id")
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreComputeHostResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreComputeHostResourceDependencies + CoreComputeHostResourceWithOptional,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_host_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_host_group_id"),
				acctest.TestCheckResourceAttributesEqual(resourceName, "compute_host_group_id", hostGroup1ResourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "newName"),
				resource.TestCheckResourceAttrSet(resourceName, "fault_domain"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "health"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "shape"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + CoreComputeHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_host", "test_compute_host", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(CoreComputeHostRepresentation, map[string]interface{}{
						"compute_host_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_compute_host_group.test_compute_host_group2.id}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compute_host_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_host_group_id"),
				acctest.TestCheckResourceAttributesEqual(resourceName, "compute_host_group_id", hostGroup2ResourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_host_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "fault_domain"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "health"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "shape"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreComputeHostResourceDependencies,
		},
		// verify datasource
		{
			Config: config + CoreComputeHostResourceDependencies + CoreComputeHostResourceWithOptional +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_hosts", "test_compute_hosts", acctest.Required, acctest.Create, CoreComputeHostDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "compute_host_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "compute_host_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + CoreComputeHostResourceDependencies + CoreComputeHostResourceWithOptional +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_host", "test_compute_host", acctest.Required, acctest.Create, CoreComputeHostSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_host_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration_data.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "configuration_state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "computebaremetalhost-ayheq"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fault_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "health"),
				resource.TestCheckNoResourceAttr(singularDatasourceName, "hpc_island_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckNoResourceAttr(singularDatasourceName, "instance_id"),
				resource.TestCheckNoResourceAttr(singularDatasourceName, "local_block_id"),
				resource.TestCheckNoResourceAttr(singularDatasourceName, "network_block_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "recycle_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckNoResourceAttr(singularDatasourceName, "time_configuration_check"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + CoreComputeHostResourceWithRequired,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreComputeHostResource_listCompartments(t *testing.T) {
	httpreplay.SetScenario("TestCoreComputeHostResource_listCompartments")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig() + `
	variable "baremetalhost_id" {
		default = "` + utils.GetEnvSettingWithBlankDefault("baremetalhost_id") + `"
	}
	variable "baremetalhost_id2" {
		default = "` + utils.GetEnvSettingWithBlankDefault("baremetalhost_id2") + `"
	}
	`
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_compute_hosts.test_compute_hosts"
	singularDatasourceName := "data.oci_core_compute_host.test_compute_host"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{

		// verify datasource
		// two hosts in two different compartments but with boolean false
		{
			Config: config + CoreComputeHostResourceDependencies + CoreComputeHostResourceWithRequired + CoreComputeHostResourceWithOptionalSubCompartment +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_hosts", "test_compute_hosts", acctest.Required, acctest.Create,
					CoreComputeHostDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "compute_host_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "compute_host_collection.0.items.#", "1"),
			),
		},

		// verify datasource
		// two hosts in two different compartments but with boolean true
		{
			Config: config + CoreComputeHostResourceDependencies + CoreComputeHostResourceWithRequired + CoreComputeHostResourceWithOptionalSubCompartment +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_hosts", "test_compute_hosts", acctest.Optional, acctest.Create,
					CoreComputeHostDataSourceRepresentationWithNoFilters) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "compute_host_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "compute_host_collection.0.items.#", "2"),
			),
		},

		// verify singular datasource
		{
			Config: config + CoreComputeHostResourceDependencies + CoreComputeHostResourceWithOptional + CoreComputeHostResourceWithOptionalSubCompartment +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_host", "test_compute_host", acctest.Required, acctest.Create, CoreComputeHostSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_host_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fault_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "health"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "shape"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
