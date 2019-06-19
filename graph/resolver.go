package graph

import (
	"context"

	gen "github.com/gogrademe/api/graph/generated"
	"github.com/gogrademe/api/model"
	"github.com/gogrademe/api/store"
)

type Resolver struct {
	db *store.Store
}

func NewRootResolvers(db *store.Store) gen.Config {
	return gen.Config{
		Resolvers: &Resolver{
			db: db,
		},
	}
}

func (r *Resolver) Assignment() gen.AssignmentResolver {
	return &assignmentResolver{r}
}
func (r *Resolver) Mutation() gen.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() gen.QueryResolver {
	return &queryResolver{r}
}

type assignmentResolver struct{ *Resolver }

func (r *assignmentResolver) Course(ctx context.Context, obj *model.Assignment) (*model.Course, error) {
	panic("not implemented")
}
func (r *assignmentResolver) Term(ctx context.Context, obj *model.Assignment) (*model.Term, error) {
	panic("not implemented")
}
func (r *assignmentResolver) MaxScore(ctx context.Context, obj *model.Assignment) (int, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateAssignment(ctx context.Context, input model.NewAssignment) (*model.Assignment, error) {
	t := model.Assignment{Name: input.Name}
	return &t, r.db.InsertAssignment(&t)
}
func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	t := model.Course{Name: input.Name, LevelID: input.LevelID}
	return &t, r.db.InsertCourse(&t)
}
func (r *mutationResolver) CreateTerm(ctx context.Context, input model.NewTerm) (*model.Term, error) {
	t := model.Term{Name: input.Name, SchoolYear: input.SchoolYear}
	return &t, r.db.InsertTerm(&t)
}
func (r *mutationResolver) CreateAssignmentGroup(ctx context.Context, input model.NewAssignmentGroup) (*model.AssignmentGroup, error) {
	t := model.AssignmentGroup{Name: input.Name}
	return &t, r.db.InsertAssignmentGroup(&t)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Assignments(ctx context.Context) ([]*model.Assignment, error) {
	list, err := r.db.GetAssignmentList()
	if err != nil {
		return nil, err
	}

	res := []*model.Assignment{}
	for _, a := range list {
		res = append(res, &a)
	}
	return res, nil
}
func (r *queryResolver) Terms(ctx context.Context) ([]*model.Term, error) {
	list, err := r.db.GetTermList()
	if err != nil {
		return nil, err
	}

	res := []*model.Term{}
	for _, a := range list {
		res = append(res, &a)
	}
	return res, nil
}

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	list, err := r.db.GetCourseList()
	if err != nil {
		return nil, err
	}

	res := []*model.Course{}
	for _, a := range list {
		res = append(res, &a)
	}
	return res, nil
}

func (r *queryResolver) People(ctx context.Context) ([]*model.Person, error) {
	list, err := r.db.GetPersonList()
	if err != nil {
		return nil, err
	}

	res := []*model.Person{}
	for _, a := range list {
		res = append(res, &a)
	}
	return res, nil
}

type termResolver struct{ *Resolver }
