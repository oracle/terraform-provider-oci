// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package container_instances

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_container_instances_container_instance", ContainerInstancesContainerInstanceDataSource())
	tfresource.RegisterDatasource("oci_container_instances_container_instance_shape", ContainerInstancesContainerInstanceShapeDataSource())
	tfresource.RegisterDatasource("oci_container_instances_container_instance_shapes", ContainerInstancesContainerInstanceShapesDataSource())
	tfresource.RegisterDatasource("oci_container_instances_container_instances", ContainerInstancesContainerInstancesDataSource())
}
