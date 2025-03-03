// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"
)

func OpsiOperationsInsightsWarehouseDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["operations_insights_warehouse_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OpsiOperationsInsightsWarehouseResource(), fieldMap, readSingularOpsiOperationsInsightsWarehouse)
}

func readSingularOpsiOperationsInsightsWarehouse(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiOperationsInsightsWarehouseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

type OpsiOperationsInsightsWarehouseDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.GetOperationsInsightsWarehouseResponse
}

func (s *OpsiOperationsInsightsWarehouseDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiOperationsInsightsWarehouseDataSourceCrud) Get() error {
	request := oci_opsi.GetOperationsInsightsWarehouseRequest{}

	if operationsInsightsWarehouseId, ok := s.D.GetOkExists("operations_insights_warehouse_id"); ok {
		tmp := operationsInsightsWarehouseId.(string)
		request.OperationsInsightsWarehouseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opsi")

	response, err := s.Client.GetOperationsInsightsWarehouse(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OpsiOperationsInsightsWarehouseDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("compute_model", s.Res.ComputeModel)

	if s.Res.CpuAllocated != nil {
		s.D.Set("cpu_allocated", *s.Res.CpuAllocated)
	}

	if s.Res.CpuUsed != nil {
		s.D.Set("cpu_used", *s.Res.CpuUsed)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DynamicGroupId != nil {
		s.D.Set("dynamic_group_id", *s.Res.DynamicGroupId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.OperationsInsightsTenancyId != nil {
		s.D.Set("operations_insights_tenancy_id", *s.Res.OperationsInsightsTenancyId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StorageAllocatedInGBs != nil {
		s.D.Set("storage_allocated_in_gbs", *s.Res.StorageAllocatedInGBs)
	}

	if s.Res.StorageUsedInGBs != nil {
		s.D.Set("storage_used_in_gbs", *s.Res.StorageUsedInGBs)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastWalletRotated != nil {
		s.D.Set("time_last_wallet_rotated", s.Res.TimeLastWalletRotated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
