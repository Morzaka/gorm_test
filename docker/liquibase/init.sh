#!/usr/bin/env bash
/opt/liquibase/liquibase \
    --driver=org.postgresql.Driver \
    --url=jdbc:postgresql://postgres:5432/users \
    --liquibaseSchemaName=liquibase \
    --classpath=/usr/share/java/postgresql.jar \
    --changeLogFile=/liquibase/my-changelog.xml \
    --username=postgres \
    --password=postgres \
    --contexts=all \
    update