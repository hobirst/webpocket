builddir = build

.prep:
	mkdir -p $(builddir)

build: .prep
	go build . -o build/webpocket

install: build
	mv build/webpocket /usr/bin/webpocket

clean:
	rm -rf build/
