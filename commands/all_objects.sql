select nsp.nspname || '.' || cls.relname as object_name
     , case cls.relkind
         when 'r' then 'TABLE'
         when 'm' then 'MATERIALIZED_VIEW'
         when 'i' then 'INDEX'
         when 'S' then 'SEQUENCE'
         when 'v' then 'VIEW'
         when 'c' then 'TYPE'
         else cls.relkind::text
       end as object_type
  from pg_class cls
  left join pg_namespace nsp on nsp.oid = cls.relnamespace
 where nsp.nspname not in ('information_schema', 'pg_catalog')
   and nsp.nspname not like 'pg_toast%'
   and cls.relkind in ('r', 'v')
 order by
       nsp.nspname
     , cls.relname
