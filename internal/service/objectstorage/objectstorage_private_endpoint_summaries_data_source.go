// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package objectstorage

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/v65/objectstorage"
)

func ObjectStoragePrivateEndpointsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readObjectStoragePrivateEndpoints,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"private_endpoint_summaries": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(ObjectStoragePrivateEndpointResource()),
			},
		},
	}
}

func readObjectStoragePrivateEndpoints(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStoragePrivateEndpointsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()

	return tfresource.ReadResource(sync)
}

type ObjectStoragePrivateEndpointsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_object_storage.ObjectStorageClient
	Res    *oci_object_storage.ListPrivateEndpointsResponse
}

func (s *ObjectStoragePrivateEndpointsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ObjectStoragePrivateEndpointsDataSourceCrud) Get() error {
	request := oci_object_storage.ListPrivateEndpointsRequest{
		Fields: oci_object_storage.GetListPrivateEndpointsFieldsEnumValues(),
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "object_storage")

	response, err := s.Client.ListPrivateEndpoints(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPrivateEndpoints(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ObjectStoragePrivateEndpointsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ObjectStoragePrivateEndpointsDataSource-", ObjectStoragePrivateEndpointsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		pe := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
			"namespace":      *r.Namespace,
		}

		if r.CreatedBy != nil {
			pe["created_by"] = *r.CreatedBy
		}

		if r.Etag != nil {
			pe["etag"] = *r.Etag
		}

		if r.Name != nil {
			pe["name"] = *r.Name
		}

		if r.TimeCreated != nil {
			pe["time_created"] = r.TimeCreated.String()
		}

		if r.Prefix != nil {
			pe["prefix"] = *r.Prefix
		}

		if r.Etag != nil {
			pe["etag"] = *r.Etag
		}

		if r.Fqdns != nil {
			fqdnsMap := make(map[string]interface{})

			// Check and set prefixFqdns if available
			if prefixFqdns, ok := s.D.GetOk("fqdns.prefixFqdns"); ok {
				prefixFqdnsMap := make(map[string]interface{})
				if m, ok := prefixFqdns.(map[string]interface{}); ok {
					if v, ok := m["objectStorageApiFqdn"].(string); ok {
						prefixFqdnsMap["objectStorageApiFqdn"] = v
					}
					if v, ok := m["s3CompatibilityApiFqdn"].(string); ok {
						prefixFqdnsMap["s3CompatibilityApiFqdn"] = v
					}
					if v, ok := m["swiftApiFqdn"].(string); ok {
						prefixFqdnsMap["swiftApiFqdn"] = v
					}
				}
				fqdnsMap["prefixFqdns"] = prefixFqdnsMap
			}

			// Check and set additionalPrefixesFqdns if available
			if additionalPrefixesFqdns, ok := s.D.GetOk("fqdns.additionalPrefixesFqdns"); ok {
				additionalPrefixesFqdnsMap := make(map[string]interface{})
				if m, ok := additionalPrefixesFqdns.(map[string]interface{}); ok {
					for key, value := range m {
						if prefixFqdns, ok := value.(map[string]interface{}); ok {
							prefixFqdnsMap := make(map[string]interface{})
							if v, ok := prefixFqdns["objectStorageApiFqdn"].(string); ok {
								prefixFqdnsMap["objectStorageApiFqdn"] = v
							}
							if v, ok := prefixFqdns["s3CompatibilityApiFqdn"].(string); ok {
								prefixFqdnsMap["s3CompatibilityApiFqdn"] = v
							}
							if v, ok := prefixFqdns["swiftApiFqdn"].(string); ok {
								prefixFqdnsMap["swiftApiFqdn"] = v
							}
							additionalPrefixesFqdnsMap[key] = prefixFqdnsMap
						}
					}
				}
				fqdnsMap["additionalPrefixesFqdns"] = additionalPrefixesFqdnsMap
			}
			pe["fqdns"] = fqdnsMap
		}

		resources = append(resources, pe)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, ObjectStoragePrivateEndpointsDataSource().Schema["private_endpoint_summaries"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("private_endpoint_summaries", resources); err != nil {
		return err
	}

	return nil
}
