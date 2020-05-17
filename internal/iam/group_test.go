package iam

import (
	"context"
	"testing"

	"github.com/hashicorp/watchtower/internal/db"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func Test_NewGroup(t *testing.T) {
	t.Parallel()
	cleanup, conn, _ := db.TestSetup(t, "postgres")
	defer func() {
		if err := cleanup(); err != nil {
			t.Error(err)
		}
	}()
	assert := assert.New(t)
	defer conn.Close()

	t.Run("valid", func(t *testing.T) {
		w := db.New(conn)
		s, err := NewOrganization()
		assert.NoError(err)
		assert.NotNil(s.Scope != nil)
		err = w.Create(context.Background(), s)
		assert.NoError(err)
		assert.NotEmpty(s.PublicId)

		grp, err := NewGroup(s.PublicId, WithDescription("this is a test group"))
		assert.NoError(err)
		assert.NotNil(grp)
		assert.Equal(grp.Description, "this is a test group")
		assert.Equal(s.PublicId, grp.ScopeId)
		err = w.Create(context.Background(), grp)
		assert.NoError(err)
		assert.NotEmpty(grp.PublicId)
	})
	t.Run("no-scope", func(t *testing.T) {
		grp, err := NewGroup("")
		assert.Error(err)
		assert.Nil(grp)
		assert.Equal(err.Error(), "error the group scope id is unset")
	})
}

func TestGroup_Members(t *testing.T) {
	t.Parallel()
	cleanup, conn, _ := db.TestSetup(t, "postgres")
	defer func() {
		if err := cleanup(); err != nil {
			t.Error(err)
		}
	}()
	assert := assert.New(t)
	defer conn.Close()

	t.Run("valid", func(t *testing.T) {
		w := db.New(conn)
		s, err := NewOrganization()
		assert.NoError(err)
		assert.NotNil(s.Scope != nil)
		err = w.Create(context.Background(), s)
		assert.NoError(err)
		assert.NotEmpty(s.PublicId)

		user, err := NewUser(s.PublicId)
		assert.NoError(err)
		err = w.Create(context.Background(), user)
		assert.NoError(err)

		grp, err := NewGroup(s.PublicId, WithDescription("this is a test group"))
		assert.NoError(err)
		assert.NotNil(grp)
		assert.Equal(grp.Description, "this is a test group")
		assert.Equal(s.PublicId, grp.ScopeId)
		err = w.Create(context.Background(), grp)
		assert.NoError(err)
		assert.NotEmpty(grp.PublicId)

		gm, err := NewGroupMember(grp, user)
		assert.NoError(err)
		assert.NotNil(gm)
		err = w.Create(context.Background(), gm)
		assert.NoError(err)

		secondUser, err := NewUser(s.PublicId)
		assert.NoError(err)
		assert.NotNil(secondUser)
		err = w.Create(context.Background(), secondUser)
		assert.NoError(err)

		gm2, err := NewGroupMember(grp, secondUser)
		assert.NoError(err)
		assert.NotNil(gm2)
		err = w.Create(context.Background(), gm2)
		assert.NoError(err)

		members, err := grp.Members(context.Background(), w)
		assert.NoError(err)
		assert.NotNil(members)
		assert.Equal(2, len(members))
		for _, m := range members {
			if m.GetMemberId() != secondUser.PublicId && m.GetMemberId() != user.PublicId {
				t.Errorf("members %s not one of the known ids %s, %s", m.GetMemberId(), secondUser.PublicId, user.PublicId)
			}
		}
	})
}

func TestGroup_Actions(t *testing.T) {
	assert := assert.New(t)
	r := &Group{}
	a := r.Actions()
	assert.Equal(a[ActionCreate.String()], ActionCreate)
	assert.Equal(a[ActionUpdate.String()], ActionUpdate)
	assert.Equal(a[ActionRead.String()], ActionRead)
	assert.Equal(a[ActionDelete.String()], ActionDelete)
}

func TestGroup_ResourceType(t *testing.T) {
	assert := assert.New(t)
	r := &Group{}
	ty := r.ResourceType()
	assert.Equal(ty, ResourceTypeGroup)
}

func TestGroup_GetScope(t *testing.T) {
	t.Parallel()
	cleanup, conn, _ := db.TestSetup(t, "postgres")
	defer func() {
		if err := cleanup(); err != nil {
			t.Error(err)
		}
	}()
	assert := assert.New(t)
	defer conn.Close()

	t.Run("valid", func(t *testing.T) {
		w := db.New(conn)
		s, err := NewOrganization()
		assert.NoError(err)
		assert.NotNil(s.Scope != nil)
		err = w.Create(context.Background(), s)
		assert.NoError(err)
		assert.NotEmpty(s.PublicId)

		grp, err := NewGroup(s.PublicId)
		assert.NoError(err)
		assert.NotNil(grp)
		assert.Equal(s.PublicId, grp.ScopeId)
		err = w.Create(context.Background(), grp)
		assert.NoError(err)
		assert.NotEmpty(grp.PublicId)

		scope, err := grp.GetScope(context.Background(), w)
		assert.NoError(err)
		assert.NotNil(scope)
	})
}

func TestGroup_AddMember(t *testing.T) {
	t.Parallel()
	cleanup, conn, _ := db.TestSetup(t, "postgres")
	defer func() {
		if err := cleanup(); err != nil {
			t.Error(err)
		}
	}()
	assert := assert.New(t)
	defer conn.Close()

	t.Run("valid", func(t *testing.T) {
		w := db.New(conn)
		s, err := NewOrganization()
		assert.NoError(err)
		assert.NotNil(s.Scope != nil)
		err = w.Create(context.Background(), s)
		assert.NoError(err)
		assert.NotEmpty(s.PublicId)

		user, err := NewUser(s.PublicId)
		assert.NoError(err)
		err = w.Create(context.Background(), user)
		assert.NoError(err)

		grp, err := NewGroup(s.PublicId, WithDescription("this is a test group"))
		assert.NoError(err)
		assert.NotNil(grp)
		assert.Equal(grp.Description, "this is a test group")
		assert.Equal(s.PublicId, grp.ScopeId)
		err = w.Create(context.Background(), grp)
		assert.NoError(err)
		assert.NotEmpty(grp.PublicId)

		gm, err := grp.AddMember(context.Background(), w, user)
		assert.NoError(err)
		assert.NotNil(gm)
		assert.Equal(gm.(*GroupMemberUser).GroupId, grp.PublicId)
		err = w.Create(context.Background(), gm)
		assert.NoError(err)
		assert.Equal("user", gm.GetType())
	})
}

func TestGroup_Clone(t *testing.T) {
	t.Parallel()
	cleanup, conn, _ := db.TestSetup(t, "postgres")
	defer func() {
		if err := cleanup(); err != nil {
			t.Error(err)
		}
	}()
	assert := assert.New(t)
	defer conn.Close()

	t.Run("valid", func(t *testing.T) {
		w := db.New(conn)
		s, err := NewOrganization()
		assert.NoError(err)
		assert.NotNil(s.Scope != nil)
		err = w.Create(context.Background(), s)
		assert.NoError(err)
		assert.NotEmpty(s.PublicId)

		grp, err := NewGroup(s.PublicId, WithDescription("this is a test group"))
		assert.NoError(err)
		assert.NotNil(grp)
		assert.Equal(grp.Description, "this is a test group")
		assert.Equal(s.PublicId, grp.ScopeId)
		err = w.Create(context.Background(), grp)
		assert.NoError(err)
		assert.NotEmpty(grp.PublicId)

		cp := grp.Clone()
		assert.True(proto.Equal(cp.(*Group).Group, grp.Group))
	})
	t.Run("not-equal", func(t *testing.T) {
		w := db.New(conn)
		s, err := NewOrganization()
		assert.NoError(err)
		assert.NotNil(s.Scope != nil)
		err = w.Create(context.Background(), s)
		assert.NoError(err)
		assert.NotEmpty(s.PublicId)

		grp, err := NewGroup(s.PublicId, WithDescription("this is a test group"))
		assert.NoError(err)
		assert.NotNil(grp)
		assert.Equal(grp.Description, "this is a test group")
		assert.Equal(s.PublicId, grp.ScopeId)
		err = w.Create(context.Background(), grp)
		assert.NoError(err)
		assert.NotEmpty(grp.PublicId)

		grp2, err := NewGroup(s.PublicId, WithDescription("second group"))
		assert.NoError(err)
		assert.NotNil(grp2)
		err = w.Create(context.Background(), grp2)
		assert.NoError(err)
		assert.NotEmpty(grp2.PublicId)

		cp := grp.Clone()
		assert.True(!proto.Equal(cp.(*Group).Group, grp2.Group))
	})
}
