JUNIT_HOME = "/usr/share/java/junit4.jar"
test :
	@rm -f bin/euclid/*.class
	@echo Compiling...
	@javac -cp "bin:$(JUNIT_HOME)" -d bin src/euclid/*
	@java -cp "bin:$(JUNIT_HOME)" org.junit.runner.JUnitCore euclid.TestEuclid
