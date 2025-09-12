create table if not exists public.carts (
    id serial primary key,
    name varchar(50) not null
);

create table if not exists public.items (
    id serial primary key,
    name varchar(50) not null,
    cart_id int not null references public.carts(id)
);