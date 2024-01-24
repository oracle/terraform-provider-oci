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
	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	MysqlChannelRequiredOnlyResource = MysqlChannelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_channel", "test_channel", acctest.Required, acctest.Create, MysqlChannelRepresentation)

	MysqlMysqlChannelSingularDataSourceRepresentation = map[string]interface{}{
		"channel_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_mysql_channel.test_channel.id}`},
	}

	MysqlMysqlChannelDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"channel_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_mysql_channel.test_channel.id}`},
		"db_system_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"is_enabled":     acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: MysqlChannelDataSourceFilterRepresentation}}
	MysqlChannelDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_mysql_channel.test_channel.id}`}},
	}

	MysqlChannelRepresentation = map[string]interface{}{
		"source":         acctest.RepresentationGroup{RepType: acctest.Required, Group: MysqlChannelSourceRepresentation},
		"target":         acctest.RepresentationGroup{RepType: acctest.Required, Group: MysqlChannelTargetRepresentation},
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_enabled":     acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesForMysqlChannel},
	}

	ignoreDefinedTagsChangesForMysqlChannel = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{"defined_tags"}},
	}

	MysqlChannelSourceRepresentation = map[string]interface{}{
		"hostname":                        acctest.Representation{RepType: acctest.Required, Create: `hostname.my.company.com`, Update: `hostname2.my.company.com`},
		"password":                        acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"source_type":                     acctest.Representation{RepType: acctest.Required, Create: `MYSQL`},
		"ssl_mode":                        acctest.Representation{RepType: acctest.Required, Create: `REQUIRED`, Update: `VERIFY_CA`},
		"username":                        acctest.Representation{RepType: acctest.Required, Create: `username`, Update: `username2`},
		"anonymous_transactions_handling": acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlChannelSourceManualAnonymousTransactionsHandlingRepresentation},
		"port":                            acctest.Representation{RepType: acctest.Optional, Create: `3300`, Update: `3306`},
	}

	MysqlChannelSourceRepresentationWithErrorOnAnonymous = acctest.GetUpdatedRepresentationCopy(
		"anonymous_transactions_handling",
		acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlChannelSourceErrorAnonymousTransactionsHandlingRepresentation},
		MysqlChannelSourceRepresentation)

	MysqlChannelSourceRepresentationWithCertificateAndErrorOnAnonymous = acctest.RepresentationCopyWithNewProperties(
		MysqlChannelSourceRepresentationWithErrorOnAnonymous,
		map[string]interface{}{"ssl_ca_certificate": acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlChannelSourceSslCaCertificateRepresentation}})

	MysqlChannelTargetRepresentation = map[string]interface{}{
		"db_system_id":                        acctest.Representation{RepType: acctest.Required, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
		"target_type":                         acctest.Representation{RepType: acctest.Required, Create: `DBSYSTEM`},
		"applier_username":                    acctest.Representation{RepType: acctest.Optional, Create: `adminUser`},
		"channel_name":                        acctest.Representation{RepType: acctest.Optional, Create: `channelname`, Update: `channelname2`},
		"delay_in_seconds":                    acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"filters":                             acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlChannelTargetFiltersRepresentation},
		"tables_without_primary_key_handling": acctest.Representation{RepType: acctest.Optional, Create: `RAISE_ERROR`, Update: `ALLOW`},
	}

	MysqlChannelSourceManualAnonymousTransactionsHandlingRepresentation = map[string]interface{}{
		"policy":                       acctest.Representation{RepType: acctest.Required, Create: `ASSIGN_MANUAL_UUID`, Update: `ERROR_ON_ANONYMOUS`},
		"last_configured_log_filename": acctest.Representation{RepType: acctest.Optional, Create: `lastConfiguredLogFilename2`},
		"last_configured_log_offset":   acctest.Representation{RepType: acctest.Optional, Create: `11`},
		"uuid":                         acctest.Representation{RepType: acctest.Optional, Create: `ae1b699d-0036-49c4-900c-4fc43335dfb2`},
	}

	MysqlChannelSourceErrorAnonymousTransactionsHandlingRepresentation = map[string]interface{}{
		"policy": acctest.Representation{RepType: acctest.Required, Update: `ERROR_ON_ANONYMOUS`},
	}

	MysqlChannelSourceSslCaCertificateRepresentation = map[string]interface{}{
		"certificate_type": acctest.Representation{RepType: acctest.Required, Update: "PEM"},
		"contents":         acctest.Representation{RepType: acctest.Required, Update: "${var.ca_certificate_value}"},
	}

	MysqlChannelTargetFiltersRepresentation = map[string]interface{}{
		"type":  acctest.Representation{RepType: acctest.Required, Create: `REPLICATE_DO_DB`, Update: `REPLICATE_IGNORE_DB`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`, Update: `value2`},
	}

	ChannelWithOptionalsResource = MysqlChannelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_channel", "test_channel", acctest.Optional, acctest.Create, MysqlChannelRepresentation)

	ChannelUpdateResource = MysqlChannelResourceDependencies + acctest.GenerateResourceFromRepresentationMap(
		"oci_mysql_channel", "test_channel", acctest.Optional, acctest.Update, acctest.GetUpdatedRepresentationCopy(
			"source", acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlChannelSourceRepresentationWithCertificateAndErrorOnAnonymous}, MysqlChannelRepresentation))

	MysqlChannelResourceDependencies = MysqlMysqlDbSystemResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Create, MysqlMysqlDbSystemRepresentation) + caCertificateVariableStr
)

// issue-routing-tag: mysql/default
func TestMysqlChannelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMysqlChannelResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_channel.test_channel"
	datasourceName := "data.oci_mysql_channels.test_channels"
	singularDatasourceName := "data.oci_mysql_channel.test_channel"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MysqlChannelResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_channel", "test_channel", acctest.Optional, acctest.Create, MysqlChannelRepresentation), "mysql", "channel", t)

	acctest.ResourceTest(t, testAccCheckMysqlChannelDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + MysqlChannelRequiredOnlyResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.hostname", "hostname.my.company.com"),
				resource.TestCheckResourceAttr(resourceName, "source.0.password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "source.0.source_type", "MYSQL"),
				resource.TestCheckResourceAttr(resourceName, "source.0.ssl_mode", "REQUIRED"),
				resource.TestCheckResourceAttr(resourceName, "source.0.username", "username"),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "target.0.target_type", "DBSYSTEM"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + MysqlChannelResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ChannelWithOptionalsResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.anonymous_transactions_handling.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.anonymous_transactions_handling.0.last_configured_log_filename", "lastConfiguredLogFilename2"),
				resource.TestCheckResourceAttr(resourceName, "source.0.anonymous_transactions_handling.0.last_configured_log_offset", "11"),
				resource.TestCheckResourceAttr(resourceName, "source.0.anonymous_transactions_handling.0.policy", "ASSIGN_MANUAL_UUID"),
				resource.TestCheckResourceAttr(resourceName, "source.0.anonymous_transactions_handling.0.uuid", "ae1b699d-0036-49c4-900c-4fc43335dfb2"),
				resource.TestCheckResourceAttr(resourceName, "source.0.hostname", "hostname.my.company.com"),
				resource.TestCheckResourceAttr(resourceName, "source.0.password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "source.0.port", "3300"),
				resource.TestCheckResourceAttr(resourceName, "source.0.source_type", "MYSQL"),
				resource.TestCheckResourceAttr(resourceName, "source.0.ssl_mode", "REQUIRED"),
				resource.TestCheckResourceAttr(resourceName, "source.0.ssl_ca_certificate.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "source.0.username", "username"),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.applier_username", "adminUser"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.channel_name"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "target.0.delay_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "target.0.filters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.filters.0.type", "REPLICATE_DO_DB"),
				resource.TestCheckResourceAttr(resourceName, "target.0.filters.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "target.0.tables_without_primary_key_handling", "RAISE_ERROR"),
				resource.TestCheckResourceAttr(resourceName, "target.0.target_type", "DBSYSTEM"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ChannelUpdateResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.anonymous_transactions_handling.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.anonymous_transactions_handling.0.policy", "ERROR_ON_ANONYMOUS"),
				resource.TestCheckResourceAttr(resourceName, "source.0.hostname", "hostname2.my.company.com"),
				resource.TestCheckResourceAttr(resourceName, "source.0.password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "source.0.port", "3306"),
				resource.TestCheckResourceAttr(resourceName, "source.0.source_type", "MYSQL"),
				resource.TestCheckResourceAttr(resourceName, "source.0.ssl_mode", "VERIFY_CA"),
				resource.TestCheckResourceAttr(resourceName, "source.0.ssl_ca_certificate.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.username", "username2"),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.applier_username", "adminUser"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.channel_name"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "target.0.delay_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "target.0.filters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.filters.0.type", "REPLICATE_IGNORE_DB"),
				resource.TestCheckResourceAttr(resourceName, "target.0.filters.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "target.0.tables_without_primary_key_handling", "ALLOW"),
				resource.TestCheckResourceAttr(resourceName, "target.0.target_type", "DBSYSTEM"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_channels", "test_channels", acctest.Optional, acctest.Update, MysqlMysqlChannelDataSourceRepresentation) +
				compartmentIdVariableStr + ChannelUpdateResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "channel_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "is_enabled", "false"),

				resource.TestCheckResourceAttr(datasourceName, "channels.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "channels.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "channels.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "channels.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "channels.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "channels.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(datasourceName, "channels.0.source.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "channels.0.source.0.anonymous_transactions_handling.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "channels.0.source.0.anonymous_transactions_handling.0.policy", "ERROR_ON_ANONYMOUS"),
				resource.TestCheckResourceAttr(datasourceName, "channels.0.source.0.hostname", "hostname2.my.company.com"),
				resource.TestCheckResourceAttr(datasourceName, "channels.0.source.0.port", "3306"),
				resource.TestCheckResourceAttr(datasourceName, "channels.0.source.0.source_type", "MYSQL"),
				resource.TestCheckResourceAttr(datasourceName, "channels.0.source.0.ssl_ca_certificate.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "channels.0.source.0.ssl_ca_certificate.0.certificate_type", "PEM"),
				resource.TestCheckResourceAttr(datasourceName, "channels.0.source.0.ssl_mode", "VERIFY_CA"),
				resource.TestCheckResourceAttr(datasourceName, "channels.0.source.0.username", "username2"),
				resource.TestCheckResourceAttrSet(datasourceName, "channels.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "channels.0.target.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "channels.0.target.0.applier_username", "adminUser"),
				resource.TestCheckResourceAttrSet(datasourceName, "channels.0.target.0.channel_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "channels.0.target.0.db_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "channels.0.target.0.delay_in_seconds", "11"),
				resource.TestCheckResourceAttr(datasourceName, "channels.0.target.0.filters.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "channels.0.target.0.filters.0.type", "REPLICATE_IGNORE_DB"),
				resource.TestCheckResourceAttr(datasourceName, "channels.0.target.0.filters.0.value", "value2"),
				resource.TestCheckResourceAttr(datasourceName, "channels.0.target.0.tables_without_primary_key_handling", "ALLOW"),
				resource.TestCheckResourceAttr(datasourceName, "channels.0.target.0.target_type", "DBSYSTEM"),
				resource.TestCheckResourceAttrSet(datasourceName, "channels.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "channels.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_channel", "test_channel", acctest.Required, acctest.Create, MysqlMysqlChannelSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ChannelUpdateResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "channel_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source.0.anonymous_transactions_handling.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source.0.anonymous_transactions_handling.0.policy", "ERROR_ON_ANONYMOUS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source.0.hostname", "hostname2.my.company.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source.0.port", "3306"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source.0.source_type", "MYSQL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source.0.ssl_ca_certificate.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source.0.ssl_ca_certificate.0.certificate_type", "PEM"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source.0.ssl_mode", "VERIFY_CA"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source.0.username", "username2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target.0.applier_username", "adminUser"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target.0.delay_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target.0.filters.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target.0.filters.0.type", "REPLICATE_IGNORE_DB"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target.0.filters.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target.0.tables_without_primary_key_handling", "ALLOW"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target.0.target_type", "DBSYSTEM"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + MysqlChannelRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"lifecycle_details",
				"source.0.password",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckMysqlChannelDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ChannelsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_mysql_channel" {
			noResourceFound = false
			request := oci_mysql.GetChannelRequest{}

			tmp := rs.Primary.ID
			request.ChannelId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "mysql")

			response, err := client.GetChannel(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_mysql.ChannelLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("MysqlChannel") {
		resource.AddTestSweepers("MysqlChannel", &resource.Sweeper{
			Name:         "MysqlChannel",
			Dependencies: acctest.DependencyGraph["channel"],
			F:            sweepMysqlChannelResource,
		})
	}
}

func sweepMysqlChannelResource(compartment string) error {
	channelsClient := acctest.GetTestClients(&schema.ResourceData{}).ChannelsClient()
	channelIds, err := getMysqlChannelIds(compartment)
	if err != nil {
		return err
	}
	for _, channelId := range channelIds {
		if ok := acctest.SweeperDefaultResourceId[channelId]; !ok {
			deleteChannelRequest := oci_mysql.DeleteChannelRequest{}

			deleteChannelRequest.ChannelId = &channelId

			deleteChannelRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "mysql")
			_, error := channelsClient.DeleteChannel(context.Background(), deleteChannelRequest)
			if error != nil {
				fmt.Printf("Error deleting Channel %s %s, It is possible that the resource is already deleted. Please verify manually \n", channelId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &channelId, MysqlChannelSweepWaitCondition, time.Duration(3*time.Minute),
				MysqlChannelSweepResponseFetchOperation, "mysql", true)
		}
	}
	return nil
}

func getMysqlChannelIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ChannelId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	channelsClient := acctest.GetTestClients(&schema.ResourceData{}).ChannelsClient()

	listChannelsRequest := oci_mysql.ListChannelsRequest{}
	listChannelsRequest.CompartmentId = &compartmentId
	listChannelsRequest.LifecycleState = oci_mysql.ChannelLifecycleStateActive
	listChannelsResponse, err := channelsClient.ListChannels(context.Background(), listChannelsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Channel list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, channel := range listChannelsResponse.Items {
		id := *channel.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ChannelId", id)
	}
	return resourceIds, nil
}

func MysqlChannelSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if channelResponse, ok := response.Response.(oci_mysql.GetChannelResponse); ok {
		return channelResponse.LifecycleState != oci_mysql.ChannelLifecycleStateDeleted
	}
	return false
}

func MysqlChannelSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ChannelsClient().GetChannel(context.Background(), oci_mysql.GetChannelRequest{
		ChannelId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
