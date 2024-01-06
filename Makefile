# Makefile

.PHONY: templ-watch air serve

# Command to start templ with the watch flag
templ-watch:
	templ generate --watch

# Command to start air for live-reloading of Go application
air:
	air

# Command to run both templ-watch and air concurrently
serve:
	make templ-watch & make air


#templ generate --watch --proxy="http://localhost:1323" --cmd="go run ."