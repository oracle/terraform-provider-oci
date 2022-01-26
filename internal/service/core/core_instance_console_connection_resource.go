// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_core "github.com/oracle/oci-go-sdk/v56/core"
)

func CoreInstanceConsoleConnectionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreInstanceConsoleConnection,
		Read:     readCoreInstanceConsoleConnection,
		Update:   updateCoreInstanceConsoleConnection,
		Delete:   deleteCoreInstanceConsoleConnection,
		Schema: map[string]*schema.Schema{
			// Required
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"public_key": {
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
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"connection_string": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_host_key_fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vnc_connection_string": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCoreInstanceConsoleConnection(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceConsoleConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreInstanceConsoleConnection(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceConsoleConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

func updateCoreInstanceConsoleConnection(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceConsoleConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreInstanceConsoleConnection(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceConsoleConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreInstanceConsoleConnectionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.ComputeClient
	Res                    *oci_core.InstanceConsoleConnection
	DisableNotFoundRetries bool
}

func (s *CoreInstanceConsoleConnectionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreInstanceConsoleConnectionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.InstanceConsoleConnectionLifecycleStateCreating),
	}
}

func (s *CoreInstanceConsoleConnectionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.InstanceConsoleConnectionLifecycleStateActive),
	}
}

func (s *CoreInstanceConsoleConnectionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.InstanceConsoleConnectionLifecycleStateDeleting),
	}
}

func (s *CoreInstanceConsoleConnectionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.InstanceConsoleConnectionLifecycleStateDeleted),
	}
}

func (s *CoreInstanceConsoleConnectionResourceCrud) Create() error {
	request := oci_core.CreateInstanceConsoleConnectionRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	if publicKey, ok := s.D.GetOkExists("public_key"); ok {
		tmp := publicKey.(string)
		request.PublicKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateInstanceConsoleConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InstanceConsoleConnection
	return nil
}

func (s *CoreInstanceConsoleConnectionResourceCrud) Get() error {
	request := oci_core.GetInstanceConsoleConnectionRequest{}

	tmp := s.D.Id()
	request.InstanceConsoleConnectionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetInstanceConsoleConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InstanceConsoleConnection
	return nil
}

func (s *CoreInstanceConsoleConnectionResourceCrud) Update() error {
	request := oci_core.UpdateInstanceConsoleConnectionRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.InstanceConsoleConnectionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateInstanceConsoleConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InstanceConsoleConnection
	return nil
}

func (s *CoreInstanceConsoleConnectionResourceCrud) Delete() error {
	request := oci_core.DeleteInstanceConsoleConnectionRequest{}

	tmp := s.D.Id()
	request.InstanceConsoleConnectionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteInstanceConsoleConnection(context.Background(), request)
	return err
}

func (s *CoreInstanceConsoleConnectionResourceCrud) SetData() error {
	if publicKey, ok := s.D.GetOkExists("public_key"); ok {
		s.D.Set("public_key", publicKey.(string))
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionString != nil {
		s.D.Set("connection_string", *s.Res.ConnectionString)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Fingerprint != nil {
		s.D.Set("fingerprint", *s.Res.Fingerprint)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InstanceId != nil {
		s.D.Set("instance_id", *s.Res.InstanceId)
	}

	if s.Res.ServiceHostKeyFingerprint != nil {
		s.D.Set("service_host_key_fingerprint", *s.Res.ServiceHostKeyFingerprint)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.VncConnectionString != nil {
		s.D.Set("vnc_connection_string", *s.Res.VncConnectionString)
	}

	return nil
}
