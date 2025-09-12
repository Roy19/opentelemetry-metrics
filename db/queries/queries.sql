-- name: GetCartDetails :one
select id, name from public.carts;

-- name: AddItemToCart :exec
insert into public.items (name, cart_id) values ($1, $2);

-- name: GetItemsInCart :many
select it.id, it.name from public.items it
inner join public.carts c
  on it.cart_id = c.id
where it.cart_id = $1;