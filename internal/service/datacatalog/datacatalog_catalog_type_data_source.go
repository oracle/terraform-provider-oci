// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datacatalog

import (
	"context"
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datacatalog "github.com/oracle/oci-go-sdk/v65/datacatalog"
)

func DatacatalogCatalogTypeDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatacatalogCatalogType,
		Schema: map[string]*schema.Schema{
			"catalog_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fields": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"type_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_type_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_approved": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_internal": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_tag": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"properties": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type_category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDatacatalogCatalogType(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogCatalogTypeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataCatalogClient()

	return tfresource.ReadResource(sync)
}

type DatacatalogCatalogTypeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datacatalog.DataCatalogClient
	Res    *oci_datacatalog.GetTypeResponse
}

func (s *DatacatalogCatalogTypeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatacatalogCatalogTypeDataSourceCrud) Get() error {
	request := oci_datacatalog.GetTypeRequest{}

	if catalogId, ok := s.D.GetOkExists("catalog_id"); ok {
		tmp := catalogId.(string)
		request.CatalogId = &tmp
	}

	if fields, ok := s.D.GetOkExists("fields"); ok {
		interfaces := fields.([]interface{})
		tmp := make([]oci_datacatalog.GetTypeFieldsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_datacatalog.GetTypeFieldsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("fields") {
			request.Fields = tmp
		}
	}

	if typeKey, ok := s.D.GetOkExists("type_key"); ok {
		tmp := typeKey.(string)
		request.TypeKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datacatalog")

	response, err := s.Client.GetType(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatacatalogCatalogTypeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatacatalogCatalogTypeDataSource-", DatacatalogCatalogTypeDataSource(), s.D))

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.ExternalTypeName != nil {
		s.D.Set("external_type_name", *s.Res.ExternalTypeName)
	}

	if s.Res.IsApproved != nil {
		s.D.Set("is_approved", *s.Res.IsApproved)
	}

	if s.Res.IsInternal != nil {
		s.D.Set("is_internal", *s.Res.IsInternal)
	}

	if s.Res.IsTag != nil {
		s.D.Set("is_tag", *s.Res.IsTag)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Properties != nil {
		s.D.Set("properties", typePropertiesToMap(s.Res.Properties))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TypeCategory != nil {
		s.D.Set("type_category", *s.Res.TypeCategory)
	}

	if s.Res.Uri != nil {
		s.D.Set("uri", *s.Res.Uri)
	}

	return nil
}

func typePropertiesToMap(properties map[string][]oci_datacatalog.PropertyDefinition) map[string]interface{} {
	var rtn = make(map[string]interface{})
	if len(properties) > 0 {
		for namespace, pds := range properties {
			for i, pd := range pds {
				if pd.Name != nil {
					rtn[namespace+"."+strconv.Itoa(i)+"."+"name"] = *pd.Name
				}

				if pd.Type != nil {
					rtn[namespace+"."+strconv.Itoa(i)+"."+"type"] = *pd.Type
				}

				if pd.IsRequired != nil {
					rtn[namespace+"."+strconv.Itoa(i)+"."+"is_required"] = strconv.FormatBool(*pd.IsRequired)
				}

				if pd.IsUpdatable != nil {
					rtn[namespace+"."+strconv.Itoa(i)+"."+"is_updatable"] = strconv.FormatBool(*pd.IsUpdatable)
				}
			}
		}
	}
	return rtn
}
