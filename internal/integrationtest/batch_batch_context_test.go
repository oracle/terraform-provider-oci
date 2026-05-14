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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	oci_batch "github.com/oracle/oci-go-sdk/v65/batch"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	BatchBatchContextDisplayName       = "batchcontextgpu" + time.Now().UTC().Format("20060102150405")
	BatchBatchContextDisplayNameUpdate = BatchBatchContextDisplayName + "2"
	BatchBatchContextCpuDisplayName    = "batchcontextcpu" + time.Now().UTC().Format("20060102150405")

	BatchBatchContextRequiredOnlyResource = BatchBatchContextResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_context", "test_batch_context", acctest.Required, acctest.Create, BatchBatchContextRepresentation)

	BatchBatchContextResourceConfig = BatchBatchContextOptionalResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_context", "test_batch_context", acctest.Optional, acctest.Update, BatchBatchContextRepresentation)

	BatchBatchContextSingularDataSourceRepresentation = map[string]interface{}{
		"batch_context_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_batch_batch_context.test_batch_context.id}`},
	}

	BatchBatchContextDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: BatchBatchContextDisplayName, Update: BatchBatchContextDisplayNameUpdate},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_batch_batch_context.test_batch_context.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: BatchBatchContextDataSourceFilterRepresentation}}
	BatchBatchContextDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_batch_batch_context.test_batch_context.id}`}},
	}

	BatchBatchContextRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"fleets":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: BatchBatchContextFleetsRepresentation},
		"network":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: BatchBatchContextNetworkRepresentation},
		"description":                 acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                acctest.Representation{RepType: acctest.Optional, Create: BatchBatchContextDisplayName, Update: BatchBatchContextDisplayNameUpdate},
		"entitlements":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"entitlements": "entitlements"}, Update: map[string]string{"entitlements2": "entitlements2"}},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"job_priority_configurations": acctest.RepresentationGroup{RepType: acctest.Optional, Group: BatchBatchContextJobPriorityConfigurationsRepresentation},
		"logging_configuration":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: BatchBatchContextLoggingConfigurationRepresentation},
		"state":                       acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`, Update: `ACTIVE`},
	}
	BatchBatchContextFleetsRepresentation = map[string]interface{}{
		"max_concurrent_tasks": acctest.Representation{RepType: acctest.Required, Create: `10`},
		"name":                 acctest.Representation{RepType: acctest.Required, Create: `name`},
		"shape":                acctest.RepresentationGroup{RepType: acctest.Required, Group: BatchBatchContextFleetsShapeRepresentation},
		"type":                 acctest.Representation{RepType: acctest.Required, Create: `SERVICE_MANAGED_GPU_FLEET`},
	}
	BatchBatchContextNetworkRepresentation = map[string]interface{}{
		"subnet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"nsg_ids":   acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
	}
	BatchBatchContextJobPriorityConfigurationsRepresentation = map[string]interface{}{
		"tag_key":       acctest.Representation{RepType: acctest.Required, Create: `tagKey`, Update: `tagKey2`},
		"tag_namespace": acctest.Representation{RepType: acctest.Required, Create: `tagNamespace`, Update: `tagNamespace2`},
		"values":        acctest.Representation{RepType: acctest.Required, Create: map[string]string{"values": "values"}, Update: map[string]string{"values2": "values2"}},
		"weight":        acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	}
	BatchBatchContextLoggingConfigurationRepresentation = map[string]interface{}{
		"log_group_id":                           acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log_group.test_log_group.id}`},
		"log_id":                                 acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log.test_log.id}`},
		"type":                                   acctest.Representation{RepType: acctest.Required, Create: `OCI_LOGGING`},
		"is_job_task_events_propagation_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	BatchBatchContextFleetsShapeRepresentation = map[string]interface{}{
		"memory_in_gbs":    acctest.Representation{RepType: acctest.Required, Create: `240`},
		"ocpus":            acctest.Representation{RepType: acctest.Required, Create: `15`},
		"type":             acctest.Representation{RepType: acctest.Required, Create: `FIXED_GPU_FLEET_SHAPE`},
		"disk_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `10`},
		"shape_name":       acctest.Representation{RepType: acctest.Required, Create: `VM.GPU.A10.1`},
	}

	BatchBatchContextResourceDependencies = BatchBatchContextShapeResourceConfig +
		acctest.GenerateDataSourceFromRepresentationMap("oci_batch_batch_context_shapes", "test_batch_context_shapes", acctest.Required, acctest.Create, BatchBatchContextShapeGpuDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
	BatchBatchContextLoggingResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", acctest.Required, acctest.Create, DevopsLogGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Required, acctest.Create, customLogRepresentation)
	BatchBatchContextOptionalResourceDependencies = BatchBatchContextResourceDependencies +
		BatchBatchContextLoggingResourceDependencies

	BatchBatchContextCpuRequiredOnlyResource = BatchBatchContextCpuResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_context", "test_batch_context_cpu", acctest.Required, acctest.Create, BatchBatchContextCpuRepresentation)

	BatchBatchContextCpuSingularDataSourceRepresentation = map[string]interface{}{
		"batch_context_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_batch_batch_context.test_batch_context_cpu.id}`},
	}

	BatchBatchContextCpuDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_batch_batch_context.test_batch_context_cpu.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: BatchBatchContextCpuDataSourceFilterRepresentation}}
	BatchBatchContextCpuDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_batch_batch_context.test_batch_context_cpu.id}`}},
	}

	BatchBatchContextCpuRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"fleets":         acctest.RepresentationGroup{RepType: acctest.Required, Group: BatchBatchContextCpuFleetsRepresentation},
		"network":        acctest.RepresentationGroup{RepType: acctest.Required, Group: BatchBatchContextCpuNetworkRepresentation},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: BatchBatchContextCpuDisplayName},
	}
	BatchBatchContextCpuFleetsRepresentation = map[string]interface{}{
		"max_concurrent_tasks": acctest.Representation{RepType: acctest.Required, Create: `1`},
		"name":                 acctest.Representation{RepType: acctest.Required, Create: `cpu-fleet`},
		"shape":                acctest.RepresentationGroup{RepType: acctest.Required, Group: BatchBatchContextCpuFleetsShapeRepresentation},
		"type":                 acctest.Representation{RepType: acctest.Required, Create: `SERVICE_MANAGED_FLEET`},
	}
	BatchBatchContextCpuNetworkRepresentation = map[string]interface{}{
		"subnet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
	}
	BatchBatchContextCpuFleetsShapeRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `16`},
		"ocpus":         acctest.Representation{RepType: acctest.Required, Create: `1`},
		"shape_name":    acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E5.Flex`},
	}

	BatchBatchContextCpuResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
)

// issue-routing-tag: batch/default
func TestBatchBatchContextResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBatchBatchContextResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_batch_batch_context.test_batch_context"
	datasourceName := "data.oci_batch_batch_contexts.test_batch_contexts"
	singularDatasourceName := "data.oci_batch_batch_context.test_batch_context"
	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BatchBatchContextOptionalResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_context", "test_batch_context", acctest.Optional, acctest.Create, BatchBatchContextRepresentation), "batch", "batchContext", t)

	acctest.ResourceTest(t, testAccCheckBatchBatchContextDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BatchBatchContextResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_context", "test_batch_context", acctest.Required, acctest.Create, BatchBatchContextRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "fleets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.max_concurrent_tasks", "10"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.memory_in_gbs", "240"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.ocpus", "15"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.shape_name", "VM.GPU.A10.1"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.type", "FIXED_GPU_FLEET_SHAPE"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.type", "SERVICE_MANAGED_GPU_FLEET"),
				resource.TestCheckResourceAttr(resourceName, "network.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network.0.subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BatchBatchContextResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BatchBatchContextOptionalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_context", "test_batch_context", acctest.Optional, acctest.Create, BatchBatchContextRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", BatchBatchContextDisplayName),
				resource.TestCheckResourceAttr(resourceName, "entitlements.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "fleets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.max_concurrent_tasks", "10"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.disk_size_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.memory_in_gbs", "240"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.ocpus", "15"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.shape_name", "VM.GPU.A10.1"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.type", "FIXED_GPU_FLEET_SHAPE"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.type", "SERVICE_MANAGED_GPU_FLEET"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "job_priority_configurations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "job_priority_configurations.0.tag_key", "tagKey"),
				resource.TestCheckResourceAttr(resourceName, "job_priority_configurations.0.tag_namespace", "tagNamespace"),
				resource.TestCheckResourceAttr(resourceName, "job_priority_configurations.0.values.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "job_priority_configurations.0.weight", "10"),
				resource.TestCheckResourceAttr(resourceName, "logging_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "logging_configuration.0.is_job_task_events_propagation_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "logging_configuration.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "logging_configuration.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "logging_configuration.0.type", "OCI_LOGGING"),
				resource.TestCheckResourceAttr(resourceName, "network.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "network.0.vnics.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "system_tags"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + BatchBatchContextOptionalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_context", "test_batch_context", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(BatchBatchContextRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", BatchBatchContextDisplayName),
				resource.TestCheckResourceAttr(resourceName, "entitlements.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "fleets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.max_concurrent_tasks", "10"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.disk_size_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.memory_in_gbs", "240"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.ocpus", "15"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.shape_name", "VM.GPU.A10.1"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.type", "FIXED_GPU_FLEET_SHAPE"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.type", "SERVICE_MANAGED_GPU_FLEET"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "job_priority_configurations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "job_priority_configurations.0.tag_key", "tagKey"),
				resource.TestCheckResourceAttr(resourceName, "job_priority_configurations.0.tag_namespace", "tagNamespace"),
				resource.TestCheckResourceAttr(resourceName, "job_priority_configurations.0.values.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "job_priority_configurations.0.weight", "10"),
				resource.TestCheckResourceAttr(resourceName, "logging_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "logging_configuration.0.is_job_task_events_propagation_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "logging_configuration.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "logging_configuration.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "logging_configuration.0.type", "OCI_LOGGING"),
				resource.TestCheckResourceAttr(resourceName, "network.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "network.0.vnics.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "system_tags"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + BatchBatchContextOptionalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_context", "test_batch_context", acctest.Optional, acctest.Update, BatchBatchContextRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", BatchBatchContextDisplayNameUpdate),
				resource.TestCheckResourceAttr(resourceName, "entitlements.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "fleets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.max_concurrent_tasks", "10"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.disk_size_in_gbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.memory_in_gbs", "240"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.ocpus", "15"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.shape_name", "VM.GPU.A10.1"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.type", "FIXED_GPU_FLEET_SHAPE"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.type", "SERVICE_MANAGED_GPU_FLEET"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "job_priority_configurations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "job_priority_configurations.0.tag_key", "tagKey2"),
				resource.TestCheckResourceAttr(resourceName, "job_priority_configurations.0.tag_namespace", "tagNamespace2"),
				resource.TestCheckResourceAttr(resourceName, "job_priority_configurations.0.values.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "job_priority_configurations.0.weight", "11"),
				resource.TestCheckResourceAttr(resourceName, "logging_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "logging_configuration.0.is_job_task_events_propagation_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "logging_configuration.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "logging_configuration.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "logging_configuration.0.type", "OCI_LOGGING"),
				resource.TestCheckResourceAttr(resourceName, "network.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "network.0.vnics.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "system_tags"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_batch_batch_contexts", "test_batch_contexts", acctest.Optional, acctest.Update, BatchBatchContextDataSourceRepresentation) +
				compartmentIdVariableStr + BatchBatchContextOptionalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_batch_batch_context", "test_batch_context", acctest.Optional, acctest.Update, BatchBatchContextRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", BatchBatchContextDisplayNameUpdate),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "batch_context_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "batch_context_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_batch_batch_context", "test_batch_context", acctest.Required, acctest.Create, BatchBatchContextSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BatchBatchContextResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "batch_context_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", BatchBatchContextDisplayNameUpdate),
				resource.TestCheckResourceAttr(singularDatasourceName, "entitlements.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.0.max_concurrent_tasks", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.0.name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.0.shape.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.0.shape.0.disk_size_in_gbs", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.0.shape.0.memory_in_gbs", "240"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.0.shape.0.ocpus", "15"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.0.shape.0.shape_name", "VM.GPU.A10.1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.0.shape.0.type", "FIXED_GPU_FLEET_SHAPE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.0.type", "SERVICE_MANAGED_GPU_FLEET"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "job_priority_configurations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "job_priority_configurations.0.tag_key", "tagKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "job_priority_configurations.0.tag_namespace", "tagNamespace2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "job_priority_configurations.0.values.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "job_priority_configurations.0.weight", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "logging_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "logging_configuration.0.is_job_task_events_propagation_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "logging_configuration.0.type", "OCI_LOGGING"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network.0.vnics.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + compartmentIdVariableStr + BatchBatchContextRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

// issue-routing-tag: batch/default
func TestBatchBatchContextResource_cpuFleet(t *testing.T) {
	httpreplay.SetScenario("TestBatchBatchContextResource_cpuFleet")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_batch_batch_context.test_batch_context_cpu"
	datasourceName := "data.oci_batch_batch_contexts.test_batch_contexts_cpu"
	singularDatasourceName := "data.oci_batch_batch_context.test_batch_context_cpu"

	acctest.SaveConfigContent(config+compartmentIdVariableStr+BatchBatchContextCpuRequiredOnlyResource, "batch", "batchContextCpu", t)

	acctest.ResourceTest(t, testAccCheckBatchBatchContextDestroy, []resource.TestStep{
		{
			Config: config + compartmentIdVariableStr + BatchBatchContextCpuRequiredOnlyResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "fleets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.max_concurrent_tasks", "1"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.name", "cpu-fleet"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.type", "SERVICE_MANAGED_FLEET"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.memory_in_gbs", "16"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.ocpus", "1"),
				resource.TestCheckResourceAttr(resourceName, "fleets.0.shape.0.shape_name", "VM.Standard.E5.Flex"),
				resource.TestCheckResourceAttr(resourceName, "network.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network.0.subnet_id"),
			),
		},
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_batch_batch_contexts", "test_batch_contexts_cpu", acctest.Optional, acctest.Create, BatchBatchContextCpuDataSourceRepresentation) +
				compartmentIdVariableStr + BatchBatchContextCpuRequiredOnlyResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "batch_context_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "batch_context_collection.0.items.#", "1"),
			),
		},
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_batch_batch_context", "test_batch_context_cpu", acctest.Required, acctest.Create, BatchBatchContextCpuSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BatchBatchContextCpuRequiredOnlyResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "batch_context_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.0.max_concurrent_tasks", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.0.name", "cpu-fleet"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.0.type", "SERVICE_MANAGED_FLEET"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.0.shape.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.0.shape.0.memory_in_gbs", "16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.0.shape.0.ocpus", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fleets.0.shape.0.shape_name", "VM.Standard.E5.Flex"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		{
			Config:                  config + compartmentIdVariableStr + BatchBatchContextCpuRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckBatchBatchContextDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).BatchComputingClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_batch_batch_context" {
			noResourceFound = false
			request := oci_batch.GetBatchContextRequest{}

			tmp := rs.Primary.ID
			request.BatchContextId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "batch")

			response, err := client.GetBatchContext(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_batch.BatchContextLifecycleStateDeleted): true,
					string(oci_batch.BatchContextLifecycleStateFailed):  true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
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
	if !acctest.InSweeperExcludeList("BatchBatchContext") {
		resource.AddTestSweepers("BatchBatchContext", &resource.Sweeper{
			Name:         "BatchBatchContext",
			Dependencies: acctest.DependencyGraph["batchContext"],
			F:            sweepBatchBatchContextResource,
		})
	}
}

func sweepBatchBatchContextResource(compartment string) error {
	batchComputingClient := acctest.GetTestClients(&schema.ResourceData{}).BatchComputingClient()
	batchContextIds, err := getBatchBatchContextIds(compartment)
	if err != nil {
		return err
	}
	for _, batchContextId := range batchContextIds {
		if ok := acctest.SweeperDefaultResourceId[batchContextId]; !ok {
			deleteBatchContextRequest := oci_batch.DeleteBatchContextRequest{}

			deleteBatchContextRequest.BatchContextId = &batchContextId

			deleteBatchContextRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "batch")
			_, error := batchComputingClient.DeleteBatchContext(context.Background(), deleteBatchContextRequest)
			if error != nil {
				fmt.Printf("Error deleting BatchContext %s %s, It is possible that the resource is already deleted. Please verify manually \n", batchContextId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &batchContextId, BatchBatchContextSweepWaitCondition, time.Duration(3*time.Minute),
				BatchBatchContextSweepResponseFetchOperation, "batch", true)
		}
	}
	return nil
}

func getBatchBatchContextIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BatchContextId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	batchComputingClient := acctest.GetTestClients(&schema.ResourceData{}).BatchComputingClient()

	listBatchContextsRequest := oci_batch.ListBatchContextsRequest{}
	listBatchContextsRequest.CompartmentId = &compartmentId
	listBatchContextsRequest.LifecycleState = oci_batch.BatchContextLifecycleStateActive
	listBatchContextsResponse, err := batchComputingClient.ListBatchContexts(context.Background(), listBatchContextsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting BatchContext list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, batchContext := range listBatchContextsResponse.Items {
		id := *batchContext.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BatchContextId", id)
	}
	return resourceIds, nil
}

func BatchBatchContextSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if batchContextResponse, ok := response.Response.(oci_batch.GetBatchContextResponse); ok {
		return batchContextResponse.LifecycleState != oci_batch.BatchContextLifecycleStateDeleted
	}
	return false
}

func BatchBatchContextSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.BatchComputingClient().GetBatchContext(context.Background(), oci_batch.GetBatchContextRequest{
		BatchContextId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
