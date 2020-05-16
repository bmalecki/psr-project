provider "google" {
  project     = var.gcp_project
  region      = "us-central1"
}

resource "google_storage_bucket" "static-site" {
  name          = "helloooosdsfewferuehfb"
  location      = "EU"
  force_destroy = true

  bucket_policy_only = true

  website {
    main_page_suffix = "index.html"
    not_found_page   = "404.html"
  }
}