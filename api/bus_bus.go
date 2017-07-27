package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BusBusService struct {
	client *Client
}

func NewBusBusService(c *Client) *BusBusService {
	return &BusBusService{client: c}
}

func (svc *BusBusService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.BusBusModel, name)
}

func (svc *BusBusService) GetByIds(ids []int) (*types.BusBuss, error) {
	b := &types.BusBuss{}
	return b, svc.client.getByIds(types.BusBusModel, ids, b)
}

func (svc *BusBusService) GetByName(name string) (*types.BusBuss, error) {
	b := &types.BusBuss{}
	return b, svc.client.getByName(types.BusBusModel, name, b)
}

func (svc *BusBusService) GetAll() (*types.BusBuss, error) {
	b := &types.BusBuss{}
	return b, svc.client.getAll(types.BusBusModel, b)
}

func (svc *BusBusService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.BusBusModel, fields, relations)
}

func (svc *BusBusService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BusBusModel, ids, fields, relations)
}

func (svc *BusBusService) Delete(ids []int) error {
	return svc.client.delete(types.BusBusModel, ids)
}