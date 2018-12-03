// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	reconfirmationSingularDataSourceRepresentation = map[string]interface{}{
		"subscription_id": Representation{repType: Required, create: `${oci_ons_subscription.test_subscription.id}`},
	}
	ReconfirmationResourceConfig = SubscriptionRequiredOnlyResource
)

func TestOnsReconfirmationResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_ons_reconfirmation.test_reconfirmation"
	resourceName := "oci_ons_subscription.test_subscription"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_ons_reconfirmation", "test_reconfirmation", Required, Create, reconfirmationSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ReconfirmationResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "subscription_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "url"),
					resource.TestCheckResourceAttr(resourceName, "state", "PENDING"),
					func(s *terraform.State) (err error) {
						url, err := fromInstanceState(s, singularDatasourceName, "url")
						_, err = http.Get(url)
						return err
					},
				),
			},
			// verify subscription state
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_ons_reconfirmation", "test_reconfirmation", Required, Create, reconfirmationSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ReconfirmationResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				),
			},
		},
	})
}
