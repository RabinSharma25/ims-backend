CREATE TABLE voc.admin_uuid(
    "id" bigserial primary key,
    email varchar(255) not null,
    uuid varchar(255) not null,
    created_at timestamp with time zone not null
);