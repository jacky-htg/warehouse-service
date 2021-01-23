package model

import (
	"context"
	"database/sql"
	"inventory-service/internal/pkg/app"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Inventory struct
type Inventory struct {
	ID               string
	CompanyID        string
	BranchID         string
	ProductID        string
	Barcode          string
	TransactionID    string
	TransactionCode  string
	TransaactionDate time.Time
	Type             string
	IsIn             bool
	ShelveID         string
}

// Get func
func (u *Inventory) Get(ctx context.Context, tx *sql.Tx) error {
	query := `
		SELECT id, company_id, branch_id, product_id, barcode, transaction_id, transaction_code, transaction_date, type, in_out, shelve_id 
		FROM inventories
		WHERE company_id = $1 AND product_id = $2 AND barcode = $3 AND transaction_id = $4
	`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare statement Get inventory: %v", err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, ctx.Value(app.Ctx("companyID")).(string), u.ProductID, u.Barcode, u.TransactionID).Scan(
		&u.ID, &u.CompanyID, &u.BranchID, &u.ProductID, &u.Barcode,
		&u.TransactionID, &u.TransactionCode, &u.TransaactionDate,
		&u.Type, &u.IsIn, &u.ShelveID,
	)

	if err == sql.ErrNoRows {
		return status.Errorf(codes.NotFound, "Query Raw get inventory: %v", err)
	}

	if err != nil {
		return status.Errorf(codes.Internal, "Query Raw get inventory: %v", err)
	}

	return nil
}

// Create Inventory
func (u *Inventory) Create(ctx context.Context, tx *sql.Tx) error {
	u.ID = uuid.New().String()
	now := time.Now().UTC()

	query := `
		INSERT INTO inventories (
			id, company_id, branch_id, product_id, barcode, 
			transaction_id, transaction_code, transaction_date, 
			type, in_out, shelve_id, createde_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare insert inventory: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		u.ID,
		ctx.Value(app.Ctx("companyID")).(string),
		u.BranchID,
		u.ProductID,
		u.Barcode,
		u.TransactionID,
		u.TransactionCode,
		u.TransaactionDate,
		u.Type,
		u.IsIn,
		u.ShelveID,
		now,
		now,
	)
	if err != nil {
		return status.Errorf(codes.Internal, "Exec insert inventory: %v", err)
	}

	return nil
}

// Update Inventory
func (u *Inventory) Update(ctx context.Context, tx *sql.Tx) error {
	query := `
		UPDATE inventories SET
		branch_id = $1, 
		product_id = $2, 
		barcode = $3, 
		transaction_id = $4, 
		transaction_code = $5, 
		transaction_date = $6, 
		type = $7, 
		in_out = $8, 
		shelve_id = $9, 
		updated_at= $10
		WHERE id = $4
	`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare update inventory: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		u.BranchID,
		u.ProductID,
		u.Barcode,
		u.TransactionID,
		u.TransactionCode,
		u.TransaactionDate,
		u.Type,
		u.IsIn,
		u.ShelveID,
		time.Now().UTC(),
		u.ID,
	)
	if err != nil {
		return status.Errorf(codes.Internal, "Exec update inventory: %v", err)
	}

	return nil
}

// Delete Inventory
func (u *Inventory) Delete(ctx context.Context, tx *sql.Tx) error {
	stmt, err := tx.PrepareContext(ctx, `DELETE FROM inventories WHERE id = $1`)
	if err != nil {
		return status.Errorf(codes.Internal, "Prepare delete inventory: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, u.ID)
	if err != nil {
		return status.Errorf(codes.Internal, "Exec delete inventory: %v", err)
	}

	return nil
}
