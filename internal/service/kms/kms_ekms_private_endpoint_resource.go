// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_kms "github.com/oracle/oci-go-sdk/v65/keymanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func KmsEkmsPrivateEndpointResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createKmsEkmsPrivateEndpoint,
		Read:     readKmsEkmsPrivateEndpoint,
		Update:   updateKmsEkmsPrivateEndpoint,
		Delete:   deleteKmsEkmsPrivateEndpoint,
		Schema: map[string]*schema.Schema{
			// Required
			"ca_bundle": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"external_key_manager_ip": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_endpoint_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createKmsEkmsPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &KmsEkmsPrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EkmClient()

	return tfresource.CreateResource(d, sync)
}

func readKmsEkmsPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &KmsEkmsPrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EkmClient()

	return tfresource.ReadResource(sync)
}

func updateKmsEkmsPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &KmsEkmsPrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EkmClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteKmsEkmsPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &KmsEkmsPrivateEndpointResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EkmClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type KmsEkmsPrivateEndpointResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_kms.EkmClient
	Res                    *oci_kms.EkmsPrivateEndpoint
	DisableNotFoundRetries bool
}

func (s *KmsEkmsPrivateEndpointResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *KmsEkmsPrivateEndpointResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_kms.EkmsPrivateEndpointLifecycleStateCreating),
	}
}

func (s *KmsEkmsPrivateEndpointResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_kms.EkmsPrivateEndpointLifecycleStateActive),
	}
}

func (s *KmsEkmsPrivateEndpointResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_kms.EkmsPrivateEndpointLifecycleStateDeleting),
	}
}

func (s *KmsEkmsPrivateEndpointResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_kms.EkmsPrivateEndpointLifecycleStateDeleted),
	}
}

func (s *KmsEkmsPrivateEndpointResourceCrud) Create() error {
	request := oci_kms.CreateEkmsPrivateEndpointRequest{}

	if caBundle, ok := s.D.GetOkExists("ca_bundle"); ok {
		tmp := caBundle.(string)
		request.CaBundle = &tmp
	}

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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if externalKeyManagerIp, ok := s.D.GetOkExists("external_key_manager_ip"); ok {
		tmp := externalKeyManagerIp.(string)
		request.ExternalKeyManagerIp = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if port, ok := s.D.GetOkExists("port"); ok {
		tmp := port.(int)
		request.Port = &tmp
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.CreateEkmsPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.EkmsPrivateEndpoint
	return nil
}

func (s *KmsEkmsPrivateEndpointResourceCrud) Get() error {
	request := oci_kms.GetEkmsPrivateEndpointRequest{}

	tmp := s.D.Id()
	request.EkmsPrivateEndpointId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.GetEkmsPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.EkmsPrivateEndpoint
	return nil
}

func (s *KmsEkmsPrivateEndpointResourceCrud) Update() error {
	request := oci_kms.UpdateEkmsPrivateEndpointRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.EkmsPrivateEndpointId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.UpdateEkmsPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.EkmsPrivateEndpoint
	return nil
}

func (s *KmsEkmsPrivateEndpointResourceCrud) Delete() error {
	request := oci_kms.DeleteEkmsPrivateEndpointRequest{}

	tmp := s.D.Id()
	request.EkmsPrivateEndpointId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	_, err := s.Client.DeleteEkmsPrivateEndpoint(context.Background(), request)
	return err
}

func (s *KmsEkmsPrivateEndpointResourceCrud) SetData() error {
	if s.Res.CaBundle != nil {
		s.D.Set("ca_bundle", *s.Res.CaBundle)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExternalKeyManagerIp != nil {
		s.D.Set("external_key_manager_ip", *s.Res.ExternalKeyManagerIp)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Port != nil {
		s.D.Set("port", *s.Res.Port)
	}

	if s.Res.PrivateEndpointIp != nil {
		s.D.Set("private_endpoint_ip", *s.Res.PrivateEndpointIp)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
