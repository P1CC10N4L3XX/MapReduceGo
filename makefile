# Definisci le directory dei programmi
PROGRAMS = mapper master reducer

# Definisci la directory di output
OUTPUT_DIR = builds

# Default: compila tutto
all: $(PROGRAMS)

# Regola generica per costruire ogni programma
$(PROGRAMS):
	@echo "Building $@..."
	@mkdir -p $(OUTPUT_DIR)
	@go build -o $(OUTPUT_DIR)/$@ ./src/$@
	@echo "$@ build complete!"

# Pulizia degli eseguibili
clean:
	@echo "Cleaning up..."
	@rm -rf $(OUTPUT_DIR)
	@echo "Clean complete!"
