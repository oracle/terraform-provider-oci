// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_network_firewall "github.com/oracle/oci-go-sdk/v65/networkfirewall"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	nfwDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}
	NetworkFirewallPolicyRequiredOnlyResource = NetworkFirewallPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall_policy", "test_network_firewall_policy", acctest.Required, acctest.Create, networkFirewallPolicyRepresentation)

	NetworkFirewallPolicyResourceConfig = NetworkFirewallPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall_policy", "test_network_firewall_policy", acctest.Optional, acctest.Update, networkFirewallPolicyRepresentation)

	networkFirewallPolicySingularDataSourceRepresentation = map[string]interface{}{
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
	}

	networkFirewallPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: networkFirewallPolicyDataSourceFilterRepresentation}}
	networkFirewallPolicyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`}},
	}

	networkFirewallPolicyRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"application_lists":   []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: networkFirewallApplicationListsRepresentation}},
		"decryption_profiles": []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: networkFirewallPolicyDecryptionProfilesRepresentation1}},
		"decryption_rules":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: networkFirewallPolicyDecryptionRulesRepresentation},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"bar-key": "value"}},
		"ip_address_lists":    []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: networkFirewallIpAddressListsRepresentation}},
		"mapped_secrets":      []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: networkFirewallPolicyMappedSecretsRepresentation}},
		"security_rules":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: networkFirewallPolicySecurityRulesRepresentation},
		"url_lists":           []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: networkFirewallUrlListsRepresentation}},
	}

	networkFirewallIpAddressListsRepresentation = map[string]interface{}{
		"ip_address_list_name": acctest.Representation{RepType: acctest.Required, Create: `hr_source`, Update: `hr_source`},
		"ip_address_list_value": acctest.Representation{RepType: acctest.Optional, Create: []string{"10.2.3.4/24",
			"10.22.2.2"}, Update: []string{"10.2.3.4/24",
			"10.22.2.2"}},
	}
	networkFirewallUrlListsRepresentation = map[string]interface{}{
		"url_list_name":   acctest.Representation{RepType: acctest.Required, Create: `hr`},
		"url_list_values": []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: networkFirewallUrlRepresentation}},
	}
	networkFirewallUrlRepresentation = map[string]interface{}{
		"type":    acctest.Representation{RepType: acctest.Required, Create: `SIMPLE`},
		"pattern": acctest.Representation{RepType: acctest.Optional, Create: `google.com`},
	}
	networkFirewallApplicationListsRepresentation = map[string]interface{}{
		"application_list_name": acctest.Representation{RepType: acctest.Required, Create: `app-1`},
		"application_values":    []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: networkFirewallApplicationRepresentation}},
	}

	networkFirewallApplicationRepresentation = map[string]interface{}{
		"type":      acctest.Representation{RepType: acctest.Required, Create: "ICMP"},
		"icmp_type": acctest.Representation{RepType: acctest.Optional, Create: `5`},
		"icmp_code": acctest.Representation{RepType: acctest.Optional, Create: `2`},
	}

	networkFirewallApplicationListsRepresentation2 = map[string]interface{}{
		"type":      acctest.Representation{RepType: acctest.Required, Create: "ICMP"},
		"key":       acctest.Representation{RepType: acctest.Required, Create: `app-2`},
		"icmp_type": acctest.Representation{RepType: acctest.Optional, Create: `5`},
		"icmp_code": acctest.Representation{RepType: acctest.Optional, Create: `2`},
	}
	networkFirewallPolicyDecryptionProfilesRepresentation1 = map[string]interface{}{
		"key":                                   acctest.Representation{RepType: acctest.Required, Create: `ssl-inbound-1`, Update: `ssl-inbound-1`},
		"is_out_of_capacity_blocked":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"is_unsupported_cipher_blocked":         acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"is_unsupported_version_blocked":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"type":                                  acctest.Representation{RepType: acctest.Required, Create: `SSL_INBOUND_INSPECTION`, Update: `SSL_INBOUND_INSPECTION`},
		"are_certificate_extensions_restricted": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"is_auto_include_alt_name":              acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"is_expired_certificate_blocked":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"is_revocation_status_timeout_blocked":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"is_unknown_revocation_status_blocked":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"is_untrusted_issuer_blocked":           acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
	}
	networkFirewallPolicyDecryptionRulesRepresentation = map[string]interface{}{
		"action":             acctest.Representation{RepType: acctest.Required, Create: `DECRYPT`}, // Update: `DECRYPT`
		"condition":          acctest.RepresentationGroup{RepType: acctest.Required, Group: networkFirewallPolicyDecryptionRulesConditionRepresentation},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `hr-inbound-inspect`, Update: `hr-inbound-inspect`},
		"decryption_profile": acctest.Representation{RepType: acctest.Optional, Create: `ssl-inbound-1`, Update: `ssl-inbound-1`},
		"secret":             acctest.Representation{RepType: acctest.Optional, Create: `cert-inbound`, Update: `cert-inbound`},
	}
	networkFirewallPolicyMappedSecretsRepresentation = map[string]interface{}{
		"key":             acctest.Representation{RepType: acctest.Required, Create: `cert-inbound`, Update: `cert-inbound`},
		"type":            acctest.Representation{RepType: acctest.Required, Create: `SSL_INBOUND_INSPECTION`, Update: `SSL_INBOUND_INSPECTION`},
		"vault_secret_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_vault_secret.test_secret.id}`},
		"version_number":  acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `10`},
	}
	networkFirewallPolicySecurityRulesRepresentation = map[string]interface{}{
		"action":     acctest.Representation{RepType: acctest.Required, Create: `INSPECT`, Update: `INSPECT`},
		"condition":  acctest.RepresentationGroup{RepType: acctest.Required, Group: networkFirewallPolicySecurityRulesConditionRepresentation},
		"name":       acctest.Representation{RepType: acctest.Required, Create: `hr_access`, Update: `hr_access`},
		"inspection": acctest.Representation{RepType: acctest.Optional, Create: `INTRUSION_DETECTION`, Update: `INTRUSION_DETECTION`},
	}
	networkFirewallPolicyDecryptionRulesConditionRepresentation = map[string]interface{}{
		"destinations": acctest.Representation{RepType: acctest.Optional, Create: []string{`hr_source`}, Update: []string{`hr_source`}},
		"sources":      acctest.Representation{RepType: acctest.Optional, Create: []string{`hr_source`}, Update: []string{`hr_source`}},
	}
	networkFirewallPolicySecurityRulesConditionRepresentation = map[string]interface{}{
		"applications": acctest.Representation{RepType: acctest.Optional, Create: []string{`app-1`}, Update: []string{`app-1`}},
		"destinations": acctest.Representation{RepType: acctest.Optional, Create: []string{`hr_source`}, Update: []string{`hr_source`}},
		"sources":      acctest.Representation{RepType: acctest.Optional, Create: []string{`hr_source`}, Update: []string{`hr_source`}},
		"urls":         acctest.Representation{RepType: acctest.Optional, Create: []string{`hr`}, Update: []string{`hr`}},
	}

	NetworkFirewallPolicyResourceDependencies = "" //DefinedTagsDependencies +
	//acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, vaultRepresentation) +
	//acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Required, acctest.Create, secretRepresentation)
)

// issue-routing-tag: network_firewall/default
func TestNetworkFirewallNetworkFirewallPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkFirewallNetworkFirewallPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	vaultId := utils.GetEnvSettingWithBlankDefault("kms_vault_ocid")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	keyId := utils.GetEnvSettingWithBlankDefault("kms_key_ocid")
	keyIdVariableStr := fmt.Sprintf("variable \"key_id\" { default = \"%s\" }\n", keyId)

	resourceName := "oci_network_firewall_network_firewall_policy.test_network_firewall_policy"
	datasourceName := "data.oci_network_firewall_network_firewall_policies.test_network_firewall_policies"
	singularDatasourceName := "data.oci_network_firewall_network_firewall_policy.test_network_firewall_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+NetworkFirewallPolicyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall_policy", "test_network_firewall_policy", acctest.Optional, acctest.Create, networkFirewallPolicyRepresentation), "network_firewall", "networkFirewallPolicy", t)

	acctest.ResourceTest(t, testAccCheckNetworkFirewallNetworkFirewallPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + vaultIdVariableStr + keyIdVariableStr + compartmentIdVariableStr + NetworkFirewallPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall_policy", "test_network_firewall_policy", acctest.Required, acctest.Create, networkFirewallPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + vaultIdVariableStr + keyIdVariableStr + compartmentIdVariableStr + NetworkFirewallPolicyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + vaultIdVariableStr + keyIdVariableStr + compartmentIdVariableStr + NetworkFirewallPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall_policy", "test_network_firewall_policy", acctest.Optional, acctest.Create, networkFirewallPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "application_lists.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "application_lists.0.application_list_name", "app-1"),
				resource.TestCheckResourceAttr(resourceName, "application_lists.0.application_values.0.icmp_type", "5"),
				resource.TestCheckResourceAttr(resourceName, "application_lists.0.application_values.0.icmp_code", "2"),
				resource.TestCheckResourceAttr(resourceName, "url_lists.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "url_lists.0.url_list_values.0.pattern", "google.com"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "decryption_profiles.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "decryption_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "decryption_rules.0.action", "DECRYPT"),
				resource.TestCheckResourceAttr(resourceName, "decryption_rules.0.condition.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "decryption_rules.0.condition.0.destinations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "decryption_rules.0.condition.0.sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "decryption_rules.0.decryption_profile", "ssl-inbound-1"),
				resource.TestCheckResourceAttr(resourceName, "decryption_rules.0.name", "hr-inbound-inspect"),
				resource.TestCheckResourceAttr(resourceName, "decryption_rules.0.secret", "cert-inbound"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ip_address_lists.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ip_address_lists.0.ip_address_list_name", "hr_source"),
				resource.TestCheckResourceAttrSet(resourceName, "is_firewall_attached"),
				resource.TestCheckResourceAttr(resourceName, "mapped_secrets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.0.action", "INSPECT"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.0.condition.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.0.condition.0.applications.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.0.condition.0.destinations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.0.condition.0.sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.0.condition.0.urls.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.0.inspection", "INTRUSION_DETECTION"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.0.name", "hr_access"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + vaultIdVariableStr + keyIdVariableStr + compartmentIdVariableStr + compartmentIdUVariableStr + NetworkFirewallPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall_policy", "test_network_firewall_policy", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(networkFirewallPolicyRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "decryption_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "decryption_rules.0.action", "DECRYPT"),
				resource.TestCheckResourceAttr(resourceName, "decryption_rules.0.condition.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "decryption_rules.0.condition.0.destinations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "decryption_rules.0.condition.0.sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "decryption_rules.0.decryption_profile", "ssl-inbound-1"),
				resource.TestCheckResourceAttr(resourceName, "decryption_rules.0.name", "hr-inbound-inspect"),
				resource.TestCheckResourceAttr(resourceName, "decryption_rules.0.secret", "cert-inbound"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ip_address_lists.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "is_firewall_attached"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.0.action", "INSPECT"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.0.condition.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.0.condition.0.applications.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.0.condition.0.destinations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.0.condition.0.sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.0.condition.0.urls.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.0.inspection", "INTRUSION_DETECTION"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.0.name", "hr_access"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + vaultIdVariableStr + keyIdVariableStr + compartmentIdVariableStr + NetworkFirewallPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall_policy", "test_network_firewall_policy", acctest.Optional, acctest.Update, networkFirewallPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "decryption_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "decryption_rules.0.action", "DECRYPT"),
				resource.TestCheckResourceAttr(resourceName, "decryption_rules.0.condition.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "decryption_rules.0.condition.0.destinations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "decryption_rules.0.condition.0.sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "decryption_rules.0.decryption_profile", "ssl-inbound-1"),
				resource.TestCheckResourceAttr(resourceName, "decryption_rules.0.name", "hr-inbound-inspect"),
				resource.TestCheckResourceAttr(resourceName, "decryption_rules.0.secret", "cert-inbound"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_firewall_attached"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.0.action", "INSPECT"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.0.condition.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.0.condition.0.applications.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.0.condition.0.destinations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.0.condition.0.sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.0.condition.0.urls.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.0.inspection", "INTRUSION_DETECTION"),
				resource.TestCheckResourceAttr(resourceName, "security_rules.0.name", "hr_access"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_network_firewall_network_firewall_policies", "test_network_firewall_policies", acctest.Optional, acctest.Update, networkFirewallPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + NetworkFirewallPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall_policy", "test_network_firewall_policy", acctest.Optional, acctest.Update, networkFirewallPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "network_firewall_policy_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "network_firewall_policy_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_network_firewall_network_firewall_policy", "test_network_firewall_policy", acctest.Required, acctest.Create, networkFirewallPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + NetworkFirewallPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "decryption_rules.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "decryption_rules.0.action", "DECRYPT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "decryption_rules.0.condition.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "decryption_rules.0.condition.0.destinations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "decryption_rules.0.condition.0.sources.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "decryption_rules.0.decryption_profile", "ssl-inbound-1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "decryption_rules.0.name", "hr-inbound-inspect"),
				resource.TestCheckResourceAttr(singularDatasourceName, "decryption_rules.0.secret", "cert-inbound"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ip_address_lists.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_firewall_attached"),
				resource.TestCheckResourceAttr(singularDatasourceName, "security_rules.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "security_rules.0.action", "INSPECT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "security_rules.0.condition.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "security_rules.0.condition.0.applications.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "security_rules.0.condition.0.destinations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "security_rules.0.condition.0.sources.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "security_rules.0.condition.0.urls.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "security_rules.0.inspection", "INTRUSION_DETECTION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "security_rules.0.name", "hr_access"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + NetworkFirewallPolicyRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckNetworkFirewallNetworkFirewallPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).NetworkFirewallClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_network_firewall_network_firewall_policy" {
			noResourceFound = false
			request := oci_network_firewall.GetNetworkFirewallPolicyRequest{}

			tmp := rs.Primary.ID
			request.NetworkFirewallPolicyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")

			response, err := client.GetNetworkFirewallPolicy(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_network_firewall.LifecycleStateDeleted): true,
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

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("NetworkFirewallNetworkFirewallPolicy") {
		resource.AddTestSweepers("NetworkFirewallNetworkFirewallPolicy", &resource.Sweeper{
			Name:         "NetworkFirewallNetworkFirewallPolicy",
			Dependencies: acctest.DependencyGraph["networkFirewallPolicy"],
			F:            sweepNetworkFirewallNetworkFirewallPolicyResource,
		})
	}
}

func sweepNetworkFirewallNetworkFirewallPolicyResource(compartment string) error {
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()
	networkFirewallPolicyIds, err := getNetworkFirewallPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, networkFirewallPolicyId := range networkFirewallPolicyIds {
		if ok := acctest.SweeperDefaultResourceId[networkFirewallPolicyId]; !ok {
			deleteNetworkFirewallPolicyRequest := oci_network_firewall.DeleteNetworkFirewallPolicyRequest{}

			deleteNetworkFirewallPolicyRequest.NetworkFirewallPolicyId = &networkFirewallPolicyId

			deleteNetworkFirewallPolicyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")
			_, error := networkFirewallClient.DeleteNetworkFirewallPolicy(context.Background(), deleteNetworkFirewallPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting NetworkFirewallPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", networkFirewallPolicyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &networkFirewallPolicyId, networkFirewallPolicySweepWaitCondition, time.Duration(3*time.Minute),
				networkFirewallPolicySweepResponseFetchOperation, "network_firewall", true)
		}
	}
	return nil
}

func getNetworkFirewallPolicyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NetworkFirewallPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()

	listNetworkFirewallPoliciesRequest := oci_network_firewall.ListNetworkFirewallPoliciesRequest{}
	listNetworkFirewallPoliciesRequest.CompartmentId = &compartmentId
	listNetworkFirewallPoliciesRequest.LifecycleState = oci_network_firewall.ListNetworkFirewallPoliciesLifecycleStateActive
	listNetworkFirewallPoliciesResponse, err := networkFirewallClient.ListNetworkFirewallPolicies(context.Background(), listNetworkFirewallPoliciesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting NetworkFirewallPolicy list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, networkFirewallPolicy := range listNetworkFirewallPoliciesResponse.Items {
		id := *networkFirewallPolicy.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NetworkFirewallPolicyId", id)
	}
	return resourceIds, nil
}

func networkFirewallPolicySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if networkFirewallPolicyResponse, ok := response.Response.(oci_network_firewall.GetNetworkFirewallPolicyResponse); ok {
		return networkFirewallPolicyResponse.LifecycleState != oci_network_firewall.LifecycleStateDeleted
	}
	return false
}

func networkFirewallPolicySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.NetworkFirewallClient().GetNetworkFirewallPolicy(context.Background(), oci_network_firewall.GetNetworkFirewallPolicyRequest{
		NetworkFirewallPolicyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
