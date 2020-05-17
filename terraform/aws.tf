provider "aws" {
  shared_credentials_file = "secrets/aws.credentials"
  version                 = "~> 2.0"
  region                  = "us-east-1"
}

resource "aws_s3_bucket" "b" {
  bucket = "${random_uuid.storage.keepers.storage}-${random_uuid.storage.result}"
  acl    = "private"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}

resource "aws_s3_bucket" "static_website" {
  bucket = "static-website-${random_uuid.static_website.result}"
  acl    = "public-read"

  website {
    index_document = "index.html"
    error_document = "error.html"
  }
}

resource "aws_s3_bucket_policy" "static_website_policy" {
  bucket = aws_s3_bucket.static_website.id

  policy = <<POLICY
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "PublicReadGetObject",
            "Effect": "Allow",
            "Principal": "*",
            "Action": [
                "s3:GetObject"
            ],
            "Resource": [
                "arn:aws:s3:::${aws_s3_bucket.static_website.id}/*"
            ]
        }
    ]
}
POLICY
}
