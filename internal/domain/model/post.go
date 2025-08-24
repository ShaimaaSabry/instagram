package model

type Media struct {
	URL string
}

func NewMedia(url string) (Media, error) {
	return Media{
		URL: url,
	}, nil
}

func (m *Media) Url() string {
	return m.URL
}

type Post struct {
	id      int
	media   []Media
	caption string
}

func NewPost(
	media []Media,
	caption string,
) (Post, error) {
	return Post{
		id:      0,
		media:   media,
		caption: caption,
	}, nil
}

func (p *Post) Id() int {
	return p.id
}

func (p *Post) Media() []Media {
	return p.media
}

func (p *Post) Caption() string {
	return p.caption
}
