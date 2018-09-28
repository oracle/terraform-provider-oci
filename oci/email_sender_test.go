// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

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
	SenderResourceConfig = SenderResourceDependencies +
		generateResourceFromRepresentationMap("oci_email_sender", "test_sender", Required, Create, senderRepresentation)

	senderSingularDataSourceRepresentation = map[string]interface{}{
		"sender_id": Representation{repType: Required, create: `${oci_email_sender.test_sender.id}`},
	}

	senderDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"email_address":  Representation{repType: Optional, create: `JohnSmith@example.com`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, senderDataSourceFilterRepresentation}}
	senderDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_email_sender.test_sender.id}`}},
	}

	senderRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"email_address":  Representation{repType: Required, create: `JohnSmith@example.com`},
	}

	SenderResourceDependencies = ""
)

func TestEmailSenderResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_email_sender.test_sender"
	datasourceName := "data.oci_email_senders.test_senders"
	singularDatasourceName := "data.oci_email_sender.test_sender"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckEmailSenderDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + SenderResourceDependencies +
					generateResourceFromRepresentationMap("oci_email_sender", "test_sender", Required, Create, senderRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "email_address", "JohnSmith@example.com"),
				),
			},

			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_email_senders", "test_senders", Optional, Update, senderDataSourceRepresentation) +
					compartmentIdVariableStr + SenderResourceDependencies +
					generateResourceFromRepresentationMap("oci_email_sender", "test_sender", Optional, Update, senderRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "email_address", "JohnSmith@example.com"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "senders.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "senders.0.email_address", "JohnSmith@example.com"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_email_sender", "test_sender", Required, Create, senderSingularDataSourceRepresentation) +
					compartmentIdVariableStr + SenderResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "sender_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "email_address", "JohnSmith@example.com"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "is_spf", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + SenderResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"compartment_id",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckEmailSenderDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).emailClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_email_sender" {
			noResourceFound = false
			request := oci_email.GetSenderRequest{}

			tmp := rs.Primary.ID
			request.SenderId = &tmp

			response, err := client.GetSender(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_email.SenderLifecycleStateDeleted): true,
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
