// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package limits

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_limits "github.com/oracle/oci-go-sdk/v56/limits"

	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func LimitsQuotaResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLimitsQuota,
		Read:     readLimitsQuota,
		Update:   updateLimitsQuota,
		Delete:   deleteLimitsQuota,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"statements": {
				Type:             schema.TypeList,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLimitsQuota(d *schema.ResourceData, m interface{}) error {
	sync := &LimitsQuotaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).QuotasClient()

	return tfresource.CreateResource(d, sync)
}

func readLimitsQuota(d *schema.ResourceData, m interface{}) error {
	sync := &LimitsQuotaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).QuotasClient()

	return tfresource.ReadResource(sync)
}

func updateLimitsQuota(d *schema.ResourceData, m interface{}) error {
	sync := &LimitsQuotaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).QuotasClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteLimitsQuota(d *schema.ResourceData, m interface{}) error {
	sync := &LimitsQuotaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).QuotasClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LimitsQuotaResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_limits.QuotasClient
	Res                    *oci_limits.Quota
	DisableNotFoundRetries bool
}

func (s *LimitsQuotaResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *LimitsQuotaResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *LimitsQuotaResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_limits.QuotaLifecycleStateActive),
	}
}

func (s *LimitsQuotaResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *LimitsQuotaResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *LimitsQuotaResourceCrud) Create() error {
	request := oci_limits.CreateQuotaRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if statements, ok := s.D.GetOkExists("statements"); ok {
		interfaces := statements.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("statements") {
			request.Statements = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "limits")

	response, err := s.Client.CreateQuota(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Quota
	return nil
}

func (s *LimitsQuotaResourceCrud) Get() error {
	request := oci_limits.GetQuotaRequest{}

	tmp := s.D.Id()
	request.QuotaId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "limits")

	response, err := s.Client.GetQuota(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Quota
	return nil
}

func (s *LimitsQuotaResourceCrud) Update() error {
	request := oci_limits.UpdateQuotaRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.QuotaId = &tmp

	if statements, ok := s.D.GetOkExists("statements"); ok {
		interfaces := statements.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("statements") {
			request.Statements = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "limits")

	response, err := s.Client.UpdateQuota(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Quota
	return nil
}

func (s *LimitsQuotaResourceCrud) Delete() error {
	request := oci_limits.DeleteQuotaRequest{}

	tmp := s.D.Id()
	request.QuotaId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "limits")

	_, err := s.Client.DeleteQuota(context.Background(), request)
	return err
}

func (s *LimitsQuotaResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("statements", s.Res.Statements)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
