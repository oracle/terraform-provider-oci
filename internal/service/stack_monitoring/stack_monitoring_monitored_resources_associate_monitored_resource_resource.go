// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"
)

func StackMonitoringMonitoredResourcesAssociateMonitoredResourceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createStackMonitoringMonitoredResourcesAssociateMonitoredResource,
		Read:     readStackMonitoringMonitoredResourcesAssociateMonitoredResource,
		Delete:   deleteStackMonitoringMonitoredResourcesAssociateMonitoredResource,
		Schema: map[string]*schema.Schema{
			// Required
			"association_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"destination_resource_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_resource_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"destination_resource_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"source_resource_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tenant_id": {
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

func createStackMonitoringMonitoredResourcesAssociateMonitoredResource(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoredResourcesAssociateMonitoredResourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.CreateResource(d, sync)
}

func readStackMonitoringMonitoredResourcesAssociateMonitoredResource(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteStackMonitoringMonitoredResourcesAssociateMonitoredResource(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoredResourcesAssociateMonitoredResourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)

}

type StackMonitoringMonitoredResourcesAssociateMonitoredResourceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_stack_monitoring.StackMonitoringClient
	Res                    *oci_stack_monitoring.MonitoredResourceAssociation
	DisableNotFoundRetries bool
}

func (s *StackMonitoringMonitoredResourcesAssociateMonitoredResourceResourceCrud) ID() string {
	return "compartmentId/" + *s.Res.CompartmentId + "associationType/" + *s.Res.AssociationType + "source/" + *s.Res.SourceResourceId + "destination/" + *s.Res.DestinationResourceId
}

func (s *StackMonitoringMonitoredResourcesAssociateMonitoredResourceResourceCrud) Create() error {
	request := oci_stack_monitoring.AssociateMonitoredResourcesRequest{}

	if associationType, ok := s.D.GetOkExists("association_type"); ok {
		tmp := associationType.(string)
		request.AssociationType = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if destinationResourceId, ok := s.D.GetOkExists("destination_resource_id"); ok {
		tmp := destinationResourceId.(string)
		request.DestinationResourceId = &tmp
	}

	if sourceResourceId, ok := s.D.GetOkExists("source_resource_id"); ok {
		tmp := sourceResourceId.(string)
		request.SourceResourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.AssociateMonitoredResources(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MonitoredResourceAssociation
	return nil
}

func (s *StackMonitoringMonitoredResourcesAssociateMonitoredResourceResourceCrud) Delete() error {
	request := oci_stack_monitoring.DisassociateMonitoredResourcesRequest{}

	if associationType, ok := s.D.GetOkExists("association_type"); ok {
		tmp := associationType.(string)
		request.AssociationType = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if destinationResourceId, ok := s.D.GetOkExists("destination_resource_id"); ok {
		tmp := destinationResourceId.(string)
		request.DestinationResourceId = &tmp
	}

	if sourceResourceId, ok := s.D.GetOkExists("source_resource_id"); ok {
		tmp := sourceResourceId.(string)
		request.SourceResourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	_, err := s.Client.DisassociateMonitoredResources(context.Background(), request)

	return err
}

func (s *StackMonitoringMonitoredResourcesAssociateMonitoredResourceResourceCrud) SetData() error {
	if s.Res.AssociationType != nil {
		s.D.Set("association_type", *s.Res.AssociationType)
	}

	s.D.Set("category", s.Res.Category)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DestinationResourceDetails != nil {
		s.D.Set("destination_resource_details", []interface{}{AssociationResourceDetailsToMap(s.Res.DestinationResourceDetails)})
	} else {
		s.D.Set("destination_resource_details", nil)
	}

	if s.Res.DestinationResourceId != nil {
		s.D.Set("destination_resource_id", *s.Res.DestinationResourceId)
	}

	if s.Res.SourceResourceDetails != nil {
		s.D.Set("source_resource_details", []interface{}{AssociationResourceDetailsToMap(s.Res.SourceResourceDetails)})
	} else {
		s.D.Set("source_resource_details", nil)
	}

	if s.Res.SourceResourceId != nil {
		s.D.Set("source_resource_id", *s.Res.SourceResourceId)
	}

	if s.Res.TenantId != nil {
		s.D.Set("tenant_id", *s.Res.TenantId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func AssociationResourceDetailsToMap(obj *oci_stack_monitoring.AssociationResourceDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}
