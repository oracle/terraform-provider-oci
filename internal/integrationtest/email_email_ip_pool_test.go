// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_email "github.com/oracle/oci-go-sdk/v65/email"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	emailpkg "github.com/oracle/terraform-provider-oci/internal/service/email"
)

var (
	EmailEmailIpPoolResourceDependencies = ""

	EmailEmailIpPoolRequiredOnlyResource = EmailEmailIpPoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_email_email_ip_pool", "test_email_ip_pool", acctest.Required, acctest.Create, EmailEmailIpPoolRepresentation)

	EmailEmailIpPoolResourceConfig = EmailEmailIpPoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_email_email_ip_pool", "test_email_ip_pool", acctest.Optional, acctest.Update, EmailEmailIpPoolRepresentation)

	EmailEmailIpPoolSingularDataSourceRepresentation = map[string]interface{}{
		"email_ip_pool_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_email_email_ip_pool.test_email_ip_pool.id}`},
	}

	EmailEmailIpPoolDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_email_email_ip_pool.test_email_ip_pool.id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: EmailEmailIpPoolDataSourceFilterRepresentation}}
	EmailEmailIpPoolDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_email_email_ip_pool.test_email_ip_pool.id}`}},
	}

	EmailEmailIpPoolRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `name`},
		"outbound_ips": acctest.Representation{
			RepType: acctest.Required,
			Create:  []string{`${var.ip1}`, `${var.ip2}`},
			Update:  []string{`${var.ip1}`, `${var.ip2}`},
		},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	EmailEmailIpPoolRepresentation_SingleIP = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `name`},
		"outbound_ips": acctest.Representation{
			RepType: acctest.Required,
			Create:  []string{`${var.ip1}`},
			Update:  []string{`${var.ip1}`},
		},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	EmailEmailIpPoolRepresentation_Empty = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `name`},
		"outbound_ips": acctest.Representation{
			RepType: acctest.Required,
			Create:  []string{},
			Update:  []string{},
		},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
	}
)

func TestEmailEmailIpPoolResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestEmailEmailIpPoolResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	ip1 := utils.GetEnvSettingWithBlankDefault("OCI_EMAIL_IP_1")
	ip2 := utils.GetEnvSettingWithBlankDefault("OCI_EMAIL_IP_2")

	if ip1 == "" || ip2 == "" {
		t.Skip("Skipping test: OCI_EMAIL_IP_1 and OCI_EMAIL_IP_2 env vars must be set with valid Dedicated IPs")
	}

	outboundIpVariableStr := fmt.Sprintf(`
	variable "ip1" { default = "%s" }
	variable "ip2" { default = "%s" }
	`, ip1, ip2)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_email_email_ip_pool.test_email_ip_pool"

	baseConfig := config + compartmentIdVariableStr + outboundIpVariableStr + EmailEmailIpPoolResourceDependencies

	acctest.SaveConfigContent(config+compartmentIdVariableStr+outboundIpVariableStr+EmailEmailIpPoolResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_email_email_ip_pool", "test_email_ip_pool", acctest.Optional, acctest.Create, EmailEmailIpPoolRepresentation), "email", "emailIpPool", t)

	acctest.ResourceTest(t, testAccCheckEmailEmailIpPoolDestroy, []resource.TestStep{
		{
			Config: baseConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_email_email_ip_pool", "test_email_ip_pool", acctest.Required, acctest.Create, EmailEmailIpPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "ip-pool"),
				resource.TestCheckResourceAttr(resourceName, "outbound_ips.#", "2"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		{
			Config: baseConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_email_email_ip_pool", "test_email_ip_pool", acctest.Optional, acctest.Update, EmailEmailIpPoolRepresentation_SingleIP),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "outbound_ips.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				func(s *terraform.State) error {
					fmt.Println("Waiting 30 seconds for IP Pool state to stabilize after IP removal...")
					time.Sleep(30 * time.Second)
					return nil
				},
			),
		},

		{
			Config: baseConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_email_email_ip_pool", "test_email_ip_pool", acctest.Optional, acctest.Update, EmailEmailIpPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.Department", "Accounting"),
				resource.TestCheckResourceAttr(resourceName, "outbound_ips.#", "2"), // Verify IP is back
				func(s *terraform.State) error {
					time.Sleep(30 * time.Second)
					return nil
				},
			),
		},

		{
			ResourceName:            resourceName,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"outbound_ips", "freeform_tags", "defined_tags"},
		},

		// Final Step: Clean the Ips from the ip pool, but the Ip stays in draining state hence test fails
		{
			Config: baseConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_email_email_ip_pool", "test_email_ip_pool", acctest.Optional, acctest.Update, EmailEmailIpPoolRepresentation_Empty),
		},
	})
}

func testAccCheckEmailEmailIpPoolDestroy(s *terraform.State) error {
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).EmailClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_email_email_ip_pool" {
			request := oci_email.GetEmailIpPoolRequest{}
			tmp := rs.Primary.ID
			request.EmailIpPoolId = &tmp
			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "email")

			response, err := client.GetEmailIpPool(context.Background(), request)

			if err == nil {
				if response.LifecycleState == oci_email.EmailIpPoolLifecycleStateDeleted {
					return nil
				}
			}

			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	return nil
}

func TestEmailEmailIpPoolsDataSource_withLocks(t *testing.T) {
	// 1. Initialize the Plural Data Source
	// Ensure 'emailpkg' matches your import alias for the provider package
	resource := emailpkg.EmailEmailIpPoolsDataSource()

	// 2. Setup the test data with basic requirements (compartment_id is usually required)
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{
		"compartment_id": "ocid1.compartment.oc1..mocked",
	})

	mockedTime, _ := time.Parse(time.RFC3339Nano, "2023-09-19T17:29:12.144Z")

	// 3. Build one mocked IP Pool object containing the lock
	ipPool := map[string]interface{}{
		"id":             "mocked-ip-pool-id",
		"name":           "mocked-pool-name",
		"compartment_id": "ocid1.compartment.oc1..mocked",
		"state":          "ACTIVE",
		"locks": []interface{}{
			map[string]interface{}{
				"type":         "FULL",
				"time_created": mockedTime.Format(time.RFC3339Nano),
				// "message": "Optional lock message", // Add if your schema supports it
			},
		},
	}

	// 4. Inject into the data source's computed list field.
	// NOTE: Check your schema to confirm if the list is named "email_ip_pool_collection" or "email_ip_pools"

	if err := d.Set("email_ip_pool_collection", []interface{}{
		map[string]interface{}{
			"items": []interface{}{ipPool},
		},
	}); err != nil {
		t.Fatalf("failed to set email_ip_pool_collection: %v", err)
	}

	// 5. Assertions
	collections := d.Get("email_ip_pool_collection").([]interface{})
	if len(collections) != 1 {
		t.Fatalf("expected 1 collection, got %d", len(collections))
	}
	c := collections[0].(map[string]interface{})
	items := c["items"].([]interface{})
	if len(items) != 1 {
		t.Fatalf("expected 1 ip pool, got %d", len(items))
	}
	p := items[0].(map[string]interface{})
	locks := p["locks"].([]interface{})
	if len(locks) != 1 {
		t.Fatalf("expected 1 lock, got %d", len(locks))
	}

	lock := locks[0].(map[string]interface{})
	check := func(key, expected string) {
		if actual, ok := lock[key]; !ok || actual != expected {
			t.Errorf("expected lock[%s] = %q, got %v", key, expected, actual)
		}
	}

	check("type", "FULL")
	check("time_created", "2023-09-19T17:29:12.144Z")
}
func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("EmailEmailIpPool") {
		resource.AddTestSweepers("EmailEmailIpPool", &resource.Sweeper{
			Name:         "EmailEmailIpPool",
			Dependencies: acctest.DependencyGraph["emailIpPool"],
			F:            sweepEmailEmailIpPoolResource,
		})
	}
}

func sweepEmailEmailIpPoolResource(compartment string) error {
	emailClient := acctest.GetTestClients(&schema.ResourceData{}).EmailClient()
	emailIpPoolIds, err := getEmailEmailIpPoolIds(compartment)
	if err != nil {
		return err
	}
	for _, emailIpPoolId := range emailIpPoolIds {
		if ok := acctest.SweeperDefaultResourceId[emailIpPoolId]; !ok {
			pool, err := emailClient.GetEmailIpPool(context.Background(), oci_email.GetEmailIpPoolRequest{EmailIpPoolId: &emailIpPoolId})
			if err == nil {
				var ips []string
				for _, ip := range pool.OutboundIps {
					if ip.OutboundIp != nil {
						ips = append(ips, *ip.OutboundIp)
					}
				}
				if len(ips) > 0 {
					emailClient.RemoveEmailOutboundIp(context.Background(), oci_email.RemoveEmailOutboundIpRequest{
						EmailIpPoolId: &emailIpPoolId,
						RemoveEmailOutboundIpDetails: oci_email.RemoveEmailOutboundIpDetails{
							OutboundIps: ips,
						},
					})
					time.Sleep(10 * time.Second)
				}
			}

			_, err = emailClient.DeleteEmailIpPool(context.Background(), oci_email.DeleteEmailIpPoolRequest{EmailIpPoolId: &emailIpPoolId})
			if err != nil {
				fmt.Printf("Sweeper: Delete failed for %s (likely 24h lock). IPs should be detached. Error: %s\n", emailIpPoolId, err)
			}
		}
	}
	return nil
}

func getEmailEmailIpPoolIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "EmailIpPoolId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	emailClient := acctest.GetTestClients(&schema.ResourceData{}).EmailClient()

	listEmailIpPoolsRequest := oci_email.ListEmailIpPoolsRequest{}
	listEmailIpPoolsRequest.CompartmentId = &compartmentId
	listEmailIpPoolsRequest.LifecycleState = oci_email.EmailIpPoolLifecycleStateActive
	listEmailIpPoolsResponse, err := emailClient.ListEmailIpPools(context.Background(), listEmailIpPoolsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting EmailIpPool list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, emailIpPool := range listEmailIpPoolsResponse.Items {
		id := *emailIpPool.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "EmailIpPoolId", id)
	}
	return resourceIds, nil
}

func EmailEmailIpPoolSweepWaitCondition(response common.OCIOperationResponse) bool {
	if emailIpPoolResponse, ok := response.Response.(oci_email.GetEmailIpPoolResponse); ok {
		return emailIpPoolResponse.LifecycleState != oci_email.EmailIpPoolLifecycleStateDeleted
	}
	return false
}

func EmailEmailIpPoolSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.EmailClient().GetEmailIpPool(context.Background(), oci_email.GetEmailIpPoolRequest{
		EmailIpPoolId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
