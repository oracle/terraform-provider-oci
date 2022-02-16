// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
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
	SteeringPolicyAttachmentRequiredOnlyResource = SteeringPolicyAttachmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dns_steering_policy_attachment", "test_steering_policy_attachment", acctest.Required, acctest.Create, steeringPolicyAttachmentRepresentation)

	SteeringPolicyAttachmentResourceConfig = SteeringPolicyAttachmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dns_steering_policy_attachment", "test_steering_policy_attachment", acctest.Optional, acctest.Update, steeringPolicyAttachmentRepresentation)

	steeringPolicyAttachmentSingularDataSourceRepresentation = map[string]interface{}{
		"steering_policy_attachment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_steering_policy_attachment.test_steering_policy_attachment.id}`},
	}

	steeringPolicyAttachmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"domain":                                acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test`},
		"id":                                    acctest.Representation{RepType: acctest.Optional, Create: `${oci_dns_steering_policy_attachment.test_steering_policy_attachment.id}`},
		"state":                                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"steering_policy_id":                    acctest.Representation{RepType: acctest.Optional, Create: `${oci_dns_steering_policy.test_steering_policy.id}`},
		"time_created_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2018-01-01T00:00:00.000Z`},
		"time_created_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `2038-01-01T00:00:00.000Z`},
		"zone_id":                               acctest.Representation{RepType: acctest.Optional, Create: `${oci_dns_zone.test_global_zone.id}`},
		"filter":                                acctest.RepresentationGroup{RepType: acctest.Required, Group: steeringPolicyAttachmentDataSourceFilterRepresentation}}

	// Used to test `domain_contains` query parameter; which cannot be simulataneously used with `domain` query param
	steeringPolicyAttachmentDataSourceRepresentationWithDomainContains = map[string]interface{}{
		"compartment_id":                        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"domain_contains":                       acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test`},
		"id":                                    acctest.Representation{RepType: acctest.Optional, Create: `${oci_dns_steering_policy_attachment.test_steering_policy_attachment.id}`},
		"state":                                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"steering_policy_id":                    acctest.Representation{RepType: acctest.Optional, Create: `${oci_dns_steering_policy.test_steering_policy.id}`},
		"time_created_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2018-01-01T00:00:00.000Z`},
		"time_created_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `2038-01-01T00:00:00.000Z`},
		"zone_id":                               acctest.Representation{RepType: acctest.Optional, Create: `${oci_dns_zone.test_global_zone.id}`},
		"filter":                                acctest.RepresentationGroup{RepType: acctest.Required, Group: steeringPolicyAttachmentDataSourceFilterRepresentation}}

	steeringPolicyAttachmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dns_steering_policy_attachment.test_steering_policy_attachment.id}`}},
	}

	steeringPolicyAttachmentRepresentation = map[string]interface{}{
		"domain_name":        acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test`},
		"steering_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_steering_policy.test_steering_policy.id}`},
		"zone_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_dns_zone.test_global_zone.id}`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
	}

	SteeringPolicyAttachmentResourceDependencies = RecordResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dns_steering_policy", "test_steering_policy", acctest.Required, acctest.Create, steeringPolicyRepresentation)
)

// issue-routing-tag: dns/default
func TestDnsSteeringPolicyAttachmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDnsSteeringPolicyAttachmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dns_steering_policy_attachment.test_steering_policy_attachment"
	datasourceName := "data.oci_dns_steering_policy_attachments.test_steering_policy_attachments"
	singularDatasourceName := "data.oci_dns_steering_policy_attachment.test_steering_policy_attachment"

	_, tokenFn := acctest.TokenizeWithHttpReplay("dns_steering")
	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(tokenFn(config+compartmentIdVariableStr+SteeringPolicyAttachmentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dns_steering_policy_attachment", "test_steering_policy_attachment", acctest.Optional, acctest.Create, steeringPolicyAttachmentRepresentation), nil), "dns", "steeringPolicyAttachment", t)

	acctest.ResourceTest(t, testAccCheckDnsSteeringPolicyAttachmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: tokenFn(config+compartmentIdVariableStr+SteeringPolicyAttachmentResourceDependencies+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_steering_policy_attachment", "test_steering_policy_attachment", acctest.Required, acctest.Create, steeringPolicyAttachmentRepresentation), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestMatchResourceAttr(resourceName, "domain_name", regexp.MustCompile("\\.oci-record-test")),
				resource.TestCheckResourceAttrSet(resourceName, "steering_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: tokenFn(config+compartmentIdVariableStr+SteeringPolicyAttachmentResourceDependencies, nil),
		},
		// verify Create with optionals
		{
			Config: tokenFn(config+compartmentIdVariableStr+SteeringPolicyAttachmentResourceDependencies+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_steering_policy_attachment", "test_steering_policy_attachment", acctest.Optional, acctest.Create, steeringPolicyAttachmentRepresentation), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestMatchResourceAttr(resourceName, "domain_name", regexp.MustCompile("\\.oci-record-test")),
				resource.TestCheckResourceAttrSet(resourceName, "steering_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_id"),

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
			Config: tokenFn(config+compartmentIdVariableStr+SteeringPolicyAttachmentResourceDependencies+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_steering_policy_attachment", "test_steering_policy_attachment", acctest.Optional, acctest.Update, steeringPolicyAttachmentRepresentation), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestMatchResourceAttr(resourceName, "domain_name", regexp.MustCompile("\\.oci-record-test")),
				resource.TestCheckResourceAttrSet(resourceName, "steering_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_id"),

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
			Config: tokenFn(config+
				acctest.GenerateDataSourceFromRepresentationMap("oci_dns_steering_policy_attachments", "test_steering_policy_attachments", acctest.Optional, acctest.Update, steeringPolicyAttachmentDataSourceRepresentation)+
				compartmentIdVariableStr+SteeringPolicyAttachmentResourceDependencies+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_steering_policy_attachment", "test_steering_policy_attachment", acctest.Optional, acctest.Update, steeringPolicyAttachmentRepresentation), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestMatchResourceAttr(datasourceName, "domain", regexp.MustCompile("\\.oci-record-test")),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "steering_policy_id"),
				resource.TestCheckResourceAttr(datasourceName, "time_created_greater_than_or_equal_to", "2018-01-01T00:00:00.000Z"),
				resource.TestCheckResourceAttr(datasourceName, "time_created_less_than", "2038-01-01T00:00:00.000Z"),
				resource.TestCheckResourceAttrSet(datasourceName, "zone_id"),

				resource.TestCheckResourceAttr(datasourceName, "steering_policy_attachments.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "steering_policy_attachments.0.display_name", "displayName2"),
				resource.TestMatchResourceAttr(datasourceName, "steering_policy_attachments.0.domain_name", regexp.MustCompile("\\.oci-record-test")),
				resource.TestCheckResourceAttrSet(datasourceName, "steering_policy_attachments.0.steering_policy_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "steering_policy_attachments.0.zone_id"),
			),
		},
		// verify datasource with domain_contains query param
		{
			Config: tokenFn(config+
				acctest.GenerateDataSourceFromRepresentationMap("oci_dns_steering_policy_attachments", "test_steering_policy_attachments", acctest.Optional, acctest.Update, steeringPolicyAttachmentDataSourceRepresentationWithDomainContains)+
				compartmentIdVariableStr+SteeringPolicyAttachmentResourceDependencies+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_steering_policy_attachment", "test_steering_policy_attachment", acctest.Optional, acctest.Update, steeringPolicyAttachmentRepresentation), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestMatchResourceAttr(datasourceName, "domain_contains", regexp.MustCompile("\\.oci-record-test")),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "steering_policy_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_greater_than_or_equal_to"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_less_than"),
				resource.TestCheckResourceAttrSet(datasourceName, "zone_id"),

				resource.TestCheckResourceAttr(datasourceName, "steering_policy_attachments.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "steering_policy_attachments.0.compartment_id"),
				resource.TestCheckResourceAttr(datasourceName, "steering_policy_attachments.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "steering_policy_attachments.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "steering_policy_attachments.0.self"),
				resource.TestCheckResourceAttrSet(datasourceName, "steering_policy_attachments.0.state"),
				resource.TestMatchResourceAttr(datasourceName, "steering_policy_attachments.0.domain_name", regexp.MustCompile("\\.oci-record-test")),
				resource.TestCheckResourceAttrSet(datasourceName, "steering_policy_attachments.0.steering_policy_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "steering_policy_attachments.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "steering_policy_attachments.0.zone_id"),
			),
		},
		// verify singular datasource
		{
			Config: tokenFn(config+
				acctest.GenerateDataSourceFromRepresentationMap("oci_dns_steering_policy_attachment", "test_steering_policy_attachment", acctest.Required, acctest.Create, steeringPolicyAttachmentSingularDataSourceRepresentation)+
				compartmentIdVariableStr+SteeringPolicyAttachmentResourceConfig, nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "steering_policy_attachment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "steering_policy_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "zone_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestMatchResourceAttr(singularDatasourceName, "domain_name", regexp.MustCompile("\\.oci-record-test")),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "self"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},

		{
			Config: tokenFn(config+compartmentIdVariableStr+SteeringPolicyAttachmentResourceDependencies+
				acctest.GenerateResourceFromRepresentationMap("oci_dns_steering_policy_attachment", "test_steering_policy_attachment", acctest.Optional, acctest.Update,
					acctest.GetUpdatedRepresentationCopy("domain_name", acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.OCI-record-test`}, steeringPolicyAttachmentRepresentation)), nil),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestMatchResourceAttr(resourceName, "domain_name", regexp.MustCompile("\\.oci-record-test")),
				resource.TestCheckResourceAttrSet(resourceName, "steering_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "zone_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: tokenFn(config+compartmentIdVariableStr+SteeringPolicyAttachmentResourceConfig, nil),
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

func testAccCheckDnsSteeringPolicyAttachmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DnsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dns_steering_policy_attachment" {
			noResourceFound = false
			request := oci_dns.GetSteeringPolicyAttachmentRequest{}

			tmp := rs.Primary.ID
			request.SteeringPolicyAttachmentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dns")

			_, err := client.GetSteeringPolicyAttachment(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("DnsSteeringPolicyAttachment") {
		resource.AddTestSweepers("DnsSteeringPolicyAttachment", &resource.Sweeper{
			Name:         "DnsSteeringPolicyAttachment",
			Dependencies: acctest.DependencyGraph["steeringPolicyAttachment"],
			F:            sweepDnsSteeringPolicyAttachmentResource,
		})
	}
}

func sweepDnsSteeringPolicyAttachmentResource(compartment string) error {
	dnsClient := acctest.GetTestClients(&schema.ResourceData{}).DnsClient()
	steeringPolicyAttachmentIds, err := getSteeringPolicyAttachmentIds(compartment)
	if err != nil {
		return err
	}
	for _, steeringPolicyAttachmentId := range steeringPolicyAttachmentIds {
		if ok := acctest.SweeperDefaultResourceId[steeringPolicyAttachmentId]; !ok {
			deleteSteeringPolicyAttachmentRequest := oci_dns.DeleteSteeringPolicyAttachmentRequest{}

			deleteSteeringPolicyAttachmentRequest.SteeringPolicyAttachmentId = &steeringPolicyAttachmentId

			deleteSteeringPolicyAttachmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dns")
			_, error := dnsClient.DeleteSteeringPolicyAttachment(context.Background(), deleteSteeringPolicyAttachmentRequest)
			if error != nil {
				fmt.Printf("Error deleting SteeringPolicyAttachment %s %s, It is possible that the resource is already deleted. Please verify manually \n", steeringPolicyAttachmentId, error)
				continue
			}
		}
	}
	return nil
}

func getSteeringPolicyAttachmentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SteeringPolicyAttachmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dnsClient := acctest.GetTestClients(&schema.ResourceData{}).DnsClient()

	listSteeringPolicyAttachmentsRequest := oci_dns.ListSteeringPolicyAttachmentsRequest{}
	listSteeringPolicyAttachmentsRequest.CompartmentId = &compartmentId
	listSteeringPolicyAttachmentsResponse, err := dnsClient.ListSteeringPolicyAttachments(context.Background(), listSteeringPolicyAttachmentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SteeringPolicyAttachment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, steeringPolicyAttachment := range listSteeringPolicyAttachmentsResponse.Items {
		id := *steeringPolicyAttachment.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SteeringPolicyAttachmentId", id)
	}
	return resourceIds, nil
}
