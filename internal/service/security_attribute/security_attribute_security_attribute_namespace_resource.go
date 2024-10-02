// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package security_attribute

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_security_attribute "github.com/oracle/oci-go-sdk/v65/securityattribute"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func SecurityAttributeSecurityAttributeNamespaceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createSecurityAttributeSecurityAttributeNamespace,
		Read:     readSecurityAttributeSecurityAttributeNamespace,
		Update:   updateSecurityAttributeSecurityAttributeNamespace,
		Delete:   deleteSecurityAttributeSecurityAttributeNamespace,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
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
			"is_retired": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
			"mode": {
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
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createSecurityAttributeSecurityAttributeNamespace(d *schema.ResourceData, m interface{}) error {
	sync := &SecurityAttributeSecurityAttributeNamespaceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SecurityAttributeClient()

	return tfresource.CreateResource(d, sync)
}

func readSecurityAttributeSecurityAttributeNamespace(d *schema.ResourceData, m interface{}) error {
	sync := &SecurityAttributeSecurityAttributeNamespaceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SecurityAttributeClient()

	return tfresource.ReadResource(sync)
}

func updateSecurityAttributeSecurityAttributeNamespace(d *schema.ResourceData, m interface{}) error {
	sync := &SecurityAttributeSecurityAttributeNamespaceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SecurityAttributeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteSecurityAttributeSecurityAttributeNamespace(d *schema.ResourceData, m interface{}) error {
	sync := &SecurityAttributeSecurityAttributeNamespaceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SecurityAttributeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type SecurityAttributeSecurityAttributeNamespaceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_security_attribute.SecurityAttributeClient
	Res                    *oci_security_attribute.SecurityAttributeNamespace
	DisableNotFoundRetries bool
}

func (s *SecurityAttributeSecurityAttributeNamespaceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *SecurityAttributeSecurityAttributeNamespaceResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *SecurityAttributeSecurityAttributeNamespaceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_security_attribute.SecurityAttributeNamespaceLifecycleStateActive),
	}
}

func (s *SecurityAttributeSecurityAttributeNamespaceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_security_attribute.SecurityAttributeNamespaceLifecycleStateDeleting),
	}
}

func (s *SecurityAttributeSecurityAttributeNamespaceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_security_attribute.SecurityAttributeNamespaceLifecycleStateDeleted),
	}
}

func (s *SecurityAttributeSecurityAttributeNamespaceResourceCrud) Create() error {
	request := oci_security_attribute.CreateSecurityAttributeNamespaceRequest{}

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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "security_attribute")

	response, err := s.Client.CreateSecurityAttributeNamespace(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SecurityAttributeNamespace
	return nil
}

func (s *SecurityAttributeSecurityAttributeNamespaceResourceCrud) Get() error {
	request := oci_security_attribute.GetSecurityAttributeNamespaceRequest{}

	tmp := s.D.Id()
	request.SecurityAttributeNamespaceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "security_attribute")

	response, err := s.Client.GetSecurityAttributeNamespace(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SecurityAttributeNamespace
	return nil
}

func (s *SecurityAttributeSecurityAttributeNamespaceResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_security_attribute.UpdateSecurityAttributeNamespaceRequest{}

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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isRetired, ok := s.D.GetOkExists("is_retired"); ok {
		tmp := isRetired.(bool)
		request.IsRetired = &tmp
	}

	tmp := s.D.Id()
	request.SecurityAttributeNamespaceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "security_attribute")

	response, err := s.Client.UpdateSecurityAttributeNamespace(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SecurityAttributeNamespace
	return nil
}

func (s *SecurityAttributeSecurityAttributeNamespaceResourceCrud) Delete() error {
	request := oci_security_attribute.DeleteSecurityAttributeNamespaceRequest{}

	tmp := s.D.Id()
	request.SecurityAttributeNamespaceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "security_attribute")

	_, err := s.Client.DeleteSecurityAttributeNamespace(context.Background(), request)
	return err
}

func (s *SecurityAttributeSecurityAttributeNamespaceResourceCrud) SetData() error {
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

	if s.Res.IsRetired != nil {
		s.D.Set("is_retired", *s.Res.IsRetired)
	}

	s.D.Set("mode", s.Res.Mode)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *SecurityAttributeSecurityAttributeNamespaceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_security_attribute.ChangeSecurityAttributeNamespaceCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SecurityAttributeNamespaceId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "security_attribute")

	_, err := s.Client.ChangeSecurityAttributeNamespaceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
