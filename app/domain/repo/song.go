package repo

type Play interface {
	Save(play *Play) error
}

type Track interface {
	Save(track *Track) error
}
