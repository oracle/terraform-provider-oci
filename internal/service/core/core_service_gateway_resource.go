// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"bytes"
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_core "github.com/oracle/oci-go-sdk/v56/core"
)

func CoreServiceGatewayResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreServiceGateway,
		Read:     readCoreServiceGateway,
		Update:   updateCoreServiceGateway,
		Delete:   deleteCoreServiceGateway,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"services": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      servicesHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"service_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"service_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
			"route_table_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"block_traffic": {
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

func createCoreServiceGateway(d *schema.ResourceData, m interface{}) error {
	sync := &CoreServiceGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreServiceGateway(d *schema.ResourceData, m interface{}) error {
	sync := &CoreServiceGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

func updateCoreServiceGateway(d *schema.ResourceData, m interface{}) error {
	sync := &CoreServiceGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreServiceGateway(d *schema.ResourceData, m interface{}) error {
	sync := &CoreServiceGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreServiceGatewayResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.ServiceGateway
	DisableNotFoundRetries bool
}

func (s *CoreServiceGatewayResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreServiceGatewayResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.ServiceGatewayLifecycleStateProvisioning),
	}
}

func (s *CoreServiceGatewayResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.ServiceGatewayLifecycleStateAvailable),
	}
}

func (s *CoreServiceGatewayResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.ServiceGatewayLifecycleStateTerminating),
	}
}

func (s *CoreServiceGatewayResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.ServiceGatewayLifecycleStateTerminated),
	}
}

func (s *CoreServiceGatewayResourceCrud) Create() error {
	request := oci_core.CreateServiceGatewayRequest{}

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if routeTableId, ok := s.D.GetOkExists("route_table_id"); ok {
		tmp := routeTableId.(string)
		request.RouteTableId = &tmp
	}

	if services, ok := s.D.GetOkExists("services"); ok {
		set := services.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_core.ServiceIdRequestDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := servicesHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "services", stateDataIndex)
			converted, err := s.mapToServiceIdRequestDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("services") {
			request.Services = tmp
		}
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateServiceGateway(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ServiceGateway
	return nil
}

func (s *CoreServiceGatewayResourceCrud) Get() error {
	request := oci_core.GetServiceGatewayRequest{}

	tmp := s.D.Id()
	request.ServiceGatewayId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetServiceGateway(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ServiceGateway
	return nil
}

func (s *CoreServiceGatewayResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateServiceGatewayRequest{}

	if blockTraffic, ok := s.D.GetOkExists("block_traffic"); ok {
		tmp := blockTraffic.(bool)
		request.BlockTraffic = &tmp
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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if routeTableId, ok := s.D.GetOkExists("route_table_id"); ok {
		tmp := routeTableId.(string)
		request.RouteTableId = &tmp
	}

	tmp := s.D.Id()
	request.ServiceGatewayId = &tmp

	if services, ok := s.D.GetOkExists("services"); ok {
		set := services.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_core.ServiceIdRequestDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := servicesHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "services", stateDataIndex)
			converted, err := s.mapToServiceIdRequestDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("services") {
			request.Services = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateServiceGateway(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ServiceGateway
	return nil
}

func (s *CoreServiceGatewayResourceCrud) Delete() error {
	request := oci_core.DeleteServiceGatewayRequest{}

	tmp := s.D.Id()
	request.ServiceGatewayId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteServiceGateway(context.Background(), request)
	return err
}

func (s *CoreServiceGatewayResourceCrud) SetData() error {
	if s.Res.BlockTraffic != nil {
		s.D.Set("block_traffic", *s.Res.BlockTraffic)
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

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.RouteTableId != nil {
		s.D.Set("route_table_id", *s.Res.RouteTableId)
	}

	services := []interface{}{}
	for _, item := range s.Res.Services {
		services = append(services, ServiceIdResponseDetailsToMap(item))
	}
	s.D.Set("services", schema.NewSet(servicesHashCodeForSets, services))

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}

func (s *CoreServiceGatewayResourceCrud) mapToServiceIdRequestDetails(fieldKeyFormat string) (oci_core.ServiceIdRequestDetails, error) {
	result := oci_core.ServiceIdRequestDetails{}

	if serviceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service_id")); ok {
		tmp := serviceId.(string)
		result.ServiceId = &tmp
	}

	return result, nil
}

func ServiceIdResponseDetailsToMap(obj oci_core.ServiceIdResponseDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ServiceId != nil {
		result["service_id"] = string(*obj.ServiceId)
	}

	if obj.ServiceName != nil {
		result["service_name"] = string(*obj.ServiceName)
	}

	return result
}

func servicesHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if serviceId, ok := m["service_id"]; ok && serviceId != "" {
		buf.WriteString(fmt.Sprintf("%v-", serviceId))
	}
	return hashcode.String(buf.String())
}
func (s *CoreServiceGatewayResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeServiceGatewayCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ServiceGatewayId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeServiceGatewayCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
