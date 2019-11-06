// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"
)

func ObjectStorageBucketDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["namespace"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(ObjectStorageBucketResource(), fieldMap, readSingularObjectStorageBucket)
}

func readSingularObjectStorageBucket(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageBucketDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return ReadResource(sync)
}

type ObjectStorageBucketDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_object_storage.ObjectStorageClient
	Res    *oci_object_storage.GetBucketResponse
}

func (s *ObjectStorageBucketDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ObjectStorageBucketDataSourceCrud) Get() error {
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

func (s *ObjectStorageBucketDataSourceCrud) SetData() error {
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

	if s.Res.ObjectEventsEnabled != nil {
		s.D.Set("object_events_enabled", *s.Res.ObjectEventsEnabled)
	}

	if s.Res.ObjectLifecyclePolicyEtag != nil {
		s.D.Set("object_lifecycle_policy_etag", *s.Res.ObjectLifecyclePolicyEtag)
	}

	s.D.Set("storage_tier", s.Res.StorageTier)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
