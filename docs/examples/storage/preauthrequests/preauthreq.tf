resource "baremetal_objectstorage_preauthrequest" "parOnBucket" {
  namespace = "${var.namespace_name}"
  bucket = "${var.bucket_name}"
  name = "parOnBucket"
  access_type = "AnyObjectWrite" //Other configurations accepted are ObjectWrite, ObjectReadWrite
  time_expires = "2019-11-10T23:00:00Z"
}

resource "baremetal_objectstorage_preauthrequest" "parOnObject" {
  namespace = "${var.namespace_name}"
  bucket = "${var.bucket_name}"
  object = "${var.object_name}"
  name = "parOnObject"
  access_type = "ObjectRead" //Other configurations accepted are ObjectWrite, ObjectReadWrite
  time_expires = "2019-11-10T23:00:00Z"
}
