select ordinal_position as position
     , column_name as column
     , case
         when character_maximum_length is not null then data_type || '(' || character_maximum_length || ')'
         else data_type
       end as type
     , character_maximum_length as length
     , case
         when is_nullable = 'NO' then 'not null'
         else ''
       end as nullable
  from information_schema.columns
 where table_schema = ?
   and table_name   = ?
 order by ordinal_position
