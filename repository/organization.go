package repository

import (
	"github.com/halfbakedio/saas/ent"
)

type IOrganizationRepository interface {
	CreateOrganization(org *ent.Organization) (*ent.Organization, error)
	FindOrganizationByName(name string) (*ent.Organization, error)
}
