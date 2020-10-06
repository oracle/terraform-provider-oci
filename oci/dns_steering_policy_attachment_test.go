// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/v26/common"
	oci_dns "github.com/oracle/oci-go-sdk/v26/dns"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	SteeringPolicyAttachmentRequiredOnlyResource = SteeringPolicyAttachmentResourceDependencies +
		generateResourceFromRepresentationMap("oci_dns_steering_policy_attachment", "test_steering_policy_attachment", Required, Create, steeringPolicyAttachmentRepresentation)

	SteeringPolicyAttachmentResourceConfig = SteeringPolicyAttachmentResourceDependencies +
		generateResourceFromRepresentationMap("oci_dns_steering_policy_attachment", "test_steering_policy_attachment", Optional, Update, steeringPolicyAttachmentRepresentation)

	steeringPolicyAttachmentSingularDataSourceRepresentation = map[string]interface{}{
		"steering_policy_attachment_id": Representation{repType: Required, create: `${oci_dns_steering_policy_attachment.test_steering_policy_attachment.id}`},
	}

	steeringPolicyAttachmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                        Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":                          Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"domain":                                Representation{repType: Optional, create: `${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test`},
		"id":                                    Representation{repType: Optional, create: `${oci_dns_steering_policy_attachment.test_steering_policy_attachment.id}`},
		"state":                                 Representation{repType: Optional, create: `ACTIVE`},
		"steering_policy_id":                    Representation{repType: Optional, create: `${oci_dns_steering_policy.test_steering_policy.id}`},
		"time_created_greater_than_or_equal_to": Representation{repType: Optional, create: `2018-01-01T00:00:00.000Z`},
		"time_created_less_than":                Representation{repType: Optional, create: `2038-01-01T00:00:00.000Z`},
		"zone_id":                               Representation{repType: Optional, create: `${oci_dns_zone.test_zone.id}`},
		"filter":                                RepresentationGroup{Required, steeringPolicyAttachmentDataSourceFilterRepresentation}}

	// Used to test `domain_contains` query parameter; which cannot be simulataneously used with `domain` query param
	steeringPolicyAttachmentDataSourceRepresentationWithDomainContains = map[string]interface{}{
		"compartment_id":                        Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":                          Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"domain_contains":                       Representation{repType: Optional, create: `${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test`},
		"id":                                    Representation{repType: Optional, create: `${oci_dns_steering_policy_attachment.test_steering_policy_attachment.id}`},
		"state":                                 Representation{repType: Optional, create: `ACTIVE`},
		"steering_policy_id":                    Representation{repType: Optional, create: `${oci_dns_steering_policy.test_steering_policy.id}`},
		"time_created_greater_than_or_equal_to": Representation{repType: Optional, create: `2018-01-01T00:00:00.000Z`},
		"time_created_less_than":                Representation{repType: Optional, create: `2038-01-01T00:00:00.000Z`},
		"zone_id":                               Representation{repType: Optional, create: `${oci_dns_zone.test_zone.id}`},
		"filter":                                RepresentationGroup{Required, steeringPolicyAttachmentDataSourceFilterRepresentation}}

	steeringPolicyAttachmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_dns_steering_policy_attachment.test_steering_policy_attachment.id}`}},
	}

	steeringPolicyAttachmentRepresentation = map[string]interface{}{
		"domain_name":        Representation{repType: Required, create: `${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test`},
		"steering_policy_id": Representation{repType: Required, create: `${oci_dns_steering_policy.test_steering_policy.id}`},
		"zone_id":            Representation{repType: Required, create: `${oci_dns_zone.test_zone.id}`},
		"display_name":       Representation{repType: Optional, create: `displayName`, update: `displayName2`},
	}

	SteeringPolicyAttachmentResourceDependencies = RecordResourceDependencies +
		generateResourceFromRepresentationMap("oci_dns_steering_policy", "test_steering_policy", Required, Create, steeringPolicyRepresentation)
)

func TestDnsSteeringPolicyAttachmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDnsSteeringPolicyAttachmentResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_dns_steering_policy_attachment.test_steering_policy_attachment"
	datasourceName := "data.oci_dns_steering_policy_attachments.test_steering_policy_attachments"
	singularDatasourceName := "data.oci_dns_steering_policy_attachment.test_steering_policy_attachment"

	_, tokenFn := tokenizeWithHttpReplay("dns_steering")
	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDnsSteeringPolicyAttachmentDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: tokenFn(config+compartmentIdVariableStr+SteeringPolicyAttachmentResourceDependencies+
					generateResourceFromRepresentationMap("oci_dns_steering_policy_attachment", "test_steering_policy_attachment", Required, Create, steeringPolicyAttachmentRepresentation), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestMatchResourceAttr(resourceName, "domain_name", regexp.MustCompile("\\.oci-record-test")),
					resource.TestCheckResourceAttrSet(resourceName, "steering_policy_id"),
					resource.TestCheckResourceAttrSet(resourceName, "zone_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: tokenFn(config+compartmentIdVariableStr+SteeringPolicyAttachmentResourceDependencies, nil),
			},
			// verify create with optionals
			{
				Config: tokenFn(config+compartmentIdVariableStr+SteeringPolicyAttachmentResourceDependencies+
					generateResourceFromRepresentationMap("oci_dns_steering_policy_attachment", "test_steering_policy_attachment", Optional, Create, steeringPolicyAttachmentRepresentation), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestMatchResourceAttr(resourceName, "domain_name", regexp.MustCompile("\\.oci-record-test")),
					resource.TestCheckResourceAttrSet(resourceName, "steering_policy_id"),
					resource.TestCheckResourceAttrSet(resourceName, "zone_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
					generateResourceFromRepresentationMap("oci_dns_steering_policy_attachment", "test_steering_policy_attachment", Optional, Update, steeringPolicyAttachmentRepresentation), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestMatchResourceAttr(resourceName, "domain_name", regexp.MustCompile("\\.oci-record-test")),
					resource.TestCheckResourceAttrSet(resourceName, "steering_policy_id"),
					resource.TestCheckResourceAttrSet(resourceName, "zone_id"),

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
				Config: tokenFn(config+
					generateDataSourceFromRepresentationMap("oci_dns_steering_policy_attachments", "test_steering_policy_attachments", Optional, Update, steeringPolicyAttachmentDataSourceRepresentation)+
					compartmentIdVariableStr+SteeringPolicyAttachmentResourceDependencies+
					generateResourceFromRepresentationMap("oci_dns_steering_policy_attachment", "test_steering_policy_attachment", Optional, Update, steeringPolicyAttachmentRepresentation), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_dns_steering_policy_attachments", "test_steering_policy_attachments", Optional, Update, steeringPolicyAttachmentDataSourceRepresentationWithDomainContains)+
					compartmentIdVariableStr+SteeringPolicyAttachmentResourceDependencies+
					generateResourceFromRepresentationMap("oci_dns_steering_policy_attachment", "test_steering_policy_attachment", Optional, Update, steeringPolicyAttachmentRepresentation), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_dns_steering_policy_attachment", "test_steering_policy_attachment", Required, Create, steeringPolicyAttachmentSingularDataSourceRepresentation)+
					compartmentIdVariableStr+SteeringPolicyAttachmentResourceConfig, nil),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateResourceFromRepresentationMap("oci_dns_steering_policy_attachment", "test_steering_policy_attachment", Optional, Update,
						getUpdatedRepresentationCopy("domain_name", Representation{repType: Required, create: `${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.OCI-record-test`}, steeringPolicyAttachmentRepresentation)), nil),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestMatchResourceAttr(resourceName, "domain_name", regexp.MustCompile("\\.oci-record-test")),
					resource.TestCheckResourceAttrSet(resourceName, "steering_policy_id"),
					resource.TestCheckResourceAttrSet(resourceName, "zone_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
		},
	})
}

func testAccCheckDnsSteeringPolicyAttachmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).dnsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dns_steering_policy_attachment" {
			noResourceFound = false
			request := oci_dns.GetSteeringPolicyAttachmentRequest{}

			tmp := rs.Primary.ID
			request.SteeringPolicyAttachmentId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "dns")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DnsSteeringPolicyAttachment") {
		resource.AddTestSweepers("DnsSteeringPolicyAttachment", &resource.Sweeper{
			Name:         "DnsSteeringPolicyAttachment",
			Dependencies: DependencyGraph["steeringPolicyAttachment"],
			F:            sweepDnsSteeringPolicyAttachmentResource,
		})
	}
}

func sweepDnsSteeringPolicyAttachmentResource(compartment string) error {
	dnsClient := GetTestClients(&schema.ResourceData{}).dnsClient()
	steeringPolicyAttachmentIds, err := getSteeringPolicyAttachmentIds(compartment)
	if err != nil {
		return err
	}
	for _, steeringPolicyAttachmentId := range steeringPolicyAttachmentIds {
		if ok := SweeperDefaultResourceId[steeringPolicyAttachmentId]; !ok {
			deleteSteeringPolicyAttachmentRequest := oci_dns.DeleteSteeringPolicyAttachmentRequest{}

			deleteSteeringPolicyAttachmentRequest.SteeringPolicyAttachmentId = &steeringPolicyAttachmentId

			deleteSteeringPolicyAttachmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "dns")
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
	ids := getResourceIdsToSweep(compartment, "SteeringPolicyAttachmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dnsClient := GetTestClients(&schema.ResourceData{}).dnsClient()

	listSteeringPolicyAttachmentsRequest := oci_dns.ListSteeringPolicyAttachmentsRequest{}
	listSteeringPolicyAttachmentsRequest.CompartmentId = &compartmentId
	listSteeringPolicyAttachmentsResponse, err := dnsClient.ListSteeringPolicyAttachments(context.Background(), listSteeringPolicyAttachmentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SteeringPolicyAttachment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, steeringPolicyAttachment := range listSteeringPolicyAttachmentsResponse.Items {
		id := *steeringPolicyAttachment.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "SteeringPolicyAttachmentId", id)
	}
	return resourceIds, nil
}
