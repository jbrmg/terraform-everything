resource ikea_kitchen k {
  name = "my-kitchen"
}

resource ikea_cabinet c1 {
  front      = "RINGHULT"
  color      = "#FFFFFF"
  kitchen_id = ikea_kitchen.k.id
}

resource ikea_cabinet c2 {
  front      = "RINGHULT"
  color      = "#000000"
  kitchen_id = ikea_kitchen.k.id
}

resource ikea_countertop top {
  type        = "KARLBY"
  cabinet_ids = [ikea_cabinet.c1.id, ikea_cabinet.c2.id]
}