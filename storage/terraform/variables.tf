variable "region" {
    type = string
    default = "eu-west-1"
}

variable "bucket_name" {
    type = string
}

variable "src_object_prefix" {
    type = string
    default = ""
}