// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FileStorageOutboundConnectorRequiredOnlyResource = FileStorageOutboundConnectorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_outbound_connector", "test_outbound_connector", acctest.Required, acctest.Create, FileStorageOutboundConnectorRepresentation)

	FileStorageOutboundConnectorResourceConfig = FileStorageOutboundConnectorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_outbound_connector", "test_outbound_connector", acctest.Optional, acctest.Update, FileStorageOutboundConnectorRepresentation)

	FileStorageOutboundConnectorSingularDataSourceRepresentation = map[string]interface{}{
		"outbound_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_file_storage_outbound_connector.test_outbound_connector.id}`},
	}

	FileStorageOutboundConnectorDataSourceRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `outbound-connector-4`, Update: `displayName2`},
		"id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_outbound_connector.test_outbound_connector.id}`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: FileStorageOutboundConnectorDataSourceFilterRepresentation}}
	FileStorageOutboundConnectorDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_file_storage_outbound_connector.test_outbound_connector.id}`}},
	}

	FileStorageOutboundConnectorRepresentation = map[string]interface{}{
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"bind_distinguished_name": acctest.Representation{RepType: acctest.Required, Create: `bindDistinguishedName`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"connector_type":          acctest.Representation{RepType: acctest.Required, Create: `LDAPBIND`},
		"endpoints":               acctest.RepresentationGroup{RepType: acctest.Required, Group: FileStorageOutboundConnectorEndpointsRepresentation},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `outbound-connector-4`, Update: `displayName2`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"password_secret_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_vault_secret.test_obc_pwd_secret.id}`},
		"password_secret_version": acctest.Representation{RepType: acctest.Required, Create: `1`},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRep},
	}
	FileStorageOutboundConnectorRepresentationWithFullLock = map[string]interface{}{
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"bind_distinguished_name": acctest.Representation{RepType: acctest.Required, Create: `bindDistinguishedName`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"connector_type":          acctest.Representation{RepType: acctest.Required, Create: `LDAPBIND`},
		"endpoints":               acctest.RepresentationGroup{RepType: acctest.Required, Group: FileStorageOutboundConnectorEndpointsRepresentation},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `outbound-connector-4`, Update: `displayName2`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"locks":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: FileStorageOutboundConnectoFullLocksRepresentation},
		"is_lock_override":        acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
		"password_secret_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_vault_secret.test_obc_pwd_secret.id}`},
		"password_secret_version": acctest.Representation{RepType: acctest.Required, Create: `1`},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRep},
	}
	FileStorageOutboundConnectorRepresentationWithDeleteLock = map[string]interface{}{
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"bind_distinguished_name": acctest.Representation{RepType: acctest.Required, Create: `bindDistinguishedName`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"connector_type":          acctest.Representation{RepType: acctest.Required, Create: `LDAPBIND`},
		"endpoints":               acctest.RepresentationGroup{RepType: acctest.Required, Group: FileStorageOutboundConnectorEndpointsRepresentation},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `outbound-connector-4`, Update: `displayName2`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"locks":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: FileStorageOutboundConnectorDeleteLocksRepresentation},
		"is_lock_override":        acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
		"password_secret_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_vault_secret.test_obc_pwd_secret.id}`},
		"password_secret_version": acctest.Representation{RepType: acctest.Required, Create: `1`},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRep},
	}
	FileStorageOutboundConnectorEndpointsRepresentation = map[string]interface{}{
		"hostname": acctest.Representation{RepType: acctest.Required, Create: `hostname`},
		"port":     acctest.Representation{RepType: acctest.Required, Create: `10`},
	}
	FileStorageOutboundConnectoFullLocksRepresentation = map[string]interface{}{
		"type":    acctest.Representation{RepType: acctest.Required, Create: `FULL`},
		"message": acctest.Representation{RepType: acctest.Optional, Create: `message`},
	}
	FileStorageOutboundConnectorDeleteLocksRepresentation = map[string]interface{}{
		"type":    acctest.Representation{RepType: acctest.Required, Create: `DELETE`},
		"message": acctest.Representation{RepType: acctest.Optional, Create: `message`},
	}

	FileStorageOutboundConnectorResourceDependencies = AvailabilityDomainConfig + DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(KmsVaultRepresentation, map[string]interface{}{
			"vault_type": acctest.Representation{RepType: acctest.Required, Create: `DEFAULT`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_key", "test_key", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(KmsKeyRepresentation, map[string]interface{}{
			"management_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_vault.test_vault.management_endpoint}`},
			"desired_state":       acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_obc_pwd_secret", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(VaultSecretRepresentation, map[string]interface{}{
			"vault_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_vault.test_vault.id}`},
			"key_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_kms_key.test_key.id}`},
			"secret_content": acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.RepresentationCopyWithNewProperties(VaultSecretSecretContentRepresentation, map[string]interface{}{
				"content": acctest.Representation{RepType: acctest.Required, Create: `dGVzdHB3ZAo=`},
			})},
		}))
)

// issue-routing-tag: file_storage/default
func TestFileStorageOutboundConnectorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFileStorageOutboundConnectorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_file_storage_outbound_connector.test_outbound_connector"
	datasourceName := "data.oci_file_storage_outbound_connectors.test_outbound_connectors"
	singularDatasourceName := "data.oci_file_storage_outbound_connector.test_outbound_connector"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+FileStorageOutboundConnectorResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_outbound_connector", "test_outbound_connector", acctest.Optional, acctest.Create, FileStorageOutboundConnectorRepresentation), "filestorage", "outboundConnector", t)

	acctest.ResourceTest(t, testAccCheckFileStorageOutboundConnectorDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + FileStorageOutboundConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_outbound_connector", "test_outbound_connector", acctest.Required, acctest.Create, FileStorageOutboundConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "bind_distinguished_name", "bindDistinguishedName"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connector_type", "LDAPBIND"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.hostname", "hostname"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.port", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "password_secret_version", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FileStorageOutboundConnectorResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + FileStorageOutboundConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_outbound_connector", "test_outbound_connector", acctest.Optional, acctest.Create, FileStorageOutboundConnectorRepresentationWithDeleteLock),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "bind_distinguished_name", "bindDistinguishedName"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connector_type", "LDAPBIND"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "outbound-connector-4"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.hostname", "hostname"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.port", "10"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "locks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "locks.0.message", "message"),
				resource.TestCheckResourceAttrSet(resourceName, "locks.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "locks.0.type", "DELETE"),
				resource.TestCheckResourceAttrSet(resourceName, "password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "password_secret_version", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + FileStorageOutboundConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_outbound_connector", "test_outbound_connector", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(FileStorageOutboundConnectorRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "bind_distinguished_name", "bindDistinguishedName"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "connector_type", "LDAPBIND"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "outbound-connector-4"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.hostname", "hostname"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.port", "10"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "locks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "locks.0.message", "message"),
				resource.TestCheckResourceAttrSet(resourceName, "locks.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "locks.0.type", "DELETE"),
				resource.TestCheckResourceAttrSet(resourceName, "password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "password_secret_version", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + FileStorageOutboundConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_outbound_connector", "test_outbound_connector", acctest.Optional, acctest.Update, FileStorageOutboundConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "bind_distinguished_name", "bindDistinguishedName"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connector_type", "LDAPBIND"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.hostname", "hostname"),
				resource.TestCheckResourceAttr(resourceName, "endpoints.0.port", "10"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "password_secret_id"),
				resource.TestCheckResourceAttr(resourceName, "password_secret_version", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_file_storage_outbound_connectors", "test_outbound_connectors", acctest.Optional, acctest.Update, FileStorageOutboundConnectorDataSourceRepresentation) +
				compartmentIdVariableStr + FileStorageOutboundConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_outbound_connector", "test_outbound_connector", acctest.Optional, acctest.Update, FileStorageOutboundConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "outbound_connectors.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "outbound_connectors.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "outbound_connectors.0.bind_distinguished_name", "bindDistinguishedName"),
				resource.TestCheckResourceAttr(datasourceName, "outbound_connectors.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "outbound_connectors.0.connector_type", "LDAPBIND"),
				resource.TestCheckResourceAttr(datasourceName, "outbound_connectors.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "outbound_connectors.0.endpoints.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "outbound_connectors.0.endpoints.0.hostname", "hostname"),
				resource.TestCheckResourceAttr(datasourceName, "outbound_connectors.0.endpoints.0.port", "10"),
				resource.TestCheckResourceAttr(datasourceName, "outbound_connectors.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "outbound_connectors.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "outbound_connectors.0.locks.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "outbound_connectors.0.locks.0.message", "message"),
				resource.TestCheckResourceAttrSet(datasourceName, "outbound_connectors.0.locks.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "outbound_connectors.0.locks.0.type", "DELETE"),
				resource.TestCheckResourceAttrSet(datasourceName, "outbound_connectors.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "outbound_connectors.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_file_storage_outbound_connector", "test_outbound_connector", acctest.Required, acctest.Create, FileStorageOutboundConnectorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + FileStorageOutboundConnectorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "outbound_connector_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "bind_distinguished_name", "bindDistinguishedName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "connector_type", "LDAPBIND"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "endpoints.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "endpoints.0.hostname", "hostname"),
				resource.TestCheckResourceAttr(singularDatasourceName, "endpoints.0.port", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "locks.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "locks.0.message", "message"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "locks.0.time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "locks.0.type", "DELETE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "password_secret_version", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + FileStorageOutboundConnectorResourceDependencies,
		},
		// 		verify Create with FULL lock
		{
			Config: config + compartmentIdVariableStr + FileStorageOutboundConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_file_storage_outbound_connector", "test_outbound_connector", acctest.Optional, acctest.Create, FileStorageOutboundConnectorRepresentationWithFullLock),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "locks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "locks.0.message", "message"),
				resource.TestCheckResourceAttrSet(resourceName, "locks.0.time_created"),
				resource.TestCheckResourceAttr(resourceName, "locks.0.type", "FULL"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify resource import
		{
			Config:                  config + FileStorageOutboundConnectorRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"is_lock_override"},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckFileStorageOutboundConnectorDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).FileStorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_file_storage_outbound_connector" {
			noResourceFound = false
			request := oci_file_storage.GetOutboundConnectorRequest{}

			tmp := rs.Primary.ID
			request.OutboundConnectorId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "file_storage")

			response, err := client.GetOutboundConnector(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_file_storage.OutboundConnectorLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
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
	if !acctest.InSweeperExcludeList("FileStorageOutboundConnector") {
		resource.AddTestSweepers("FileStorageOutboundConnector", &resource.Sweeper{
			Name:         "FileStorageOutboundConnector",
			Dependencies: acctest.DependencyGraph["outboundConnector"],
			F:            sweepFileStorageOutboundConnectorResource,
		})
	}
}

func sweepFileStorageOutboundConnectorResource(compartment string) error {
	fileStorageClient := acctest.GetTestClients(&schema.ResourceData{}).FileStorageClient()
	outboundConnectorIds, err := getFileStorageOutboundConnectorIds(compartment)
	if err != nil {
		return err
	}
	for _, outboundConnectorId := range outboundConnectorIds {
		if ok := acctest.SweeperDefaultResourceId[outboundConnectorId]; !ok {
			deleteOutboundConnectorRequest := oci_file_storage.DeleteOutboundConnectorRequest{}

			deleteOutboundConnectorRequest.OutboundConnectorId = &outboundConnectorId

			deleteOutboundConnectorRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "file_storage")
			_, error := fileStorageClient.DeleteOutboundConnector(context.Background(), deleteOutboundConnectorRequest)
			if error != nil {
				fmt.Printf("Error deleting OutboundConnector %s %s, It is possible that the resource is already deleted. Please verify manually \n", outboundConnectorId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &outboundConnectorId, FileStorageOutboundConnectorSweepWaitCondition, time.Duration(3*time.Minute),
				FileStorageOutboundConnectorSweepResponseFetchOperation, "file_storage", true)
		}
	}
	return nil
}

func getFileStorageOutboundConnectorIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OutboundConnectorId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	fileStorageClient := acctest.GetTestClients(&schema.ResourceData{}).FileStorageClient()

	listOutboundConnectorsRequest := oci_file_storage.ListOutboundConnectorsRequest{}
	listOutboundConnectorsRequest.CompartmentId = &compartmentId

	availabilityDomains, err := acctest.GetAvalabilityDomains(compartment)
	if err != nil {
		return resourceIds, fmt.Errorf("Error getting availabilityDomains required for OutboundConnector list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, availabilityDomainName := range availabilityDomains {
		listOutboundConnectorsRequest.AvailabilityDomain = &availabilityDomainName

		listOutboundConnectorsRequest.LifecycleState = oci_file_storage.ListOutboundConnectorsLifecycleStateActive
		listOutboundConnectorsResponse, err := fileStorageClient.ListOutboundConnectors(context.Background(), listOutboundConnectorsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting OutboundConnector list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, outboundConnector := range listOutboundConnectorsResponse.Items {
			id := *outboundConnector.GetId()
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OutboundConnectorId", id)
		}

	}
	return resourceIds, nil
}

func FileStorageOutboundConnectorSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if outboundConnectorResponse, ok := response.Response.(oci_file_storage.GetOutboundConnectorResponse); ok {
		return outboundConnectorResponse.GetLifecycleState() != oci_file_storage.OutboundConnectorLifecycleStateDeleted
	}
	return false
}

func FileStorageOutboundConnectorSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.FileStorageClient().GetOutboundConnector(context.Background(), oci_file_storage.GetOutboundConnectorRequest{
		OutboundConnectorId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
