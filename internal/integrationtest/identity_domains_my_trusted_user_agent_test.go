// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"

	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsIdentityDomainsMyTrustedUserAgentSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":            acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_trusted_user_agent_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domains_my_trusted_user_agents.test_my_trusted_user_agents.my_trusted_user_agents.0.id}`},
		"attribute_sets":           acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityDomainsMyTrustedUserAgentDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":                acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_trusted_user_agent_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"my_trusted_user_agent_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"attribute_sets":               acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":                  acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsMyTrustedUserAgentResourceConfig = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsMyTrustedUserAgentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsMyTrustedUserAgentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_domains_my_trusted_user_agents.test_my_trusted_user_agents"
	singularDatasourceName := "data.oci_identity_domains_my_trusted_user_agent.test_my_trusted_user_agent"

	acctest.SaveConfigContent("", "", "", t)

	print(acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_trusted_user_agents", "test_my_trusted_user_agents", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsMyTrustedUserAgentDataSourceRepresentation))

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_trusted_user_agents", "test_my_trusted_user_agents", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsMyTrustedUserAgentDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMyTrustedUserAgentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestMatchResourceAttr(datasourceName, "my_trusted_user_agents.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(datasourceName, "my_trusted_user_agents.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_trusted_user_agents", "test_my_trusted_user_agents", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsMyTrustedUserAgentDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_trusted_user_agent", "test_my_trusted_user_agent", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsMyTrustedUserAgentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMyTrustedUserAgentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "my_trusted_user_agent_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_ocid"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "domain_ocid"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "expiry_time"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_created_by.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "idcs_last_modified_by.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "last_used_on"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "location"),
				resource.TestCheckResourceAttr(singularDatasourceName, "meta.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ocid"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "platform"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tenancy_ocid"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "token_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "trust_token"),
				resource.TestCheckResourceAttr(singularDatasourceName, "trusted_factors.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user.#", "1"),
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("IdentityDomainsMyTrustedUserAgent") {
		resource.AddTestSweepers("IdentityDomainsMyTrustedUserAgent", &resource.Sweeper{
			Name:         "IdentityDomainsMyTrustedUserAgent",
			Dependencies: acctest.DependencyGraph["myTrustedUserAgent"],
			F:            sweepIdentityDomainsMyTrustedUserAgentResource,
		})
	}
}

func sweepIdentityDomainsMyTrustedUserAgentResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	myTrustedUserAgentIds, err := getIdentityDomainsMyTrustedUserAgentIds(compartment)
	if err != nil {
		return err
	}
	for _, myTrustedUserAgentId := range myTrustedUserAgentIds {
		if ok := acctest.SweeperDefaultResourceId[myTrustedUserAgentId]; !ok {
			deleteMyTrustedUserAgentRequest := oci_identity_domains.DeleteMyTrustedUserAgentRequest{}

			deleteMyTrustedUserAgentRequest.MyTrustedUserAgentId = &myTrustedUserAgentId

			deleteMyTrustedUserAgentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteMyTrustedUserAgent(context.Background(), deleteMyTrustedUserAgentRequest)
			if error != nil {
				fmt.Printf("Error deleting MyTrustedUserAgent %s %s, It is possible that the resource is already deleted. Please verify manually \n", myTrustedUserAgentId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsMyTrustedUserAgentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MyTrustedUserAgentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listMyTrustedUserAgentsRequest := oci_identity_domains.ListMyTrustedUserAgentsRequest{}
	listMyTrustedUserAgentsResponse, err := identityDomainsClient.ListMyTrustedUserAgents(context.Background(), listMyTrustedUserAgentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MyTrustedUserAgent list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, myTrustedUserAgent := range listMyTrustedUserAgentsResponse.Resources {
		id := *myTrustedUserAgent.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MyTrustedUserAgentId", id)
	}
	return resourceIds, nil
}
