CREATE OR REPLACE FUNCTION "public"."click_activity"("_link_id" uuid)
RETURNS TABLE("count" int8, "date" date) AS $BODY$
SELECT
COUNT(s.*),
d. DAY

FROM
(
  SELECT
  DAY :: DATE
  FROM
  generate_series(
    (
      SELECT
      created_at
      FROM
      links
      WHERE
      id = $1
    ) :: DATE,
    CURRENT_DATE :: DATE
    , INTERVAL '1 day'
  ) DAY
  ) d LEFT JOIN clicks s ON(
  s.created_at :: DATE = d.DAY AND s.link_id = $1
) GROUP BY d.DAY ORDER BY d.DAY
$BODY$
LANGUAGE sql VOLATILE
COST 100
ROWS 1000
