terraform {
  backend "gcs" {
    bucket = "terraform-remote-backend-e472"
    prefix = "catbot/prod/googlecloud"
  }
}