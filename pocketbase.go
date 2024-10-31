package pocketclient

type PocketBase struct {
	URL string

	Admins []*AdminCredentials
}

func (p *PocketBase) FirstAdmin() *AdminCredentials {
	if len(p.Admins) > 0 {
		return p.Admins[0]
	}

	return nil
}
