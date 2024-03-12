CREATE TABLE public."shortener" (
    "id"           varchar(40) NOT NULL,
    "short_url"    varchar(50) NOT NULL,
    "original_url" varchar(50) NOT NULL,
    "created"      timestamp(0) NOT NULL
);
