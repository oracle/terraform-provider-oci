// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/crud"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func ServiceGatewayResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createServiceGateway,
		Read:     readServiceGateway,
		Update:   updateServiceGateway,
		Delete:   deleteServiceGateway,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"services": {
				Type:     schema.TypeList,
				Required: true,
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
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"block_traffic": {
				Type:     schema.TypeBool,
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
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createServiceGateway(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.CreateResource(d, sync)
}

func readServiceGateway(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

func updateServiceGateway(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.UpdateResource(d, sync)
}

func deleteServiceGateway(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type ServiceGatewayResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.ServiceGateway
	DisableNotFoundRetries bool
}

func (s *ServiceGatewayResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ServiceGatewayResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.ServiceGatewayLifecycleStateProvisioning),
	}
}

func (s *ServiceGatewayResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.ServiceGatewayLifecycleStateAvailable),
	}
}

func (s *ServiceGatewayResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.ServiceGatewayLifecycleStateTerminating),
	}
}

func (s *ServiceGatewayResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.ServiceGatewayLifecycleStateTerminated),
	}
}

func (s *ServiceGatewayResourceCrud) Create() error {
	request := oci_core.CreateServiceGatewayRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.Services = []oci_core.ServiceIdRequestDetails{}
	if services, ok := s.D.GetOkExists("services"); ok {
		interfaces := services.([]interface{})
		tmp := make([]oci_core.ServiceIdRequestDetails, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = mapToServiceIdRequestDetails(toBeConverted.(map[string]interface{}))
		}
		request.Services = tmp
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateServiceGateway(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ServiceGateway
	return nil
}

func (s *ServiceGatewayResourceCrud) Get() error {
	request := oci_core.GetServiceGatewayRequest{}

	tmp := s.D.Id()
	request.ServiceGatewayId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetServiceGateway(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ServiceGateway
	return nil
}

func (s *ServiceGatewayResourceCrud) Update() error {
	request := oci_core.UpdateServiceGatewayRequest{}

	if blockTraffic, ok := s.D.GetOkExists("block_traffic"); ok {
		tmp := blockTraffic.(bool)
		request.BlockTraffic = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.ServiceGatewayId = &tmp

	request.Services = []oci_core.ServiceIdRequestDetails{}
	if services, ok := s.D.GetOkExists("services"); ok {
		interfaces := services.([]interface{})
		tmp := make([]oci_core.ServiceIdRequestDetails, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = mapToServiceIdRequestDetails(toBeConverted.(map[string]interface{}))
		}
		request.Services = tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateServiceGateway(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ServiceGateway
	return nil
}

func (s *ServiceGatewayResourceCrud) Delete() error {
	request := oci_core.DeleteServiceGatewayRequest{}

	tmp := s.D.Id()
	request.ServiceGatewayId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteServiceGateway(context.Background(), request)
	return err
}

func (s *ServiceGatewayResourceCrud) SetData() {
	if s.Res.BlockTraffic != nil {
		s.D.Set("block_traffic", *s.Res.BlockTraffic)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	services := []interface{}{}
	for _, item := range s.Res.Services {
		services = append(services, ServiceIdResponseDetailsToMap(item))
	}
	s.D.Set("services", services)

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("time_created", s.Res.TimeCreated.String())

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

}

func mapToServiceIdRequestDetails(raw map[string]interface{}) oci_core.ServiceIdRequestDetails {
	result := oci_core.ServiceIdRequestDetails{}

	if serviceId, ok := raw["service_id"]; ok {
		tmp := serviceId.(string)
		result.ServiceId = &tmp
	}

	return result
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
