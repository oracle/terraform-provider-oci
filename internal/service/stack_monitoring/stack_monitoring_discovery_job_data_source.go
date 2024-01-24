// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringDiscoveryJobDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["discovery_job_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(StackMonitoringDiscoveryJobResource(), fieldMap, readSingularStackMonitoringDiscoveryJob)
}

func readSingularStackMonitoringDiscoveryJob(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringDiscoveryJobDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

type StackMonitoringDiscoveryJobDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_stack_monitoring.StackMonitoringClient
	Res    *oci_stack_monitoring.GetDiscoveryJobResponse
}

func (s *StackMonitoringDiscoveryJobDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StackMonitoringDiscoveryJobDataSourceCrud) Get() error {
	request := oci_stack_monitoring.GetDiscoveryJobRequest{}

	if discoveryJobId, ok := s.D.GetOkExists("discovery_job_id"); ok {
		tmp := discoveryJobId.(string)
		request.DiscoveryJobId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "stack_monitoring")

	response, err := s.Client.GetDiscoveryJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *StackMonitoringDiscoveryJobDataSourceCrud) SetData() error {
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

	if s.Res.DiscoveryClient != nil {
		s.D.Set("discovery_client", *s.Res.DiscoveryClient)
	}

	if s.Res.DiscoveryDetails != nil {
		s.D.Set("discovery_details", []interface{}{DiscoveryDetailsToMap(s.Res.DiscoveryDetails)})
	} else {
		s.D.Set("discovery_details", nil)
	}

	s.D.Set("discovery_type", s.Res.DiscoveryType)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.StatusMessage != nil {
		s.D.Set("status_message", *s.Res.StatusMessage)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TenantId != nil {
		s.D.Set("tenant_id", *s.Res.TenantId)
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.UserId != nil {
		s.D.Set("user_id", *s.Res.UserId)
	}

	return nil
}
