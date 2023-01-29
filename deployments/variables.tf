variable "region" {
  description = "AWS region"
  type        = string
}

variable "vpc_name" {
  description = "Name of the VPC"
  type        = string
}

variable "vpc_cidr" {
  description = "VPC cidr"
  type        = string
  default     = "10.0.0.0/16"
}

variable "private_subnets" {
  description = "Private Subnets cidr"
  type        = list(string)
  default     = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
}

variable "public_subnets" {
  description = "Public Subnets cidr"
  type        = list(string)
  default     = ["10.0.4.0/24", "10.0.5.0/24", "10.0.6.0/24"]
}

variable "cluster_name" {
  description = "EKS Cluster name"
  type        = string
}

variable "cluster_version" {
  description = "EKS Cluster version"
  type        = string
}

variable "instance_types" {
  description = "Nodes instance types"
  type        = list(string)
  default     = ["t2.micro"]
}

variable "ami_type" {
  description = "AMI type"
  type        = string
  default     = "AL2_x86_64"
}