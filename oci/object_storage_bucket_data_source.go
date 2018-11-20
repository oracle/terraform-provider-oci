// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"
)

func BucketDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularBucket,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"access_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"approximate_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"approximate_size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metadata": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"object_lifecycle_policy_etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_tier": {
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

func readSingularBucket(d *schema.ResourceData, m interface{}) error {
	sync := &BucketDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return ReadResource(sync)
}

type BucketDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_object_storage.ObjectStorageClient
	Res    *oci_object_storage.GetBucketResponse
}

func (s *BucketDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BucketDataSourceCrud) Get() error {
	request := oci_object_storage.GetBucketRequest{}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.Fields = oci_object_storage.GetGetBucketFieldsEnumValues()
	request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "object_storage")

	response, err := s.Client.GetBucket(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BucketDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	s.D.Set("access_type", s.Res.PublicAccessType)

	if s.Res.ApproximateCount != nil {
		s.D.Set("approximate_count", strconv.FormatInt(*s.Res.ApproximateCount, 10))
	}

	if s.Res.ApproximateSize != nil {
		s.D.Set("approximate_size", strconv.FormatInt(*s.Res.ApproximateSize, 10))
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Etag != nil {
		s.D.Set("etag", *s.Res.Etag)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	s.D.Set("metadata", s.Res.Metadata)

	if s.Res.ObjectLifecyclePolicyEtag != nil {
		s.D.Set("object_lifecycle_policy_etag", *s.Res.ObjectLifecyclePolicyEtag)
	}

	s.D.Set("storage_tier", s.Res.StorageTier)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
