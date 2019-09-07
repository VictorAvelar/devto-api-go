# devto
--
    import "github.com/VictorAvelar/devto-api-go/devto"

Package devto is a wrapper around dev.to REST API

Where programmers share ideas and help each other grow. It is an online
community for sharing and discovering great ideas, having debates, and making
friends. Anyone can share articles, questions, discussions, etc. as long as they
have the rights to the words they are sharing. Cross-posting from your own blog
is welcome.

See: https://docs.dev.to/api

## Usage

```go
const (
	BaseURL      string = "https://dev.to"
	APIVersion   string = "0.5.1"
	APIKeyHeader string = "api-key"
)
```
Configuration constants

```go
var (
	ErrMissingConfig     = errors.New("missing configuration")
	ErrProtectedEndpoint = errors.New("to use this resource you need to provide an authentication method")
)
```
devto client errors

```go
var (
	ErrMissingRequiredParameter = errors.New("a required parameter is missing")
)
```
Configuration errors

#### type Article

```go
type Article struct {
	TypeOf                 string       `json:"type_of,omitempty"`
	ID                     uint32       `json:"id,omitempty"`
	Title                  string       `json:"title,omitempty"`
	Description            string       `json:"description,omitempty"`
	CoverImage             *WebURL      `json:"cover_image,omitempty"`
	SocialImage            *WebURL      `json:"social_image,omitempty"`
	PublishedAt            *time.Time   `json:"published_at,omitempty"`
	EditedAt               *time.Time   `json:"edited_at,omitempty"`
	CrossPostedAt          *time.Time   `json:"crossposted_at,omitempty"`
	LastCommentAt          *time.Time   `json:"last_comment_at,omitempty"`
	TagList                Tags         `json:"tag_list,omitempty"`
	Tags                   string       `json:"tags,omitempty"`
	Slug                   string       `json:"slug,omitempty"`
	Path                   *WebURL      `json:"path,omitempty"`
	URL                    *WebURL      `json:"url,omitempty"`
	CanonicalURL           *WebURL      `json:"canonical_url,omitempty"`
	CommentsCount          uint         `json:"comments_count,omitempty"`
	PositiveReactionsCount uint         `json:"positive_reactions_count,omitempty"`
	PublishedTimestamp     *time.Time   `json:"published_timestamp,omitempty"`
	User                   User         `json:"user,omitempty"`
	Organization           Organization `json:"organization,omitempty"`
	BodyHTML               string       `json:"body_html,omitempty"`
	BodyMarkdown           string       `json:"body_markdown,omitempty"`
	Published              bool         `json:"published,omitempty"`
}
```

Article contains all the information related to a single information resource
from devto.

#### type ArticleListOptions

```go
type ArticleListOptions struct {
	Tags     string `url:"tag,omitempty"`
	Username string `url:"username,omitempty"`
	State    string `url:"state,omitempty"`
	Top      string `url:"top,omitempty"`
	Page     int    `url:"page,omitempty"`
}
```

ArticleListOptions holds the query values to pass as query string parameter to
the Articles List action.

#### type ArticlesResource

```go
type ArticlesResource struct {
	API *Client
}
```

ArticlesResource implements the APIResource interface for devto articles.

#### func (*ArticlesResource) Find

```go
func (ar *ArticlesResource) Find(ctx context.Context, id uint32) (Article, error)
```
Find will retrieve an Article matching the ID passed.

#### func (*ArticlesResource) List

```go
func (ar *ArticlesResource) List(ctx context.Context, opt ArticleListOptions) ([]Article, error)
```
List will return the articles uploaded to devto, the result can be narrowed
down, filtered or enhanced using query parameters as specified on the
documentation. See: https://docs.dev.to/api/#tag/articles/paths/~1articles/get

#### func (*ArticlesResource) New

```go
func (ar *ArticlesResource) New(ctx context.Context, a Article) (Article, error)
```
New will create a new article on dev.to

#### func (*ArticlesResource) Update

```go
func (ar *ArticlesResource) Update(ctx context.Context, a Article) (Article, error)
```
Update will mutate the resource by id, and all the changes performed to the
Article will be applied, thus validation on the API side.

#### type Client

```go
type Client struct {
	Context    context.Context
	BaseURL    *url.URL
	HTTPClient httpClient
	Config     *Config
	Articles   *ArticlesResource
}
```

Client is the main data structure for performing actions against dev.to API

#### func  NewClient

```go
func NewClient(ctx context.Context, conf *Config, bc httpClient, bu string) (dev *Client, err error)
```
NewClient takes a context, a configuration pointer and optionally a base http
client (bc) to build an Client instance.

#### func (*Client) NewRequest

```go
func (c *Client) NewRequest(method string, uri string, body io.Reader) (*http.Request, error)
```
NewRequest build the request relative to the client BaseURL

#### type Config

```go
type Config struct {
	APIKey       string
	InsecureOnly bool
}
```

Config contains the elements required to initialize a devto client.

#### func  NewConfig

```go
func NewConfig(p bool, k string) (c *Config, err error)
```
NewConfig build a devto configuration instance with the required parameters.

It takes a boolean (p) as first parameter to indicate if you need access to
endpoints which require authentication, and a API key as second parameter, if p
is set to true and you don't provide an API key, it will return an error.

#### type Organization

```go
type Organization struct {
	Name           string  `json:"name,omitempty"`
	Username       string  `json:"username,omitempty"`
	Slug           string  `json:"slug,omitempty"`
	ProfileImage   *WebURL `json:"profile_image,omitempty"`
	ProfileImage90 *WebURL `json:"profile_image_90,omitempty"`
}
```

Organization describes a company or group that publishes content to devto.

#### type Tags

```go
type Tags []string
```

Tags are a group of topics related to an article

#### type User

```go
type User struct {
	Name            string  `json:"name,omitempty"`
	Username        string  `json:"username,omitempty"`
	TwitterUsername string  `json:"twitter_username,omitempty"`
	GithubUsername  string  `json:"github_username,omitempty"`
	WebsiteURL      *WebURL `json:"website_url,omitempty"`
	ProfileImage    *WebURL `json:"profile_image,omitempty"`
	ProfileImage90  *WebURL `json:"profile_image_90,omitempty"`
}
```

User contains information about a devto account

#### type WebURL

```go
type WebURL struct {
	*url.URL
}
```


#### func (*WebURL) UnmarshalJSON

```go
func (s *WebURL) UnmarshalJSON(b []byte) error
```
UnmarshalJSON overrides the default unmarshal behaviour for URL
