    CREATE OR REPLACE FUNCTION uniq(int[]) RETURNS int[] AS $BODY$
        SELECT array( SELECT unnest($1) group by 1);  
    $BODY$ LANGUAGE sql strict;  

    CREATE OR REPLACE FUNCTION array_uniq_cat(anyarray,anyarray) RETURNS anyarray AS $BODY$ 
        SELECT uniq(array_cat($1,$2));   
    $BODY$  LANGUAGE sql strict;

    DROP AGGREGATE  IF EXISTS  arragg (anyarray);
    CREATE  AGGREGATE  arragg (anyarray) (sfunc = array_uniq_cat, stype=anyarray, PARALLEL=safe);
