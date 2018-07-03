// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_audit "github.com/oracle/oci-go-sdk/audit"

	"github.com/oracle/terraform-provider-oci/crud"
)

func ConfigurationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularConfiguration,
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

func readSingularConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &ConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).auditClient

	return crud.ReadResource(sync)
}

type ConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_audit.AuditClient
	Res    *oci_audit.GetConfigurationResponse
}

func (s *ConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ConfigurationDataSourceCrud) Get() error {
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

func (s *ConfigurationDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())

	if s.Res.RetentionPeriodDays != nil {
		s.D.Set("retention_period_days", *s.Res.RetentionPeriodDays)
	}

	return
}
