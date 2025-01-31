variable "gcp_project_id" {
  description = "The project to deploy the cluster in"
  type        = string
}

variable "gcp_region" {
  type = string
}

variable "custom_envs_repository_name" {
  type = string
}

variable "orchestration_repository_name" {
  type = string
}

variable "prefix" {
  type        = string
  description = "The prefix to use for all resources in this module"
  default     = "khulnasoft-"
}

variable "labels" {
  description = "The labels to attach to resources created by this module"
  type        = map(string)
  default = {
    "app"       = "khulnasoft"
    "terraform" = "true"
  }
}
