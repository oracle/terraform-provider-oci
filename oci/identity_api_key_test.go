// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_identity "github.com/oracle/oci-go-sdk/identity"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	apiKeyDataSourceRepresentation = map[string]interface{}{
		"user_id": Representation{repType: Required, create: `${oci_identity_user.test_user.id}`},
		"filter":  RepresentationGroup{Required, apiKeyDataSourceFilterRepresentation}}
	apiKeyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_identity_api_key.test_api_key.id}`}},
	}

	apiKeyRepresentation = map[string]interface{}{
		"key_value": Representation{repType: Required, create: `${var.api_key_value}`},
		"user_id":   Representation{repType: Required, create: `${oci_identity_user.test_user.id}`},
	}

	ApiKeyResourceDependencies = generateResourceFromRepresentationMap("oci_identity_user", "test_user", Required, Create, userRepresentation)

	apiKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4fGHcxbEs3VaWoKaGUiP
HGZ5ILiOXCcWN4nOgLr6CSzUjtgjmN3aA6rsT2mYiD+M5EecDbEUMectUhNtLl5L
PABN9kpjuR0zxCJXvYYQiCBtdjb1/YxrZI9T/9Jtd+cTabCahJHR/cR8jFmvO4cK
JCa/0+Y00zvktrqniHIn3edGAKC4Ttlwj/1NqT0ZVePMXg3rWHPsIW6ONfdn6FNf
Met8Qa8K3C9xVvzImlYx8PQBy/44Ilu5T3A+puwb2QMeZnQZGDALOY4MvrBTTA1T
djFpg1N/Chj2rGYzreysqlnKFu+1qg64wel39kHkppz4Fv2vaLXF9qIeDjeo3G4s
HQIDAQAB
-----END PUBLIC KEY-----`

	apiKey2 = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvLA8ZvgZBJy1nNvFAc7V
qocUbYTg3skMJqEn6N9iH9le7Isvgc/owePuH4eP6AOIvKZA4g9TdxJoJIuh06J1
KpMmRbvA8556zIUjaGwF7dL0qfp2Llv3KEAcWfmWQxtfy/IBh9FgA+xHl6QXDp+O
nsRc4FBQSw9Ldp36h9JLQrXo9PcGkD8IGmsJ/7gvdh/tvccSYhJ1vYYLtq5WZnn6
Di9EjV2cP2F43YE1wlrRjzliZOB8M2neUjF7IG3Rszd6Ij3jYL1W1N5GZj+E+Yiu
27Z+8kUy/d4s9TVKr6BWaH2xL/YirrE2ARM57WBOXciqaE9PUGs8bdKjRzImfp/4
pQIDAQAB
-----END PUBLIC KEY-----`
)

func TestIdentityApiKeyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityApiKeyResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	apiKeyVarStr := fmt.Sprintf("variable \"api_key_value\" { \n\tdefault = <<EOF\n%s\nEOF\n}\n", apiKey)

	resourceName := "oci_identity_api_key.test_api_key"
	datasourceName := "data.oci_identity_api_keys.test_api_keys"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckIdentityApiKeyDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + apiKeyVarStr + compartmentIdVariableStr + ApiKeyResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_api_key", "test_api_key", Required, Create, apiKeyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "key_value", apiKey),
					resource.TestCheckResourceAttrSet(resourceName, "user_id"),
				),
			},

			// verify datasource
			{
				Config: config + apiKeyVarStr +
					generateDataSourceFromRepresentationMap("oci_identity_api_keys", "test_api_keys", Optional, Update, apiKeyDataSourceRepresentation) +
					compartmentIdVariableStr + ApiKeyResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_api_key", "test_api_key", Optional, Update, apiKeyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "user_id"),

					resource.TestCheckResourceAttr(datasourceName, "api_keys.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "api_keys.0.fingerprint"),
					resource.TestCheckResourceAttrSet(datasourceName, "api_keys.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "api_keys.0.key_value", apiKey),
					resource.TestCheckResourceAttrSet(datasourceName, "api_keys.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "api_keys.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "api_keys.0.user_id"),
				),
			},
		},
	})
}

func testAccCheckIdentityApiKeyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).identityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_api_key" {
			noResourceFound = false
			request := oci_identity.ListApiKeysRequest{}

			if value, ok := rs.Primary.Attributes["user_id"]; ok {
				request.UserId = &value
			}

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "identity")
			response, err := client.ListApiKeys(context.Background(), request)

			if err == nil {
				fingerprint := rs.Primary.Attributes["fingerprint"]
				for _, item := range response.Items {
					if *item.Fingerprint == fingerprint {
						return fmt.Errorf("item still exists")
					}
				}
				// no error and item not found, that means item is deleted. continue checking next one
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
