-- +goose Up
INSERT INTO urls(url) VALUES
  ("https://chaster-club.ru/catalog/product/orient-ra-ac0m11y"),
  ("https://www.ozon.ru/product/roland-fp-30x-wh-tsifrovoe-pianino-beloe-1280857545/?avtc=1&avte=4&avts=1738784206"),
  ("https://www.dns-shop.ru/product/9c399838380bed20/videokarta-asus-geforce-rtx-4070-dual-white-oc-edition-dual-rtx4070-o12g-white/");

-- +goose Down
DELETE FROM urls WHERE TRUE;
