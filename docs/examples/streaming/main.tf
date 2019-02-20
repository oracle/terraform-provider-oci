variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

resource "oci_streaming_stream" "stream" {
  compartment_id     = "${var.tenancy_ocid}"
  name               = "stream1"
  partitions         = "1"
  retention_in_hours = "24"
}

data "oci_streaming_stream" "stream1" {
  stream_id = "${oci_streaming_stream.stream.id}"
}

# Output the result
output "stream" {
  value = <<EOF

  id = "${data.oci_streaming_stream.stream1.id}"
  compartment_id = "${data.oci_streaming_stream.stream1.compartment_id}"
  messages_endpoint = "${data.oci_streaming_stream.stream1.messages_endpoint}"
  name = "${data.oci_streaming_stream.stream1.name}"
  partitions = "${data.oci_streaming_stream.stream1.partitions}"
  retention_in_hours = "${data.oci_streaming_stream.stream1.retention_in_hours}"
  state = "${data.oci_streaming_stream.stream1.state}"
  time_created = "${data.oci_streaming_stream.stream1.time_created}"
EOF

  # This value is not always present--when state is FAILED it may contain an explanation.
  #lifecycle_state_details = "${data.oci_streaming_stream.stream1.lifecycle_state_details}"
}

data "oci_streaming_streams" "streams" {
  compartment_id = "${oci_streaming_stream.stream.compartment_id}"

  # optional
  state = "ACTIVE"

  //  id    = "${oci_streaming_stream.stream.id}"
  //  name  = "${oci_streaming_stream.stream.name}"
}

output "streams" {
  value = "${data.oci_streaming_streams.streams.streams}"
}
