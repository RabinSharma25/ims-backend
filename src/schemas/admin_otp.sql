CREATE TABLE voc.admin_otp(
    "id" bigserial primary key,
    email varchar(255) not null,
    otp bigint not null,
    created_at timestamp with time zone not null
);
