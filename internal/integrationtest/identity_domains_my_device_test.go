// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsIdentityDomainsMyDeviceSingularDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_device_id":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domains_my_devices.test_my_devices.my_devices.0.id}`},
		"attribute_sets": acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
	}

	IdentityDomainsIdentityDomainsMyDeviceDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_device_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"my_device_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"attribute_sets":   acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":      acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsMyDeviceResourceConfig = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsMyDeviceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsMyDeviceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_domains_my_devices.test_my_devices"
	singularDatasourceName := "data.oci_identity_domains_my_device.test_my_device"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_devices", "test_my_devices", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsMyDeviceDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMyDeviceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestCheckResourceAttrSet(datasourceName, "my_devices.#"),
				resource.TestCheckResourceAttr(datasourceName, "my_devices.0.schemas.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_devices", "test_my_devices", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsMyDeviceDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_device", "test_my_device", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsMyDeviceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMyDeviceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "idcs_endpoint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "my_device_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "additional_attributes.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "authentication_factors.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_acc_rec_enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "last_sync_time"),
				resource.TestCheckResourceAttr(singularDatasourceName, "non_compliances.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ocid"),
				resource.TestCheckResourceAttr(singularDatasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttr(singularDatasourceName, "third_party_factor.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user.#", "1"),
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("IdentityDomainsMyDevice") {
		resource.AddTestSweepers("IdentityDomainsMyDevice", &resource.Sweeper{
			Name:         "IdentityDomainsMyDevice",
			Dependencies: acctest.DependencyGraph["myDevice"],
			F:            sweepIdentityDomainsMyDeviceResource,
		})
	}
}

func sweepIdentityDomainsMyDeviceResource(compartment string) error {
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()
	myDeviceIds, err := getIdentityDomainsMyDeviceIds(compartment)
	if err != nil {
		return err
	}
	for _, myDeviceId := range myDeviceIds {
		if ok := acctest.SweeperDefaultResourceId[myDeviceId]; !ok {
			deleteMyDeviceRequest := oci_identity_domains.DeleteMyDeviceRequest{}

			deleteMyDeviceRequest.MyDeviceId = &myDeviceId

			deleteMyDeviceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "identity_domains")
			_, error := identityDomainsClient.DeleteMyDevice(context.Background(), deleteMyDeviceRequest)
			if error != nil {
				fmt.Printf("Error deleting MyDevice %s %s, It is possible that the resource is already deleted. Please verify manually \n", myDeviceId, error)
				continue
			}
		}
	}
	return nil
}

func getIdentityDomainsMyDeviceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MyDeviceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	identityDomainsClient := acctest.GetTestClients(&schema.ResourceData{}).IdentityDomainsClient()

	listMyDevicesRequest := oci_identity_domains.ListMyDevicesRequest{}
	listMyDevicesResponse, err := identityDomainsClient.ListMyDevices(context.Background(), listMyDevicesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MyDevice list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, myDevice := range listMyDevicesResponse.Resources {
		id := *myDevice.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MyDeviceId", id)
	}
	return resourceIds, nil
}
