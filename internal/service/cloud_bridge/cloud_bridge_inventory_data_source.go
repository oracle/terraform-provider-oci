// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_bridge

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_bridge "github.com/oracle/oci-go-sdk/v65/cloudbridge"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudBridgeInventoryDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["inventory_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CloudBridgeInventoryResource(), fieldMap, readSingularCloudBridgeInventory)
}

func readSingularCloudBridgeInventory(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeInventoryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InventoryClient()

	return tfresource.ReadResource(sync)
}

type CloudBridgeInventoryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_bridge.InventoryClient
	Res    *oci_cloud_bridge.GetInventoryResponse
}

func (s *CloudBridgeInventoryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudBridgeInventoryDataSourceCrud) Get() error {
	request := oci_cloud_bridge.GetInventoryRequest{}

	if inventoryId, ok := s.D.GetOkExists("inventory_id"); ok {
		tmp := inventoryId.(string)
		request.InventoryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_bridge")

	response, err := s.Client.GetInventory(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CloudBridgeInventoryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
