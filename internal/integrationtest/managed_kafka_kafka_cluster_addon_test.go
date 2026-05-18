// Copyright (c) 2017, 2024, Oracle and/or its affiliates.

package integrationtest

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_managed_kafka "github.com/oracle/oci-go-sdk/v65/managedkafka"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

const (
	managedKafkaKafkaClusterAddonName = "terraform-public-addon-test-mtls"
)

var (
	ManagedKafkaKafkaClusterAddonRepresentation = map[string]interface{}{
		"addon_type": acctest.Representation{
			RepType: acctest.Required,
			Create:  `PUBLICCONNECTIVITY`,
		},
		"authentication_mechanism": acctest.Representation{
			RepType: acctest.Required,
			Create:  `MTLS`,
		},
		"kafka_cluster_id": acctest.Representation{
			RepType: acctest.Required,
			Create:  `${var.kafka_cluster_id}`,
			Update:  `${var.kafka_cluster_id}`,
		},
		"name": acctest.Representation{
			RepType: acctest.Required,
			Create:  managedKafkaKafkaClusterAddonName,
		},
		"network_cidrs": acctest.Representation{
			RepType: acctest.Required,
			Create:  []string{`10.0.0.0/24`},
			Update:  []string{`10.0.1.0/24`},
		},
		"description": acctest.Representation{
			RepType: acctest.Required,
			Create:  `description`,
			Update:  `description2`,
		},
	}
	ManagedKafkaKafkaClusterAddonResourceDependencies = ""
)

func TestManagedKafkaKafkaClusterAddonResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagedKafkaKafkaClusterAddonResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	clusterId := utils.GetEnvSettingWithBlankDefault("managed_kafka_cluster_ocid")
	if clusterId == "" {
		t.Fatal("managed_kafka_cluster_ocid must be set")
	}
	clusterIdVariableStr := fmt.Sprintf("variable \"kafka_cluster_id\" { default = \"%s\" }\n", clusterId)

	resourceName := "oci_managed_kafka_kafka_cluster_addon.test_kafka_cluster_addon"

	var resId, resId2 string

	// This removes any pre-existing addon with the same name before the first create step.
	// It also prints every addon found on the cluster.
	acctest.ResourceTest(t, testAccCheckManagedKafkaKafkaClusterAddonDestroy, []resource.TestStep{
		{
			PreConfig: func() {
				if err := cleanupManagedKafkaKafkaClusterAddon(clusterId, managedKafkaKafkaClusterAddonName); err != nil {
					t.Fatalf("failed to delete existing addons: %v", err)
				}
			},
			Config: config + clusterIdVariableStr + ManagedKafkaKafkaClusterAddonResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_managed_kafka_kafka_cluster_addon",
					"test_kafka_cluster_addon",
					acctest.Required,
					acctest.Create,
					ManagedKafkaKafkaClusterAddonRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "addon_type", "PUBLICCONNECTIVITY"),
				resource.TestCheckResourceAttr(resourceName, "authentication_mechanism", "MTLS"),
				resource.TestCheckResourceAttr(resourceName, "name", managedKafkaKafkaClusterAddonName),
				resource.TestCheckResourceAttr(resourceName, "network_cidrs.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "kafka_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		//verify updates to updatable parameters
		{
			Config: config + clusterIdVariableStr + ManagedKafkaKafkaClusterAddonResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_managed_kafka_kafka_cluster_addon",
					"test_kafka_cluster_addon",
					acctest.Optional,
					acctest.Update,
					ManagedKafkaKafkaClusterAddonRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "addon_type", "PUBLICCONNECTIVITY"),
				resource.TestCheckResourceAttr(resourceName, "authentication_mechanism", "MTLS"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "name", managedKafkaKafkaClusterAddonName),
				resource.TestCheckResourceAttr(resourceName, "network_cidrs.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_cidrs.0", "10.0.1.0/24"),
				resource.TestCheckResourceAttrSet(resourceName, "bootstrap_url"),
				resource.TestCheckResourceAttrSet(resourceName, "kafka_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// final delete
		{
			Config: config,
		},
	})
}

func cleanupManagedKafkaKafkaClusterAddon(clusterID, addonName string) error {
	client := acctest.GetTestClients(&schema.ResourceData{}).KafkaClusterClient()

	if err := listAndPrintManagedKafkaKafkaClusterAddons(clusterID); err != nil {
		return err
	}

	req := oci_managed_kafka.UninstallAddonRequest{
		AddonName:      &addonName,
		KafkaClusterId: &clusterID,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: tfresource.GetRetryPolicy(true, "managed_kafka"),
		},
	}

	_, err := client.UninstallAddon(context.Background(), req)
	if err != nil {
		// If it was already deleted, that's fine.
		if failure, ok := common.IsServiceError(err); ok && failure.GetHTTPStatusCode() == 404 {
			return nil
		}
		return err
	}

	return waitForManagedKafkaKafkaClusterAddonGone(clusterID, addonName)
}

func listAndPrintManagedKafkaKafkaClusterAddons(clusterID string) error {
	client := acctest.GetTestClients(&schema.ResourceData{}).KafkaClusterClient()

	req := oci_managed_kafka.ListAddonsRequest{
		KafkaClusterId: &clusterID,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: tfresource.GetRetryPolicy(true, "managed_kafka"),
		},
	}

	resp, err := client.ListAddons(context.Background(), req)
	if err != nil {
		return err
	}

	for _, addon := range resp.Items {
		name := "<nil>"
		if addon.Name != nil {
			name = *addon.Name
		}

		state := string(addon.LifecycleState)
		fmt.Printf("found addon: id=%s name=%s state=%s\n", getKafkaClusterAddonCompositeId(name, clusterID), name, state)
	}

	return nil
}

func getKafkaClusterAddonCompositeId(addonName, clusterID string) string {
	addonName = url.PathEscape(addonName)
	clusterID = url.PathEscape(clusterID)
	return "kafkaClusters/" + clusterID + "/addons/" + addonName
}

func waitForManagedKafkaKafkaClusterAddonGone(clusterID, addonName string) error {
	client := acctest.GetTestClients(&schema.ResourceData{}).KafkaClusterClient()

	deadline := time.Now().Add(15 * time.Minute)

	for time.Now().Before(deadline) {
		req := oci_managed_kafka.GetAddonRequest{
			AddonName:      &addonName,
			KafkaClusterId: &clusterID,
			RequestMetadata: common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(true, "managed_kafka"),
			},
		}

		resp, err := client.GetAddon(context.Background(), req)
		if err != nil {
			if failure, ok := common.IsServiceError(err); ok && failure.GetHTTPStatusCode() == 404 {
				return nil
			}
			return err
		}

		if resp.GetLifecycleState() == oci_managed_kafka.KafkaClusterAddonLifecycleStateDeleted {
			return nil
		}

		time.Sleep(5 * time.Second)
	}

	return fmt.Errorf("addon %s on cluster %s did not disappear in time", addonName, clusterID)
}

func testAccCheckManagedKafkaKafkaClusterAddonDestroy(s *terraform.State) error {
	if len(s.RootModule().Resources) == 0 {
		return nil
	}

	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).KafkaClusterClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "oci_managed_kafka_kafka_cluster_addon" {
			continue
		}

		req := oci_managed_kafka.GetAddonRequest{}

		if v, ok := rs.Primary.Attributes["name"]; ok {
			req.AddonName = &v
		}
		if v, ok := rs.Primary.Attributes["kafka_cluster_id"]; ok {
			req.KafkaClusterId = &v
		}

		req.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "managed_kafka")

		resp, err := client.GetAddon(context.Background(), req)
		if err == nil {
			if resp.GetLifecycleState() != oci_managed_kafka.KafkaClusterAddonLifecycleStateDeleted {
				return fmt.Errorf("addon still exists: %s", resp.GetLifecycleState())
			}
			continue
		}

		if failure, ok := common.IsServiceError(err); ok && failure.GetHTTPStatusCode() == 404 {
			continue
		}

		return err
	}

	return nil
}

func parseManagedKafkaKafkaClusterAddonCompositeId(id string) (string, string, error) {
	parts := strings.Split(id, "/")
	if len(parts) != 4 || parts[0] != "kafkaClusters" || parts[2] != "addons" {
		return "", "", fmt.Errorf("invalid id")
	}

	clusterId, err := url.PathUnescape(parts[1])
	if err != nil {
		return "", "", err
	}

	addonName, err := url.PathUnescape(parts[3])
	if err != nil {
		return "", "", err
	}

	return addonName, clusterId, nil
}
