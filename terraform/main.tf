resource "random_uuid" "static_website" {
  keepers = {
    # Generate a new id each time we switch to a new AMI id
    static_website = "${var.static_website}"
  }
}


