// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package objectstorage

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func BucketResource() *schema.Resource {
	return &schema.Resource{
		Create: createBucket,
		Read:   readBucket,
		Update: updateBucket,
		Delete: deleteBucket,
		Schema: bucketSchema,
	}
}

func createBucket(d *schema.ResourceData, m interface{}) (e error) {
	sync := &BucketResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.CreateResource(d, sync)
}

func readBucket(d *schema.ResourceData, m interface{}) (e error) {
	sync := &BucketResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.ReadResource(sync)
}

func updateBucket(d *schema.ResourceData, m interface{}) (e error) {
	sync := &BucketResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.UpdateResource(d, sync)
}

func deleteBucket(d *schema.ResourceData, m interface{}) (e error) {
	sync := &BucketResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.DeleteResource(sync)
}
