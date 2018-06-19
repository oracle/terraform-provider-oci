// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func InternetGatewayResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createInternetGateway,
		Read:     readInternetGateway,
		Update:   updateInternetGateway,
		Delete:   deleteInternetGateway,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enabled": {
				Type: schema.TypeBool,
				// Keep 'enabled' optional & set the default to true to avoid a breaking change.
				Optional: true,
				Default:  true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
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
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// @Deprecated 01/2018: time_modified (removed)
			"time_modified": {
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: crud.FieldDeprecated("time_modified"),
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createInternetGateway(d *schema.ResourceData, m interface{}) error {
	sync := &InternetGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.CreateResource(d, sync)
}

func readInternetGateway(d *schema.ResourceData, m interface{}) error {
	sync := &InternetGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

func updateInternetGateway(d *schema.ResourceData, m interface{}) error {
	sync := &InternetGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.UpdateResource(d, sync)
}

func deleteInternetGateway(d *schema.ResourceData, m interface{}) error {
	sync := &InternetGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type InternetGatewayResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.InternetGateway
	DisableNotFoundRetries bool
}

func (s *InternetGatewayResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *InternetGatewayResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.InternetGatewayLifecycleStateProvisioning),
	}
}

func (s *InternetGatewayResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.InternetGatewayLifecycleStateAvailable),
	}
}

func (s *InternetGatewayResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.InternetGatewayLifecycleStateTerminating),
	}
}

func (s *InternetGatewayResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.InternetGatewayLifecycleStateTerminated),
	}
}

func (s *InternetGatewayResourceCrud) Create() error {
	request := oci_core.CreateInternetGatewayRequest{}

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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	// TODO: GetOk malfunction with this bool: 'ok' is always the value of the bool
	// newer versions of terraform support GetOkExists which should resolve this problem
	enabledTmp := s.D.Get("enabled").(bool)
	request.IsEnabled = &enabledTmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateInternetGateway(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InternetGateway
	return nil
}

func (s *InternetGatewayResourceCrud) Get() error {
	request := oci_core.GetInternetGatewayRequest{}

	tmp := s.D.Id()
	request.IgId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetInternetGateway(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InternetGateway
	return nil
}

func (s *InternetGatewayResourceCrud) Update() error {
	request := oci_core.UpdateInternetGatewayRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	// TODO: GetOk malfunction with this bool: 'ok' is always the value of the bool
	// newer versions of terraform support GetOkExists which should resolve this problem
	enabledTmp := s.D.Get("enabled").(bool)
	request.IsEnabled = &enabledTmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.IgId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateInternetGateway(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.InternetGateway
	return nil
}

func (s *InternetGatewayResourceCrud) Delete() error {
	request := oci_core.DeleteInternetGatewayRequest{}

	tmp := s.D.Id()
	request.IgId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteInternetGateway(context.Background(), request)
	return err
}

func (s *InternetGatewayResourceCrud) SetData() {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.IsEnabled != nil {
		s.D.Set("enabled", *s.Res.IsEnabled)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

}
