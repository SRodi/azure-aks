output "ca_certificate" {
  value     = module.aks.ca_certificate
  sensitive = true
}

output "client_certificate" {
  value     = module.aks.client_certificate
  sensitive = true
}

output "client_key" {
  value     = module.aks.client_key
  sensitive = true
}

output "host" {
  value     = module.aks.host
  sensitive = true
}