// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package em_warehouse

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_em_warehouse "github.com/oracle/oci-go-sdk/v65/emwarehouse"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func EmWarehouseEmWarehouseDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["em_warehouse_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(EmWarehouseEmWarehouseResource(), fieldMap, readSingularEmWarehouseEmWarehouse)
}

func readSingularEmWarehouseEmWarehouse(d *schema.ResourceData, m interface{}) error {
	sync := &EmWarehouseEmWarehouseDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmWarehouseClient()

	return tfresource.ReadResource(sync)
}

type EmWarehouseEmWarehouseDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_em_warehouse.EmWarehouseClient
	Res    *oci_em_warehouse.GetEmWarehouseResponse
}

func (s *EmWarehouseEmWarehouseDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *EmWarehouseEmWarehouseDataSourceCrud) Get() error {
	request := oci_em_warehouse.GetEmWarehouseRequest{}

	if emWarehouseId, ok := s.D.GetOkExists("em_warehouse_id"); ok {
		tmp := emWarehouseId.(string)
		request.EmWarehouseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "emwarehouse")

	response, err := s.Client.GetEmWarehouse(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *EmWarehouseEmWarehouseDataSourceCrud) SetData() error {
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

	if s.Res.EmBridgeId != nil {
		s.D.Set("em_bridge_id", *s.Res.EmBridgeId)
	}

	if s.Res.EmWarehouseType != nil {
		s.D.Set("em_warehouse_type", *s.Res.EmWarehouseType)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LatestEtlRunMessage != nil {
		s.D.Set("latest_etl_run_message", *s.Res.LatestEtlRunMessage)
	}

	if s.Res.LatestEtlRunStatus != nil {
		s.D.Set("latest_etl_run_status", *s.Res.LatestEtlRunStatus)
	}

	if s.Res.LatestEtlRunTime != nil {
		s.D.Set("latest_etl_run_time", *s.Res.LatestEtlRunTime)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.OperationsInsightsWarehouseId != nil {
		s.D.Set("operations_insights_warehouse_id", *s.Res.OperationsInsightsWarehouseId)
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
