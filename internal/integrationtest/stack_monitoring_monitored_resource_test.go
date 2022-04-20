// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

/**
  Dependency variables:
      hostname = var.stack_mon_hostname_resource1
      management_agent_id = var.stack_mon_management_agent_id_resource1
      hostname2 = var.stack_mon_hostname_resource2
      management_agent_id2 = var.stack_mon_management_agent_id_resource2
*/
var (
	MonitoredResourceRequiredOnlyResource = MonitoredResourceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource", "test_monitored_resource", acctest.Required, acctest.Create, monitoredResourceRepresentation)

	MonitoredResourceResourceConfig = MonitoredResourceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource", "test_monitored_resource", acctest.Optional, acctest.Update, monitoredResourceRepresentation)

	monitoredResourceSingularDataSourceRepresentation = map[string]interface{}{
		"monitored_resource_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_stack_monitoring_monitored_resource.test_monitored_resource.id}`},
	}

	monitoredResourceRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":                        acctest.Representation{RepType: acctest.Required, Create: `terraformResource`},
		"type":                        acctest.Representation{RepType: acctest.Required, Create: `host`},
		"display_name":                acctest.Representation{RepType: acctest.Optional, Create: `displayNameTerra`, Update: `displayNameTerra2`},
		"host_name":                   acctest.Representation{RepType: acctest.Optional, Create: `${var.stack_mon_hostname_resource1}`},
		"management_agent_id":         acctest.Representation{RepType: acctest.Optional, Create: `${var.stack_mon_management_agent_id_resource1}`},
		"credentials":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: monitoredResourceCredentialsRepresentation},
		"database_connection_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: monitoredResourceDatabaseConnectionDetailsRepresentation},
		"properties":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: monitoredResourcePropertiesRepresentation},
		"resource_time_zone":          acctest.Representation{RepType: acctest.Optional, Create: `en`},
		"lifecycle":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSensitiveDataRepresentation},
	}
	//Get API does not return sensitive data, it returns null
	ignoreSensitiveDataRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`credentials`, `database_connection_details`}},
	}

	monitoredResourceRepresentation2 = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":                acctest.Representation{RepType: acctest.Required, Create: `terraformSecondaryResource`},
		"type":                acctest.Representation{RepType: acctest.Required, Create: `host`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displaySecondaryNameTerra`, Update: `displaySecondaryNameTerra2`},
		"host_name":           acctest.Representation{RepType: acctest.Optional, Create: `${var.stack_mon_hostname_resource2}`},
		"management_agent_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.stack_mon_management_agent_id_resource2}`},
		"properties":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: monitoredResourcePropertiesRepresentation},
		"resource_time_zone":  acctest.Representation{RepType: acctest.Optional, Create: `en`},
	}

	monitoredResourceAliasesRepresentation = map[string]interface{}{
		"credential": acctest.RepresentationGroup{RepType: acctest.Required, Group: monitoredResourceAliasesCredentialRepresentation},
		"name":       acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"source":     acctest.Representation{RepType: acctest.Required, Create: `source`, Update: `source2`},
	}
	monitoredResourceCredentialsRepresentation = map[string]interface{}{
		"credential_type": acctest.Representation{RepType: acctest.Optional, Create: `PLAINTEXT`, Update: `PLAINTEXT`},
		"description":     acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"key_id":          acctest.Representation{RepType: acctest.Optional, Create: `somekeyid`},
		"name":            acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"properties":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: monitoredResourceCredentialsPropertiesRepresentation},
		"source":          acctest.Representation{RepType: acctest.Optional, Create: `host.terraformName`},
		"type":            acctest.Representation{RepType: acctest.Optional, Create: `type`, Update: `type2`},
	}
	monitoredResourceDatabaseConnectionDetailsRepresentation = map[string]interface{}{
		"port":           acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"protocol":       acctest.Representation{RepType: acctest.Required, Create: `TCP`, Update: `TCPS`},
		"service_name":   acctest.Representation{RepType: acctest.Required, Create: `service.name`},
		"connector_id":   acctest.Representation{RepType: acctest.Optional, Create: `connector.id`},
		"db_id":          acctest.Representation{RepType: acctest.Optional, Create: `db_id`},
		"db_unique_name": acctest.Representation{RepType: acctest.Optional, Create: `dbUniqueName`, Update: `dbUniqueName2`},
	}
	monitoredResourcePropertiesRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Optional, Create: `OS`, Update: `OS`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `Linux`, Update: `Linux`},
	}
	monitoredResourceAliasesCredentialRepresentation = map[string]interface{}{
		"name":    acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"service": acctest.Representation{RepType: acctest.Required, Create: `service`, Update: `service2`},
		"source":  acctest.Representation{RepType: acctest.Required, Create: `source`, Update: `source2`},
	}
	monitoredResourceCredentialsPropertiesRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `value`, Update: `value2`},
	}

	MonitoredResourceResourceDependencies = ""
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringMonitoredResourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringMonitoredResourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	managementAgentId1 := utils.GetEnvSettingWithBlankDefault("stack_mon_management_agent_id_resource1")
	if managementAgentId1 == "" {
		t.Skip("Setting environmental variable stack_mon_management_agent_id_resource1 that represents management agent with resource monitoring plugin is pre-requisite for this test")
	}
	managementAgentId1VariableStr := fmt.Sprintf("variable \"stack_mon_management_agent_id_resource1\" { default = \"%s\" }\n", managementAgentId1)

	hostname1 := utils.GetEnvSettingWithBlankDefault("stack_mon_hostname_resource1")
	if hostname1 == "" {
		t.Skip("Setting environmental variable stack_mon_hostname_resource1 that host accessible by agent defined by stack_mon_management_agent_id_resource1 variable is pre-requisite for this test")
	}
	hostname1VariableStr := fmt.Sprintf("variable \"stack_mon_hostname_resource1\" { default = \"%s\" }\n", hostname1)

	resourceName := "oci_stack_monitoring_monitored_resource.test_monitored_resource"

	singularDatasourceName := "data.oci_stack_monitoring_monitored_resource.test_monitored_resource"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MonitoredResourceResourceDependencies+managementAgentId1VariableStr+hostname1VariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource", "test_monitored_resource", acctest.Optional, acctest.Create, monitoredResourceRepresentation), "stackmonitoring", "monitoredResource", t)

	acctest.ResourceTest(t, testAccCheckStackMonitoringMonitoredResourceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + MonitoredResourceResourceDependencies + managementAgentId1VariableStr + hostname1VariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource", "test_monitored_resource", acctest.Optional, acctest.Create, monitoredResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "terraformResource"),
				resource.TestCheckResourceAttr(resourceName, "type", "host"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + MonitoredResourceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + MonitoredResourceResourceDependencies + managementAgentId1VariableStr + hostname1VariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource", "test_monitored_resource", acctest.Optional, acctest.Create, monitoredResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayNameTerra"),
				resource.TestCheckResourceAttrSet(resourceName, "host_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_agent_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "terraformResource"),
				resource.TestCheckResourceAttr(resourceName, "properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "properties.0.name", "OS"),
				resource.TestCheckResourceAttr(resourceName, "properties.0.value", "Linux"),
				resource.TestCheckResourceAttr(resourceName, "resource_time_zone", "en"),
				resource.TestCheckResourceAttr(resourceName, "type", "host"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + MonitoredResourceResourceDependencies + managementAgentId1VariableStr + hostname1VariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource", "test_monitored_resource", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(monitoredResourceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayNameTerra"),
				resource.TestCheckResourceAttrSet(resourceName, "host_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_agent_id"),
				resource.TestCheckResourceAttr(resourceName, "properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "properties.0.name", "OS"),
				resource.TestCheckResourceAttr(resourceName, "properties.0.value", "Linux"),
				resource.TestCheckResourceAttr(resourceName, "resource_time_zone", "en"),
				resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "host"),

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
			Config: config + compartmentIdVariableStr + MonitoredResourceResourceDependencies + managementAgentId1VariableStr + hostname1VariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource", "test_monitored_resource", acctest.Optional, acctest.Update, monitoredResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayNameTerra2"),
				resource.TestCheckResourceAttrSet(resourceName, "host_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "management_agent_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "terraformResource"),
				resource.TestCheckResourceAttr(resourceName, "properties.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "properties.0.name", "OS"),
				resource.TestCheckResourceAttr(resourceName, "properties.0.value", "Linux"),
				resource.TestCheckResourceAttr(resourceName, "resource_time_zone", "en"),
				resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "host"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_monitored_resource", "test_monitored_resource", acctest.Required, acctest.Create, monitoredResourceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr + MonitoredResourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "monitored_resource_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayNameTerra2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "terraformResource"),
				resource.TestCheckResourceAttr(singularDatasourceName, "properties.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "properties.0.name", "OS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "properties.0.value", "Linux"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_time_zone", "en"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenant_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "host"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr + MonitoredResourceResourceConfig,
		},
		// verify resource import
		{
			Config:            config + compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr + MonitoredResourceResourceConfig,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"external_resource_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckStackMonitoringMonitoredResourceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).StackMonitoringClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_stack_monitoring_monitored_resource" {
			noResourceFound = false
			request := oci_stack_monitoring.GetMonitoredResourceRequest{}

			tmp := rs.Primary.ID
			request.MonitoredResourceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "stack_monitoring")

			response, err := client.GetMonitoredResource(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_stack_monitoring.ResourceLifecycleStateDeleted): true,
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
