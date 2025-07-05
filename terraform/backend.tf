terraform {
  backend "gcs" {
    bucket = "my-bucket"
    prefix = "catbot/prod/googlecloud"
  }
}