provider "google" {
  project     = var.gcp_project
  region      = "us-central1"
  credentials = "secrets/gcp.json"
}

resource "google_storage_bucket" "example" {
  name          = "${random_uuid.storage.keepers.storage}-${random_uuid.storage.result}"
  location      = "EU"
  force_destroy = true

  bucket_policy_only = true
}
