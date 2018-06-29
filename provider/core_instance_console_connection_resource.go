// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func InstanceConsoleConnectionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createInstanceConsoleConnection,
		Read:     readInstanceConsoleConnection,
		Delete:   deleteInstanceConsoleConnection,
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
				ForceNew:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
			"id": {
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

func createInstanceConsoleConnection(d *schema.ResourceData, m interface{}) error {
	sync := &InstanceConsoleConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return crud.CreateResource(d, sync)
}

func readInstanceConsoleConnection(d *schema.ResourceData, m interface{}) error {
	sync := &InstanceConsoleConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient

	return crud.ReadResource(sync)
}

func deleteInstanceConsoleConnection(d *schema.ResourceData, m interface{}) error {
	sync := &InstanceConsoleConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type InstanceConsoleConnectionResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_core.ComputeClient
	Res                    *oci_core.InstanceConsoleConnection
	DisableNotFoundRetries bool
}

func (s *InstanceConsoleConnectionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *InstanceConsoleConnectionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.InstanceConsoleConnectionLifecycleStateCreating),
	}
}

func (s *InstanceConsoleConnectionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.InstanceConsoleConnectionLifecycleStateActive),
	}
}

func (s *InstanceConsoleConnectionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.InstanceConsoleConnectionLifecycleStateDeleting),
	}
}

func (s *InstanceConsoleConnectionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.InstanceConsoleConnectionLifecycleStateDeleted),
	}
}

func (s *InstanceConsoleConnectionResourceCrud) Create() error {
	request := oci_core.CreateInstanceConsoleConnectionRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	if publicKey, ok := s.D.GetOkExists("public_key"); ok {
		tmp := publicKey.(string)
		request.PublicKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateInstanceConsoleConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InstanceConsoleConnection
	return nil
}

func (s *InstanceConsoleConnectionResourceCrud) Get() error {
	request := oci_core.GetInstanceConsoleConnectionRequest{}

	tmp := s.D.Id()
	request.InstanceConsoleConnectionId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetInstanceConsoleConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InstanceConsoleConnection
	return nil
}

func (s *InstanceConsoleConnectionResourceCrud) Delete() error {
	request := oci_core.DeleteInstanceConsoleConnectionRequest{}

	tmp := s.D.Id()
	request.InstanceConsoleConnectionId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteInstanceConsoleConnection(context.Background(), request)
	return err
}

func (s *InstanceConsoleConnectionResourceCrud) SetData() {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionString != nil {
		s.D.Set("connection_string", *s.Res.ConnectionString)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Fingerprint != nil {
		s.D.Set("fingerprint", *s.Res.Fingerprint)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.InstanceId != nil {
		s.D.Set("instance_id", *s.Res.InstanceId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.VncConnectionString != nil {
		s.D.Set("vnc_connection_string", *s.Res.VncConnectionString)
	}

}
