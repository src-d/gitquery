package plan

import (
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gitql/gitql/mem"
	"github.com/gitql/gitql/sql"
	"github.com/gitql/gitql/sql/expression"
)

func TestProject(t *testing.T) {
	require := require.New(t)
	childSchema := sql.Schema{
		sql.Field{"col1", sql.String},
		sql.Field{"col2", sql.String},
	}
	child := mem.NewTable("test", childSchema)
	child.Insert("col1_1", "col2_1")
	child.Insert("col1_2", "col2_2")
	p := NewProject([]sql.Expression{expression.NewGetField(1, sql.String, "col2")}, child)
	require.Equal(1, len(p.Children()))
	schema := sql.Schema{
		sql.Field{"col2", sql.String},
	}
	require.Equal(schema, p.Schema())
	iter, err := p.RowIter()
	require.Nil(err)
	require.NotNil(iter)
	row, err := iter.Next()
	require.Nil(err)
	require.NotNil(row)
	fields := row.Fields()
	require.Equal(1, len(fields))
	require.Equal("col2_1", fields[0])
	row, err = iter.Next()
	require.Nil(err)
	require.NotNil(row)
	fields = row.Fields()
	require.Equal(1, len(fields))
	require.Equal("col2_2", fields[0])
	row, err = iter.Next()
	require.Equal(io.EOF, err)
	require.Nil(row)

	p = NewProject(nil, child)
	require.Equal(0, len(p.Schema()))

	p = NewProject([]sql.Expression{
		expression.NewAlias(
			expression.NewGetField(1, sql.String, "col2"),
			"foo",
		),
	}, child)
	schema = sql.Schema{
		sql.Field{"foo", sql.String},
	}
	require.Equal(schema, p.Schema())
}
