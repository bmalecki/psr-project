output "static_website_endpoint_url" {
  value       = "http://${aws_s3_bucket.static_website.website_endpoint}"
  description = "The static website endpoint"
}

output "static_website_id" {
  value       =  aws_s3_bucket.static_website.id
  description = "The static website id"
}