variable "destructive" {
  type    = bool
  default = false
}

env "local" {
  src = "ent://ent/schema"
  url = "postgres://postgres:postgres@:5432/saas?search_path=public&sslmode=disable"
  dev = "docker://postgres/16/dev"

  migration {
    dir = "file://ent/migrate/migrations"
  }

  diff {
    skip {
      drop_schema = !var.destructive
      drop_table  = !var.destructive
    }
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

env "test" {
  src = "ent://ent/schema"
  url = "postgres://postgres:postgres@:5432/saas_test?search_path=public&sslmode=disable"
  dev = "docker://postgres/16/test"

  migration {
    dir = "file://ent/migrate/migrations"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
