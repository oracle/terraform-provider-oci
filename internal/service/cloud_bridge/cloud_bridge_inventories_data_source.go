// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_bridge

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_bridge "github.com/oracle/oci-go-sdk/v65/cloudbridge"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudBridgeInventoriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCloudBridgeInventories,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"inventory_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(CloudBridgeInventoryResource()),
						},
					},
				},
			},
		},
	}
}

func readCloudBridgeInventories(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeInventoriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InventoryClient()

	return tfresource.ReadResource(sync)
}

type CloudBridgeInventoriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_bridge.InventoryClient
	Res    *oci_cloud_bridge.ListInventoriesResponse
}

func (s *CloudBridgeInventoriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudBridgeInventoriesDataSourceCrud) Get() error {
	request := oci_cloud_bridge.ListInventoriesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_cloud_bridge.InventoryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_bridge")

	response, err := s.Client.ListInventories(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInventories(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CloudBridgeInventoriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CloudBridgeInventoriesDataSource-", CloudBridgeInventoriesDataSource(), s.D))
	resources := []map[string]interface{}{}
	inventory := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, InventorySummaryToMap(item))
	}
	inventory["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CloudBridgeInventoriesDataSource().Schema["inventory_collection"].Elem.(*schema.Resource).Schema)
		inventory["items"] = items
	}

	resources = append(resources, inventory)
	if err := s.D.Set("inventory_collection", resources); err != nil {
		return err
	}

	return nil
}
