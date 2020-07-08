// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_waas "github.com/oracle/oci-go-sdk/waas"
)

func init() {
	RegisterResource("oci_waas_custom_protection_rule", WaasCustomProtectionRuleResource())
}

func WaasCustomProtectionRuleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createWaasCustomProtectionRule,
		Read:     readWaasCustomProtectionRule,
		Update:   updateWaasCustomProtectionRule,
		Delete:   deleteWaasCustomProtectionRule,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"template": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"mod_security_rule_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
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

func createWaasCustomProtectionRule(d *schema.ResourceData, m interface{}) error {
	sync := &WaasCustomProtectionRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).waasClient()

	return CreateResource(d, sync)
}

func readWaasCustomProtectionRule(d *schema.ResourceData, m interface{}) error {
	sync := &WaasCustomProtectionRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).waasClient()

	return ReadResource(sync)
}

func updateWaasCustomProtectionRule(d *schema.ResourceData, m interface{}) error {
	sync := &WaasCustomProtectionRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).waasClient()

	return UpdateResource(d, sync)
}

func deleteWaasCustomProtectionRule(d *schema.ResourceData, m interface{}) error {
	sync := &WaasCustomProtectionRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).waasClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type WaasCustomProtectionRuleResourceCrud struct {
	BaseCrud
	Client                 *oci_waas.WaasClient
	Res                    *oci_waas.CustomProtectionRule
	DisableNotFoundRetries bool
}

func (s *WaasCustomProtectionRuleResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *WaasCustomProtectionRuleResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_waas.LifecycleStatesCreating),
	}
}

func (s *WaasCustomProtectionRuleResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_waas.LifecycleStatesActive),
	}
}

func (s *WaasCustomProtectionRuleResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_waas.LifecycleStatesDeleting),
	}
}

func (s *WaasCustomProtectionRuleResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_waas.LifecycleStatesDeleted),
	}
}

func (s *WaasCustomProtectionRuleResourceCrud) Create() error {
	request := oci_waas.CreateCustomProtectionRuleRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if template, ok := s.D.GetOkExists("template"); ok {
		tmp := template.(string)
		request.Template = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "waas")

	response, err := s.Client.CreateCustomProtectionRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CustomProtectionRule
	return nil
}

func (s *WaasCustomProtectionRuleResourceCrud) Get() error {
	request := oci_waas.GetCustomProtectionRuleRequest{}

	tmp := s.D.Id()
	request.CustomProtectionRuleId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "waas")

	response, err := s.Client.GetCustomProtectionRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CustomProtectionRule
	return nil
}

func (s *WaasCustomProtectionRuleResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_waas.UpdateCustomProtectionRuleRequest{}

	tmp := s.D.Id()
	request.CustomProtectionRuleId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if template, ok := s.D.GetOkExists("template"); ok {
		tmp := template.(string)
		request.Template = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "waas")

	response, err := s.Client.UpdateCustomProtectionRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CustomProtectionRule
	return nil
}

func (s *WaasCustomProtectionRuleResourceCrud) Delete() error {
	request := oci_waas.DeleteCustomProtectionRuleRequest{}

	tmp := s.D.Id()
	request.CustomProtectionRuleId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "waas")

	_, err := s.Client.DeleteCustomProtectionRule(context.Background(), request)
	return err
}

func (s *WaasCustomProtectionRuleResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("mod_security_rule_ids", s.Res.ModSecurityRuleIds)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Template != nil {
		s.D.Set("template", *s.Res.Template)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *WaasCustomProtectionRuleResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_waas.ChangeCustomProtectionRuleCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.CustomProtectionRuleId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "waas")

	_, err := s.Client.ChangeCustomProtectionRuleCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
