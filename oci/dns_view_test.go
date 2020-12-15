// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v31/common"
	oci_dns "github.com/oracle/oci-go-sdk/v31/dns"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ViewRequiredOnlyResource = ViewResourceDependencies +
		generateResourceFromRepresentationMap("oci_dns_view", "test_view", Required, Create, viewRepresentation)

	ViewResourceConfig = ViewResourceDependencies +
		generateResourceFromRepresentationMap("oci_dns_view", "test_view", Optional, Update, viewRepresentation)

	viewSingularDataSourceRepresentation = map[string]interface{}{
		"view_id": Representation{repType: Required, create: `${oci_dns_view.test_view.id}`},
		"scope":   Representation{repType: Required, create: `PRIVATE`},
	}

	viewDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"id":             Representation{repType: Optional, create: `${oci_dns_view.test_view.id}`},
		"scope":          Representation{repType: Required, create: `PRIVATE`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, viewDataSourceFilterRepresentation}}
	viewDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_dns_view.test_view.id}`}},
	}

	viewRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"freeformTags": "freeformTags"}, update: map[string]string{"freeformTags2": "freeformTags2"}},
		"scope":          Representation{repType: Required, create: `PRIVATE`},
	}
	viewRepresentationDefault = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"freeformTags": "freeformTags"}, update: map[string]string{"freeformTags2": "freeformTags2"}},
	}

	ViewResourceDependencies = DefinedTagsDependencies
)

func TestDnsViewResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDnsViewResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_dns_view.test_view"
	datasourceName := "data.oci_dns_views.test_views"
	singularDatasourceName := "data.oci_dns_view.test_view"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDnsViewDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ViewResourceDependencies +
					generateResourceFromRepresentationMap("oci_dns_view", "test_view", Required, Create, viewRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ViewResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ViewResourceDependencies +
					generateResourceFromRepresentationMap("oci_dns_view", "test_view", Optional, Create, viewRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId, err = fromInstanceState(s, resourceName, "id")
						// Resource discovery is disabled for Views
						//if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						//	if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
						//		return errExport
						//	}
						//}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ViewResourceDependencies +
					generateResourceFromRepresentationMap("oci_dns_view", "test_view", Optional, Create,
						representationCopyWithNewProperties(viewRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_dns_view", "test_view", Optional, Update, viewRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_dns_views", "test_views", Optional, Update, viewDataSourceRepresentation) +
					compartmentIdVariableStr + ViewResourceDependencies +
					generateResourceFromRepresentationMap("oci_dns_view", "test_view", Optional, Update, viewRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "scope", "PRIVATE"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttr(datasourceName, "views.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "views.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "views.0.defined_tags.%", "1"),
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
					generateDataSourceFromRepresentationMap("oci_dns_view", "test_view", Required, Create, viewSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ViewResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(singularDatasourceName, "scope", "PRIVATE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "view_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
	client := testAccProvider.Meta().(*OracleClients).dnsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dns_view" {
			noResourceFound = false
			request := oci_dns.GetViewRequest{}

			if value, ok := rs.Primary.Attributes["scope"]; ok {
				request.Scope = oci_dns.GetViewScopeEnum(value)
			}

			tmp := rs.Primary.ID
			request.ViewId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "dns")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DnsView") {
		resource.AddTestSweepers("DnsView", &resource.Sweeper{
			Name:         "DnsView",
			Dependencies: DependencyGraph["view"],
			F:            sweepDnsViewResource,
		})
	}
}

func sweepDnsViewResource(compartment string) error {
	dnsClient := GetTestClients(&schema.ResourceData{}).dnsClient()
	viewIds, err := getViewIds(compartment)
	if err != nil {
		return err
	}
	for _, viewId := range viewIds {
		if ok := SweeperDefaultResourceId[viewId]; !ok {
			deleteViewRequest := oci_dns.DeleteViewRequest{}

			deleteViewRequest.ViewId = &viewId

			deleteViewRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "dns")
			_, error := dnsClient.DeleteView(context.Background(), deleteViewRequest)
			if error != nil {
				fmt.Printf("Error deleting View %s %s, It is possible that the resource is already deleted. Please verify manually \n", viewId, error)
				continue
			}
			waitTillCondition(testAccProvider, &viewId, viewSweepWaitCondition, time.Duration(3*time.Minute),
				viewSweepResponseFetchOperation, "dns", true)
		}
	}
	return nil
}

func getViewIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ViewId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dnsClient := GetTestClients(&schema.ResourceData{}).dnsClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "ViewId", id)
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

func viewSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.dnsClient().GetView(context.Background(), oci_dns.GetViewRequest{
		ViewId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
