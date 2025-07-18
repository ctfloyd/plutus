// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: queries.sql

package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAuth = `-- name: CreateAuth :one

INSERT INTO auth (user_id, password_hash, salt, created_at)
VALUES ($1, $2, $3, $4)
    RETURNING user_id, password_hash, salt, created_at
`

type CreateAuthParams struct {
	UserID       string
	PasswordHash []byte
	Salt         []byte
	CreatedAt    time.Time
}

// ============================================================================
// AUTH TABLE QUERIES
// ============================================================================
func (q *Queries) CreateAuth(ctx context.Context, arg CreateAuthParams) (Auth, error) {
	row := q.db.QueryRow(ctx, createAuth,
		arg.UserID,
		arg.PasswordHash,
		arg.Salt,
		arg.CreatedAt,
	)
	var i Auth
	err := row.Scan(
		&i.UserID,
		&i.PasswordHash,
		&i.Salt,
		&i.CreatedAt,
	)
	return i, err
}

const createInventoryTx = `-- name: CreateInventoryTx :one

INSERT INTO inventory_tx (id, user_id, product_id, unit_id, location_id, action, quantity, occurred_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING id, user_id, product_id, unit_id, location_id, action, quantity, occurred_at
`

type CreateInventoryTxParams struct {
	ID         string
	UserID     *string
	ProductID  *string
	UnitID     *string
	LocationID *string
	Action     string
	Quantity   pgtype.Numeric
	OccurredAt time.Time
}

// ============================================================================
// INVENTORY_TX TABLE QUERIES
// ============================================================================
func (q *Queries) CreateInventoryTx(ctx context.Context, arg CreateInventoryTxParams) (InventoryTx, error) {
	row := q.db.QueryRow(ctx, createInventoryTx,
		arg.ID,
		arg.UserID,
		arg.ProductID,
		arg.UnitID,
		arg.LocationID,
		arg.Action,
		arg.Quantity,
		arg.OccurredAt,
	)
	var i InventoryTx
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ProductID,
		&i.UnitID,
		&i.LocationID,
		&i.Action,
		&i.Quantity,
		&i.OccurredAt,
	)
	return i, err
}

const createLocation = `-- name: CreateLocation :one

INSERT INTO location (id, name, description, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
    RETURNING id, name, description, created_at, updated_at
`

type CreateLocationParams struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// ============================================================================
// LOCATION TABLE QUERIES
// ============================================================================
func (q *Queries) CreateLocation(ctx context.Context, arg CreateLocationParams) (Location, error) {
	row := q.db.QueryRow(ctx, createLocation,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Location
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createProduct = `-- name: CreateProduct :one

INSERT INTO product (id, name, description, image_url, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING id, name, description, image_url, created_at, updated_at
`

type CreateProductParams struct {
	ID          string
	Name        string
	Description string
	ImageUrl    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// ============================================================================
// PRODUCT TABLE QUERIES
// ============================================================================
func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, createProduct,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.ImageUrl,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createProductUnit = `-- name: CreateProductUnit :one

INSERT INTO product_unit (id, product_id, unit, default_unit)
VALUES ($1, $2, $3, $4)
    RETURNING id, product_id, unit, default_unit
`

type CreateProductUnitParams struct {
	ID          string
	ProductID   *string
	Unit        string
	DefaultUnit bool
}

// ============================================================================
// PRODUCT_UNIT TABLE QUERIES
// ============================================================================
func (q *Queries) CreateProductUnit(ctx context.Context, arg CreateProductUnitParams) (ProductUnit, error) {
	row := q.db.QueryRow(ctx, createProductUnit,
		arg.ID,
		arg.ProductID,
		arg.Unit,
		arg.DefaultUnit,
	)
	var i ProductUnit
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.Unit,
		&i.DefaultUnit,
	)
	return i, err
}

const createToken = `-- name: CreateToken :one
INSERT INTO token (token, user_id, revoked)
VALUES ($1, $2, $3)
RETURNING token, user_id, revoked, created_at
`

type CreateTokenParams struct {
	Token   string
	UserID  string
	Revoked bool
}

// ============================================================================
// TOKEN QUERIES
// ============================================================================
func (q *Queries) CreateToken(ctx context.Context, arg CreateTokenParams) (Token, error) {
	row := q.db.QueryRow(ctx, createToken, arg.Token, arg.UserID, arg.Revoked)
	var i Token
	err := row.Scan(
		&i.Token,
		&i.UserID,
		&i.Revoked,
		&i.CreatedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, first_name, last_name, email, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING id, first_name, last_name, email, created_at, updated_at
`

type CreateUserParams struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAuth = `-- name: DeleteAuth :exec
DELETE FROM auth WHERE user_id = $1
`

func (q *Queries) DeleteAuth(ctx context.Context, userID string) error {
	_, err := q.db.Exec(ctx, deleteAuth, userID)
	return err
}

const deleteInventoryTx = `-- name: DeleteInventoryTx :exec
DELETE FROM inventory_tx WHERE id = $1
`

func (q *Queries) DeleteInventoryTx(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteInventoryTx, id)
	return err
}

const deleteLocation = `-- name: DeleteLocation :exec
DELETE FROM location WHERE id = $1
`

func (q *Queries) DeleteLocation(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteLocation, id)
	return err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM product WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteProduct, id)
	return err
}

const deleteProductUnit = `-- name: DeleteProductUnit :exec
DELETE FROM product_unit WHERE id = $1
`

func (q *Queries) DeleteProductUnit(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteProductUnit, id)
	return err
}

const deleteToken = `-- name: DeleteToken :exec
DELETE FROM token
WHERE token = $1
`

func (q *Queries) DeleteToken(ctx context.Context, token string) error {
	_, err := q.db.Exec(ctx, deleteToken, token)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getAuth = `-- name: GetAuth :one
SELECT user_id, password_hash, salt, created_at FROM auth WHERE user_id = $1
`

func (q *Queries) GetAuth(ctx context.Context, userID string) (Auth, error) {
	row := q.db.QueryRow(ctx, getAuth, userID)
	var i Auth
	err := row.Scan(
		&i.UserID,
		&i.PasswordHash,
		&i.Salt,
		&i.CreatedAt,
	)
	return i, err
}

const getAuthByUserEmail = `-- name: GetAuthByUserEmail :one
SELECT a.user_id, a.password_hash, a.salt, a.created_at FROM auth a
                    JOIN users u ON a.user_id = u.id
WHERE u.email = $1
`

func (q *Queries) GetAuthByUserEmail(ctx context.Context, email string) (Auth, error) {
	row := q.db.QueryRow(ctx, getAuthByUserEmail, email)
	var i Auth
	err := row.Scan(
		&i.UserID,
		&i.PasswordHash,
		&i.Salt,
		&i.CreatedAt,
	)
	return i, err
}

const getCurrentInventoryByLocation = `-- name: GetCurrentInventoryByLocation :many
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
ORDER BY p.name ASC
`

type GetCurrentInventoryByLocationRow struct {
	ProductID          string
	ProductName        string
	ProductDescription string
	Unit               string
	LocationID         string
	LocationName       string
	CurrentQuantity    int64
}

func (q *Queries) GetCurrentInventoryByLocation(ctx context.Context, id string) ([]GetCurrentInventoryByLocationRow, error) {
	rows, err := q.db.Query(ctx, getCurrentInventoryByLocation, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCurrentInventoryByLocationRow
	for rows.Next() {
		var i GetCurrentInventoryByLocationRow
		if err := rows.Scan(
			&i.ProductID,
			&i.ProductName,
			&i.ProductDescription,
			&i.Unit,
			&i.LocationID,
			&i.LocationName,
			&i.CurrentQuantity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCurrentInventoryByProduct = `-- name: GetCurrentInventoryByProduct :many
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
ORDER BY l.name ASC
`

type GetCurrentInventoryByProductRow struct {
	ProductID          string
	ProductName        string
	ProductDescription string
	Unit               string
	LocationID         string
	LocationName       string
	CurrentQuantity    int64
}

func (q *Queries) GetCurrentInventoryByProduct(ctx context.Context, id string) ([]GetCurrentInventoryByProductRow, error) {
	rows, err := q.db.Query(ctx, getCurrentInventoryByProduct, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCurrentInventoryByProductRow
	for rows.Next() {
		var i GetCurrentInventoryByProductRow
		if err := rows.Scan(
			&i.ProductID,
			&i.ProductName,
			&i.ProductDescription,
			&i.Unit,
			&i.LocationID,
			&i.LocationName,
			&i.CurrentQuantity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCurrentInventoryTotal = `-- name: GetCurrentInventoryTotal :many
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
ORDER BY p.name ASC
`

type GetCurrentInventoryTotalRow struct {
	ProductID          string
	ProductName        string
	ProductDescription string
	Unit               string
	TotalQuantity      int64
}

func (q *Queries) GetCurrentInventoryTotal(ctx context.Context) ([]GetCurrentInventoryTotalRow, error) {
	rows, err := q.db.Query(ctx, getCurrentInventoryTotal)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCurrentInventoryTotalRow
	for rows.Next() {
		var i GetCurrentInventoryTotalRow
		if err := rows.Scan(
			&i.ProductID,
			&i.ProductName,
			&i.ProductDescription,
			&i.Unit,
			&i.TotalQuantity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getDefaultProductUnit = `-- name: GetDefaultProductUnit :one
SELECT id, product_id, unit, default_unit FROM product_unit WHERE product_id = $1 AND default_unit = true
`

func (q *Queries) GetDefaultProductUnit(ctx context.Context, productID *string) (ProductUnit, error) {
	row := q.db.QueryRow(ctx, getDefaultProductUnit, productID)
	var i ProductUnit
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.Unit,
		&i.DefaultUnit,
	)
	return i, err
}

const getInventoryTx = `-- name: GetInventoryTx :one
SELECT id, user_id, product_id, unit_id, location_id, action, quantity, occurred_at FROM inventory_tx WHERE id = $1
`

func (q *Queries) GetInventoryTx(ctx context.Context, id string) (InventoryTx, error) {
	row := q.db.QueryRow(ctx, getInventoryTx, id)
	var i InventoryTx
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ProductID,
		&i.UnitID,
		&i.LocationID,
		&i.Action,
		&i.Quantity,
		&i.OccurredAt,
	)
	return i, err
}

const getInventoryTxByAction = `-- name: GetInventoryTxByAction :many
SELECT id, user_id, product_id, unit_id, location_id, action, quantity, occurred_at FROM inventory_tx WHERE action = $1 ORDER BY occurred_at DESC
`

func (q *Queries) GetInventoryTxByAction(ctx context.Context, action string) ([]InventoryTx, error) {
	rows, err := q.db.Query(ctx, getInventoryTxByAction, action)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []InventoryTx
	for rows.Next() {
		var i InventoryTx
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ProductID,
			&i.UnitID,
			&i.LocationID,
			&i.Action,
			&i.Quantity,
			&i.OccurredAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInventoryTxByLocationId = `-- name: GetInventoryTxByLocationId :many
SELECT id, user_id, product_id, unit_id, location_id, action, quantity, occurred_at FROM inventory_tx WHERE location_id = $1 ORDER BY occurred_at DESC
`

func (q *Queries) GetInventoryTxByLocationId(ctx context.Context, locationID *string) ([]InventoryTx, error) {
	rows, err := q.db.Query(ctx, getInventoryTxByLocationId, locationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []InventoryTx
	for rows.Next() {
		var i InventoryTx
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ProductID,
			&i.UnitID,
			&i.LocationID,
			&i.Action,
			&i.Quantity,
			&i.OccurredAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInventoryTxByProductId = `-- name: GetInventoryTxByProductId :many
SELECT id, user_id, product_id, unit_id, location_id, action, quantity, occurred_at FROM inventory_tx WHERE product_id = $1 ORDER BY occurred_at DESC
`

func (q *Queries) GetInventoryTxByProductId(ctx context.Context, productID *string) ([]InventoryTx, error) {
	rows, err := q.db.Query(ctx, getInventoryTxByProductId, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []InventoryTx
	for rows.Next() {
		var i InventoryTx
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ProductID,
			&i.UnitID,
			&i.LocationID,
			&i.Action,
			&i.Quantity,
			&i.OccurredAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInventoryTxByUserId = `-- name: GetInventoryTxByUserId :many
SELECT id, user_id, product_id, unit_id, location_id, action, quantity, occurred_at FROM inventory_tx WHERE user_id = $1 ORDER BY occurred_at DESC
`

func (q *Queries) GetInventoryTxByUserId(ctx context.Context, userID *string) ([]InventoryTx, error) {
	rows, err := q.db.Query(ctx, getInventoryTxByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []InventoryTx
	for rows.Next() {
		var i InventoryTx
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ProductID,
			&i.UnitID,
			&i.LocationID,
			&i.Action,
			&i.Quantity,
			&i.OccurredAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInventoryTxWithDetails = `-- name: GetInventoryTxWithDetails :many
SELECT
    it.id, it.user_id, it.product_id, it.unit_id, it.location_id, it.action, it.quantity, it.occurred_at,
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
ORDER BY it.occurred_at DESC
`

type GetInventoryTxWithDetailsRow struct {
	ID                  string
	UserID              *string
	ProductID           *string
	UnitID              *string
	LocationID          *string
	Action              string
	Quantity            pgtype.Numeric
	OccurredAt          time.Time
	FirstName           string
	LastName            string
	Email               string
	ProductName         string
	ProductDescription  string
	Unit                string
	LocationName        string
	LocationDescription string
}

func (q *Queries) GetInventoryTxWithDetails(ctx context.Context) ([]GetInventoryTxWithDetailsRow, error) {
	rows, err := q.db.Query(ctx, getInventoryTxWithDetails)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetInventoryTxWithDetailsRow
	for rows.Next() {
		var i GetInventoryTxWithDetailsRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ProductID,
			&i.UnitID,
			&i.LocationID,
			&i.Action,
			&i.Quantity,
			&i.OccurredAt,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.ProductName,
			&i.ProductDescription,
			&i.Unit,
			&i.LocationName,
			&i.LocationDescription,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInventoryTxWithDetailsByUserId = `-- name: GetInventoryTxWithDetailsByUserId :many
SELECT
    it.id, it.user_id, it.product_id, it.unit_id, it.location_id, it.action, it.quantity, it.occurred_at,
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
ORDER BY it.occurred_at DESC
`

type GetInventoryTxWithDetailsByUserIdRow struct {
	ID                  string
	UserID              *string
	ProductID           *string
	UnitID              *string
	LocationID          *string
	Action              string
	Quantity            pgtype.Numeric
	OccurredAt          time.Time
	FirstName           string
	LastName            string
	Email               string
	ProductName         string
	ProductDescription  string
	Unit                string
	LocationName        string
	LocationDescription string
}

func (q *Queries) GetInventoryTxWithDetailsByUserId(ctx context.Context, userID *string) ([]GetInventoryTxWithDetailsByUserIdRow, error) {
	rows, err := q.db.Query(ctx, getInventoryTxWithDetailsByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetInventoryTxWithDetailsByUserIdRow
	for rows.Next() {
		var i GetInventoryTxWithDetailsByUserIdRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ProductID,
			&i.UnitID,
			&i.LocationID,
			&i.Action,
			&i.Quantity,
			&i.OccurredAt,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.ProductName,
			&i.ProductDescription,
			&i.Unit,
			&i.LocationName,
			&i.LocationDescription,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLocation = `-- name: GetLocation :one
SELECT id, name, description, created_at, updated_at FROM location WHERE id = $1
`

func (q *Queries) GetLocation(ctx context.Context, id string) (Location, error) {
	row := q.db.QueryRow(ctx, getLocation, id)
	var i Location
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getLocationByName = `-- name: GetLocationByName :one
SELECT id, name, description, created_at, updated_at FROM location WHERE name = $1
`

func (q *Queries) GetLocationByName(ctx context.Context, name string) (Location, error) {
	row := q.db.QueryRow(ctx, getLocationByName, name)
	var i Location
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProduct = `-- name: GetProduct :one
SELECT id, name, description, image_url, created_at, updated_at FROM product WHERE id = $1
`

func (q *Queries) GetProduct(ctx context.Context, id string) (Product, error) {
	row := q.db.QueryRow(ctx, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProductByName = `-- name: GetProductByName :one
SELECT id, name, description, image_url, created_at, updated_at FROM product WHERE name = $1
`

func (q *Queries) GetProductByName(ctx context.Context, name string) (Product, error) {
	row := q.db.QueryRow(ctx, getProductByName, name)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProductUnit = `-- name: GetProductUnit :one
SELECT id, product_id, unit, default_unit FROM product_unit WHERE id = $1
`

func (q *Queries) GetProductUnit(ctx context.Context, id string) (ProductUnit, error) {
	row := q.db.QueryRow(ctx, getProductUnit, id)
	var i ProductUnit
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.Unit,
		&i.DefaultUnit,
	)
	return i, err
}

const getProductUnitsByProductId = `-- name: GetProductUnitsByProductId :many
SELECT id, product_id, unit, default_unit FROM product_unit WHERE product_id = $1 ORDER BY default_unit DESC, unit ASC
`

func (q *Queries) GetProductUnitsByProductId(ctx context.Context, productID *string) ([]ProductUnit, error) {
	rows, err := q.db.Query(ctx, getProductUnitsByProductId, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProductUnit
	for rows.Next() {
		var i ProductUnit
		if err := rows.Scan(
			&i.ID,
			&i.ProductID,
			&i.Unit,
			&i.DefaultUnit,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getToken = `-- name: GetToken :one
SELECT token, user_id, revoked, created_at FROM token
WHERE token = $1
`

func (q *Queries) GetToken(ctx context.Context, token string) (Token, error) {
	row := q.db.QueryRow(ctx, getToken, token)
	var i Token
	err := row.Scan(
		&i.Token,
		&i.UserID,
		&i.Revoked,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, first_name, last_name, email, created_at, updated_at FROM users WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, first_name, last_name, email, created_at, updated_at FROM users where email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const isEmailTaken = `-- name: IsEmailTaken :one
SELECT EXISTS(SELECT 1 FROM users where email = $1)
`

func (q *Queries) IsEmailTaken(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRow(ctx, isEmailTaken, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const listAuth = `-- name: ListAuth :many
SELECT user_id, password_hash, salt, created_at FROM auth ORDER BY created_at DESC
`

func (q *Queries) ListAuth(ctx context.Context) ([]Auth, error) {
	rows, err := q.db.Query(ctx, listAuth)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Auth
	for rows.Next() {
		var i Auth
		if err := rows.Scan(
			&i.UserID,
			&i.PasswordHash,
			&i.Salt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listInventoryTx = `-- name: ListInventoryTx :many
SELECT id, user_id, product_id, unit_id, location_id, action, quantity, occurred_at FROM inventory_tx ORDER BY occurred_at DESC
`

func (q *Queries) ListInventoryTx(ctx context.Context) ([]InventoryTx, error) {
	rows, err := q.db.Query(ctx, listInventoryTx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []InventoryTx
	for rows.Next() {
		var i InventoryTx
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ProductID,
			&i.UnitID,
			&i.LocationID,
			&i.Action,
			&i.Quantity,
			&i.OccurredAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listInventoryTxPaginated = `-- name: ListInventoryTxPaginated :many
SELECT id, user_id, product_id, unit_id, location_id, action, quantity, occurred_at FROM inventory_tx
ORDER BY occurred_at DESC
    LIMIT $1 OFFSET $2
`

type ListInventoryTxPaginatedParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListInventoryTxPaginated(ctx context.Context, arg ListInventoryTxPaginatedParams) ([]InventoryTx, error) {
	rows, err := q.db.Query(ctx, listInventoryTxPaginated, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []InventoryTx
	for rows.Next() {
		var i InventoryTx
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ProductID,
			&i.UnitID,
			&i.LocationID,
			&i.Action,
			&i.Quantity,
			&i.OccurredAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLocations = `-- name: ListLocations :many
SELECT id, name, description, created_at, updated_at FROM location ORDER BY name ASC
`

func (q *Queries) ListLocations(ctx context.Context) ([]Location, error) {
	rows, err := q.db.Query(ctx, listLocations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Location
	for rows.Next() {
		var i Location
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProductUnits = `-- name: ListProductUnits :many
SELECT id, product_id, unit, default_unit FROM product_unit ORDER BY product_id, default_unit DESC, unit ASC
`

func (q *Queries) ListProductUnits(ctx context.Context) ([]ProductUnit, error) {
	rows, err := q.db.Query(ctx, listProductUnits)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProductUnit
	for rows.Next() {
		var i ProductUnit
		if err := rows.Scan(
			&i.ID,
			&i.ProductID,
			&i.Unit,
			&i.DefaultUnit,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProducts = `-- name: ListProducts :many
SELECT id, name, description, image_url, created_at, updated_at FROM product ORDER BY created_at DESC
`

func (q *Queries) ListProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.Query(ctx, listProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.ImageUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT id, first_name, last_name, email, created_at, updated_at FROM users ORDER BY created_at DESC
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const revokeAllUserTokens = `-- name: RevokeAllUserTokens :exec
UPDATE token
SET revoked = true
WHERE user_id = $1
`

func (q *Queries) RevokeAllUserTokens(ctx context.Context, userID string) error {
	_, err := q.db.Exec(ctx, revokeAllUserTokens, userID)
	return err
}

const revokeToken = `-- name: RevokeToken :exec
UPDATE token
SET revoked = true
WHERE token = $1
`

func (q *Queries) RevokeToken(ctx context.Context, token string) error {
	_, err := q.db.Exec(ctx, revokeToken, token)
	return err
}

const searchLocations = `-- name: SearchLocations :many
SELECT id, name, description, created_at, updated_at FROM location
WHERE name ILIKE $1 OR description ILIKE $1
ORDER BY name ASC
`

func (q *Queries) SearchLocations(ctx context.Context, name string) ([]Location, error) {
	rows, err := q.db.Query(ctx, searchLocations, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Location
	for rows.Next() {
		var i Location
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchProducts = `-- name: SearchProducts :many
SELECT id, name, description, image_url, created_at, updated_at FROM product
WHERE name ILIKE $1 OR description ILIKE $1
ORDER BY created_at DESC
`

func (q *Queries) SearchProducts(ctx context.Context, name string) ([]Product, error) {
	rows, err := q.db.Query(ctx, searchProducts, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.ImageUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setProductUnitAsDefault = `-- name: SetProductUnitAsDefault :exec
UPDATE product_unit
SET default_unit = CASE WHEN id = $2 THEN true ELSE false END
WHERE product_id = $1
`

type SetProductUnitAsDefaultParams struct {
	ProductID *string
	ID        string
}

func (q *Queries) SetProductUnitAsDefault(ctx context.Context, arg SetProductUnitAsDefaultParams) error {
	_, err := q.db.Exec(ctx, setProductUnitAsDefault, arg.ProductID, arg.ID)
	return err
}

const updateAuth = `-- name: UpdateAuth :one
UPDATE auth
SET password_hash = $2, salt = $3
WHERE user_id = $1
    RETURNING user_id, password_hash, salt, created_at
`

type UpdateAuthParams struct {
	UserID       string
	PasswordHash []byte
	Salt         []byte
}

func (q *Queries) UpdateAuth(ctx context.Context, arg UpdateAuthParams) (Auth, error) {
	row := q.db.QueryRow(ctx, updateAuth, arg.UserID, arg.PasswordHash, arg.Salt)
	var i Auth
	err := row.Scan(
		&i.UserID,
		&i.PasswordHash,
		&i.Salt,
		&i.CreatedAt,
	)
	return i, err
}

const updateInventoryTx = `-- name: UpdateInventoryTx :one
UPDATE inventory_tx
SET action = $2, quantity = $3, occurred_at = $4
WHERE id = $1
    RETURNING id, user_id, product_id, unit_id, location_id, action, quantity, occurred_at
`

type UpdateInventoryTxParams struct {
	ID         string
	Action     string
	Quantity   pgtype.Numeric
	OccurredAt time.Time
}

func (q *Queries) UpdateInventoryTx(ctx context.Context, arg UpdateInventoryTxParams) (InventoryTx, error) {
	row := q.db.QueryRow(ctx, updateInventoryTx,
		arg.ID,
		arg.Action,
		arg.Quantity,
		arg.OccurredAt,
	)
	var i InventoryTx
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ProductID,
		&i.UnitID,
		&i.LocationID,
		&i.Action,
		&i.Quantity,
		&i.OccurredAt,
	)
	return i, err
}

const updateLocation = `-- name: UpdateLocation :one
UPDATE location
SET name = $2, description = $3, updated_at = $4
WHERE id = $1
    RETURNING id, name, description, created_at, updated_at
`

type UpdateLocationParams struct {
	ID          string
	Name        string
	Description string
	UpdatedAt   time.Time
}

func (q *Queries) UpdateLocation(ctx context.Context, arg UpdateLocationParams) (Location, error) {
	row := q.db.QueryRow(ctx, updateLocation,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.UpdatedAt,
	)
	var i Location
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE product
SET name = $2, description = $3, image_url = $4, updated_at = $5
WHERE id = $1
    RETURNING id, name, description, image_url, created_at, updated_at
`

type UpdateProductParams struct {
	ID          string
	Name        string
	Description string
	ImageUrl    string
	UpdatedAt   time.Time
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, updateProduct,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.ImageUrl,
		arg.UpdatedAt,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateProductUnit = `-- name: UpdateProductUnit :one
UPDATE product_unit
SET unit = $2, default_unit = $3
WHERE id = $1
    RETURNING id, product_id, unit, default_unit
`

type UpdateProductUnitParams struct {
	ID          string
	Unit        string
	DefaultUnit bool
}

func (q *Queries) UpdateProductUnit(ctx context.Context, arg UpdateProductUnitParams) (ProductUnit, error) {
	row := q.db.QueryRow(ctx, updateProductUnit, arg.ID, arg.Unit, arg.DefaultUnit)
	var i ProductUnit
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.Unit,
		&i.DefaultUnit,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET first_name = $2, last_name = $3, email = $4, updated_at = $5
WHERE id = $1
    RETURNING id, first_name, last_name, email, created_at, updated_at
`

type UpdateUserParams struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	UpdatedAt time.Time
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
