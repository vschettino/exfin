create table brokers
(
    id          serial  not null
        constraint bank_account_pk
            primary key,
    name        varchar not null,
    institution varchar not null,
    account_id  integer not null
        constraint bank_accounts_accounts_id_fk
            references public.accounts
            on update cascade on delete cascade
);

create table companies
(
    id     serial  not null
        constraint company_pk
            primary key,
    name   varchar not null,
    sector varchar
);


create table assets
(
    id             serial      not null
        constraint assets_pk
            primary key,
    code           varchar(10) not null,
    type           varchar(30),
    classification varchar(5)  not null,
    company_id     integer     not null
        constraint assets_companies_id_fk
            references companies
            on update cascade on delete cascade
);


create unique index asset_code_uindex
    on assets (code);

create table operations
(
    id           serial               not null
        constraint operations_pk
            primary key,
    asset_id     integer              not null
        constraint operations_assets_id_fk
            references assets
            on update cascade on delete cascade,
    broker_id    integer
        constraint operations_brokers_id_fk
            references brokers
            on update cascade on delete cascade,
    unit_price   integer,
    market_taxes integer    default 0 not null,
    other_taxes  integer    default 0 not null,
    total        integer    default 0 not null,
    type         varchar              not null,
    quantity     integer    default 0 not null,
    currency     varchar(3) default 'BRL'::character varying
);


create table yields
(
    id             serial            not null
        constraint yields_pk
            primary key,
    asset_id       integer
        constraint yields_assets_id_fk
            references assets
            on update cascade on delete cascade,
    account_id     integer
        constraint yields_accounts_id_fk
            references public.accounts
            on update cascade on delete cascade,
    broker_id      integer
        constraint yields_brokers_id_fk
            references brokers
            on update cascade on delete cascade,
    ref_start_date date,
    ref_end_date   date,
    total          integer default 0 not null,
    asset_quantity integer default 0 not null,
    grant_date     date              not null
);

