// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_core_virtual_circuit_public_prefixes" "virtual_circuit_public_prefixes" {
  #Required
  virtual_circuit_id = oci_core_virtual_circuit.virtual_circuit_public.id

  #Optional
  verification_state = var.virtual_circuit_public_prefix_verification_state
}

output "virtual_circuit_public_prefixes" {
  value = data.oci_core_virtual_circuit_public_prefixes.virtual_circuit_public_prefixes.virtual_circuit_public_prefixes
}

