-- name: CreateUser :one
INSERT INTO users (id, first_name, last_name, email, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY created_at DESC;

-- name: UpdateUser :one
UPDATE users
SET first_name = $2, last_name = $3, email = $4, updated_at = $5
WHERE id = $1
    RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- ============================================================================
-- AUTH TABLE QUERIES
-- ============================================================================

-- name: CreateAuth :one
INSERT INTO auth (user_id, password_hash, salt, created_at)
VALUES ($1, $2, $3, $4)
    RETURNING *;

-- name: GetAuth :one
SELECT * FROM auth WHERE user_id = $1;

-- name: GetAuthByUserEmail :one
SELECT a.* FROM auth a
                    JOIN users u ON a.user_id = u.id
WHERE u.email = $1;

-- name: ListAuth :many
SELECT * FROM auth ORDER BY created_at DESC;

-- name: UpdateAuth :one
UPDATE auth
SET password_hash = $2, salt = $3
WHERE user_id = $1
    RETURNING *;

-- name: DeleteAuth :exec
DELETE FROM auth WHERE user_id = $1;

-- ============================================================================
-- PRODUCT TABLE QUERIES
-- ============================================================================

-- name: CreateProduct :one
INSERT INTO product (id, name, description, image_url, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING *;

-- name: GetProduct :one
SELECT * FROM product WHERE id = $1;

-- name: GetProductByName :one
SELECT * FROM product WHERE name = $1;

-- name: ListProducts :many
SELECT * FROM product ORDER BY created_at DESC;

-- name: SearchProducts :many
SELECT * FROM product
WHERE name ILIKE $1 OR description ILIKE $1
ORDER BY created_at DESC;

-- name: UpdateProduct :one
UPDATE product
SET name = $2, description = $3, image_url = $4, updated_at = $5
WHERE id = $1
    RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM product WHERE id = $1;

-- ============================================================================
-- PRODUCT_UNIT TABLE QUERIES
-- ============================================================================

-- name: CreateProductUnit :one
INSERT INTO product_unit (id, product_id, unit, default_unit)
VALUES ($1, $2, $3, $4)
    RETURNING *;

-- name: GetProductUnit :one
SELECT * FROM product_unit WHERE id = $1;

-- name: GetProductUnitsByProductId :many
SELECT * FROM product_unit WHERE product_id = $1 ORDER BY default_unit DESC, unit ASC;

-- name: GetDefaultProductUnit :one
SELECT * FROM product_unit WHERE product_id = $1 AND default_unit = true;

-- name: ListProductUnits :many
SELECT * FROM product_unit ORDER BY product_id, default_unit DESC, unit ASC;

-- name: UpdateProductUnit :one
UPDATE product_unit
SET unit = $2, default_unit = $3
WHERE id = $1
    RETURNING *;

-- name: SetProductUnitAsDefault :exec
UPDATE product_unit
SET default_unit = CASE WHEN id = $2 THEN true ELSE false END
WHERE product_id = $1;

-- name: DeleteProductUnit :exec
DELETE FROM product_unit WHERE id = $1;

-- ============================================================================
-- LOCATION TABLE QUERIES
-- ============================================================================

-- name: CreateLocation :one
INSERT INTO location (id, name, description, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
    RETURNING *;

-- name: GetLocation :one
SELECT * FROM location WHERE id = $1;

-- name: GetLocationByName :one
SELECT * FROM location WHERE name = $1;

-- name: ListLocations :many
SELECT * FROM location ORDER BY name ASC;

-- name: SearchLocations :many
SELECT * FROM location
WHERE name ILIKE $1 OR description ILIKE $1
ORDER BY name ASC;

-- name: UpdateLocation :one
UPDATE location
SET name = $2, description = $3, updated_at = $4
WHERE id = $1
    RETURNING *;

-- name: DeleteLocation :exec
DELETE FROM location WHERE id = $1;

-- ============================================================================
-- INVENTORY_TX TABLE QUERIES
-- ============================================================================

-- name: CreateInventoryTx :one
INSERT INTO inventory_tx (id, user_id, product_id, unit_id, location_id, action, quantity, occurred_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING *;

-- name: GetInventoryTx :one
SELECT * FROM inventory_tx WHERE id = $1;

-- name: GetInventoryTxByUserId :many
SELECT * FROM inventory_tx WHERE user_id = $1 ORDER BY occurred_at DESC;

-- name: GetInventoryTxByProductId :many
SELECT * FROM inventory_tx WHERE product_id = $1 ORDER BY occurred_at DESC;

-- name: GetInventoryTxByLocationId :many
SELECT * FROM inventory_tx WHERE location_id = $1 ORDER BY occurred_at DESC;

-- name: GetInventoryTxByAction :many
SELECT * FROM inventory_tx WHERE action = $1 ORDER BY occurred_at DESC;

-- name: ListInventoryTx :many
SELECT * FROM inventory_tx ORDER BY occurred_at DESC;

-- name: ListInventoryTxPaginated :many
SELECT * FROM inventory_tx
ORDER BY occurred_at DESC
    LIMIT $1 OFFSET $2;

-- name: GetInventoryTxWithDetails :many
SELECT
    it.*,
    u.first_name,
    u.last_name,
    u.email,
    p.name as product_name,
    p.description as product_description,
    pu.unit,
    l.name as location_name,
    l.description as location_description
FROM inventory_tx it
         JOIN users u ON it.user_id = u.id
         JOIN product p ON it.product_id = p.id
         JOIN product_unit pu ON it.unit_id = pu.id
         JOIN location l ON it.location_id = l.id
ORDER BY it.occurred_at DESC;

-- name: GetInventoryTxWithDetailsByUserId :many
SELECT
    it.*,
    u.first_name,
    u.last_name,
    u.email,
    p.name as product_name,
    p.description as product_description,
    pu.unit,
    l.name as location_name,
    l.description as location_description
FROM inventory_tx it
         JOIN users u ON it.user_id = u.id
         JOIN product p ON it.product_id = p.id
         JOIN product_unit pu ON it.unit_id = pu.id
         JOIN location l ON it.location_id = l.id
WHERE it.user_id = $1
ORDER BY it.occurred_at DESC;

-- name: GetCurrentInventoryByLocation :many
SELECT
    p.id as product_id,
    p.name as product_name,
    p.description as product_description,
    pu.unit,
    l.id as location_id,
    l.name as location_name,
    SUM(CASE WHEN it.action = 'add' THEN it.quantity ELSE -it.quantity END) as current_quantity
FROM inventory_tx it
         JOIN product p ON it.product_id = p.id
         JOIN product_unit pu ON it.unit_id = pu.id
         JOIN location l ON it.location_id = l.id
WHERE l.id = $1
GROUP BY p.id, p.name, p.description, pu.unit, l.id, l.name
HAVING SUM(CASE WHEN it.action = 'add' THEN it.quantity ELSE -it.quantity END) > 0
ORDER BY p.name ASC;

-- name: GetCurrentInventoryByProduct :many
SELECT
    p.id as product_id,
    p.name as product_name,
    p.description as product_description,
    pu.unit,
    l.id as location_id,
    l.name as location_name,
    SUM(CASE WHEN it.action = 'add' THEN it.quantity ELSE -it.quantity END) as current_quantity
FROM inventory_tx it
         JOIN product p ON it.product_id = p.id
         JOIN product_unit pu ON it.unit_id = pu.id
         JOIN location l ON it.location_id = l.id
WHERE p.id = $1
GROUP BY p.id, p.name, p.description, pu.unit, l.id, l.name
HAVING SUM(CASE WHEN it.action = 'add' THEN it.quantity ELSE -it.quantity END) > 0
ORDER BY l.name ASC;

-- name: GetCurrentInventoryTotal :many
SELECT
    p.id as product_id,
    p.name as product_name,
    p.description as product_description,
    pu.unit,
    SUM(CASE WHEN it.action = 'add' THEN it.quantity ELSE -it.quantity END) as total_quantity
FROM inventory_tx it
         JOIN product p ON it.product_id = p.id
         JOIN product_unit pu ON it.unit_id = pu.id
GROUP BY p.id, p.name, p.description, pu.unit
HAVING SUM(CASE WHEN it.action = 'add' THEN it.quantity ELSE -it.quantity END) > 0
ORDER BY p.name ASC;

-- name: UpdateInventoryTx :one
UPDATE inventory_tx
SET action = $2, quantity = $3, occurred_at = $4
WHERE id = $1
    RETURNING *;

-- name: DeleteInventoryTx :exec
DELETE FROM inventory_tx WHERE id = $1;

-- ============================================================================
-- UTILITY QUERIES
-- ============================================================================

-- name: GetUserWithAuth :one
SELECT
    u.*,
    a.password_hash,
    a.salt
FROM users u
         LEFT JOIN auth a ON u.id = a.user_id
WHERE u.id = $1;

-- name: GetProductWithUnits :one
SELECT
    p.*,
    COALESCE(
            json_agg(
                    json_build_object(
                            'id', pu.id,
                            'unit', pu.unit,
                            'default_unit', pu.default_unit
                    )
            ) FILTER (WHERE pu.id IS NOT NULL),
            '[]'
    ) as units
FROM product p
         LEFT JOIN product_unit pu ON p.id = pu.product_id
WHERE p.id = $1
GROUP BY p.id, p.name, p.description, p.image_url, p.created_at, p.updated_at;

-- name: CountInventoryTxByUser :one
SELECT COUNT(*) FROM inventory_tx WHERE user_id = $1;

-- name: CountInventoryTxByProduct :one
SELECT COUNT(*) FROM inventory_tx WHERE product_id = $1;

-- name: CountInventoryTxByLocation :one
SELECT COUNT(*) FROM inventory_tx WHERE location_id = $1;