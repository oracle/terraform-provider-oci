// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	FileStorageMountTargetRequiredOnlyResource = FileStorageMountTargetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target", acctest.Required, acctest.Create, FileStorageMountTargetRepresentation)

	FileStorageFileStorageMountTargetDataSourceRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `mount-target-5`, Update: `displayName2`},
		"id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_mount_target.test_mount_target.id}`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: FileStorageMountTargetDataSourceFilterRepresentation}}
	FileStorageMountTargetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_file_storage_mount_target.test_mount_target.id}`}},
	}

	FileStorageMountTargetRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subnet_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `mount-target-5`, Update: `displayName2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"hostname_label":      acctest.Representation{RepType: acctest.Optional, Create: `hostnamelabel`},
		"idmap_type":          acctest.Representation{RepType: acctest.Optional, Create: `LDAP`, Update: `LDAP`},
		"ip_address":          acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.5`},
		"kerberos":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: FileStorageMountTargetKerberosRepresentation},
		"ldap_idmap":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: FileStorageMountTargetLdapIdmapRepresentation},
		"nsg_ids":             acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{}},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
	}
	FileStorageMountTargetKerberosRepresentation = map[string]interface{}{
		"kerberos_realm":                 acctest.Representation{RepType: acctest.Required, Create: `kerberosRealm`, Update: `kerberosRealm2`},
		"backup_key_tab_secret_version":  acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `0`},
		"current_key_tab_secret_version": acctest.Representation{RepType: acctest.Optional, Create: `1`, Update: `1`},
		"is_kerberos_enabled":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"key_tab_secret_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_keytab_secret.id}`},
	}
	FileStorageMountTargetLdapIdmapRepresentation = map[string]interface{}{
		"group_search_base":               acctest.Representation{RepType: acctest.Required, Create: `groupSearchBase`, Update: `groupSearchBase2`},
		"user_search_base":                acctest.Representation{RepType: acctest.Required, Create: `userSearchBase`, Update: `userSearchBase2`},
		"cache_lifetime_seconds":          acctest.Representation{RepType: acctest.Optional, Create: `300`, Update: `400`},
		"cache_refresh_interval_seconds":  acctest.Representation{RepType: acctest.Optional, Create: `300`, Update: `400`},
		"negative_cache_lifetime_seconds": acctest.Representation{RepType: acctest.Optional, Create: `300`, Update: `400`},
		"outbound_connector1id":           acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_outbound_connector.test_outbound_connector1.id}`},
		"outbound_connector2id":           acctest.Representation{RepType: acctest.Optional, Update: `${oci_file_storage_outbound_connector.test_outbound_connector2.id}`},
		"schema_type":                     acctest.Representation{RepType: acctest.Optional, Create: `RFC2307`},
	}

	FileStorageMountTargetResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
			"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`},
			"dns_label":           acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		AvailabilityDomainConfig + DefinedTagsDependencies

	FileStorageMountTargetResourceKerberosDependencies = acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(KmsVaultRepresentation, map[string]interface{}{
		"vault_type": acctest.Representation{RepType: acctest.Required, Create: `DEFAULT`},
	})) +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(KmsKeyRepresentation, map[string]interface{}{
			"management_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_vault.test_vault.management_endpoint}`},
			"desired_state":       acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_keytab_secret", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(VaultSecretRepresentation, map[string]interface{}{
			"vault_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_vault.test_vault.id}`},
			"key_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_key.test_key.id}`},
			"secret_name": acctest.Representation{RepType: acctest.Required, Create: strconv.Itoa(int(time.Now().Unix())) + "_keytab"},
			"secret_content": acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.RepresentationCopyWithNewProperties(VaultSecretSecretContentRepresentation, map[string]interface{}{
				"content": acctest.Representation{RepType: acctest.Required, Create: `BQIAAAClAAIAI0FEMkNBTkFSWS5QSFhERVZQQ0FOUy5PUkFDTEVWQ04uQ09NAANuZnMARmtlcmJlcm9zLWFwaS1jYW5hcnktbW91bnQtdGFyZ2V0LTEuYWQyY2FuYXJ5LnBoeGRldnBjYW5zLm9yYWNsZXZjbi5jb20AAAABYgMUPgIAEgAgIvKmyzN+v/xsEQpwSzwxfFCEwtbV5ozYkk8VAmx9NhQAAAAC`},
			})},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_obc_pwd_secret", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(VaultSecretRepresentation, map[string]interface{}{
			"vault_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_vault.test_vault.id}`},
			"key_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_key.test_key.id}`},
			"secret_name": acctest.Representation{RepType: acctest.Required, Create: strconv.Itoa(int(time.Now().Unix())) + "_pwd"},
			"secret_content": acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.RepresentationCopyWithNewProperties(VaultSecretSecretContentRepresentation, map[string]interface{}{
				"content": acctest.Representation{RepType: acctest.Required, Create: `dGVzdHB3ZAo=`},
			})},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_outbound_connector", "test_outbound_connector1", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(FileStorageOutboundConnectorRepresentation, map[string]interface{}{
			"password_secret_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_vault_secret.test_obc_pwd_secret.id}`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_outbound_connector", "test_outbound_connector2", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(FileStorageOutboundConnectorRepresentation, map[string]interface{}{
			"password_secret_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_vault_secret.test_obc_pwd_secret.id}`},
		}))
)

// issue-routing-tag: file_storage/default
func TestFileStorageMountTargetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFileStorageMountTargetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_file_storage_mount_target.test_mount_target"
	datasourceName := "data.oci_file_storage_mount_targets.test_mount_targets"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FileStorageMountTargetResourceDependencies+FileStorageMountTargetResourceKerberosDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target", acctest.Optional, acctest.Create, FileStorageMountTargetRepresentation), "filestorage", "mountTarget", t)

	acctest.ResourceTest(t, testAccCheckFileStorageMountTargetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + FileStorageMountTargetResourceDependencies + FileStorageMountTargetResourceKerberosDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target", acctest.Required, acctest.Create, FileStorageMountTargetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "export_set_id"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FileStorageMountTargetResourceDependencies + FileStorageMountTargetResourceKerberosDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + FileStorageMountTargetResourceDependencies + FileStorageMountTargetResourceKerberosDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target", acctest.Optional, acctest.Create, FileStorageMountTargetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "mount-target-5"),
				resource.TestCheckResourceAttrSet(resourceName, "export_set_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnamelabel"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "idmap_type", "LDAP"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.5"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "private_ip_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "private_ip_ids.0"),
				resource.TestCheckResourceAttr(resourceName, "state", string(oci_file_storage.MountTargetLifecycleStateActive)),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "kerberos.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "kerberos.0.backup_key_tab_secret_version", "0"),
				resource.TestCheckResourceAttr(resourceName, "kerberos.0.current_key_tab_secret_version", "1"),
				resource.TestCheckResourceAttr(resourceName, "kerberos.0.is_kerberos_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "kerberos.0.kerberos_realm", "kerberosRealm"),
				resource.TestCheckResourceAttrSet(resourceName, "kerberos.0.key_tab_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "ldap_idmap.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ldap_idmap.0.cache_lifetime_seconds", "300"),
				resource.TestCheckResourceAttr(resourceName, "ldap_idmap.0.cache_refresh_interval_seconds", "300"),
				resource.TestCheckResourceAttr(resourceName, "ldap_idmap.0.group_search_base", "groupSearchBase"),
				resource.TestCheckResourceAttr(resourceName, "ldap_idmap.0.user_search_base", "userSearchBase"),
				resource.TestCheckResourceAttr(resourceName, "ldap_idmap.0.negative_cache_lifetime_seconds", "300"),
				resource.TestCheckResourceAttrSet(resourceName, "ldap_idmap.0.outbound_connector1id"),
				resource.TestCheckResourceAttr(resourceName, "ldap_idmap.0.schema_type", "RFC2307"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + FileStorageMountTargetResourceDependencies + FileStorageMountTargetResourceKerberosDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(FileStorageMountTargetRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "mount-target-5"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnamelabel"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "idmap_type", "LDAP"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.5"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "private_ip_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "state", string(oci_file_storage.MountTargetLifecycleStateActive)),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "kerberos.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "kerberos.0.backup_key_tab_secret_version", "0"),
				resource.TestCheckResourceAttr(resourceName, "kerberos.0.current_key_tab_secret_version", "1"),
				resource.TestCheckResourceAttr(resourceName, "kerberos.0.is_kerberos_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "kerberos.0.kerberos_realm", "kerberosRealm"),
				resource.TestCheckResourceAttrSet(resourceName, "kerberos.0.key_tab_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "ldap_idmap.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ldap_idmap.0.cache_lifetime_seconds", "300"),
				resource.TestCheckResourceAttr(resourceName, "ldap_idmap.0.cache_refresh_interval_seconds", "300"),
				resource.TestCheckResourceAttr(resourceName, "ldap_idmap.0.group_search_base", "groupSearchBase"),
				resource.TestCheckResourceAttr(resourceName, "ldap_idmap.0.user_search_base", "userSearchBase"),
				resource.TestCheckResourceAttr(resourceName, "ldap_idmap.0.negative_cache_lifetime_seconds", "300"),
				resource.TestCheckResourceAttrSet(resourceName, "ldap_idmap.0.outbound_connector1id"),
				resource.TestCheckResourceAttr(resourceName, "ldap_idmap.0.schema_type", "RFC2307"),

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
			Config: config + compartmentIdVariableStr + FileStorageMountTargetResourceDependencies + FileStorageMountTargetResourceKerberosDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target", acctest.Optional, acctest.Update, FileStorageMountTargetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "export_set_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnamelabel"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "idmap_type", "LDAP"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.5"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "private_ip_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "state", string(oci_file_storage.MountTargetLifecycleStateActive)),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "kerberos.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "kerberos.0.backup_key_tab_secret_version", "0"),
				resource.TestCheckResourceAttr(resourceName, "kerberos.0.current_key_tab_secret_version", "1"),
				resource.TestCheckResourceAttr(resourceName, "kerberos.0.is_kerberos_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "kerberos.0.kerberos_realm", "kerberosRealm2"),
				resource.TestCheckResourceAttrSet(resourceName, "kerberos.0.key_tab_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "ldap_idmap.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ldap_idmap.0.cache_lifetime_seconds", "400"),
				resource.TestCheckResourceAttr(resourceName, "ldap_idmap.0.cache_refresh_interval_seconds", "400"),
				resource.TestCheckResourceAttr(resourceName, "ldap_idmap.0.group_search_base", "groupSearchBase2"),
				resource.TestCheckResourceAttr(resourceName, "ldap_idmap.0.user_search_base", "userSearchBase2"),
				resource.TestCheckResourceAttr(resourceName, "ldap_idmap.0.negative_cache_lifetime_seconds", "400"),
				resource.TestCheckResourceAttrSet(resourceName, "ldap_idmap.0.outbound_connector1id"),
				resource.TestCheckResourceAttrSet(resourceName, "ldap_idmap.0.outbound_connector2id"),
				resource.TestCheckResourceAttr(resourceName, "ldap_idmap.0.schema_type", "RFC2307"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_file_storage_mount_targets", "test_mount_targets", acctest.Optional, acctest.Update, FileStorageFileStorageMountTargetDataSourceRepresentation) +
				compartmentIdVariableStr + FileStorageMountTargetResourceDependencies + FileStorageMountTargetResourceKerberosDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target", acctest.Optional, acctest.Update, FileStorageMountTargetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "mount_targets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "mount_targets.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "mount_targets.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "mount_targets.0.export_set_id"),
				resource.TestCheckResourceAttr(datasourceName, "mount_targets.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "mount_targets.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "mount_targets.nsg_ids.#", "0"),
				resource.TestCheckResourceAttrSet(datasourceName, "mount_targets.0.private_ip_ids.#"),
				resource.TestCheckResourceAttr(datasourceName, "mount_targets.0.state", string(oci_file_storage.MountTargetLifecycleStateActive)),
				resource.TestCheckResourceAttrSet(datasourceName, "mount_targets.0.subnet_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "mount_targets.0.time_created"),
			),
		},
		// verify updates to kerberos and ldap parameters
		{
			Config: config + compartmentIdVariableStr + FileStorageMountTargetResourceDependencies + FileStorageMountTargetResourceKerberosDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(FileStorageMountTargetRepresentation, map[string]interface{}{
					"kerberos": acctest.RepresentationGroup{RepType: acctest.Optional, Group: acctest.RepresentationCopyWithNewProperties(FileStorageMountTargetKerberosRepresentation, map[string]interface{}{
						"is_kerberos_enabled":            acctest.Representation{RepType: acctest.Optional, Update: `false`},
						"key_tab_secret_id":              acctest.Representation{RepType: acctest.Optional, Update: ``},
						"current_key_tab_secret_version": acctest.Representation{RepType: acctest.Optional, Update: `0`},
					})},
					"ldap_idmap": acctest.RepresentationGroup{RepType: acctest.Optional, Group: acctest.RepresentationCopyWithNewProperties(FileStorageMountTargetLdapIdmapRepresentation, map[string]interface{}{
						"outbound_connector2id": acctest.Representation{RepType: acctest.Optional, Update: ``},
					})},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "kerberos.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "kerberos.0.backup_key_tab_secret_version", "0"),
				resource.TestCheckResourceAttr(resourceName, "kerberos.0.current_key_tab_secret_version", "0"),
				resource.TestCheckResourceAttr(resourceName, "kerberos.0.is_kerberos_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "kerberos.0.key_tab_secret_id", ""),
				resource.TestCheckResourceAttr(resourceName, "ldap_idmap.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ldap_idmap.0.outbound_connector2id", ""),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify resource import
		{
			Config:                  config + FileStorageMountTargetRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

// issue-routing-tag: file_storage/default
func TestFileStorageMountTargetResource_failedWorkRequest(t *testing.T) {
	httpreplay.SetScenario("TestFileStorageMountTargetResource_failedWorkRequest")
	defer httpreplay.SaveScenario()
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_file_storage_mount_target.test_mount_target2"

	acctest.ResourceTest(t, testAccCheckFileStorageMountTargetDestroy, []resource.TestStep{
		// verify resource creation fails for the second mount target with the same ip_address
		{
			Config: config + compartmentIdVariableStr + FileStorageMountTargetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target1", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithRemovedProperties(FileStorageMountTargetRepresentation, []string{"idmap_type", "kerberos", "ldap_idmap"})) +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target2", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithRemovedProperties(FileStorageMountTargetRepresentation, []string{"idmap_type", "kerberos", "ldap_idmap"})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.5"),
			),
			ExpectError: regexp.MustCompile("Resource creation failed"),
		},
	})
}

func testAccCheckFileStorageMountTargetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FileStorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_file_storage_mount_target" {
			noResourceFound = false
			request := oci_file_storage.GetMountTargetRequest{}

			tmp := rs.Primary.ID
			request.MountTargetId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "file_storage")

			response, err := client.GetMountTarget(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_file_storage.MountTargetLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("FileStorageMountTarget") {
		resource.AddTestSweepers("FileStorageMountTarget", &resource.Sweeper{
			Name:         "FileStorageMountTarget",
			Dependencies: acctest.DependencyGraph["mountTarget"],
			F:            sweepFileStorageMountTargetResource,
		})
	}
}

func sweepFileStorageMountTargetResource(compartment string) error {
	fileStorageClient := acctest.GetTestClients(&schema.ResourceData{}).FileStorageClient()
	mountTargetIds, err := getFileStorageMountTargetIds(compartment)
	if err != nil {
		return err
	}
	for _, mountTargetId := range mountTargetIds {
		if ok := acctest.SweeperDefaultResourceId[mountTargetId]; !ok {
			deleteMountTargetRequest := oci_file_storage.DeleteMountTargetRequest{}

			deleteMountTargetRequest.MountTargetId = &mountTargetId

			deleteMountTargetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "file_storage")
			_, error := fileStorageClient.DeleteMountTarget(context.Background(), deleteMountTargetRequest)
			if error != nil {
				fmt.Printf("Error deleting MountTarget %s %s, It is possible that the resource is already deleted. Please verify manually \n", mountTargetId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &mountTargetId, FileStorageMountTargetSweepWaitCondition, time.Duration(3*time.Minute),
				FileStorageMountTargetSweepResponseFetchOperation, "file_storage", true)
		}
	}
	return nil
}

func getFileStorageMountTargetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MountTargetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fileStorageClient := acctest.GetTestClients(&schema.ResourceData{}).FileStorageClient()

	listMountTargetsRequest := oci_file_storage.ListMountTargetsRequest{}
	listMountTargetsRequest.CompartmentId = &compartmentId

	availabilityDomains, err := acctest.GetAvalabilityDomains(compartment)
	if err != nil {
		return resourceIds, fmt.Errorf("Error getting availabilityDomains required for MountTarget list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, availabilityDomainName := range availabilityDomains {
		listMountTargetsRequest.AvailabilityDomain = &availabilityDomainName

		listMountTargetsRequest.LifecycleState = oci_file_storage.ListMountTargetsLifecycleStateActive
		listMountTargetsResponse, err := fileStorageClient.ListMountTargets(context.Background(), listMountTargetsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting MountTarget list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, mountTarget := range listMountTargetsResponse.Items {
			id := *mountTarget.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MountTargetId", id)
		}

	}
	return resourceIds, nil
}

func FileStorageMountTargetSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if mountTargetResponse, ok := response.Response.(oci_file_storage.GetMountTargetResponse); ok {
		return mountTargetResponse.LifecycleState != oci_file_storage.MountTargetLifecycleStateDeleted
	}
	return false
}

func FileStorageMountTargetSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FileStorageClient().GetMountTarget(context.Background(), oci_file_storage.GetMountTargetRequest{
		MountTargetId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
