// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	DedicatedVmHostResourceConfig_E3Shape = DedicatedVmHostResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Update, dedicatedVmHostRepresentation_E3Shape)

	DedicatedVmHostResourceConfig_E2Shape = DedicatedVmHostResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Update, dedicatedVmHostRepresentation_E2Shape)

	DedicatedVmHostResourceConfig_DenseIO2Shape = DedicatedVmHostResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Update, dedicatedVmHostRepresentation_DenseIO2Shape)

	dedicatedVmHostDataSourceRepresentation_E3Shape = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"instance_shape_name": acctest.Representation{RepType: acctest.Optional, Create: `VM.Standard.E3.Flex`},
		"remaining_memory_in_gbs_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `16.0`},
		"remaining_ocpus_greater_than_or_equal_to":         acctest.Representation{RepType: acctest.Optional, Create: `1.0`},
		"state":  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: dedicatedVmHostDataSourceFilterRepresentation}}
	dedicatedVmHostDataSourceFilterRepresentation_E3Shape = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`}},
	}

	dedicatedVmHostDataSourceRepresentation_E2Shape = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"instance_shape_name": acctest.Representation{RepType: acctest.Optional, Create: `VM.Standard.E2.1`},
		"remaining_memory_in_gbs_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `8.0`},
		"remaining_ocpus_greater_than_or_equal_to":         acctest.Representation{RepType: acctest.Optional, Create: `1.0`},
		"state":  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: dedicatedVmHostDataSourceFilterRepresentation}}
	dedicatedVmHostDataSourceFilterRepresentation_E2Shape = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`}},
	}

	dedicatedVmHostDataSourceRepresentation_DenseIO2Shape = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"instance_shape_name": acctest.Representation{RepType: acctest.Optional, Create: `VM.DenseIO2.8`},
		"remaining_memory_in_gbs_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `120.0`},
		"remaining_ocpus_greater_than_or_equal_to":         acctest.Representation{RepType: acctest.Optional, Create: `8.0`},
		"state":  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: dedicatedVmHostDataSourceFilterRepresentation}}
	dedicatedVmHostDataSourceFilterRepresentation_DenseIO2Shape = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`}},
	}

	dedicatedVmHostRepresentation_E3Shape = map[string]interface{}{
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"dedicated_vm_host_shape": acctest.Representation{RepType: acctest.Required, Create: `DVH.Standard.E3.128`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"fault_domain":            acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-3`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	dedicatedVmHostRepresentation_E2Shape = map[string]interface{}{
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"dedicated_vm_host_shape": acctest.Representation{RepType: acctest.Required, Create: `DVH.Standard.E2.64`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"fault_domain":            acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-3`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	dedicatedVmHostRepresentation_DenseIO2Shape = map[string]interface{}{
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"dedicated_vm_host_shape": acctest.Representation{RepType: acctest.Required, Create: `DVH.DenseIO2.52`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"fault_domain":            acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-3`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
)

// issue-routing-tag: core/default
func TestResourceCoreDedicatedVmHost_DenseIO2ShapeDVH(t *testing.T) {
	httpreplay.SetScenario("TestResourceCoreDedicatedVmHost_DenseIO2ShapeDVH")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_dedicated_vm_host.test_dedicated_vm_host"
	datasourceName := "data.oci_core_dedicated_vm_hosts.test_dedicated_vm_hosts"
	singularDatasourceName := "data.oci_core_dedicated_vm_host.test_dedicated_vm_host"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DedicatedVmHostResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Create, dedicatedVmHostRepresentation_DenseIO2Shape), "core", "dedicatedVmHost", t)

	acctest.ResourceTest(t, testAccCheckCoreDedicatedVmHostDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DedicatedVmHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Required, acctest.Create, dedicatedVmHostRepresentation_DenseIO2Shape),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dedicated_vm_host_shape", "DVH.DenseIO2.52"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DedicatedVmHostResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DedicatedVmHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Create, dedicatedVmHostRepresentation_DenseIO2Shape),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dedicated_vm_host_shape", "DVH.DenseIO2.52"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-3"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "remaining_ocpus"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "total_ocpus"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DedicatedVmHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(dedicatedVmHostRepresentation_DenseIO2Shape, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "dedicated_vm_host_shape", "DVH.DenseIO2.52"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-3"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "remaining_ocpus"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "total_ocpus"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DedicatedVmHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Update, dedicatedVmHostRepresentation_DenseIO2Shape),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dedicated_vm_host_shape", "DVH.DenseIO2.52"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-3"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "remaining_ocpus"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "total_ocpus"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_dedicated_vm_hosts", "test_dedicated_vm_hosts", acctest.Optional, acctest.Update, dedicatedVmHostDataSourceRepresentation_DenseIO2Shape) +
				compartmentIdVariableStr + DedicatedVmHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Update, dedicatedVmHostRepresentation_DenseIO2Shape),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "instance_shape_name", "VM.DenseIO2.8"),
				resource.TestCheckResourceAttr(datasourceName, "remaining_memory_in_gbs_greater_than_or_equal_to", "120"),
				resource.TestCheckResourceAttr(datasourceName, "remaining_ocpus_greater_than_or_equal_to", "8"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.0.dedicated_vm_host_shape", "DVH.DenseIO2.52"),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.0.fault_domain", "FAULT-DOMAIN-3"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.remaining_memory_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.remaining_ocpus"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.total_memory_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.total_ocpus"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Required, acctest.Create, dedicatedVmHostSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DedicatedVmHostResourceConfig_DenseIO2Shape,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "dedicated_vm_host_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "dedicated_vm_host_shape", "DVH.DenseIO2.52"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fault_domain", "FAULT-DOMAIN-3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "remaining_memory_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "remaining_ocpus"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_memory_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_ocpus"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DedicatedVmHostResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

// issue-routing-tag: core/default
func TestResourceCoreDedicatedVmHost_E2ShapeDVH(t *testing.T) {
	httpreplay.SetScenario("TestCoreDedicatedVmHostResource_E2ShapeDVH")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_dedicated_vm_host.test_dedicated_vm_host"
	datasourceName := "data.oci_core_dedicated_vm_hosts.test_dedicated_vm_hosts"
	singularDatasourceName := "data.oci_core_dedicated_vm_host.test_dedicated_vm_host"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DedicatedVmHostResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Create, dedicatedVmHostRepresentation_E2Shape), "core", "dedicatedVmHost", t)

	acctest.ResourceTest(t, testAccCheckCoreDedicatedVmHostDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DedicatedVmHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Required, acctest.Create, dedicatedVmHostRepresentation_E2Shape),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dedicated_vm_host_shape", "DVH.Standard.E2.64"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DedicatedVmHostResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DedicatedVmHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Create, dedicatedVmHostRepresentation_E2Shape),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dedicated_vm_host_shape", "DVH.Standard.E2.64"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-3"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "remaining_ocpus"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "total_ocpus"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DedicatedVmHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(dedicatedVmHostRepresentation_E2Shape, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "dedicated_vm_host_shape", "DVH.Standard.E2.64"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-3"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "remaining_ocpus"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "total_ocpus"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DedicatedVmHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Update, dedicatedVmHostRepresentation_E2Shape),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dedicated_vm_host_shape", "DVH.Standard.E2.64"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-3"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "remaining_ocpus"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "total_ocpus"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_dedicated_vm_hosts", "test_dedicated_vm_hosts", acctest.Optional, acctest.Update, dedicatedVmHostDataSourceRepresentation_E2Shape) +
				compartmentIdVariableStr + DedicatedVmHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Update, dedicatedVmHostRepresentation_E2Shape),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "instance_shape_name", "VM.Standard.E2.1"),
				resource.TestCheckResourceAttr(datasourceName, "remaining_memory_in_gbs_greater_than_or_equal_to", "8"),
				resource.TestCheckResourceAttr(datasourceName, "remaining_ocpus_greater_than_or_equal_to", "1"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.0.dedicated_vm_host_shape", "DVH.Standard.E2.64"),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.0.fault_domain", "FAULT-DOMAIN-3"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.remaining_memory_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.remaining_ocpus"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.total_memory_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.total_ocpus"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Required, acctest.Create, dedicatedVmHostSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DedicatedVmHostResourceConfig_E2Shape,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "dedicated_vm_host_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "dedicated_vm_host_shape", "DVH.Standard.E2.64"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fault_domain", "FAULT-DOMAIN-3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "remaining_memory_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "remaining_ocpus"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_memory_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_ocpus"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DedicatedVmHostResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

// issue-routing-tag: core/default
func TestResourceCoreDedicatedVmHost_E3ShapeDVH(t *testing.T) {
	httpreplay.SetScenario("TestCoreDedicatedVmHostResource_E3ShapeDVH")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_dedicated_vm_host.test_dedicated_vm_host"
	datasourceName := "data.oci_core_dedicated_vm_hosts.test_dedicated_vm_hosts"
	singularDatasourceName := "data.oci_core_dedicated_vm_host.test_dedicated_vm_host"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DedicatedVmHostResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Create, dedicatedVmHostRepresentation_E3Shape), "core", "dedicatedVmHost", t)

	acctest.ResourceTest(t, testAccCheckCoreDedicatedVmHostDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DedicatedVmHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Required, acctest.Create, dedicatedVmHostRepresentation_E3Shape),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dedicated_vm_host_shape", "DVH.Standard.E3.128"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DedicatedVmHostResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DedicatedVmHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Create, dedicatedVmHostRepresentation_E3Shape),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dedicated_vm_host_shape", "DVH.Standard.E3.128"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-3"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "remaining_ocpus"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "total_ocpus"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DedicatedVmHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(dedicatedVmHostRepresentation_E3Shape, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "dedicated_vm_host_shape", "DVH.Standard.E3.128"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-3"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "remaining_ocpus"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "total_ocpus"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DedicatedVmHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Update, dedicatedVmHostRepresentation_E3Shape),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dedicated_vm_host_shape", "DVH.Standard.E3.128"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "fault_domain", "FAULT-DOMAIN-3"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "remaining_ocpus"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "total_ocpus"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_dedicated_vm_hosts", "test_dedicated_vm_hosts", acctest.Optional, acctest.Update, dedicatedVmHostDataSourceRepresentation_E3Shape) +
				compartmentIdVariableStr + DedicatedVmHostResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Optional, acctest.Update, dedicatedVmHostRepresentation_E3Shape),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "instance_shape_name", "VM.Standard.E3.Flex"),
				resource.TestCheckResourceAttr(datasourceName, "remaining_memory_in_gbs_greater_than_or_equal_to", "16"),
				resource.TestCheckResourceAttr(datasourceName, "remaining_ocpus_greater_than_or_equal_to", "1"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.0.dedicated_vm_host_shape", "DVH.Standard.E3.128"),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_hosts.0.fault_domain", "FAULT-DOMAIN-3"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.remaining_memory_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.remaining_ocpus"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.total_memory_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_hosts.0.total_ocpus"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_dedicated_vm_host", "test_dedicated_vm_host", acctest.Required, acctest.Create, dedicatedVmHostSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DedicatedVmHostResourceConfig_E3Shape,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "dedicated_vm_host_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "dedicated_vm_host_shape", "DVH.Standard.E3.128"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fault_domain", "FAULT-DOMAIN-3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "remaining_memory_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "remaining_ocpus"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_memory_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_ocpus"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DedicatedVmHostResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
