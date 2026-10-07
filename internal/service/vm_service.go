package service

import (
	"context"
	"fmt"

	"vm-crud/internal/model"
	"vm-crud/internal/repository"
)

type VMService struct {
	repo *repository.VMRepository
}

func NewVMService(repo *repository.VMRepository) *VMService {
	return &VMService{
		repo: repo,
	}
}

func (s *VMService) CreateVM(
	ctx context.Context,
	req model.CreateVMRequest,
) (*model.VM, error) {

	vm, err := s.repo.CreateVM(ctx, req)
	if err != nil {
		fmt.Println("Error creating VM in service:", err)
		return nil, err
	}

	return vm, nil
}
func (s *VMService) ListVMs(ctx context.Context) ([]model.VM, error) {
	vms, err := s.repo.ListVMs(ctx)
	if err != nil {
		fmt.Println("Error listing VMs in service:", err)
		return nil, err
	}
	return vms, nil
}

// get vm by id service
func (s *VMService) GetVMByID(ctx context.Context, id int) (*model.VM, error) {
	vm, err := s.repo.GetVMByID(ctx, id)
	if err != nil {
		fmt.Println("Error getting VM by ID in service:", err)
		return nil, err
	}
	
	return vm, nil
}

// update vm service
func (s *VMService) UpdateVM(ctx context.Context, id int, req model.UpdateVMRequest) (*model.VM, error) {
	vm, err := s.repo.UpdateVM(ctx, id, req)
	if err != nil {
		fmt.Println("Error updating VM in service:", err)
		return nil, err
	}
	return vm, nil
}

// delete vm service
func (s *VMService) DeleteVM(ctx context.Context, id int) error {
	err := s.repo.DeleteVM(ctx, id)
	if err != nil {
		fmt.Println("Error deleting VM in service:", err)
		return err
	}
	return nil
}	


