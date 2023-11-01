// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringMonitoredResourcesListMemberResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createStackMonitoringMonitoredResourcesListMember,
		Read:     readStackMonitoringMonitoredResourcesListMember,
		Delete:   deleteStackMonitoringMonitoredResourcesListMember,
		Schema: map[string]*schema.Schema{
			// Required
			"monitored_resource_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"destination_resource_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"limit_level": {
				Type:     schema.TypeInt,
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
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"defined_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"external_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"host_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"license": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parent_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"system_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
					},
				},
			},
		},
	}
}

func createStackMonitoringMonitoredResourcesListMember(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoredResourcesListMemberResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.CreateResource(d, sync)
}

func readStackMonitoringMonitoredResourcesListMember(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteStackMonitoringMonitoredResourcesListMember(d *schema.ResourceData, m interface{}) error {
	return nil
}

type StackMonitoringMonitoredResourcesListMemberResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_stack_monitoring.StackMonitoringClient
	Res                    *oci_stack_monitoring.MonitoredResourceMembersCollection
	DisableNotFoundRetries bool
}

func (s *StackMonitoringMonitoredResourcesListMemberResourceCrud) ID() string {
	var id = "/monitoredResources/"
	if monitoredResourceId, ok := s.D.GetOkExists("monitored_resource_id"); ok {
		id = id + monitoredResourceId.(string)
	}
	id = id + "/actions/listMembers"
	if destinationResourceId, ok := s.D.GetOkExists("destination_resource_id"); ok {
		id = id + "/destinationResource" + destinationResourceId.(string)
	}
	if limitLevel, ok := s.D.GetOkExists("limit_level"); ok {
		id = id + "/limitLevel" + limitLevel.(string)
	}
	return id
}

func (s *StackMonitoringMonitoredResourcesListMemberResourceCrud) Create() error {
	request := oci_stack_monitoring.SearchMonitoredResourceMembersRequest{}

	if destinationResourceId, ok := s.D.GetOkExists("destination_resource_id"); ok {
		tmp := destinationResourceId.(string)
		request.DestinationResourceId = &tmp
	}

	if limitLevel, ok := s.D.GetOkExists("limit_level"); ok {
		tmp := limitLevel.(int)
		request.LimitLevel = &tmp
	}

	if monitoredResourceId, ok := s.D.GetOkExists("monitored_resource_id"); ok {
		tmp := monitoredResourceId.(string)
		request.MonitoredResourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.SearchMonitoredResourceMembers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MonitoredResourceMembersCollection
	return nil
}

func (s *StackMonitoringMonitoredResourcesListMemberResourceCrud) SetData() error {
	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MonitoredResourceMemberSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func MonitoredResourceMemberSummaryToMap(obj oci_stack_monitoring.MonitoredResourceMemberSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.ExternalId != nil {
		result["external_id"] = string(*obj.ExternalId)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.HostName != nil {
		result["host_name"] = string(*obj.HostName)
	}

	result["license"] = string(obj.License)

	if obj.ParentId != nil {
		result["parent_id"] = string(*obj.ParentId)
	}

	if obj.ResourceDisplayName != nil {
		result["resource_display_name"] = string(*obj.ResourceDisplayName)
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	if obj.ResourceName != nil {
		result["resource_name"] = string(*obj.ResourceName)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	return result
}
