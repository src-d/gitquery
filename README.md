[![Build Status](https://travis-ci.org/gitql/gitql.svg?branch=master)](https://travis-ci.org/gitql/gitql) [![codecov](https://codecov.io/gh/gitql/gitql/branch/master/graph/badge.svg)](https://codecov.io/gh/gitql/gitql) [![GoDoc](https://godoc.org/github.com/gitql/gitql?status.svg)](https://godoc.org/github.com/gitql/gitql)

<center> <img src="https://rawgit.com/gitql/gitql/master/gitql-logo.svg"></center>

**gitql** is a SQL interface to Git repositories, written in Go.

## Installation

Check the [Releases](https://github.com/gitql/gitql/releases) page to download
the gitql binary.

## Usage

```bash
Usage:
  gitql [OPTIONS] <query | shell | version>

Help Options:
  -h, --help  Show this help message

Available commands:
  query    Execute a SQL query a repository.
  shell    Start an interactive session.
  version  Show the version information.
```

For example:

```bash
$ cd my_git_repo
$ gitql query 'SELECT hash, author_email, author_name FROM commits LIMIT 2;' 
SELECT hash, author_email, author_name FROM commits LIMIT 2;
+------------------------------------------+---------------------+-----------------------+
|                   HASH                   |    AUTHOR EMAIL     |      AUTHOR NAME      |
+------------------------------------------+---------------------+-----------------------+
| 003dc36e0067b25333cb5d3a5ccc31fd028a1c83 | user1@test.io       | Santiago M. Mola      |
| 01ace9e4d144aaeb50eb630fed993375609bcf55 | user2@test.io       | Antonio Navarro Perez |
+------------------------------------------+---------------------+-----------------------+
```

Also you can use the interactive shell like you usually do to explore tables in postgreSQL per example:

```bash
$ gitql shell

           gitQL SHELL
           -----------
You must end your queries with ';'

!> SELECT hash, author_email, author_name FROM commits LIMIT 2;

--> Executing query: SELECT hash, author_email, author_name FROM commits LIMIT 2;

+------------------------------------------+---------------------+-----------------------+
|                   HASH                   |    AUTHOR EMAIL     |      AUTHOR NAME      |
+------------------------------------------+---------------------+-----------------------+
| 003dc36e0067b25333cb5d3a5ccc31fd028a1c83 | user1@test.io       | Santiago M. Mola      |
| 01ace9e4d144aaeb50eb630fed993375609bcf55 | user2@test.io       | Antonio Navarro Perez |
+------------------------------------------+---------------------+-----------------------+
!>  
```

## Tables

gitql exposes the following tables:

|     Name     |                                               Columns                                               |
|:------------:|:---------------------------------------------------------------------------------------------------:|
|    commits   | hash, author_name, author_email, author_time, comitter_name, comitter_email, comitter_time, message |
|     blobs    | hash, size                                                                                          |
|  references  | hash,hash, name, is_branch, is_note, is_remote, is_tag, target                                      |
|     tags     | hash, name, tagger_email, tagger_name, tagger_when, message, target                                 |
| tree_entries | tree_hash, entry_hash, mode, name                                                                   |

## SQL syntax

We are adding continuously more functionality to gitql. Actually we supports a subset of the SQL standard, currently including:

|                        |                                     Supported                                     |        Planned       |
|:----------------------:|:---------------------------------------------------------------------------------:|:--------------------:|
| Comparison expressions |                                !=, ==, >, <, >=,<=                                |                      |
|  Grouping expressions  |                                    COUNT, FIRST                                   |    LAST, SUM, AVG    |
|  Standard expressions  |                              ALIAS, LITERAL, STAR (*)                             |     LIKE, REGEXP     |
|       Statements       | CROSS JOIN, DESCRIBE, FILTER (WHERE), GROUP BY, LIMIT, SELECT, SHOW TABLES, SORT  | more JOINS, DISTINCT |

## License

gitql is licensed under the [MIT License](https://github.com/gitql/gitql/blob/master/LICENSE).
