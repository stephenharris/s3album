

resource "aws_s3_bucket" "photos" {
  bucket = var.bucket_name
  acl    = "private"

  cors_rule {
    allowed_headers = ["*"]
    allowed_methods = ["PUT", "GET", "POST", "DELETE"]
    allowed_origins = ["*"]
    expose_headers  = []
    max_age_seconds = 0
  }

  versioning {
    enabled    = true
    mfa_delete = false
  }
}