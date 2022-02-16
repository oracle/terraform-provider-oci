// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_waas "github.com/oracle/oci-go-sdk/v58/waas"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	HttpRedirectRequiredOnlyResource = HttpRedirectResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_waas_http_redirect", "test_http_redirect", acctest.Required, acctest.Create, httpRedirectRepresentation)

	HttpRedirectResourceConfig = HttpRedirectResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_waas_http_redirect", "test_http_redirect", acctest.Optional, acctest.Update, httpRedirectRepresentation)

	httpRedirectSingularDataSourceRepresentation = map[string]interface{}{
		"http_redirect_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_waas_http_redirect.test_http_redirect.id}`},
	}
	domainName = utils.RandomString(6, utils.CharsetWithoutDigits) + ".com"

	httpRedirectDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_names":                         acctest.Representation{RepType: acctest.Optional, Create: []string{`displayName2`}},
		"ids":                                   acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_waas_http_redirect.test_http_redirect.id}`}},
		"states":                                acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"time_created_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2018-01-01T00:00:00.000Z`},
		"time_created_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `2038-01-01T00:00:00.000Z`},
		"filter":                                acctest.RepresentationGroup{RepType: acctest.Required, Group: httpRedirectDataSourceFilterRepresentation}}
	httpRedirectDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_waas_http_redirect.test_http_redirect.id}`}},
	}

	httpRedirectRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"domain":         acctest.Representation{RepType: acctest.Required, Create: domainName},
		"target":         acctest.RepresentationGroup{RepType: acctest.Required, Group: httpRedirectTargetRepresentation},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"response_code":  acctest.Representation{RepType: acctest.Optional, Create: `301`, Update: `302`},
	}
	httpRedirectTargetRepresentation = map[string]interface{}{
		"host":     acctest.Representation{RepType: acctest.Required, Create: `example1.com`, Update: `example2.com`},
		"path":     acctest.Representation{RepType: acctest.Required, Create: `/test{path}`, Update: `/test2{path}`},
		"protocol": acctest.Representation{RepType: acctest.Required, Create: `HTTP`, Update: `HTTPS`},
		"query":    acctest.Representation{RepType: acctest.Required, Create: ``, Update: `{query}`},
		"port":     acctest.Representation{RepType: acctest.Optional, Create: `8080`, Update: `8082`},
	}

	HttpRedirectResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: waas/default
func TestWaasHttpRedirectResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWaasHttpRedirectResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_waas_http_redirect.test_http_redirect"
	datasourceName := "data.oci_waas_http_redirects.test_http_redirects"
	singularDatasourceName := "data.oci_waas_http_redirect.test_http_redirect"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+HttpRedirectResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_waas_http_redirect", "test_http_redirect", acctest.Optional, acctest.Create, httpRedirectRepresentation), "waas", "httpRedirect", t)

	acctest.ResourceTest(t, testAccCheckWaasHttpRedirectDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + HttpRedirectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waas_http_redirect", "test_http_redirect", acctest.Required, acctest.Create, httpRedirectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "domain", domainName),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.host", "example1.com"),
				resource.TestCheckResourceAttr(resourceName, "target.0.path", "/test{path}"),
				resource.TestCheckResourceAttr(resourceName, "target.0.protocol", "HTTP"),
				resource.TestCheckResourceAttr(resourceName, "target.0.query", ""),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + HttpRedirectResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + HttpRedirectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waas_http_redirect", "test_http_redirect", acctest.Optional, acctest.Create, httpRedirectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "domain", domainName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "response_code", "301"),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.host", "example1.com"),
				resource.TestCheckResourceAttr(resourceName, "target.0.path", "/test{path}"),
				resource.TestCheckResourceAttr(resourceName, "target.0.port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "target.0.protocol", "HTTP"),
				resource.TestCheckResourceAttr(resourceName, "target.0.query", ""),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + HttpRedirectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waas_http_redirect", "test_http_redirect", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(httpRedirectRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "domain", domainName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "response_code", "301"),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.host", "example1.com"),
				resource.TestCheckResourceAttr(resourceName, "target.0.path", "/test{path}"),
				resource.TestCheckResourceAttr(resourceName, "target.0.port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "target.0.protocol", "HTTP"),
				resource.TestCheckResourceAttr(resourceName, "target.0.query", ""),

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
			Config: config + compartmentIdVariableStr + HttpRedirectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waas_http_redirect", "test_http_redirect", acctest.Optional, acctest.Update, httpRedirectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "domain", domainName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "response_code", "302"),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.host", "example2.com"),
				resource.TestCheckResourceAttr(resourceName, "target.0.path", "/test2{path}"),
				resource.TestCheckResourceAttr(resourceName, "target.0.port", "8082"),
				resource.TestCheckResourceAttr(resourceName, "target.0.protocol", "HTTPS"),
				resource.TestCheckResourceAttr(resourceName, "target.0.query", "{query}"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_waas_http_redirects", "test_http_redirects", acctest.Optional, acctest.Update, httpRedirectDataSourceRepresentation) +
				compartmentIdVariableStr + HttpRedirectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waas_http_redirect", "test_http_redirect", acctest.Optional, acctest.Update, httpRedirectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_names.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "ids.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "states.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_greater_than_or_equal_to"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_less_than"),

				resource.TestCheckResourceAttr(datasourceName, "http_redirects.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "http_redirects.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "http_redirects.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "http_redirects.0.domain", domainName),
				resource.TestCheckResourceAttr(datasourceName, "http_redirects.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_redirects.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "http_redirects.0.response_code", "302"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_redirects.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "http_redirects.0.target.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "http_redirects.0.target.0.host", "example2.com"),
				resource.TestCheckResourceAttr(datasourceName, "http_redirects.0.target.0.path", "/test2{path}"),
				resource.TestCheckResourceAttr(datasourceName, "http_redirects.0.target.0.port", "8082"),
				resource.TestCheckResourceAttr(datasourceName, "http_redirects.0.target.0.protocol", "HTTPS"),
				resource.TestCheckResourceAttr(datasourceName, "http_redirects.0.target.0.query", "{query}"),
				resource.TestCheckResourceAttrSet(datasourceName, "http_redirects.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_waas_http_redirect", "test_http_redirect", acctest.Required, acctest.Create, httpRedirectSingularDataSourceRepresentation) +
				compartmentIdVariableStr + HttpRedirectResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "http_redirect_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "domain", domainName),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "response_code", "302"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target.0.host", "example2.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target.0.path", "/test2{path}"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target.0.port", "8082"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target.0.protocol", "HTTPS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target.0.query", "{query}"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + HttpRedirectResourceConfig,
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

func testAccCheckWaasHttpRedirectDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).RedirectClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_waas_http_redirect" {
			noResourceFound = false
			request := oci_waas.GetHttpRedirectRequest{}

			tmp := rs.Primary.ID
			request.HttpRedirectId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "waas")

			response, err := client.GetHttpRedirect(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_waas.LifecycleStatesDeleted): true,
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
	if !acctest.InSweeperExcludeList("WaasHttpRedirect") {
		resource.AddTestSweepers("WaasHttpRedirect", &resource.Sweeper{
			Name:         "WaasHttpRedirect",
			Dependencies: acctest.DependencyGraph["httpRedirect"],
			F:            sweepWaasHttpRedirectResource,
		})
	}
}

func sweepWaasHttpRedirectResource(compartment string) error {
	redirectClient := acctest.GetTestClients(&schema.ResourceData{}).RedirectClient()
	httpRedirectIds, err := getHttpRedirectIds(compartment)
	if err != nil {
		return err
	}
	for _, httpRedirectId := range httpRedirectIds {
		if ok := acctest.SweeperDefaultResourceId[httpRedirectId]; !ok {
			deleteHttpRedirectRequest := oci_waas.DeleteHttpRedirectRequest{}

			deleteHttpRedirectRequest.HttpRedirectId = &httpRedirectId

			deleteHttpRedirectRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "waas")
			_, error := redirectClient.DeleteHttpRedirect(context.Background(), deleteHttpRedirectRequest)
			if error != nil {
				fmt.Printf("Error deleting HttpRedirect %s %s, It is possible that the resource is already deleted. Please verify manually \n", httpRedirectId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &httpRedirectId, httpRedirectSweepWaitCondition, time.Duration(3*time.Minute),
				httpRedirectSweepResponseFetchOperation, "waas", true)
		}
	}
	return nil
}

func getHttpRedirectIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "HttpRedirectId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	redirectClient := acctest.GetTestClients(&schema.ResourceData{}).RedirectClient()

	listHttpRedirectsRequest := oci_waas.ListHttpRedirectsRequest{}
	listHttpRedirectsRequest.CompartmentId = &compartmentId
	listHttpRedirectsResponse, err := redirectClient.ListHttpRedirects(context.Background(), listHttpRedirectsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting HttpRedirect list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, httpRedirect := range listHttpRedirectsResponse.Items {
		id := *httpRedirect.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "HttpRedirectId", id)
	}
	return resourceIds, nil
}

func httpRedirectSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if httpRedirectResponse, ok := response.Response.(oci_waas.GetHttpRedirectResponse); ok {
		return httpRedirectResponse.LifecycleState != oci_waas.LifecycleStatesDeleted
	}
	return false
}

func httpRedirectSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.RedirectClient().GetHttpRedirect(context.Background(), oci_waas.GetHttpRedirectRequest{
		HttpRedirectId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
