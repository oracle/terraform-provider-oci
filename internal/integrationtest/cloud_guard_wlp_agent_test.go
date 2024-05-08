// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CloudGuardWlpAgentRequiredOnlyResource = CloudGuardWlpAgentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_wlp_agent", "test_wlp_agent", acctest.Required, acctest.Create, CloudGuardWlpAgentRepresentation)

	CloudGuardWlpAgentResourceConfig = CloudGuardWlpAgentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_wlp_agent", "test_wlp_agent", acctest.Optional, acctest.Update, CloudGuardWlpAgentRepresentation)

	CloudGuardWlpAgentSingularDataSourceRepresentation = map[string]interface{}{
		"wlp_agent_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_guard_wlp_agent.test_wlp_agent.id}`},
	}

	CloudGuardWlpAgentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudGuardWlpAgentDataSourceFilterRepresentation}}
	CloudGuardWlpAgentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_cloud_guard_wlp_agent.test_wlp_agent.id}`}},
	}

	CloudGuardWlpAgentRepresentation = map[string]interface{}{
		"agent_version":              acctest.Representation{RepType: acctest.Required, Create: "1.0.147"},
		"certificate_signed_request": acctest.Representation{RepType: acctest.Required, Create: `-----BEGIN CERTIFICATE REQUEST-----\nMIID1DCCArwCAQAwggGNMQswCQYDVQQGEwJVUzEPMA0GA1UEChMGT3JhY2xlMYIB\nKDCCASQGA1UECxOCARtvcGMtaW5zdGFuY2U6b2NpZDEuaW5zdGFuY2Uub2MxLnBo\neC5hbnlocWxqdDd4bTQ1Y2ljeWkya2E1cHVtcWd5dWhkZXVhaHJjYXh6NHN3bGtv\nZmo2dXhjdmtubnhkaGEsb3BjLWNvbXBhcnRtZW50Om9jaWQxLnRlbmFuY3kub2Mx\nLi5hYWFhYWFhYXFvZ2d6c2p1dDJ1NjR3cWxpeWQ0ZXlkM2RsNGlwc3UyNmxncXg0\nYmlob2ZudmU1bGk1aHEsb3BjLXRlbmFudDpvY2lkMS50ZW5hbmN5Lm9jMS4uYWFh\nYWFhYWFxb2dnenNqdXQydTY0d3FsaXlkNGV5ZDNkbDRpcHN1MjZsZ3F4NGJpaG9m\nbnZlNWxpNWhxMQwwCgYDVQQDEwN3bHAxMzAxBgkqhkiG9w0BCQEWJHdvcmtsb2Fk\ncHJvdGVjdGlvbl91c19ncnBAb3JhY2xlLmNvbTCCASIwDQYJKoZIhvcNAQEBBQAD\nggEPADCCAQoCggEBAMDLbqoECIIh02HvkusRyGGI/cqK9Wrg7xDn/Wwg1C9noOo+\nbHmU5sBervLUHKXuC3IUwM0GgytjLsOjMWI9ex0ZunQONwwAe/MDD+YQcnqbOnmb\naUrdp0gB231SRqCUST1xf9y8shlK3zXrav+qgtF1bDihsGh6O4DMLPYIsOZAXo6M\nrGPokj1nViLdvFaBBG4Q1sgximufh/eqFCaUawIUOeQ7XcDqeWM+G8IA3vIuWqbr\nSoI61/COgq6eDsUMu/ZcMNF0UYRV4bWwVM18Cx8Tlp0kH/mbnlHxBMxz1x/cbHmQ\nEwPrSKWo8Gn2B1HeXWhVGNPa4Xs0xn/kaW1QaS8CAwEAAaAAMA0GCSqGSIb3DQEB\nCwUAA4IBAQABiABQPOngTCA24KzY6GcyVi/4H6nhOu6smAgnPM2PoJEoog5yvnLR\nTvoyec0TTIIiRZtDIYejRMUyGZxR1o1Hgrkq80OmqfRZW57e2WPRgpHcp87Yfp0B\nRmkobQMRSAypZDGCdco2cuQ4F7GG0KFMb1Tf+b/XQnf6L3cd9PCHPECOVe1LFJV3\nqxhNkkxd+REI8iihLjzslqJFufYTkfmL2xamhS2nzGbG5XcfURdqx6S2ZDVoCkNy\nikohM9PlBrWAXWYALRqgcy1KFH9lQ9+tIqpnGbOHOyIqFPmoMKX2ugisTWMpgTp9\nxICh2HMz77KABXXf/t58HDODI4Wx8yJA\n-----END CERTIFICATE REQUEST-----\n`, Update: `-----BEGIN CERTIFICATE REQUEST-----\nMIID1DCCArwCAQAwggGNMQswCQYDVQQGEwJVUzEPMA0GA1UEChMGT3JhY2xlMYIB\nKDCCASQGA1UECxOCARtvcGMtaW5zdGFuY2U6b2NpZDEuaW5zdGFuY2Uub2MxLnBo\neC5hbnlocWxqdDd4bTQ1Y2ljeWkya2E1cHVtcWd5dWhkZXVhaHJjYXh6NHN3bGtv\nZmo2dXhjdmtubnhkaGEsb3BjLWNvbXBhcnRtZW50Om9jaWQxLnRlbmFuY3kub2Mx\nLi5hYWFhYWFhYXFvZ2d6c2p1dDJ1NjR3cWxpeWQ0ZXlkM2RsNGlwc3UyNmxncXg0\nYmlob2ZudmU1bGk1aHEsb3BjLXRlbmFudDpvY2lkMS50ZW5hbmN5Lm9jMS4uYWFh\nYWFhYWFxb2dnenNqdXQydTY0d3FsaXlkNGV5ZDNkbDRpcHN1MjZsZ3F4NGJpaG9m\nbnZlNWxpNWhxMQwwCgYDVQQDEwN3bHAxMzAxBgkqhkiG9w0BCQEWJHdvcmtsb2Fk\ncHJvdGVjdGlvbl91c19ncnBAb3JhY2xlLmNvbTCCASIwDQYJKoZIhvcNAQEBBQAD\nggEPADCCAQoCggEBAMDLbqoECIIh02HvkusRyGGI/cqK9Wrg7xDn/Wwg1C9noOo+\nbHmU5sBervLUHKXuC3IUwM0GgytjLsOjMWI9ex0ZunQONwwAe/MDD+YQcnqbOnmb\naUrdp0gB231SRqCUST1xf9y8shlK3zXrav+qgtF1bDihsGh6O4DMLPYIsOZAXo6M\nrGPokj1nViLdvFaBBG4Q1sgximufh/eqFCaUawIUOeQ7XcDqeWM+G8IA3vIuWqbr\nSoI61/COgq6eDsUMu/ZcMNF0UYRV4bWwVM18Cx8Tlp0kH/mbnlHxBMxz1x/cbHmQ\nEwPrSKWo8Gn2B1HeXWhVGNPa4Xs0xn/kaW1QaS8CAwEAAaAAMA0GCSqGSIb3DQEB\nCwUAA4IBAQABiABQPOngTCA24KzY6GcyVi/4H6nhOu6smAgnPM2PoJEoog5yvnLR\nTvoyec0TTIIiRZtDIYejRMUyGZxR1o1Hgrkq80OmqfRZW57e2WPRgpHcp87Yfp0B\nRmkobQMRSAypZDGCdco2cuQ4F7GG0KFMb1Tf+b/XQnf6L3cd9PCHPECOVe1LFJV3\nqxhNkkxd+REI8iihLjzslqJFufYTkfmL2xamhS2nzGbG5XcfURdqx6S2ZDVoCkNy\nikohM9PlBrWAXWYALRqgcy1KFH9lQ9+tIqpnGbOHOyIqFPmoMKX2ugisTWMpgTp9\nxICh2HMz77KABXXf/t58HDODI4Wx8yJA\n-----END CERTIFICATE REQUEST-----\n`},
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"os_info":                    acctest.Representation{RepType: acctest.Required, Create: `Oracle Linux Server_8.5_amd64`},
	}

	CloudGuardWlpAgentResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: cloud_guard/default
func TestCloudGuardWlpAgentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudGuardWlpAgentResource_basic")
	defer httpreplay.SaveScenario()
	csr := "-----BEGIN CERTIFICATE REQUEST-----\nMIID1DCCArwCAQAwggGNMQswCQYDVQQGEwJVUzEPMA0GA1UEChMGT3JhY2xlMYIB\nKDCCASQGA1UECxOCARtvcGMtaW5zdGFuY2U6b2NpZDEuaW5zdGFuY2Uub2MxLnBo\neC5hbnlocWxqdDd4bTQ1Y2ljeWkya2E1cHVtcWd5dWhkZXVhaHJjYXh6NHN3bGtv\nZmo2dXhjdmtubnhkaGEsb3BjLWNvbXBhcnRtZW50Om9jaWQxLnRlbmFuY3kub2Mx\nLi5hYWFhYWFhYXFvZ2d6c2p1dDJ1NjR3cWxpeWQ0ZXlkM2RsNGlwc3UyNmxncXg0\nYmlob2ZudmU1bGk1aHEsb3BjLXRlbmFudDpvY2lkMS50ZW5hbmN5Lm9jMS4uYWFh\nYWFhYWFxb2dnenNqdXQydTY0d3FsaXlkNGV5ZDNkbDRpcHN1MjZsZ3F4NGJpaG9m\nbnZlNWxpNWhxMQwwCgYDVQQDEwN3bHAxMzAxBgkqhkiG9w0BCQEWJHdvcmtsb2Fk\ncHJvdGVjdGlvbl91c19ncnBAb3JhY2xlLmNvbTCCASIwDQYJKoZIhvcNAQEBBQAD\nggEPADCCAQoCggEBAMDLbqoECIIh02HvkusRyGGI/cqK9Wrg7xDn/Wwg1C9noOo+\nbHmU5sBervLUHKXuC3IUwM0GgytjLsOjMWI9ex0ZunQONwwAe/MDD+YQcnqbOnmb\naUrdp0gB231SRqCUST1xf9y8shlK3zXrav+qgtF1bDihsGh6O4DMLPYIsOZAXo6M\nrGPokj1nViLdvFaBBG4Q1sgximufh/eqFCaUawIUOeQ7XcDqeWM+G8IA3vIuWqbr\nSoI61/COgq6eDsUMu/ZcMNF0UYRV4bWwVM18Cx8Tlp0kH/mbnlHxBMxz1x/cbHmQ\nEwPrSKWo8Gn2B1HeXWhVGNPa4Xs0xn/kaW1QaS8CAwEAAaAAMA0GCSqGSIb3DQEB\nCwUAA4IBAQABiABQPOngTCA24KzY6GcyVi/4H6nhOu6smAgnPM2PoJEoog5yvnLR\nTvoyec0TTIIiRZtDIYejRMUyGZxR1o1Hgrkq80OmqfRZW57e2WPRgpHcp87Yfp0B\nRmkobQMRSAypZDGCdco2cuQ4F7GG0KFMb1Tf+b/XQnf6L3cd9PCHPECOVe1LFJV3\nqxhNkkxd+REI8iihLjzslqJFufYTkfmL2xamhS2nzGbG5XcfURdqx6S2ZDVoCkNy\nikohM9PlBrWAXWYALRqgcy1KFH9lQ9+tIqpnGbOHOyIqFPmoMKX2ugisTWMpgTp9\nxICh2HMz77KABXXf/t58HDODI4Wx8yJA\n-----END CERTIFICATE REQUEST-----\n"
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_cloud_guard_wlp_agent.test_wlp_agent"
	datasourceName := "data.oci_cloud_guard_wlp_agents.test_wlp_agents"
	singularDatasourceName := "data.oci_cloud_guard_wlp_agent.test_wlp_agent"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CloudGuardWlpAgentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_wlp_agent", "test_wlp_agent", acctest.Optional, acctest.Create, CloudGuardWlpAgentRepresentation), "cloudguard", "wlpAgent", t)

	acctest.ResourceTest(t, testAccCheckCloudGuardWlpAgentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CloudGuardWlpAgentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_wlp_agent", "test_wlp_agent", acctest.Required, acctest.Create, CloudGuardWlpAgentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "agent_version", "1.0.147"),
				resource.TestCheckResourceAttr(resourceName, "certificate_signed_request", csr),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "os_info", "Oracle Linux Server_8.5_amd64"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CloudGuardWlpAgentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CloudGuardWlpAgentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_wlp_agent", "test_wlp_agent", acctest.Optional, acctest.Create, CloudGuardWlpAgentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "agent_version", "1.0.147"),
				resource.TestCheckResourceAttrSet(resourceName, "certificate_id"),
				resource.TestCheckResourceAttr(resourceName, "certificate_signed_request", csr),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "os_info", "Oracle Linux Server_8.5_amd64"),

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
			Config: config + compartmentIdVariableStr + CloudGuardWlpAgentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_wlp_agent", "test_wlp_agent", acctest.Optional, acctest.Update, CloudGuardWlpAgentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "agent_version", "1.0.147"),
				resource.TestCheckResourceAttrSet(resourceName, "certificate_id"),
				resource.TestCheckResourceAttr(resourceName, "certificate_signed_request", csr),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "os_info", "Oracle Linux Server_8.5_amd64"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_wlp_agents", "test_wlp_agents", acctest.Optional, acctest.Update, CloudGuardWlpAgentDataSourceRepresentation) +
				compartmentIdVariableStr + CloudGuardWlpAgentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_guard_wlp_agent", "test_wlp_agent", acctest.Optional, acctest.Update, CloudGuardWlpAgentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "wlp_agent_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "wlp_agent_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_wlp_agent", "test_wlp_agent", acctest.Required, acctest.Create, CloudGuardWlpAgentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CloudGuardWlpAgentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "wlp_agent_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "agent_version", "1.0.147"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "certificate_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "certificate_signed_request", csr),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenant_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + CloudGuardWlpAgentRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"os_info",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckCloudGuardWlpAgentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).CloudGuardClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cloud_guard_wlp_agent" {
			noResourceFound = false
			request := oci_cloud_guard.GetWlpAgentRequest{}

			tmp := rs.Primary.ID
			request.WlpAgentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_guard")

			_, err := client.GetWlpAgent(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("CloudGuardWlpAgent") {
		resource.AddTestSweepers("CloudGuardWlpAgent", &resource.Sweeper{
			Name:         "CloudGuardWlpAgent",
			Dependencies: acctest.DependencyGraph["wlpAgent"],
			F:            sweepCloudGuardWlpAgentResource,
		})
	}
}

func sweepCloudGuardWlpAgentResource(compartment string) error {
	cloudGuardClient := acctest.GetTestClients(&schema.ResourceData{}).CloudGuardClient()
	wlpAgentIds, err := getCloudGuardWlpAgentIds(compartment)
	if err != nil {
		return err
	}
	for _, wlpAgentId := range wlpAgentIds {
		if ok := acctest.SweeperDefaultResourceId[wlpAgentId]; !ok {
			deleteWlpAgentRequest := oci_cloud_guard.DeleteWlpAgentRequest{}

			deleteWlpAgentRequest.WlpAgentId = &wlpAgentId

			deleteWlpAgentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_guard")
			_, error := cloudGuardClient.DeleteWlpAgent(context.Background(), deleteWlpAgentRequest)
			if error != nil {
				fmt.Printf("Error deleting WlpAgent %s %s, It is possible that the resource is already deleted. Please verify manually \n", wlpAgentId, error)
				continue
			}
		}
	}
	return nil
}

func getCloudGuardWlpAgentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "WlpAgentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	cloudGuardClient := acctest.GetTestClients(&schema.ResourceData{}).CloudGuardClient()

	listWlpAgentsRequest := oci_cloud_guard.ListWlpAgentsRequest{}
	listWlpAgentsRequest.CompartmentId = &compartmentId
	listWlpAgentsResponse, err := cloudGuardClient.ListWlpAgents(context.Background(), listWlpAgentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting WlpAgent list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, wlpAgent := range listWlpAgentsResponse.Items {
		id := *wlpAgent.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "WlpAgentId", id)
	}
	return resourceIds, nil
}
