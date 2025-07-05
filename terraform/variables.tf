variable "project_id" {
  description = "Google Cloud Project ID"
  type        = string
}

variable "region" {
  description = "Google Cloud Region"
  type        = string
  default     = "asia-northeast1"
}

variable "bucket_name" {
  description = "GCS bucket name for Terraform state"
  type        = string
  default     = "my-bucket"
}