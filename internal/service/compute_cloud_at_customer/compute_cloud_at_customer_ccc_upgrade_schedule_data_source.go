// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package compute_cloud_at_customer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_compute_cloud_at_customer "github.com/oracle/oci-go-sdk/v65/computecloudatcustomer"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ComputeCloudAtCustomerCccUpgradeScheduleDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["ccc_upgrade_schedule_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ComputeCloudAtCustomerCccUpgradeScheduleResource(), fieldMap, readSingularComputeCloudAtCustomerCccUpgradeSchedule)
}

func readSingularComputeCloudAtCustomerCccUpgradeSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &ComputeCloudAtCustomerCccUpgradeScheduleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeCloudAtCustomerClient()

	return tfresource.ReadResource(sync)
}

type ComputeCloudAtCustomerCccUpgradeScheduleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_compute_cloud_at_customer.ComputeCloudAtCustomerClient
	Res    *oci_compute_cloud_at_customer.GetCccUpgradeScheduleResponse
}

func (s *ComputeCloudAtCustomerCccUpgradeScheduleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ComputeCloudAtCustomerCccUpgradeScheduleDataSourceCrud) Get() error {
	request := oci_compute_cloud_at_customer.GetCccUpgradeScheduleRequest{}

	if cccUpgradeScheduleId, ok := s.D.GetOkExists("ccc_upgrade_schedule_id"); ok {
		tmp := cccUpgradeScheduleId.(string)
		request.CccUpgradeScheduleId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "compute_cloud_at_customer")

	response, err := s.Client.GetCccUpgradeSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ComputeCloudAtCustomerCccUpgradeScheduleDataSourceCrud) SetData() error {
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

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	events := []interface{}{}
	for _, item := range s.Res.Events {
		events = append(events, CccScheduleEventToMap(item))
	}
	s.D.Set("events", events)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("infrastructure_ids", s.Res.InfrastructureIds)

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
