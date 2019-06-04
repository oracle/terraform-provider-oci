// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func CoreInternetGatewayResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createCoreInternetGateway,
		Read:     readCoreInternetGateway,
		Update:   updateCoreInternetGateway,
		Delete:   deleteCoreInternetGateway,
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
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// @Deprecated 01/2018: time_modified (removed)
			"time_modified": {
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: FieldDeprecated("time_modified"),
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCoreInternetGateway(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInternetGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return CreateResource(d, sync)
}

func readCoreInternetGateway(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInternetGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return ReadResource(sync)
}

func updateCoreInternetGateway(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInternetGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return UpdateResource(d, sync)
}

func deleteCoreInternetGateway(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInternetGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type CoreInternetGatewayResourceCrud struct {
	BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.InternetGateway
	DisableNotFoundRetries bool
}

func (s *CoreInternetGatewayResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreInternetGatewayResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.InternetGatewayLifecycleStateProvisioning),
	}
}

func (s *CoreInternetGatewayResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.InternetGatewayLifecycleStateAvailable),
	}
}

func (s *CoreInternetGatewayResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.InternetGatewayLifecycleStateTerminating),
	}
}

func (s *CoreInternetGatewayResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.InternetGatewayLifecycleStateTerminated),
	}
}

func (s *CoreInternetGatewayResourceCrud) Create() error {
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

	if enabled, ok := s.D.GetOkExists("enabled"); ok {
		tmp := enabled.(bool)
		request.IsEnabled = &tmp
	}

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

func (s *CoreInternetGatewayResourceCrud) Get() error {
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

func (s *CoreInternetGatewayResourceCrud) Update() error {
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

	if enabled, ok := s.D.GetOkExists("enabled"); ok {
		tmp := enabled.(bool)
		request.IsEnabled = &tmp
	}

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

func (s *CoreInternetGatewayResourceCrud) Delete() error {
	request := oci_core.DeleteInternetGatewayRequest{}

	tmp := s.D.Id()
	request.IgId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteInternetGateway(context.Background(), request)
	return err
}

func (s *CoreInternetGatewayResourceCrud) SetData() error {
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

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}
