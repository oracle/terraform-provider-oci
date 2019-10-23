// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"time"

	"github.com/hashicorp/terraform/helper/schema"

	oci_audit "github.com/oracle/oci-go-sdk/audit"
)

func AuditConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: DefaultTimeout,
		Create:   createAuditConfiguration,
		Read:     readAuditConfiguration,
		Update:   updateAuditConfiguration,
		Delete:   deleteAuditConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"retention_period_days": {
				Type:     schema.TypeInt,
				Required: true,
			},

			// Optional

			// Computed
		},
	}
}

func createAuditConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &AuditConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).auditClient

	return CreateResource(d, sync)
}

func readAuditConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &AuditConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).auditClient

	return ReadResource(sync)
}

func updateAuditConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &AuditConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).auditClient

	return UpdateResource(d, sync)
}

func deleteAuditConfiguration(d *schema.ResourceData, m interface{}) error {
	return nil
}

type AuditConfigurationResourceCrud struct {
	BaseCrud
	Client                 *oci_audit.AuditClient
	Res                    *oci_audit.Configuration
	DisableNotFoundRetries bool
}

func (s *AuditConfigurationResourceCrud) ID() string {
	return s.D.Get("compartment_id").(string)
}

func (s *AuditConfigurationResourceCrud) Create() error {
	// This resource can't actually be created. So treat it as an update instead.
	return s.Update()
}

func (s *AuditConfigurationResourceCrud) Get() error {
	request := oci_audit.GetConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "audit")

	response, err := s.Client.GetConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Configuration
	return nil
}

func (s *AuditConfigurationResourceCrud) Update() error {
	request := oci_audit.UpdateConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if retentionPeriodDays, ok := s.D.GetOkExists("retention_period_days"); ok {
		tmp := retentionPeriodDays.(int)
		request.RetentionPeriodDays = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "audit")

	_, err := s.Client.UpdateConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	// Workaround: Sleep for some time before polling the configuration. Because update happens asynchronously, polling too
	// soon may result in service returning stale configuration values.
	time.Sleep(time.Second * 5)

	// Requests to update the retention policy may succeed instantly but may not see the actual update take effect
	// until minutes later. Add polling here to return only when the change has taken effect.
	retentionPolicyFunc := func() bool { return *s.Res.RetentionPeriodDays == *request.RetentionPeriodDays }
	return WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *AuditConfigurationResourceCrud) SetData() error {
	if s.Res.RetentionPeriodDays != nil {
		s.D.Set("retention_period_days", *s.Res.RetentionPeriodDays)
	}

	return nil
}
