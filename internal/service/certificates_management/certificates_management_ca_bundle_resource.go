// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package certificates_management

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_certificates_management "github.com/oracle/oci-go-sdk/v65/certificatesmanagement"
)

func CertificatesManagementCaBundleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCertificatesManagementCaBundle,
		Read:     readCertificatesManagementCaBundle,
		Update:   updateCertificatesManagementCaBundle,
		Delete:   deleteCertificatesManagementCaBundle,
		Schema: map[string]*schema.Schema{
			// Required
			"ca_bundle_pem": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
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
			"lifecycle_details": {
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
		},
	}
}

func createCertificatesManagementCaBundle(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesManagementCaBundleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CertificatesManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readCertificatesManagementCaBundle(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesManagementCaBundleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CertificatesManagementClient()

	return tfresource.ReadResource(sync)
}

func updateCertificatesManagementCaBundle(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesManagementCaBundleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CertificatesManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCertificatesManagementCaBundle(d *schema.ResourceData, m interface{}) error {
	sync := &CertificatesManagementCaBundleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CertificatesManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CertificatesManagementCaBundleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_certificates_management.CertificatesManagementClient
	Res                    *oci_certificates_management.CaBundle
	DisableNotFoundRetries bool
}

func (s *CertificatesManagementCaBundleResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CertificatesManagementCaBundleResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_certificates_management.CaBundleLifecycleStateCreating),
	}
}

func (s *CertificatesManagementCaBundleResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_certificates_management.CaBundleLifecycleStateActive),
	}
}

func (s *CertificatesManagementCaBundleResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_certificates_management.CaBundleLifecycleStateUpdating),
	}
}

func (s *CertificatesManagementCaBundleResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_certificates_management.CaBundleLifecycleStateActive),
	}
}

func (s *CertificatesManagementCaBundleResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_certificates_management.CaBundleLifecycleStateDeleting),
	}
}

func (s *CertificatesManagementCaBundleResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_certificates_management.CaBundleLifecycleStateDeleted),
	}
}

func (s *CertificatesManagementCaBundleResourceCrud) Create() error {
	request := oci_certificates_management.CreateCaBundleRequest{}

	if caBundlePem, ok := s.D.GetOkExists("ca_bundle_pem"); ok {
		tmp := caBundlePem.(string)
		request.CaBundlePem = &tmp
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "certificates_management")

	response, err := s.Client.CreateCaBundle(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CaBundle
	return nil
}

func (s *CertificatesManagementCaBundleResourceCrud) Get() error {
	request := oci_certificates_management.GetCaBundleRequest{}

	tmp := s.D.Id()
	request.CaBundleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "certificates_management")

	response, err := s.Client.GetCaBundle(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CaBundle
	return nil
}

func (s *CertificatesManagementCaBundleResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_certificates_management.UpdateCaBundleRequest{}

	tmp := s.D.Id()
	request.CaBundleId = &tmp

	if caBundlePem, ok := s.D.GetOkExists("ca_bundle_pem"); ok && s.D.HasChange("ca_bundle_pem") {
		tmp := caBundlePem.(string)
		request.CaBundlePem = &tmp
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

	if request.CaBundlePem != nil || request.DefinedTags != nil || request.Description != nil || request.FreeformTags != nil {
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "certificates_management")

		response, err := s.Client.UpdateCaBundle(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res = &response.CaBundle
	}

	return nil
}

func (s *CertificatesManagementCaBundleResourceCrud) Delete() error {
	request := oci_certificates_management.DeleteCaBundleRequest{}

	tmp := s.D.Id()
	request.CaBundleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "certificates_management")

	_, err := s.Client.DeleteCaBundle(context.Background(), request)
	return err
}

func (s *CertificatesManagementCaBundleResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func CaBundleSummaryToMap(obj oci_certificates_management.CaBundleSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = *obj.CompartmentId
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = *obj.Description
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = *obj.Id
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = *obj.LifecycleDetails
	}

	if obj.Name != nil {
		result["name"] = *obj.Name
	}

	result["state"] = obj.LifecycleState

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func (s *CertificatesManagementCaBundleResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_certificates_management.ChangeCaBundleCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.CaBundleId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "certificates_management")

	_, err := s.Client.ChangeCaBundleCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
