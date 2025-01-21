// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreByoasnResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreByoasn,
		Read:     readCoreByoasn,
		Update:   updateCoreByoasn,
		Delete:   deleteCoreByoasn,
		Schema: map[string]*schema.Schema{
			// Required
			"asn": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
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
			"byoip_ranges": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"as_path_prepend_length": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"byoip_range_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cidr_block": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv6cidr_block": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_validated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"validation_token": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCoreByoasn(d *schema.ResourceData, m interface{}) error {
	sync := &CoreByoasnResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreByoasn(d *schema.ResourceData, m interface{}) error {
	sync := &CoreByoasnResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

func updateCoreByoasn(d *schema.ResourceData, m interface{}) error {
	sync := &CoreByoasnResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreByoasn(d *schema.ResourceData, m interface{}) error {
	sync := &CoreByoasnResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreByoasnResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.Byoasn
	DisableNotFoundRetries bool
}

func (s *CoreByoasnResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreByoasnResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.ByoasnLifecycleStateCreating),
	}
}

func (s *CoreByoasnResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.ByoasnLifecycleStateActive),
		string(oci_core.ByoasnLifecycleStateCreating),
	}
}

func (s *CoreByoasnResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *CoreByoasnResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.ByoasnLifecycleStateDeleted),
	}
}

func (s *CoreByoasnResourceCrud) Create() error {
	request := oci_core.CreateByoasnRequest{}

	if asn, ok := s.D.GetOkExists("asn"); ok {
		tmp := asn.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert asn string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.Asn = &tmpInt64
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateByoasn(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Byoasn
	return nil
}

func (s *CoreByoasnResourceCrud) Get() error {
	request := oci_core.GetByoasnRequest{}

	tmp := s.D.Id()
	request.ByoasnId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetByoasn(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Byoasn
	return nil
}

func (s *CoreByoasnResourceCrud) Update() error {

	if _, ok := s.D.GetOkExists(""); ok && s.D.HasChange("") {
		err := s.ValidateByoasn()
		if err != nil {
			return err
		}
	}
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateByoasnRequest{}

	tmp := s.D.Id()
	request.ByoasnId = &tmp

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateByoasn(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Byoasn
	return nil
}

func (s *CoreByoasnResourceCrud) Delete() error {
	request := oci_core.DeleteByoasnRequest{}

	tmp := s.D.Id()
	request.ByoasnId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteByoasn(context.Background(), request)
	return err
}

func (s *CoreByoasnResourceCrud) SetData() error {
	if s.Res.Asn != nil {
		s.D.Set("asn", strconv.FormatInt(*s.Res.Asn, 10))
	}

	byoipRanges := []interface{}{}
	for _, item := range s.Res.ByoipRanges {
		byoipRanges = append(byoipRanges, ByoasnByoipRangeToMap(item))
	}
	s.D.Set("byoip_ranges", byoipRanges)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TimeValidated != nil {
		s.D.Set("time_validated", s.Res.TimeValidated.String())
	}

	if s.Res.ValidationToken != nil {
		s.D.Set("validation_token", *s.Res.ValidationToken)
	}

	return nil
}

func (s *CoreByoasnResourceCrud) ValidateByoasn() error {
	request := oci_core.ValidateByoasnRequest{}

	idTmp := s.D.Id()
	request.ByoasnId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ValidateByoasn(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func ByoasnByoipRangeToMap(obj oci_core.ByoasnByoipRange) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AsPathPrependLength != nil {
		result["as_path_prepend_length"] = int(*obj.AsPathPrependLength)
	}

	if obj.ByoipRangeId != nil {
		result["byoip_range_id"] = string(*obj.ByoipRangeId)
	}

	if obj.CidrBlock != nil {
		result["cidr_block"] = string(*obj.CidrBlock)
	}

	if obj.Ipv6CidrBlock != nil {
		result["ipv6cidr_block"] = string(*obj.Ipv6CidrBlock)
	}

	return result
}

func ByoasnSummaryToMap(obj oci_core.ByoasnSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Asn != nil {
		result["asn"] = strconv.FormatInt(*obj.Asn, 10)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func (s *CoreByoasnResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeByoasnCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.ByoasnId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeByoasnCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
