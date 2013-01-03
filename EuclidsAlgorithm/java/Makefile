JUNIT_HOME = "/usr/share/java/junit4.jar"
test :
	@rm -f bin/main/*.class
	@echo Compiling...
	@javac -cp "bin:$(JUNIT_HOME)" -d bin src/main/*
	@java -cp "bin:$(JUNIT_HOME)" org.junit.runner.JUnitCore main.TestEuclid
