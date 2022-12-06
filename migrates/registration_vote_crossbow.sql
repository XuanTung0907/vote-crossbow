create table registration_vote_crossbow
(
    created_at   timestamp with time zone,
    created_by   bigint,
    updated_at   timestamp with time zone,
    updated_by   bigint,
    deleted_at   timestamp with time zone,
    id           bigserial primary key,
    name         varchar(255),
    phone_number varchar(20),
    email        varchar(255),
    content      text
)
