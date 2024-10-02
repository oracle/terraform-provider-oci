// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package security_attribute

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_security_attribute_security_attribute", SecurityAttributeSecurityAttributeDataSource())
	tfresource.RegisterDatasource("oci_security_attribute_security_attribute_namespace", SecurityAttributeSecurityAttributeNamespaceDataSource())
	tfresource.RegisterDatasource("oci_security_attribute_security_attribute_namespaces", SecurityAttributeSecurityAttributeNamespacesDataSource())
	tfresource.RegisterDatasource("oci_security_attribute_security_attributes", SecurityAttributeSecurityAttributesDataSource())
}
