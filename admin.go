package pocketclient

import "github.com/avila-r/pocketclient/validation"

type AdminProfile struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created"`
	UpdatedAt string `json:"updated"`
	Email     string `json:"email"`
	Avatar    int    `json:"avatar"`
}

type AdminRequest struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Avatar   int    `json:"avatar"`
}

func (m *ModuleAdmin) ListAll(p ...PaginationParams) (*Pagination[AdminProfile], error) {
	params := PaginationParams{}
	if len(p) > 0 {
		params = p[0]
	}

	res, err := RequestListAdmins(params)

	if err != nil {
		return nil, err
	}

	if err := validation.VerifyResponse(res); err != nil {
		return nil, err
	}

	admins := Pagination[AdminProfile]{}
	if err := Unmarshal(res.Body(), &admins); err != nil {
		return nil, err
	}

	return &admins, nil
}

func (m *ModuleAdmin) GetByID(id string) (*AdminProfile, error) {
	res, err := RequestGetAdmin(id)

	if err != nil {
		return nil, err
	}

	if err := validation.VerifyResponse(res); err != nil {
		return nil, err
	}

	admin := AdminProfile{}
	if err := Unmarshal(res.Body(), &admin); err != nil {
		return nil, err
	}

	return &admin, nil
}

func (m *ModuleAdmin) New(new AdminRequest) (*AdminProfile, error) {
	res, err := RequestPostAdmin(new)

	if err != nil {
		return nil, err
	}

	if err := validation.VerifyResponse(res); err != nil {
		return nil, err
	}

	created := AdminProfile{}
	if err := Unmarshal(res.Body(), &created); err != nil {
		return nil, err
	}

	return &created, nil
}

func (m *ModuleAdmin) Update(id string, patch AdminRequest) (*AdminProfile, error) {
	res, err := RequestUpdateAdmin(id, patch)

	if err != nil {
		return nil, err
	}

	if err := validation.VerifyResponse(res); err != nil {
		return nil, err
	}

	updated := AdminProfile{}
	if err := Unmarshal(res.Body(), &updated); err != nil {
		return nil, err
	}

	return &updated, nil
}

func (m *ModuleAdmin) DeleteByID(id string) error {
	res, err := RequestDeleteAdmin(id)

	if err != nil {
		return err
	}

	if err := validation.VerifyResponse(res); err != nil {
		return err
	}

	return nil
}
