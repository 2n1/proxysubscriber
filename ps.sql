CREATE TABLE groups
(
    id   INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name TEXT    NOT NULL UNIQUE,
    url  TEXT    NOT NULL UNIQUE
);

CREATE TABLE nodes
(
    id        INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name      TEXT    NOT NULL,
    group_id  INTEGER NOT NULL,
    node_type TEXT    NOT NULL,
    server    TEXT    NOT NULL,
    port      INTEGER NOT NULL,
    passwd    TEXT    NOT NULL,
    cipher    TEXT    NOT NULL,
    sni       TEXT    NOT NULL,
    alter_id  TEXT    NOT NULL,
    ws_path   TEXT    NOT NULL,
    ws_host   TEXT    NOT NULL,
    cf_ip     INTEGER NOT NULL
);

CREATE TABLE cfips
(
    cu_ip    TEXT NOT NULL,
    cu_label TEXT NOT NULL,
    ct_ip    TEXT NOT NULL,
    ct_label TEXT NOT NULL,
    cm_ip    TEXT NOT NULL,
    cm_label TEXT NOT NULL
);

CREATE TABLE auths
(
    id     INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    email  TEXT    NOT NULL UNIQUE,
    passwd TEXT    NOT NULL
);