package repository

import (
	"context"
	"fmt"
	"vm-crud/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type VMRepository struct {
	dbpool *pgxpool.Pool
}

func NewVMRepository(dbpool *pgxpool.Pool) *VMRepository {
	return &VMRepository{
		dbpool: dbpool,
	}
}

func (r *VMRepository) CreateVM(ctx context.Context, req model.CreateVMRequest) (*model.VM, error) {
	query := `Insert into vms (name , cpu, ram)
	values ($1, $2, $3)
	returning id, name, cpu, ram, status, created_at, updated_at
	`
	var vm model.VM
	err := r.dbpool.QueryRow(ctx, query, req.Name, req.CPU, req.RAM).Scan(
		&vm.ID,
		&vm.Name,
		&vm.CPU,
		&vm.RAM,
		&vm.Status,
		&vm.CreatedAt,
		&vm.UpdatedAt,
	)
	if err != nil {
		fmt.Println("Error creating VM in repository:", err)
		return nil, fmt.Errorf("failed to create vm: %w", err)
	}
	return &vm, nil
}


func (r *VMRepository) ListVMs(ctx context.Context) ([]model.VM, error) {
	query :=  `SELECT id, name, cpu, ram, status, created_at, updated_at FROM vms`
	var vms []model.VM
	rows, err := r.dbpool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list vms: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var vm model.VM
		if err := rows.Scan(
			&vm.ID,
			&vm.Name,
			&vm.CPU,
			&vm.RAM,
			&vm.Status,
			&vm.CreatedAt,
			&vm.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan vm: %w", err)
		}
		vms = append(vms, vm)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", rows.Err())
	}
	return vms, nil
}


func  (r *VMRepository) GetVMByID(ctx context.Context, id int) (*model.VM, error) {
	query := `SELECT id, name, cpu, ram, status, created_at, updated_at FROM vms WHERE id = $1`
	var vm model.VM
	err := r.dbpool.QueryRow(ctx, query, id).Scan(
		&vm.ID,
		&vm.Name,
		&vm.CPU,
		&vm.RAM,
		&vm.Status,
		&vm.CreatedAt,
		&vm.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get vm by id: %w", err)
	}
	return &vm, nil
}



func (r *VMRepository) UpdateVM(ctx context.Context, id int, req model.UpdateVMRequest) (*model.VM, error) {
	query := `UPDATE vms SET name = $1, cpu = $2, ram = $3, updated_at = NOW() WHERE id = $4 RETURNING id, name, cpu, ram, status, created_at, updated_at`
	var vm model.VM

	err := r.dbpool.QueryRow(ctx, query, req.Name, req.CPU, req.RAM, id).Scan(
		&vm.ID,
		&vm.Name,
		&vm.CPU,
		&vm.RAM,
		&vm.Status,
		&vm.CreatedAt,
		&vm.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update vm: %w", err)
	}
	return &vm, nil
}

func (r *VMRepository) DeleteVM(ctx context.Context, id int) error {
	query := `DELETE FROM vms WHERE id = $1`
	_, err := r.dbpool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete vm: %w", err)
	}
	return nil
}	