// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsOciConsoleSignOnPolicyConsentSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":                         acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"oci_console_sign_on_policy_consent_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domains_oci_console_sign_on_policy_consents.test_oci_console_sign_on_policy_consents.resources.0.id}`},
	}

	IdentityDomainsOciConsoleSignOnPolicyConsentDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"oci_console_sign_on_policy_consent_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"oci_console_sign_on_policy_consent_filter": acctest.Representation{RepType: acctest.Optional, Create: `ociConsoleSignOnPolicyConsentFilter`},
		"start_index": acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsOciConsoleSignOnPolicyConsentResourceConfig = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsOciConsoleSignOnPolicyConsentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsOciConsoleSignOnPolicyConsentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_domains_oci_console_sign_on_policy_consents.test_oci_console_sign_on_policy_consents"
	singularDatasourceName := "data.oci_identity_domains_oci_console_sign_on_policy_consent.test_oci_console_sign_on_policy_consent"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_oci_console_sign_on_policy_consents", "test_oci_console_sign_on_policy_consents", acctest.Required, acctest.Create, IdentityDomainsOciConsoleSignOnPolicyConsentDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsOciConsoleSignOnPolicyConsentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestCheckResourceAttrSet(datasourceName, "resources.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "resources.0.client_ip"),
				resource.TestCheckResourceAttrSet(datasourceName, "resources.0.change_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "resources.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "resources.0.justification"),
				resource.TestCheckResourceAttrSet(datasourceName, "resources.0.notification_recipients.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "resources.0.reason"),
				resource.TestCheckResourceAttrSet(datasourceName, "resources.0.time_consent_signed"),
				//resource.TestCheckResourceAttr(datasourceName, "oci_console_sign_on_policy_consents.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_oci_console_sign_on_policy_consents", "test_oci_console_sign_on_policy_consents", acctest.Required, acctest.Create, IdentityDomainsOciConsoleSignOnPolicyConsentDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_oci_console_sign_on_policy_consent", "test_oci_console_sign_on_policy_consent", acctest.Required, acctest.Create, IdentityDomainsOciConsoleSignOnPolicyConsentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsOciConsoleSignOnPolicyConsentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "change_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "client_ip"),
				resource.TestCheckResourceAttr(singularDatasourceName, "consent_signed_by.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "justification"),
				resource.TestCheckResourceAttr(singularDatasourceName, "modified_resource.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notification_recipients.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_resource.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "reason"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_consent_signed"),
			),
		},
	})
}
