package models

type DefaultData struct {
	UID        int    `json:"uid"`
}

type DefaultError struct {
	Message string
}

func (d *DefaultError) Error() string {
	return d.Message
}
