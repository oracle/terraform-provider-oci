// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_audit "github.com/oracle/oci-go-sdk/audit"
)

func AuditConfigurationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularAuditConfiguration,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"retention_period_days": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readSingularAuditConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &AuditConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).auditClient

	return ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "audit")

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

	s.D.SetId(GenerateDataSourceID())

	if s.Res.RetentionPeriodDays != nil {
		s.D.Set("retention_period_days", *s.Res.RetentionPeriodDays)
	}

	return nil
}
