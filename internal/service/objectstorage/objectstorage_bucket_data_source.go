// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package objectstorage

import (
	"context"
	"strconv"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/v58/objectstorage"
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
	return tfresource.GetSingularDataSourceItemSchema(ObjectStorageBucketResource(), fieldMap, readSingularObjectStorageBucket)
}

func readSingularObjectStorageBucket(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageBucketDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()

	return tfresource.ReadResource(sync)
}

type ObjectStorageBucketDataSourceCrud struct {
	D                      *schema.ResourceData
	Client                 *oci_object_storage.ObjectStorageClient
	Res                    *oci_object_storage.GetBucketResponse
	RetentionRuleRes       []*oci_object_storage.RetentionRule
	DisableNotFoundRetries bool
}

func (s *ObjectStorageBucketDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ObjectStorageBucketDataSourceCrud) Get() error {
	request := oci_object_storage.GetBucketRequest{}
	listRetentionRulesRequest := oci_object_storage.ListRetentionRulesRequest{}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.BucketName = &tmp
		listRetentionRulesRequest.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
		listRetentionRulesRequest.NamespaceName = &tmp
	}

	request.Fields = oci_object_storage.GetGetBucketFieldsEnumValues()
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "object_storage")

	response, err := s.Client.GetBucket(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response

	// using list call as summary and get response is same for a retention rule
	listRetentionRulesRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")
	listRetentionRulesResponse, e := s.Client.ListRetentionRules(context.Background(), listRetentionRulesRequest)
	if e != nil {
		return e
	}

	s.RetentionRuleRes = listResponseToRetentionRuleRes(listRetentionRulesResponse)

	return nil
}

func (s *ObjectStorageBucketDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceID())

	s.D.Set("access_type", s.Res.PublicAccessType)

	if s.Res.ApproximateCount != nil {
		s.D.Set("approximate_count", strconv.FormatInt(*s.Res.ApproximateCount, 10))
	}

	if s.Res.ApproximateSize != nil {
		s.D.Set("approximate_size", strconv.FormatInt(*s.Res.ApproximateSize, 10))
	}

	s.D.Set("auto_tiering", s.Res.AutoTiering)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Etag != nil {
		s.D.Set("etag", *s.Res.Etag)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsReadOnly != nil {
		s.D.Set("is_read_only", *s.Res.IsReadOnly)
	}

	if s.Res.Id != nil {
		s.D.Set("bucket_id", *s.Res.Id)
	}

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

	if s.Res.ReplicationEnabled != nil {
		s.D.Set("replication_enabled", *s.Res.ReplicationEnabled)
	}

	s.D.Set("storage_tier", s.Res.StorageTier)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	s.D.Set("retention_rules", retentionRulesResToSet(s.RetentionRuleRes, true))

	s.D.Set("versioning", s.Res.Versioning)

	return nil
}
