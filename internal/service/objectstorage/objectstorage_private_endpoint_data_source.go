// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package objectstorage

import (
	"context"
	"fmt"
	"log"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/v65/objectstorage"
)

func ObjectStoragePrivateEndpointDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["namespace"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ObjectStoragePrivateEndpointResource(), fieldMap, readSingularObjectStoragePrivateEndpoint)
}

func readSingularObjectStoragePrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStoragePrivateEndpointDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()

	return tfresource.ReadResource(sync)
}

type ObjectStoragePrivateEndpointDataSourceCrud struct {
	D                      *schema.ResourceData
	Client                 *oci_object_storage.ObjectStorageClient
	Res                    *oci_object_storage.GetPrivateEndpointResponse
	DisableNotFoundRetries bool
}

func (s *ObjectStoragePrivateEndpointDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ObjectStoragePrivateEndpointDataSourceCrud) Get() error {
	request := oci_object_storage.GetPrivateEndpointRequest{}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.PeName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "object_storage")

	response, err := s.Client.GetPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response

	return nil
}

func (s *ObjectStoragePrivateEndpointDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceID())

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeModified != nil {
		s.D.Set("time_modified", s.Res.TimeModified.String())
	}

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.PrivateEndpointIp != nil {
		s.D.Set("private_endpoint_ip", s.Res.PrivateEndpointIp)
	}

	if s.Res.Prefix != nil {
		s.D.Set("prefix", s.Res.Prefix)
	}

	if s.Res.ETag != nil {
		s.D.Set("etag", *s.Res.ETag)
	}

	if s.Res.AccessTargets != nil {
		var accessTargets []map[string]interface{}
		for _, target := range s.Res.AccessTargets {
			accessTarget := map[string]interface{}{
				"namespace":      target.Namespace,
				"compartment_id": target.CompartmentId,
				"bucket":         target.Bucket,
				// Add other fields as needed
			}
			accessTargets = append(accessTargets, accessTarget)
		}
		s.D.Set("access_targets", accessTargets)
	} else {
		log.Printf("[WARN] SetData() unable to parse accessTargets : %s", s.Res.AccessTargets)
	}

	if s.Res.Fqdns != nil {
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
		// Set fqdnsMap in ResourceData
		if err := s.D.Set("fqdns", fqdnsMap); err != nil {
			return fmt.Errorf("error setting fqdns attribute: %w", err)
		}
	} else {
		return fmt.Errorf("s.Res.Fqdns is nil")
	}

	s.D.Set("additional_prefixes", s.Res.AdditionalPrefixes)

	s.D.Set("nsg_ids", s.Res.NsgIds)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("id", *s.Res.Id)

	return nil
}
