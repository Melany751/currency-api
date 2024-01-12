create table currencies
(
    id          uuid             not null
        primary key,
    code        varchar(10)      not null,
    value       double precision not null,
    create_date timestamp        not null
);

create table registries
(
    transaction varchar,
    date        timestamp,
    duration    double precision
);