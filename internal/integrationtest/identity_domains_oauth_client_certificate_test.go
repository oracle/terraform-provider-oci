// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsOAuthClientCertificateRequiredOnlyResource = IdentityDomainsOAuthClientCertificateResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_oauth_client_certificate", "test_oauth_client_certificate", acctest.Required, acctest.Create, IdentityDomainsOAuthClientCertificateRepresentation)

	IdentityDomainsOAuthClientCertificateResourceConfig = IdentityDomainsOAuthClientCertificateResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_oauth_client_certificate", "test_oauth_client_certificate", acctest.Optional, acctest.Update, IdentityDomainsOAuthClientCertificateRepresentation)

	IdentityDomainsOAuthClientCertificateSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":                acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"o_auth_client_certificate_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domains_oauth_client_certificate.test_oauth_client_certificate.id}`},
	}

	IdentityDomainsOAuthClientCertificateDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":                   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"oauth_client_certificate_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"oauth_client_certificate_filter": acctest.Representation{RepType: acctest.Optional, Create: `certificateAlias eq \"certificateAlias\"`},
		"start_index":                     acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsOAuthClientCertificateRepresentation = map[string]interface{}{
		"certificate_alias":     acctest.Representation{RepType: acctest.Required, Create: `certificateAlias`},
		"idcs_endpoint":         acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"schemas":               acctest.Representation{RepType: acctest.Required, Create: []string{`urn:ietf:params:scim:schemas:oracle:idcs:OAuthClientCertificate`}},
		"x509base64certificate": acctest.Representation{RepType: acctest.Required, Create: `MIIC3DCCAkWgAwIBAgIIaW0BvzFmG4EwDQYJKoZIhvcNAQEFBQAwVzETMBEGCgmSJomT8ixkARkWA2NvbTEWMBQGCgmSJomT8ixkARkWBm9yYWNsZTEVMBMGCgmSJomT8ixkARkWBWNsb3VkMREwDwYDVQQDEwhDbG91ZDlDQTAeFw0xNTA5MDIxNTA3MDVaFw0yNTA4MzAxNTA3MDVaMGMxEzARBgoJkiaJk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/IsZAEZFgVjbG91ZDEdMBsGA1UEAwwUSm9lX1RlbmFudDFfTUNTX21jczEwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCG4Y4dk7jxwHonjDZHlny7C7QoyyaNbKf/CA/FiLi89gO7CwH3Rdlai4BHCeTi3VjYqoIgwzCU0rz7DQkWTxrRLYFAgGrGBZ4E23UJ7IPLxJytlJQ9tqdjGT/qqGG5O9lmH0hS1LzHIsKPPjpO02pDpl9VRK6IpGnwgJ/ngzvkReZB25JcO+80Wkw0L5G/jTDHcBfQ+3b/uk8TP+fZTnHj4uvRF362l7rhKmgEuK/xcDexI0v4Rq8Cdp+nty2atKxab9rDVxMkmrfxnlMEyaVFf2BboPosaJfmnUehjWKzPrjp4It6b4GX2Lyu+vKR1hCfJxM3K3nb4Jdbwq4p6trLAgMBAAGjITAfMB0GA1UdDgQWBBT6cdF6e2hC6jglFglA/YC0RGKBODANBgkqhkiG9w0BAQUFAAOBgQB+zVtnzr+h+1lL1xhyrc3mYsUJGvAH5ZSAqLWFNRbaVR8kUme1FJkT6UaI9PEISDhspSY9fqF2eqYSlqMt5ZGWbDn5ZkQ8B6QB1yrvjUndzgLGy+ksIpwN34GDJyarrupwSrd4h5QMhJI4Fmjf+tDwsTmMESyDm8NJCkWfza2KgQ==`},
		"external_id":           acctest.Representation{RepType: acctest.Optional, Create: `externalId`},
	}

	IdentityDomainsOAuthClientCertificateResourceDependencies = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsOAuthClientCertificateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsOAuthClientCertificateResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domains_oauth_client_certificate.test_oauth_client_certificate"
	datasourceName := "data.oci_identity_domains_oauth_client_certificates.test_oauth_client_certificates"
	singularDatasourceName := "data.oci_identity_domains_oauth_client_certificate.test_oauth_client_certificate"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainsOAuthClientCertificateResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_oauth_client_certificate", "test_oauth_client_certificate", acctest.Optional, acctest.Create, IdentityDomainsOAuthClientCertificateRepresentation), "identitydomains", "oAuthClientCertificate", t)

	acctest.ResourceTest(t, testAccCheckIdentityDomainsOAuthClientCertificateDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsOAuthClientCertificateResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_oauth_client_certificate", "test_oauth_client_certificate", acctest.Required, acctest.Create, IdentityDomainsOAuthClientCertificateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "certificate_alias", "certificateAlias"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "x509base64certificate", "MIIC3DCCAkWgAwIBAgIIaW0BvzFmG4EwDQYJKoZIhvcNAQEFBQAwVzETMBEGCgmSJomT8ixkARkWA2NvbTEWMBQGCgmSJomT8ixkARkWBm9yYWNsZTEVMBMGCgmSJomT8ixkARkWBWNsb3VkMREwDwYDVQQDEwhDbG91ZDlDQTAeFw0xNTA5MDIxNTA3MDVaFw0yNTA4MzAxNTA3MDVaMGMxEzARBgoJkiaJk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/IsZAEZFgVjbG91ZDEdMBsGA1UEAwwUSm9lX1RlbmFudDFfTUNTX21jczEwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCG4Y4dk7jxwHonjDZHlny7C7QoyyaNbKf/CA/FiLi89gO7CwH3Rdlai4BHCeTi3VjYqoIgwzCU0rz7DQkWTxrRLYFAgGrGBZ4E23UJ7IPLxJytlJQ9tqdjGT/qqGG5O9lmH0hS1LzHIsKPPjpO02pDpl9VRK6IpGnwgJ/ngzvkReZB25JcO+80Wkw0L5G/jTDHcBfQ+3b/uk8TP+fZTnHj4uvRF362l7rhKmgEuK/xcDexI0v4Rq8Cdp+nty2atKxab9rDVxMkmrfxnlMEyaVFf2BboPosaJfmnUehjWKzPrjp4It6b4GX2Lyu+vKR1hCfJxM3K3nb4Jdbwq4p6trLAgMBAAGjITAfMB0GA1UdDgQWBBT6cdF6e2hC6jglFglA/YC0RGKBODANBgkqhkiG9w0BAQUFAAOBgQB+zVtnzr+h+1lL1xhyrc3mYsUJGvAH5ZSAqLWFNRbaVR8kUme1FJkT6UaI9PEISDhspSY9fqF2eqYSlqMt5ZGWbDn5ZkQ8B6QB1yrvjUndzgLGy+ksIpwN34GDJyarrupwSrd4h5QMhJI4Fmjf+tDwsTmMESyDm8NJCkWfza2KgQ=="),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsOAuthClientCertificateResourceDependencies,
		},
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsOAuthClientCertificateResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_oauth_client_certificate", "test_oauth_client_certificate", acctest.Optional, acctest.Create, IdentityDomainsOAuthClientCertificateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cert_end_date"),
				resource.TestCheckResourceAttrSet(resourceName, "cert_start_date"),
				resource.TestCheckResourceAttr(resourceName, "certificate_alias", "certificateAlias"),
				resource.TestCheckResourceAttr(resourceName, "external_id", "externalId"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "sha1thumbprint"),
				resource.TestCheckResourceAttrSet(resourceName, "sha256thumbprint"),
				resource.TestCheckResourceAttr(resourceName, "x509base64certificate", "MIIC3DCCAkWgAwIBAgIIaW0BvzFmG4EwDQYJKoZIhvcNAQEFBQAwVzETMBEGCgmSJomT8ixkARkWA2NvbTEWMBQGCgmSJomT8ixkARkWBm9yYWNsZTEVMBMGCgmSJomT8ixkARkWBWNsb3VkMREwDwYDVQQDEwhDbG91ZDlDQTAeFw0xNTA5MDIxNTA3MDVaFw0yNTA4MzAxNTA3MDVaMGMxEzARBgoJkiaJk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/IsZAEZFgVjbG91ZDEdMBsGA1UEAwwUSm9lX1RlbmFudDFfTUNTX21jczEwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCG4Y4dk7jxwHonjDZHlny7C7QoyyaNbKf/CA/FiLi89gO7CwH3Rdlai4BHCeTi3VjYqoIgwzCU0rz7DQkWTxrRLYFAgGrGBZ4E23UJ7IPLxJytlJQ9tqdjGT/qqGG5O9lmH0hS1LzHIsKPPjpO02pDpl9VRK6IpGnwgJ/ngzvkReZB25JcO+80Wkw0L5G/jTDHcBfQ+3b/uk8TP+fZTnHj4uvRF362l7rhKmgEuK/xcDexI0v4Rq8Cdp+nty2atKxab9rDVxMkmrfxnlMEyaVFf2BboPosaJfmnUehjWKzPrjp4It6b4GX2Lyu+vKR1hCfJxM3K3nb4Jdbwq4p6trLAgMBAAGjITAfMB0GA1UdDgQWBBT6cdF6e2hC6jglFglA/YC0RGKBODANBgkqhkiG9w0BAQUFAAOBgQB+zVtnzr+h+1lL1xhyrc3mYsUJGvAH5ZSAqLWFNRbaVR8kUme1FJkT6UaI9PEISDhspSY9fqF2eqYSlqMt5ZGWbDn5ZkQ8B6QB1yrvjUndzgLGy+ksIpwN34GDJyarrupwSrd4h5QMhJI4Fmjf+tDwsTmMESyDm8NJCkWfza2KgQ=="),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					idcsEndpoint, err := acctest.FromInstanceState(s, "data.oci_identity_domain.test_domain", "url")
					if err != nil {
						return err
					}

					compositeId := getIdentityDomainsCompositeId(idcsEndpoint, "oauthClientCertificates", resId)
					log.Printf("[DEBUG] Composite ID to import: %s", compositeId)
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_oauth_client_certificates", "test_oauth_client_certificates", acctest.Optional, acctest.Update, IdentityDomainsOAuthClientCertificateDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsOAuthClientCertificateResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_oauth_client_certificate", "test_oauth_client_certificate", acctest.Optional, acctest.Update, IdentityDomainsOAuthClientCertificateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "oauth_client_certificate_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "oauth_client_certificate_filter", "certificateAlias eq \"certificateAlias\""),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestCheckResourceAttr(datasourceName, "oauth_client_certificates.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "oauth_client_certificates.0.schemas.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "oauth_client_certificates.0.cert_end_date"),
				resource.TestCheckResourceAttrSet(datasourceName, "oauth_client_certificates.0.cert_start_date"),
				resource.TestCheckResourceAttr(datasourceName, "oauth_client_certificates.0.certificate_alias", "certificateAlias"),
				resource.TestCheckResourceAttrSet(datasourceName, "oauth_client_certificates.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "oauth_client_certificates.0.sha1thumbprint"),
				resource.TestCheckResourceAttrSet(datasourceName, "oauth_client_certificates.0.sha256thumbprint"),
				resource.TestCheckResourceAttr(datasourceName, "oauth_client_certificates.0.x509base64certificate", "MIIC3DCCAkWgAwIBAgIIaW0BvzFmG4EwDQYJKoZIhvcNAQEFBQAwVzETMBEGCgmSJomT8ixkARkWA2NvbTEWMBQGCgmSJomT8ixkARkWBm9yYWNsZTEVMBMGCgmSJomT8ixkARkWBWNsb3VkMREwDwYDVQQDEwhDbG91ZDlDQTAeFw0xNTA5MDIxNTA3MDVaFw0yNTA4MzAxNTA3MDVaMGMxEzARBgoJkiaJk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/IsZAEZFgVjbG91ZDEdMBsGA1UEAwwUSm9lX1RlbmFudDFfTUNTX21jczEwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCG4Y4dk7jxwHonjDZHlny7C7QoyyaNbKf/CA/FiLi89gO7CwH3Rdlai4BHCeTi3VjYqoIgwzCU0rz7DQkWTxrRLYFAgGrGBZ4E23UJ7IPLxJytlJQ9tqdjGT/qqGG5O9lmH0hS1LzHIsKPPjpO02pDpl9VRK6IpGnwgJ/ngzvkReZB25JcO+80Wkw0L5G/jTDHcBfQ+3b/uk8TP+fZTnHj4uvRF362l7rhKmgEuK/xcDexI0v4Rq8Cdp+nty2atKxab9rDVxMkmrfxnlMEyaVFf2BboPosaJfmnUehjWKzPrjp4It6b4GX2Lyu+vKR1hCfJxM3K3nb4Jdbwq4p6trLAgMBAAGjITAfMB0GA1UdDgQWBBT6cdF6e2hC6jglFglA/YC0RGKBODANBgkqhkiG9w0BAQUFAAOBgQB+zVtnzr+h+1lL1xhyrc3mYsUJGvAH5ZSAqLWFNRbaVR8kUme1FJkT6UaI9PEISDhspSY9fqF2eqYSlqMt5ZGWbDn5ZkQ8B6QB1yrvjUndzgLGy+ksIpwN34GDJyarrupwSrd4h5QMhJI4Fmjf+tDwsTmMESyDm8NJCkWfza2KgQ=="),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_oauth_client_certificate", "test_oauth_client_certificate", acctest.Required, acctest.Create, IdentityDomainsOAuthClientCertificateSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsOAuthClientCertificateResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "o_auth_client_certificate_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "certificate_alias", "certificateAlias"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "x509base64certificate", "MIIC3DCCAkWgAwIBAgIIaW0BvzFmG4EwDQYJKoZIhvcNAQEFBQAwVzETMBEGCgmSJomT8ixkARkWA2NvbTEWMBQGCgmSJomT8ixkARkWBm9yYWNsZTEVMBMGCgmSJomT8ixkARkWBWNsb3VkMREwDwYDVQQDEwhDbG91ZDlDQTAeFw0xNTA5MDIxNTA3MDVaFw0yNTA4MzAxNTA3MDVaMGMxEzARBgoJkiaJk/IsZAEZFgNjb20xFjAUBgoJkiaJk/IsZAEZFgZvcmFjbGUxFTATBgoJkiaJk/IsZAEZFgVjbG91ZDEdMBsGA1UEAwwUSm9lX1RlbmFudDFfTUNTX21jczEwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCG4Y4dk7jxwHonjDZHlny7C7QoyyaNbKf/CA/FiLi89gO7CwH3Rdlai4BHCeTi3VjYqoIgwzCU0rz7DQkWTxrRLYFAgGrGBZ4E23UJ7IPLxJytlJQ9tqdjGT/qqGG5O9lmH0hS1LzHIsKPPjpO02pDpl9VRK6IpGnwgJ/ngzvkReZB25JcO+80Wkw0L5G/jTDHcBfQ+3b/uk8TP+fZTnHj4uvRF362l7rhKmgEuK/xcDexI0v4Rq8Cdp+nty2atKxab9rDVxMkmrfxnlMEyaVFf2BboPosaJfmnUehjWKzPrjp4It6b4GX2Lyu+vKR1hCfJxM3K3nb4Jdbwq4p6trLAgMBAAGjITAfMB0GA1UdDgQWBBT6cdF6e2hC6jglFglA/YC0RGKBODANBgkqhkiG9w0BAQUFAAOBgQB+zVtnzr+h+1lL1xhyrc3mYsUJGvAH5ZSAqLWFNRbaVR8kUme1FJkT6UaI9PEISDhspSY9fqF2eqYSlqMt5ZGWbDn5ZkQ8B6QB1yrvjUndzgLGy+ksIpwN34GDJyarrupwSrd4h5QMhJI4Fmjf+tDwsTmMESyDm8NJCkWfza2KgQ=="),
			),
		},
		// verify resource import
		{
			Config:            config + IdentityDomainsOAuthClientCertificateRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateIdFunc: getIdentityDomainsImportIdFn("oci_identity_domains_oauth_client_certificate", "oauthClientCertificates"),
			ImportStateVerifyIgnore: []string{
				"authorization",
				"idcs_endpoint",
				"resource_type_schema_version",
				// the following attributes are not present in the response of GET single datasource api
				"cert_end_date",
				"cert_start_date",
				"compartment_ocid",
				"domain_ocid",
				"external_id",
				"sha1thumbprint",
				"sha256thumbprint",
				"tenancy_ocid",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityDomainsOAuthClientCertificateDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).IdentityDomainsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_domains_oauth_client_certificate" {
			noResourceFound = false
			request := oci_identity_domains.GetOAuthClientCertificateRequest{}

			if value, ok := rs.Primary.Attributes["authorization"]; ok {
				request.Authorization = &value
			}

			if value, ok := rs.Primary.Attributes["idcs_endpoint"]; ok {
				client.Host = value
			}

			tmp := rs.Primary.ID
			request.OAuthClientCertificateId = &tmp

			if value, ok := rs.Primary.Attributes["resource_type_schema_version"]; ok {
				request.ResourceTypeSchemaVersion = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")

			res, err := client.GetOAuthClientCertificate(context.Background(), request)

			// The API returns an empty object when the resource is deleted,
			// so if res doesn't have an id, that means the resource is deleted successfully.
			if err == nil && res.Id != nil {
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
	if !acctest.InSweeperExcludeList("IdentityDomainsOAuthClientCertificate") {
		resource.AddTestSweepers("IdentityDomainsOAuthClientCertificate", &resource.Sweeper{
			Name:         "IdentityDomainsOAuthClientCertificate",
			Dependencies: acctest.DependencyGraph["oAuthClientCertificate"],
			F:            sweepIdentityDomainsOAuthClientCertificateResource,
		})
	}
}

func sweepIdentityDomainsOAuthClientCertificateResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	oAuthClientCertificateIds, err := getIdentityDomainsOAuthClientCertificateIds(compartment)
	if err != nil {
		return err
	}
	for _, oAuthClientCertificateId := range oAuthClientCertificateIds {
		if ok := acctest.SweeperDefaultResourceId[oAuthClientCertificateId]; !ok {
			deleteOAuthClientCertificateRequest := oci_identity_domains.DeleteOAuthClientCertificateRequest{}

			deleteOAuthClientCertificateRequest.OAuthClientCertificateId = &oAuthClientCertificateId

			deleteOAuthClientCertificateRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteOAuthClientCertificate(context.Background(), deleteOAuthClientCertificateRequest)
			if error != nil {
				fmt.Printf("Error deleting OAuthClientCertificate %s %s, It is possible that the resource is already deleted. Please verify manually \n", oAuthClientCertificateId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsOAuthClientCertificateIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OAuthClientCertificateId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listOAuthClientCertificatesRequest := oci_identity_domains.ListOAuthClientCertificatesRequest{}
	listOAuthClientCertificatesResponse, err := identityDomainsClient.ListOAuthClientCertificates(context.Background(), listOAuthClientCertificatesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OAuthClientCertificate list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, oAuthClientCertificate := range listOAuthClientCertificatesResponse.Resources {
		id := *oAuthClientCertificate.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OAuthClientCertificateId", id)
	}
	return resourceIds, nil
}
