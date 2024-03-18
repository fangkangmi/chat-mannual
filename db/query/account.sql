-- name: CreateAuthor :execresult
INSERT INTO accounts (
    owner,
    balance,
    currency,
    username,
    password
) VALUES (
  ?, ?, ?, ?, ?
);