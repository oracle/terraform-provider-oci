// Copyright (c) 2025, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	oci_apm_config "github.com/oracle/oci-go-sdk/v65/apmconfig"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	ApmConfigMacsApmExtensionRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_macs_apm_extension", acctest.Required, acctest.Create, configMacsApmExtensionRepresentation)

	ApmConfigMacsApmExtensionResource = ApmConfigConfigResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_macs_apm_extension", acctest.Optional, acctest.Update, configMacsApmExtensionRepresentation)

	ApmConfigMacsApmExtensionDataResource = acctest.GenerateDataSourceFromRepresentationMap("oci_apm_config_config", "test_macs_apm_extension", acctest.Required, acctest.Create, ApmConfigMacsApmExtensionSingularDataSourceRepresentation)

	ApmConfigMacsApmExtensionSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"config_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_config_config.test_macs_apm_extension.id}`},
	}

	ApmConfigMacsApmExtensionDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"config_type":   acctest.Representation{RepType: acctest.Optional, Create: configTypeMacsApmExtension, Update: configTypeMacsApmExtension},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: ApmConfigMacsApmExtensionFilterDataSourceFilterRepresentation},
	}

	ApmConfigMacsApmExtensionFilterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${data.oci_apm_config_config.test_macs_apm_extension.id}`}},
	}

	configTypeMacsApmExtension = "MACS_APM_EXTENSION"

	configManagementAgentOcid = utils.GetEnvSettingWithBlankDefault("management_agent_ocid")

	configMacsApmExtensionRepresentation = map[string]interface{}{
		"apm_domain_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_apm_domain.test_apm_domain.id}`},
		"config_type":         acctest.Representation{RepType: acctest.Required, Create: configTypeMacsApmExtension, Update: configTypeMacsApmExtension},
		"management_agent_id": acctest.Representation{RepType: acctest.Required, Create: configManagementAgentOcid},
		"process_filter": acctest.Representation{RepType: acctest.Required, Create: []string{".*org.apache.catalina.startup.Bootstrap.*", ".*jetty.*"},
			Update: []string{".*org.jboss.*", ".*org.apache.bootstrap.*"}},
		"run_as_user":        acctest.Representation{RepType: acctest.Required, Create: "tomcat", Update: "tomcat2"},
		"service_name":       acctest.Representation{RepType: acctest.Required, Create: "Tomcat", Update: "Tomcat2"},
		"agent_version":      acctest.Representation{RepType: acctest.Required, Create: "1.16.0.585", Update: "1.16.0.586"},
		"attach_install_dir": acctest.Representation{RepType: acctest.Required, Create: "/opt/oracle/apm_attach_process", Update: "/opt/oracle/apm_attach_process2"},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: "Display name 1", Update: "Display name 2"},
		"defined_tags": acctest.Representation{RepType: acctest.Optional,
			Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`,
			Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}
)

// issue-routing-tag: apm_config/default
func TestApmConfigMacsApmExtensionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmConfigMacsApmExtensionResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_apm_config_config.test_macs_apm_extension"
	datasourceName := "data.oci_apm_config_configs.test_macs_apm_extensions"
	singularDatasourceName := "data.oci_apm_config_config.test_macs_apm_extension"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ApmConfigConfigResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_macs_apm_extension", acctest.Optional, acctest.Create, configMacsApmExtensionRepresentation), "apmconfig", "config", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckApmConfigMacsApmExtensionDestroy,
		Steps: []resource.TestStep{
			// Find these steps in the test log easily with "Executing step (number)"
			// Step 1: verify create Macs APM Extension
			{
				Config: config + compartmentIdVariableStr + ApmConfigConfigResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_macs_apm_extension", acctest.Required, acctest.Create, configMacsApmExtensionRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "config_type", configTypeMacsApmExtension),
					resource.TestCheckResourceAttr(resourceName, "management_agent_id", configManagementAgentOcid),
					resource.TestCheckResourceAttr(resourceName, "process_filter.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "process_filter.0", ".*org.apache.catalina.startup.Bootstrap.*"),
					resource.TestCheckResourceAttr(resourceName, "process_filter.1", ".*jetty.*"),
					resource.TestCheckResourceAttr(resourceName, "run_as_user", "tomcat"),
					resource.TestCheckResourceAttr(resourceName, "service_name", "Tomcat"),
					resource.TestCheckResourceAttr(resourceName, "agent_version", "1.16.0.585"),
					resource.TestCheckResourceAttr(resourceName, "attach_install_dir", "/opt/oracle/apm_attach_process"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_by"),
					resource.TestCheckResourceAttrSet(resourceName, "etag"),
					resource.TestCheckResourceAttr(resourceName, "in_use_by.0.items.#", "0"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						utils.Logf("This is error")
						return err
					},
				),
			},
			// Step 2: delete Macs APM Extension before next create
			{
				Config: config + compartmentIdVariableStr + ApmConfigConfigResourceDependencies,
			},
			// Step 3: verify create Macs APM Extension with optionals
			{
				Config: config + compartmentIdVariableStr + ApmConfigConfigResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_macs_apm_extension", acctest.Optional, acctest.Create, configMacsApmExtensionRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "config_type", configTypeMacsApmExtension),
					resource.TestCheckResourceAttr(resourceName, "management_agent_id", configManagementAgentOcid),
					resource.TestCheckResourceAttr(resourceName, "process_filter.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "process_filter.0", ".*org.apache.catalina.startup.Bootstrap.*"),
					resource.TestCheckResourceAttr(resourceName, "process_filter.1", ".*jetty.*"),
					resource.TestCheckResourceAttr(resourceName, "run_as_user", "tomcat"),
					resource.TestCheckResourceAttr(resourceName, "service_name", "Tomcat"),
					resource.TestCheckResourceAttr(resourceName, "agent_version", "1.16.0.585"),
					resource.TestCheckResourceAttr(resourceName, "attach_install_dir", "/opt/oracle/apm_attach_process"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "Display name 1"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_by"),
					resource.TestCheckResourceAttrSet(resourceName, "etag"),
					resource.TestCheckResourceAttr(resourceName, "in_use_by.0.items.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),

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
			// Step 4: verify updates to Macs APM Extension updatable parameters
			{
				Config: config + compartmentIdVariableStr + ApmConfigConfigResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_apm_config_config", "test_macs_apm_extension", acctest.Optional, acctest.Update, configMacsApmExtensionRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "config_type", configTypeMacsApmExtension),
					resource.TestCheckResourceAttr(resourceName, "management_agent_id", configManagementAgentOcid),
					resource.TestCheckResourceAttr(resourceName, "process_filter.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "process_filter.0", ".*org.jboss.*"),
					resource.TestCheckResourceAttr(resourceName, "process_filter.1", ".*org.apache.bootstrap.*"),
					resource.TestCheckResourceAttr(resourceName, "run_as_user", "tomcat2"),
					resource.TestCheckResourceAttr(resourceName, "service_name", "Tomcat2"),
					resource.TestCheckResourceAttr(resourceName, "agent_version", "1.16.0.586"),
					resource.TestCheckResourceAttr(resourceName, "attach_install_dir", "/opt/oracle/apm_attach_process2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "Display name 2"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_by"),
					resource.TestCheckResourceAttrSet(resourceName, "etag"),
					resource.TestCheckResourceAttr(resourceName, "in_use_by.0.items.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// Step 5: verify datasource (Macs APM Extension)
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_apm_config_configs", "test_macs_apm_extensions", acctest.Optional, acctest.Update, ApmConfigMacsApmExtensionDataSourceRepresentation) +
					compartmentIdVariableStr + ApmConfigMacsApmExtensionResource + ApmConfigMacsApmExtensionDataResource,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
					resource.TestCheckResourceAttr(datasourceName, "config_type", configTypeMacsApmExtension),
					resource.TestCheckResourceAttr(datasourceName, "config_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "config_collection.0.items.#", "1"),
				),
			},
			// Step 6: verify singular datasource (Macs APM Extension)
			{
				Config: config + compartmentIdVariableStr + ApmConfigMacsApmExtensionResource + ApmConfigMacsApmExtensionDataResource,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "config_type", configTypeMacsApmExtension),
					resource.TestCheckResourceAttr(singularDatasourceName, "management_agent_id", configManagementAgentOcid),
					resource.TestCheckResourceAttr(singularDatasourceName, "process_filter.#", "2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "process_filter.0", ".*org.jboss.*"),
					resource.TestCheckResourceAttr(singularDatasourceName, "process_filter.1", ".*org.apache.bootstrap.*"),
					resource.TestCheckResourceAttr(singularDatasourceName, "run_as_user", "tomcat2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "service_name", "Tomcat2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "agent_version", "1.16.0.586"),
					resource.TestCheckResourceAttr(singularDatasourceName, "attach_install_dir", "/opt/oracle/apm_attach_process2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "Display name 2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "updated_by"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "etag"),
					resource.TestCheckResourceAttr(singularDatasourceName, "in_use_by.0.items.#", "0"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				),
			},
			// Step 7: remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + ApmConfigMacsApmExtensionResource,
			},
			// Step 8: verify resource import
			{
				Config:            config + ApmConfigMacsApmExtensionRequiredOnlyResource,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"apm_domain_id",
					"opc_dry_run",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckApmConfigMacsApmExtensionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ConfigClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_apm_config_config" {
			noResourceFound = false
			request := oci_apm_config.GetConfigRequest{}

			if value, ok := rs.Primary.Attributes["apm_domain_id"]; ok {
				request.ApmDomainId = &value
			}

			tmp, _ := parseConfigCompositeId(rs.Primary.ID)
			request.ConfigId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apm_config")

			_, err := client.GetConfig(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !acctest.InSweeperExcludeList("ApmConfigMacsApmExtension") {
		resource.AddTestSweepers("ApmConfigMacsApmExtension", &resource.Sweeper{
			Name:         "ApmConfigMacsApmExtension",
			Dependencies: acctest.DependencyGraph["config"],
			F:            sweepApmConfigMacsApmExtensionResource,
		})
	}
}

func sweepApmConfigMacsApmExtensionResource(compartment string) error {
	configClient := acctest.GetTestClients(&schema.ResourceData{}).ConfigClient()
	configIds, err := getMacsApmExtensionIds(compartment)
	if err != nil {
		return err
	}
	for _, configId := range configIds {
		if ok := acctest.SweeperDefaultResourceId[configId]; !ok {
			deleteConfigRequest := oci_apm_config.DeleteConfigRequest{}

			deleteConfigRequest.ConfigId = &configId

			deleteConfigRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apm_config")
			_, error := configClient.DeleteConfig(context.Background(), deleteConfigRequest)
			if error != nil {
				fmt.Printf("Error deleting Config %s %s, It is possible that the resource is already deleted. Please verify manually \n", configId, error)
				continue
			}
		}
	}
	return nil
}

func getMacsApmExtensionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ConfigId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	configClient := acctest.GetTestClients(&schema.ResourceData{}).ConfigClient()

	listConfigsRequest := oci_apm_config.ListConfigsRequest{}
	//listConfigsRequest.CompartmentId = &compartmentId

	apmDomainIds, error := getApmDomainIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting apmDomainId required for Config resource requests \n")
	}
	for _, apmDomainId := range apmDomainIds {
		listConfigsRequest.ApmDomainId = &apmDomainId

		listConfigsResponse, err := configClient.ListConfigs(context.Background(), listConfigsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Config list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, config := range listConfigsResponse.Items {
			id := *config.GetId()
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ConfigId", id)
		}

	}
	return resourceIds, nil
}
