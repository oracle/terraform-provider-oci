// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	ApiKeyResourceConfigOnly = `
resource "oci_identity_api_key" "test_api_key" {
	#Required
	user_id = "${oci_identity_user.test_user.id}"
	key_value = "${var.api_key_value}"
}`

	ApiKeyResourceConfig = ApiKeyResourceDependencies + ApiKeyResourceConfigOnly

	ApiKeyResourceDependencies = UserPropertyVariables + UserResourceConfig

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
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	apiKeyVarStr := fmt.Sprintf("variable \"api_key_value\" { \n\tdefault = <<EOF\n%sEOF\n}\n", apiKey)

	resourceName := "oci_identity_api_key.test_api_key"
	datasourceName := "data.oci_identity_api_keys.test_api_keys"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + apiKeyVarStr + compartmentIdVariableStr + ApiKeyResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "key_value", apiKey),
					resource.TestCheckResourceAttrSet(resourceName, "user_id"),
				),
			},

			// verify datasource
			{
				Config: config + apiKeyVarStr + `

data "oci_identity_api_keys" "test_api_keys" {
	#Required
	user_id = "${oci_identity_user.test_user.id}"

    filter {
    	name = "id"
    	values = ["${oci_identity_api_key.test_api_key.id}"]
    }
}
                ` + compartmentIdVariableStr + ApiKeyResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "user_id"),

					resource.TestCheckResourceAttr(datasourceName, "api_keys.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "api_keys.0.user_id"),
				),
			},
		},
	})
}
