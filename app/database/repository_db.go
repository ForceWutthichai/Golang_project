package database //เชื่อม database

import (
	"context"
	"fmt"

	"todo/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	CreatePatient(ctx context.Context, createPatientRequest *models.CreatePatientRequest) error
	UpdatePatient(ctx context.Context, updatePatientRequest *models.UpdatePatientRequest) error
	ReadPatient(ctx context.Context, req *models.ResponseReadPatient) (*[]models.ResponseReadPatient, error)
	ReadPatientAll(ctx context.Context) (*[]models.ResponseReadPatientAll, error)
}

type RepositoryDB struct {
	pool *pgxpool.Pool
}

func NewRepositoryDB(pool *pgxpool.Pool) Repository {
	return &RepositoryDB{
		pool: pool,
	}
}

func (r *RepositoryDB) CreatePatient(ctx context.Context, createPatientRequest *models.CreatePatientRequest) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	stmt := `INSERT INTO patient (first_name, last_name, address,phone, gender, id_card, date_birth) 
	VALUES(@first_name, @last_name, @address,@phone, @gender, @id_card, @date_birth);`
	args := pgx.NamedArgs{
		"first_name": createPatientRequest.FirstName,
		"last_name":  createPatientRequest.LastName,
		"address":    createPatientRequest.Address,
		"phone":      createPatientRequest.Phone,
		"gender":     createPatientRequest.Gender,
		"id_card":    createPatientRequest.IdCard,
		"date_birth": createPatientRequest.DateBirth,
	}

	_, err = tx.Exec(ctx, stmt, args)
	return err
}

func (r *RepositoryDB) UpdatePatient(ctx context.Context, editPatientRequest *models.UpdatePatientRequest) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	stmt := `UPDATE patient SET first_name = @first_name, last_name = @last_name, address = @address, phone = @phone, 
	gender = @gender, id_card = @id_card, date_birth = @date_birth 
	WHERE id = @id;`
	args := pgx.NamedArgs{
		"id":         editPatientRequest.Id,
		"first_name": editPatientRequest.FirstName,
		"last_name":  editPatientRequest.LastName,
		"address":    editPatientRequest.Address,
		"phone":      editPatientRequest.Phone,
		"gender":     editPatientRequest.Gender,
		"id_card":    editPatientRequest.IdCard,
		"date_birth": editPatientRequest.DateBirth,
	}

	_, err = tx.Exec(ctx, stmt, args)
	return err
}

func (r *RepositoryDB) ReadPatient(ctx context.Context, req *models.ResponseReadPatient) (*[]models.ResponseReadPatient, error) {
	query := `SELECT p.id, p.first_name, p.last_name, p.address, p.phone, p.gender, p.id_card, p.date_birth FROM patient p WHERE 1=1`

	// Add other optional query parameters here
	if req.FirstName != "" {
		query += " AND p.first_name = @first_name" + fmt.Sprint(*&req.FirstName)
	}
	if req.LastName != "" {
		query += " AND p.last_name = @last_name" + fmt.Sprint(*&req.LastName)
	}
	if req.Address != "" {
		query += " AND p.address = @address" + fmt.Sprint(*&req.Address)
	}
	if req.Phone != "" {
		query += " AND p.phone = @phone" + fmt.Sprint(*&req.Phone)
	}
	if req.Gender != "" {
		query += " AND p.gender = @gender" + fmt.Sprint(*&req.Gender)
	}
	if req.IdCard != "" {
		query += " AND p.id_card = @id_card" + fmt.Sprint(*&req.IdCard)
	}
	if req.DateBirth != "" {
		query += " AND p.date_birth = @date_birth" + fmt.Sprint(*&req.DateBirth)
	}

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []models.ResponseReadPatient
	for rows.Next() {
		var patient models.ResponseReadPatient
		if err := rows.Scan(&patient.Id, &patient.FirstName, &patient.LastName, &patient.Address, &patient.Phone, &patient.Gender, &patient.IdCard, &patient.DateBirth); err != nil {
			return nil, err
		}
		patients = append(patients, patient)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(patients) == 0 {
		return &[]models.ResponseReadPatient{}, nil
	}

	return &patients, nil
}

func (r *RepositoryDB) ReadPatientAll(ctx context.Context) (*[]models.ResponseReadPatientAll, error) {
	query := `SELECT p.id, p.first_name, p.last_name, p.address, p.phone, p.gender, p.id_card, p.date_birth FROM patient p; `

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var responseReadPatientList []models.ResponseReadPatientAll
	for rows.Next() {
		var responseReadPatient models.ResponseReadPatientAll
		err := rows.Scan(
			&responseReadPatient.Id,
			&responseReadPatient.FirstName,
			&responseReadPatient.LastName,
			&responseReadPatient.Address,
			&responseReadPatient.Phone,
			&responseReadPatient.Gender,
			&responseReadPatient.IdCard,
			&responseReadPatient.DateBirth,
		)
		if err != nil {
			return nil, err
		}
		responseReadPatientList = append(responseReadPatientList, responseReadPatient)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(responseReadPatientList) == 0 {
		return &[]models.ResponseReadPatientAll{}, nil
	}

	return &responseReadPatientList, nil
}
