// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"time"

	"github.com/hashicorp/terraform/helper/schema"

	oci_audit "github.com/oracle/oci-go-sdk/audit"
)

func ConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: DefaultTimeout,
		Create:   createConfiguration,
		Read:     readConfiguration,
		Update:   updateConfiguration,
		Delete:   deleteConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"retention_period_days": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &ConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).auditClient

	return CreateResource(d, sync)
}

func readConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &ConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).auditClient

	return ReadResource(sync)
}

func updateConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &ConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).auditClient

	return UpdateResource(d, sync)
}

func deleteConfiguration(d *schema.ResourceData, m interface{}) error {
	return nil
}

type ConfigurationResourceCrud struct {
	BaseCrud
	Client                 *oci_audit.AuditClient
	Res                    *oci_audit.Configuration
	DisableNotFoundRetries bool
}

func (s *ConfigurationResourceCrud) ID() string {
	return s.D.Get("compartment_id").(string)
}

func (s *ConfigurationResourceCrud) Create() error {
	// This resource can't actually be created. So treat it as an update instead.
	return s.Update()
}

func (s *ConfigurationResourceCrud) Get() error {
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

func (s *ConfigurationResourceCrud) Update() error {
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

func (s *ConfigurationResourceCrud) SetData() error {
	if s.Res.RetentionPeriodDays != nil {
		s.D.Set("retention_period_days", *s.Res.RetentionPeriodDays)
	}

	return nil
}
