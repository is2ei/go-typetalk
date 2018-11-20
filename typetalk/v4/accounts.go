package v4

import (
	"context"
	"time"

	. "github.com/nulab/go-typetalk/typetalk/internal"
	. "github.com/nulab/go-typetalk/typetalk/shared"
)

type AccountsService service

type Account struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	FullName   string     `json:"fullName"`
	Suggestion string     `json:"suggestion"`
	ImageURL   string     `json:"imageUrl"`
	IsBot      bool       `json:"isBot"`
	CreatedAt  *time.Time `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}

type Status struct {
	Presence *string     `json:"presence"`
	Web      interface{} `json:"web"`
	Mobile   interface{} `json:"mobile"`
}

type AccountStatus struct {
	Account *Account `json:"account"`
	Status  *Status  `json:"status"`
}

type Friends struct {
	Count    int              `json:"count"`
	Accounts []*AccountStatus `json:"accounts"`
}

type GetMyFriendsOptions struct {
	Offset int `json:"offset,omitempty"`
	Count  int `json:"count,omitempty"`
}

type getMyFriendsOptions struct {
	*GetMyFriendsOptions
	SpaceKey string `json:"spaceKey"`
	Q        string `json:"q"`
}

// https://developer.nulab-inc.com/docs/typetalk/api/4/get-friends
func (s *AccountsService) GetMyFriends(ctx context.Context, spaceKey, q string, opt *GetMyFriendsOptions) (*Friends, *Response, error) {
	u, err := AddQueries("search/friends", &getMyFriendsOptions{GetMyFriendsOptions: opt, SpaceKey: spaceKey, Q: q})
	if err != nil {
		return nil, nil, err
	}
	var result *Friends
	if resp, err := s.client.Get(ctx, u, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}