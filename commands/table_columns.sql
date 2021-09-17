select ordinal_position
     , column_name
     , data_type
     , is_nullable
  from information_schema.columns
 where table_schema = $1
   and table_name   = $2
 order by ordinal_position
