resource "random_uuid" "storage" {
  keepers = {
    # Generate a new id each time we switch to a new AMI id
    storage = "${var.storage}"
  }
}


