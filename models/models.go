package models

type BlogEntry struct {
	Title                   string   `json:"title"`
	Content                 string   `json:"content"`
	Id                      int      `json:"id"`
	OriginalLocale          string   `json:"originalLocale"`
	CreationTimeSeconds     int      `json:"creationTimeSeconds"`
	AuthorHandle            string   `json:"authorHandle"`
	Locale                  string   `json:"locale"`
	ModificationTimeSeconds int      `json:"modificationTimeSeconds"`
	AllowViewHistory        bool     `json:"allowViewHistory"`
	Tags                    []string `json:"tags"`
	Rating                  int      `json:"rating"`
}

type Comment struct {
	Text                string `json:"text"`
	Rating              int    `json:"rating"`
	Id                  int    `json:"id"`
	CreationTimeSeconds int    `json:"creationTimeSeconds"`
	CommentatorHandle   string `json:"commentatorHandle"`
	Locale              string `json:"locale"`
	ParentCommentId     int    `json:"parentCommentId"`
}

type RecentActions struct {
	TimeSeconds int       `json:"timeSeconds"`
	BlogEntry   BlogEntry `json:"blogEntry"`
	Comment     Comment   `json:"comment"`
}

type Wrapper struct {
	Status string          `json:"status"`
	Result []RecentActions `json:"result"`
}
