// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LogAnalyticsNamespaceDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularLogAnalyticsNamespaceWithContext,
		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_archiving_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_onboarded": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_logset_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_data_ever_ingested": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func readSingularLogAnalyticsNamespaceWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &LogAnalyticsNamespaceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type LogAnalyticsNamespaceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.GetNamespaceResponse
}

func (s *LogAnalyticsNamespaceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsNamespaceDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_log_analytics.GetNamespaceRequest{}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.GetNamespace(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LogAnalyticsNamespaceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LogAnalyticsNamespaceDataSource-", LogAnalyticsNamespaceDataSource(), s.D))

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.IsArchivingEnabled != nil {
		s.D.Set("is_archiving_enabled", *s.Res.IsArchivingEnabled)
	}

	if s.Res.IsOnboarded != nil {
		s.D.Set("is_onboarded", *s.Res.IsOnboarded)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.IsLogSetEnabled != nil {
		s.D.Set("is_logset_enabled", *s.Res.IsLogSetEnabled)
	}

	if s.Res.IsDataEverIngested != nil {
		s.D.Set("is_data_ever_ingested", *s.Res.IsDataEverIngested)
	}

	return nil
}
