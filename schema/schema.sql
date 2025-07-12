create table users (
    id varchar(26) primary key,
    first_name varchar(128) not null,
    last_name varchar(128) not null,
    email varchar(320) not null,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp
);

create table auth (
    user_id varchar(26) primary key references users(id),
    password_hash bytea not null,
    salt bytea not null,
    created_at timestamptz not null default current_timestamp
);

create table token (
    token varchar(2048) not null primary key,
    user_id varchar(26) not null references users(id),
    revoked bool not null,
    created_at timestamptz not null default current_timestamp
);

create table product (
    id varchar(26) primary key,
    name varchar(256) not null,
    description varchar(512) not null,
    image_url varchar(1024) not null,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp
);

create table product_unit (
    id varchar(26) primary key,
    product_id varchar(26) references product(id),
    unit varchar(50) not null,
    default_unit bool not null default false
);

create table location (
    id varchar(26) primary key,
    name varchar(256) not null,
    description varchar(256) not null,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp
);

create table inventory_tx (
    id varchar(26) primary key,
    user_id varchar(26) references users(id),
    product_id varchar(26) references product(id),
    unit_id varchar(26) references product_unit(id),
    location_id varchar(26) references location(id),
    action varchar(50) not null,
    quantity numeric(10) not null,
    occurred_at timestamptz not null default current_timestamp
);