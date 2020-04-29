.DEFAULT_GOAL := all

go-fmt:
	@files=$$(gofmt -l .); \
        if [ -n "$$files" ]; then \
                echo "Go fmt issues $$files"; \
                exit 1; \
	else \
		echo "<<<<<<<< Go fmt is okay >>>>>>>>"; \
        fi;

