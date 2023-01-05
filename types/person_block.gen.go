package types

type PersonBlock struct {
	ID        int       `json:"id" url:"id,omitempty"`
	PersonID  int       `json:"person_id" url:"person_id,omitempty"`
	TargetID  int       `json:"target_id" url:"target_id,omitempty"`
	Published LemmyTime `json:"published" url:"published,omitempty"`
}
type PersonBlockForm struct {
	PersonID int `json:"person_id" url:"person_id,omitempty"`
	TargetID int `json:"target_id" url:"target_id,omitempty"`
}
