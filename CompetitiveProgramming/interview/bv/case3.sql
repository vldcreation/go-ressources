WITH quantumsafe_detections_summary AS (
    SELECT
        ft.extension,
        COUNT(qd.filetype_id) AS quantumsafe_total_detections
    FROM
        file_types ft
    LEFT JOIN
        quantumsafe_detections qd ON ft.id = qd.filetype_id
    WHERE
        qd.dt BETWEEN '2023-06-31' AND '2023-08-01'
    GROUP BY
        ft.extension
),
webguardian_detections_summary AS (
    SELECT
        ft.extension,
        COUNT(wgd.filetype_id) AS webguardian_total_detections
    FROM
        file_types ft
    LEFT JOIN
        webguardian_detections wgd ON ft.id = wgd.filetype_id
    WHERE
        wgd.dt BETWEEN '2023-06-31' AND '2023-08-01'
    GROUP BY
        ft.extension
)
SELECT
    qs.extension,
    qs.quantumsafe_total_detections,
    wg.webguardian_total_detections
FROM
    quantumsafe_detections_summary qs
LEFT JOIN
    webguardian_detections_summary wg ON qs.extension = wg.extension
ORDER BY
    qs.extension ASC;