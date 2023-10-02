data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./loader",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  // dev = "docker://postgres"
  dev = "postgres://bb:bb@localhost/bb?sslmode=disable"
  migration {
    dir = "file://migrations/test/d1/ddl"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}


env "gorm-prod" {
  src = data.external_schema.gorm.url
  // dev = "docker://postgres"
  dev = "postgres://bb:bb@localhost/bb?sslmode=disable"
  migration {
    dir = "file://migrations/prod/d2/ddl"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
