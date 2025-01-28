resource "azurerm_resource_group" "aks_rg" {
  name     = var.resource_group_name
  location = var.location
}

resource "azurerm_kubernetes_cluster" "aks" {
  name                = "${var.prefix}-aks"
  location            = azurerm_resource_group.aks_rg.location
  resource_group_name = azurerm_resource_group.aks_rg.name
  dns_prefix          = "${var.prefix}-aks-dns"
  kubernetes_version  = "1.29.8"

  default_node_pool {
    name        = "default"
    node_count  = 1
    vm_size     = "Standard_D2_v2"
    node_labels = var.labels
  }

  identity {
    type = "SystemAssigned"
  }

  tags = var.labels
}