// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	StackMonitoringMetricExtensionRequiredOnlyResource = StackMonitoringMetricExtensionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_metric_extension", "test_metric_extension", acctest.Required, acctest.Create, StackMonitoringMetricExtensionRepresentation)

	StackMonitoringMetricExtensionResourceConfig = StackMonitoringMetricExtensionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_metric_extension", "test_metric_extension", acctest.Optional, acctest.Update, StackMonitoringMetricExtensionRepresentation)

	StackMonitoringMetricExtensionSingularDataSourceRepresentation = map[string]interface{}{
		"metric_extension_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_stack_monitoring_metric_extension.test_metric_extension.id}`},
	}

	StackMonitoringMetricExtensionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `ME_CountOfRunningInstances`},
		"resource_type":  acctest.Representation{RepType: acctest.Optional, Create: `ebs_instance`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"status":         acctest.Representation{RepType: acctest.Optional, Create: `DRAFT`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringMetricExtensionDataSourceFilterRepresentation}}

	StackMonitoringMetricExtensionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_stack_monitoring_metric_extension.test_metric_extension.id}`}},
	}

	StackMonitoringMetricExtensionRepresentation = map[string]interface{}{
		"collection_recurrences": acctest.Representation{RepType: acctest.Required, Create: `FREQ=MINUTELY;INTERVAL=10`, Update: `FREQ=MINUTELY;INTERVAL=5`},
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `Count of Running Instances`, Update: `All Server Instances Running Factor`},
		"metric_list":            acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringMetricExtensionMetricListRepresentation},
		"name":                   acctest.Representation{RepType: acctest.Required, Create: `ME_CountOfRunningInstances`},
		"query_properties":       acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringMetricExtensionQueryPropertiesRepresentation},
		"resource_type":          acctest.Representation{RepType: acctest.Required, Create: `ebs_instance`},
		"description":            acctest.Representation{RepType: acctest.Optional, Create: `Collects count of instances in 'UP' status in staging compartments from monitoring table`, Update: `Gives value 1 when All servers are in 'UP' status in production compartments in monitoring table`},
		"publish_trigger":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
	}
	StackMonitoringMetricExtensionMetricListRepresentation = map[string]interface{}{
		"data_type":          acctest.Representation{RepType: acctest.Required, Create: `NUMBER`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `CountOfRunningInstances`, Update: `AllServerStatus`},
		"compute_expression": acctest.Representation{RepType: acctest.Optional, Create: `CountOfRunningInstances - _CountOfRunningInstances`, Update: `(AllServerStatus >= _AllServerStatus) ? 1 : 0`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `Count of Running Instances`, Update: `All Server Running Factor`},
		"is_dimension":       acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"is_hidden":          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"metric_category":    acctest.Representation{RepType: acctest.Optional, Create: `AVAILABILITY`, Update: `UTILIZATION`},
		"unit":               acctest.Representation{RepType: acctest.Optional, Create: `count`, Update: ` `},
	}
	StackMonitoringMetricExtensionQueryPropertiesRepresentation = map[string]interface{}{
		"collection_method": acctest.Representation{RepType: acctest.Required, Create: `SQL`},
		"in_param_details":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: StackMonitoringMetricExtensionQueryPropertiesInParamDetailsRepresentation},
		"out_param_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: StackMonitoringMetricExtensionQueryPropertiesOutParamDetailsRepresentation},
		"sql_details":       acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringMetricExtensionQueryPropertiesSqlDetailsRepresentation},
		"sql_type":          acctest.Representation{RepType: acctest.Required, Create: `STATEMENT`, Update: `STATEMENT`},
	}
	StackMonitoringMetricExtensionQueryPropertiesInParamDetailsRepresentation = map[string]interface{}{
		"in_param_position": acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `2`},
		"in_param_value":    acctest.Representation{RepType: acctest.Optional, Create: `staging`, Update: `production`},
	}
	StackMonitoringMetricExtensionQueryPropertiesOutParamDetailsRepresentation = map[string]interface{}{
		"out_param_position": acctest.Representation{RepType: acctest.Optional, Create: `2`, Update: `3`},
		"out_param_type":     acctest.Representation{RepType: acctest.Optional, Create: `SQL_CURSOR`, Update: `ARRAY`},
	}

	StackMonitoringMetricExtensionQueryPropertiesSqlDetailsRepresentation = map[string]interface{}{
		"content":          acctest.Representation{RepType: acctest.Required, Create: `U0VMRUNUIGNvdW50KGluc3RhbmNlX2lkKSBGUk9NIG1vbml0b3JpbmdfdGFibGUgV0hFUkUgc3RhdHVzID0gJ1VQJyBBTkQgY29tcGFydG1lbnRfdHlwZSA9IDox`, Update: `U0VMRUNUIGNvdW50KGluc3RhbmNlX2lkKSBGUk9NIG1vbml0b3JpbmdfdGFibGUgV0hFUkUgc3RhdHVzID0gJ1VQJyBBTkQgY29tcGFydG1lbnRfdHlwZSA9IDoy`},
		"script_file_name": acctest.Representation{RepType: acctest.Optional, Create: `No-File`, Update: ` `},
	}

	StackMonitoringMetricExtensionForPublishRepresentation = map[string]interface{}{
		"collection_recurrences": acctest.Representation{RepType: acctest.Required, Create: `FREQ=HOURLY;INTERVAL=6`},
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `OS File System Utilization`},
		"metric_list":            []acctest.RepresentationGroup{{RepType: acctest.Required, Group: StackMonitoringMetricExtensionMetricList0ForPublishRepresentation}, {RepType: acctest.Required, Group: StackMonitoringMetricExtensionMetricList1ForPublishRepresentation}, {RepType: acctest.Required, Group: StackMonitoringMetricExtensionMetricList2ForPublishRepresentation}, {RepType: acctest.Required, Group: StackMonitoringMetricExtensionMetricList3ForPublishRepresentation}},
		"name":                   acctest.Representation{RepType: acctest.Required, Create: `ME_OsFileSystemUtilization`},
		"query_properties":       acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringMetricExtensionQueryPropertiesForPublishRepresentation},
		"resource_type":          acctest.Representation{RepType: acctest.Required, Create: `host_linux`},
		"description":            acctest.Representation{RepType: acctest.Optional, Create: `Computes File System Utilization Percentage of various mount points`},
		"publish_trigger":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	StackMonitoringMetricExtensionMetricList0ForPublishRepresentation = map[string]interface{}{
		"data_type":          acctest.Representation{RepType: acctest.Required, Create: `STRING`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `MountPoint`},
		"compute_expression": acctest.Representation{RepType: acctest.Optional, Create: ` `, Update: ` `},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: ` `, Update: ` `},
		"is_dimension":       acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"is_hidden":          acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"metric_category":    acctest.Representation{RepType: acctest.Optional, Create: ``},
		"unit":               acctest.Representation{RepType: acctest.Optional, Create: ` `, Update: ` `},
	}

	StackMonitoringMetricExtensionMetricList1ForPublishRepresentation = map[string]interface{}{
		"data_type":          acctest.Representation{RepType: acctest.Required, Create: `NUMBER`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `FileSystemSize`},
		"compute_expression": acctest.Representation{RepType: acctest.Optional, Create: ` `, Update: ` `},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: ` `, Update: ` `},
		"is_dimension":       acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_hidden":          acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"metric_category":    acctest.Representation{RepType: acctest.Optional, Create: ``},
		"unit":               acctest.Representation{RepType: acctest.Optional, Create: ` `, Update: ` `},
	}

	StackMonitoringMetricExtensionMetricList2ForPublishRepresentation = map[string]interface{}{
		"data_type":          acctest.Representation{RepType: acctest.Required, Create: `NUMBER`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `FileSystemUsed`},
		"compute_expression": acctest.Representation{RepType: acctest.Optional, Create: ` `, Update: ` `},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: ` `, Update: ` `},
		"is_dimension":       acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_hidden":          acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"metric_category":    acctest.Representation{RepType: acctest.Optional, Create: ``},
		"unit":               acctest.Representation{RepType: acctest.Optional, Create: ` `, Update: ` `},
	}

	StackMonitoringMetricExtensionMetricList3ForPublishRepresentation = map[string]interface{}{
		"data_type":          acctest.Representation{RepType: acctest.Required, Create: `NUMBER`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `FileSystemUsage`},
		"compute_expression": acctest.Representation{RepType: acctest.Optional, Create: `(FileSystemUsed / FileSystemSize) * 100`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `File System Usage`},
		"is_dimension":       acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_hidden":          acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"metric_category":    acctest.Representation{RepType: acctest.Optional, Create: `UTILIZATION`},
		"unit":               acctest.Representation{RepType: acctest.Optional, Create: `percent`},
	}
	StackMonitoringMetricExtensionQueryPropertiesForPublishRepresentation = map[string]interface{}{
		"collection_method": acctest.Representation{RepType: acctest.Required, Create: `OS_COMMAND`},
		"command":           acctest.Representation{RepType: acctest.Optional, Create: `/bin/bash`},
		"delimiter":         acctest.Representation{RepType: acctest.Optional, Create: `|`},
		"script_details":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: StackMonitoringMetricExtensionQueryPropertiesScriptDetailsForPublishRepresentation},
		"starts_with":       acctest.Representation{RepType: acctest.Optional, Create: `oci_result=`},
	}

	StackMonitoringMetricExtensionQueryPropertiesScriptDetailsForPublishRepresentation = map[string]interface{}{
		"content": acctest.Representation{RepType: acctest.Optional, Create: `IyEvYmluL2Jhc2gKIyBDb3B5cmlnaHQgKGMpIDIwMjIsIE9yYWNsZSBhbmQvb3IgaXRzIGFmZmlsaWF0ZXMuIEFsbCByaWdodHMgcmVzZXJ2ZWQuCiMKIyBTdGFjayBNb25pdG9yaW5nIC8gSG9zdDogY29sbGVjdCBmaWxlc3lzdGVtIHN0YXRpc3RpY3MgZnJvbSBMaW51eCBob3N0cwojCiMgT3V0cHV0IGZvcm1hdDoKIwojIHJlc3VsdD1tb3VudHxzaXplfHVzZWQKCmV4ZWMgMTA+JjEKZXhlYyAxPiYyCgoKd2hpbGUgcmVhZCAtciBkZXYgc2l6ZSB1c2VkIGF2YWlsIHVzZWRwIG1vdW50IG90aGVyCmRvCiAgICBpZiBbWyAiJHtkZXZ9IiA9fiAvIF1dCiAgICB0aGVuCiAgICAgICAgaWYgWyAiJHt0b3R9IiA9PSAiMCIgXQogICAgICAgIHRoZW4KICAgICAgICAgICAgIyBQcmV2ZW50IGRldmlzaW9uIGJ5IHplcm8KICAgICAgICAgICAgdXNlZD0wCiAgICAgICAgICAgIHVzZWRwPTAKICAgICAgICBmaQoKICAgICAgICBwcmludGYgIm9jaV9yZXN1bHQ9JXN8JXN8JXNcbiIgIiR7bW91bnR9IiAiJHtzaXplfSIgIiR7dXNlZH0iID4mMTAKICAgIGZpCmRvbmUgPCA8KGRmIC1rIDI+L2Rldi9udWxsKQ==`},
		"name":    acctest.Representation{RepType: acctest.Optional, Create: `fileSystem.sh`},
	}

	StackMonitoringMetricExtensionJmxRequiredOnlyResource = StackMonitoringMetricExtensionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_metric_extension", "test_metric_extension_jmx", acctest.Required, acctest.Create, StackMonitoringMetricExtensionJmxRepresentation)

	StackMonitoringMetricExtensionJmxRepresentation = map[string]interface{}{
		"collection_recurrences": acctest.Representation{RepType: acctest.Required, Create: `FREQ=MINUTELY;INTERVAL=10`, Update: `FREQ=MINUTELY;INTERVAL=5`},
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `Total Physical Memory In GigaBytes`, Update: `Total Physical Memory In GBs`},
		"metric_list":            []acctest.RepresentationGroup{{RepType: acctest.Required, Group: StackMonitoringMetricExtensionJmxMetricList0Representation}, {RepType: acctest.Required, Group: StackMonitoringMetricExtensionJmxMetricList1Representation}},
		"name":                   acctest.Representation{RepType: acctest.Required, Create: `ME_TotalPhysicalMemoryInGB`},
		"query_properties":       acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringMetricExtensionJmxQueryPropertiesRepresentation},
		"resource_type":          acctest.Representation{RepType: acctest.Required, Create: `weblogic_j2eeserver`},
		"description":            acctest.Representation{RepType: acctest.Optional, Create: `Collects TotalPhysicalMemoryInGB for server named PIA in Giga bytes`, Update: `Collects TotalPhysicalMemoryInGB for server named PIA in GBs`},
		"publish_trigger":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
	}
	StackMonitoringMetricExtensionJmxMetricList0Representation = map[string]interface{}{
		"data_type":    acctest.Representation{RepType: acctest.Required, Create: `NUMBER`},
		"name":         acctest.Representation{RepType: acctest.Required, Create: `TotalPhysicalMemorySize`, Update: `TotalPhysicalMemorySizeInBytes`},
		"is_dimension": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"is_hidden":    acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
	}
	StackMonitoringMetricExtensionJmxMetricList1Representation = map[string]interface{}{
		"data_type":          acctest.Representation{RepType: acctest.Required, Create: `NUMBER`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `TotalPhysicalMemorySizeGigaBytes`, Update: `TotalPhysicalMemorySizeGB`},
		"compute_expression": acctest.Representation{RepType: acctest.Optional, Create: `TotalPhysicalMemorySize > 0 ? (TotalPhysicalMemorySize / (1024 * 1024 * 1024))`, Update: `TotalPhysicalMemorySizeInBytes > 0 ? (TotalPhysicalMemorySizeInBytes / (1024 * 1024 * 1024))`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `TotalPhysicalMemorySizeInGigaBytes`, Update: `TotalPhysicalMemorySizeInGBs`},
		"is_dimension":       acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"is_hidden":          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
	}
	StackMonitoringMetricExtensionJmxQueryPropertiesRepresentation = map[string]interface{}{
		"collection_method":  acctest.Representation{RepType: acctest.Required, Create: `JMX`},
		"managed_bean_query": acctest.Representation{RepType: acctest.Required, Create: `java.lang:type=OperatingSystem,Location=PIA`},
		"jmx_attributes":     acctest.Representation{RepType: acctest.Required, Create: `TotalPhysicalMemorySize`},
		"auto_row_prefix":    acctest.Representation{RepType: acctest.Optional, Create: `PseudoKey`},
	}

	StackMonitoringMetricExtensionResourceDependencies = ""
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringMetricExtensionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringMetricExtensionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_stack_monitoring_metric_extension.test_metric_extension"
	datasourceName := "data.oci_stack_monitoring_metric_extensions.test_metric_extensions"
	singularDatasourceName := "data.oci_stack_monitoring_metric_extension.test_metric_extension"

	resourceNameForPublish := "oci_stack_monitoring_metric_extension.test_metric_extension_for_publish"

	var resId, resId2, resId3, resId4 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+StackMonitoringMetricExtensionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_metric_extension", "test_metric_extension", acctest.Optional, acctest.Create, StackMonitoringMetricExtensionRepresentation), "stackmonitoring", "metricExtension", t)

	acctest.ResourceTest(t, testAccCheckStackMonitoringMetricExtensionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + StackMonitoringMetricExtensionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_metric_extension", "test_metric_extension", acctest.Required, acctest.Create, StackMonitoringMetricExtensionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "collection_recurrences", "FREQ=MINUTELY;INTERVAL=10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Count of Running Instances"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.data_type", "NUMBER"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.name", "CountOfRunningInstances"),
				resource.TestCheckResourceAttr(resourceName, "name", "ME_CountOfRunningInstances"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.collection_method", "SQL"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.sql_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.sql_details.0.content", "U0VMRUNUIGNvdW50KGluc3RhbmNlX2lkKSBGUk9NIG1vbml0b3JpbmdfdGFibGUgV0hFUkUgc3RhdHVzID0gJ1VQJyBBTkQgY29tcGFydG1lbnRfdHlwZSA9IDox"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.sql_type", "STATEMENT"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "ebs_instance"),
				resource.TestCheckResourceAttr(resourceName, "status", "DRAFT"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + StackMonitoringMetricExtensionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + StackMonitoringMetricExtensionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_metric_extension", "test_metric_extension", acctest.Optional, acctest.Create, StackMonitoringMetricExtensionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "collection_method"),
				resource.TestCheckResourceAttr(resourceName, "collection_recurrences", "FREQ=MINUTELY;INTERVAL=10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "Collects count of instances in 'UP' status in staging compartments from monitoring table"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Count of Running Instances"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.compute_expression", "CountOfRunningInstances - _CountOfRunningInstances"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.data_type", "NUMBER"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.display_name", "Count of Running Instances"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.is_dimension", "false"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.is_hidden", "false"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.metric_category", "AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.name", "CountOfRunningInstances"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.unit", "count"),
				resource.TestCheckResourceAttr(resourceName, "name", "ME_CountOfRunningInstances"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.collection_method", "SQL"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.in_param_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.in_param_details.0.in_param_position", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.in_param_details.0.in_param_value", "staging"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.out_param_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.out_param_details.0.out_param_position", "2"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.out_param_details.0.out_param_type", "SQL_CURSOR"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.sql_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.sql_details.0.content", "U0VMRUNUIGNvdW50KGluc3RhbmNlX2lkKSBGUk9NIG1vbml0b3JpbmdfdGFibGUgV0hFUkUgc3RhdHVzID0gJ1VQJyBBTkQgY29tcGFydG1lbnRfdHlwZSA9IDox"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.sql_details.0.script_file_name", "No-File"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.sql_type", "STATEMENT"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "ebs_instance"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + StackMonitoringMetricExtensionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_metric_extension", "test_metric_extension", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(StackMonitoringMetricExtensionRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "collection_method"),
				resource.TestCheckResourceAttr(resourceName, "collection_recurrences", "FREQ=MINUTELY;INTERVAL=10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "Collects count of instances in 'UP' status in staging compartments from monitoring table"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Count of Running Instances"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.compute_expression", "CountOfRunningInstances - _CountOfRunningInstances"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.data_type", "NUMBER"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.display_name", "Count of Running Instances"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.is_dimension", "false"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.is_hidden", "false"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.metric_category", "AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.name", "CountOfRunningInstances"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.unit", "count"),
				resource.TestCheckResourceAttr(resourceName, "name", "ME_CountOfRunningInstances"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.collection_method", "SQL"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.in_param_details.0.in_param_position", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.sql_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.sql_details.0.content", "U0VMRUNUIGNvdW50KGluc3RhbmNlX2lkKSBGUk9NIG1vbml0b3JpbmdfdGFibGUgV0hFUkUgc3RhdHVzID0gJ1VQJyBBTkQgY29tcGFydG1lbnRfdHlwZSA9IDox"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.sql_details.0.script_file_name", "No-File"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.sql_type", "STATEMENT"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "ebs_instance"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),

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
			Config: config + compartmentIdVariableStr + StackMonitoringMetricExtensionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_metric_extension", "test_metric_extension", acctest.Optional, acctest.Update, StackMonitoringMetricExtensionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "collection_method"),
				resource.TestCheckResourceAttr(resourceName, "collection_recurrences", "FREQ=MINUTELY;INTERVAL=5"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "Gives value 1 when All servers are in 'UP' status in production compartments in monitoring table"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "All Server Instances Running Factor"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.compute_expression", "(AllServerStatus >= _AllServerStatus) ? 1 : 0"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.data_type", "NUMBER"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.display_name", "All Server Running Factor"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.is_dimension", "false"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.is_hidden", "false"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.metric_category", "UTILIZATION"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.name", "AllServerStatus"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.unit", " "),
				resource.TestCheckResourceAttr(resourceName, "name", "ME_CountOfRunningInstances"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.collection_method", "SQL"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.in_param_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.in_param_details.0.in_param_position", "2"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.in_param_details.0.in_param_value", "production"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.out_param_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.out_param_details.0.out_param_position", "3"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.out_param_details.0.out_param_type", "ARRAY"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.sql_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.sql_details.0.content", "U0VMRUNUIGNvdW50KGluc3RhbmNlX2lkKSBGUk9NIG1vbml0b3JpbmdfdGFibGUgV0hFUkUgc3RhdHVzID0gJ1VQJyBBTkQgY29tcGFydG1lbnRfdHlwZSA9IDoy"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.sql_details.0.script_file_name", " "),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.sql_type", "STATEMENT"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "ebs_instance"),
				resource.TestCheckResourceAttr(resourceName, "status", "DRAFT"),
				resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_metric_extensions", "test_metric_extensions", acctest.Optional, acctest.Update, StackMonitoringMetricExtensionDataSourceRepresentation) +
				compartmentIdVariableStr + StackMonitoringMetricExtensionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_metric_extension", "test_metric_extension", acctest.Optional, acctest.Update, StackMonitoringMetricExtensionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "ME_CountOfRunningInstances"),
				resource.TestCheckResourceAttr(datasourceName, "resource_type", "ebs_instance"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "status", "DRAFT"),
				resource.TestCheckResourceAttr(datasourceName, "metric_extension_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "metric_extension_collection.0.items.#", "1"),
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_metric_extension", "test_metric_extension", acctest.Required, acctest.Create, StackMonitoringMetricExtensionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + StackMonitoringMetricExtensionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "metric_extension_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "collection_method"),
				resource.TestCheckResourceAttr(singularDatasourceName, "collection_recurrences", "FREQ=MINUTELY;INTERVAL=5"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "Gives value 1 when All servers are in 'UP' status in production compartments in monitoring table"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "All Server Instances Running Factor"),
				resource.TestCheckResourceAttr(singularDatasourceName, "enabled_on_resources.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "enabled_on_resources_count", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "last_updated_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metric_list.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metric_list.0.data_type", "NUMBER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metric_list.0.display_name", "All Server Running Factor"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metric_list.0.is_dimension", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metric_list.0.is_hidden", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metric_list.0.metric_category", "UTILIZATION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metric_list.0.name", "AllServerStatus"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metric_list.0.unit", " "),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "ME_CountOfRunningInstances"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_properties.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_properties.0.collection_method", "SQL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_properties.0.in_param_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_properties.0.in_param_details.0.in_param_position", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_properties.0.in_param_details.0.in_param_value", "production"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_properties.0.out_param_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_properties.0.out_param_details.0.out_param_position", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_properties.0.out_param_details.0.out_param_type", "ARRAY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_properties.0.sql_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_properties.0.sql_details.0.content", "U0VMRUNUIGNvdW50KGluc3RhbmNlX2lkKSBGUk9NIG1vbml0b3JpbmdfdGFibGUgV0hFUkUgc3RhdHVzID0gJ1VQJyBBTkQgY29tcGFydG1lbnRfdHlwZSA9IDoy"),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_properties.0.sql_details.0.script_file_name", " "),
				resource.TestCheckResourceAttr(singularDatasourceName, "query_properties.0.sql_type", "STATEMENT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_type", "ebs_instance"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_uri"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenant_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},

		// verify resource import
		{
			Config:                  config + StackMonitoringMetricExtensionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},

		// verify Create for publish with optionals
		{
			Config: config + compartmentIdVariableStr + StackMonitoringMetricExtensionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_metric_extension", "test_metric_extension_for_publish", acctest.Optional, acctest.Create, StackMonitoringMetricExtensionForPublishRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameForPublish, "collection_method"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "collection_recurrences", "FREQ=HOURLY;INTERVAL=6"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceNameForPublish, "description", "Computes File System Utilization Percentage of various mount points"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "display_name", "OS File System Utilization"),
				resource.TestCheckResourceAttrSet(resourceNameForPublish, "id"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.#", "4"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.0.compute_expression", " "),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.0.data_type", "STRING"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.0.display_name", " "),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.0.is_dimension", "true"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.0.is_hidden", "false"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.0.metric_category", ""),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.0.name", "MountPoint"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.0.unit", " "),
				resource.TestCheckResourceAttr(resourceNameForPublish, "name", "ME_OsFileSystemUtilization"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.1.compute_expression", " "),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.1.data_type", "NUMBER"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.1.display_name", " "),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.1.is_dimension", "false"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.1.is_hidden", "true"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.1.metric_category", ""),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.1.name", "FileSystemSize"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.1.unit", " "),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.2.compute_expression", " "),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.2.data_type", "NUMBER"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.2.display_name", " "),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.2.is_dimension", "false"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.2.is_hidden", "true"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.2.metric_category", ""),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.2.name", "FileSystemUsed"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.2.unit", " "),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.3.compute_expression", "(FileSystemUsed / FileSystemSize) * 100"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.3.data_type", "NUMBER"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.3.display_name", "File System Usage"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.3.is_dimension", "false"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.3.is_hidden", "false"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.3.metric_category", "UTILIZATION"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.3.name", "FileSystemUsage"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.3.unit", "percent"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "query_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "query_properties.0.collection_method", "OS_COMMAND"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "query_properties.0.command", "/bin/bash"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "query_properties.0.delimiter", "|"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "query_properties.0.starts_with", "oci_result="),
				resource.TestCheckResourceAttr(resourceNameForPublish, "query_properties.0.script_details.0.content", "IyEvYmluL2Jhc2gKIyBDb3B5cmlnaHQgKGMpIDIwMjIsIE9yYWNsZSBhbmQvb3IgaXRzIGFmZmlsaWF0ZXMuIEFsbCByaWdodHMgcmVzZXJ2ZWQuCiMKIyBTdGFjayBNb25pdG9yaW5nIC8gSG9zdDogY29sbGVjdCBmaWxlc3lzdGVtIHN0YXRpc3RpY3MgZnJvbSBMaW51eCBob3N0cwojCiMgT3V0cHV0IGZvcm1hdDoKIwojIHJlc3VsdD1tb3VudHxzaXplfHVzZWQKCmV4ZWMgMTA+JjEKZXhlYyAxPiYyCgoKd2hpbGUgcmVhZCAtciBkZXYgc2l6ZSB1c2VkIGF2YWlsIHVzZWRwIG1vdW50IG90aGVyCmRvCiAgICBpZiBbWyAiJHtkZXZ9IiA9fiAvIF1dCiAgICB0aGVuCiAgICAgICAgaWYgWyAiJHt0b3R9IiA9PSAiMCIgXQogICAgICAgIHRoZW4KICAgICAgICAgICAgIyBQcmV2ZW50IGRldmlzaW9uIGJ5IHplcm8KICAgICAgICAgICAgdXNlZD0wCiAgICAgICAgICAgIHVzZWRwPTAKICAgICAgICBmaQoKICAgICAgICBwcmludGYgIm9jaV9yZXN1bHQ9JXN8JXN8JXNcbiIgIiR7bW91bnR9IiAiJHtzaXplfSIgIiR7dXNlZH0iID4mMTAKICAgIGZpCmRvbmUgPCA8KGRmIC1rIDI+L2Rldi9udWxsKQ=="),
				resource.TestCheckResourceAttr(resourceNameForPublish, "query_properties.0.script_details.0.name", "fileSystem.sh"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "status", "DRAFT"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "resource_type", "host_linux"),
				resource.TestCheckResourceAttrSet(resourceNameForPublish, "tenant_id"),

				func(s *terraform.State) (err error) {
					resId3, err = acctest.FromInstanceState(s, resourceNameForPublish, "id")
					return err
				},
			),
		},

		// verify updates to publish trigger parameter
		{
			Config: config + compartmentIdVariableStr + StackMonitoringMetricExtensionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_metric_extension", "test_metric_extension_for_publish", acctest.Optional, acctest.Update, StackMonitoringMetricExtensionForPublishRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameForPublish, "collection_method"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "collection_recurrences", "FREQ=HOURLY;INTERVAL=6"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceNameForPublish, "description", "Computes File System Utilization Percentage of various mount points"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "display_name", "OS File System Utilization"),
				resource.TestCheckResourceAttrSet(resourceNameForPublish, "id"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.#", "4"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.0.compute_expression", " "),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.0.data_type", "STRING"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.0.display_name", " "),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.0.is_dimension", "true"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.0.is_hidden", "false"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.0.metric_category", ""),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.0.name", "MountPoint"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.0.unit", " "),
				resource.TestCheckResourceAttr(resourceNameForPublish, "name", "ME_OsFileSystemUtilization"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.1.compute_expression", " "),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.1.data_type", "NUMBER"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.1.display_name", " "),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.1.is_dimension", "false"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.1.is_hidden", "true"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.1.metric_category", ""),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.1.name", "FileSystemSize"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.1.unit", " "),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.2.compute_expression", " "),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.2.data_type", "NUMBER"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.2.display_name", " "),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.2.is_dimension", "false"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.2.is_hidden", "true"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.2.metric_category", ""),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.2.name", "FileSystemUsed"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.2.unit", " "),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.3.compute_expression", "(FileSystemUsed / FileSystemSize) * 100"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.3.data_type", "NUMBER"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.3.display_name", "File System Usage"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.3.is_dimension", "false"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.3.is_hidden", "false"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.3.metric_category", "UTILIZATION"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.3.name", "FileSystemUsage"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "metric_list.3.unit", "percent"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "query_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "query_properties.0.collection_method", "OS_COMMAND"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "query_properties.0.command", "/bin/bash"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "query_properties.0.delimiter", "|"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "query_properties.0.script_details.#", "1"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "query_properties.0.script_details.0.content", "IyEvYmluL2Jhc2gKIyBDb3B5cmlnaHQgKGMpIDIwMjIsIE9yYWNsZSBhbmQvb3IgaXRzIGFmZmlsaWF0ZXMuIEFsbCByaWdodHMgcmVzZXJ2ZWQuCiMKIyBTdGFjayBNb25pdG9yaW5nIC8gSG9zdDogY29sbGVjdCBmaWxlc3lzdGVtIHN0YXRpc3RpY3MgZnJvbSBMaW51eCBob3N0cwojCiMgT3V0cHV0IGZvcm1hdDoKIwojIHJlc3VsdD1tb3VudHxzaXplfHVzZWQKCmV4ZWMgMTA+JjEKZXhlYyAxPiYyCgoKd2hpbGUgcmVhZCAtciBkZXYgc2l6ZSB1c2VkIGF2YWlsIHVzZWRwIG1vdW50IG90aGVyCmRvCiAgICBpZiBbWyAiJHtkZXZ9IiA9fiAvIF1dCiAgICB0aGVuCiAgICAgICAgaWYgWyAiJHt0b3R9IiA9PSAiMCIgXQogICAgICAgIHRoZW4KICAgICAgICAgICAgIyBQcmV2ZW50IGRldmlzaW9uIGJ5IHplcm8KICAgICAgICAgICAgdXNlZD0wCiAgICAgICAgICAgIHVzZWRwPTAKICAgICAgICBmaQoKICAgICAgICBwcmludGYgIm9jaV9yZXN1bHQ9JXN8JXN8JXNcbiIgIiR7bW91bnR9IiAiJHtzaXplfSIgIiR7dXNlZH0iID4mMTAKICAgIGZpCmRvbmUgPCA8KGRmIC1rIDI+L2Rldi9udWxsKQ=="),
				resource.TestCheckResourceAttr(resourceNameForPublish, "query_properties.0.script_details.0.name", "fileSystem.sh"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "query_properties.0.starts_with", "oci_result="),
				resource.TestCheckResourceAttr(resourceNameForPublish, "resource_type", "host_linux"),
				resource.TestCheckResourceAttr(resourceNameForPublish, "status", "PUBLISHED"),
				resource.TestCheckResourceAttrSet(resourceNameForPublish, "tenant_id"),

				func(s *terraform.State) (err error) {
					resId4, err = acctest.FromInstanceState(s, resourceNameForPublish, "id")
					if resId3 != resId4 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// delete resource for publish from previous step
		{
			Config: config + compartmentIdVariableStr + StackMonitoringMetricExtensionResourceDependencies,
		},
	})
}

func TestStackMonitoringMetricExtensionResource_jmx(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringMetricExtensionResource_jmx")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_stack_monitoring_metric_extension.test_metric_extension_jmx"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+StackMonitoringMetricExtensionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_metric_extension", "test_metric_extension_jmx", acctest.Optional, acctest.Create, StackMonitoringMetricExtensionJmxRepresentation), "stackmonitoring", "metricExtension", t)

	acctest.ResourceTest(t, testAccCheckStackMonitoringMetricExtensionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + StackMonitoringMetricExtensionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_metric_extension", "test_metric_extension_jmx", acctest.Required, acctest.Create, StackMonitoringMetricExtensionJmxRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "collection_recurrences", "FREQ=MINUTELY;INTERVAL=10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Total Physical Memory In GigaBytes"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.data_type", "NUMBER"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.name", "TotalPhysicalMemorySize"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.1.data_type", "NUMBER"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.1.name", "TotalPhysicalMemorySizeGigaBytes"),
				resource.TestCheckResourceAttr(resourceName, "name", "ME_TotalPhysicalMemoryInGB"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.collection_method", "JMX"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.managed_bean_query", "java.lang:type=OperatingSystem,Location=PIA"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.jmx_attributes", "TotalPhysicalMemorySize"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "weblogic_j2eeserver"),
				resource.TestCheckResourceAttr(resourceName, "status", "DRAFT"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + StackMonitoringMetricExtensionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + StackMonitoringMetricExtensionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_metric_extension", "test_metric_extension_jmx", acctest.Optional, acctest.Create, StackMonitoringMetricExtensionJmxRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "collection_method"),
				resource.TestCheckResourceAttr(resourceName, "collection_recurrences", "FREQ=MINUTELY;INTERVAL=10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Total Physical Memory In GigaBytes"),
				resource.TestCheckResourceAttr(resourceName, "description", "Collects TotalPhysicalMemoryInGB for server named PIA in Giga bytes"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.data_type", "NUMBER"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.name", "TotalPhysicalMemorySize"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.is_hidden", "true"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.is_dimension", "false"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.1.data_type", "NUMBER"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.1.is_dimension", "false"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.1.is_hidden", "false"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.1.name", "TotalPhysicalMemorySizeGigaBytes"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.1.display_name", "TotalPhysicalMemorySizeInGigaBytes"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.1.compute_expression", "TotalPhysicalMemorySize > 0 ? (TotalPhysicalMemorySize / (1024 * 1024 * 1024))"),
				resource.TestCheckResourceAttr(resourceName, "name", "ME_TotalPhysicalMemoryInGB"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.collection_method", "JMX"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.managed_bean_query", "java.lang:type=OperatingSystem,Location=PIA"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.jmx_attributes", "TotalPhysicalMemorySize"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.auto_row_prefix", "PseudoKey"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "weblogic_j2eeserver"),
				resource.TestCheckResourceAttr(resourceName, "status", "DRAFT"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),

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
			Config: config + compartmentIdVariableStr + StackMonitoringMetricExtensionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_metric_extension", "test_metric_extension_jmx", acctest.Optional, acctest.Update, StackMonitoringMetricExtensionJmxRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "collection_method"),
				resource.TestCheckResourceAttr(resourceName, "collection_recurrences", "FREQ=MINUTELY;INTERVAL=5"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "Collects TotalPhysicalMemoryInGB for server named PIA in GBs"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Total Physical Memory In GBs"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.data_type", "NUMBER"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.name", "TotalPhysicalMemorySizeInBytes"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.is_hidden", "true"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.0.is_dimension", "false"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.1.data_type", "NUMBER"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.1.is_dimension", "false"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.1.is_hidden", "false"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.1.name", "TotalPhysicalMemorySizeGB"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.1.display_name", "TotalPhysicalMemorySizeInGBs"),
				resource.TestCheckResourceAttr(resourceName, "metric_list.1.compute_expression", "TotalPhysicalMemorySizeInBytes > 0 ? (TotalPhysicalMemorySizeInBytes / (1024 * 1024 * 1024))"),
				resource.TestCheckResourceAttr(resourceName, "name", "ME_TotalPhysicalMemoryInGB"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.collection_method", "JMX"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.managed_bean_query", "java.lang:type=OperatingSystem,Location=PIA"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.jmx_attributes", "TotalPhysicalMemorySize"),
				resource.TestCheckResourceAttr(resourceName, "query_properties.0.auto_row_prefix", "PseudoKey"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "weblogic_j2eeserver"),
				resource.TestCheckResourceAttr(resourceName, "status", "DRAFT"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),

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
			Config:                  config + StackMonitoringMetricExtensionJmxRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},

		// delete jmx resource from previous step
		{
			Config: config + compartmentIdVariableStr + StackMonitoringMetricExtensionResourceDependencies,
		},
	})
}

func testAccCheckStackMonitoringMetricExtensionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).StackMonitoringClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_stack_monitoring_metric_extension" {
			noResourceFound = false
			request := oci_stack_monitoring.GetMetricExtensionRequest{}

			tmp := rs.Primary.ID
			request.MetricExtensionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "stack_monitoring")

			response, err := client.GetMetricExtension(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_stack_monitoring.MetricExtensionLifeCycleStatesDeleted): true,
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
	if !acctest.InSweeperExcludeList("StackMonitoringMetricExtension") {
		resource.AddTestSweepers("StackMonitoringMetricExtension", &resource.Sweeper{
			Name:         "StackMonitoringMetricExtension",
			Dependencies: acctest.DependencyGraph["metricExtension"],
			F:            sweepStackMonitoringMetricExtensionResource,
		})
	}
}

func sweepStackMonitoringMetricExtensionResource(compartment string) error {
	stackMonitoringClient := acctest.GetTestClients(&schema.ResourceData{}).StackMonitoringClient()
	metricExtensionIds, err := getStackMonitoringMetricExtensionIds(compartment)
	if err != nil {
		return err
	}
	for _, metricExtensionId := range metricExtensionIds {
		if ok := acctest.SweeperDefaultResourceId[metricExtensionId]; !ok {
			deleteMetricExtensionRequest := oci_stack_monitoring.DeleteMetricExtensionRequest{}

			deleteMetricExtensionRequest.MetricExtensionId = &metricExtensionId

			deleteMetricExtensionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "stack_monitoring")
			_, error := stackMonitoringClient.DeleteMetricExtension(context.Background(), deleteMetricExtensionRequest)
			if error != nil {
				fmt.Printf("Error deleting MetricExtension %s %s, It is possible that the resource is already deleted. Please verify manually \n", metricExtensionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &metricExtensionId, StackMonitoringMetricExtensionSweepWaitCondition, time.Duration(3*time.Minute),
				StackMonitoringMetricExtensionSweepResponseFetchOperation, "stack_monitoring", true)
		}
	}
	return nil
}

func getStackMonitoringMetricExtensionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MetricExtensionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	stackMonitoringClient := acctest.GetTestClients(&schema.ResourceData{}).StackMonitoringClient()

	listMetricExtensionsRequest := oci_stack_monitoring.ListMetricExtensionsRequest{}
	listMetricExtensionsRequest.CompartmentId = &compartmentId
	listMetricExtensionsRequest.LifecycleState = oci_stack_monitoring.ListMetricExtensionsLifecycleStateActive
	listMetricExtensionsResponse, err := stackMonitoringClient.ListMetricExtensions(context.Background(), listMetricExtensionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MetricExtension list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, metricExtension := range listMetricExtensionsResponse.Items {
		id := *metricExtension.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MetricExtensionId", id)
	}
	return resourceIds, nil
}

func StackMonitoringMetricExtensionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if metricExtensionResponse, ok := response.Response.(oci_stack_monitoring.GetMetricExtensionResponse); ok {
		return metricExtensionResponse.LifecycleState != oci_stack_monitoring.MetricExtensionLifeCycleStatesDeleted
	}
	return false
}

func StackMonitoringMetricExtensionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.StackMonitoringClient().GetMetricExtension(context.Background(), oci_stack_monitoring.GetMetricExtensionRequest{
		MetricExtensionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
