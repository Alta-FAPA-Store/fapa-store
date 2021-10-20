package pet

import "time"

//Pet product Pet that available to rent or sell
type Pet struct {
	ID         int
	UserID     int
	Name       string
	Kind       string
	CreatedAt  time.Time
	CreatedBy  string
	ModifiedAt time.Time
	ModifiedBy string
	Version    int
}

//NewPet create new Pet
func NewPet(
	id int,
	userID int,
	name string,
	kind string,
	creator string,
	createdAt time.Time) Pet {

	return Pet{
		ID:         id,
		UserID:     userID,
		Name:       name,
		Kind:       kind,
		CreatedAt:  createdAt,
		CreatedBy:  creator,
		ModifiedAt: createdAt,
		ModifiedBy: creator,
		Version:    1,
	}
}

//ModifyPet update existing Pet data
func (oldData *Pet) ModifyPet(newName string, modifiedAt time.Time, updater string) Pet {
	return Pet{
		ID:         oldData.ID,
		UserID:     oldData.UserID,
		Name:       newName,
		Kind:       oldData.Kind,
		CreatedAt:  oldData.CreatedAt,
		CreatedBy:  oldData.CreatedBy,
		ModifiedAt: modifiedAt,
		ModifiedBy: updater,
		Version:    oldData.Version + 1,
	}
}
