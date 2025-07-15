-- name: UpdateTicketQuantity :exec
UPDATE
    tickets
SET
    quantity = ?
WHERE
    id = ?;

-- name: GetAllTickets :many
SELECT
    *
FROM
    tickets;

-- name: GetTicketByID :one
SELECT
    *
FROM
    tickets
WHERE
    id = ?
LIMIT
    1;
