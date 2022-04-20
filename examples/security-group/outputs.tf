output "arn" {
  description = "The ARN for the EFS"
  value       = module.efs.arn
}

output "id" {
  description = "The ID for the EFS"
  value       = module.efs.arn
}

output "dns_name" {
  description = "The DNS name of the EFS"
  value       = module.efs.dns_name
}
