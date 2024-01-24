// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringMonitoredResourcesSearchAssociationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createStackMonitoringMonitoredResourcesSearchAssociation,
		Read:     readStackMonitoringMonitoredResourcesSearchAssociation,
		Delete:   deleteStackMonitoringMonitoredResourcesSearchAssociation,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"association_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"destination_resource_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"destination_resource_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"destination_resource_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"source_resource_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"source_resource_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"source_resource_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"association_type": {
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
						"destination_resource_id": {
							Type:     schema.TypeString,
							Computed: true,
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
						"source_resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createStackMonitoringMonitoredResourcesSearchAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoredResourcesSearchAssociationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.CreateResource(d, sync)
}

func readStackMonitoringMonitoredResourcesSearchAssociation(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteStackMonitoringMonitoredResourcesSearchAssociation(d *schema.ResourceData, m interface{}) error {
	return nil
}

type StackMonitoringMonitoredResourcesSearchAssociationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_stack_monitoring.StackMonitoringClient
	Res                    *oci_stack_monitoring.MonitoredResourceAssociationsCollection
	DisableNotFoundRetries bool
}

func (s *StackMonitoringMonitoredResourcesSearchAssociationResourceCrud) ID() string {
	var id = "/monitoredResources/actions/associateMonitoredResources"

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		id = id + "/compartmentId" + compartmentId.(string)
	}
	if associationType, ok := s.D.GetOkExists("association_type"); ok {
		id = id + "/associationType" + associationType.(string)
	}
	if destinationResourceId, ok := s.D.GetOkExists("destination_resource_id"); ok {
		id = id + "/destinationResourceId" + destinationResourceId.(string)
	}
	if destinationResourceName, ok := s.D.GetOkExists("destination_resource_name"); ok {
		id = id + "/destinationResourceName" + destinationResourceName.(string)
	}
	if destinationResourceType, ok := s.D.GetOkExists("destination_resource_type"); ok {
		id = id + "/destinationResourceType" + destinationResourceType.(string)
	}
	if sourceResourceId, ok := s.D.GetOkExists("source_resource_id"); ok {
		id = id + "/sourceResourceId" + sourceResourceId.(string)
	}
	if sourceResourceName, ok := s.D.GetOkExists("source_resource_name"); ok {
		id = id + "/sourceResourceName" + sourceResourceName.(string)
	}
	if sourceResourceType, ok := s.D.GetOkExists("source_resource_type"); ok {
		id = id + "/sourceResourceType" + sourceResourceType.(string)
	}

	return id
}

func (s *StackMonitoringMonitoredResourcesSearchAssociationResourceCrud) Create() error {
	request := oci_stack_monitoring.SearchMonitoredResourceAssociationsRequest{}

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

	if destinationResourceName, ok := s.D.GetOkExists("destination_resource_name"); ok {
		tmp := destinationResourceName.(string)
		request.DestinationResourceName = &tmp
	}

	if destinationResourceType, ok := s.D.GetOkExists("destination_resource_type"); ok {
		tmp := destinationResourceType.(string)
		request.DestinationResourceType = &tmp
	}

	if sourceResourceId, ok := s.D.GetOkExists("source_resource_id"); ok {
		tmp := sourceResourceId.(string)
		request.SourceResourceId = &tmp
	}

	if sourceResourceName, ok := s.D.GetOkExists("source_resource_name"); ok {
		tmp := sourceResourceName.(string)
		request.SourceResourceName = &tmp
	}

	if sourceResourceType, ok := s.D.GetOkExists("source_resource_type"); ok {
		tmp := sourceResourceType.(string)
		request.SourceResourceType = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.SearchMonitoredResourceAssociations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MonitoredResourceAssociationsCollection
	return nil
}

func (s *StackMonitoringMonitoredResourcesSearchAssociationResourceCrud) SetData() error {
	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MonitoredResourceAssociationSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func MonitoredResourceAssociationSummaryToMap(obj oci_stack_monitoring.MonitoredResourceAssociationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AssociationType != nil {
		result["association_type"] = string(*obj.AssociationType)
	}

	if obj.DestinationResourceDetails != nil {
		result["destination_resource_details"] = []interface{}{AssociationResourceDetailsToMap(obj.DestinationResourceDetails)}
	}

	if obj.DestinationResourceId != nil {
		result["destination_resource_id"] = string(*obj.DestinationResourceId)
	}

	if obj.SourceResourceDetails != nil {
		result["source_resource_details"] = []interface{}{AssociationResourceDetailsToMap(obj.SourceResourceDetails)}
	}

	if obj.SourceResourceId != nil {
		result["source_resource_id"] = string(*obj.SourceResourceId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}
