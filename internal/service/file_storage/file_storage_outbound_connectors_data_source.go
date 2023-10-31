// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FileStorageOutboundConnectorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFileStorageOutboundConnectors,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"outbound_connectors": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(FileStorageOutboundConnectorResource()),
			},
		},
	}
}

func readFileStorageOutboundConnectors(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageOutboundConnectorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.ReadResource(sync)
}

type FileStorageOutboundConnectorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_file_storage.FileStorageClient
	Res    *oci_file_storage.ListOutboundConnectorsResponse
}

func (s *FileStorageOutboundConnectorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FileStorageOutboundConnectorsDataSourceCrud) Get() error {
	request := oci_file_storage.ListOutboundConnectorsRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_file_storage.ListOutboundConnectorsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "file_storage")

	response, err := s.Client.ListOutboundConnectors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOutboundConnectors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FileStorageOutboundConnectorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FileStorageOutboundConnectorsDataSource-", FileStorageOutboundConnectorsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		result := map[string]interface{}{}
		switch v := (r).(type) {
		case oci_file_storage.LdapBindAccountSummary:
			result["connector_type"] = "LDAPBIND"

			if v.BindDistinguishedName != nil {
				result["bind_distinguished_name"] = string(*v.BindDistinguishedName)
			}

			endpoints := []interface{}{}
			for _, item := range v.Endpoints {
				endpoints = append(endpoints, EndpointToMap(item))
			}
			result["endpoints"] = endpoints

			if v.AvailabilityDomain != nil {
				result["availability_domain"] = string(*v.AvailabilityDomain)
			}

			if v.CompartmentId != nil {
				result["compartment_id"] = string(*v.CompartmentId)
			}

			if v.DefinedTags != nil {
				result["defined_tags"] = tfresource.DefinedTagsToMap(v.DefinedTags)
			}

			if v.DisplayName != nil {
				result["display_name"] = string(*v.DisplayName)
			}

			result["freeform_tags"] = v.FreeformTags

			if v.Id != nil {
				result["id"] = string(*v.Id)
			}

			result["state"] = string(v.LifecycleState)

			if v.TimeCreated != nil {
				result["time_created"] = v.TimeCreated.String()
			}
		default:
			log.Printf("[WARN] Received 'connector_type' of unknown type %v", r)
			return nil
		}

		resources = append(resources, result)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, FileStorageOutboundConnectorsDataSource().Schema["outbound_connectors"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("outbound_connectors", resources); err != nil {
		return err
	}

	return nil
}
