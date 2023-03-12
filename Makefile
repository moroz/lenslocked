migrate:
	scripts/migrate_db.sh

rollback:
	scripts/migrate_db.sh down

recreate_db:
	dropdb --if-exists lenslocked_dev
	createdb lenslocked_dev
	make migrate

debug: FORCE
	dlv debug main.go

install_tools:
	scripts/install_tools.sh

FORCE: ;
