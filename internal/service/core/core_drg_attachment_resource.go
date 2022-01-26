// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_core "github.com/oracle/oci-go-sdk/v56/core"
)

func CoreDrgAttachmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreDrgAttachment,
		Read:     readCoreDrgAttachment,
		Update:   updateCoreDrgAttachment,
		Delete:   deleteCoreDrgAttachment,
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
			"drg_route_table_id": {
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
			"network_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"VCN",
							}, true),
						},

						// Optional
						"route_table_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"ipsec_connection_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"route_table_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"export_drg_route_distribution_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"remove_export_drg_route_distribution_trigger": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_cross_tenancy": {
				Type:     schema.TypeBool,
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

func createCoreDrgAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if _, ok := sync.D.GetOkExists("remove_export_drg_route_distribution_trigger"); ok {
		err := sync.removeExportDrgRouteDistribution()
		if err != nil {
			return err
		}
	}

	if _, ok := sync.D.GetOkExists("export_drg_route_distribution_id"); ok && sync.D.HasChange("export_drg_route_distribution_id") {
		return sync.Update()
	}

	return nil
}

func readCoreDrgAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

func updateCoreDrgAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	if _, ok := sync.D.GetOkExists("remove_export_drg_route_distribution_trigger"); ok &&
		sync.D.HasChange("remove_export_drg_route_distribution_trigger") {
		err := sync.removeExportDrgRouteDistribution()
		if err != nil {
			return err
		}
	}

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreDrgAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreDrgAttachmentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.DrgAttachment
	DisableNotFoundRetries bool
}

func (s *CoreDrgAttachmentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreDrgAttachmentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.DrgAttachmentLifecycleStateAttaching),
	}
}

func (s *CoreDrgAttachmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.DrgAttachmentLifecycleStateAttached),
	}
}

func (s *CoreDrgAttachmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.DrgAttachmentLifecycleStateDetaching),
	}
}

func (s *CoreDrgAttachmentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.DrgAttachmentLifecycleStateDetached),
	}
}

func (s *CoreDrgAttachmentResourceCrud) Create() error {
	request := oci_core.CreateDrgAttachmentRequest{}

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

	if drgRouteTableId, ok := s.D.GetOkExists("drg_route_table_id"); ok {
		tmp := drgRouteTableId.(string)
		request.DrgRouteTableId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if networkDetails, ok := s.D.GetOkExists("network_details"); ok {
		if tmpList := networkDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "network_details", 0)
			tmp, err := s.mapToDrgAttachmentNetworkCreateDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NetworkDetails = tmp
		}
	}

	if routeTableId, ok := s.D.GetOkExists("route_table_id"); ok {
		tmp := routeTableId.(string)
		request.RouteTableId = &tmp
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateDrgAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DrgAttachment

	return nil
}

func (s *CoreDrgAttachmentResourceCrud) Get() error {
	request := oci_core.GetDrgAttachmentRequest{}

	tmp := s.D.Id()
	request.DrgAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetDrgAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DrgAttachment
	return nil
}

func (s *CoreDrgAttachmentResourceCrud) Update() error {
	request := oci_core.UpdateDrgAttachmentRequest{}

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
	request.DrgAttachmentId = &tmp

	if drgRouteTableId, ok := s.D.GetOkExists("drg_route_table_id"); ok && s.D.HasChange("drg_route_table_id") {
		tmp := drgRouteTableId.(string)
		request.DrgRouteTableId = &tmp
	}

	if exportRouteDistributionId, ok := s.D.GetOkExists("export_drg_route_distribution_id"); ok && s.D.HasChange("export_drg_route_distribution_id") {
		if exportRouteDistributionId != nil && exportRouteDistributionId != "" {
			tmp := exportRouteDistributionId.(string)
			request.ExportDrgRouteDistributionId = &tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok && s.D.HasChange("freeform_tags") {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if networkDetails, ok := s.D.GetOkExists("network_details"); ok && s.D.HasChange("network_details") {
		if tmpList := networkDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "network_details", 0)
			tmp, err := s.mapToDrgAttachmentNetworkUpdateDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NetworkDetails = tmp
		}
	}

	if routeTableId, ok := s.D.GetOkExists("route_table_id"); ok && s.D.HasChange("route_table_id") {
		tmp := routeTableId.(string)
		request.RouteTableId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateDrgAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DrgAttachment
	return nil
}

func (s *CoreDrgAttachmentResourceCrud) Delete() error {
	request := oci_core.DeleteDrgAttachmentRequest{}

	tmp := s.D.Id()
	request.DrgAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteDrgAttachment(context.Background(), request)
	return err
}

func (s *CoreDrgAttachmentResourceCrud) SetData() error {
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

	if s.Res.DrgRouteTableId != nil {
		s.D.Set("drg_route_table_id", *s.Res.DrgRouteTableId)
	}

	if s.Res.ExportDrgRouteDistributionId != nil {
		s.D.Set("export_drg_route_distribution_id", *s.Res.ExportDrgRouteDistributionId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsCrossTenancy != nil {
		s.D.Set("is_cross_tenancy", *s.Res.IsCrossTenancy)
	}

	if s.Res.NetworkDetails != nil {
		networkDetailsArray := []interface{}{}
		if networkDetailsMap := DrgAttachmentNetworkDetailsToMap(&s.Res.NetworkDetails); networkDetailsMap != nil {
			networkDetailsArray = append(networkDetailsArray, networkDetailsMap)
		}
		s.D.Set("network_details", networkDetailsArray)
	} else {
		s.D.Set("network_details", nil)
	}

	if s.Res.RouteTableId != nil {
		s.D.Set("route_table_id", *s.Res.RouteTableId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}

func (s *CoreDrgAttachmentResourceCrud) mapToDrgAttachmentNetworkCreateDetails(fieldKeyFormat string) (oci_core.DrgAttachmentNetworkCreateDetails, error) {
	var baseObject oci_core.DrgAttachmentNetworkCreateDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("VCN"):
		details := oci_core.VcnDrgAttachmentNetworkCreateDetails{}
		if routeTableId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "route_table_id")); ok {
			tmp := routeTableId.(string)
			details.RouteTableId = &tmp
		}
		if network_detail_id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
			tmp := network_detail_id.(string)
			details.Id = &tmp
		}

		baseObject = details

	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func (s *CoreDrgAttachmentResourceCrud) mapToDrgAttachmentNetworkUpdateDetails(fieldKeyFormat string) (oci_core.DrgAttachmentNetworkUpdateDetails, error) {
	var baseObject oci_core.DrgAttachmentNetworkUpdateDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("VCN"):
		details := oci_core.VcnDrgAttachmentNetworkUpdateDetails{}
		if routeTableId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "route_table_id")); ok {
			tmp := routeTableId.(string)
			details.RouteTableId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func DrgAttachmentNetworkDetailsToMap(obj *oci_core.DrgAttachmentNetworkDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_core.VcnDrgAttachmentNetworkDetails:
		result["type"] = "VCN"

		if v.RouteTableId != nil {
			result["route_table_id"] = string(*v.RouteTableId)
		}

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}
	return result
}

func (s *CoreDrgAttachmentResourceCrud) removeExportDrgRouteDistribution() error {
	request := oci_core.RemoveExportDrgRouteDistributionRequest{}

	tmp := s.D.Id()
	request.DrgAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.RemoveExportDrgRouteDistribution(context.Background(), request)
	if err != nil {
		return err
	}

	val := s.D.Get("remove_export_drg_route_distribution_trigger")
	s.D.Set("remove_export_drg_route_distribution_trigger", val)

	id := response.ExportDrgRouteDistributionId
	s.D.Set("export_drg_route_distribution_id", id)

	s.Res = &response.DrgAttachment
	return nil
}
