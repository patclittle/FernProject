package model

type Petition struct {
	CreatedAt   DateTime `json:"createdAt,omitempty"`
	Description string   `json:"description,omitempty"`
	Categories  []string `json:"categories,omitempty"`
	Image       string   `json:"image,omitempty"`
	ID          string   `json:"id,omitempty"`
	RelatedTo   []string `json:"relatedTo,omitempty"`
	Upvotes     int      `json:"upvotes,omitempty"`
	Downvotes   int      `json:"downvotes,omitempty"`
	Comments    []string `json:"comments,omitempty"`
	NeedsInfo   int      `json:"needsInfo,omitempty"`
}
