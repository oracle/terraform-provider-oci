// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataflow

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dataflow "github.com/oracle/oci-go-sdk/v65/dataflow"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataflowPoolDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["pool_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataflowPoolResource(), fieldMap, readSingularDataflowPool)
}

func readSingularDataflowPool(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowPoolDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()

	return tfresource.ReadResource(sync)
}

type DataflowPoolDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dataflow.DataFlowClient
	Res    *oci_dataflow.GetPoolResponse
}

func (s *DataflowPoolDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataflowPoolDataSourceCrud) Get() error {
	request := oci_dataflow.GetPoolRequest{}

	if poolId, ok := s.D.GetOkExists("pool_id"); ok {
		tmp := poolId.(string)
		request.PoolId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dataflow")

	response, err := s.Client.GetPool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataflowPoolDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	configurations := []interface{}{}
	for _, item := range s.Res.Configurations {
		configurations = append(configurations, PoolConfigToMap(item))
	}
	s.D.Set("configurations", configurations)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IdleTimeoutInMinutes != nil {
		s.D.Set("idle_timeout_in_minutes", *s.Res.IdleTimeoutInMinutes)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.OwnerPrincipalId != nil {
		s.D.Set("owner_principal_id", *s.Res.OwnerPrincipalId)
	}

	if s.Res.OwnerUserName != nil {
		s.D.Set("owner_user_name", *s.Res.OwnerUserName)
	}

	if s.Res.PoolMetrics != nil {
		s.D.Set("pool_metrics", []interface{}{PoolMetricsToMap(s.Res.PoolMetrics)})
	} else {
		s.D.Set("pool_metrics", nil)
	}

	schedules := []interface{}{}
	for _, item := range s.Res.Schedules {
		schedules = append(schedules, PoolScheduleToMap(item))
	}
	s.D.Set("schedules", schedules)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
