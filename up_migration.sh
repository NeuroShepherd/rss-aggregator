#!/bin/bash

source .env
cd sql/schema
goose postgres $POSTGRES_DB up