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

func CoreDrgRouteTableResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreDrgRouteTable,
		Read:     readCoreDrgRouteTable,
		Update:   updateCoreDrgRouteTable,
		Delete:   deleteCoreDrgRouteTable,
		Schema: map[string]*schema.Schema{
			// Required
			"drg_id": {
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
			"import_drg_route_distribution_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"is_ecmp_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"remove_import_trigger": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			// Computed
			"compartment_id": {
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

func createCoreDrgRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgRouteTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if _, ok := sync.D.GetOkExists("remove_import_trigger"); ok {
		err := sync.RemoveImportRouteDistribution()
		if err != nil {
			return err
		}
	}

	return nil
}

func readCoreDrgRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgRouteTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

func updateCoreDrgRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgRouteTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	if _, ok := sync.D.GetOkExists("remove_import_trigger"); ok &&
		sync.D.HasChange("remove_import_trigger") {
		err := sync.RemoveImportRouteDistribution()
		if err != nil {
			return err
		}
	}

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreDrgRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgRouteTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreDrgRouteTableResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.DrgRouteTable
	DisableNotFoundRetries bool
}

func (s *CoreDrgRouteTableResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreDrgRouteTableResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.DrgRouteTableLifecycleStateProvisioning),
	}
}

func (s *CoreDrgRouteTableResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.DrgRouteTableLifecycleStateAvailable),
	}
}

func (s *CoreDrgRouteTableResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.DrgRouteTableLifecycleStateTerminating),
	}
}

func (s *CoreDrgRouteTableResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.DrgRouteTableLifecycleStateTerminated),
	}
}

func (s *CoreDrgRouteTableResourceCrud) Create() error {
	request := oci_core.CreateDrgRouteTableRequest{}

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

	if drgId, ok := s.D.GetOkExists("drg_id"); ok {
		tmp := drgId.(string)
		request.DrgId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if importDrgRouteDistributionId, ok := s.D.GetOkExists("import_drg_route_distribution_id"); ok {
		tmp := importDrgRouteDistributionId.(string)
		request.ImportDrgRouteDistributionId = &tmp
	}

	if isEcmpEnabled, ok := s.D.GetOkExists("is_ecmp_enabled"); ok {
		tmp := isEcmpEnabled.(bool)
		request.IsEcmpEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateDrgRouteTable(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DrgRouteTable
	return nil
}

func (s *CoreDrgRouteTableResourceCrud) Get() error {
	request := oci_core.GetDrgRouteTableRequest{}

	tmp := s.D.Id()
	request.DrgRouteTableId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetDrgRouteTable(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DrgRouteTable
	return nil
}

func (s *CoreDrgRouteTableResourceCrud) Update() error {
	request := oci_core.UpdateDrgRouteTableRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok && s.D.HasChange("defined_tags") {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok && s.D.HasChange("display_name") {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.DrgRouteTableId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok && s.D.HasChange("freeform_tags") {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if importDrgRouteDistributionId, ok := s.D.GetOkExists("import_drg_route_distribution_id"); ok && s.D.HasChange("import_drg_route_distribution_id") {
		if importDrgRouteDistributionId != nil && importDrgRouteDistributionId != "" {
			tmp := importDrgRouteDistributionId.(string)
			request.ImportDrgRouteDistributionId = &tmp
		}
	}

	if isEcmpEnabled, ok := s.D.GetOkExists("is_ecmp_enabled"); ok && s.D.HasChange("is_ecmp_enabled") {
		tmp := isEcmpEnabled.(bool)
		request.IsEcmpEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateDrgRouteTable(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DrgRouteTable
	return nil
}

func (s *CoreDrgRouteTableResourceCrud) RemoveImportRouteDistribution() error {
	request := oci_core.RemoveImportDrgRouteDistributionRequest{}

	tmp := s.D.Id()
	request.DrgRouteTableId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.RemoveImportDrgRouteDistribution(context.Background(), request)
	if err != nil {
		return err
	}

	val := s.D.Get("remove_import_trigger")
	s.D.Set("remove_import_trigger", val)

	id := response.ImportDrgRouteDistributionId
	s.D.Set("import_drg_route_distribution_id", id)

	s.Res = &response.DrgRouteTable
	return nil

}
func (s *CoreDrgRouteTableResourceCrud) Delete() error {
	request := oci_core.DeleteDrgRouteTableRequest{}

	tmp := s.D.Id()
	request.DrgRouteTableId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.DeleteDrgRouteTable(context.Background(), request)
	if response.HTTPResponse().StatusCode == 400 {
		retryDeleteFunc := func() bool {
			response, err := s.Client.DeleteDrgRouteTable(context.Background(), request)
			return err == nil && response.HTTPResponse().StatusCode == 200
		}
		return tfresource.WaitForResourceCondition(s, retryDeleteFunc, s.D.Timeout(schema.TimeoutDelete))
	}
	return err
}

func (s *CoreDrgRouteTableResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DrgId != nil {
		s.D.Set("drg_id", *s.Res.DrgId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ImportDrgRouteDistributionId != nil {
		s.D.Set("import_drg_route_distribution_id", *s.Res.ImportDrgRouteDistributionId)
	}

	if s.Res.IsEcmpEnabled != nil {
		s.D.Set("is_ecmp_enabled", *s.Res.IsEcmpEnabled)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
