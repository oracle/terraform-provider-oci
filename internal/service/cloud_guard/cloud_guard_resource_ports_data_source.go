// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

//resource not exposed to user through Terraform, but generated.
//Hence TF team suggested to keep the file commented as codeGen patch build fails if file not present

package cloud_guard

/*import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudGuardResourcePortsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCloudGuardResourcePorts,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"open_port": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_port_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"port_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"process": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readCloudGuardResourcePorts(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardResourcePortsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

type CloudGuardResourcePortsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.ListResourcePortsResponse
}

func (s *CloudGuardResourcePortsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardResourcePortsDataSourceCrud) Get() error {
	request := oci_cloud_guard.ListResourcePortsRequest{}

	if openPort, ok := s.D.GetOkExists("open_port"); ok {
		tmp := openPort.(string)
		request.OpenPort = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_guard")

	response, err := s.Client.ListResourcePorts(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListResourcePorts(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CloudGuardResourcePortsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CloudGuardResourcePortsDataSource-", CloudGuardResourcePortsDataSource(), s.D))
	resources := []map[string]interface{}{}
	resourcePort := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ResourcePortSummaryToMap(item))
	}
	resourcePort["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CloudGuardResourcePortsDataSource().Schema["resource_port_collection"].Elem.(*schema.Resource).Schema)
		resourcePort["items"] = items
	}

	resources = append(resources, resourcePort)
	if err := s.D.Set("resource_port_collection", resources); err != nil {
		return err
	}

	return nil
}

func ResourcePortSummaryToMap(obj oci_cloud_guard.ResourcePortSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.PortNumber != nil {
		result["port_number"] = string(*obj.PortNumber)
	}

	if obj.Process != nil {
		result["process"] = string(*obj.Process)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}*/
