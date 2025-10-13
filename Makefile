override TARGET = spin

override RM = rm -f
override MAKEFLAGS = --no-print-directory
override GO = $$HOME/local/go/go/bin/go

.PHONY: all
all:
	@$(MAKE) clean > /dev/null || true
	$(GO) build -o $(TARGET) main.go

.PHONY: run
run: $(TARGET)
	./$(TARGET)

.PHONY: test
test:
	true

.PHONY: clean
clean:
	$(RM) $(TARGET)

.PHONY: fmt
fmt:
	$(GO) fmt ./...
