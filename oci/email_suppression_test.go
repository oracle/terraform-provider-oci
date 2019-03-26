// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_email "github.com/oracle/oci-go-sdk/email"
)

var (
	SuppressionResourceConfig = SuppressionResourceDependencies +
		generateResourceFromRepresentationMap("oci_email_suppression", "test_suppression", Optional, Update, suppressionRepresentation)

	suppressionSingularDataSourceRepresentation = map[string]interface{}{
		"suppression_id": Representation{repType: Required, create: `${oci_email_suppression.test_suppression.id}`},
	}

	suppressionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                        Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"email_address":                         Representation{repType: Optional, create: `JohnSmith@example.com`},
		"time_created_greater_than_or_equal_to": Representation{repType: Optional, create: `2018-01-01T00:00:00.000Z`},
		"time_created_less_than":                Representation{repType: Optional, create: `2038-01-01T00:00:00.000Z`},
		"filter":                                RepresentationGroup{Required, suppressionDataSourceFilterRepresentation}}
	suppressionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_email_suppression.test_suppression.id}`}},
	}

	suppressionRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"email_address":  Representation{repType: Required, create: `JohnSmith@example.com`},
	}

	SuppressionResourceDependencies = ""
)

func TestEmailSuppressionResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_email_suppression.test_suppression"
	datasourceName := "data.oci_email_suppressions.test_suppressions"
	singularDatasourceName := "data.oci_email_suppression.test_suppression"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckEmailSuppressionDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + SuppressionResourceDependencies +
					generateResourceFromRepresentationMap("oci_email_suppression", "test_suppression", Required, Create, suppressionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					// email address is converted to lower case by the service
					resource.TestCheckResourceAttr(resourceName, "email_address", "johnsmith@example.com"),
				),
			},

			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_email_suppressions", "test_suppressions", Optional, Update, suppressionDataSourceRepresentation) +
					compartmentIdVariableStr + SuppressionResourceDependencies +
					generateResourceFromRepresentationMap("oci_email_suppression", "test_suppression", Optional, Update, suppressionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(datasourceName, "email_address", "JohnSmith@example.com"),
					resource.TestCheckResourceAttr(datasourceName, "time_created_greater_than_or_equal_to", "2018-01-01T00:00:00.000Z"),
					resource.TestCheckResourceAttr(datasourceName, "time_created_less_than", "2038-01-01T00:00:00.000Z"),

					resource.TestCheckResourceAttr(datasourceName, "suppressions.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "suppressions.0.compartment_id", tenancyId),
					// email address is converted to lower case by the service
					resource.TestCheckResourceAttr(datasourceName, "suppressions.0.email_address", "johnsmith@example.com"),
					resource.TestCheckResourceAttrSet(datasourceName, "suppressions.0.time_created"),
					resource.TestCheckResourceAttr(datasourceName, "suppressions.0.reason", "MANUAL"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_email_suppression", "test_suppression", Required, Create, suppressionSingularDataSourceRepresentation) +
					compartmentIdVariableStr + SuppressionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "suppression_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
					// email address is converted to lower case by the service
					resource.TestCheckResourceAttr(singularDatasourceName, "email_address", "johnsmith@example.com"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "reason", "MANUAL"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + SuppressionResourceConfig,
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

func testAccCheckEmailSuppressionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).emailClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_email_suppression" {
			noResourceFound = false
			request := oci_email.GetSuppressionRequest{}

			tmp := rs.Primary.ID
			request.SuppressionId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "email")

			_, err := client.GetSuppression(context.Background(), request)

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
