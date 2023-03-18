create table bundles
(
    id                    serial  not null primary key,
    symbol                varchar not null,
    exchange_from         varchar not null,
    exchange_to           varchar not null,
    percentage_difference decimal not null,
    constraint bundle_unique unique (symbol, exchange_from, exchange_to)
);


