-- name: CreateSale :exec
INSERT INTO sales (
    id,
    ticket_id,
    user_id,
    payment_type
) 
VALUES (
    ?, ?, ?, ?
);

-- name: GetSaleByID :one
SELECT
    *
FROM
    sales
WHERE
    id = ?;