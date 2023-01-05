package types

type SortType string

const (
	SortTypeActive       = "Active"
	SortTypeHot          = "Hot"
	SortTypeNew          = "New"
	SortTypeTopDay       = "TopDay"
	SortTypeTopWeek      = "TopWeek"
	SortTypeTopMonth     = "TopMonth"
	SortTypeTopYear      = "TopYear"
	SortTypeTopAll       = "TopAll"
	SortTypeMostComments = "MostComments"
	SortTypeNewComments  = "NewComments"
)

type ListingType string

const (
	ListingTypeAll        = "All"
	ListingTypeLocal      = "Local"
	ListingTypeSubscribed = "Subscribed"
	ListingTypeCommunity  = "Community"
)

type SearchType string

const (
	SearchTypeAll         = "All"
	SearchTypeComments    = "Comments"
	SearchTypePosts       = "Posts"
	SearchTypeCommunities = "Communities"
	SearchTypeUsers       = "Users"
	SearchTypeUrl         = "Url"
)
