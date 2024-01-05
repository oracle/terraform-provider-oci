// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_network_firewall "github.com/oracle/oci-go-sdk/v65/networkfirewall"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	createMappedSecretResourceConfig = mappedSecretDependencies + createMappedSecretResource

	createMappedSecretResource = acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy_mapped_secret",
		"test_network_firewall_policy_mapped_secret",
		acctest.Required, acctest.Create,
		mappedSecretRepresentation,
	)

	mappedSecretResourceConfig = mappedSecretDependencies + acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy_mapped_secret",
		"test_network_firewall_policy_mapped_secret",
		acctest.Optional, acctest.Update,
		mappedSecretRepresentation,
	)

	mappedSecretSingularDataSourceRepresentation = map[string]interface{}{
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy_mapped_secret.test_network_firewall_policy_mapped_secret.name}`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
	}

	mappedSecretDataSourceRepresentation = map[string]interface{}{
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
	}

	mappedSecretRepresentation = map[string]interface{}{
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `mapped_secret_1`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
		"source":                     acctest.Representation{RepType: acctest.Required, Create: `OCI_VAULT`, Update: `OCI_VAULT`},
		"type":                       acctest.Representation{RepType: acctest.Required, Create: `SSL_INBOUND_INSPECTION`, Update: `SSL_INBOUND_INSPECTION`},
		"vault_secret_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_vault_secret.test_secret.id}`},
		"version_number":             acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	}

	mappedSecretDependencies = vaultSecretResource + acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy",
		"test_network_firewall_policy",
		acctest.Required, acctest.Create,
		networkFirewallPolicyRepresentation,
	)

	vaultSecretResource = acctest.GenerateResourceFromRepresentationMap(
		"oci_vault_secret",
		"test_secret",
		acctest.Required, acctest.Create,
		VaultSecretRepresentation,
	)
)

// issue-routing-tag: network_firewall/default
func TestNetworkFirewallNetworkFirewallPolicyMappedSecretResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkFirewallNetworkFirewallPolicyMappedSecretResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	vaultId := utils.GetEnvSettingWithBlankDefault("kms_vault_id")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	keyId := utils.GetEnvSettingWithBlankDefault("kms_key_id")
	keyIdVariableStr := fmt.Sprintf("variable \"key_id\" { default = \"%s\" }\n", keyId)

	resourceName := "oci_network_firewall_network_firewall_policy_mapped_secret.test_network_firewall_policy_mapped_secret"
	datasourceName := "data.oci_network_firewall_network_firewall_policy_mapped_secrets.test_network_firewall_policy_mapped_secrets"
	singularDatasourceName := "data.oci_network_firewall_network_firewall_policy_mapped_secret.test_network_firewall_policy_mapped_secret"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+createMappedSecretResourceConfig, "networkfirewall", "networkFirewallPolicyMappedSecret", t)

	configWithVariables := config + vaultIdVariableStr + keyIdVariableStr + compartmentIdVariableStr

	acctest.ResourceTest(t, testAccCheckNetworkFirewallNetworkFirewallPolicyMappedSecretDestroy, []resource.TestStep{
		// verify Create
		{
			Config: configWithVariables + createMappedSecretResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "name", "mapped_secret_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "source", "OCI_VAULT"),
				resource.TestCheckResourceAttr(resourceName, "type", "SSL_INBOUND_INSPECTION"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "version_number", "10"),

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

		// verify updates to updatable parameters
		{
			Config: configWithVariables + mappedSecretResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "name", "mapped_secret_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "source", "OCI_VAULT"),
				resource.TestCheckResourceAttr(resourceName, "type", "SSL_INBOUND_INSPECTION"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "version_number", "11"),

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
			Config: configWithVariables + acctest.GenerateDataSourceFromRepresentationMap(
				"oci_network_firewall_network_firewall_policy_mapped_secrets",
				"test_network_firewall_policy_mapped_secrets",
				acctest.Optional, acctest.Update,
				mappedSecretDataSourceRepresentation,
			) + mappedSecretResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "network_firewall_policy_id"),

				resource.TestCheckResourceAttr(datasourceName, "mapped_secret_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "mapped_secret_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: configWithVariables +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_network_firewall_network_firewall_policy_mapped_secret",
					"test_network_firewall_policy_mapped_secret",
					acctest.Required, acctest.Create,
					mappedSecretSingularDataSourceRepresentation,
				) + mappedSecretResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_firewall_policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "name", "mapped_secret_1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parent_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vault_secret_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source", "OCI_VAULT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "SSL_INBOUND_INSPECTION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "version_number", "11"),
			),
		},
		// verify resource import
		{
			Config:                  config + mappedSecretResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckNetworkFirewallNetworkFirewallPolicyMappedSecretDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).NetworkFirewallClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_network_firewall_network_firewall_policy_mapped_secret" {
			noResourceFound = false
			request := oci_network_firewall.GetMappedSecretRequest{}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.MappedSecretName = &value
			}

			if value, ok := rs.Primary.Attributes["network_firewall_policy_id"]; ok {
				request.NetworkFirewallPolicyId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")

			_, err := client.GetMappedSecret(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("NetworkFirewallNetworkFirewallPolicyMappedSecret") {
		resource.AddTestSweepers("NetworkFirewallNetworkFirewallPolicyMappedSecret", &resource.Sweeper{
			Name:         "NetworkFirewallNetworkFirewallPolicyMappedSecret",
			Dependencies: acctest.DependencyGraph["networkFirewallPolicyMappedSecret"],
			F:            sweepNetworkFirewallNetworkFirewallPolicyMappedSecretResource,
		})
	}
}

func sweepNetworkFirewallNetworkFirewallPolicyMappedSecretResource(compartment string) error {
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()
	networkFirewallPolicyMappedSecretIds, err := getNetworkFirewallNetworkFirewallPolicyMappedSecretIds(compartment)
	if err != nil {
		return err
	}
	for _, networkFirewallPolicyMappedSecretId := range networkFirewallPolicyMappedSecretIds {
		if ok := acctest.SweeperDefaultResourceId[networkFirewallPolicyMappedSecretId]; !ok {
			deleteMappedSecretRequest := oci_network_firewall.DeleteMappedSecretRequest{}

			deleteMappedSecretRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")
			_, error := networkFirewallClient.DeleteMappedSecret(context.Background(), deleteMappedSecretRequest)
			if error != nil {
				fmt.Printf("Error deleting NetworkFirewallPolicyMappedSecret %s %s, It is possible that the resource is already deleted. Please verify manually \n", networkFirewallPolicyMappedSecretId, error)
				continue
			}
		}
	}
	return nil
}

func getNetworkFirewallNetworkFirewallPolicyMappedSecretIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NetworkFirewallPolicyMappedSecretId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()

	listMappedSecretsRequest := oci_network_firewall.ListMappedSecretsRequest{}

	networkFirewallPolicyIds, error := getNetworkFirewallPolicyIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting networkFirewallPolicyId required for NetworkFirewallPolicyMappedSecret resource requests \n")
	}
	for _, networkFirewallPolicyId := range networkFirewallPolicyIds {
		listMappedSecretsRequest.NetworkFirewallPolicyId = &networkFirewallPolicyId

		listMappedSecretsResponse, err := networkFirewallClient.ListMappedSecrets(context.Background(), listMappedSecretsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting NetworkFirewallPolicyMappedSecret list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, networkFirewallPolicyMappedSecret := range listMappedSecretsResponse.Items {
			id := *networkFirewallPolicyMappedSecret.Name
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NetworkFirewallPolicyMappedSecretId", id)
		}

	}
	return resourceIds, nil
}
