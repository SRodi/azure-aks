module "aks" {
  source              = "../../modules/aks"
  resource_group_name = var.resource_group_name
  location            = var.location
  prefix              = var.prefix
  labels              = var.labels
}