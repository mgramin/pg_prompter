select d.datname as name
     , pg_catalog.pg_get_userbyid(d.datdba) as owner
     , pg_encoding_to_char(encoding) as encoding
     , datcollate as "collate"
     , datctype as ctype
     , null as privileges
  from pg_database d
 order by 1
