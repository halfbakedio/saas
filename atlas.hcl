env "local" {
  src = "ent://ent/schema"
  url = "postgres://postgres:postgres@:5432/saas?search_path=public&sslmode=disable"
  dev = "docker://postgres/16/dev"

  migration {
    dir = "file://ent/migrate/migrations"
    # format = golang-migrate
  }

  diff {
    skip {
      drop_table = true
      drop_schema = true
    }
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
