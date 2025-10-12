override TARGET = spin

override RM = rm
override GO = go

.PHONY: all
all: $(TARGET)

$(TARGET):
	$(GO) build -o $(TARGET) main.go

.PHONY: run
run: $(TARGET)
	./$(TARGET)

.PHONY: re
re: clean all

.PHONY: test
test:
	true

.PHONY: clean
clean:
	$(RM) $(TARGET)
