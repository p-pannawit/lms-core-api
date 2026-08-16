data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./internal/database/model",
    "--dialect", "postgres",
  ]
}

env "local" {
  src = data.external_schema.gorm.url

  url = getenv("DATABASE_URL")

  dev = "postgres://postgres:postgres@localhost:5432/salung_dev?sslmode=disable"

  migration {
    dir = "file://internal/database/migration"
  }
}
