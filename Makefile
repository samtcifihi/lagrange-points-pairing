compile:
	go build -o "./bin/lagrange-points-pairing" ./src

run:
	./bin/lagrange-points-pairing

test: compile
	./bin/lagrange-points-pairing

clean:
	rm ./bin/lagrange-points-pairing
