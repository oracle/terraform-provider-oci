// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package analytics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_analytics "github.com/oracle/oci-go-sdk/v65/analytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AnalyticsAnalyticsInstanceResourceGroupDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["analytics_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["analytics_instance_resource_group_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(AnalyticsAnalyticsInstanceResourceGroupResource(), fieldMap, readSingularAnalyticsAnalyticsInstanceResourceGroupWithContext)
}

func readSingularAnalyticsAnalyticsInstanceResourceGroupWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &AnalyticsAnalyticsInstanceResourceGroupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnalyticsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type AnalyticsAnalyticsInstanceResourceGroupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_analytics.AnalyticsClient
	Res    *oci_analytics.GetResourceGroupResponse
}

func (s *AnalyticsAnalyticsInstanceResourceGroupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AnalyticsAnalyticsInstanceResourceGroupDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_analytics.GetResourceGroupRequest{}

	if analyticsInstanceId, ok := s.D.GetOkExists("analytics_instance_id"); ok {
		tmp := analyticsInstanceId.(string)
		request.AnalyticsInstanceId = &tmp
	}

	if analyticsInstanceResourceGroupId, ok := s.D.GetOkExists("analytics_instance_resource_group_id"); ok {
		tmp := analyticsInstanceResourceGroupId.(string)
		request.AnalyticsInstanceResourceGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "analytics")

	response, err := s.Client.GetResourceGroup(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AnalyticsAnalyticsInstanceResourceGroupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Capacity != nil {
		s.D.Set("capacity", *s.Res.Capacity)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ResourceName != nil {
		s.D.Set("resource_name", *s.Res.ResourceName)
	}

	return nil
}
