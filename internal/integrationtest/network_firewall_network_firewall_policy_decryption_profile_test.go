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
	createDecryptionProfileResourceConfig = decryptionProfileResourceDependencies + createDecryptionProfileResource

	createDecryptionProfileResource = acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy_decryption_profile",
		"test_network_firewall_policy_decryption_profile",
		acctest.Required, acctest.Create,
		decryptionProfileRepresentation,
	)

	createDecryptionProfileOptionalResourceConfig = decryptionProfileResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap(
			"oci_network_firewall_network_firewall_policy_decryption_profile",
			"test_network_firewall_policy_decryption_profile",
			acctest.Optional, acctest.Create,
			decryptionProfileRepresentation,
		)

	decryptionProfileResourceConfig = decryptionProfileResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap(
			"oci_network_firewall_network_firewall_policy_decryption_profile",
			"test_network_firewall_policy_decryption_profile",
			acctest.Optional, acctest.Update,
			decryptionProfileRepresentation,
		)

	decryptionProfileSingularDataSourceRepresentation = map[string]interface{}{
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy_decryption_profile.test_network_firewall_policy_decryption_profile.name}`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
	}

	decryptionProfileDataSourceRepresentation = map[string]interface{}{
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
	}

	decryptionProfileRepresentation = map[string]interface{}{
		"name":                           acctest.Representation{RepType: acctest.Required, Create: `decryption_profile_1`},
		"network_firewall_policy_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
		"type":                           acctest.Representation{RepType: acctest.Required, Create: `SSL_INBOUND_INSPECTION`, Update: `SSL_INBOUND_INSPECTION`},
		"is_out_of_capacity_blocked":     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_unsupported_cipher_blocked":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_unsupported_version_blocked": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	decryptionProfileResourceDependencies = acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy",
		"test_network_firewall_policy",
		acctest.Required, acctest.Create,
		networkFirewallPolicyRepresentation,
	)
)

// issue-routing-tag: network_firewall/default
func TestNetworkFirewallNetworkFirewallPolicyDecryptionProfileResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkFirewallNetworkFirewallPolicyDecryptionProfileResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_network_firewall_network_firewall_policy_decryption_profile.test_network_firewall_policy_decryption_profile"
	datasourceName := "data.oci_network_firewall_network_firewall_policy_decryption_profiles.test_network_firewall_policy_decryption_profiles"
	singularDatasourceName := "data.oci_network_firewall_network_firewall_policy_decryption_profile.test_network_firewall_policy_decryption_profile"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+decryptionProfileResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap(
			"oci_network_firewall_network_firewall_policy_decryption_profile",
			"test_network_firewall_policy_decryption_profile",
			acctest.Optional, acctest.Create,
			decryptionProfileRepresentation,
		), "networkfirewall", "networkFirewallPolicyDecryptionProfile", t)

	acctest.ResourceTest(t, testAccCheckNetworkFirewallNetworkFirewallPolicyDecryptionProfileDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + createDecryptionProfileResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "name", "decryption_profile_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "SSL_INBOUND_INSPECTION"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + decryptionProfileResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + createDecryptionProfileOptionalResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "is_out_of_capacity_blocked", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_unsupported_cipher_blocked", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_unsupported_version_blocked", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "decryption_profile_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "SSL_INBOUND_INSPECTION"),

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
			Config: config + compartmentIdVariableStr + decryptionProfileResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "is_out_of_capacity_blocked", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_unsupported_cipher_blocked", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_unsupported_version_blocked", "true"),
				resource.TestCheckResourceAttr(resourceName, "name", "decryption_profile_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "SSL_INBOUND_INSPECTION"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_network_firewall_network_firewall_policy_decryption_profiles", "test_network_firewall_policy_decryption_profiles", acctest.Optional, acctest.Update, decryptionProfileDataSourceRepresentation) +
				compartmentIdVariableStr + decryptionProfileResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall_policy_decryption_profile", "test_network_firewall_policy_decryption_profile", acctest.Optional, acctest.Update, decryptionProfileRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "network_firewall_policy_id"),

				resource.TestCheckResourceAttr(datasourceName, "decryption_profile_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "decryption_profile_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_network_firewall_network_firewall_policy_decryption_profile", "test_network_firewall_policy_decryption_profile", acctest.Required, acctest.Create, decryptionProfileSingularDataSourceRepresentation) +
				compartmentIdVariableStr + decryptionProfileResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_firewall_policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "is_out_of_capacity_blocked", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_unsupported_version_blocked", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_unsupported_cipher_blocked", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "decryption_profile_1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parent_resource_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "SSL_INBOUND_INSPECTION"),
			),
		},
		// verify resource import
		{
			Config:                  config + createDecryptionProfileResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckNetworkFirewallNetworkFirewallPolicyDecryptionProfileDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).NetworkFirewallClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_network_firewall_network_firewall_policy_decryption_profile" {
			noResourceFound = false
			request := oci_network_firewall.GetDecryptionProfileRequest{}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.DecryptionProfileName = &value
			}

			if value, ok := rs.Primary.Attributes["network_firewall_policy_id"]; ok {
				request.NetworkFirewallPolicyId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")

			_, err := client.GetDecryptionProfile(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("NetworkFirewallNetworkFirewallPolicyDecryptionProfile") {
		resource.AddTestSweepers("NetworkFirewallNetworkFirewallPolicyDecryptionProfile", &resource.Sweeper{
			Name:         "NetworkFirewallNetworkFirewallPolicyDecryptionProfile",
			Dependencies: acctest.DependencyGraph["networkFirewallPolicyDecryptionProfile"],
			F:            sweepNetworkFirewallNetworkFirewallPolicyDecryptionProfileResource,
		})
	}
}

func sweepNetworkFirewallNetworkFirewallPolicyDecryptionProfileResource(compartment string) error {
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()
	networkFirewallPolicyDecryptionProfileIds, err := getNetworkFirewallNetworkFirewallPolicyDecryptionProfileIds(compartment)
	if err != nil {
		return err
	}
	for _, networkFirewallPolicyDecryptionProfileId := range networkFirewallPolicyDecryptionProfileIds {
		if ok := acctest.SweeperDefaultResourceId[networkFirewallPolicyDecryptionProfileId]; !ok {
			deleteDecryptionProfileRequest := oci_network_firewall.DeleteDecryptionProfileRequest{}

			deleteDecryptionProfileRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")
			_, error := networkFirewallClient.DeleteDecryptionProfile(context.Background(), deleteDecryptionProfileRequest)
			if error != nil {
				fmt.Printf("Error deleting NetworkFirewallPolicyDecryptionProfile %s %s, It is possible that the resource is already deleted. Please verify manually \n", networkFirewallPolicyDecryptionProfileId, error)
				continue
			}
		}
	}
	return nil
}

func getNetworkFirewallNetworkFirewallPolicyDecryptionProfileIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NetworkFirewallPolicyDecryptionProfileId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()

	listDecryptionProfilesRequest := oci_network_firewall.ListDecryptionProfilesRequest{}

	networkFirewallPolicyIds, error := getNetworkFirewallPolicyIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting networkFirewallPolicyId required for NetworkFirewallPolicyDecryptionProfile resource requests \n")
	}
	for _, networkFirewallPolicyId := range networkFirewallPolicyIds {
		listDecryptionProfilesRequest.NetworkFirewallPolicyId = &networkFirewallPolicyId

		listDecryptionProfilesResponse, err := networkFirewallClient.ListDecryptionProfiles(context.Background(), listDecryptionProfilesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting NetworkFirewallPolicyDecryptionProfile list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, networkFirewallPolicyDecryptionProfile := range listDecryptionProfilesResponse.Items {
			id := *networkFirewallPolicyDecryptionProfile.Name
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NetworkFirewallPolicyDecryptionProfileId", id)
		}

	}
	return resourceIds, nil
}
