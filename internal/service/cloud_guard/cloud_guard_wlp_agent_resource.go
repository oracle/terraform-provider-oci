// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudGuardWlpAgentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCloudGuardWlpAgent,
		Read:     readCloudGuardWlpAgent,
		Update:   updateCloudGuardWlpAgent,
		Delete:   deleteCloudGuardWlpAgent,
		Schema: map[string]*schema.Schema{
			// Required
			"agent_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"certificate_signed_request": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"os_info": {
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

			// Computed
			"certificate_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"tenant_id": {
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

func createCloudGuardWlpAgent(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardWlpAgentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.CreateResource(d, sync)
}

func readCloudGuardWlpAgent(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardWlpAgentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

func updateCloudGuardWlpAgent(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardWlpAgentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCloudGuardWlpAgent(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardWlpAgentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CloudGuardWlpAgentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_cloud_guard.CloudGuardClient
	Res                    *oci_cloud_guard.WlpAgent
	DisableNotFoundRetries bool
}

func (s *CloudGuardWlpAgentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CloudGuardWlpAgentResourceCrud) Create() error {
	request := oci_cloud_guard.CreateWlpAgentRequest{}

	if agentVersion, ok := s.D.GetOkExists("agent_version"); ok {
		tmp := agentVersion.(string)
		request.AgentVersion = &tmp
	}

	if certificateSignedRequest, ok := s.D.GetOkExists("certificate_signed_request"); ok {
		tmp := certificateSignedRequest.(string)
		request.CertificateSignedRequest = &tmp
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if osInfo, ok := s.D.GetOkExists("os_info"); ok {
		tmp := osInfo.(string)
		request.OsInfo = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.CreateWlpAgent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.WlpAgent
	return nil
}

func (s *CloudGuardWlpAgentResourceCrud) Get() error {
	request := oci_cloud_guard.GetWlpAgentRequest{}

	tmp := s.D.Id()
	request.WlpAgentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.GetWlpAgent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.WlpAgent
	return nil
}

func (s *CloudGuardWlpAgentResourceCrud) Update() error {
	request := oci_cloud_guard.UpdateWlpAgentRequest{}

	if certificateSignedRequest, ok := s.D.GetOkExists("certificate_signed_request"); ok {
		tmp := certificateSignedRequest.(string)
		request.CertificateSignedRequest = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.WlpAgentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.UpdateWlpAgent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.WlpAgent
	return nil
}

func (s *CloudGuardWlpAgentResourceCrud) Delete() error {
	request := oci_cloud_guard.DeleteWlpAgentRequest{}

	tmp := s.D.Id()
	request.WlpAgentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	_, err := s.Client.DeleteWlpAgent(context.Background(), request)
	return err
}

func (s *CloudGuardWlpAgentResourceCrud) SetData() error {
	if s.Res.AgentVersion != nil {
		s.D.Set("agent_version", *s.Res.AgentVersion)
	}

	if s.Res.CertificateId != nil {
		s.D.Set("certificate_id", *s.Res.CertificateId)
	}

	if s.Res.CertificateSignedRequest != nil {
		s.D.Set("certificate_signed_request", *s.Res.CertificateSignedRequest)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HostId != nil {
		s.D.Set("host_id", *s.Res.HostId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TenantId != nil {
		s.D.Set("tenant_id", *s.Res.TenantId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func WlpAgentSummaryToMap(obj oci_cloud_guard.WlpAgentSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AgentVersion != nil {
		result["agent_version"] = string(*obj.AgentVersion)
	}

	if obj.CertificateId != nil {
		result["certificate_id"] = string(*obj.CertificateId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.HostId != nil {
		result["host_id"] = string(*obj.HostId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TenantId != nil {
		result["tenant_id"] = string(*obj.TenantId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
