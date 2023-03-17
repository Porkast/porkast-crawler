create table
    if not exists rss_feed_channel (
        id varchar(64) not null primary key,
        title varchar(128) null,
        channel_desc mediumtext null,
        image_url varchar(128) null,
        link varchar(128) null,
        rss_link varchar(128) null,
        copyright varchar(128) null,
        language varchar(128) null,
        author varchar(128) null,
        owner_name varchar(128) null,
        owner_email varchar(128) null,
        feed_type varchar(128) null,
        categories varchar(128) null,
    );

-- MySQL fulltext search table

create table
    if not exists rss_feed_item (
        id varchar(64) not null primary key,
        channel_id varchar(64) not null,
        title mediumtext null,
        link varchar(128) null,
        pub_date date null,
        author varchar(128) null,
        input_date datetime null,
        image_url varchar(256) null,
        enclosure_url varchar(256) null,
        enclosure_type varchar(256) null,
        enclosure_length varchar(256) null,
        duration varchar(256) null,
        episode varchar(64) null,
        explicit varchar(64) null,
        season varchar(64) null,
        episodeType varchar(64) null,
        description mediumtext null,
    );

create index rfi_idx_channel_id on rss_feed_item (channel_id);