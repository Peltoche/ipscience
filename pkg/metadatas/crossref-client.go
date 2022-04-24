package metadatas

import (
	"context"
	"fmt"

	"github.com/cgxeiji/crossref"
)

const defaultMailto = "jeanne.do@ipscience.com"

type CrossRefClient struct {
	api *crossref.Client
}

func NewCrossRefClient() *CrossRefClient {
	api := crossref.NewClient("ipscience", defaultMailto)

	return &CrossRefClient{api}
}

func (c *CrossRefClient) FetchWork(ctx context.Context, doi string) (*Work, error) {
	work, err := c.api.DOI(doi)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from the remote: %w", err)
	}

	res := Work{
		Type:      work.Type,
		DOI:       doi,
		Title:     work.Title,
		BookTitle: work.BookTitle,
		Authors:   []*Contributor{},
		Editors:   []*Contributor{},
		Date:      work.Date,
		Publisher: work.Publisher,
		Volume:    work.Volume,
		Pages:     work.Pages,
		ISSNs:     work.ISSNs,
		ISSN:      work.ISSN,
		ISBNs:     work.ISBNs,
		ISBN:      work.ISSN,
	}

	for _, author := range work.Authors {
		res.Authors = append(res.Authors, &Contributor{
			First: author.First,
			Last:  author.Last,
		})
	}

	for _, editor := range work.Editors {
		res.Editors = append(res.Editors, &Contributor{
			First: editor.First,
			Last:  editor.Last,
		})
	}

	return &res, nil
}
