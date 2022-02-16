// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_dns "github.com/oracle/oci-go-sdk/v58/dns"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ViewRequiredOnlyResource = ViewResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dns_view", "test_view", acctest.Required, acctest.Create, viewRepresentation)

	ViewResourceConfig = ViewResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dns_view", "test_view", acctest.Optional, acctest.Update, viewRepresentation)

	viewSingularDataSourceRepresentation = map[string]interface{}{
		"view_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_view.test_view.id}`},
		"scope":   acctest.Representation{RepType: acctest.Required, Create: `PRIVATE`},
	}

	viewDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_dns_view.test_view.id}`},
		"scope":          acctest.Representation{RepType: acctest.Required, Create: `PRIVATE`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: viewDataSourceFilterRepresentation}}
	viewDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dns_view.test_view.id}`}},
	}

	viewRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"scope":          acctest.Representation{RepType: acctest.Required, Create: `PRIVATE`},
	}
	viewRepresentationDefault = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
	}

	ViewResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: dns/default
func TestDnsViewResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDnsViewResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_dns_view.test_view"
	datasourceName := "data.oci_dns_views.test_views"
	singularDatasourceName := "data.oci_dns_view.test_view"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ViewResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dns_view", "test_view", acctest.Optional, acctest.Create, viewRepresentation), "dns", "view", t)

	acctest.ResourceTest(t, testAccCheckDnsViewDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ViewResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_view", "test_view", acctest.Required, acctest.Create, viewRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ViewResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ViewResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_view", "test_view", acctest.Optional, acctest.Create, viewRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_protected"),
				resource.TestCheckResourceAttr(resourceName, "scope", "PRIVATE"),
				resource.TestCheckResourceAttrSet(resourceName, "self"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					// Resource discovery is disabled for Views
					//if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					//	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					//		return errExport
					//	}
					//}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ViewResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_view", "test_view", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(viewRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_protected"),
				resource.TestCheckResourceAttr(resourceName, "scope", "PRIVATE"),
				resource.TestCheckResourceAttrSet(resourceName, "self"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + ViewResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_view", "test_view", acctest.Optional, acctest.Update, viewRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_protected"),
				resource.TestCheckResourceAttr(resourceName, "scope", "PRIVATE"),
				resource.TestCheckResourceAttrSet(resourceName, "self"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dns_views", "test_views", acctest.Optional, acctest.Update, viewDataSourceRepresentation) +
				compartmentIdVariableStr + ViewResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dns_view", "test_view", acctest.Optional, acctest.Update, viewRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "scope", "PRIVATE"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "views.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "views.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "views.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "views.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "views.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "views.0.is_protected"),
				resource.TestCheckResourceAttrSet(datasourceName, "views.0.self"),
				resource.TestCheckResourceAttrSet(datasourceName, "views.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "views.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "views.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dns_view", "test_view", acctest.Required, acctest.Create, viewSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ViewResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "scope", "PRIVATE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "view_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_protected"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "self"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + ViewResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getDnsViewImportId(resourceName),
			ImportStateVerifyIgnore: []string{
				"scope",
			},
			ResourceName: resourceName,
		},
	})
}

func getDnsViewImportId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("viewId/" + rs.Primary.Attributes["id"] + "/scope/" + rs.Primary.Attributes["scope"]), nil
	}
}

func testAccCheckDnsViewDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DnsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dns_view" {
			noResourceFound = false
			request := oci_dns.GetViewRequest{}

			if value, ok := rs.Primary.Attributes["scope"]; ok {
				request.Scope = oci_dns.GetViewScopeEnum(value)
			}

			tmp := rs.Primary.ID
			request.ViewId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dns")

			_, err := client.GetView(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
			}
			//Verify that exception is for 404.
			// after destruction
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
	if !acctest.InSweeperExcludeList("DnsView") {
		resource.AddTestSweepers("DnsView", &resource.Sweeper{
			Name:         "DnsView",
			Dependencies: acctest.DependencyGraph["view"],
			F:            sweepDnsViewResource,
		})
	}
}

func sweepDnsViewResource(compartment string) error {
	dnsClient := acctest.GetTestClients(&schema.ResourceData{}).DnsClient()
	viewIds, err := getViewIds(compartment)
	if err != nil {
		return err
	}
	for _, viewId := range viewIds {
		if ok := acctest.SweeperDefaultResourceId[viewId]; !ok {
			deleteViewRequest := oci_dns.DeleteViewRequest{}

			deleteViewRequest.ViewId = &viewId

			deleteViewRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dns")
			_, error := dnsClient.DeleteView(context.Background(), deleteViewRequest)
			if error != nil {
				fmt.Printf("Error deleting View %s %s, It is possible that the resource is already deleted. Please verify manually \n", viewId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &viewId, viewSweepWaitCondition, time.Duration(3*time.Minute),
				viewSweepResponseFetchOperation, "dns", true)
		}
	}
	return nil
}

func getViewIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ViewId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dnsClient := acctest.GetTestClients(&schema.ResourceData{}).DnsClient()

	listViewsRequest := oci_dns.ListViewsRequest{}
	listViewsRequest.CompartmentId = &compartmentId
	listViewsRequest.LifecycleState = oci_dns.ViewSummaryLifecycleStateActive
	listViewsResponse, err := dnsClient.ListViews(context.Background(), listViewsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting View list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, view := range listViewsResponse.Items {
		id := *view.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ViewId", id)
	}
	return resourceIds, nil
}

func viewSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if viewResponse, ok := response.Response.(oci_dns.GetViewResponse); ok {
		return viewResponse.LifecycleState != oci_dns.ViewLifecycleStateDeleted
	}
	return false
}

func viewSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DnsClient().GetView(context.Background(), oci_dns.GetViewRequest{
		ViewId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
