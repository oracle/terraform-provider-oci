// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package audit

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"

	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_audit "github.com/oracle/oci-go-sdk/v56/audit"
)

func AuditConfigurationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["compartment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(AuditConfigurationResource(), fieldMap, readSingularAuditConfiguration)
}

func readSingularAuditConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &AuditConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AuditClient()

	return tfresource.ReadResource(sync)
}

type AuditConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_audit.AuditClient
	Res    *oci_audit.GetConfigurationResponse
}

func (s *AuditConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AuditConfigurationDataSourceCrud) Get() error {
	request := oci_audit.GetConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "audit")

	response, err := s.Client.GetConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AuditConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("AuditConfigurationDataSource-", AuditConfigurationDataSource(), s.D))

	if s.Res.RetentionPeriodDays != nil {
		s.D.Set("retention_period_days", *s.Res.RetentionPeriodDays)
	}

	return nil
}
