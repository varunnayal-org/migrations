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
  dev = "postgres://d1:d1@localhost/d1?sslmode=disable"
  migration {
    dir = "file://migrations/atlas-dir"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
