# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

build-container:
	echo "[Makefile] >> Building container..."
	docker build --tag screenshot:latest .
	echo "[Makefile] >> Done!"

run: build-container
	echo "[Makefile] >> Running container..."
	docker container run --rm -p 8080:8080 screenshot:latest
