CREATE TABLE offer_categories
(
    id         uuid      NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP NOT NULL                    DEFAULT current_timestamp,
    code       varchar   not null unique
);

CREATE TABLE offers
(
    id                 uuid      NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
    user_id            uuid      NOT NULL REFERENCES users ON DELETE CASCADE,
    category_id        uuid      NOT NULL REFERENCES offer_categories ON DELETE NO ACTION,
    created_at         TIMESTAMP NOT NULL                    DEFAULT current_timestamp,
    hourly_price_fiat  money,
    hours_price_tokens money,
    published_state    varchar   NOT NULL                    DEFAULT 'draft'
);

CREATE TABLE offer_descriptions
(
    id          uuid NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
    language    varchar,
    offer_id    uuid NOT NULL REFERENCES offers ON DELETE CASCADE,
    title       varchar,
    description varchar,
    unique (offer_id, language)
)