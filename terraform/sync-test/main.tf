resource "random_pet" "p" {
  keepers = {
    x = "i"
  }
}

output "pet" {
  value = random_pet.p.id
}
