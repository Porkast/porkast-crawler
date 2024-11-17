-- Active: 1731851622041@@127.0.0.1@5432@porkastdb@public
CREATE TABLE feed_channel (
    id varchar(64) NOT NULL,
    title varchar(128),
    channel_desc text,
    image_url varchar(128),
    link varchar(128),
    feed_link varchar(128),
    copyright varchar(128),
    language varchar(128),
    author varchar(128),
    owner_name varchar(128),
    owner_email varchar(128),
    feed_type varchar(128),
    categories varchar(128),
    source varchar(64),
    feed_id varchar(64),
    PRIMARY KEY (id)
);

CREATE TABLE feed_item (
    id varchar(64) NOT NULL,
    channel_id varchar(64) NOT NULL,
    guid varchar(256),
    title text,
    link text,
    pub_date date,
    author varchar(128),
    input_date timestamp,
    image_url varchar(256),
    enclosure_url text,
    enclosure_type varchar(256),
    enclosure_length varchar(256),
    duration varchar(256),
    episode varchar(64),
    explicit varchar(64),
    season varchar(64),
    episodeType varchar(64),
    description text,
    channel_title text,
    feed_id varchar(64) NOT NULL,
    feed_link varchar(255),
    source varchar(255),
    PRIMARY KEY (id)
);

CREATE EXTENSION pgroonga;

CREATE INDEX rfi_idx_channel_id ON feed_item (channel_id);

CREATE INDEX idx_feed_item_full_text_search ON feed_item USING pgroonga (
    title,
    channel_title,
    description
);