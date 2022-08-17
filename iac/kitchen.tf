resource ikea_cabinet c1 {
  front = "RINGHULT"
  color = "#FFFFFF"
}

resource ikea_cabinet c2 {
  front = "RINGHULT"
  color = "#000000"
}

resource ikea_countertop top {
  type        = "KARLBY"
  cabinet_ids = [ikea_cabinet.c1.id, ikea_cabinet.c2.id]
}