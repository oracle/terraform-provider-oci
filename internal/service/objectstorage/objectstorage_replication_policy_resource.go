// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package objectstorage

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_object_storage "github.com/oracle/oci-go-sdk/v58/objectstorage"
)

func ObjectStorageReplicationPolicyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createObjectStorageReplicationPolicy,
		Read:     readObjectStorageReplicationPolicy,
		Delete:   deleteObjectStorageReplicationPolicy,
		Schema: map[string]*schema.Schema{
			// Required
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"destination_bucket_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"destination_region_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"delete_object_in_destination_bucket": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"ACCEPT",
				}, true),
				Deprecated: tfresource.FieldDeprecated("delete_object_in_destination_bucket"),
			},

			// Optional

			// Computed
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status_message": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_sync": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createObjectStorageReplicationPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageReplicationPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()

	return tfresource.CreateResource(d, sync)
}

func readObjectStorageReplicationPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageReplicationPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()

	return tfresource.ReadResource(sync)
}

func deleteObjectStorageReplicationPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageReplicationPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ObjectStorageReplicationPolicyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_object_storage.ObjectStorageClient
	Res                    *oci_object_storage.ReplicationPolicy
	DisableNotFoundRetries bool
}

func (s *ObjectStorageReplicationPolicyResourceCrud) ID() string {
	return GetReplicationPolicyCompositeId(s.D.Get("bucket").(string), s.D.Get("namespace").(string), *s.Res.Id)
}

func (s *ObjectStorageReplicationPolicyResourceCrud) Create() error {
	request := oci_object_storage.CreateReplicationPolicyRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if destinationBucketName, ok := s.D.GetOkExists("destination_bucket_name"); ok {
		tmp := destinationBucketName.(string)
		request.DestinationBucketName = &tmp
	}

	if destinationRegionName, ok := s.D.GetOkExists("destination_region_name"); ok {
		tmp := destinationRegionName.(string)
		request.DestinationRegionName = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.CreateReplicationPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ReplicationPolicy
	return nil
}

func (s *ObjectStorageReplicationPolicyResourceCrud) Get() error {
	request := oci_object_storage.GetReplicationPolicyRequest{}

	bucket, namespace, replicationId, err := ParseReplicationPolicyCompositeId(s.D.Id())
	if err == nil {
		request.BucketName = &bucket
		request.NamespaceName = &namespace
		request.ReplicationId = &replicationId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.GetReplicationPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ReplicationPolicy
	return nil
}

func (s *ObjectStorageReplicationPolicyResourceCrud) Delete() error {
	request := oci_object_storage.DeleteReplicationPolicyRequest{}

	bucket, namespace, replicationId, err := ParseReplicationPolicyCompositeId(s.D.Id())
	if err == nil {
		request.BucketName = &bucket
		request.NamespaceName = &namespace
		request.ReplicationId = &replicationId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	_, err = s.Client.DeleteReplicationPolicy(context.Background(), request)
	return err
}

func (s *ObjectStorageReplicationPolicyResourceCrud) SetData() error {

	bucket, namespace, replicationId, err := ParseReplicationPolicyCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("bucket", &bucket)
		s.D.Set("namespace", &namespace)
		s.D.Set("replication_id", &replicationId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.DestinationBucketName != nil {
		s.D.Set("destination_bucket_name", *s.Res.DestinationBucketName)
	}

	if s.Res.DestinationRegionName != nil {
		s.D.Set("destination_region_name", *s.Res.DestinationRegionName)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("status", s.Res.Status)

	if s.Res.StatusMessage != nil {
		s.D.Set("status_message", *s.Res.StatusMessage)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastSync != nil {
		s.D.Set("time_last_sync", s.Res.TimeLastSync.String())
	}

	return nil
}

func GetReplicationPolicyCompositeId(bucket string, namespace string, replicationId string) string {
	bucket = url.PathEscape(bucket)
	namespace = url.PathEscape(namespace)
	replicationId = url.PathEscape(replicationId)
	compositeId := "n/" + namespace + "/b/" + bucket + "/replicationPolicies/" + replicationId
	return compositeId
}

func ParseReplicationPolicyCompositeId(compositeId string) (bucket string, namespace string, replicationId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("n/.*/b/.*/replicationPolicies/.*", compositeId)
	if !match || len(parts) != 6 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	namespace, _ = url.PathUnescape(parts[1])
	bucket, _ = url.PathUnescape(parts[3])
	replicationId, _ = url.PathUnescape(parts[5])

	return
}
